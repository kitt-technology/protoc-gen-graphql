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
	Name: "{{ .ObjectName }}",
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

var {{ .Descriptor.GetName }}_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "{{ .ObjectName }}",
	Fields: graphql.InputObjectConfigFieldMap{
		{{- range $field := .Fields }}
		"{{ $field.Key }}": &graphql.InputObjectFieldConfig{
			{{- if $field.Optional }}
			Type: {{ replace $field.Type "type" "input_type" }},
			{{- else }}
			Type: graphql.NewNonNull({{ replace $field.Type "type" "input_type" }}),
			{{- end }}
		},
		{{- end }}
	},
})

var {{ .Descriptor.GetName }}_args = graphql.FieldConfigArgument{
	{{- range $field := .Fields }}
	"{{ $field.Key }}": &graphql.ArgumentConfig{
		{{- if $field.Optional }}
		Type: {{ replace $field.Type "type" "input_type" }},
		{{- else }}
		Type: graphql.NewNonNull({{ replace $field.Type "type" "input_type" }}),
		{{- end }}
	},
	{{- end }}
}

func {{ .Descriptor.GetName }}_from_args(args map[string]interface{}) *{{ .Descriptor.GetName }} {
	objectFromArgs := {{ .Descriptor.GetName }}{}
	{{- range $field := .Fields }}
		{{- if $field.StructKey }}	
			
			{{- if $field.IsList }}
			if args["{{ $field.Key }}"] != nil {
	
				{{ $field.Key }}InterfaceList := args["{{ $field.Key }}"].([]interface{})
			
				var {{ $field.Key }} []{{ $field.StructType }}
				for _, item := range {{ $field.Key }}InterfaceList {
					{{ $field.Key }} = append({{ $field.Key }}, item.({{ $field.StructType }}))
				}
				objectFromArgs.{{ $field.StructKey }} = {{ $field.Key }}

			}
		
			{{- else }}
				{{ if $field.IsTimestamp }}
				
				{{ else }}

					{{ if $field.WrapperType }}
						if args["{{ $field.Key }}"] != nil {
							objectFromArgs.{{ $field.StructKey }} = wrapperspb.{{ $field.WrapperType.Type }}(args["{{ $field.Key }}"].({{ $field.WrapperType.Primitive }}))
						}
					{{ end }}
	
					{{ if not $field.WrapperType }}
						{{ if $field.IsObject }}
							if args["{{ $field.Key }}"] != nil {
								objectFromArgs.{{ $field.StructKey }} = {{ replace $field.StructType "*" "" }}_from_args(args["{{ $field.Key }}"].(map[string]interface{}))
							}
						{{ else }}
							objectFromArgs.{{ $field.StructKey }} = args["{{ $field.Key }}"].({{ $field.StructType }})
						{{ end }}
					{{ end }}

					
				{{ end }}
				
			{{- end }}
			
		{{- end }}
	{{- end }}


	return &objectFromArgs
}

func (objectFromArgs *{{ .Descriptor.GetName }}) From_args(args map[string]interface{}) {
		objectFromArgs = {{ .Descriptor.GetName }}_from_args(args)

}

func (msg *{{ .Descriptor.GetName }}) XXX_type() *graphql.Object {
	return {{ .Descriptor.GetName }}_type
}

