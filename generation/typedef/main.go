package typedef

import (
	"bytes"
	"fmt"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"html/template"
	"log"
	"strings"
)

const msgTpl = `

var {{ .Descriptor.GetName }}_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "{{ .Descriptor.GetName }}",
	Fields: graphql.Fields{
		{{- range $field := .Fields }}
		"{{ $field.Key }}": &graphql.Field{
			{{- if $field.Optional }}
			Type: {{ $field.Type }},
			{{- else }}
			Type: graphql.NewNonNull({{ $field.Type }}),
			{{- end }}
		},
		{{- end }}
	},
})

var {{ .Descriptor.GetName }}_args = graphql.FieldConfigArgument{
	{{- range $field := .Fields }}
	"{{ $field.Key }}": &graphql.ArgumentConfig{
		{{- if $field.Optional }}
		Type: {{ $field.Type }},
		{{- else }}
		Type: graphql.NewNonNull({{ $field.Type }}),
		{{- end }}
	},
	{{- end }}
}

func {{ .Descriptor.GetName }}_from_args(args map[string]interface{}) *{{ .Descriptor.GetName }} {
	objectFromArgs := {{ .Descriptor.GetName }}{}
	{{- range $field := .Fields }}
		{{- if $field.StructKey }}	
			
			{{- if $field.IsList }}

			{{ $field.Key }}InterfaceList := args["{{ $field.Key }}"].([]interface{})
		
			var {{ $field.Key }} []{{ $field.StructType }}
			for _, item := range {{ $field.Key }}InterfaceList {
				{{ $field.Key }} = append({{ $field.Key }}, item.({{ $field.StructType }}))
			}
			objectFromArgs.{{ $field.StructKey }} = {{ $field.Key }}

			{{- else }}
			objectFromArgs.{{ $field.StructKey }} = args["{{ $field.Key }}"].({{ $field.StructType }})

			{{- end }}
			
		{{- end }}
	{{- end }}


	return &objectFromArgs
}
`

type Message struct {
	Descriptor *descriptorpb.DescriptorProto
	Options    *graphql.MutationOption
	Fields     []Field
}

func New(msg *descriptorpb.DescriptorProto) (m Message) {
	return Message{
		Options:    proto.GetExtension(msg.Options, graphql.E_MutationOptions).(*graphql.MutationOption),
		Descriptor: msg,
	}
}

func last(path string) string {
	t := strings.Split(path, ".")
	return t[len(t)-1]
}

func (m Message) Imports() []string {
	return []string{}
}

func (m Message) Generate() string {
	for _, field := range m.Descriptor.Field {
		switch field.Label.String() {

		// It's a list or a map
		case "LABEL_REPEATED":
			switch field.Type.String() {
			case "TYPE_MESSAGE":
				// Is it a map?
				hasMapEntry := false
				nestedTypeKey := last(*field.TypeName)
				for _, nestedType := range m.Descriptor.NestedType {
					if *nestedType.Name == nestedTypeKey {
						// Maps which we havent dealt with yet TODO
						//if nestedType.Field[1].TypeName != nil {
						//	m.Fields = append(m.Fields, Field{
						//		Key:      *field.JsonName,
						//		Type:     fmt.Sprintf("graphql.NewList(%s_tuple)", last(*nestedType.Field[1].TypeName)),
						//		Optional: false,
						//	})
						//} else {
						//	m.Fields = append(m.Fields, Field{
						//		Key:      *field.JsonName,
						//		Type:     fmt.Sprintf("graphql.NewList(%s_tuple)", protoToGraphqlType(nestedType.Field[1].Type.String())),
						//		Optional: false,
						//	})
						//}
						hasMapEntry = true
					}
				}
				if !hasMapEntry {
					m.Fields = append(m.Fields, Field{
						Key:        *field.JsonName,
						Type:       fmt.Sprintf("graphql.NewList(%s_type)", nestedTypeKey),
						Optional:   false,
						StructKey:  toGoStruct(field),
						StructType: "*" + toGoType(field),
						IsList:     true,
					})
				}
				break
			default:
				m.Fields = append(m.Fields, Field{
					Key:        *field.JsonName,
					Type:       fmt.Sprintf("graphql.NewList(graphql.%s)", protoToGraphqlType(field.Type.String())),
					Optional:   false,
					StructKey:  toGoStruct(field),
					StructType: toGoType(field),
					IsList:     true,
				})
				log.Printf("%d is a list of %s", *field.Number, field.Type.String())
			}
			break
		// Its a normal type
		default:
			if field.TypeName != nil {

				m.Fields = append(m.Fields, Field{
					Key:        *field.JsonName,
					Type:       fmt.Sprintf("graphql.%s", protoToGraphqlType(*field.TypeName)),
					Optional:   true,
					StructKey:  toGoStruct(field),
					StructType: toGoType(field),
				})

			} else {
				m.Fields = append(m.Fields, Field{
					Key:        *field.JsonName,
					Type:       fmt.Sprintf("graphql.%s", protoToGraphqlType(field.Type.String())),
					Optional:   false,
					StructKey:  toGoStruct(field),
					StructType: toGoType(field),
				})
			}

		}
	}

	if len(m.Fields) == 0 {
		m.Fields = append(m.Fields, Field{
			Key:        "message",
			Type:       "graphql.String",
			Optional:   true,
			StructKey:  "",
			StructType: "",
		})
	}

	var buf bytes.Buffer
	mTpl, err := template.New("msg").Parse(msgTpl)
	if err != nil {
		panic(err)
	}
	mTpl.Execute(&buf, m)

	return buf.String()
}

func protoToGraphqlType(protoType string) string {
	switch protoType {
	case "TYPE_STRING":
		return "String"
	case "TYPE_INT32":
		return "Int"
	case "TYPE_BOOL":
		return "Boolean"
	case ".google.protobuf.StringValue":
		return "String"
	}
	return protoType
}

func toGoStruct(field *descriptorpb.FieldDescriptorProto) string {
	name := *field.JsonName
	return strings.ToUpper(string(name[0])) + name[1:]
}

func toGoType(field *descriptorpb.FieldDescriptorProto) string {
	if field.TypeName != nil {
		switch *field.TypeName {
		case ".google.protobuf.StringValue":
			return "*wrappers.StringValue"

		}

		return last(*field.TypeName)
	}

	switch field.Type.String() {
	case "TYPE_STRING":
		return "string"
	case "TYPE_INT32":
		return "int32"
	case "TYPE_BOOL":
		return "bool"
	}
	return ""
}

type Field struct {
	Optional   bool
	Key        string
	Type       string
	StructKey  string
	StructType string
	IsList     bool
}
