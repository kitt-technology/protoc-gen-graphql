package types

import "text/template"

const typeTpl = `
{{- if not .GraphqlOptional }}gql.NewNonNull({{- end }}
{{- if .IsList }}gql.NewList(gql.NewNonNull({{- end }}
{{- if eq .TypeOfType "Object" }}{{ .GqlType }}{{ .Suffix }}{{- end }}
{{- if eq .TypeOfType "Wrapper" }}{{ .GqlType }}{{- end }}
{{- if eq .TypeOfType "Primitive" }}{{ .GqlType }}{{- end }}
{{- if eq .TypeOfType "Enum" }}{{ .GqlType }}GraphqlEnum{{- end }}
{{- if eq .TypeOfType "Timestamp" }}pg.Timestamp{{ .Suffix }}{{- end }}
{{- if eq .TypeOfType "WrappedString" }}pg.WrappedString{{ .Suffix }}{{- end }}
{{- if eq .TypeOfType "Common" }}{{ .GqlType }}{{ .Suffix }}{{- end }}
{{- if .IsList }})){{- end }}
{{- if not .GraphqlOptional }}){{- end }}`

const goFromArgs = `
{{- if eq .TypeOfType "Object" }}{{ .GoType  }}FromArgs(val.(map[string]interface{})){{- end }}
{{- if eq .TypeOfType "Primitive" }}{{  .GoType }}(val.({{ strip_precision .GoType }})){{- end }}
{{- if eq .TypeOfType "Wrapper" }}{{  .GoType }}({{ primitive_to_wrapper .GoType }}(val.({{ wrapper_to_primitive .GoType }}))){{- end }}
{{- if eq .TypeOfType "Enum" }}val.({{ .GoType }}){{- end }}
{{- if eq .TypeOfType "Timestamp" }}pg.ToTimestamp(val){{- end }}
{{- if eq .TypeOfType "Common" }}{{ .GoType  }}FromArgs(val.(map[string]interface{})){{- end }}`

