package main

import (
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
	packageImportMap := make(map[string]string, 0)

	var req pluginpb.CodeGeneratorRequest
	err := proto.Unmarshal(bytes, &req)
	if err != nil {
		panic(err)
	}

	opts := protogen.Options{}
	plugin, _ := opts.New(&req)

	for _, file := range plugin.Files {
		if shouldProcess(file) {
			if proto.HasExtension(file.Proto.Options, graphql.E_Package) {
				// Using graphql.package, could fall back to go_package?
				packageImportMap[*file.Proto.Package] = proto.GetExtension(file.Proto.Options, graphql.E_Package).(string)
			}
			// I'm drilling packageImportMap all the way down here, maybe a better way of doing it?
			parsedFile := generation.New(file, packageImportMap)
			generateFile := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+".graphql.go", ".")
			_, err = generateFile.Write([]byte(parsedFile.ToString()))
			if err != nil {
				panic(err)
			}
		}
	}

	stdout := plugin.Response()
	out, _ := proto.Marshal(stdout)

	_, err = os.Stdout.Write(out)
	if err != nil {
		panic(err)
	}
}

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
