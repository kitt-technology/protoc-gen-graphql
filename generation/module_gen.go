package generation

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/kitt-technology/protoc-gen-graphql/generation/templates"
	"github.com/kitt-technology/protoc-gen-graphql/generation/types"
)

const (
	loaderAccessorTemplate = `
{{- range .Loaders }}
// {{ .MethodName }} loads a single {{ .ResultsType }} using the {{ .LowerServiceName }} service dataloader
func (m *{{ $.ModuleName }}) {{ .MethodName }}(p gql.ResolveParams, {{ if .Custom }}key *{{ .KeysType }}{{ else }}key string{{ end }}) (func() (interface{}, error), error) {
	return {{ .Method }}(p, key)
}

// {{ .MethodNameMany }} loads multiple {{ .ResultsType }} using the {{ .LowerServiceName }} service dataloader
func (m *{{ $.ModuleName }}) {{ .MethodNameMany }}(p gql.ResolveParams, {{ if .Custom }}keys []*{{ .KeysType }}{{ else }}keys []string{{ end }}) (func() (interface{}, error), error) {
	return {{ .Method }}Many(p, keys)
}
{{ end -}}
`

	typeOnlyModuleTemplate = `
// {{ .ModuleName }} implements the Module interface for the {{ .PackageName }} package (types only, no services)
type {{ .ModuleName }} struct{}

// New{{ .ModuleName }} creates a new module instance
func New{{ .ModuleName }}() *{{ .ModuleName }} {
	return &{{ .ModuleName }}{}
}

// Fields returns an empty map (no services in this module)
func (m *{{ .ModuleName }}) Fields() gql.Fields {
	return gql.Fields{}
}

// Messages returns all message types from this package
func (m *{{ .ModuleName }}) Messages() []pg.GraphqlMessage {
	return allMessages
}

// WithLoaders returns the context unchanged (no loaders in this module)
func (m *{{ .ModuleName }}) WithLoaders(ctx context.Context) context.Context {
	return ctx
}

// PackageName returns the proto package name
func (m *{{ .ModuleName }}) PackageName() string {
	return {{ printf "%q" .PackageName }}
}
`

	moduleStructTemplate = `
// {{ .ModuleName }} implements the Module interface for the {{ .PackageName }} package
type {{ .ModuleName }} struct {
{{- range .Services }}
	{{ .LowerServiceName }}Client  {{ .ServiceName }}Client
	{{ .LowerServiceName }}Service {{ .ServiceName }}Server
{{- end }}

	dialOpts []grpc.DialOption
}
`

	moduleOptionTemplate = `
// {{ .ModuleName }}Option configures the {{ .ModuleName }}
type {{ .ModuleName }}Option func(*{{ .ModuleName }})
`

	serviceOptionTemplate = `
// WithModule{{ .ServiceName }}Client sets the gRPC client for the {{ .ServiceName }} service
func WithModule{{ .ServiceName }}Client(client {{ .ServiceName }}Client) {{ .ModuleName }}Option {
	return func(m *{{ .ModuleName }}) {
		m.{{ .LowerServiceName }}Client = client
	}
}

// WithModule{{ .ServiceName }}Service sets the direct service implementation for the {{ .ServiceName }} service
func WithModule{{ .ServiceName }}Service(service {{ .ServiceName }}Server) {{ .ModuleName }}Option {
	return func(m *{{ .ModuleName }}) {
		m.{{ .LowerServiceName }}Service = service
	}
}
`

	dialOptionsTemplate = `
// WithDialOptions sets dial options for lazy client creation for all services in this module
func WithDialOptions(opts ...grpc.DialOption) {{ .ModuleName }}Option {
	return func(m *{{ .ModuleName }}) {
		m.dialOpts = opts
	}
}
`

	constructorTemplate = `
// New{{ .ModuleName }} creates a new module with optional service configurations
func New{{ .ModuleName }}(opts ...{{ .ModuleName }}Option) *{{ .ModuleName }} {
	m := &{{ .ModuleName }}{}
	for _, opt := range opts {
		opt(m)
	}

	// Initialize ClientInstance variables for backward compatibility
{{- range .Services }}
	if m.dialOpts != nil || m.{{ .LowerServiceName }}Client != nil || m.{{ .LowerServiceName }}Service != nil {
		if m.{{ .LowerServiceName }}Client != nil {
			{{ .ServiceName }}ClientInstance = &{{ .LowerServiceName }}ClientAdapter{client: m.{{ .LowerServiceName }}Client}
		} else if m.{{ .LowerServiceName }}Service != nil {
			{{ .ServiceName }}ClientInstance = &{{ .LowerServiceName }}ServerAdapter{server: m.{{ .LowerServiceName }}Service}
		} else {
			{{ .ServiceName }}ClientInstance = &{{ .LowerServiceName }}ClientAdapter{client: m.get{{ .ServiceName }}Client()}
		}
	}
	if {{ .ServiceName }}ClientInstance == nil {
		{{ .ServiceName }}ClientInstance = &{{ .LowerServiceName }}ClientAdapter{client: m.get{{ .ServiceName }}Client()}
	}
{{- end }}
	return m
}
`

	getClientTemplate = `
// get{{ .ServiceName }}Client returns the client, creating it lazily if needed
func (m *{{ .ModuleName }}) get{{ .ServiceName }}Client() {{ .ServiceName }}Client {
	if m.{{ .LowerServiceName }}Client == nil {
		host := os.Getenv("GRPC_SERVER_HOST")
		if host == "" {
			host = {{ printf "%q" .Dns }}
		}
		m.{{ .LowerServiceName }}Client = New{{ .ServiceName }}Client(pg.GrpcConnection(host, m.dialOpts...))
	}
	return m.{{ .LowerServiceName }}Client
}
`

	basicMethodsTemplate = `
// Messages returns all message types from this package
func (m *{{ .ModuleName }}) Messages() []pg.GraphqlMessage {
	return allMessages
}

// PackageName returns the proto package name
func (m *{{ .ModuleName }}) PackageName() string {
	return {{ printf "%q" .PackageName }}
}
`

	fieldCustomizerTemplate = `
// AddFieldTo{{ .TypeName }} adds a custom field to the {{ .TypeName }} GraphQL type
func (m *{{ .ModuleName }}) AddFieldTo{{ .TypeName }}(fieldName string, field *gql.Field) {
	{{ .TypeName }}GraphqlType.AddFieldConfig(fieldName, field)
}
`

	typeAccessorTemplate = `
// {{ .TypeName }}Type returns the GraphQL type for {{ .TypeName }}
func (m *{{ .ModuleName }}) {{ .TypeName }}Type() *gql.Object {
	return {{ .TypeName }}GraphqlType
}
`

	serviceInstanceInterfaceTemplate = `
// {{ .ServiceName }}Instance is a unified interface for calling {{ .ServiceName }} methods
// It works with both gRPC clients and direct service implementations
type {{ .ServiceName }}Instance interface {
{{- range .Methods }}
	{{ .Name }}(ctx context.Context, req *{{ .InputType }}) (*{{ .Output }}, error)
{{- end }}
{{- range .Loaders }}
	{{ .Method }}(ctx context.Context, req *{{ .InputType }}) (*{{ .ResponseType }}, error)
	{{ .Method }}Batch(p gql.ResolveParams, key {{ .KeyType }}) (func() (interface{}, error), error)
	{{ .Method }}BatchMany(p gql.ResolveParams, keys {{ .KeysType }}) (func() (interface{}, error), error)
{{- end }}
}
`

	serverAdapterTemplate = `
type {{ .LowerServiceName }}ServerAdapter struct {
	server {{ .ServiceName }}Server
}
{{- range .Methods }}

func (a *{{ $.LowerServiceName }}ServerAdapter) {{ .Name }}(ctx context.Context, req *{{ .InputType }}) (*{{ .Output }}, error) {
	return a.server.{{ .Name }}(ctx, req)
}
{{- end }}
{{- range .Loaders }}

func (a *{{ $.LowerServiceName }}ServerAdapter) {{ .Method }}(ctx context.Context, req *{{ .InputType }}) (*{{ .ResponseType }}, error) {
	return a.server.{{ .Method }}(ctx, req)
}

func (a *{{ $.LowerServiceName }}ServerAdapter) {{ .Method }}Batch(p gql.ResolveParams, key {{ .KeyType }}) (func() (interface{}, error), error) {
	return {{ .Method }}(p, key)
}

func (a *{{ $.LowerServiceName }}ServerAdapter) {{ .Method }}BatchMany(p gql.ResolveParams, keys {{ .KeysType }}) (func() (interface{}, error), error) {
	return {{ .Method }}Many(p, keys)
}
{{- end }}
`

	clientAdapterTemplate = `
type {{ .LowerServiceName }}ClientAdapter struct {
	client {{ .ServiceName }}Client
}
{{- range .Methods }}

func (a *{{ $.LowerServiceName }}ClientAdapter) {{ .Name }}(ctx context.Context, req *{{ .InputType }}) (*{{ .Output }}, error) {
	return a.client.{{ .Name }}(ctx, req)
}
{{- end }}
{{- range .Loaders }}

func (a *{{ $.LowerServiceName }}ClientAdapter) {{ .Method }}(ctx context.Context, req *{{ .InputType }}) (*{{ .ResponseType }}, error) {
	return a.client.{{ .Method }}(ctx, req)
}

func (a *{{ $.LowerServiceName }}ClientAdapter) {{ .Method }}Batch(p gql.ResolveParams, key {{ .KeyType }}) (func() (interface{}, error), error) {
	return {{ .Method }}(p, key)
}

func (a *{{ $.LowerServiceName }}ClientAdapter) {{ .Method }}BatchMany(p gql.ResolveParams, keys {{ .KeysType }}) (func() (interface{}, error), error) {
	return {{ .Method }}Many(p, keys)
}
{{- end }}
`

	serviceGetterTemplate = `
// {{ .ServiceName }} returns a unified {{ .ServiceName }}Instance that works with both clients and services
// Returns nil if neither client nor service is configured
func (m *{{ .ModuleName }}) {{ .ServiceName }}() {{ .ServiceName }}Instance {
	if m.{{ .LowerServiceName }}Client != nil {
		return &{{ .LowerServiceName }}ClientAdapter{client: m.{{ .LowerServiceName }}Client}
	}
	if m.{{ .LowerServiceName }}Service != nil {
		return &{{ .LowerServiceName }}ServerAdapter{server: m.{{ .LowerServiceName }}Service}
	}
	return &{{ .LowerServiceName }}ClientAdapter{client: m.get{{ .ServiceName }}Client()}
}
`
)

