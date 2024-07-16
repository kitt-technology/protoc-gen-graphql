package dataloaders

import (
	"bytes"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/kitt-technology/protoc-gen-graphql/generation/types"
	"github.com/kitt-technology/protoc-gen-graphql/generation/util"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"strings"
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

	dns := proto.GetExtension(msg.Options, graphql.E_Host).(string)

	for _, method := range msg.Method {
		// Get output type of method
		output := util.GetMessageType(root, *method.OutputType)
		// Get input type of method
		input := util.GetMessageType(root, *method.InputType)

		if util.Last(*method.OutputType) == "Empty" {
			continue
		}

		// See if method is a batch loader
		custom := proto.HasExtension(method.Options, graphql.E_BatchLoader)
		if *method.InputType == ".graphql.BatchRequest" || custom {
			// Find type of map
			var resultType string

			if len(output.Field) == 0 || output.Field[0].Label.String() != "LABEL_REPEATED" || !strings.Contains(*output.Field[0].TypeName, "Entry") {
				panic(fmt.Sprintf("batch loaders must have one field of the type: map<string, Result> for %s.%s", *msg.Name, *method.Name))
			}

			if custom && (len(input.Field) != 1 || input.Field[0].Label.String() != "LABEL_REPEATED") {
				panic(fmt.Sprintf("custom batch loaders must have only one field of the type: repeated %s for %s.%s", util.Last(*method.InputType), *msg.Name, *method.Name))
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
						rt, _, _ := types.Types(nestedType.Field[1], root, map[string]types.GraphqlImport{}, []*descriptorpb.FileDescriptorProto{root})
						resultType = string(rt)
					}
				}
			}

			keysField := strcase.ToCamel("Keys")
			keysType := "string"
			if custom {
				keysField = strcase.ToCamel(*input.Field[0].Name)
				keysType = util.Last(*input.Field[0].TypeName)
			}

			m.Loaders = append(m.Loaders, LoaderVars{
				Method:       util.Title(*method.Name),
				RequestType:  util.Title(util.Last(*method.InputType)),
				KeysField:    keysField,
				KeysType:     keysType,
				ResultsField: strcase.ToCamel(*field.Name),
				ResultsType:  resultType,
				Custom:       custom,
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
	if len(m.Loaders) > 0 {
		return []string{"context", "github.com/graph-gophers/dataloader"}
	}
	return []string{}
}

func (m Message) Generate() string {
	var buf bytes.Buffer
	tpl.Execute(&buf, m)
	return buf.String()
}

type Method struct {
	Input  string
	Output string
	Name   string
}
