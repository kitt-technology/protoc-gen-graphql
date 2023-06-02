package enum

import (
	"bytes"
	"fmt"
	"github.com/kitt-technology/protoc-gen-graphql/generation/imports"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"html/template"
	"os"
)

type Message struct {
	Descriptor   *descriptorpb.EnumDescriptorProto
	Import       map[string]string
	Values       map[string]string
	DisplayNames map[string]string
	EnumName     string
}

func New(msg *descriptorpb.EnumDescriptorProto) (m Message) {
	if proto.HasExtension(msg.Options, graphql.E_EnumName) {
		m.EnumName = proto.GetExtension(msg.Options, graphql.E_EnumName).(string)
	} else {
		m.EnumName = *msg.Name
	}

	m = Message{
		Import:       make(map[string]string),
		Values:       make(map[string]string),
		DisplayNames: make(map[string]string),
		Descriptor:   msg,
		EnumName:     m.EnumName,
	}

	for _, value := range msg.Value {
		displayName := proto.GetExtension(value.Options, graphql.E_DisplayName)
		if displayName == "" {
			m.DisplayNames[*value.Name] = *value.Name
		} else {
			m.DisplayNames[*value.Name] = displayName.(string)
		}

	}

	return m
}

func (m Message) Imports() []string {
	return []string{imports.GraphqlAst}
}

func (m Message) Generate() string {
	for _, field := range m.Descriptor.Value {
		m.Values[*field.Name] = fmt.Sprint(*field.Number)
	}

	fmt.Fprintf(os.Stderr, "Display names: %s\n", m.DisplayNames)
	var buf bytes.Buffer
	mTpl, err := template.New("msg").Parse(msgTpl)
	if err != nil {
		panic(err)
	}
	mTpl.Execute(&buf, m)

	return buf.String()
}

const msgTpl = `
var {{ .Descriptor.GetName }}GraphqlEnum = gql.NewEnum(gql.EnumConfig{
	Name: "{{ .EnumName }}",
	Values: gql.EnumValueConfigMap{
		{{- range $key, $val := .Values }}
		"{{ $key }}": &gql.EnumValueConfig{
			Value: {{ $.Descriptor.GetName }}({{ $val }}),
		},
		{{- end }}
	},
})

var {{ .Descriptor.GetName }}GraphqlType = gql.NewScalar(gql.ScalarConfig{
	Name: "{{ .EnumName }}",
	ParseValue: func(value interface{}) interface{} {
		return nil
	},
	Serialize: func(value interface{}) interface{} {
		return value.({{ .Descriptor.GetName }}).String()
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		return nil
	},
})
`