type loaderAccessorData struct {
	ModuleName string
	Loaders    []loaderMethodData
}

type loaderMethodData struct {
	MethodName       string
	MethodNameMany   string
	ServiceName      string
	LowerServiceName string
	Method           string
	ResultsType      string
	KeysType         string
	Custom           bool
}

// generateTypeOnlyModule generates a module for proto files with no services (just types)
func (f File) generateTypeOnlyModule() string {
	moduleName := strcase.ToCamel(string(f.Package)) + "Module"
	pkgName := string(f.Package)

	tmpl := template.Must(template.New("typeOnlyModule").Parse(typeOnlyModuleTemplate))

	var buf bytes.Buffer
	data := struct {
		ModuleName  string
		PackageName string
	}{
		ModuleName:  moduleName,
		PackageName: pkgName,
	}

	if err := tmpl.Execute(&buf, data); err != nil {
		panic(fmt.Sprintf("failed to execute type-only module template: %v", err))
	}

	return buf.String()
}

// generateServiceModule generates a module for proto files with one or more services
func (f File) generateServiceModule(services []templates.Message) string {
	var out string

	moduleName := strcase.ToCamel(string(f.Package)) + "Module"
	pkgName := string(f.Package)

	// 1. Generate module struct with fields for each service
	out += fmt.Sprintf("\n// %s implements the Module interface for the %s package\n", moduleName, pkgName)
	out += fmt.Sprintf("type %s struct {\n", moduleName)

	for _, svc := range services {
		serviceName := svc.Descriptor.GetName()
		lowerServiceName := strcase.ToLowerCamel(serviceName)

		out += fmt.Sprintf("\t%sClient  %sClient\n", lowerServiceName, serviceName)
		out += fmt.Sprintf("\t%sService %sServer\n", lowerServiceName, serviceName)
	}

	// Add single dialOpts field
	out += "\n\tdialOpts []grpc.DialOption\n"
	out += "}\n\n"

	// 2. Generate option type
	out += fmt.Sprintf("// %sOption configures the %s\n", moduleName, moduleName)
	out += fmt.Sprintf("type %sOption func(*%s)\n\n", moduleName, moduleName)

	// 3. Generate option functions for each service
	for _, svc := range services {
		serviceName := svc.Descriptor.GetName()
		lowerServiceName := strcase.ToLowerCamel(serviceName)

		// WithXModuleClient option
		out += fmt.Sprintf("// WithModule%sClient sets the gRPC client for the %s service\n", serviceName, serviceName)
		out += fmt.Sprintf("func WithModule%sClient(client %sClient) %sOption {\n", serviceName, serviceName, moduleName)
		out += fmt.Sprintf("\treturn func(m *%s) {\n", moduleName)
		out += fmt.Sprintf("\t\tm.%sClient = client\n", lowerServiceName)
		out += "\t}\n}\n\n"

		// WithXModuleService option
		out += fmt.Sprintf("// WithModule%sService sets the direct service implementation for the %s service\n", serviceName, serviceName)
		out += fmt.Sprintf("func WithModule%sService(service %sServer) %sOption {\n", serviceName, serviceName, moduleName)
		out += fmt.Sprintf("\treturn func(m *%s) {\n", moduleName)
		out += fmt.Sprintf("\t\tm.%sService = service\n", lowerServiceName)
		out += "\t}\n}\n\n"
	}

	// Generate single WithDialOptions function that accepts ...grpc.DialOption
	out += "// WithDialOptions sets dial options for lazy client creation for all services in this module\n"
	out += fmt.Sprintf("func WithDialOptions(opts ...grpc.DialOption) %sOption {\n", moduleName)
	out += fmt.Sprintf("\treturn func(m *%s) {\n", moduleName)
	out += "\t\tm.dialOpts = opts\n"
	out += "\t}\n}\n\n"

	// 4. Generate constructor (returns concrete type)
	out += fmt.Sprintf("// New%s creates a new module with optional service configurations\n", moduleName)
	out += fmt.Sprintf("func New%s(opts ...%sOption) *%s {\n", moduleName, moduleName, moduleName)
	out += fmt.Sprintf("\tm := &%s{}\n", moduleName)
	out += "\tfor _, opt := range opts {\n"
	out += "\t\topt(m)\n"
	out += "\t}\n\n"

	// Initialize ClientInstance variables if module has dial options
	out += "\t// Initialize ClientInstance variables for backward compatibility\n"
	for _, svc := range services {
		serviceName := svc.Descriptor.GetName()
		lowerServiceName := strcase.ToLowerCamel(serviceName)
		out += fmt.Sprintf("\tif m.dialOpts != nil || m.%sClient != nil || m.%sService != nil {\n", lowerServiceName, lowerServiceName)
		out += fmt.Sprintf("\t\tif m.%sClient != nil {\n", lowerServiceName)
		out += fmt.Sprintf("\t\t\t%sClientInstance = &%sClientAdapter{client: m.%sClient}\n", serviceName, lowerServiceName, lowerServiceName)
		out += fmt.Sprintf("\t\t} else if m.%sService != nil {\n", lowerServiceName)
		out += fmt.Sprintf("\t\t\t%sClientInstance = &%sServerAdapter{server: m.%sService}\n", serviceName, lowerServiceName, lowerServiceName)
		out += "\t\t} else {\n"
		out += fmt.Sprintf("\t\t\t%sClientInstance = &%sClientAdapter{client: m.get%sClient()}\n", serviceName, lowerServiceName, serviceName)
		out += "\t\t}\n"
		out += "\t}\n"
		out += fmt.Sprintf("\tif %sClientInstance == nil {\n", serviceName)
		out += fmt.Sprintf("\t\t%sClientInstance = &%sClientAdapter{client: m.get%sClient()}\n", serviceName, lowerServiceName, serviceName)
		out += fmt.Sprintf("\t}\n")

	}

	out += "\treturn m\n"
	out += "}\n\n"

	// 5. Generate getClient methods for each service (lazy loading)
	for _, svc := range services {
		serviceName := svc.Descriptor.GetName()
		lowerServiceName := strcase.ToLowerCamel(serviceName)
		dns := svc.Dns
		if dns == "" {
			dns = strings.ToLower(serviceName)
		}

		out += fmt.Sprintf("// get%sClient returns the client, creating it lazily if needed\n", serviceName)
		out += fmt.Sprintf("func (m *%s) get%sClient() %sClient {\n", moduleName, serviceName, serviceName)
		out += fmt.Sprintf("\tif m.%sClient == nil {\n", lowerServiceName)
		out += "\t\thost := os.Getenv(\"GRPC_SERVER_HOST\")\n"
		out += "\t\tif host == \"\" {\n"
		out += fmt.Sprintf("\t\t\thost = %q\n", dns)
		out += "\t\t}\n"
		out += fmt.Sprintf("\t\tm.%sClient = New%sClient(pg.GrpcConnection(host, m.dialOpts...))\n",
			lowerServiceName, serviceName)
		out += "\t}\n"
		out += fmt.Sprintf("\treturn m.%sClient\n", lowerServiceName)
		out += "}\n\n"
	}

	// 6. Generate Fields() method
	out += f.generateFieldsMethod(moduleName, services)

	// 7. Generate WithLoaders() method
	out += f.generateWithLoadersMethod(moduleName, services)

	// 8. Generate Messages() and PackageName() methods
	out += f.generateBasicMethods(moduleName, pkgName, services)

	// 9. Generate service instance accessor methods
	out += f.generateServiceAccessors(moduleName, services)

	// 10. Generate backward compatibility layer (deprecated)
	out += f.generateBackwardCompatLayer(moduleName, services)

	return out
}

