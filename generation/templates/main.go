package templates

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/kitt-technology/protoc-gen-graphql/generation/types"
	"github.com/kitt-technology/protoc-gen-graphql/generation/util"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type Message struct {
	Package     string
	Descriptor  *descriptorpb.ServiceDescriptorProto
	Methods     []Method
	ServiceName string
	Dns         string
	Loaders     []LoaderVars
}

func New(msg *descriptorpb.ServiceDescriptorProto, root *descriptorpb.FileDescriptorProto) (m Message) {
	var methods []Method

	dns, _ := proto.GetExtension(msg.Options, graphql.E_Host).(string)

	for _, method := range msg.Method {
		// Get output type of method
		output := util.GetMessageType(root, *method.OutputType)
		// Get input type of method
		input := util.GetMessageType(root, *method.InputType)

		if util.Last(*method.OutputType) == "Empty" {
			continue
		}

		// See if method is a batch loader
		// A method is a batch loader if:
		// 1. It has the explicit batch_loader option, OR
		// 2. It uses graphql.BatchRequest as input
		isBatchLoader := proto.HasExtension(method.Options, graphql.E_BatchLoader) ||
			strings.HasSuffix(*method.InputType, ".BatchRequest")
		if isBatchLoader {
			// Find type of map
			var resultType string

			if len(output.Field) == 0 || output.Field[0].Label.String() != "LABEL_REPEATED" || !strings.Contains(*output.Field[0].TypeName, "Entry") {
				panic(fmt.Sprintf("batch loaders must have one field of the type: map<string, Result> for %s.%s", *msg.Name, *method.Name))
			}

			// Check if using graphql.BatchRequest (which may be nil if from external package)
			isExternalBatchRequest := input == nil && strings.HasSuffix(*method.InputType, ".BatchRequest")

			if !isExternalBatchRequest {
				if input == nil || len(input.Field) != 1 || input.Field[0].Label.String() != "LABEL_REPEATED" {
					panic(fmt.Sprintf("batch loaders must have only one repeated field for %s.%s", *msg.Name, *method.Name))
				}
			}

			var field = output.Field[0]

			resultType = util.Title(util.Last(*field.TypeName))
			nestedTypeKey := util.Last(*field.TypeName)
			for _, nestedType := range output.NestedType {
				if *nestedType.Name == nestedTypeKey {
					if nestedType.Field[1].TypeName != nil {

						resultType = util.Last(*nestedType.Field[1].TypeName)

						if !strings.Contains(resultType, "*") {
							resultType = "*" + resultType
						}
					} else {
						rt, _, _, _ := types.Types(nestedType.Field[1], root, map[string]types.GraphqlImport{}, map[string]bool{})
						resultType = string(rt)
					}
				}
			}

			var keysField string
			var keysType string
			var isStringKey bool

			if isExternalBatchRequest {
				// graphql.BatchRequest always has: repeated string keys = 1
				keysField = "Keys"
				keysType = "string"
				isStringKey = true
			} else {
				keysField = strcase.ToCamel(*input.Field[0].Name)

				if input.Field[0].TypeName != nil && *input.Field[0].TypeName != "" {
					// Complex type (message, enum, etc.)
					keysType = util.Last(*input.Field[0].TypeName)
				} else {
					// Scalar type (string, int, etc.)
					goType, _, _, _ := types.Types(input.Field[0], root, map[string]types.GraphqlImport{}, map[string]bool{})
					keysType = string(goType)
				}

				// Determine if this is a simple string key (can use dataloader.StringKey)
				// or a custom type (needs custom Key wrapper)
				isStringKey = keysType == "string"
			}

			m.Loaders = append(m.Loaders, LoaderVars{
				Method:       util.Title(*method.Name),
				RequestType:  util.Title(util.Last(*method.InputType)),
				ResponseType: util.Title(util.Last(*method.OutputType)),
				KeysField:    keysField,
				KeysType:     keysType,
				ResultsField: strcase.ToCamel(*field.Name),
				ResultsType:  resultType,
				Custom:       !isStringKey,
			})

		} else {
			methods = append(methods, Method{Input: util.Last(*method.InputType), Output: util.Last(*method.OutputType), Name: util.Title(*method.Name)})
		}
	}

	var pkg string
	if root.Package != nil {
		pkg = *root.Package
	}
	pkgPath := strings.Split(pkg, ".")

	return Message{
		Package:     pkg,
		Descriptor:  msg,
		Methods:     methods,
		ServiceName: pkgPath[len(pkgPath)-1],
		Dns:         dns,
		Loaders:     m.Loaders,
	}
}

func (m Message) Imports() []string {
	imports := []string{"context", "os", "google.golang.org/grpc"}
	if len(m.Loaders) > 0 {
		imports = append(imports, "github.com/graph-gophers/dataloader")
	}
	return imports
}

func (m Message) Generate() string {
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, m); err != nil {
		panic(err)
	}
	return buf.String()
}

type Method struct {
	Input  string
	Output string
	Name   string
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

// GenerateUnified generates the unified WithLoaders and Fields functions
func GenerateUnified(services []Message) string {
	if len(services) == 0 {
		return ""
	}

	var out string

	// Generate unified WithLoaders function
	hasLoaders := false
	for _, svc := range services {
		if len(svc.Loaders) > 0 {
			hasLoaders = true
			break
		}
	}

	if hasLoaders {
		out += "\n// WithLoaders adds all batch loaders from all services to the context\n"
		out += "func WithLoaders(ctx context.Context) context.Context {\n"
		for _, svc := range services {
			if len(svc.Loaders) > 0 {
				out += "\tctx = " + *svc.Descriptor.Name + "WithLoaders(ctx)\n"
			}
		}
		out += "\treturn ctx\n"
		out += "}\n"
	}

	// Generate unified Fields function
	out += "\n// Fields returns all GraphQL fields from all services\n"
	out += "func Fields(ctx context.Context) []*gql.Field {\n"
	out += "\tvar fields []*gql.Field\n"
	out += "\tvar serviceFields []*gql.Field\n\n"

	for _, svc := range services {
		out += "\tctx, serviceFields = " + *svc.Descriptor.Name + "Init(ctx)\n"
		out += "\tfields = append(fields, serviceFields...)\n\n"
	}

	out += "\treturn fields\n"
	out += "}\n"

	return out
}
