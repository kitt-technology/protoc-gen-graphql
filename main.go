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
	"strings"
)

func main() {
	panic("POOOOOOOO")
	bytes, _ := ioutil.ReadAll(os.Stdin)

	SupportedFeatures := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	var req pluginpb.CodeGeneratorRequest
	err := proto.Unmarshal(bytes, &req)
	if err != nil {
		panic(err)
	}

	opts := protogen.Options{}
	plugin, _ := opts.New(&req)

	filesGroupedByPackage := make(map[string][]*protogen.File)

	for _, file := range plugin.Files {
		if shouldProcess(file) {
			goPkgNameString := string(file.GoPackageName)
			if _, ok := filesGroupedByPackage[goPkgNameString]; !ok {
				filesGroupedByPackage[goPkgNameString] = make([]*protogen.File, 0)
			}
			filesGroupedByPackage[goPkgNameString] = append(filesGroupedByPackage[goPkgNameString], file)
		}
	}

	for _, packageFiles := range filesGroupedByPackage {
		parsedFile := generation.NewFromMultiple2(packageFiles)
		prefix := packageFiles[0].GeneratedFilenamePrefix

		//if there are multiple files, rename the generated graphql file to 'combined'
		if len(packageFiles) > 1 {
			newPart := "combined"

			// Split the string by "/"
			parts := strings.Split(prefix, "/")

			// Replace the last part
			parts[len(parts)-1] = newPart

			// Join the parts back together
			prefix = strings.Join(parts, "/")
		}

		generateFile := plugin.NewGeneratedFile(prefix+".graphql.go", ".")
		_, err = generateFile.Write([]byte(parsedFile.ToString()))
		if err != nil {
			panic(err)
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
