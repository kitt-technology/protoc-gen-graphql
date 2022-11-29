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


## Release

1. Commit, merge, and pull
2. Tag the new version: `git tag v0.X.X` (NOTE: has to be triple digit version v.0.X.X)
3. Push the tag: `git push --tags`
4. Update `KITT_REPO/build/common/docker/deps/Dockerfile`:
   ```
   ENV GEN_GRAPHQL_VERSION v0.X.X
   ```
5. Rebuild the docker image:
   ```
   docker build --build-arg GITHUB_TOKEN -t gcr.io/kitt-220208/deps .
   ```
6. Push that docker image (optional):
   ```
   docker push gcr.io/kitt-220208/deps
   ```

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
