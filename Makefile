PACKAGE := github.com/kitt-technology/protoc-gen-graphql

.PHONE: all
all: test

.PHONY: deps
deps:
	GO111MODULE=off go get github.com/kitt-technology/protoc-gen-graphql
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.11

.PHONY: build
build:
	@go install .

.PHONY: examples
build-examples:
	@go install .


	protoc \
		--proto_path ./example/authors/ \
		-I . \
		-I ${GOPATH}/src \
		./example/authors/authors.proto \
		--go_out=. \
		--go-grpc_out=. \
		--graphql_out="lang=go:."
	protoc \
		--proto_path ./example/books \
		-I . \
		-I ${GOPATH}/src \
		./example/books/books.proto \
		--go_out=. \
		--go-grpc_out=. \
		--graphql_out="lang=go:."
	protoc \
		--proto_path ./example/common-example \
		-I . \
		-I ${GOPATH}/src \
		./example/common-example/common-example.proto \
		--go_out=. \
		--go-grpc_out=. \
		--graphql_out="lang=go:."
	rm -rf ./github.com

run-examples:
	cd example; parallel -u ::: 'go run ./cmd/books' 'go run ./cmd/authors' 'go run ./cmd/graphql'; cd -

.PHONY: docker
docker:
	docker build . -t kittoffices/protoc-gen-graphql
	docker push kittoffices/protoc-gen-graphql

.PHONY: generate
generate:
	protoc \
		--proto_path . \
		-I=. \
		--graphql_out="module=${PACKAGE}:./" \
		--go_out="module=${PACKAGE}:./" \
		graphql/graphql.proto

.PHONY: clone
clone:
	GO111MODULE=off go get -d github.com/kitt-technology/protoc-gen-graphql \
	  && cd ${GOPATH}/src/github.com/kitt-technology/protoc-gen-graphql \
	  && echo ${GOPATH}

.PHONY: test
test:
	rm -rf tests/out || true
	mkdir tests/out/
	go install .
	protoc \
		--proto_path . \
		-I=. \
		-I=./example/common-example \
		-I ${GOPATH}/src \
		--go_out="module=${PACKAGE}:./tests/out" \
		--go-grpc_out="module=${PACKAGE}:./tests/out" \
		--graphql_out="module=${PACKAGE},lang=go:./tests/out" \
		./tests/cases/messages.proto
	go fmt tests/out/cases/messages.graphql.go
	go run tests/run.go
