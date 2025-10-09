package dataloaders

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
		isBatchLoader := proto.HasExtension(method.Options, graphql.E_BatchLoader)
		if isBatchLoader {
			// Find type of map
			var resultType string

			if len(output.Field) == 0 || output.Field[0].Label.String() != "LABEL_REPEATED" || !strings.Contains(*output.Field[0].TypeName, "Entry") {
				panic(fmt.Sprintf("batch loaders must have one field of the type: map<string, Result> for %s.%s", *msg.Name, *method.Name))
			}

			if len(input.Field) != 1 || input.Field[0].Label.String() != "LABEL_REPEATED" {
				panic(fmt.Sprintf("batch loaders must have only one repeated field for %s.%s", *msg.Name, *method.Name))
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
						rt, _, _, _ := types.Types(nestedType.Field[1], root, map[string]types.GraphqlImport{})
						resultType = string(rt)
					}
				}
			}

			keysField := strcase.ToCamel(*input.Field[0].Name)

			var keysType string
			if input.Field[0].TypeName != nil && *input.Field[0].TypeName != "" {
				// Complex type (message, enum, etc.)
				keysType = util.Last(*input.Field[0].TypeName)
			} else {
				// Scalar type (string, int, etc.)
				goType, _, _, _ := types.Types(input.Field[0], root, map[string]types.GraphqlImport{})
				keysType = string(goType)
			}

			// Determine if this is a simple string key (can use dataloader.StringKey)
			// or a custom type (needs custom Key wrapper)
			isStringKey := keysType == "string"

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
	imports := []string{"os", "google.golang.org/grpc"}
	if len(m.Loaders) > 0 {
		imports = append(imports, "context", "github.com/graph-gophers/dataloader")
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
