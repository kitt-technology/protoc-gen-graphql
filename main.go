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
	fmt.Println("os stdin", os.Stdin)

	bytes, _ := ioutil.ReadAll(os.Stdin)

	SupportedFeatures := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	var req pluginpb.CodeGeneratorRequest
	err := proto.Unmarshal(bytes, &req)
	if err != nil {
		panic(err)
	}

	opts := protogen.Options{}
	plugin, _ := opts.New(&req)

	fmt.Println("to gen", req.FileToGenerate)
	fmt.Println("proto files", req.ProtoFile)
	for _, file := range plugin.Files {
		fmt.Println("file gopkg: ", file.GoPackageName)
		fmt.Println("file name: ", file.GeneratedFilenamePrefix)
		fmt.Println("file message length: ", len(file.Messages))
	}

	filesToProcess := make([]*protogen.File, 0)
	for _, file := range plugin.Files {
		if shouldProcess(file) {
			filesToProcess = append(filesToProcess, file)
			//parsedFile := generation.New(file)
			//fmt.Println("<<<<", "*1*", parsedFile.Package, "*2*", len(parsedFile.Message), "*3*", file.GeneratedFilenamePrefix, "*4*", len(file.Messages),
			//	"*5*", len(parsedFile.TypeDefs), "*6*", len(parsedFile.Imports),
			//	"*7*", len(parsedFile.ImportMap),
			//	">>>>>")
			//generateFile := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+".graphql.go", ".")
			//_, err = generateFile.Write([]byte(parsedFile.ToString()))
			//if err != nil {
			//	panic(err)
			//}
		}
	}

	parsedFile := generation.NewFromMultiple(filesToProcess)
	fmt.Println("<<<<", "*1*", parsedFile.Package, "*2*", len(parsedFile.Message),
		"*5*", len(parsedFile.TypeDefs), "*6*", len(parsedFile.Imports),
		"*7*", len(parsedFile.ImportMap),
		">>>>>")
	generateFile := plugin.NewGeneratedFile(filesToProcess[0].GeneratedFilenamePrefix+".graphql.go", ".")
	_, err = generateFile.Write([]byte(parsedFile.ToString()))
	if err != nil {
		fmt.Println("<<<<ERROR1>>>>")
		panic(err)
	}

	stdout := plugin.Response()
	stdout.SupportedFeatures = &SupportedFeatures
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
