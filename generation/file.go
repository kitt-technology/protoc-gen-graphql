package generation

import (
	"bytes"
	"github.com/kitt-technology/protoc-gen-graphql/generation/enum"
	"github.com/kitt-technology/protoc-gen-graphql/generation/query"
	"github.com/kitt-technology/protoc-gen-graphql/generation/typedef"
	"google.golang.org/protobuf/compiler/protogen"
	"text/template"
)

const fileTpl = `
package {{ .Package }}

import (
	"github.com/graphql-go/graphql"
	pg "github.com/kitt-technology/protoc-gen-graphql/graphql"
	{{ range .Imports }}
	"{{ . }}"
	{{end}}
)


var Fields []*graphql.Field
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

	for _, dep := range file.Proto.Dependency {
		switch dep {
		case "google/protobuf/wrappers.proto":
			//f.Imports = append(f.Imports, "github.com/golang/protobuf/ptypes/wrappers")
		}
	}

	for _, service := range file.Proto.Service {
		f.Message = append(f.Message, query.New(service, file.Proto))
	}

	for _, e := range file.Proto.EnumType {
		f.TypeDefs = append(f.TypeDefs, enum.New(e))
	}

	for _, msg := range file.Proto.MessageType {
		f.TypeDefs = append(f.TypeDefs, typedef.New(msg, file.Proto))

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

	var buf bytes.Buffer
	tpl, err := template.New("file").Parse(fileTpl)
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