// generateFieldsMethod generates the Fields() method that returns all GraphQL fields from all services
func (f File) generateFieldsMethod(moduleName string, services []templates.Message) string {
	var out string

	out += "// Fields returns all GraphQL query/mutation fields from all services in this module\n"
	out += fmt.Sprintf("func (m *%s) Fields() gql.Fields {\n", moduleName)
	out += "\tfields := gql.Fields{}\n\n"

	// Add fields from each service
	for _, svc := range services {
		serviceName := svc.Descriptor.GetName()
		lowerServiceName := strcase.ToLowerCamel(serviceName)
		servicePackage := svc.ServiceName

		// Generate fields for each method in this service
		for _, method := range svc.Methods {
			fieldName := fmt.Sprintf("%s_%s", servicePackage, method.Name)

			out += fmt.Sprintf("\t// %s service: %s method\n", serviceName, method.Name)
			out += fmt.Sprintf("\tfields[%q] = &gql.Field{\n", fieldName)
			out += fmt.Sprintf("\t\tName: %q,\n", fieldName)
			out += fmt.Sprintf("\t\tType: %sGraphqlType,\n", method.Output)
			if method.Input == "BatchRequest" {
				out += fmt.Sprintf("\t\tArgs: pg.%sGraphqlArgs,\n", method.Input)
			} else {
				out += fmt.Sprintf("\t\tArgs: %sGraphqlArgs,\n", method.Input)
			}
			out += "\t\tResolve: func(p gql.ResolveParams) (interface{}, error) {\n"

			// Parse args
			if method.Input == "BatchRequest" {
				out += fmt.Sprintf("\t\t\treq := pg.%sFromArgs(p.Args)\n", method.Input)
			} else {
				out += fmt.Sprintf("\t\t\treq := %sFromArgs(p.Args)\n", method.Input)
			}

			// Prefer service (direct) over client (gRPC)
			out += fmt.Sprintf("\t\t\tif m.%sService != nil {\n", lowerServiceName)
			out += fmt.Sprintf("\t\t\t\treturn m.%sService.%s(p.Context, req)\n", lowerServiceName, method.Name)
			out += "\t\t\t}\n"
			out += fmt.Sprintf("\t\t\treturn m.get%sClient().%s(p.Context, req)\n", serviceName, method.Name)
			out += "\t\t},\n"
			out += "\t}\n\n"
		}
	}

	out += "\treturn fields\n"
	out += "}\n\n"

	return out
}

