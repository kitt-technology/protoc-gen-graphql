package util

import (
	"testing"

	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestLast(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "simple path",
			input: "foo.bar.baz",
			want:  "baz",
		},
		{
			name:  "single element",
			input: "foo",
			want:  "foo",
		},
		{
			name:  "package path",
			input: "com.example.mypackage.MyMessage",
			want:  "MyMessage",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Last(tt.input)
			if got != tt.want {
				t.Errorf("Last(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestParseGraphqlPackage(t *testing.T) {
	tests := []struct {
		name           string
		graphqlPackage string
		wantImportPath string
		wantPkg        string
		wantOk         bool
		hasExtension   bool
	}{
		{
			name:           "simple package name",
			graphqlPackage: "mypkg",
			wantImportPath: "",
			wantPkg:        "mypkg",
			wantOk:         true,
			hasExtension:   true,
		},
		{
			name:           "import path with package",
			graphqlPackage: "github.com/example/repo;mypkg",
			wantImportPath: "github.com/example/repo",
			wantPkg:        "mypkg",
			wantOk:         true,
			hasExtension:   true,
		},
		{
			name:           "import path without semicolon",
			graphqlPackage: "github.com/example/mypkg",
			wantImportPath: "github.com/example/mypkg",
			wantPkg:        "mypkg",
			wantOk:         true,
			hasExtension:   true,
		},
		{
			name:           "no extension present",
			graphqlPackage: "",
			wantImportPath: "",
			wantPkg:        "",
			wantOk:         false,
			hasExtension:   false,
		},
		{
			name:           "package with dashes",
			graphqlPackage: "github.com/example/my-pkg;my_pkg",
			wantImportPath: "github.com/example/my-pkg",
			wantPkg:        "my_pkg",
			wantOk:         true,
			hasExtension:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := &descriptorpb.FileDescriptorProto{
				Options: &descriptorpb.FileOptions{},
			}

			if tt.hasExtension {
				proto.SetExtension(file.Options, graphql.E_Package, tt.graphqlPackage)
			}

			importPath, pkg, ok := ParseGraphqlPackage(file)
			if ok != tt.wantOk {
				t.Errorf("ParseGraphqlPackage() ok = %v, want %v", ok, tt.wantOk)
			}
			if importPath != tt.wantImportPath {
				t.Errorf("ParseGraphqlPackage() importPath = %q, want %q", importPath, tt.wantImportPath)
			}
			if pkg != tt.wantPkg {
				t.Errorf("ParseGraphqlPackage() pkg = %q, want %q", pkg, tt.wantPkg)
			}
		})
	}
}

func TestCleanPackageName(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "simple name",
			input: "mypackage",
			want:  "mypackage",
		},
		{
			name:  "name with dashes",
			input: "my-package",
			want:  "my_package",
		},
		{
			name:  "name with dots",
			input: "my.package",
			want:  "my_package",
		},
		{
			name:  "go keyword - package",
			input: "package",
			want:  "_package",
		},
		{
			name:  "go keyword - func",
			input: "func",
			want:  "_func",
		},
		{
			name:  "go keyword - type",
			input: "type",
			want:  "_type",
		},
		{
			name:  "starts with digit",
			input: "123package",
			want:  "_123package",
		},
		{
			name:  "multiple special characters",
			input: "my-special.package!",
			want:  "my_special_package_",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cleanPackageName(tt.input)
			if got != tt.want {
				t.Errorf("cleanPackageName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestBadToUnderscore(t *testing.T) {
	tests := []struct {
		name  string
		input rune
		want  rune
	}{
		{
			name:  "letter",
			input: 'a',
			want:  'a',
		},
		{
			name:  "uppercase letter",
			input: 'A',
			want:  'A',
		},
		{
			name:  "digit",
			input: '5',
			want:  '5',
		},
		{
			name:  "underscore",
			input: '_',
			want:  '_',
		},
		{
			name:  "dash",
			input: '-',
			want:  '_',
		},
		{
			name:  "dot",
			input: '.',
			want:  '_',
		},
		{
			name:  "exclamation",
			input: '!',
			want:  '_',
		},
		{
			name:  "space",
			input: ' ',
			want:  '_',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := badToUnderscore(tt.input)
			if got != tt.want {
				t.Errorf("badToUnderscore(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestTitle(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "lowercase word",
			input: "hello",
			want:  "Hello",
		},
		{
			name:  "already capitalized",
			input: "Hello",
			want:  "Hello",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "mixed case",
			input: "helloWorld",
			want:  "HelloWorld",
		},
		{
			name:  "all caps",
			input: "HELLO",
			want:  "HELLO",
		},
		{
			name:  "single character",
			input: "a",
			want:  "A",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Title(tt.input)
			if got != tt.want {
				t.Errorf("Title(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestGetMessageType(t *testing.T) {
	tests := []struct {
		name        string
		messageType string
		wantFound   bool
		wantName    string
	}{
		{
			name:        "existing message",
			messageType: "MyMessage",
			wantFound:   true,
			wantName:    "MyMessage",
		},
		{
			name:        "existing message with package prefix",
			messageType: "package.MyMessage",
			wantFound:   true,
			wantName:    "MyMessage",
		},
		{
			name:        "non-existing message",
			messageType: "NonExistent",
			wantFound:   false,
		},
		{
			name:        "another existing message",
			messageType: "AnotherMessage",
			wantFound:   true,
			wantName:    "AnotherMessage",
		},
	}

	root := &descriptorpb.FileDescriptorProto{
		MessageType: []*descriptorpb.DescriptorProto{
			{Name: proto.String("MyMessage")},
			{Name: proto.String("AnotherMessage")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetMessageType(root, tt.messageType)
			if tt.wantFound {
				if got == nil {
					t.Errorf("GetMessageType(%q) = nil, want message", tt.messageType)
				} else if got.GetName() != tt.wantName {
					t.Errorf("GetMessageType(%q) name = %q, want %q", tt.messageType, got.GetName(), tt.wantName)
				}
			} else {
				if got != nil {
					t.Errorf("GetMessageType(%q) = %v, want nil", tt.messageType, got)
				}
			}
		})
	}
}

func TestIsGoKeyword(t *testing.T) {
	// Test that all go keywords are in the map
	keywords := []string{
		"break", "case", "chan", "const", "continue",
		"default", "else", "defer", "fallthrough", "for",
		"func", "go", "goto", "if", "import",
		"interface", "map", "package", "range", "return",
		"select", "struct", "switch", "type", "var",
	}

	for _, keyword := range keywords {
		t.Run(keyword, func(t *testing.T) {
			if !isGoKeyword[keyword] {
				t.Errorf("expected %q to be in isGoKeyword map", keyword)
			}
		})
	}

	// Test that non-keywords are not in the map
	nonKeywords := []string{"myvar", "hello", "world", "foo123"}
	for _, nonKeyword := range nonKeywords {
		t.Run("not_"+nonKeyword, func(t *testing.T) {
			if isGoKeyword[nonKeyword] {
				t.Errorf("expected %q to NOT be in isGoKeyword map", nonKeyword)
			}
		})
	}
}
