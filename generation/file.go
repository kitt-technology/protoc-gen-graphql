package generation

import (
	"bytes"
	"github.com/kitt-technology/protoc-gen-graphql/generation/dataloaders"
	"github.com/kitt-technology/protoc-gen-graphql/generation/imports"
	"github.com/kitt-technology/protoc-gen-graphql/generation/types"
	"github.com/kitt-technology/protoc-gen-graphql/generation/types/enum"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"sort"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"
)

const fileTpl = `
package {{ .Package }}

import (
	gql "github.com/graphql-go/graphql"
	"google.golang.org/grpc"
	{{- range .Imports }}
	{{ if has_alias . }}{{ . }}{{else}}"{{ . }}"{{end}}{{ end }}
)

var fieldInits []func(...grpc.DialOption)

func Fields(opts ...grpc.DialOption) []*gql.Field {
	for _, fieldInit := range fieldInits {
		fieldInit(opts...)
	}
	return fields
}

var fields []*gql.Field
`

var GraphqlImportMap = make(map[string]types.GraphqlImport, 0)

type Message interface {
	Generate() string
	Imports() []string
}

type File struct {
	Package   protogen.GoPackageName
	Message   []Message
	TypeDefs  []Message
	Imports   []string
	ImportMap map[string]string
}

func New(file *protogen.File) (f File) {
	f.Package = file.GoPackageName

	for _, service := range file.Proto.Service {
		f.Message = append(f.Message, dataloaders.New(service, file.Proto))
		f.Imports = append(f.Imports, imports.PggImport)
	}

	for _, e := range file.Proto.EnumType {
		f.TypeDefs = append(f.TypeDefs, enum.New(e))
	}

	if proto.HasExtension(file.Proto.Options, graphql.E_Package) {
		importPath, gqlPkg, ok := parseGraphqlPackage(file.Proto)
		if !ok {
			panic("invalid graphql.package: " + file.Proto.GetName())
		}
		// Using graphql.package, could fall back to go_package?
		GraphqlImportMap[*file.Proto.Package] = types.GraphqlImport{
			ImportPath: importPath,
			GoPackage:  gqlPkg,
		}
	}

	for _, msg := range file.Proto.MessageType {
		if proto.HasExtension(msg.Options, graphql.E_SkipMessage) &&
			proto.GetExtension(msg.Options, graphql.E_SkipMessage).(bool) {
			continue
		}
		f.TypeDefs = append(f.TypeDefs, types.New(msg, file.Proto, GraphqlImportMap))

	}
	return f
}

func (f File) ToString() string {
	var extraImportMap = map[string]string{}
	var extraImports = []string{}

	for _, msg := range append(f.TypeDefs, f.Message...) {
		for _, imp := range msg.Imports() {
			extraImportMap[imp] = imp
		}
	}
	for _, val := range f.Imports {
		extraImportMap[val] = val
	}

	for _, val := range extraImportMap {
		extraImports = append(extraImports, val)
	}

	f.Imports = extraImports

	// Sort so that we're deterministic for testing
	sort.Strings(f.Imports)

	var buf bytes.Buffer
	tpl, err := template.New("file").Funcs(map[string]interface{}{
		"has_alias": func(impt string) bool {
			return strings.Contains(impt, "\"")
		},
	}).Parse(fileTpl)
	if err != nil {
		panic(err)
	}

	tpl.Execute(&buf, f)

	out := buf.String()

	for _, msg := range f.TypeDefs {
		out += msg.Generate()
	}

	for _, msg := range f.Message {
		out += msg.Generate()
	}
	return out
}

// The parsing for graphql.package is copied from protoc-gen-go's package parsing
// https://github.com/golang/protobuf/blob/ae97035608a719c7a1c1c41bed0ae0744bdb0c6f/protoc-gen-go/generator/generator.go#L275

func parseGraphqlPackage(file *descriptorpb.FileDescriptorProto) (importPath string, pkg string, ok bool) {
	if !proto.HasExtension(file.Options, graphql.E_Package) {
		return "", "", false
	}
	graphqlPackage := proto.GetExtension(file.Options, graphql.E_Package).(string)

	sc := strings.Index(graphqlPackage, ";")
	if sc >= 0 {
		return graphqlPackage[:sc], cleanPackageName(graphqlPackage[sc+1:]), true
	}
	slash := strings.LastIndex(graphqlPackage, "/")
	if slash >= 0 {
		return graphqlPackage, cleanPackageName(graphqlPackage[slash+1:]), true
	}
	return "", cleanPackageName(graphqlPackage), true
}

func cleanPackageName(name string) string {
	name = strings.Map(badToUnderscore, name)
	// Identifier must not be keyword or predeclared identifier: insert _.
	if isGoKeyword[name] {
		name = "_" + name
	}
	// Identifier must not begin with digit: insert _.
	if r, _ := utf8.DecodeRuneInString(name); unicode.IsDigit(r) {
		name = "_" + name
	}
	return name
}

var isGoKeyword = map[string]bool{
	"break":       true,
	"case":        true,
	"chan":        true,
	"const":       true,
	"continue":    true,
	"default":     true,
	"else":        true,
	"defer":       true,
	"fallthrough": true,
	"for":         true,
	"func":        true,
	"go":          true,
	"goto":        true,
	"if":          true,
	"import":      true,
	"interface":   true,
	"map":         true,
	"package":     true,
	"range":       true,
	"return":      true,
	"select":      true,
	"struct":      true,
	"switch":      true,
	"type":        true,
	"var":         true,
}

// badToUnderscore is the mapping function used to generate Go names from package names,
// which can be dotted in the input .proto file.  It replaces non-identifier characters such as
// dot or dash with underscore.
func badToUnderscore(r rune) rune {
	if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
		return r
	}
	return '_'
}
