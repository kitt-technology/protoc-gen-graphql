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
	"github.com/kitt-technology/protoc-gen-graphql/generation"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	_ "github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
	_ "google.golang.org/protobuf/types/pluginpb"
	"io"
	"os"
)

func main() {
	bytes, _ := io.ReadAll(os.Stdin)

	SupportedFeatures := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	var req pluginpb.CodeGeneratorRequest
	err := proto.Unmarshal(bytes, &req)
	if err != nil {
		panic(err)
	}

	opts := protogen.Options{}
	plugin, _ := opts.New(&req)

	for _, file := range plugin.Files {
		if shouldProcess(file) {
			parsedFile := generation.New(file)
			generateFile := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+".graphql.go", ".")
			_, err = generateFile.Write([]byte(parsedFile.ToString()))
			if err != nil {
				panic(err)
			}
		}
	}

	stdout := plugin.Response()
	stdout.SupportedFeatures = &SupportedFeatures
	out, _ := proto.Marshal(stdout)

	_, err = os.Stdout.Write(out)
	if err != nil {
		panic(err)
	}
}

// shouldProcess determines whether a proto file should have GraphQL code generated.
// It filters out well-known Google proto files and files marked with the
// (graphql.disabled) option.
func shouldProcess(file *protogen.File) bool {
	ignoredFiles := []string{
		"google/protobuf/descriptor.proto",
		"google/protobuf/wrappers.proto",
		"google/protobuf/timestamp.proto",
		"github.com/kitt-technology/protoc-gen-graphql/graphql/graphql.proto",
	}
	for _, ignored := range ignoredFiles {
		if *file.Proto.Name == ignored {
			return false
		}
	}
	if proto.HasExtension(file.Proto.Options, graphql.E_Disabled) {
		return !proto.GetExtension(file.Proto.Options, graphql.E_Disabled).(bool)
	}
	return true
}
