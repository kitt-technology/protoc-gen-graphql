package enum

import (
	"bytes"
	"fmt"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/types/descriptorpb"
	"html/template"
	"strings"
)

const msgTpl = `

var {{ .Descriptor.GetName }}_enum = graphql.NewEnum(graphql.EnumConfig{
	Name: "{{ .Descriptor.GetName }}",
	Values: graphql.EnumValueConfigMap{
		{{- range $key, $val := .Values }}
		"{{ $key }}": &graphql.EnumValueConfig{
			Value: {{ $.Descriptor.GetName }}({{ $val }}),
		},
		{{- end }}
	},
})

var {{ .Descriptor.GetName }}_type = graphql.NewScalar(graphql.ScalarConfig{
	Name: "{{ .Descriptor.GetName }}",
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

type Message struct {
	Descriptor *descriptorpb.EnumDescriptorProto
	Options    *graphql.MutationOption
	Import     map[string]string
	Values     map[string]string
}

func New(msg *descriptorpb.EnumDescriptorProto) (m Message) {
	return Message{
		Import:     make(map[string]string),
		Values:     make(map[string]string),
		Descriptor: msg,
	}
}

func last(path string) string {
	t := strings.Split(path, ".")
	return t[len(t)-1]
}

func (m Message) Imports() []string {
	return []string{"github.com/graphql-go/graphql/language/ast"}
}

func (m Message) Generate() string {
	for _, field := range m.Descriptor.Value {
		m.Values[*field.Name] = fmt.Sprint(*field.Number)
	}
	var buf bytes.Buffer
	mTpl, err := template.New("msg").Parse(msgTpl)
	if err != nil {
		panic(err)
	}
	mTpl.Execute(&buf, m)

	return buf.String()
}
