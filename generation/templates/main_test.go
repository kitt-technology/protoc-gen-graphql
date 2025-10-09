package templates

import (
	"strings"
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestNew_BasicService(t *testing.T) {
	root := &descriptorpb.FileDescriptorProto{
		Package: proto.String("example.service"),
		MessageType: []*descriptorpb.DescriptorProto{
			{Name: proto.String("GetUserRequest")},
			{Name: proto.String("GetUserResponse")},
		},
	}

	service := &descriptorpb.ServiceDescriptorProto{
		Name: proto.String("UserService"),
		Method: []*descriptorpb.MethodDescriptorProto{
			{
				Name:       proto.String("GetUser"),
				InputType:  proto.String(".example.service.GetUserRequest"),
				OutputType: proto.String(".example.service.GetUserResponse"),
			},
		},
	}

	m := New(service, root)

	if m.Package != "example.service" {
		t.Errorf("New() Package = %q, want %q", m.Package, "example.service")
	}

	if m.Descriptor != service {
		t.Error("New() Descriptor should match input service")
	}

	if len(m.Methods) != 1 {
		t.Fatalf("New() should have 1 method, got %d", len(m.Methods))
	}

	method := m.Methods[0]
	if method.Name != "GetUser" {
		t.Errorf("Method Name = %q, want %q", method.Name, "GetUser")
	}
	if method.Input != "GetUserRequest" {
		t.Errorf("Method Input = %q, want %q", method.Input, "GetUserRequest")
	}
	if method.Output != "GetUserResponse" {
		t.Errorf("Method Output = %q, want %q", method.Output, "GetUserResponse")
	}

	if m.ServiceName != "service" {
		t.Errorf("ServiceName = %q, want %q", m.ServiceName, "service")
	}
}

func TestNew_SkipsEmptyOutput(t *testing.T) {
	root := &descriptorpb.FileDescriptorProto{
		Package: proto.String("example.service"),
		MessageType: []*descriptorpb.DescriptorProto{
			{Name: proto.String("DeleteRequest")},
		},
	}

	service := &descriptorpb.ServiceDescriptorProto{
		Name: proto.String("UserService"),
		Method: []*descriptorpb.MethodDescriptorProto{
			{
				Name:       proto.String("Delete"),
				InputType:  proto.String(".example.service.DeleteRequest"),
				OutputType: proto.String(".google.protobuf.Empty"),
			},
		},
	}

	m := New(service, root)

	if len(m.Methods) != 0 {
		t.Errorf("New() should skip methods with Empty output, got %d methods", len(m.Methods))
	}
}

func TestNew_DetectsBatchRequestAsLoader(t *testing.T) {
	// Test that using graphql.BatchRequest as input auto-detects a batch loader
	// even without explicit (graphql.batch_loader) option
	root := &descriptorpb.FileDescriptorProto{
		Package: proto.String("example.service"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: proto.String("LoadItemsResponse"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:     proto.String("results"),
						Number:   proto.Int32(1),
						Label:    descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
						Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
						TypeName: proto.String(".example.service.LoadItemsResponse.ResultsEntry"),
					},
				},
				NestedType: []*descriptorpb.DescriptorProto{
					{
						Name: proto.String("ResultsEntry"),
						Field: []*descriptorpb.FieldDescriptorProto{
							{
								Name:   proto.String("key"),
								Number: proto.Int32(1),
								Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
							},
							{
								Name:     proto.String("value"),
								Number:   proto.Int32(2),
								Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
								TypeName: proto.String(".example.service.Item"),
							},
						},
					},
				},
			},
			{Name: proto.String("Item")},
		},
	}

	service := &descriptorpb.ServiceDescriptorProto{
		Name: proto.String("ItemService"),
		Method: []*descriptorpb.MethodDescriptorProto{
			{
				Name:       proto.String("loadItems"),
				InputType:  proto.String(".graphql.BatchRequest"),
				OutputType: proto.String(".example.service.LoadItemsResponse"),
			},
		},
	}

	m := New(service, root)

	// Method should NOT be in Methods array (it's a loader)
	if len(m.Methods) != 0 {
		t.Errorf("New() should not include batch loader in Methods, got %d methods", len(m.Methods))
	}

	// Should be detected as a loader
	if len(m.Loaders) != 1 {
		t.Fatalf("New() should detect BatchRequest as batch loader, got %d loaders", len(m.Loaders))
	}

	loader := m.Loaders[0]

	// Verify loader properties
	if loader.Method != "LoadItems" {
		t.Errorf("Loader Method = %q, want %q", loader.Method, "LoadItems")
	}

	if loader.RequestType != "BatchRequest" {
		t.Errorf("Loader RequestType = %q, want %q", loader.RequestType, "BatchRequest")
	}

	if loader.ResponseType != "LoadItemsResponse" {
		t.Errorf("Loader ResponseType = %q, want %q", loader.ResponseType, "LoadItemsResponse")
	}

	// BatchRequest uses "Keys" field by default
	if loader.KeysField != "Keys" {
		t.Errorf("Loader KeysField = %q, want %q", loader.KeysField, "Keys")
	}

	// BatchRequest always uses string keys
	if loader.KeysType != "string" {
		t.Errorf("Loader KeysType = %q, want %q", loader.KeysType, "string")
	}

	if loader.ResultsField != "Results" {
		t.Errorf("Loader ResultsField = %q, want %q", loader.ResultsField, "Results")
	}

	if loader.ResultsType != "*Item" {
		t.Errorf("Loader ResultsType = %q, want %q", loader.ResultsType, "*Item")
	}

	// String keys should not require custom key type
	if loader.Custom != false {
		t.Errorf("Loader Custom = %v, want false for string keys", loader.Custom)
	}
}