// generateWithLoadersMethod generates the WithLoaders() method for all services
func (f File) generateWithLoadersMethod(moduleName string, services []templates.Message) string {
	var out string

	// Check if any service has loaders
	hasLoaders := false
	for _, svc := range services {
		if len(svc.Loaders) > 0 {
			hasLoaders = true
			break
		}
	}

	out += "// WithLoaders registers all dataloaders from all services into the context\n"
	out += fmt.Sprintf("func (m *%s) WithLoaders(ctx context.Context) context.Context {\n", moduleName)

	if !hasLoaders {
		out += "\t// No loaders in this module\n"
		out += "\treturn ctx\n"
		out += "}\n\n"
		return out
	}

	// Generate loaders for each service
	for _, svc := range services {
		if len(svc.Loaders) == 0 {
			continue
		}

		serviceName := svc.Descriptor.GetName()
		lowerServiceName := strcase.ToLowerCamel(serviceName)

		for _, loader := range svc.Loaders {
			out += fmt.Sprintf("\t// %s service: %s loader\n", serviceName, loader.Method)
			out += fmt.Sprintf("\tctx = context.WithValue(ctx, %q, dataloader.NewBatchedLoader(\n", loader.Method+"Loader")
			out += "\t\tfunc(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {\n"
			out += "\t\t\tvar results []*dataloader.Result\n\n"

			// Build the request
			if loader.Custom {
				// Custom key type
				out += fmt.Sprintf("\t\t\tvar requests []*%s\n", loader.KeysType)
				out += "\t\t\tfor _, key := range keys {\n"
				out += fmt.Sprintf("\t\t\t\trequests = append(requests, key.(*%sKey).%s)\n", loader.KeysType, loader.KeysType)
				out += "\t\t\t}\n"
			}

			// Make the request
			out += fmt.Sprintf("\t\t\tvar resp *%s\n", loader.ResponseType)
			out += "\t\t\tvar err error\n\n"

			// Request construction
			if loader.Custom {
				out += fmt.Sprintf("\t\t\treq := &%s{\n", loader.RequestType)
				out += fmt.Sprintf("\t\t\t\t%s: requests,\n", loader.KeysField)
				out += "\t\t\t}\n"
			} else {
				if loader.RequestType == "BatchRequest" {
					out += "\t\t\treq := &pg.BatchRequest{\n"
				} else {
					out += fmt.Sprintf("\t\t\treq := &%s{\n", loader.RequestType)
				}
				out += fmt.Sprintf("\t\t\t\t%s: keys.Keys(),\n", loader.KeysField)
				out += "\t\t\t}\n"
			}

			// Call service or client
			out += fmt.Sprintf("\t\t\tif m.%sService != nil {\n", lowerServiceName)
			out += fmt.Sprintf("\t\t\t\tresp, err = m.%sService.%s(ctx, req)\n", lowerServiceName, loader.Method)
			out += "\t\t\t} else {\n"
			out += fmt.Sprintf("\t\t\t\tresp, err = m.get%sClient().%s(ctx, req)\n", serviceName, loader.Method)
			out += "\t\t\t}\n\n"

			// Handle errors - must return error results for ALL keys
			out += "\t\t\tif err != nil {\n"
			out += "\t\t\t\t// Return error result for each key - dataloader requires same number of results as keys\n"
			out += "\t\t\t\tfor range keys {\n"
			out += "\t\t\t\t\tresults = append(results, &dataloader.Result{Error: err})\n"
			out += "\t\t\t\t}\n"
			out += "\t\t\t\treturn results\n"
			out += "\t\t\t}\n\n"

			// Build results
			out += "\t\t\tfor _, key := range keys.Keys() {\n"
			out += fmt.Sprintf("\t\t\t\tif val, ok := resp.%s[key]; ok {\n", loader.ResultsField)
			out += "\t\t\t\t\tresults = append(results, &dataloader.Result{Data: val})\n"
			out += "\t\t\t\t} else {\n"
			out += fmt.Sprintf("\t\t\t\t\tvar empty %s\n", loader.ResultsType)
			out += "\t\t\t\t\tresults = append(results, &dataloader.Result{Data: empty})\n"
			out += "\t\t\t\t}\n"
			out += "\t\t\t}\n\n"

			out += "\t\t\treturn results\n"
			out += "\t\t},\n"
			out += "\t))\n\n"
		}
	}

	out += "\treturn ctx\n"
	out += "}\n\n"

	return out
}

