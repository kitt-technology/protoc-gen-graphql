package query

import (
	"bytes"
	"fmt"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"html/template"
	"strings"
)

const msgTpl = `

var {{ .Descriptor.Name }} {{ .Descriptor.Name }}Client

func init() {
	{{ .Descriptor.Name }} = New{{ .Descriptor.Name }}Client(pg.GrpcConnection("{{ .Dns }}"))
	{{- range $method := .Methods }}
	queries = append(queries, &graphql.Field{
		Name: "{{ $.ServiceName }}_{{ $method.Name }}",
		Type: {{ $method.Output }}_type,
		Args: {{ $method.Input }}_args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return {{ $.ServiceName }}.{{ $method.Name }}(p.Context, {{ $method.Input }}_from_args(p.Args))
		},
	})
	{{ end }}

	{{- range $dataloader := .RequiredDataloaders }}
	if dataloadersToRegister == nil {
		dataloadersToRegister = make(map[string][]pg.RegisterDataloaderFn)
	}

	if _, ok := dataloadersToRegister["{{ $dataloader.Name }}"]; !ok {
		dataloadersToRegister["{{ $dataloader.Name }}"] = []pg.RegisterDataloaderFn{}
	}

	dataloadersToRegister["{{ $dataloader.Name }}"] = append(dataloadersToRegister["{{ $dataloader.Name }}"], func(dl pg.Dataloader) {
		{{ $dataloader.Source }}_type.AddFieldConfig("{{ $dataloader.FieldName }}", &graphql.Field{
			Type: dl.Output,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				source, _ := p.Source.(*{{ $dataloader.Source }})
				return dl.Fn(p.Context, {{ $dataloader.IdFieldName }})
			},
		})
	})

	{{ end }}

	{{- range $dataloader := .DataloadersToProvide }}
	if dataloadersToProvide == nil {
		dataloadersToProvide = make(map[string]pg.Dataloader)
	}
	dataloadersToProvide["{{ $dataloader.Name }}"] = pg.Dataloader{
		Fn: func(ctx context.Context, ids []string) (interface{}, error) {
			resp, err := {{ $.Descriptor.Name }}.{{ $dataloader.MethodName }}(ctx, &{{ $dataloader.Input }}{ {{ $dataloader.IdField }}: ids})

			if err != nil {
				return nil, err
			}
			return resp.{{ $dataloader.ObjectField }}, nil
		},
		Output: graphql.NewList({{ $dataloader.ReturnType }}_type),
	}
	{{- end }}

}
`

type DataloaderConfig struct {
	Name        string
	IdFieldName string
	FieldName   string
	Source      string
}

type DataloaderProviderConfig struct {
	Name        string
	IdField     string
	MethodName  string
	Input       string
	Output      string
	ObjectField string
	ReturnType  string
}

type Message struct {
	Descriptor           *descriptorpb.ServiceDescriptorProto
	Options              *graphql.MutationOption
	Methods              []Method
	DataloadersToProvide []DataloaderProviderConfig
	RequiredDataloaders  []DataloaderConfig
	ServiceName          string
	Dns                  string
}

func New(msg *descriptorpb.ServiceDescriptorProto, root *descriptorpb.FileDescriptorProto) (m Message) {
	var methods []Method

	dns := proto.GetExtension(msg.Options, graphql.E_Host).(string)

	var dataloaders []DataloaderConfig
	var dataloadersToProvider []DataloaderProviderConfig
	for _, method := range msg.Method {
		// See if output has any dataloaded fields
		var output *descriptorpb.DescriptorProto
		for _, msgType := range root.MessageType {
			if last(*method.OutputType) == *msgType.Name {
				output = msgType
			}
		}

		// See if method is a dataloader service method
		if proto.HasExtension(method.Options, graphql.E_DataloaderService) {
			serviceOptions := proto.GetExtension(method.Options, graphql.E_DataloaderService).(*graphql.DataloaderServiceOptions)

			providerConfig := DataloaderProviderConfig{
				Name:       serviceOptions.Name,
				MethodName: strings.Title(*method.Name),
				Input:      strings.Title(last(*method.InputType)),
				Output:     strings.Title(last(*method.OutputType)),
			}

			var input *descriptorpb.DescriptorProto
			for _, msgType := range root.MessageType {
				if last(*method.InputType) == *msgType.Name {
					input = msgType
				}
			}
			for _, field := range input.Field {
				if proto.HasExtension(field.Options, graphql.E_DataloaderIds) {
					dataloader := proto.GetExtension(field.Options, graphql.E_DataloaderIds).(bool)
					if dataloader {
						providerConfig.IdField = strings.Title(*field.JsonName)
					}
				}
			}
			for _, field := range output.Field {
				if proto.HasExtension(field.Options, graphql.E_DataloaderObject) {
					object := proto.GetExtension(field.Options, graphql.E_DataloaderObject).(bool)
					if object {
						providerConfig.ObjectField = strings.Title(*field.JsonName)
						providerConfig.ReturnType = last(*field.TypeName)
					}
				}
			}
			dataloadersToProvider = append(dataloadersToProvider, providerConfig)
			// Find id field

		}

		for _, field := range output.Field {
			dataloaders = getDataloaders(root, *output.Name, field, dataloaders)
		}

		methods = append(methods, Method{Input: last(*method.InputType), Output: last(*method.OutputType), Name: strings.Title(*method.Name)})
	}
	return Message{
		Descriptor:           msg,
		Methods:              methods,
		ServiceName:          *msg.Name,
		RequiredDataloaders:  dataloaders,
		DataloadersToProvide: dataloadersToProvider,
		Dns:                  dns,
	}
}

func getDataloaders(root *descriptorpb.FileDescriptorProto, outputName string, field *descriptorpb.FieldDescriptorProto, dataloaders []DataloaderConfig) []DataloaderConfig {

	if field.TypeName != nil {
		for _, msgType := range root.MessageType {
			if last(*field.TypeName) == *msgType.Name {
				for _, field := range msgType.Field {
					dataloaders = getDataloaders(root, *msgType.Name, field, dataloaders)
				}
			}
		}
	}

	if proto.HasExtension(field.Options, graphql.E_FieldResolver) {
		fieldResolver := proto.GetExtension(field.Options, graphql.E_FieldResolver).(*graphql.FieldResolver)

		IdFieldName := fmt.Sprintf("[]string{source.%s}", strings.Title(*field.JsonName))
		if field.Label.String() == "LABEL_REPEATED" {
			IdFieldName = fmt.Sprintf("source.%s", strings.Title(*field.JsonName))
		}

		dataloaders = append(dataloaders, DataloaderConfig{
			Name:        fieldResolver.DataloaderName,
			IdFieldName: IdFieldName,
			FieldName:   fieldResolver.FieldName,
			Source:      outputName,
		})
	}
	return dataloaders
}

func (m Message) Imports() []string {
	if len(m.DataloadersToProvide) > 0 {
		return []string{"context"}
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