func TestNew_PackageNameParsing(t *testing.T) {
	tests := []struct {
		name            string
		packageName     string
		wantServiceName string
	}{
		{
			name:            "simple package",
			packageName:     "myservice",
			wantServiceName: "myservice",
		},
		{
			name:            "nested package",
			packageName:     "com.example.myservice",
			wantServiceName: "myservice",
		},
		{
			name:            "deep nested package",
			packageName:     "com.example.v1.myservice",
			wantServiceName: "myservice",
		},
		{
			name:            "empty package",
			packageName:     "",
			wantServiceName: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &descriptorpb.FileDescriptorProto{
				Package:     proto.String(tt.packageName),
				MessageType: []*descriptorpb.DescriptorProto{},
			}

			service := &descriptorpb.ServiceDescriptorProto{
				Name:   proto.String("TestService"),
				Method: []*descriptorpb.MethodDescriptorProto{},
			}

			m := New(service, root)

			if m.ServiceName != tt.wantServiceName {
				t.Errorf("ServiceName = %q, want %q", m.ServiceName, tt.wantServiceName)
			}
		})
	}
}

func TestMessage_Imports_WithoutLoaders(t *testing.T) {
	m := Message{
		Methods: []Method{{Name: "Test", Input: "Request", Output: "Response"}},
		Loaders: []LoaderVars{},
	}

	imports := m.Imports()

	expectedImports := []string{"context", "os"}
	if len(imports) != len(expectedImports) {
		t.Fatalf("Imports() should return %d imports, got %d", len(expectedImports), len(imports))
	}

	for i, exp := range expectedImports {
		if imports[i] != exp {
			t.Errorf("Import[%d] = %q, want %q", i, imports[i], exp)
		}
	}
}

func TestMessage_Imports_WithLoaders(t *testing.T) {
	m := Message{
		Methods: []Method{},
		Loaders: []LoaderVars{{Method: "LoadUser"}},
	}

	imports := m.Imports()

	// Should include os, context, and dataloader
	expectedImports := map[string]bool{
		"os":                                  true,
		"context":                             true,
		"github.com/graph-gophers/dataloader": true,
	}

	if len(imports) != len(expectedImports) {
		t.Fatalf("Imports() with loaders should return %d imports, got %d", len(expectedImports), len(imports))
	}

	for _, imp := range imports {
		if !expectedImports[imp] {
			t.Errorf("Unexpected import: %q", imp)
		}
	}
}