// generateBasicMethods generates Messages() and PackageName() methods
func (f File) generateBasicMethods(moduleName, pkgName string, services []templates.Message) string {
	tmpl := template.Must(template.New("basicMethods").Parse(basicMethodsTemplate))

	var buf bytes.Buffer
	data := struct {
		ModuleName  string
		PackageName string
	}{
		ModuleName:  moduleName,
		PackageName: pkgName,
	}

	if err := tmpl.Execute(&buf, data); err != nil {
		panic(fmt.Sprintf("failed to execute basic methods template: %v", err))
	}

	out := buf.String()

	// Generate type-safe field customization methods for each message type
	out += "\n// Type-safe field customization methods\n"
	out += f.generateFieldCustomizers(moduleName)

	// Generate type accessor methods
	out += "\n// GraphQL type accessors\n"
	out += f.generateTypeAccessors(moduleName)

	// Generate loader accessor methods
	out += "\n// DataLoader accessor methods\n"
	out += f.generateLoaderAccessors(moduleName, services)

	return out
}

// generateFieldCustomizers generates type-safe methods to add fields on message types
func (f File) generateFieldCustomizers(moduleName string) string {
	tmpl := template.Must(template.New("fieldCustomizer").Parse(fieldCustomizerTemplate))

	var buf bytes.Buffer
	for _, typedef := range f.TypeDefs {
		// Only generate for message types (not enums)
		msg, ok := typedef.(types.Message)
		if !ok {
			continue
		}

		typeName := msg.Descriptor.GetName()
		data := struct {
			ModuleName string
			TypeName   string
		}{
			ModuleName: moduleName,
			TypeName:   typeName,
		}

		if err := tmpl.Execute(&buf, data); err != nil {
			panic(fmt.Sprintf("failed to execute field customizer template: %v", err))
		}
	}

	return buf.String()
}

