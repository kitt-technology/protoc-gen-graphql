# Security Policy

## Supported Versions

We release patches for security vulnerabilities. Currently supported versions:

| Version | Supported          |
| ------- | ------------------ |
| latest  | :white_check_mark: |
| < 1.0   | :x:                |

## Reporting a Vulnerability

We take the security of protoc-gen-graphql seriously. If you believe you have found a security vulnerability, please report it to us as described below.

### Please DO NOT:

- Open a public GitHub issue for security vulnerabilities
- Discuss the vulnerability in public forums, social media, or mailing lists

### Please DO:

**Report security vulnerabilities via GitHub Security Advisories:**

1. Go to the [Security Advisories page](https://github.com/kitt-technology/protoc-gen-graphql/security/advisories)
2. Click "Report a vulnerability"
3. Fill out the form with details about the vulnerability

**What to include in your report:**

- Type of vulnerability (e.g., code injection, information disclosure, etc.)
- Full paths of source file(s) related to the vulnerability
- Location of the affected source code (tag/branch/commit or direct URL)
- Step-by-step instructions to reproduce the issue
- Proof-of-concept or exploit code (if possible)
- Impact of the vulnerability, including how an attacker might exploit it

### What to expect:

- **Initial Response**: We will acknowledge your email within 48 hours
- **Status Updates**: We will send you regular updates about our progress
- **Disclosure Timeline**: We aim to disclose vulnerabilities within 90 days of the initial report
- **Credit**: We will credit you in the security advisory (unless you prefer to remain anonymous)

## Security Best Practices

When using protoc-gen-graphql:

1. **Keep Up to Date**: Always use the latest version of protoc-gen-graphql
2. **Review Generated Code**: Review generated GraphQL code before deploying to production
3. **Input Validation**: Validate all inputs in your gRPC services
4. **Authentication**: Implement proper authentication and authorization
5. **Rate Limiting**: Apply rate limiting to your GraphQL endpoints
6. **Dependency Management**: Keep all dependencies up to date

## Known Security Considerations

- **Code Generation**: This tool generates code based on proto files. Ensure your proto files come from trusted sources
- **Generated Code**: Review generated code for your specific use case
- **gRPC Security**: Follow gRPC security best practices for your backend services

## Security Update Process

When a security vulnerability is confirmed:

1. We will create a patch and release a new version
2. We will publish a GitHub Security Advisory
3. We will update the CHANGELOG.md
4. We will notify users through GitHub releases

## Questions

If you have questions about this security policy, please open a discussion on GitHub.

Thank you for helping keep protoc-gen-graphql and its users safe!