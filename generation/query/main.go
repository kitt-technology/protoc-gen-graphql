package query

import (
	"bytes"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"html/template"
	"strings"
	"unicode"
)

const msgTpl = `

var {{ lower .Descriptor.Name }}ClientInstance {{ .Descriptor.Name }}Client

func init() {
	fieldInits = append(fieldInits, func(opts ...grpc.DialOption) {
		{{ lower .Descriptor.Name }}ClientInstance = New{{ .Descriptor.Name }}Client(pg.GrpcConnection("{{ .Dns }}", opts...))
	})
	
	{{- range $method := .Methods }}
	fields = append(fields, &graphql.Field{
		Name: "{{ $.ServiceName }}_{{ $method.Name }}",
		Type: {{ $method.Output }}_type,
		Args: {{ $method.Input }}_args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return {{ lower $.Descriptor.Name }}ClientInstance.{{ $method.Name }}(p.Context, {{ $method.Input }}_from_args(p.Args))
		},
	})
	{{ end }}
}


func WithLoaders(ctx context.Context) context.Context {
	{{- range $loader :=.Loaders }}
	ctx = context.WithValue(ctx, "{{ $loader.ResultsType }}Loader", dataloader.NewBatchedLoader(
		func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
			var results []*dataloader.Result
	
			resp, err := {{ lower $.Descriptor.Name }}ClientInstance.{{ $loader.Method }}(ctx, &pg.BatchRequest{
				{{ $loader.KeysField }}: keys.Keys(),
			})
	
			if err != nil {
				return results
			}
	
			for _, key := range keys.Keys() {
				results = append(results, &dataloader.Result{Data: resp.{{ $loader.ResultsField }}[key]})
			}
	
			return results
		},
	))
	{{ end }}
	return ctx
}

{{ range $loader :=.Loaders }}
func Load{{ $loader.ResultsType }}(p graphql.ResolveParams, key string) (func() (interface{}, error), error) {
	loader := p.Context.Value("{{ $loader.ResultsType }}Loader").(*dataloader.Loader)
	thunk := loader.Load(p.Context, dataloader.StringKey(key))
	return func() (interface{}, error) {
				res, err := thunk()
				if err != nil {
					return nil, err
				}
				return res.(*{{ $loader.ResultsType }}), nil
	}, nil
}

func LoadMany{{ $loader.ResultsType }}(p graphql.ResolveParams, keys []string) (func() (interface{}, error), error) {
	loader := p.Context.Value("{{ $loader.ResultsType }}Loader").(*dataloader.Loader)
	thunk := loader.LoadMany(p.Context, dataloader.NewKeysFromStrings(keys))
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

		if last(*method.OutputType) == "Empty" {
			continue
		}

		// See if method is a batch loader
		if *method.InputType == ".graphql.BatchRequest" {
			// Find type of map
			var resultType string

			if len(output.Field) == 0 || output.Field[0].Label.String() != "LABEL_REPEATED" || !strings.Contains(*output.Field[0].TypeName, "Entry") {
				panic(fmt.Sprintf("batch loaders must have one field of the type: map<string, Result> for %s.%s", *msg.Name, *method.Name))
			}

			var field = output.Field[0]

			resultType = strings.Title(last(*field.TypeName))
			nestedTypeKey := last(*field.TypeName)
			for _, nestedType := range output.NestedType {
				if *nestedType.Name == nestedTypeKey {
					if nestedType.Field[1].TypeName != nil {
						resultType = last(*nestedType.Field[1].TypeName)
					} else {
						resultType = nestedType.Field[1].Type.String()
					}
				}
			}

			m.Loaders = append(m.Loaders, Loader{
				Method:       strings.Title(*method.Name),
				RequestType:  strings.Title(last(*method.InputType)),
				KeysField:    strcase.ToCamel("Keys"),
				ResultsField: strcase.ToCamel(*field.Name),
				ResultsType:  resultType,
			})

		} else {
			methods = append(methods, Method{Input: last(*method.InputType), Output: last(*method.OutputType), Name: strings.Title(*method.Name)})
		}
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
	mTpl, err := template.New("msg").Funcs(map[string]interface{}{
		"lower": func(input string) string {
			for i, v := range input {
				return string(unicode.ToLower(v)) + input[i+1:]
			}
			return ""
		},
	}).Parse(msgTpl)
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
