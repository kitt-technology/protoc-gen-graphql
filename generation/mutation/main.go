package mutation

import (
	"bytes"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"html/template"
)

const msgTpl = `
func init() {
	mutations = append(mutations, &graphql.Field{
		Type: {{ .Options.Success }}_type,
		Args: {{ .Descriptor.Name }}_args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return mutationResolver({{ .Descriptor.Name }}_from_args(p.Args), &TestCommand{})
		},
	})
}
`

type Message struct {
	Descriptor *descriptorpb.DescriptorProto
	Options    *graphql.MutationOption
}

func New(msg *descriptorpb.DescriptorProto) (m Message) {
	return Message{
		Options:    proto.GetExtension(msg.Options, graphql.E_MutationOptions).(*graphql.MutationOption),
		Descriptor: msg,
	}
}

func (m Message) Imports() []string {
	return []string{}
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

type Field struct {
	Optional bool
	Key      string
	Type     string
}
