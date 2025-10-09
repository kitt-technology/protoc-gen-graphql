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

const loaderAccessorTemplate = `
{{- range .Loaders }}
// {{ .MethodName }} loads a single {{ .ResultsType }} using the {{ .LowerServiceName }} service dataloader
func (m *{{ $.ModuleName }}) {{ .MethodName }}(p gql.ResolveParams, {{ if .Custom }}key *{{ .KeysType }}{{ else }}key string{{ end }}) (func() (interface{}, error), error) {
	return {{ .ServiceName }}{{ .Method }}(p, key)
}

// {{ .MethodNameMany }} loads multiple {{ .ResultsType }} using the {{ .LowerServiceName }} service dataloader
func (m *{{ $.ModuleName }}) {{ .MethodNameMany }}(p gql.ResolveParams, {{ if .Custom }}keys []*{{ .KeysType }}{{ else }}keys []string{{ end }}) (func() (interface{}, error), error) {
	return {{ .ServiceName }}{{ .Method }}Many(p, keys)
}
{{ end -}}
`

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

	return fmt.Sprintf(`
// %s implements the Module interface for the %s package (types only, no services)
type %s struct{}

// New%s creates a new module instance
func New%s() pg.Module {
	return &%s{}
}

// Fields returns an empty map (no services in this module)
func (m *%s) Fields() gql.Fields {
	return gql.Fields{}
}

// Messages returns all message types from this package
func (m *%s) Messages() []pg.GraphqlMessage {
	return allMessages
}

// WithLoaders returns the context unchanged (no loaders in this module)
func (m *%s) WithLoaders(ctx context.Context) context.Context {
	return ctx
}

// PackageName returns the proto package name
func (m *%s) PackageName() string {
	return %q
}
`,
		moduleName, pkgName,
		moduleName,
		moduleName, moduleName, moduleName,
		moduleName,
		moduleName,
		moduleName,
		moduleName, pkgName,
	)
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

		out += fmt.Sprintf("\t%sClient   %sClient\n", lowerServiceName, serviceName)
		out += fmt.Sprintf("\t%sService  %sServer\n", lowerServiceName, serviceName)
		out += fmt.Sprintf("\t%sDialOpts []grpc.DialOption\n\n", lowerServiceName)
	}

	out += "}\n\n"

	// 2. Generate option type
	out += fmt.Sprintf("// %sOption configures the %s\n", moduleName, moduleName)
	out += fmt.Sprintf("type %sOption func(*%s)\n\n", moduleName, moduleName)

	// 3. Generate option functions for each service (renamed to avoid conflicts with old Init functions)
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

		// WithXModuleDialOptions option
		out += fmt.Sprintf("// WithModule%sDialOptions sets dial options for lazy client creation for the %s service\n", serviceName, serviceName)
		out += fmt.Sprintf("func WithModule%sDialOptions(opts ...grpc.DialOption) %sOption {\n", serviceName, moduleName)
		out += fmt.Sprintf("\treturn func(m *%s) {\n", moduleName)
		out += fmt.Sprintf("\t\tm.%sDialOpts = opts\n", lowerServiceName)
		out += "\t}\n}\n\n"
	}

	// 4. Generate constructor (returns concrete type)
	out += fmt.Sprintf("// New%s creates a new module with optional service configurations\n", moduleName)
	out += fmt.Sprintf("func New%s(opts ...%sOption) *%s {\n", moduleName, moduleName, moduleName)
	out += fmt.Sprintf("\tm := &%s{}\n", moduleName)
	out += "\tfor _, opt := range opts {\n"
	out += "\t\topt(m)\n"
	out += "\t}\n"
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
		out += fmt.Sprintf("\t\thost := os.Getenv(%q)\n", strings.ToUpper(serviceName)+"_SERVICE_HOST")
		out += "\t\tif host == \"\" {\n"
		out += fmt.Sprintf("\t\t\thost = %q\n", dns)
		out += "\t\t}\n"
		out += fmt.Sprintf("\t\tm.%sClient = New%sClient(pg.GrpcConnection(host, m.%sDialOpts...))\n",
			lowerServiceName, serviceName, lowerServiceName)
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
			out += fmt.Sprintf("\tfields[\"%s\"] = &gql.Field{\n", fieldName)
			out += fmt.Sprintf("\t\tName: \"%s\",\n", fieldName)
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

			// Handle errors
			out += "\t\t\tif err != nil {\n"
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
	var out string

	out += fmt.Sprintf(`// Messages returns all message types from this package
func (m *%s) Messages() []pg.GraphqlMessage {
	return allMessages
}

// PackageName returns the proto package name
func (m *%s) PackageName() string {
	return %q
}
`,
		moduleName,
		moduleName, pkgName,
	)

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
	var out string

	for _, typedef := range f.TypeDefs {
		// Only generate for message types (not enums)
		msg, ok := typedef.(types.Message)
		if !ok {
			continue
		}

		typeName := msg.Descriptor.GetName()

		out += fmt.Sprintf("\n// AddFieldTo%s adds a custom field to the %s GraphQL type\n", typeName, typeName)
		out += fmt.Sprint("// This provides a type-safe way to extend types with cross-service relationships\n")
		out += fmt.Sprintf("func (m *%s) AddFieldTo%s(fieldName string, field *gql.Field) {\n", moduleName, typeName)
		out += fmt.Sprintf("\t%sGraphqlType.AddFieldConfig(fieldName, field)\n", typeName)
		out += "}\n"
	}

	return out
}

// generateTypeAccessors generates methods to access GraphQL types through the module
func (f File) generateTypeAccessors(moduleName string) string {
	var out string

	for _, typedef := range f.TypeDefs {
		// Only generate for message types (not enums)
		msg, ok := typedef.(types.Message)
		if !ok {
			continue
		}

		typeName := msg.Descriptor.GetName()

		out += fmt.Sprintf("\n// %sType returns the GraphQL type for %s\n", typeName, typeName)
		out += fmt.Sprintf("func (m *%s) %sType() *gql.Object {\n", moduleName, typeName)
		out += fmt.Sprintf("\treturn %sGraphqlType\n", typeName)
		out += "}\n"
	}

	return out
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
