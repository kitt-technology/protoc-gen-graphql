package typedef

import (
	"bytes"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"html/template"
	"io/ioutil"
	"strings"
)

const typeTpl = `
{{- if not .Optional }}graphql.NewNonNull({{- end }}
			{{- if .IsList }}graphql.NewList(graphql.NewNonNull({{- end }}
			{{- if eq .TypeOfType "Object" }}{{ .GqlType }}_{{ .Suffix }}{{- end }}
			{{- if eq .TypeOfType "Wrapper" }}{{ .GqlType }}{{- end }}
			{{- if eq .TypeOfType "Primitive" }}{{ .GqlType }}{{- end }}
			{{- if eq .TypeOfType "Enum" }}{{ .GqlType }}_enum{{- end }}
			{{- if eq .TypeOfType "Timestamp" }}pg.Timestamp_{{ .Suffix }}{{- end }}
 			{{- if .IsList }})){{- end }}
			{{- if not .Optional }}){{- end }}`

const goFromArgs = `
{{- if eq .TypeOfType "Object" }}{{ .GoType  }}_from_args(val.(map[string]interface{})){{- end }}
{{- if eq .TypeOfType "Primitive" }}{{  .GoType }}(val.({{ strip_precision .GoType }})){{- end }}
{{- if eq .TypeOfType "Wrapper" }}{{  .GoType }}({{ primitive_to_wrapper .GoType }}(val.({{ wrapper_to_primitive .GoType }}))){{- end }}
{{- if eq .TypeOfType "Enum" }}val.({{ .GoType }}){{- end }}
{{- if eq .TypeOfType "Timestamp" }}pg.ToTimestamp(val){{- end }}`


const msgTpl = `


var {{ .Descriptor.GetName }}_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "{{ .ObjectName }}",
	Fields: graphql.Fields{
		{{- range $field := .Fields }}
		"{{ $field.GqlKey }}": &graphql.Field{
			Type: {{- $field.Type }},
		},
		{{- end }}
	},
})

var {{ .Descriptor.GetName }}_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "{{ .ObjectName }}Input",
	Fields: graphql.InputObjectConfigFieldMap{
		{{- range $field := .Fields }}
		"{{ $field.GqlKey }}": &graphql.InputObjectFieldConfig{
			Type: {{- $field.InputType }},
		},
		{{- end }}
	},
})

var {{ .Descriptor.GetName }}_args = graphql.FieldConfigArgument{
	{{- range $field := .Fields }}
	"{{ $field.GqlKey }}": &graphql.ArgumentConfig{
			Type: {{- $field.InputType }},
	},
	{{- end }}
}

func {{ .Descriptor.GetName }}_from_args(args map[string]interface{}) *{{ .Descriptor.GetName }} {
	return {{ .Descriptor.GetName }}_instance_from_args(&{{ .Descriptor.GetName }}{}, args)
}

func {{ .Descriptor.GetName }}_instance_from_args(objectFromArgs *{{ .Descriptor.GetName }}, args map[string]interface{}) *{{ .Descriptor.GetName }} {
	{{- range $field := .Fields }}
		{{- if $field.GoKey }}	
			
			{{- if $field.IsList }}
			if args["{{ $field.GqlKey }}"] != nil {
	
				{{ $field.GqlKey }}InterfaceList := args["{{ $field.GqlKey }}"].([]interface{})

				var {{ $field.GqlKey }} []
			{{- if eq $field.TypeOfType "Object" }}*{{- end }}
			{{- if eq $field.TypeOfType "Wrapper" }}*{{- end }}
			{{- if eq $field.TypeOfType "Timestamp" }}*{{- end }}
			{{- $field.GoType }}

				for _, val := range {{ $field.GqlKey }}InterfaceList {
					itemResolved := {{ $field.GoFromArgs }}
					{{ $field.GqlKey }} = append({{ $field.GqlKey }}, itemResolved)
				}
				objectFromArgs.{{ $field.GoKey }} = {{ $field.GqlKey }}
			}
		
			{{- else }}
				if args["{{ $field.GqlKey }}"] != nil {
					val := args["{{  $field.GqlKey }}"]
					objectFromArgs.{{ $field.GoKey }} = {{ $field.GoFromArgs }}
				}
			{{- end }}
			
		{{- end }}
	{{- end }}
	return objectFromArgs
}


func (objectFromArgs *{{ .Descriptor.GetName }}) From_args(args map[string]interface{}) {
	{{ .Descriptor.GetName }}_instance_from_args(objectFromArgs, args)
}

func (msg *{{ .Descriptor.GetName }}) XXX_type() *graphql.Object {
	return {{ .Descriptor.GetName }}_type
}

func (msg *{{ .Descriptor.GetName }}) XXX_args() graphql.FieldConfigArgument {
	return {{ .Descriptor.GetName }}_args
}
`

