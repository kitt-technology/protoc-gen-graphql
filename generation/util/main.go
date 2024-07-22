package util

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func Last(path string) string {
	t := strings.Split(path, ".")
	return t[len(t)-1]
}

// The parsing for graphql.package is copied from protoc-gen-go's package parsing
// https://github.com/golang/protobuf/blob/ae97035608a719c7a1c1c41bed0ae0744bdb0c6f/protoc-gen-go/generator/generator.go#L275

func ParseGraphqlPackage(file *descriptorpb.FileDescriptorProto) (importPath string, pkg string, ok bool) {
	if !proto.HasExtension(file.Options, graphql.E_Package) {
		return "", "", false
	}
	graphqlPackage := proto.GetExtension(file.Options, graphql.E_Package).(string)

	sc := strings.Index(graphqlPackage, ";")
	if sc >= 0 {
		return graphqlPackage[:sc], cleanPackageName(graphqlPackage[sc+1:]), true
	}
	slash := strings.LastIndex(graphqlPackage, "/")
	if slash >= 0 {
		return graphqlPackage, cleanPackageName(graphqlPackage[slash+1:]), true
	}
	return "", cleanPackageName(graphqlPackage), true
}

func cleanPackageName(name string) string {
	name = strings.Map(badToUnderscore, name)
	// Identifier must not be keyword or predeclared identifier: insert _.
	if isGoKeyword[name] {
		name = "_" + name
	}
	// Identifier must not begin with digit: insert _.
	if r, _ := utf8.DecodeRuneInString(name); unicode.IsDigit(r) {
		name = "_" + name
	}
	return name
}

var isGoKeyword = map[string]bool{
	"break":       true,
	"case":        true,
	"chan":        true,
	"const":       true,
	"continue":    true,
	"default":     true,
	"else":        true,
	"defer":       true,
	"fallthrough": true,
	"for":         true,
	"func":        true,
	"go":          true,
	"goto":        true,
	"if":          true,
	"import":      true,
	"interface":   true,
	"map":         true,
	"package":     true,
	"range":       true,
	"return":      true,
	"select":      true,
	"struct":      true,
	"switch":      true,
	"type":        true,
	"var":         true,
}

// badToUnderscore is the mapping function used to generate Go names from package names,
// which can be dotted in the input .proto file.  It replaces non-identifier characters such as
// dot or dash with underscore.
func badToUnderscore(r rune) rune {
	if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
		return r
	}
	return '_'
}

func Title(str string) string {
	return cases.Title(language.Und, cases.NoLower).String(str)
}

func GetMessageType(allRoots []*descriptorpb.FileDescriptorProto, messageType string) *descriptorpb.DescriptorProto {
	for _, root := range allRoots {
		for _, msgType := range root.MessageType {
			if Last(messageType) == *msgType.Name {
				return msgType
			}
		}
	}

	return nil
}
