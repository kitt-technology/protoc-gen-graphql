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
	go test ./enforce/...
	rm -rf tests/out || true
	mkdir tests/out
	protoc \
		--proto_path ./auth \
		-I=./auth \
		./auth/auth.proto \
		--go_out=./auth/
	@go install .
	protoc \
		--proto_path tests/cases \
		-I=. \
		./tests/cases/messages.proto \
		--go_out=./tests/out \
		--auth_out="lang=go:./tests/out"
	go run tests/run.go