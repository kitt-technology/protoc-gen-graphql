package generation

import (
	"bytes"
	"github.com/kitt-technology/protoc-gen-graphql/generation/mutation"
	"github.com/kitt-technology/protoc-gen-graphql/generation/query"
	"github.com/kitt-technology/protoc-gen-graphql/generation/typedef"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"text/template"
)

const fileTpl = `
package {{ .Package }}

import "github.com/graphql-go/graphql"
import pg "github.com/kitt-technology/protoc-gen-graphql/graphql"
import "google.golang.org/protobuf/proto"

{{- range $import := .Imports }}
import "{{ $import }}"
{{ end }}

var mutations []*graphql.Field
var queries []*graphql.Field

var mutationResolver func(command proto.Message, success proto.Message) (proto.Message, error)
var dataloadersToRegister map[string][]pg.RegisterDataloaderFn
var dataloadersToProvide map[string]pg.Dataloader

func Register(config pg.ProtoConfig) pg.ProtoConfig {
	config.Queries = append(config.Queries, queries...)
	return config
}

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
			f.Imports = append(f.Imports, "github.com/golang/protobuf/ptypes/wrappers")
		}
	}

	for _, service := range file.Proto.Service {
		f.Message = append(f.Message, query.New(service, file.Proto))
	}
	for _, msg := range file.Proto.MessageType {
		if msg.Options != nil {
			if proto.HasExtension(msg.Options, graphql.E_MutationOptions) {
				f.Message = append(f.Message, mutation.New(msg))
			}
		}
		f.TypeDefs = append(f.TypeDefs, typedef.New(msg))

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
