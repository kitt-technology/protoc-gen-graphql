package enum

import (
	"strings"
	"testing"

	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name         string
		descriptor   *descriptorpb.EnumDescriptorProto
		wantEnumName string
	}{
		{
			name: "basic enum without custom name",
			descriptor: &descriptorpb.EnumDescriptorProto{
				Name: proto.String("MyEnum"),
			},
			wantEnumName: "MyEnum",
		},
		{
			name: "enum with custom graphql name",
			descriptor: &descriptorpb.EnumDescriptorProto{
				Name: proto.String("MyEnum"),
				Options: &descriptorpb.EnumOptions{},
			},
			wantEnumName: "CustomName",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up custom name extension if needed
			if tt.wantEnumName == "CustomName" {
				proto.SetExtension(tt.descriptor.Options, graphql.E_EnumName, "CustomName")
			}

			got := New(tt.descriptor)

			if got.EnumName != tt.wantEnumName {
				t.Errorf("New() EnumName = %q, want %q", got.EnumName, tt.wantEnumName)
			}

			if got.Descriptor != tt.descriptor {
				t.Error("New() Descriptor should match input descriptor")
			}

			if got.Import == nil {
				t.Error("New() Import map should be initialized")
			}

			if got.Values == nil {
				t.Error("New() Values map should be initialized")
			}
		})
	}
}

func TestMessage_Imports(t *testing.T) {
	m := Message{
		Descriptor: &descriptorpb.EnumDescriptorProto{
			Name: proto.String("TestEnum"),
		},
		EnumName: "TestEnum",
		Import:   make(map[string]string),
		Values:   make(map[string]string),
	}

	imports := m.Imports()

	if len(imports) != 1 {
		t.Fatalf("Imports() should return 1 import, got %d", len(imports))
	}

	expectedImport := "github.com/graphql-go/graphql/language/ast"
	if imports[0] != expectedImport {
		t.Errorf("Imports() = %q, want %q", imports[0], expectedImport)
	}
}

func TestMessage_Generate(t *testing.T) {
	tests := []struct {
		name           string
		descriptor     *descriptorpb.EnumDescriptorProto
		enumName       string
		wantContains   []string
		wantNotContain []string
	}{
		{
			name: "basic enum generation",
			descriptor: &descriptorpb.EnumDescriptorProto{
				Name: proto.String("Status"),
				Value: []*descriptorpb.EnumValueDescriptorProto{
					{Name: proto.String("ACTIVE"), Number: proto.Int32(0)},
					{Name: proto.String("INACTIVE"), Number: proto.Int32(1)},
				},
			},
			enumName: "Status",
			wantContains: []string{
				"var StatusGraphqlEnum = gql.NewEnum",
				`Name: "Status"`,
				`"ACTIVE"`,
				`"INACTIVE"`,
				"Status(0)",
				"Status(1)",
				"var StatusGraphqlType = gql.NewScalar",
			},
		},
		{
			name: "enum with custom name",
			descriptor: &descriptorpb.EnumDescriptorProto{
				Name: proto.String("Priority"),
				Value: []*descriptorpb.EnumValueDescriptorProto{
					{Name: proto.String("LOW"), Number: proto.Int32(0)},
					{Name: proto.String("HIGH"), Number: proto.Int32(1)},
				},
			},
			enumName: "PriorityLevel",
			wantContains: []string{
				"var PriorityGraphqlEnum",
				`Name: "PriorityLevel"`,
				`"LOW"`,
				`"HIGH"`,
			},
		},
		{
			name: "empty enum",
			descriptor: &descriptorpb.EnumDescriptorProto{
				Name:  proto.String("Empty"),
				Value: []*descriptorpb.EnumValueDescriptorProto{},
			},
			enumName: "Empty",
			wantContains: []string{
				"var EmptyGraphqlEnum",
				`Name: "Empty"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Message{
				Descriptor: tt.descriptor,
				EnumName:   tt.enumName,
				Import:     make(map[string]string),
				Values:     make(map[string]string),
			}

			got := m.Generate()

			for _, want := range tt.wantContains {
				if !strings.Contains(got, want) {
					t.Errorf("Generate() should contain %q, but doesn't.\nGot:\n%s", want, got)
				}
			}

			for _, notWant := range tt.wantNotContain {
				if strings.Contains(got, notWant) {
					t.Errorf("Generate() should NOT contain %q, but does.\nGot:\n%s", notWant, got)
				}
			}
		})
	}
}

func TestMessage_Generate_ValueMapping(t *testing.T) {
	// Test that enum values are correctly mapped to numbers
	m := Message{
		Descriptor: &descriptorpb.EnumDescriptorProto{
			Name: proto.String("TestEnum"),
			Value: []*descriptorpb.EnumValueDescriptorProto{
				{Name: proto.String("ZERO"), Number: proto.Int32(0)},
				{Name: proto.String("ONE"), Number: proto.Int32(1)},
				{Name: proto.String("TEN"), Number: proto.Int32(10)},
			},
		},
		EnumName: "TestEnum",
		Import:   make(map[string]string),
		Values:   make(map[string]string),
	}

	output := m.Generate()

	// Verify that each enum value is mapped to the correct number
	if !strings.Contains(output, "TestEnum(0)") {
		t.Error("Generate() should map ZERO to TestEnum(0)")
	}
	if !strings.Contains(output, "TestEnum(1)") {
		t.Error("Generate() should map ONE to TestEnum(1)")
	}
	if !strings.Contains(output, "TestEnum(10)") {
		t.Error("Generate() should map TEN to TestEnum(10)")
	}
}

func TestMessage_Generate_GraphqlEnumAndType(t *testing.T) {
	// Test that both enum and scalar type are generated
	m := Message{
		Descriptor: &descriptorpb.EnumDescriptorProto{
			Name: proto.String("MyEnum"),
			Value: []*descriptorpb.EnumValueDescriptorProto{
				{Name: proto.String("VALUE1"), Number: proto.Int32(0)},
			},
		},
		EnumName: "MyEnum",
		Import:   make(map[string]string),
		Values:   make(map[string]string),
	}

	output := m.Generate()

	// Should generate both GraphqlEnum and GraphqlType
	if !strings.Contains(output, "var MyEnumGraphqlEnum =") {
		t.Error("Generate() should create GraphqlEnum variable")
	}

	if !strings.Contains(output, "var MyEnumGraphqlType =") {
		t.Error("Generate() should create GraphqlType variable")
	}

	// GraphqlType should have correct methods
	expectedMethods := []string{"ParseValue", "Serialize", "ParseLiteral"}
	for _, method := range expectedMethods {
		if !strings.Contains(output, method) {
			t.Errorf("Generate() GraphqlType should have %s method", method)
		}
	}
}
