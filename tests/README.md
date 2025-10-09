# Integration Tests

This directory contains integration tests for the protoc-gen-graphql plugin.

## Test Structure

- `cases/` - Contains proto files used for testing code generation
- `golden/` - Contains expected generated output (golden files)
- `out/` - Temporary directory for generated code (gitignored)
- `generation_test.go` - Main test file using golden file testing pattern

## Running Tests

### Run all tests
```bash
make test
```

Or from the project root:
```bash
go test ./...
```

### Update golden files

When you make intentional changes to the code generator, update the golden files:

```bash
make update-golden
```

Or:
```bash
cd tests && go test -update .
```

## Golden File Testing

This test suite uses the [golden file pattern](https://softwareengineering.stackexchange.com/questions/358786/what-is-golden-file-testing):

1. The test runs `protoc` with the plugin to generate GraphQL code
2. It compares the generated code to a saved "golden" file
3. If they don't match, the test fails and shows a diff
4. When generator output changes intentionally, update golden files with `-update` flag

## Adding New Tests

1. Add a new `.proto` file in `cases/`
2. Run `make update-golden` to generate the initial golden file
3. The test will automatically pick up new proto files in `cases/`
4. Verify the generated code is correct before committing

## Test Cases

The current test cases cover:

- `messages.proto` - Comprehensive test of message types, enums, services, batch loaders, and field options
