package main

import (
	"testing"

	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestShouldProcess(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		disabled *bool
		want     bool
	}{
		{
			name:     "regular proto file",
			filename: "myservice.proto",
			disabled: nil,
			want:     true,
		},
		{
			name:     "google protobuf descriptor",
			filename: "google/protobuf/descriptor.proto",
			disabled: nil,
			want:     false,
		},
		{
			name:     "google protobuf wrappers",
			filename: "google/protobuf/wrappers.proto",
			disabled: nil,
			want:     false,
		},
		{
			name:     "google protobuf timestamp",
			filename: "google/protobuf/timestamp.proto",
			disabled: nil,
			want:     false,
		},
		{
			name:     "graphql proto itself",
			filename: "github.com/kitt-technology/protoc-gen-graphql/graphql/graphql.proto",
			disabled: nil,
			want:     false,
		},
		{
			name:     "file with disabled option true",
			filename: "myservice.proto",
			disabled: proto.Bool(true),
			want:     false,
		},
		{
			name:     "file with disabled option false",
			filename: "myservice.proto",
			disabled: proto.Bool(false),
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileDesc := &descriptorpb.FileDescriptorProto{
				Name: proto.String(tt.filename),
			}

			if tt.disabled != nil {
				opts := &descriptorpb.FileOptions{}
				proto.SetExtension(opts, graphql.E_Disabled, *tt.disabled)
				fileDesc.Options = opts
			}

			file := &protogen.File{
				Proto: fileDesc,
			}

			got := shouldProcess(file)
			if got != tt.want {
				t.Errorf("shouldProcess() = %v, want %v", got, tt.want)
			}
		})
	}
}
