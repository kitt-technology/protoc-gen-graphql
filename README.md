# protoc-gen-graphql (PGG)

[![CircleCI](https://circleci.com/gh/kitt-technology/protoc-gen-graphql.svg?style=svg)](https://circleci.com/gh/kitt-technology/protoc-gen-graphql)

PGG is a protoc plugin to generate a performant GraphQL server to knit together your gRPC services.

## Installation

Assuming you have `$GOBIN` on your path, the following with enable the generator as a plugin.

```bash
go install github.com/kitt-technology/protoc-gen-graphql
```

## Usage

```bash
protoc \
	-I ${GOPATH}/src \
	--go_out=./out \
	--go-grpc_out=./out \
	--graphql_out="lang=go:./out" \
	--proto_path . \
	./path/to/your/file.proto
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
