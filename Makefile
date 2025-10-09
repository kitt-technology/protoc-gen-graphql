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

.PHONY: regenerate-examples
regenerate-examples: build
	@echo "🔄 Regenerating examples with buf..."
	@buf generate --path example/
	@echo "✅ Examples regenerated successfully!"

.PHONY: check-examples
check-examples: build
	@echo "🔍 Checking if examples are up-to-date..."
	@buf generate --path example/
	@if git diff --quiet example/; then \
		echo "✅ Examples are up-to-date!"; \
	else \
		echo "❌ Examples are out of date. Run 'make regenerate-examples' to update them."; \
		git diff example/; \
		exit 1; \
	fi

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
