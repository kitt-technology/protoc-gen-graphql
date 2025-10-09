package generation

import (
	"bytes"
	"sort"
	"strings"
	"text/template"

	"github.com/kitt-technology/protoc-gen-graphql/generation/imports"
	"github.com/kitt-technology/protoc-gen-graphql/generation/templates"
	"github.com/kitt-technology/protoc-gen-graphql/generation/types"
	"github.com/kitt-technology/protoc-gen-graphql/generation/types/enum"
	"github.com/kitt-technology/protoc-gen-graphql/generation/util"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

const fileTpl = `
package {{ .Package }}

import (
	gql "github.com/graphql-go/graphql"
	{{- range .Imports }}
	{{ if has_alias . }}{{ . }}{{else}}"{{ . }}"{{end}}{{ end }}
)
`

var GraphqlImportMap = make(map[string]types.GraphqlImport, 0)
var SkippedMessages = make(map[string]bool, 0) // Tracks skipped messages by fully qualified name (e.g., "common.DateTime")

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
		f.Message = append(f.Message, templates.New(service, file.Proto))
		f.Imports = append(f.Imports, imports.PggImport)
	}

	for _, e := range file.Proto.EnumType {
		f.TypeDefs = append(f.TypeDefs, enum.New(e))
	}

	if proto.HasExtension(file.Proto.Options, graphql.E_Package) {
		importPath, gqlPkg, ok := util.ParseGraphqlPackage(file.Proto)
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
		if proto.HasExtension(msg.Options, graphql.E_SkipMessage) {
			if skip, ok := proto.GetExtension(msg.Options, graphql.E_SkipMessage).(bool); ok && skip {
				// Track skipped messages so other packages can skip fields referencing them
				if file.Proto.Package != nil {
					fullyQualifiedName := *file.Proto.Package + "." + *msg.Name
					SkippedMessages[fullyQualifiedName] = true
				}
				continue
			}
		}
		f.TypeDefs = append(f.TypeDefs, types.New(msg, file.Proto, GraphqlImportMap, SkippedMessages))

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

	if err := tpl.Execute(&buf, f); err != nil {
		panic(err)
	}

	out := buf.String()

	for _, msg := range f.TypeDefs {
		out += msg.Generate()
	}

	for _, msg := range f.Message {
		out += msg.Generate()
	}
	return out
}
