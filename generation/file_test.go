package generation

import (
	"strings"
	"testing"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestFile_ToString(t *testing.T) {
	tests := []struct {
		name           string
		pkg            protogen.GoPackageName
		imports        []string
		wantContains   []string
		wantNotContain []string
	}{
		{
			name: "basic file with imports",
			pkg:  "mypackage",
			imports: []string{
				"context",
				"github.com/graphql-go/graphql",
			},
			wantContains: []string{
				"package mypackage",
				`gql "github.com/graphql-go/graphql"`,
				`"context"`,
			},
		},
		{
			name:    "file without extra imports",
			pkg:     "testpkg",
			imports: []string{},
			wantContains: []string{
				"package testpkg",
				`gql "github.com/graphql-go/graphql"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := File{
				Package: tt.pkg,
				Imports: tt.imports,
			}

			got := f.ToString()

			for _, want := range tt.wantContains {
				if !strings.Contains(got, want) {
					t.Errorf("ToString() output should contain %q, but doesn't.\nGot:\n%s", want, got)
				}
			}

			for _, notWant := range tt.wantNotContain {
				if strings.Contains(got, notWant) {
					t.Errorf("ToString() output should NOT contain %q, but does.\nGot:\n%s", notWant, got)
				}
			}
		})
	}
}

func TestFile_ImportDeduplication(t *testing.T) {
	f := File{
		Package: "testpkg",
		Imports: []string{
			"context",
			"context", // duplicate
			"fmt",
		},
	}

	output := f.ToString()

	// Count occurrences of "context"
	contextCount := strings.Count(output, `"context"`)
	if contextCount != 1 {
		t.Errorf("expected 'context' import to appear exactly once, got %d occurrences", contextCount)
	}
}

func TestFile_ImportSorting(t *testing.T) {
	f := File{
		Package: "testpkg",
		Imports: []string{
			"z_package",
			"a_package",
			"m_package",
		},
	}

	output := f.ToString()

	// Find positions of imports
	posA := strings.Index(output, "a_package")
	posM := strings.Index(output, "m_package")
	posZ := strings.Index(output, "z_package")

	if posA == -1 || posM == -1 || posZ == -1 {
		t.Fatal("not all imports found in output")
	}

	// Verify they are in alphabetical order
	if !(posA < posM && posM < posZ) {
		t.Error("imports are not sorted alphabetically")
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name            string
		proto           *descriptorpb.FileDescriptorProto
		wantServicesLen int
		wantTypeDefsLen int
	}{
		{
			name: "file with service",
			proto: &descriptorpb.FileDescriptorProto{
				Name:    proto.String("test.proto"),
				Package: proto.String("test"),
				Service: []*descriptorpb.ServiceDescriptorProto{
					{Name: proto.String("MyService")},
				},
			},
			wantServicesLen: 1,
			wantTypeDefsLen: 0,
		},
		{
			name: "file with enum",
			proto: &descriptorpb.FileDescriptorProto{
				Name:    proto.String("test.proto"),
				Package: proto.String("test"),
				EnumType: []*descriptorpb.EnumDescriptorProto{
					{Name: proto.String("MyEnum")},
				},
			},
			wantServicesLen: 0,
			wantTypeDefsLen: 1,
		},
		{
			name: "file with message",
			proto: &descriptorpb.FileDescriptorProto{
				Name:    proto.String("test.proto"),
				Package: proto.String("test"),
				MessageType: []*descriptorpb.DescriptorProto{
					{Name: proto.String("MyMessage")},
				},
			},
			wantServicesLen: 0,
			wantTypeDefsLen: 1,
		},
		{
			name: "empty file",
			proto: &descriptorpb.FileDescriptorProto{
				Name:    proto.String("test.proto"),
				Package: proto.String("test"),
			},
			wantServicesLen: 0,
			wantTypeDefsLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := &protogen.File{
				Proto:         tt.proto,
				GoPackageName: "testpkg",
			}

			got := New(file)

			if len(got.Message) != tt.wantServicesLen {
				t.Errorf("New() services count = %d, want %d", len(got.Message), tt.wantServicesLen)
			}

			if len(got.TypeDefs) != tt.wantTypeDefsLen {
				t.Errorf("New() typedefs count = %d, want %d", len(got.TypeDefs), tt.wantTypeDefsLen)
			}

			if got.Package != file.GoPackageName {
				t.Errorf("New() package = %q, want %q", got.Package, file.GoPackageName)
			}
		})
	}
}

func TestFile_WithMessages(t *testing.T) {
	mockMsg := &mockMessage{
		generate: "// generated code\n",
		imports:  []string{"context", "fmt"},
	}

	f := File{
		Package: "testpkg",
		Message: []Message{mockMsg},
	}

	output := f.ToString()

	if !strings.Contains(output, "// generated code") {
		t.Error("ToString() should contain generated message code")
	}

	if !strings.Contains(output, `"context"`) {
		t.Error("ToString() should contain message imports")
	}
}

func TestFile_WithTypeDefs(t *testing.T) {
	mockTypeDef := &mockMessage{
		generate: "// type definition\n",
		imports:  []string{"github.com/example/types"},
	}

	f := File{
		Package:  "testpkg",
		TypeDefs: []Message{mockTypeDef},
	}

	output := f.ToString()

	if !strings.Contains(output, "// type definition") {
		t.Error("ToString() should contain typedef code")
	}

	if !strings.Contains(output, "github.com/example/types") {
		t.Error("ToString() should contain typedef imports")
	}
}

// mockMessage is a mock implementation of the Message interface for testing
type mockMessage struct {
	generate string
	imports  []string
}

func (m *mockMessage) Generate() string {
	return m.generate
}

func (m *mockMessage) Imports() []string {
	return m.imports
}
