package main

import (
    "fmt"
    _ "github.com/kitt-technology/protoc-gen-auth/auth"
    "github.com/kitt-technology/protoc-gen-auth/generation"
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
            generateFile := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix + ".auth.go", ".")
            generateFile.Write([]byte(generation.New(file).ToString()))
        }
    }

    stdout := plugin.Response()
    out, _ := proto.Marshal(stdout)

    fmt.Fprintf(os.Stdout, string(out))
}

func shouldProcess(file *protogen.File) bool {
    ignoredFiles := []string{"auth/auth.proto", "auth.proto", "google/protobuf/descriptor.proto"}
    for _, ignored := range ignoredFiles {
        if *file.Proto.Name == ignored {
            return false
        }
    }
    return true
}