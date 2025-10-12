package types

import (
	"bytes"
	"io"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/kitt-technology/protoc-gen-graphql/generation/imports"
	"github.com/kitt-technology/protoc-gen-graphql/generation/util"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type GraphqlImport struct {
	ImportPath string
	GoPackage  string
}

type Field struct {
	GqlKey string

	GoKey  string
	GoType GoType

	Type       string
	GoFromArgs string
	InputType  string
	ArgType    string

	IsList         bool
	TypeOfType     string
	IsPointer      bool
	Proto3Optional bool
}

type Message struct {
	Descriptor       *descriptorpb.DescriptorProto
	Package          string
	Root             *descriptorpb.FileDescriptorProto
	Fields           []Field
	OneOfFields      map[string]map[string]Field
	Import           map[string]string
	ObjectName       string
	InputTypeName    string
	PackageImportMap map[string]GraphqlImport
	SkippedMessages  map[string]bool
}

func New(msg *descriptorpb.DescriptorProto, file *descriptorpb.FileDescriptorProto, graphqlImportMap map[string]GraphqlImport, skippedMessages map[string]bool) (m Message) {
	pkg := file.Package

	var actualPkg string
	if pkg != nil {
		pkgPath := strings.Split(*pkg, ".")
		if len(pkgPath) > 0 {
			actualPkg = pkgPath[len(pkgPath)-1]
		}
	}
	return Message{
		Import:           make(map[string]string),
		Descriptor:       msg,
		Root:             file,
		Package:          actualPkg,
		OneOfFields:      make(map[string]map[string]Field, 0),
		PackageImportMap: graphqlImportMap,
		SkippedMessages:  skippedMessages,
	}
}

func (m Message) Imports() []string {
	m.Generate()
	var imps []string
	for _, val := range m.Import {
		imps = append(imps, val)
	}
	return imps
}

func (m Message) Generate() string {
	// If there's a custom name, grab it
	if proto.HasExtension(m.Descriptor.Options, graphql.E_ObjectName) {
		if name, ok := proto.GetExtension(m.Descriptor.Options, graphql.E_ObjectName).(string); ok {
			m.ObjectName = name
		} else {
			m.ObjectName = *m.Descriptor.Name
		}
	} else {
		m.ObjectName = *m.Descriptor.Name
	}

	// Check for custom input type name override
	if proto.HasExtension(m.Descriptor.Options, graphql.E_InputTypeName) {
		if name, ok := proto.GetExtension(m.Descriptor.Options, graphql.E_InputTypeName).(string); ok {
			m.InputTypeName = name
		}
	}

	for _, field := range m.Descriptor.Field {
		if proto.HasExtension(field.Options, graphql.E_SkipField) {
			if skip, ok := proto.GetExtension(field.Options, graphql.E_SkipField).(bool); ok && skip {
				continue
			}
		}
		isList := false

		// It's a list or a map
		if field.Label.String() == "LABEL_REPEATED" {
			switch field.Type.String() {
			case "TYPE_MESSAGE":
				// Is it a map?
				isMap := false
				nestedTypeKey := util.Last(*field.TypeName)
				for _, nestedType := range m.Descriptor.NestedType {
					if *nestedType.Name == nestedTypeKey {
						// If it's a map, continue - we don't support maps yet
						isMap = true
					}
				}

				if isMap {
					// Maps are not supported as they there is no corresponding Graphql type
					continue
				} else {
					// It's a list of objects .e.g.
					// repeated Object objects = 1;
					isList = true
				}

			default:
				isList = true
			}
		}

		goType, gqlType, typeOfType, ok := Types(field, m.Root, m.PackageImportMap, m.SkippedMessages)
		if !ok {
			// Skip fields with unknown types
			continue
		}
		if isList {
			goType, gqlType, typeOfType, ok = Types(field, m.Root, m.PackageImportMap, m.SkippedMessages)
			if !ok {
				continue
			}
		}

		switch {
		case typeOfType == Wrapper:
			m.Import[imports.WrappersPbImport] = imports.WrappersPbImport
		case typeOfType == Timestamp:
			if isList {
				m.Import[imports.TimestampPbImport] = imports.TimestampPbImport
			}
			m.Import[imports.PggImport] = imports.PggImport
		case typeOfType == Common:

			for key, importPath := range m.PackageImportMap {
				typeNameWithProtoImport := field.GetTypeName()[1:]
				if strings.HasPrefix(typeNameWithProtoImport, key+".") {
					m.Import[importPath.ImportPath] = importPath.ImportPath
				}
			}
		}

		graphqlOptional := field.TypeName != nil
		if proto.HasExtension(field.Options, graphql.E_Optional) {
			val := proto.GetExtension(field.Options, graphql.E_Optional)
			if optional, ok := val.(bool); ok && optional {
				graphqlOptional = true
			}
		}

		if field.GetProto3Optional() {
			graphqlOptional = true
		}

		isPointer := false
		pointerTypes := []string{"Object", "Wrapper", "Timestamp", "Money", "Common"}
		for _, pointerType := range pointerTypes {
			if string(typeOfType) == pointerType {
				isPointer = true
			}
		}

		fieldVars := Field{
			GqlKey:         *field.JsonName,
			GoKey:          goKey(field),
			GoType:         goType,
			TypeOfType:     string(typeOfType),
			IsList:         isList,
			IsPointer:      isPointer,
			Proto3Optional: field.GetProto3Optional(),
		}

		// Generate input type
		typeVars := FieldTypeVars{
			TypeOfType:      string(typeOfType),
			IsList:          isList,
			GoType:          goType,
			GqlType:         gqlType,
			GqlKey:          *field.JsonName,
			Suffix:          "GraphqlInputType",
			GraphqlOptional: graphqlOptional,
		}
		var buf bytes.Buffer
		if err := typeTemplate.Execute(&buf, typeVars); err != nil {
			panic(err)
		}
		inputType, err := io.ReadAll(&buf)
		if err != nil {
			panic(err)
		}

		// Generate generic type
		typeVars.Suffix = "GraphqlType"
		if err := typeTemplate.Execute(&buf, typeVars); err != nil {
			panic(err)
		}
		normalType, err := io.ReadAll(&buf)
		if err != nil {
			panic(err)
		}
		typeVars.Suffix = "GraphqlType"

		// Generate "FromArg" string
		if err := goFromArgsTemplate.Execute(&buf, typeVars); err != nil {
			panic(err)
		}
		goFromArg, err := io.ReadAll(&buf)
		if err != nil {
			panic(err)
		}
		fieldVars.GoFromArgs = string(goFromArg)
		fieldVars.InputType = string(inputType)
		fieldVars.Type = string(normalType)

		if field.OneofIndex != nil && !field.GetProto3Optional() {
			key := *m.Descriptor.OneofDecl[*field.OneofIndex].Name

			if _, ok := m.OneOfFields[key]; !ok {
				m.OneOfFields[key] = make(map[string]Field, 0)
			}
			m.OneOfFields[key][*field.Name] = fieldVars

		} else {
			m.Fields = append(m.Fields, fieldVars)
		}
	}

	// Can't have messages with empty fields, so this is a hackaround
	if len(m.Fields) == 0 {
		m.Fields = append(m.Fields, Field{
			GqlKey:    "_null",
			Type:      "gql.Boolean",
			InputType: "gql.Boolean",
		})
	}

	var buf bytes.Buffer
	if err := messageTemplate.Execute(&buf, m); err != nil {
		panic(err)
	}

	return buf.String()
}

type Descriptor interface {
	GetType() descriptorpb.FieldDescriptorProto_Type
	GetTypeName() string
}

type (
	FieldType string
	GoType    string
	GqlType   string
)

const (
	Wrapper   FieldType = "Wrapper"
	Object    FieldType = "Object"
	Primitive FieldType = "Primitive"
	Enum      FieldType = "Enum"
	Timestamp FieldType = "Timestamp"
	Common    FieldType = "Common"
)

func Types(field *descriptorpb.FieldDescriptorProto, root *descriptorpb.FileDescriptorProto, packageImportMap map[string]GraphqlImport, skippedMessages map[string]bool) (GoType, GqlType, FieldType, bool) {
	if field.GetTypeName() != "" {
		switch field.GetTypeName() {
		case ".google.protobuf.StringValue":
			return "wrapperspb.String", "gql.String", Wrapper, true
		case ".google.protobuf.BoolValue":
			return "wrapperspb.Bool", "gql.Boolean", Wrapper, true
		case ".google.protobuf.FloatValue":
			return "wrapperspb.Float", "gql.Float", Wrapper, true
		case ".google.protobuf.Timestamp":
			return "timestamppb.Timestamp", "pg.Timestamp", Timestamp, true
		case ".google.protobuf.Int32Value":
			return "wrapperspb.Int32", "gql.Int", Wrapper, true
		case ".google.protobuf.Int64Value":
			return "wrapperspb.Int64", "gql.Int", Wrapper, true
		}
	}

	switch field.GetType() {
	case descriptorpb.FieldDescriptorProto_TYPE_STRING:
		return "string", "gql.String", Primitive, true
	case descriptorpb.FieldDescriptorProto_TYPE_INT32:
		return "int32", "gql.Int", Primitive, true
	case descriptorpb.FieldDescriptorProto_TYPE_INT64:
		return "int64", "gql.Int", Primitive, true
	case descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
		return "float32", "gql.Float", Primitive, true
	case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
		return "bool", "gql.Boolean", Primitive, true
	case descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:
		return "float64", "gql.Float", Primitive, true
	case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
		return "[]byte", "gql.String", Primitive, true
	}

	// Check if this is a skipped message
	if field.GetTypeName() != "" {
		typeNameWithProtoImport := field.GetTypeName()[1:]
		if skippedMessages[typeNameWithProtoImport] {
			// Check if this is a cross-package reference
			fieldPackage := strings.TrimSuffix(typeNameWithProtoImport, "."+util.Last(field.GetTypeName()))
			currentPackage := root.GetPackage()

			// Only skip if it's in the SAME package
			// For cross-package references, try to resolve via packageImportMap
			if fieldPackage == currentPackage {
				return "", "", "", false
			}

			// Cross-package reference to a skipped message
			// Try to resolve it via packageImportMap (it should have manual GraphQL types)
			for pkg, graphqlType := range packageImportMap {
				if pkg != root.GetPackage() && strings.HasPrefix(typeNameWithProtoImport, pkg+".") {
					typeName := strings.TrimPrefix(typeNameWithProtoImport, pkg+".")
					typeNameWithGoImport := graphqlType.GoPackage + "." + typeName
					return GoType(typeNameWithGoImport), GqlType(typeNameWithGoImport), Common, true
				}
			}

			// If we reach here, the package isn't in packageImportMap
			// This shouldn't happen if the proto file has the graphql.package option set
			// Skip the field to avoid generating broken references
			return "", "", "", false
		}
	}

	// Handle non-skipped cross-package references
	for pkg, graphqlType := range packageImportMap {
		typeNameWithProtoImport := field.GetTypeName()[1:]
		if pkg != root.GetPackage() && strings.HasPrefix(typeNameWithProtoImport, pkg+".") {
			typeName := strings.TrimPrefix(typeNameWithProtoImport, pkg+".")
			typeNameWithGoImport := graphqlType.GoPackage + "." + typeName
			return GoType(typeNameWithGoImport), GqlType(typeNameWithGoImport), Common, true
		}
	}

	// Search through message descriptors
	for _, messageType := range root.MessageType {
		if *messageType.Name == util.Last(field.GetTypeName()) {
			return GoType(util.Last(field.GetTypeName())), GqlType(*messageType.Name), Object, true
		}
	}

	// Search through enums
	for _, enumType := range root.EnumType {
		if *enumType.Name == util.Last(field.GetTypeName()) {
			return GoType(util.Last(field.GetTypeName())), GqlType(*enumType.Name), Enum, true
		}
	}

	if field.GetTypeName() == ".graphql.PageInfo" {
		return "pg.PageInfo", "pg.PageInfo", Object, true
	}

	// Unknown type - return false to indicate it should be skipped
	return "", "", "", false
}

func goKey(field *descriptorpb.FieldDescriptorProto) string {
	return strcase.ToCamel(*field.Name)
}
