package generation

import (
	"bytes"
	"sort"
	"strings"
	"text/template"

	"github.com/kitt-technology/protoc-gen-graphql/generation/imports"
	"github.com/kitt-technology/protoc-gen-graphql/generation/templates"
	"github.com/kitt-technology/protoc-gen-graphql/generation/types"
	"github.com/kitt-technology/protoc-gen-graphql/generation/types/enum"
	"github.com/kitt-technology/protoc-gen-graphql/generation/util"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

const fileTpl = `
package {{ .Package }}

import (
	gql "github.com/graphql-go/graphql"
	{{- range .Imports }}
	{{ if has_alias . }}{{ . }}{{else}}"{{ . }}"{{end}}{{ end }}
)
`

var GraphqlImportMap = make(map[string]types.GraphqlImport, 0)
var SkippedMessages = make(map[string]bool, 0) // Tracks skipped messages by fully qualified name (e.g., "common.DateTime")

type Message interface {
	Generate() string
	Imports() []string
}

type File struct {
	Package   protogen.GoPackageName
	Message   []Message
	TypeDefs  []Message
	Imports   []string
	ImportMap map[string]string
}

func New(file *protogen.File) (f File) {
	f.Package = file.GoPackageName

	hasServices := false
	hasLoaders := false
	for _, service := range file.Proto.Service {
		svc := templates.New(service, file.Proto)
		f.Message = append(f.Message, svc)
		hasServices = true
		if len(svc.Loaders) > 0 {
			hasLoaders = true
		}
	}

	// Add imports needed for all generated files with services or just types
	// These are always needed because we always generate a module
	f.Imports = append(f.Imports, imports.PggImport, imports.ContextImport)

	// Only add these if we have services (modules with services need them)
	if hasServices {
		f.Imports = append(f.Imports, imports.StringsImport)
	}

	// Only add dataloader if we have services with batch loaders
	if hasLoaders {
		f.Imports = append(f.Imports, imports.DataloaderImport)
	}

	for _, e := range file.Proto.EnumType {
		f.TypeDefs = append(f.TypeDefs, enum.New(e))
	}

	if proto.HasExtension(file.Proto.Options, graphql.E_Package) {
		importPath, gqlPkg, ok := util.ParseGraphqlPackage(file.Proto)
		if !ok {
			panic("invalid graphql.package: " + file.Proto.GetName())
		}
		// Using graphql.package, could fall back to go_package?
		GraphqlImportMap[*file.Proto.Package] = types.GraphqlImport{
			ImportPath: importPath,
			GoPackage:  gqlPkg,
		}
	}

	for _, msg := range file.Proto.MessageType {
		if proto.HasExtension(msg.Options, graphql.E_SkipMessage) {
			if skip, ok := proto.GetExtension(msg.Options, graphql.E_SkipMessage).(bool); ok && skip {
				// Track skipped messages so other packages can skip fields referencing them
				if file.Proto.Package != nil {
					fullyQualifiedName := *file.Proto.Package + "." + *msg.Name
					SkippedMessages[fullyQualifiedName] = true
				}
				continue
			}
		}
		f.TypeDefs = append(f.TypeDefs, types.New(msg, file.Proto, GraphqlImportMap, SkippedMessages))

	}
	return f
}

func (f File) Messages() []Message {
	return f.Message
}

func (f File) ToString() string {
	var extraImportMap = map[string]string{}
	var extraImports = []string{}

	for _, msg := range append(f.TypeDefs, f.Message...) {
		for _, imp := range msg.Imports() {
			extraImportMap[imp] = imp
		}
	}
	for _, val := range f.Imports {
		extraImportMap[val] = val
	}

	// Sort keys before iterating to ensure deterministic order
	var keys []string
	for key := range extraImportMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		extraImports = append(extraImports, extraImportMap[key])
	}

	f.Imports = extraImports

	var buf bytes.Buffer
	tpl, err := template.New("file").Funcs(map[string]interface{}{
		"has_alias": func(impt string) bool {
			return strings.Contains(impt, "\"")
		},
	}).Parse(fileTpl)
	if err != nil {
		panic(err)
	}

	if err := tpl.Execute(&buf, f); err != nil {
		panic(err)
	}

	out := buf.String()

	for _, msg := range f.TypeDefs {
		out += msg.Generate()
	}

	for _, msg := range f.Message {
		out += msg.Generate()
	}

	// Generate unified WithLoaders and Fields functions
	out += f.GenerateUnifiedFunctions()

	return out
}

func (f File) GenerateUnifiedFunctions() string {
	var out string

	// 1. Generate allMessages slice - all messages from this proto file
	out += f.generateAllMessages()

	// 2. Generate the Module implementation
	out += f.generateModule()

	return out
}

func (f File) generateAllMessages() string {
	var out string

	out += "\n// allMessages contains all message types from this proto package\n"
	out += "var allMessages = []pg.GraphqlMessage{\n"

	for _, typedef := range f.TypeDefs {
		// Add each message type to the slice
		// TypeDefs includes all messages and enums
		// We check if it's a types.Message by looking at the concrete type
		if msg, ok := typedef.(types.Message); ok {
			// This is a message type - get its name from the descriptor
			out += "\t&" + msg.Descriptor.GetName() + "{},\n"
		}
		// Skip enums and other types
	}

	out += "}\n\n"

	return out
}

func (f File) generateModule() string {
	var out string

	// Collect all services
	var services []templates.Message
	for _, msg := range f.Message {
		if svc, ok := msg.(templates.Message); ok {
			services = append(services, svc)
		}
	}

	// Generate module based on whether there are services or not
	if len(services) == 0 {
		// Type-only module (no services)
		out += f.generateTypeOnlyModule()
	} else {
		// Module with services
		out += f.generateServiceModule(services)
	}

	return out
}
