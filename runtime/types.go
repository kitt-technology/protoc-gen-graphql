// Package runtime provides plain Go types that mirror the protobuf types in the graphql package.
// These types are provided for reference and documentation purposes, allowing users to
// understand the structure without importing the graphql package (which contains .proto files).
//
// Note: The actual gRPC service implementations should use the proto-generated types from
// the graphql package (github.com/kitt-technology/protoc-gen-graphql/graphql) for proper
// protobuf serialization. These runtime types are structurally identical but are not suitable
// for gRPC communication.
package runtime

// BatchRequest is a common request type for batch loaders.
// It contains a list of string keys to batch load.
//
// This type mirrors graphql.BatchRequest from the proto definition:
//
//	message BatchRequest {
//	  repeated string keys = 1;
//	}
//
// Use this type for reference, but gRPC service implementations should use
// graphql.BatchRequest from github.com/kitt-technology/protoc-gen-graphql/graphql.
type BatchRequest struct {
	Keys []string
}

// PageInfo contains pagination information.
//
// This type mirrors graphql.PageInfo from the proto definition:
//
//	message PageInfo {
//	  int32 total_count = 1;
//	  string end_cursor = 2;
//	  bool has_next_page = 3;
//	}
type PageInfo struct {
	TotalCount  int32
	EndCursor   string
	HasNextPage bool
}

// FieldMask is used for field selection optimization.
//
// This type mirrors graphql.FieldMask from the proto definition:
//
//	message FieldMask {
//	  repeated string paths = 1;
//	  map<string, bool> paths_map = 2;
//	}
type FieldMask struct {
	Paths    []string
	PathsMap map[string]bool
}