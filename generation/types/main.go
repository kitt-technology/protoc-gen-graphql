package types

import (
	"bytes"
	"io/ioutil"
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
	PackageImportMap map[string]GraphqlImport
	AllRoots         []*descriptorpb.FileDescriptorProto
}

func New(msg *descriptorpb.DescriptorProto, file *descriptorpb.FileDescriptorProto, graphqlImportMap map[string]GraphqlImport, allRoots []*descriptorpb.FileDescriptorProto) (m Message) {
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
		AllRoots:         allRoots,
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
		m.ObjectName = proto.GetExtension(m.Descriptor.Options, graphql.E_ObjectName).(string)
	} else {
		m.ObjectName = *m.Descriptor.Name
	}

	for _, field := range m.Descriptor.Field {
		if proto.HasExtension(field.Options, graphql.E_SkipField) &&
			proto.GetExtension(field.Options, graphql.E_SkipField).(bool) {
			continue
		}
		isList := false

		switch field.Label.String() {
		// It's a list or a map
		case "LABEL_REPEATED":
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

				break
			default:
				isList = true
			}
		}

		goType, gqlType, typeOfType := Types(field, m.AllRoots, m.PackageImportMap)
		if isList {
			goType, gqlType, typeOfType = Types(field, m.AllRoots, m.PackageImportMap)
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
				if strings.HasPrefix(typeNameWithProtoImport, key) {
					m.Import[importPath.ImportPath] = importPath.ImportPath
				}
			}
		}

		graphqlOptional := field.TypeName != nil
		if proto.HasExtension(field.Options, graphql.E_Optional) {
			val := proto.GetExtension(field.Options, graphql.E_Optional)
			if val.(bool) {
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
		typeTemplate.Execute(&buf, typeVars)
		inputType, err := ioutil.ReadAll(&buf)
		if err != nil {
			panic(err)
		}

		// Generate generic type
		typeVars.Suffix = "GraphqlType"
		typeTemplate.Execute(&buf, typeVars)
		normalType, err := ioutil.ReadAll(&buf)
		typeVars.Suffix = "GraphqlType"

		// Generate "FromArg" string
		goFromArgsTemplate.Execute(&buf, typeVars)
		goFromArg, err := ioutil.ReadAll(&buf)
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
	messageTemplate.Execute(&buf, m)

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
	Object              = "Object"
	Primitive           = "Primitive"
	Enum                = "Enum"
	Timestamp           = "Timestamp"
	Common              = "Common"
)

func Types(field *descriptorpb.FieldDescriptorProto, allRoots []*descriptorpb.FileDescriptorProto, packageImportMap map[string]GraphqlImport) (GoType, GqlType, FieldType) {
	if field.GetTypeName() != "" {
		switch field.GetTypeName() {
		case ".google.protobuf.StringValue":
			return "wrapperspb.String", "gql.String", Wrapper
		case ".google.protobuf.BoolValue":
			return "wrapperspb.Bool", "gql.Boolean", Wrapper
		case ".google.protobuf.FloatValue":
			return "wrapperspb.Float", "gql.Float", Wrapper
		case ".google.protobuf.Timestamp":
			return "timestamppb.Timestamp", "pg.Timestamp", Timestamp
		case ".google.protobuf.Int32Value":
			return "wrapperspb.Int32", "gql.Int", Wrapper
		case ".google.protobuf.Int64Value":
			return "wrapperspb.Int64", "gql.Int", Wrapper
		}
	}

	switch field.GetType() {
	case descriptorpb.FieldDescriptorProto_TYPE_STRING:
		return "string", "gql.String", Primitive
	case descriptorpb.FieldDescriptorProto_TYPE_INT32:
		return "int32", "gql.Int", Primitive
	case descriptorpb.FieldDescriptorProto_TYPE_INT64:
		return "int64", "gql.Int", Primitive
	case descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
		return "float32", "gql.Float", Primitive
	case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
		return "bool", "gql.Boolean", Primitive
	case descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:
		return "float64", "gql.Float", Primitive
	case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
		return "[]byte", "gql.String", Primitive
	}

	for _, root := range allRoots {
		for pkg, graphqlType := range packageImportMap {
			typeNameWithProtoImport := field.GetTypeName()[1:]
			if pkg != root.GetPackage() && strings.HasPrefix(typeNameWithProtoImport, pkg) {
				typeName := strings.TrimPrefix(typeNameWithProtoImport, pkg)
				typeNameWithGoImport := graphqlType.GoPackage + typeName
				return GoType(typeNameWithGoImport), GqlType(typeNameWithGoImport), Common
			}
		}

		// Search through message descriptors
		for _, messageType := range root.MessageType {
			if *messageType.Name == util.Last(field.GetTypeName()) {
				return GoType(util.Last(field.GetTypeName())), GqlType(*messageType.Name), Object
			}
		}

		// Search through enums
		for _, enumType := range root.EnumType {
			if *enumType.Name == util.Last(field.GetTypeName()) {
				return GoType(util.Last(field.GetTypeName())), GqlType(*enumType.Name), Enum
			}
		}

		if field.GetTypeName() == ".graphql.PageInfo" {
			return "pg.PageInfo", "pg.PageInfo", Object
		}

		for _, messageType := range root.MessageType {
			if *messageType.Name == util.Last(field.GetTypeName()) {
				return GoType(util.Last(field.GetTypeName())), GqlType(*messageType.Name), Object
			}
		}

	}

	panic(field)
}

func goKey(field *descriptorpb.FieldDescriptorProto) string {
	return strcase.ToCamel(*field.Name)
}
