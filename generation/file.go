package generation

import (
	"bytes"
	"github.com/kitt-technology/protoc-gen-graphql/generation/dataloaders"
	"github.com/kitt-technology/protoc-gen-graphql/generation/types"
	"github.com/kitt-technology/protoc-gen-graphql/generation/types/enum"
	"google.golang.org/protobuf/compiler/protogen"
	"sort"
	"strings"
	"text/template"
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

type Message interface {
	Generate() string
	Imports() []string
}

type File struct {
	Package  protogen.GoPackageName
	Message  []Message
	TypeDefs []Message
	Imports  []string
}

func New(file *protogen.File) (f File) {
	f.Package = file.GoPackageName

	for _, service := range file.Proto.Service {
		f.Message = append(f.Message, dataloaders.New(service, file.Proto))
		f.Imports = append(f.Imports, "pg \"github.com/kitt-technology/protoc-gen-graphql/graphql\"")
	}

	for _, e := range file.Proto.EnumType {
		f.TypeDefs = append(f.TypeDefs, enum.New(e))
	}

	for _, msg := range file.Proto.MessageType {
		f.TypeDefs = append(f.TypeDefs, types.New(msg, file.Proto))

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
	for _, val := range extraImportMap {
		extraImports = append(extraImports, val)
	}

	f.Imports = append(f.Imports, extraImports...)

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
