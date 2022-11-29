package dataloaders

import "text/template"

const msgTpl = `

var {{ .Descriptor.Name }}ClientInstance {{ .Descriptor.Name }}Client

func init() {
	fieldInits = append(fieldInits, func(opts ...grpc.DialOption) {
		{{ .Descriptor.Name }}ClientInstance = New{{ .Descriptor.Name }}Client(pg.GrpcConnection("{{ .Dns }}", opts...))
	})
	
	{{- range $method := .Methods }}
	fields = append(fields, &gql.Field{
		Name: "{{ $.ServiceName }}_{{ $method.Name }}",
		Type: {{ $method.Output }}GraphqlType,
		Args: {{ $method.Input }}GraphqlArgs,
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			return {{ $.Descriptor.Name }}ClientInstance.{{ $method.Name }}(p.Context, {{ $method.Input }}FromArgs(p.Args))
		},
	})
	{{ end }}
}

{{ if .Loaders }}
func WithLoaders(ctx context.Context) context.Context {
	{{- range $loader :=.Loaders }}
	ctx = context.WithValue(ctx, "{{ $loader.Method }}Loader", dataloader.NewBatchedLoader(
		func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
			var results []*dataloader.Result

			{{ if $loader.Custom }}
			var requests []*{{ $loader.KeysType }}
			for _, key := range keys {
				requests = append(requests, key.(*{{ $loader.KeysType }}Key).{{ $loader.KeysType }})
			}
			resp, err := {{ $.Descriptor.Name }}ClientInstance.{{ $loader.Method }}(ctx, &{{ $loader.RequestType }}{
				{{ $loader.KeysField }}: requests,
			})
			{{- else }}
			resp, err := {{ $.Descriptor.Name }}ClientInstance.{{ $loader.Method }}(ctx, &pg.BatchRequest{
				{{ $loader.KeysField }}: keys.Keys(),
			})
			{{- end }}
		
			if err != nil {
				return results
			}
	
			for _, key := range keys.Keys() {
				if val, ok := resp.{{ $loader.ResultsField }}[key]; ok {
					results = append(results, &dataloader.Result{Data: val})
				{{- if $loader.Custom }}
				} else if err, ok := resp.Errors[key]; ok {
					results = append(results, &dataloader.Result{Error: errors.New(err)})
				{{- end }}
				} else {
					var empty {{ $loader.ResultsType }}
					results = append(results, &dataloader.Result{Data: empty})
				}
			}
	
			return results
		},
	))
	{{ end }}
	return ctx
}
{{ end }}

{{ range $loader :=.Loaders }}
{{ if $loader.Custom }}

type {{ $loader.KeysType }}Key struct {
	*{{ $loader.KeysType }}
}

func (key *{{ $loader.KeysType }}Key) String() string {
	return pg.ProtoKey(key)
}

func (key *{{ $loader.KeysType }}Key) Raw() interface{} {
	return key
}

{{ end }}
func {{ $loader.Method }}(p gql.ResolveParams, {{ if $loader.Custom }} key *{{ $loader.KeysType }} {{ else }} key string {{ end }}) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("{{ $loader.Method }}Loader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("{{ $loader.Method }}Loader").(*dataloader.Loader)
	default:
		panic("Please call {{ $.Package }}.WithLoaders with the current context first")
	}

	thunk := loader.Load(p.Context,  {{ if $loader.Custom }} &{{ $loader.KeysType }}Key{key} {{ else }} dataloader.StringKey(key) {{ end }})
	return func() (interface{}, error) {
				res, err := thunk()
				if err != nil {
					return nil, err
				}
				return res.({{ $loader.ResultsType }}), nil
	}, nil
}

func {{ $loader.Method }}Many(p gql.ResolveParams, {{ if $loader.Custom }} keys []*{{ $loader.KeysType }} {{ else }} keys []string{{ end }}) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("{{ $loader.Method }}Loader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("{{ $loader.Method }}Loader").(*dataloader.Loader)
	default:
		panic("Please call {{ $.Package }}.WithLoaders with the current context first")
	}

	{{ if $loader.Custom }}
	loaderKeys := make(dataloader.Keys, len(keys))
	for ix := range keys {
		loaderKeys[ix] = &{{ $loader.KeysType }}Key{keys[ix]}
	}
	{{ end }}

	thunk := loader.LoadMany(p.Context, {{ if $loader.Custom }} loaderKeys{{ else }} dataloader.NewKeysFromStrings(keys) {{ end }})
	return func() (interface{}, error) {
		resSlice, errSlice := thunk()
		
		for _, err := range errSlice {
			if err != nil {
				return nil, err
			}
		}
		
		var results []{{ $loader.ResultsType }}
		for _, res := range resSlice {
			results = append(results, res.({{ $loader.ResultsType }}))
		}

		return results, nil
	}, nil
}
{{ end }}
`

var tpl *template.Template

func init() {
	var err error
	tpl, err = template.New("msg").Parse(msgTpl)
	if err != nil {
		panic(err)
	}
}

type LoaderVars struct {
	Method       string
	RequestType  string
	KeysField    string
	KeysType     string
	ResultsField string
	ResultsType  string
	Custom       bool
}
