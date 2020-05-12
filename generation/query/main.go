package query

import (
	"bytes"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/types/descriptorpb"
	"html/template"
	"strings"
)

const msgTpl = `

var {{ .Descriptor.Name }} {{ .Descriptor.Name }}Client

func init() {
	{{ .Descriptor.Name }} = New{{ .Descriptor.Name }}Client(pg.GrpcConnection("service"))
	{{- range $method := .Methods }}
	queries = append(queries, &graphql.Field{
		Type: {{ $method.Output }}_type,
		Args: {{ $method.Input }}_args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return {{ $.ServiceName }}.{{ $method.Name }}(p.Context, {{ $method.Input }}_from_args(p.Args))
		},
	})
	{{ end }}	
}

`

type Message struct {
	Descriptor  *descriptorpb.ServiceDescriptorProto
	Options     *graphql.MutationOption
	Methods     []Method
	ServiceName string
}

func New(msg *descriptorpb.ServiceDescriptorProto) (m Message) {
	var methods []Method
	for _, method := range msg.Method {
		methods = append(methods, Method{Input: last(*method.InputType), Output: last(*method.OutputType), Name: strings.Title(*method.Name)})
	}
	return Message{
		Descriptor:  msg,
		Methods:     methods,
		ServiceName: *msg.Name,
	}
}

func (m Message) Generate() string {
	var buf bytes.Buffer
	mTpl, err := template.New("msg").Parse(msgTpl)
	if err != nil {
		panic(err)
	}
	mTpl.Execute(&buf, m)

	return buf.String()
}

type Method struct {
	Input  string
	Output string
	Name   string
}

func last(path string) string {
	t := strings.Split(path, ".")
	return t[len(t)-1]
}
