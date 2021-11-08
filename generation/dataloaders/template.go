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
	
			resp, err := {{ $.Descriptor.Name }}ClientInstance.{{ $loader.Method }}(ctx, &pg.BatchRequest{
				{{ $loader.KeysField }}: keys.Keys(),
			})
	
			if err != nil {
				return results
			}
	
			for _, key := range keys.Keys() {
				if val, ok := resp.{{ $loader.ResultsField }}[key]; ok {
					results = append(results, &dataloader.Result{Data: val})
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
func {{ $loader.Method }}(p gql.ResolveParams, key string) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("{{ $loader.Method }}Loader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("{{ $loader.Method }}Loader").(*dataloader.Loader)
	default:
		panic("Please call {{ $.Package }}.WithLoaders with the current context first")
	}

	thunk := loader.Load(p.Context, dataloader.StringKey(key))
	return func() (interface{}, error) {
				res, err := thunk()
				if err != nil {
					return nil, err
				}
				return res.({{ $loader.ResultsType }}), nil
	}, nil
}

func {{ $loader.Method }}Many(p gql.ResolveParams, keys []string) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("{{ $loader.Method }}Loader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("{{ $loader.Method }}Loader").(*dataloader.Loader)
	default:
		panic("Please call {{ $.Package }}.WithLoaders with the current context first")
	}

	thunk := loader.LoadMany(p.Context, dataloader.NewKeysFromStrings(keys))
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
	ResultsField string
	ResultsType  string
}
