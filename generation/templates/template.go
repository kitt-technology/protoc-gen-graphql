package templates

import "text/template"

const msgTpl = `

var {{ .Descriptor.Name }}ClientInstance {{ .Descriptor.Name }}Client
var {{ .Descriptor.Name }}ServiceInstance {{ .Descriptor.Name }}Server
var {{ .Descriptor.Name }}DialOpts []grpc.DialOption

type {{ .Descriptor.Name }}Option func(*{{ .Descriptor.Name }}Config)

type {{ .Descriptor.Name }}Config struct {
	service {{ .Descriptor.Name }}Server
	client  {{ .Descriptor.Name }}Client
	dialOpts []grpc.DialOption
}

// WithService sets the service implementation for direct calls (no gRPC)
func WithService(service {{ .Descriptor.Name }}Server) {{ .Descriptor.Name }}Option {
	return func(cfg *{{ .Descriptor.Name }}Config) {
		cfg.service = service
	}
}

// WithClient sets the gRPC client for remote calls
func WithClient(client {{ .Descriptor.Name }}Client) {{ .Descriptor.Name }}Option {
	return func(cfg *{{ .Descriptor.Name }}Config) {
		cfg.client = client
	}
}

// WithDialOptions sets the dial options for the gRPC client
func WithDialOptions(opts ...grpc.DialOption) {{ .Descriptor.Name }}Option {
	return func(cfg *{{ .Descriptor.Name }}Config) {
		cfg.dialOpts = opts
	}
}

func Init(ctx context.Context, opts ...{{ .Descriptor.Name }}Option) (context.Context, []*gql.Field) {
	cfg := &{{ .Descriptor.Name }}Config{}
	for _, opt := range opts {
		opt(cfg)
	}

	{{ .Descriptor.Name }}ServiceInstance = cfg.service
	{{ .Descriptor.Name }}ClientInstance = cfg.client
	{{ .Descriptor.Name }}DialOpts = cfg.dialOpts

	var fields []*gql.Field
	{{- range $method := .Methods }}
	fields = append(fields, &gql.Field{
		Name: "{{ $.ServiceName }}_{{ $method.Name }}",
		Type: {{ $method.Output }}GraphqlType,
		Args: {{ $method.Input }}GraphqlArgs,
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			if {{ $.Descriptor.Name }}ServiceInstance != nil {
				return {{ $.Descriptor.Name }}ServiceInstance.{{ $method.Name }}(p.Context, {{ $method.Input }}FromArgs(p.Args))
			}
			if {{ $.Descriptor.Name }}ClientInstance == nil {
				{{ $.Descriptor.Name }}ClientInstance = get{{ $.Descriptor.Name }}Client()
			}
			return {{ $.Descriptor.Name }}ClientInstance.{{ $method.Name }}(p.Context, {{ $method.Input }}FromArgs(p.Args))
		},
	})
	{{ end }}

	{{ if .Loaders }}
	ctx = {{ .Descriptor.Name }}WithLoaders(ctx)
	{{ end }}

	return ctx, fields
}

func get{{ .Descriptor.Name }}Client() {{ .Descriptor.Name }}Client {
	host := "{{ .Dns }}"
	envHost := os.Getenv("SERVICE_HOST")
	if envHost != "" {
		host = envHost
	}
	return New{{ .Descriptor.Name }}Client(pg.GrpcConnection(host, {{ .Descriptor.Name }}DialOpts...))
}

// Set{{ .Descriptor.Name }}Service sets the service implementation for direct calls (no gRPC)
func Set{{ .Descriptor.Name }}Service(service {{ .Descriptor.Name }}Server) {
	{{ .Descriptor.Name }}ServiceInstance = service
}

// Set{{ .Descriptor.Name }}Client sets the gRPC client for remote calls
func Set{{ .Descriptor.Name }}Client(client {{ .Descriptor.Name }}Client) {
	{{ .Descriptor.Name }}ClientInstance = client
}

{{ if .Loaders }}
func {{ .Descriptor.Name }}WithLoaders(ctx context.Context) context.Context {
	{{- range $loader :=.Loaders }}
	ctx = context.WithValue(ctx, "{{ $loader.Method }}Loader", dataloader.NewBatchedLoader(
		func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
			var results []*dataloader.Result

			{{ if $loader.Custom }}
			var requests []*{{ $loader.KeysType }}
			for _, key := range keys {
				requests = append(requests, key.(*{{ $loader.KeysType }}Key).{{ $loader.KeysType }})
			}
			var resp *{{ $loader.ResponseType }}
			var err error
			if {{ $.Descriptor.Name }}ServiceInstance != nil {
				resp, err = {{ $.Descriptor.Name }}ServiceInstance.{{ $loader.Method }}(ctx, &{{ $loader.RequestType }}{
					{{ $loader.KeysField }}: requests,
				})
			} else {
				if {{ $.Descriptor.Name }}ClientInstance == nil {
					{{ $.Descriptor.Name }}ClientInstance = get{{ $.Descriptor.Name }}Client()
				}
				resp, err = {{ $.Descriptor.Name }}ClientInstance.{{ $loader.Method }}(ctx, &{{ $loader.RequestType }}{
					{{ $loader.KeysField }}: requests,
				})
			}
			{{- else }}
			var resp *{{ $loader.ResponseType }}
			var err error
			if {{ $.Descriptor.Name }}ServiceInstance != nil {
				resp, err = {{ $.Descriptor.Name }}ServiceInstance.{{ $loader.Method }}(ctx, &{{ if eq $loader.RequestType "BatchRequest" }}pg.{{ end }}{{ $loader.RequestType }}{
					{{ $loader.KeysField }}: keys.Keys(),
				})
			} else {
				if {{ $.Descriptor.Name }}ClientInstance == nil {
					{{ $.Descriptor.Name }}ClientInstance = get{{ $.Descriptor.Name }}Client()
				}
				resp, err = {{ $.Descriptor.Name }}ClientInstance.{{ $loader.Method }}(ctx, &{{ if eq $loader.RequestType "BatchRequest" }}pg.{{ end }}{{ $loader.RequestType }}{
					{{ $loader.KeysField }}: keys.Keys(),
				})
			}
			{{- end }}

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
func {{ $loader.Method }}(p gql.ResolveParams, {{ if $loader.Custom }}key *{{ $loader.KeysType }}{{ else }}key string{{ end }}) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("{{ $loader.Method }}Loader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("{{ $loader.Method }}Loader").(*dataloader.Loader)
	default:
		panic("Please call {{ $.Package }}.WithLoaders with the current context first")
	}

	thunk := loader.Load(p.Context, {{ if $loader.Custom }}&{{ $loader.KeysType }}Key{key}{{ else }}dataloader.StringKey(key){{ end }})
	return func() (interface{}, error) {
				res, err := thunk()
				if err != nil {
					return nil, err
				}
				return res.({{ $loader.ResultsType }}), nil
	}, nil
}

func {{ $loader.Method }}Many(p gql.ResolveParams, {{ if $loader.Custom }}keys []*{{ $loader.KeysType }}{{ else }}keys []string{{ end }}) (func() (interface{}, error), error) {
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

	thunk := loader.LoadMany(p.Context, loaderKeys)
	{{ else }}
	thunk := loader.LoadMany(p.Context, dataloader.NewKeysFromStrings(keys))
	{{ end }}
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
	ResponseType string
	KeysField    string
	KeysType     string
	ResultsField string
	ResultsType  string
	Custom       bool
}