const msgTpl = `
var {{ .Descriptor.GetName }}GraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "{{ .ObjectName }}",
	Fields: gql.Fields{
		{{- range $field := .Fields }}
		"{{ $field.GqlKey }}": &gql.Field{
			Type: {{- $field.Type }},
			{{- if eq .TypeOfType "Wrapper" }}
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				if p.Source.(*{{ $.Descriptor.GetName }}) == nil || p.Source.(*{{ $.Descriptor.GetName }}).{{ $field.GoKey }} == nil{
					return nil, nil
				}
				return p.Source.(*{{ $.Descriptor.GetName }}).{{ $field.GoKey }}.Value, nil
			},
			{{ end }}
		},
		{{- end }}
		{{- range $name, $fields := .OneOfFields }}
		"{{ $name }}": &gql.Field{
			Type: {{ $name }}GraphqlType,
		},
		{{- end }}
	},
})
{{- if not .InputTypeName }}
var {{ .Descriptor.GetName }}GraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "{{ .ObjectName }}Input",
	Fields: gql.InputObjectConfigFieldMap{
		{{- range $field := .Fields }}
		"{{ $field.GqlKey }}": &gql.InputObjectFieldConfig{
			Type: {{- $field.InputType }},
		},
		{{- end }}
	},
})

var {{ .Descriptor.GetName }}GraphqlArgs = gql.FieldConfigArgument{
	{{- range $field := .Fields }}
	"{{ $field.GqlKey }}": &gql.ArgumentConfig{
			Type: {{- $field.InputType }},
	},
	{{- end }}
}
{{- else }}
// Using custom input type: {{ .InputTypeName }}
var {{ .Descriptor.GetName }}GraphqlInputType = {{ .InputTypeName }}

var {{ .Descriptor.GetName }}GraphqlArgs = gql.FieldConfigArgument{
	{{- range $field := .Fields }}
	"{{ $field.GqlKey }}": &gql.ArgumentConfig{
			Type: {{- $field.InputType }},
	},
	{{- end }}
}
{{- end }}

func {{ .Descriptor.GetName }}FromArgs(args map[string]interface{}) *{{ .Descriptor.GetName }} {
	return {{ .Descriptor.GetName }}InstanceFromArgs(&{{ .Descriptor.GetName }}{}, args)
}

func {{ .Descriptor.GetName }}InstanceFromArgs(objectFromArgs *{{ .Descriptor.GetName }}, args map[string]interface{}) *{{ .Descriptor.GetName }} {
	{{- range $field := .Fields }}
		{{- if $field.GoKey }}
			{{- if $field.IsList }}
			if args["{{ $field.GqlKey }}"] != nil {
				{{ $field.GqlKey }}InterfaceList := args["{{ $field.GqlKey }}"].([]interface{})
				{{ $field.GqlKey }} := make([]{{- if $field.IsPointer }}*{{- end}}{{- $field.GoType }}, 0)

				for _, val := range {{ $field.GqlKey }}InterfaceList {
					itemResolved := {{ $field.GoFromArgs }}
					{{ $field.GqlKey }} = append({{ $field.GqlKey }}, itemResolved)
				}
				objectFromArgs.{{ $field.GoKey }} = {{ $field.GqlKey }}
			}
			{{- else }}
			if args["{{ $field.GqlKey }}"] != nil {
				val := args["{{  $field.GqlKey }}"]
				{{- if $field.Proto3Optional }}
				ptr := {{ $field.GoFromArgs }}
				objectFromArgs.{{ $field.GoKey }} = &ptr
				{{- else }}
				objectFromArgs.{{ $field.GoKey }} = {{ $field.GoFromArgs }}
				{{- end}}
			}
			{{- end }}
		{{- end }}
	{{- end }}
	return objectFromArgs
}


func (objectFromArgs *{{ .Descriptor.GetName }}) FromArgs(args map[string]interface{}) {
	{{ .Descriptor.GetName }}InstanceFromArgs(objectFromArgs, args)
}

func (msg *{{ .Descriptor.GetName }}) XXX_GraphqlType() *gql.Object {
	return {{ .Descriptor.GetName }}GraphqlType
}

func (msg *{{ .Descriptor.GetName }}) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return {{ .Descriptor.GetName }}GraphqlArgs
}

func (msg *{{ .Descriptor.GetName }}) XXX_Package() string {
	return "{{ .Package }}"
}

{{- range $name, $fields := .OneOfFields }}
var {{ $name }}GraphqlType = gql.NewUnion(gql.UnionConfig{
	Name: "{{ $name }}",
	Types: []*gql.Object{
		{{- range $field := $fields }}
		{{- $field.Type }},
		{{- end }}
	},
	ResolveType: (func(p gql.ResolveTypeParams) *gql.Object {
		switch p.Value.(type) {
		{{- range $field := $fields }}
		case *{{ $.Descriptor.GetName }}_{{- $field.GoKey }}:
			fields := gql.Fields{}
			for name, field := range {{ $field.GoType }}GraphqlType.Fields() {
				fields[name] = &gql.Field{
					Name: field.Name,
					Type: field.Type,
					Resolve: func(p gql.ResolveParams) (interface{}, error) {
						wrapper := p.Source.(*{{ $.Descriptor.GetName }}_{{- $field.GoKey }})
						p.Source = wrapper.{{- $field.GoKey }}
						return gql.DefaultResolveFn(p)
					},
				}
			}
			return  gql.NewObject(gql.ObjectConfig{
				Name: {{- $field.GoType }}GraphqlType.Name(),
				Fields: fields,
			})
		{{- end }}
		}
		return nil
	}),
})
{{- end }}
`

type FieldTypeVars struct {
	TypeOfType      string
	GraphqlOptional bool
	IsList          bool
	GqlType         GqlType
	Suffix          string
	GoType          GoType
	GqlKey          string
}

var (
	messageTemplate    *template.Template
	typeTemplate       *template.Template
	goFromArgsTemplate *template.Template
)

func init() {
	funcMap := template.FuncMap{
		"strip_precision":      stripPrecision,
		"wrapper_to_primitive": wrapperToPrimitive,
		"primitive_to_wrapper": primitiveToWrapper,
	}
	var err error
	messageTemplate, err = template.New("msg").Funcs(funcMap).Parse(msgTpl)
	if err != nil {
		panic(err)
	}
	typeTemplate, err = template.New("msg").Funcs(funcMap).Parse(typeTpl)
	if err != nil {
		panic(err)
	}
	goFromArgsTemplate, err = template.New("msg").Funcs(funcMap).Parse(goFromArgs)
	if err != nil {
		panic(err)
	}
}