type Field struct {
	GqlKey    string

	GoKey  string
	GoType  GoType

	Type  string
	GoFromArgs  string
	InputType  string
	ArgType  string

	IsList      bool
	TypeOfType string
}

type FieldTypeVars struct {
	TypeOfType string
	Optional bool
	IsList bool
	GqlType GqlType
	Suffix string
	GoType GoType
	GqlKey string
}

type Message struct {
	Descriptor *descriptorpb.DescriptorProto
	Root *descriptorpb.FileDescriptorProto
	Fields     []Field
	Import     map[string]string
	ObjectName    string
}

func New(msg *descriptorpb.DescriptorProto, file *descriptorpb.FileDescriptorProto) (m Message) {
	return Message{
		Import:     make(map[string]string),
		Descriptor: msg,
		Root: file,
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
	// If there's a custom name, grab it
	if proto.HasExtension(m.Descriptor.Options, graphql.E_ObjectName) {
		m.ObjectName = proto.GetExtension(m.Descriptor.Options, graphql.E_ObjectName).(string)
	} else {
		m.ObjectName = *m.Descriptor.Name
	}

	funcMap := template.FuncMap{
		"strip_precision": stripPrecision,
		"wrapper_to_primitive": wrapperToPrimitive,
		"primitive_to_wrapper": primitiveToWrapper,
	}

	for _, field := range m.Descriptor.Field {
		isList := false
		switch field.Label.String() {
		// It's a list or a map
		case "LABEL_REPEATED":
			switch field.Type.String() {
			case "TYPE_MESSAGE":
				// Is it a map?
				isMap := false
				nestedTypeKey := last(*field.TypeName)
				for _, nestedType := range m.Descriptor.NestedType {
					if *nestedType.Name == nestedTypeKey {
						// If it's a map, continue - we don't support maps yet
						isMap = true
					}
				}

				if isMap {
					// Maps are not yet supported
					continue
				} else {
					// It's a list of objects .e.g.
					// repeated Object objects = 1;
					isList = true
				}

				break
			default:
				isList = true
			}
		}

		goType, gqlType, typeOfType := types(field, m.Root)
		if isList {
			goType, gqlType, typeOfType = types(field, m.Root)
		}

		if typeOfType == Wrapper {
			m.Import["google.golang.org/protobuf/types/known/wrapperspb"] ="google.golang.org/protobuf/types/known/wrapperspb"
		}
		if typeOfType == Timestamp && isList {
			m.Import["github.com/golang/protobuf/ptypes/timestamp"] = "github.com/golang/protobuf/ptypes/timestamp"
		}

		optional := field.TypeName != nil
		if proto.HasExtension(field.Options, graphql.E_Optional)   {
			val := proto.GetExtension(field.Options, graphql.E_Optional)
			if val.(bool) {
				optional = true
			}
		}

		fieldVars := Field{
			GqlKey: *field.JsonName,
			GoKey: goKey(field),
			GoType: goType,
			TypeOfType: string(typeOfType),
			IsList: isList,
		}

		// Generate input type
		typeVars := FieldTypeVars{
			TypeOfType: string(typeOfType),
			IsList: isList,
			GoType: goType,
			GqlType: gqlType,
			GqlKey: *field.JsonName,
			Suffix: "input_type",
			Optional: optional,
		}
		var buf bytes.Buffer
		mTpl, err := template.New("input_type").Funcs(funcMap).Parse(typeTpl)
		if err != nil {
			panic(err)
		}
		mTpl.Execute(&buf, typeVars)
		inputType, err := ioutil.ReadAll(&buf)

		// Generate generic type
		typeVars.Suffix = "type"
		mTpl, err = template.New("type").Funcs(funcMap).Parse(typeTpl)
		if err != nil {
			panic(err)
		}
		mTpl.Execute(&buf, typeVars)
		normalType, err := ioutil.ReadAll(&buf)
		typeVars.Suffix = "type"

		// Generate "from_arg" string
		mTpl, err = template.New("from_arg").Funcs(funcMap).Parse(goFromArgs)

		if err != nil {
			panic(err)
		}
		mTpl.Execute(&buf, typeVars)
		goFromArg, err := ioutil.ReadAll(&buf)
		fieldVars.GoFromArgs = string(goFromArg)
		fieldVars.InputType = string(inputType)
		fieldVars.Type = string(normalType)

		m.Fields = append(m.Fields, fieldVars)
	}


	if len(m.Fields) == 0 {
		m.Fields = append(m.Fields, Field{
			GqlKey: "_null",
			Type: "graphql.Boolean",
			InputType: "graphql.Boolean",
		})
	}

	var buf bytes.Buffer

	mTpl, err := template.New("msg").Funcs(funcMap).Parse(msgTpl)
	if err != nil {
		panic(err)
	}
	mTpl.Execute(&buf, m)

	return buf.String()
}

func goKey(field *descriptorpb.FieldDescriptorProto) string {
	name := *field.JsonName
	return strings.ToUpper(string(name[0])) + name[1:]
}


type Descriptor interface {
	GetType() descriptorpb.FieldDescriptorProto_Type
	GetTypeName() string
}

type FieldType string
type GoType string
type GqlType string

const (
	Wrapper FieldType = "Wrapper"
	Object = "Object"
	Primitive = "Primitive"
	Enum = "Enum"
	Timestamp = "Timestamp"
)

func wrapperToPrimitive(wrapperType GoType) string {
	switch wrapperType {
	case "wrapperspb.String":
		return "string"
	case "wrapperspb.Bool":
		return "bool"
	case "wrapperspb.Float":
		return "float64"
	}
	return ""
}
func stripPrecision(arg GoType) string {
	if strings.Contains(string(arg), "int") {
		output := strings.Replace(string(arg), "64", "", -1)
		return strings.Replace(output, "32", "", -1)
	}
	if strings.Contains(string(arg), "float") {
		return "float64"
	}
	return string(arg)
}
func primitiveToWrapper(wrapperType GoType) string {
	switch wrapperType {
	case "wrapperspb.Float":
		return "float32"
	case "wrapperspb.String":
		return "string"
	case "wrapperspb.Bool":
		return "bool"
	}
	return string(wrapperType)
}

func types(field Descriptor, root *descriptorpb.FileDescriptorProto) (GoType, GqlType, FieldType) {
	if field.GetTypeName() != "" {
		switch field.GetTypeName() {
		case ".google.protobuf.StringValue":
			return "wrapperspb.String", "graphql.String", Wrapper
		case ".google.protobuf.BoolValue":
			return "wrapperspb.Bool", "graphql.Boolean", Wrapper
		case ".google.protobuf.FloatValue":
			return "wrapperspb.Float", "graphql.Float", Wrapper
		case ".google.protobuf.Timestamp":
			return "timestamp.Timestamp", "pg.Timestamp", Timestamp
		case ".google.protobuf.Int32":
			return "wrapperspb.Int32", "graphql.Int", Wrapper
		case ".google.protobuf.Int64":
			return "wrapperspb.Int64", "graphql.Int", Wrapper
		}
	}

	switch field.GetType().String() {
	case "TYPE_STRING":
		return "string", "graphql.String", Primitive
	case "TYPE_INT32":
		return "int32", "graphql.Int", Primitive
	case "TYPE_INT64":
		return "int64", "graphql.Int", Primitive
	case "TYPE_FLOAT":
		return "float32", "graphql.Float", Primitive
	case "TYPE_BOOL":
		return "bool", "graphql.Boolean", Primitive
	case "TYPE_DOUBLE":
		return "float64", "graphql.Float", Primitive
	case "TYPE_BYTES":
		return "[]byte", "graphql.String", Primitive
	}

	// Search through message descriptors
	for _, messageType := range root.MessageType {
		if *messageType.Name == last(field.GetTypeName()) {
			return GoType(last(field.GetTypeName())), GqlType(*messageType.Name), Object
		}
	}

	// Search through enums
	for _, enumType := range root.EnumType {
		if *enumType.Name == last(field.GetTypeName()) {
			return GoType(last(field.GetTypeName())), GqlType(*enumType.Name), Enum
		}
	}
	panic(field)
}
