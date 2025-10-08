# Contributing to protoc-gen-graphql

Thank you for your interest in contributing to protoc-gen-graphql! We welcome contributions from the community.

## Getting Started

### Prerequisites

- Go 1.23 or later
- `protoc` (Protocol Buffers compiler) version 3.15.0 or later
- Make

### Development Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/kitt-technology/protoc-gen-graphql.git
   cd protoc-gen-graphql
   ```

2. **Install dependencies**
   ```bash
   make deps
   ```

3. **Build the plugin**
   ```bash
   make build
   ```

4. **Run tests**
   ```bash
   make test
   ```

## Development Workflow

### Making Changes

1. **Create a new branch**
   ```bash
   git checkout -b feature/your-feature-name
   # or
   git checkout -b fix/your-bug-fix
   ```

2. **Make your changes**
   - Write clear, readable code
   - Follow Go conventions and idioms
   - Add tests for new functionality
   - Update documentation as needed

3. **Test your changes**
   ```bash
   # Run tests
   make test

   # Run linter (if installed)
   golangci-lint run

   # Build examples to ensure nothing breaks
   make build-examples
   ```

4. **Commit your changes**
   ```bash
   git add .
   git commit -m "Brief description of your changes"
   ```

   Follow these commit message guidelines:
   - Use the imperative mood ("Add feature" not "Added feature")
   - Keep the first line under 72 characters
   - Reference issues and pull requests where appropriate

### Testing

- All code changes should include tests
- Tests should cover both happy paths and error cases
- Golden file tests are located in `tests/cases/`
- To update golden files after intentional changes:
  ```bash
  make test
  # If the output is correct, copy from tests/out/cases/ to tests/cases/
  ```

### Code Style

- Run `gofmt` on your code (or use `goimports`)
- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Keep functions small and focused
- Write meaningful variable names
- Add comments for exported functions and types

### Running the Linter

If you have `golangci-lint` installed:

```bash
golangci-lint run
```

This will check for common issues, code style problems, and potential bugs.

## Submitting Changes

### Pull Request Process

1. **Push your branch**
   ```bash
   git push origin feature/your-feature-name
   ```

2. **Open a Pull Request**
   - Go to the GitHub repository
   - Click "New Pull Request"
   - Select your branch
   - Fill out the PR template with:
     - Description of changes
     - Related issues (if any)
     - Testing performed
     - Screenshots (if applicable)

3. **Code Review**
   - A maintainer will review your PR
   - Address any feedback or requested changes
   - Once approved, your PR will be merged

### PR Guidelines

- Keep PRs focused on a single feature or fix
- Update the README.md if you're adding new features
- Ensure all tests pass
- Add tests for new functionality
- Update documentation as needed

## Project Structure

```
protoc-gen-graphql/
├── main.go              # Entry point for the protoc plugin
├── generation/          # Code generation logic
├── graphql/             # GraphQL proto definitions and generated code
├── example/             # Example proto files and servers
│   ├── authors/
│   ├── books/
│   ├── common-example/
│   └── cmd/            # Example server implementations
├── tests/              # Test files
│   ├── cases/          # Golden files for comparison
│   └── out/            # Generated test output
└── Makefile            # Build automation
```

## Adding New Features

When adding a new feature:

1. **Discuss first** - For major changes, open an issue to discuss your idea
2. **Add proto definitions** - Update `graphql/graphql.proto` if adding new options
3. **Update generation logic** - Modify code in `generation/` package
4. **Add tests** - Create test cases in `tests/cases/`
5. **Add examples** - Demonstrate usage in the `example/` directory
6. **Update documentation** - Add to README.md and code comments

## Publishing to Buf Schema Registry

If you're a maintainer publishing proto definitions:

```bash
cd graphql
buf push
```

## Questions?

- Open an issue for bugs or feature requests
- Start a discussion for questions or ideas
- Check existing issues before creating a new one

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

Thank you for contributing to protoc-gen-graphql! 🎉