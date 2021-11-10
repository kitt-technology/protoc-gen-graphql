PACKAGE := github.com/kitt-technology/protoc-gen-graphql

.PHONE: all
all: test

.PHONY: deps
deps:
	@go mod download

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

	git status
	git config --global user.email "you@example.com"
	git config --global user.name "Your Name"
	git stash
	protoc \
		--proto_path . \
		-I=. \
		-I ${GOPATH}/src \
		--go_out="module=${PACKAGE}:./tests/out" \
		--go-grpc_out="module=${PACKAGE}:./tests/out" \
		--graphql_out="module=${PACKAGE},lang=go:./tests/out" \
		./tests/cases/messages.proto
	go run tests/run.go
