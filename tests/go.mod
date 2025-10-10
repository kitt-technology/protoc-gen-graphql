module github.com/kitt-technology/protoc-gen-graphql/tests

go 1.25.2

require (
	github.com/graph-gophers/dataloader v5.0.0+incompatible
	github.com/graphql-go/graphql v0.8.1
	github.com/kitt-technology/protoc-gen-graphql v0.52.10
	github.com/sergi/go-diff v1.4.0
	google.golang.org/protobuf v1.36.10
)

require (
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250804133106-a7a43d27e69b // indirect
	google.golang.org/grpc v1.76.0 // indirect
)

replace github.com/kitt-technology/protoc-gen-graphql => ../