// generateTypeAccessors generates methods to access GraphQL types through the module
func (f File) generateTypeAccessors(moduleName string) string {
	tmpl := template.Must(template.New("typeAccessor").Parse(typeAccessorTemplate))

	var buf bytes.Buffer
	for _, typedef := range f.TypeDefs {
		// Only generate for message types (not enums)
		msg, ok := typedef.(types.Message)
		if !ok {
			continue
		}

		typeName := msg.Descriptor.GetName()
		data := struct {
			ModuleName string
			TypeName   string
		}{
			ModuleName: moduleName,
			TypeName:   typeName,
		}

		if err := tmpl.Execute(&buf, data); err != nil {
			panic(fmt.Sprintf("failed to execute type accessor template: %v", err))
		}
	}

	return buf.String()
}

// generateLoaderAccessors generates methods to access loader functions through the module
func (f File) generateLoaderAccessors(moduleName string, services []templates.Message) string {
	// Collect all loaders from all services
	var loaders []loaderMethodData

	for _, svc := range services {
		serviceName := svc.Descriptor.GetName()
		lowerServiceName := strcase.ToLowerCamel(serviceName)

		for _, loader := range svc.Loaders {
			// Prefix with service name to avoid collisions when multiple services have loaders with same name
			// e.g., "Users" service + "LoadUsers" method = "UsersLoadUsers"
			methodName := serviceName + loader.Method
			methodNameMany := methodName + "Many"

			loaders = append(loaders, loaderMethodData{
				MethodName:       methodName,
				MethodNameMany:   methodNameMany,
				ServiceName:      serviceName,
				LowerServiceName: lowerServiceName,
				Method:           loader.Method,
				ResultsType:      loader.ResultsType,
				KeysType:         loader.KeysType,
				Custom:           loader.Custom,
			})
		}
	}

	if len(loaders) == 0 {
		return ""
	}

	// Execute template
	tmpl, err := template.New("loaderAccessors").Parse(loaderAccessorTemplate)
	if err != nil {
		panic(fmt.Sprintf("failed to parse loader accessor template: %v", err))
	}

	var buf bytes.Buffer
	data := loaderAccessorData{
		ModuleName: moduleName,
		Loaders:    loaders,
	}

	if err := tmpl.Execute(&buf, data); err != nil {
		panic(fmt.Sprintf("failed to execute loader accessor template: %v", err))
	}

	return buf.String()
}

