// Package main implements the protoc-gen-graphql plugin.
//
// protoc-gen-graphql is a protoc plugin that generates GraphQL server code
// from Protocol Buffer definitions. It creates GraphQL schemas, types, and
// resolvers that integrate with gRPC services, enabling you to expose your
// gRPC services through a unified GraphQL API.
//
// The plugin supports:
//   - Automatic GraphQL schema generation
//   - DataLoader integration for N+1 query prevention
//   - Batch loading with the batch_loader option
//   - Proto annotations for customizing GraphQL output
//
// Usage with protoc:
//
//	protoc --graphql_out="lang=go:." your.proto
//
// Usage with buf:
//
//	buf generate
package main

import (
	"io"
	"os"
	"strings"

	"github.com/kitt-technology/protoc-gen-graphql/generation"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	SupportedFeatures := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	var req pluginpb.CodeGeneratorRequest
	err = proto.Unmarshal(bytes, &req)
	if err != nil {
		panic(err)
	}

	opts := protogen.Options{}
	plugin, err := opts.New(&req)
	if err != nil {
		panic(err)
	}

	// First pass: Populate GraphqlImportMap from ALL files (including dependencies)
	for _, file := range plugin.Files {
		// Process ALL files to populate GraphqlImportMap, not just files being generated
		if shouldProcess(file) {
			// Just call New() to populate the GraphqlImportMap, don't generate yet
			_ = generation.New(file)
		}
	}

	// Second pass: Generate code with complete GraphqlImportMap
	for _, file := range plugin.Files {
		// Only generate code for files that are part of this generation request,
		// not for imported dependencies
		if !file.Generate {
			continue
		}
		if shouldProcess(file) {
			parsedFile := generation.New(file)
			generateFile := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+".graphql.go", ".")
			if _, err = generateFile.Write([]byte(parsedFile.ToString())); err != nil {
				panic(err)
			}
		}
	}

	stdout := plugin.Response()
	stdout.SupportedFeatures = &SupportedFeatures
	out, err := proto.Marshal(stdout)
	if err != nil {
		panic(err)
	}

	if _, err = os.Stdout.Write(out); err != nil {
		panic(err)
	}
}

// shouldProcess determines whether a proto file should have GraphQL code generated.
// It filters out well-known Google proto files and files marked with the
// (graphql.disabled) option.
func shouldProcess(file *protogen.File) bool {
	filename := *file.Proto.Name

	// Skip Google well-known types
	if strings.HasPrefix(filename, "google/protobuf/") {
		return false
	}

	// Skip graphql.proto itself (can appear as different paths depending on import)
	if strings.HasSuffix(filename, "graphql/graphql.proto") || filename == "graphql.proto" {
		return false
	}

	// Check for disabled option
	if proto.HasExtension(file.Proto.Options, graphql.E_Disabled) {
		if disabled, ok := proto.GetExtension(file.Proto.Options, graphql.E_Disabled).(bool); ok {
			return !disabled
		}
	}
	return true
}
