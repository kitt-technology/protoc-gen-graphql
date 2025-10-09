PACKAGE := github.com/kitt-technology/protoc-gen-graphql

.PHONY: all
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
		--proto_path ./example/users/ \
		-I . \
		-I ./graphql \
		-I ./example \
		-I ${GOPATH}/src \
		./example/users/users.proto \
		--go_out=. \
		--go-grpc_out=. \
		--graphql_out="lang=go:."
	protoc \
		--proto_path ./example/products \
		-I . \
		-I ./graphql \
		-I ./example \
		-I ${GOPATH}/src \
		./example/products/products.proto \
		--go_out=. \
		--go-grpc_out=. \
		--graphql_out="lang=go:."
	protoc \
		--proto_path ./example/common-example \
		-I . \
		-I ./graphql \
		-I ${GOPATH}/src \
		./example/common-example/common-example.proto \
		--go_out=. \
		--go-grpc_out=. \
		--graphql_out="lang=go:."
	rm -rf ./github.com

.PHONY: docker
docker:
	docker build . -t kittoffices/protoc-gen-graphql --platform linux/amd64
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
	go install .
	cd tests && go test -v .

.PHONY: update-golden
update-golden:
	go install .
	cd tests && go test -v -update .