func (msg *{{ .Descriptor.GetName }}) XXX_args() graphql.FieldConfigArgument {
	return {{ .Descriptor.GetName }}_args
}
`

type Message struct {
	Descriptor *descriptorpb.DescriptorProto
	Fields     []Field
	Import     map[string]string
	ObjectName    string
}

func New(msg *descriptorpb.DescriptorProto) (m Message) {
	return Message{
		Import:     make(map[string]string),
		Descriptor: msg,
	}
}

func last(path string) string {
	t := strings.Split(path, ".")
	return t[len(t)-1]
}

func (m Message) Imports() []string {
	m.Generate()
	var imports []string
	for _, val := range m.Import {
		imports = append(imports, val)
	}
	return imports
}

func (m Message) Generate() string {
	if proto.HasExtension(m.Descriptor.Options, graphql.E_ObjectName) {
		m.ObjectName = proto.GetExtension(m.Descriptor.Options, graphql.E_ObjectName).(string)
	} else {
		m.ObjectName = *m.Descriptor.Name
	}

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
						Optional:   true,
						StructKey:  toGoStruct(field),
						StructType: toGoType(field),
						IsList:     true,
					})
				}
				break
			default:
				t := fmt.Sprintf("graphql.NewList(%s)", protoToGraphqlType(field.Type.String()))

				isEnum := false
				structType := toGoType(field)
				if field.Type.String() == "TYPE_ENUM" {
					t = fmt.Sprintf("%s_type", protoToGraphqlType(*field.TypeName))
					t = strings.Replace(t, "_type", "_enum", -1)
					t = fmt.Sprintf("graphql.NewList(%s)", t)
					structType = strings.Replace(structType, "*", "", -1)
					isEnum = true
				}

				m.Fields = append(m.Fields, Field{
					Key:        *field.JsonName,
					Type:       t,
					Optional:   true,
					StructKey:  toGoStruct(field),
					StructType: structType,
					IsList:     true,
					IsEnum: isEnum,
				})
			}
			break
		// Its a normal type
		default:
			if field.TypeName != nil {
				wrapperType := IsWrapper(field)
				isObject := IsObject(field)

				t := fmt.Sprintf("%s_type", protoToGraphqlType(*field.TypeName))
				if wrapperType != nil {
					m.Import["google.golang.org/protobuf/types/known/wrapperspb"] = "google.golang.org/protobuf/types/known/wrapperspb"
					t = fmt.Sprintf("graphql.%s", wrapperType.GraphqlType)
				}

				structType := toGoType(field)
				isEnum := false
				if field.Type.String() == "TYPE_ENUM" {
					structType = strings.Replace(structType, "*", "", -1)
					t = strings.Replace(t, "_type", "_enum", -1)
					isEnum = true
				}

				isTimestamp := false
				if structType == "*Timestamp" {
					isTimestamp = true
				}

				m.Fields = append(m.Fields, Field{
					Key:         *field.JsonName,
					Type:        t,
					Optional:    true,
					StructKey:   toGoStruct(field),
					StructType:  structType,
					WrapperType: wrapperType,
					IsTimestamp: isTimestamp,
					IsObject: isObject,
					IsEnum: isEnum,
				})

			} else {
				wrapperType := IsWrapper(field)
				isObject := IsObject(field)

				m.Fields = append(m.Fields, Field{
					Key:         *field.JsonName,
					Type:        fmt.Sprintf("%s", protoToGraphqlType(field.Type.String())),
					Optional:    false,
					StructKey:   toGoStruct(field),
					StructType:  toGoType(field),
					WrapperType: wrapperType,
					IsObject: isObject,
				})
			}

		}
	}

	log.Println(m.Fields)

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

	funcMap := template.FuncMap{
		"replace": replace,
	}

	mTpl, err := template.New("msg").Funcs(funcMap).Parse(msgTpl)
	if err != nil {
		panic(err)
	}
	mTpl.Execute(&buf, m)

	return buf.String()
}

func replace(input, from, to string) string {
	return strings.Replace(input, from, to, -1)
}
func protoToGraphqlType(protoType string) string {
	log.Println(protoType)
	switch protoType {
	case "TYPE_STRING":
		return "graphql.String"
	case "TYPE_INT32":
		return "graphql.Int"
	case "TYPE_FLOAT":
		return "graphql.Float"
	case "TYPE_DOUBLE":
		return "graphql.Float"
	case "TYPE_BYTES":
		return "graphql.String"
	case "TYPE_ENUM":
		return protoType + "_enum"
	case "TYPE_BOOL":
		return "graphql.Boolean"
	case ".google.protobuf.StringValue":
		return "graphql.String"
	case ".google.protobuf.BoolValue":
		return "graphql.Bool"
	case ".google.protobuf.FloatValue":
		return "graphql.Float"
	case ".google.protobuf.Timestamp":
		return "pg.Timestamp"
	}
	return last(protoType)
}

func toGoStruct(field *descriptorpb.FieldDescriptorProto) string {
	name := *field.JsonName

	return strings.ToUpper(string(name[0])) + name[1:]
}

func IsObject(field *descriptorpb.FieldDescriptorProto) bool {
	return field.TypeName != nil && field.Type.String() != "TYPE_ENUM"
}

func IsWrapper(field *descriptorpb.FieldDescriptorProto) *WrapperType {
	if field.TypeName != nil {
		if strings.Contains(*field.TypeName, "google.protobuf.") {
			switch *field.TypeName {
			case ".google.protobuf.StringValue":
				return &WrapperType{Type: "String", Primitive: "string", GraphqlType: "String"}
			case ".google.protobuf.BoolValue":
				return &WrapperType{Type: "Bool", Primitive: "bool", GraphqlType: "Boolean"}
			case ".google.protobuf.FloatValue":
				return &WrapperType{Type: "Float", Primitive: "float32", GraphqlType: "Float"}
			}
		}
	}
	return nil
}

func toGoType(field *descriptorpb.FieldDescriptorProto) string {
	if field.TypeName != nil {
		switch *field.TypeName {
		case ".google.protobuf.StringValue":
			return "*wrappers.StringValue"
		case ".google.protobuf.BoolValue":
			return "*wrappers.BoolValue"
		case ".google.protobuf.FloatValue":
			return "*wrappers.FloatValue"
		}

		if strings.Contains(last(*field.TypeName), "*") {
			return last(*field.TypeName)
		}

		return "*" + last(*field.TypeName)
	}

	switch field.Type.String() {
	case "TYPE_STRING":
		return "string"
	case "TYPE_INT32":
		return "int"
	case "TYPE_FLOAT":
		return "float32"
	case "TYPE_BOOL":
		return "bool"
	case "TYPE_DOUBLE":
		return "float64"
	case "TYPE_BYTES":
		return "[]byte"
	}
	return ""
}

type Field struct {
	Optional    bool
	Key         string
	Type        string
	StructKey   string
	StructType  string
	IsList      bool
	WrapperType *WrapperType
	IsTimestamp bool
	IsObject bool
	IsEnum bool
}

type WrapperType struct {
	Type        string
	Primitive   string
	GraphqlType string
}
