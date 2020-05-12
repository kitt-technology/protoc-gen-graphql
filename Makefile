GO_IMPORT := google/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor

.PHONE: all
all: test

.PHONY: deps
deps:
	@go mod download

.PHONY: build
build:
	@go install .

.PHONY: test
test:
	protoc \
		--proto_path ./graphql \
		-I=./graphql \
		./graphql/graphql.proto \
		--go_out=./graphql/
	rm -rf tests/out || true
	mkdir tests/out/
	@go install .
	protoc \
		--proto_path tests/cases \
		-I=. \
		./tests/cases/messages.proto \
		--go_out=./tests/out \
		--go-grpc_out=./tests/out \
		--graphql_out="lang=go:./tests/out"
	go run tests/run.go
