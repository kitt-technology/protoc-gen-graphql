package main

import (
	"fmt"
	"github.com/kitt-technology/protoc-gen-graphql/generation"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	_ "github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
	_ "google.golang.org/protobuf/types/pluginpb"
	"io/ioutil"
	"os"
)

func main() {
	bytes, _ := ioutil.ReadAll(os.Stdin)

	var req pluginpb.CodeGeneratorRequest
	proto.Unmarshal(bytes, &req)

	opts := protogen.Options{}
	plugin, _ := opts.New(&req)

	for _, file := range plugin.Files {
		if shouldProcess(file) {
			parsedFile := generation.New(file)
			generateFile := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+".graphql.go", ".")
			generateFile.Write([]byte(parsedFile.ToString()))
		}
	}

	stdout := plugin.Response()
	out, _ := proto.Marshal(stdout)

	fmt.Fprintf(os.Stdout, string(out))
}

func shouldProcess(file *protogen.File) bool {
	ignoredFiles := []string{"graphql/graphql.proto", "graphql.proto", "google/protobuf/descriptor.proto", "google/protobuf/wrappers.proto", "google/protobuf/timestamp.proto", "github.com/kitt-technology/protoc-gen-graphql/graphql/graphql.proto"}
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