func TestMessage_Generate(t *testing.T) {
	tests := []struct {
		name         string
		message      Message
		wantContains []string
	}{
		{
			name: "basic service",
			message: Message{
				Package:     "test",
				ServiceName: "test",
				Descriptor: &descriptorpb.ServiceDescriptorProto{
					Name: proto.String("UserService"),
				},
				Methods: []Method{
					{Name: "GetUser", Input: "GetUserRequest", Output: "GetUserResponse"},
				},
				Loaders: []LoaderVars{},
			},
			// Note: Generate() now only generates loader helper functions
			// The Module pattern is generated at the file level
			wantContains: []string{},
		},
		{
			name: "service with loaders",
			message: Message{
				Package:     "test",
				ServiceName: "test",
				Descriptor: &descriptorpb.ServiceDescriptorProto{
					Name: proto.String("ProductService"),
				},
				Methods: []Method{},
				Loaders: []LoaderVars{
					{
						Method:       "LoadProduct",
						RequestType:  "LoadProductRequest",
						ResponseType: "LoadProductResponse",
						KeysField:    "Ids",
						KeysType:     "string",
						ResultsField: "Products",
						ResultsType:  "*Product",
						Custom:       false,
					},
				},
			},
			wantContains: []string{
				"ProductServiceLoadProduct(p gql.ResolveParams, key string)",
				"ProductServiceLoadProductMany(p gql.ResolveParams, keys []string)",
			},
		},
		{
			name: "service with custom key loader",
			message: Message{
				Package:     "test",
				ServiceName: "test",
				Descriptor: &descriptorpb.ServiceDescriptorProto{
					Name: proto.String("OrderService"),
				},
				Methods: []Method{},
				Loaders: []LoaderVars{
					{
						Method:       "LoadOrder",
						RequestType:  "LoadOrderRequest",
						ResponseType: "LoadOrderResponse",
						KeysField:    "OrderIds",
						KeysType:     "OrderId",
						ResultsField: "Orders",
						ResultsType:  "*Order",
						Custom:       true,
					},
				},
			},
			wantContains: []string{
				"type OrderIdKey struct",
				"func (key *OrderIdKey) String() string",
				"OrderServiceLoadOrder(p gql.ResolveParams, key *OrderId)",
				"OrderServiceLoadOrderMany(p gql.ResolveParams, keys []*OrderId)",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.message.Generate()

			for _, want := range tt.wantContains {
				if !strings.Contains(got, want) {
					t.Errorf("Generate() should contain %q, but doesn't.\nGot:\n%s", want, got)
				}
			}
		})
	}
}

// Note: DNS configuration, options pattern, and setter functions are now generated
// at the file/module level (in generation/module_gen.go), not per-service.
// These tests have been removed as they no longer apply to Message.Generate().

func TestMethod(t *testing.T) {
	// Test the Method struct
	m := Method{
		Name:   "GetUser",
		Input:  "GetUserRequest",
		Output: "GetUserResponse",
	}

	if m.Name != "GetUser" {
		t.Errorf("Method.Name = %q, want %q", m.Name, "GetUser")
	}
	if m.Input != "GetUserRequest" {
		t.Errorf("Method.Input = %q, want %q", m.Input, "GetUserRequest")
	}
	if m.Output != "GetUserResponse" {
		t.Errorf("Method.Output = %q, want %q", m.Output, "GetUserResponse")
	}
}

func TestLoaderVars(t *testing.T) {
	// Test the LoaderVars struct
	l := LoaderVars{
		Method:       "LoadUser",
		RequestType:  "LoadUserRequest",
		ResponseType: "LoadUserResponse",
		KeysField:    "Ids",
		KeysType:     "string",
		ResultsField: "Users",
		ResultsType:  "*User",
		Custom:       false,
	}

	if l.Method != "LoadUser" {
		t.Errorf("LoaderVars.Method = %q, want %q", l.Method, "LoadUser")
	}
	if l.RequestType != "LoadUserRequest" {
		t.Errorf("LoaderVars.RequestType = %q, want %q", l.RequestType, "LoadUserRequest")
	}
	if l.ResponseType != "LoadUserResponse" {
		t.Errorf("LoaderVars.ResponseType = %q, want %q", l.ResponseType, "LoadUserResponse")
	}
	if l.KeysField != "Ids" {
		t.Errorf("LoaderVars.KeysField = %q, want %q", l.KeysField, "Ids")
	}
	if l.KeysType != "string" {
		t.Errorf("LoaderVars.KeysType = %q, want %q", l.KeysType, "string")
	}
	if l.ResultsField != "Users" {
		t.Errorf("LoaderVars.ResultsField = %q, want %q", l.ResultsField, "Users")
	}
	if l.ResultsType != "*User" {
		t.Errorf("LoaderVars.ResultsType = %q, want %q", l.ResultsType, "*User")
	}
	if l.Custom != false {
		t.Errorf("LoaderVars.Custom = %v, want false", l.Custom)
	}
}
