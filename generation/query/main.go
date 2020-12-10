package query

import (
	"bytes"
	"github.com/iancoleman/strcase"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"html/template"
	"strings"
)

const msgTpl = `

var client {{ .Descriptor.Name }}Client

func init() {
	client = New{{ .Descriptor.Name }}Client(pg.GrpcConnection("{{ .Dns }}"))
	{{- range $method := .Methods }}
	Fields = append(Fields, &graphql.Field{
		Name: "{{ $.ServiceName }}_{{ $method.Name }}",
		Type: {{ $method.Output }}_type,
		Args: {{ $method.Input }}_args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return client.{{ $method.Name }}(p.Context, {{ $method.Input }}_from_args(p.Args))
		},
	})
	{{ end }}
}


{{ range $loader :=.Loaders }}
func Load{{ $loader.ResultsType }}(originalContext context.Context, key string) (func() (interface{}, error), error) {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		resp, err := client.{{ $loader.Method }}(ctx, &{{ $loader.RequestType }}{
			{{ $loader.KeysField }}: keys.Keys(),
		})

		if err != nil {
			return results
		}

		for _, item := range resp.{{ $loader.ResultsField }} {
			results = append(results, &dataloader.Result{Data: item})
		}

		return results
	}

	loader := dataloader.NewBatchedLoader(batchFn)

	thunk := loader.Load(originalContext, dataloader.StringKey(key))
	return func() (interface{}, error) {
				res, err := thunk()
				if err != nil {
					return nil, err
				}
				return res.(*{{ $loader.ResultsType }}), nil
	}, nil
}

func LoadMany{{ $loader.ResultsType }}(originalContext context.Context, keys []string) (func() (interface{}, error), error) {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		resp, err := client.{{ $loader.Method }}(ctx, &{{ $loader.RequestType }}{
			{{ $loader.KeysField }}: keys.Keys(),
		})

		if err != nil {
			return results
		}

		for _, item := range resp.{{ $loader.ResultsField }} {
			results = append(results, &dataloader.Result{Data: item})
		}

		return results
	}

	loader := dataloader.NewBatchedLoader(batchFn)

	thunk := loader.LoadMany(originalContext, dataloader.NewKeysFromStrings(keys))
	return func() (interface{}, error) {
		resSlice, errSlice := thunk()
		
		for _, err := range errSlice {
			if err != nil {
				return nil, err
			}
		}
		
		var results []*{{ $loader.ResultsType }}
		for _, res := range resSlice {
			results = append(results, res.(*{{ $loader.ResultsType }}))
		}

		return results, nil
	}, nil
}
{{ end }}
`

type Loader struct {
	Method       string
	RequestType  string
	KeysField    string
	ResultsField string
	ResultsType  string
}

type Message struct {
	Descriptor  *descriptorpb.ServiceDescriptorProto
	Options     *graphql.MutationOption
	Methods     []Method
	ServiceName string
	Dns         string
	Loaders     []Loader
}

func New(msg *descriptorpb.ServiceDescriptorProto, root *descriptorpb.FileDescriptorProto) (m Message) {
	var methods []Method

	dns := proto.GetExtension(msg.Options, graphql.E_Host).(string)

	for _, method := range msg.Method {
		// Get output type of method
		var output *descriptorpb.DescriptorProto
		for _, msgType := range root.MessageType {
			if last(*method.OutputType) == *msgType.Name {
				output = msgType
			}
		}

		// See if method is a batch loader
		if proto.HasExtension(method.Options, graphql.E_Batch) {
			batchOptions := proto.GetExtension(method.Options, graphql.E_Batch).(*graphql.BatchOptions)

			// Find result field type
			var resultType string
			for _, field := range output.Field {
				if *field.Name == batchOptions.Results {
					resultType = strings.Title(last(*field.TypeName))
				}
			}

			m.Loaders = append(m.Loaders, Loader{
				Method:       strings.Title(*method.Name),
				RequestType:  strings.Title(last(*method.InputType)),
				KeysField:    strcase.ToCamel(batchOptions.Keys),
				ResultsField: strcase.ToCamel(batchOptions.Results),
				ResultsType:  resultType,
			})

		}

		methods = append(methods, Method{Input: last(*method.InputType), Output: last(*method.OutputType), Name: strings.Title(*method.Name)})
	}
	return Message{
		Descriptor:  msg,
		Methods:     methods,
		ServiceName: *msg.Name,
		Dns:         dns,
		Loaders:     m.Loaders,
	}
}

func (m Message) Imports() []string {
	if len(m.Loaders) > 0 {
		return []string{"context", "github.com/graph-gophers/dataloader"}
	}
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

type Method struct {
	Input  string
	Output string
	Name   string
}

func last(path string) string {
	t := strings.Split(path, ".")
	return t[len(t)-1]
}