// generateServiceAccessors generates methods to access the underlying client and service instances
func (f File) generateServiceAccessors(moduleName string, services []templates.Message) string {
	var out string

	out += "\n// Service instance accessors\n"

	for _, svc := range services {
		serviceName := svc.Descriptor.GetName()
		lowerServiceName := strcase.ToLowerCamel(serviceName)

		// Generate a unified interface that wraps both client and server
		out += fmt.Sprintf("\n// %sInstance is a unified interface for calling %s methods\n", serviceName, serviceName)
		out += "// It works with both gRPC clients and direct service implementations\n"
		out += fmt.Sprintf("type %sInstance interface {\n", serviceName)

		// Add methods from the service (normal RPC methods)
		for _, method := range svc.Methods {
			inputType := method.Input
			if inputType == "BatchRequest" {
				inputType = "pg.BatchRequest"
			}
			out += fmt.Sprintf("\t%s(ctx context.Context, req *%s) (*%s, error)\n",
				method.Name, inputType, method.Output)
		}

		// Add batch loader methods with BOTH normal RPC signature and batch signature
		for _, loader := range svc.Loaders {
			// Add normal RPC signature for the loader method
			inputType := loader.RequestType
			if inputType == "BatchRequest" {
				inputType = "pg.BatchRequest"
			}
			out += fmt.Sprintf("\t%s(ctx context.Context, req *%s) (*%s, error)\n",
				loader.Method, inputType, loader.ResponseType)

			// Single item batch loader (suffixed with "Batch")
			keyType := "string"
			if loader.Custom {
				keyType = "*" + loader.KeysType
			}
			out += fmt.Sprintf("\t%sBatch(p gql.ResolveParams, key %s) (func() (interface{}, error), error)\n",
				loader.Method, keyType)

			// Many items batch loader (suffixed with "BatchMany")
			keysType := "[]string"
			if loader.Custom {
				keysType = "[]*" + loader.KeysType
			}
			out += fmt.Sprintf("\t%sBatchMany(p gql.ResolveParams, keys %s) (func() (interface{}, error), error)\n",
				loader.Method, keysType)
		}

		out += "}\n"

		// Generate adapter for server to match the interface
		out += fmt.Sprintf("\ntype %sServerAdapter struct {\n", lowerServiceName)
		out += fmt.Sprintf("\tserver %sServer\n", serviceName)
		out += "}\n"

		// Implement interface methods for the adapter
		for _, method := range svc.Methods {
			inputType := method.Input
			if inputType == "BatchRequest" {
				inputType = "pg.BatchRequest"
			}
			out += fmt.Sprintf("\nfunc (a *%sServerAdapter) %s(ctx context.Context, req *%s) (*%s, error) {\n",
				lowerServiceName, method.Name, inputType, method.Output)
			out += fmt.Sprintf("\treturn a.server.%s(ctx, req)\n", method.Name)
			out += "}\n"
		}

		// Implement batch loader methods for the server adapter
		for _, loader := range svc.Loaders {
			// Normal RPC signature for loader method
			inputType := loader.RequestType
			if inputType == "BatchRequest" {
				inputType = "pg.BatchRequest"
			}
			out += fmt.Sprintf("\nfunc (a *%sServerAdapter) %s(ctx context.Context, req *%s) (*%s, error) {\n",
				lowerServiceName, loader.Method, inputType, loader.ResponseType)
			out += fmt.Sprintf("\treturn a.server.%s(ctx, req)\n", loader.Method)
			out += "}\n"

			// Single item batch loader (suffixed with "Batch")
			keyType := "string"
			if loader.Custom {
				keyType = "*" + loader.KeysType
			}
			out += fmt.Sprintf("\nfunc (a *%sServerAdapter) %sBatch(p gql.ResolveParams, key %s) (func() (interface{}, error), error) {\n",
				lowerServiceName, loader.Method, keyType)
			out += fmt.Sprintf("\treturn %s(p, key)\n", loader.Method)
			out += "}\n"

			// Many items batch loader (suffixed with "BatchMany")
			keysType := "[]string"
			if loader.Custom {
				keysType = "[]*" + loader.KeysType
			}
			out += fmt.Sprintf("\nfunc (a *%sServerAdapter) %sBatchMany(p gql.ResolveParams, keys %s) (func() (interface{}, error), error) {\n",
				lowerServiceName, loader.Method, keysType)
			out += fmt.Sprintf("\treturn %sMany(p, keys)\n", loader.Method)
			out += "}\n"
		}

		// Generate adapter for client to match the interface
		out += fmt.Sprintf("\ntype %sClientAdapter struct {\n", lowerServiceName)
		out += fmt.Sprintf("\tclient %sClient\n", serviceName)
		out += "}\n"

		// Implement interface methods for the client adapter
		for _, method := range svc.Methods {
			inputType := method.Input
			if inputType == "BatchRequest" {
				inputType = "pg.BatchRequest"
			}
			out += fmt.Sprintf("\nfunc (a *%sClientAdapter) %s(ctx context.Context, req *%s) (*%s, error) {\n",
				lowerServiceName, method.Name, inputType, method.Output)
			out += fmt.Sprintf("\treturn a.client.%s(ctx, req)\n", method.Name)
			out += "}\n"
		}

		// Implement batch loader methods for the client adapter
		for _, loader := range svc.Loaders {
			// Normal RPC signature for loader method
			inputType := loader.RequestType
			if inputType == "BatchRequest" {
				inputType = "pg.BatchRequest"
			}
			out += fmt.Sprintf("\nfunc (a *%sClientAdapter) %s(ctx context.Context, req *%s) (*%s, error) {\n",
				lowerServiceName, loader.Method, inputType, loader.ResponseType)
			out += fmt.Sprintf("\treturn a.client.%s(ctx, req)\n", loader.Method)
			out += "}\n"

			// Single item batch loader (suffixed with "Batch")
			keyType := "string"
			if loader.Custom {
				keyType = "*" + loader.KeysType
			}
			out += fmt.Sprintf("\nfunc (a *%sClientAdapter) %sBatch(p gql.ResolveParams, key %s) (func() (interface{}, error), error) {\n",
				lowerServiceName, loader.Method, keyType)
			out += fmt.Sprintf("\treturn %s(p, key)\n", loader.Method)
			out += "}\n"

			// Many items batch loader (suffixed with "BatchMany")
			keysType := "[]string"
			if loader.Custom {
				keysType = "[]*" + loader.KeysType
			}
			out += fmt.Sprintf("\nfunc (a *%sClientAdapter) %sBatchMany(p gql.ResolveParams, keys %s) (func() (interface{}, error), error) {\n",
				lowerServiceName, loader.Method, keysType)
			out += fmt.Sprintf("\treturn %sMany(p, keys)\n", loader.Method)
			out += "}\n"
		}

		// Generate the getter that returns the unified interface
		out += fmt.Sprintf("\n// %s returns a unified %sInstance that works with both clients and services\n", serviceName, serviceName)
		out += "// Returns nil if neither client nor service is configured\n"
		out += fmt.Sprintf("func (m *%s) %s() %sInstance {\n", moduleName, serviceName, serviceName)
		out += fmt.Sprintf("\tif m.%sClient != nil {\n", lowerServiceName)
		out += fmt.Sprintf("\t\treturn &%sClientAdapter{client: m.%sClient}\n", lowerServiceName, lowerServiceName)
		out += "\t}\n"
		out += fmt.Sprintf("\tif m.%sService != nil {\n", lowerServiceName)
		out += fmt.Sprintf("\t\treturn &%sServerAdapter{server: m.%sService}\n", lowerServiceName, lowerServiceName)
		out += "\t}\n"
		out += fmt.Sprintf("\treturn &%sClientAdapter{client: m.get%sClient()}", lowerServiceName, serviceName)
		out += "}\n"
	}

	return out
}

// generateBackwardCompatLayer generates deprecated global functions for backward compatibility with v0.51.7
func (f File) generateBackwardCompatLayer(moduleName string, services []templates.Message) string {
	var out string

	// Generate package-level default module instance
	out += "// Backward compatibility layer for v0.51.7 API\n"
	out += "// All functions below are deprecated and will be removed in a future version.\n"
	out += "// Please migrate to the module-based API using New" + moduleName + "()\n\n"

	out += "var defaultModule *" + moduleName + "\n\n"

	// Generate root-level ClientInstance variables for each service
	for _, svc := range services {
		serviceName := svc.Descriptor.GetName()
		out += fmt.Sprintf("// %sClientInstance provides a unified %s client interface\n", serviceName, serviceName)
		out += "// Deprecated: Use New" + moduleName + "()." + serviceName + "() instead\n"
		out += fmt.Sprintf("var %sClientInstance %sInstance\n\n", serviceName, serviceName)
	}

	// Generate function to set default module
	out += "// SetDefaultModule allows you to set a custom module instance as the default\n"
	out += "// for use with deprecated package-level functions.\n"
	out += "// This allows you to configure a module once and have all deprecated functions use it.\n"
	out += "// Example:\n"
	out += "//   module := New" + moduleName + "(WithDialOptions(...))\n"
	out += "//   SetDefaultModule(module)\n"
	out += "//   // Now all deprecated Init(), WithLoaders(), etc. will use your module\n"
	out += fmt.Sprintf("func SetDefaultModule(module *%s) {\n", moduleName)
	out += "\tdefaultModule = module\n"
	out += "}\n\n"

	// Generate getter for default module (lazy initialization)
	out += "func getDefaultModule() *" + moduleName + " {\n"
	out += "\tif defaultModule == nil {\n"
	out += "\t\tdefaultModule = New" + moduleName + "()\n"
	out += "\t}\n"
	out += "\treturn defaultModule\n"
	out += "}\n\n"

	// Generate deprecated Init() function for each service
	for _, svc := range services {
		serviceName := svc.Descriptor.GetName()

		out += fmt.Sprintf("// %sInit initializes the %s service.\n", serviceName, serviceName)
		out += "// Deprecated: Use New" + moduleName + "() and configure with WithModule" + serviceName + "Client() or WithModule" + serviceName + "Service() instead.\n"
		out += fmt.Sprintf("func %sInit(ctx context.Context, opts ...%sOption) (context.Context, []*gql.Field) {\n", serviceName, moduleName)
		out += "\t// Apply options to default module\n"
		out += "\tm := getDefaultModule()\n"
		out += "\tfor _, opt := range opts {\n"
		out += "\t\topt(m)\n"
		out += "\t}\n\n"

		out += "\t// Get fields from the module\n"
		out += "\tfields := m.Fields()\n\n"

		out += "\t// Register loaders in context\n"
		out += "\tctx = m.WithLoaders(ctx)\n\n"

		out += "\t// Convert fields map to slice for this service only\n"
		out += "\tvar serviceFields []*gql.Field\n"
		out += fmt.Sprintf("\tservicePrefix := %q\n", svc.ServiceName+"_")
		out += "\t// Sort field names for deterministic order\n"
		out += "\tvar fieldNames []string\n"
		out += "\tfor name := range fields {\n"
		out += "\t\tfieldNames = append(fieldNames, name)\n"
		out += "\t}\n"
		out += "\tsort.Strings(fieldNames)\n"
		out += "\tfor _, name := range fieldNames {\n"
		out += "\t\tfield := fields[name]\n"
		out += "\t\tif strings.HasPrefix(name, servicePrefix) {\n"
		out += "\t\t\tserviceFields = append(serviceFields, field)\n"
		out += "\t\t}\n"
		out += "\t}\n\n"

		out += "\treturn ctx, serviceFields\n"
		out += "}\n\n"

		// Generate deprecated WithLoaders() function if this service has loaders
		if len(svc.Loaders) > 0 {
			out += fmt.Sprintf("// %sWithLoaders registers dataloaders for the %s service into the context.\n", serviceName, serviceName)
			out += "// Deprecated: Use New" + moduleName + "().WithLoaders(ctx) instead.\n"
			out += fmt.Sprintf("func %sWithLoaders(ctx context.Context) context.Context {\n", serviceName)
			out += "\treturn getDefaultModule().WithLoaders(ctx)\n"
			out += "}\n\n"
		}
	}

	// Generate deprecated global WithLoaders() function
	out += "// WithLoaders registers all dataloaders from all services into the context.\n"
	out += "// Deprecated: Use New" + moduleName + "().WithLoaders(ctx) instead.\n"
	out += "func WithLoaders(ctx context.Context) context.Context {\n"
	out += "\treturn getDefaultModule().WithLoaders(ctx)\n"
	out += "}\n\n"

	// Generate deprecated global Fields() function
	out += "// Fields returns all GraphQL query/mutation fields from all services.\n"
	out += "// Deprecated: Use New" + moduleName + "().Fields() instead.\n"
	out += "func Fields() gql.Fields {\n"
	out += "\treturn getDefaultModule().Fields()\n"
	out += "}\n\n"

	return out
}
