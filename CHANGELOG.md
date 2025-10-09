git# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- GitHub Actions CI/CD workflows
- GoReleaser configuration for cross-platform releases
- Comprehensive GitHub issue and PR templates
- Security policy (SECURITY.md)
- Dependabot configuration for automated dependency updates
- golangci-lint configuration
- CONTRIBUTING.md with development guidelines
- MIT License file
- Modern multi-stage Dockerfile

### Changed
- Replaced deprecated `ioutil` package with `io` and `os`
- Modernized Dockerfile from Ubuntu Xenial to Alpine-based multi-stage build
- Updated protoc include paths to support simplified imports

### Fixed
- Fixed `.PHONE` typo to `.PHONY` in Makefile

## Previous Releases

For changes prior to this changelog, please see the [commit history](https://github.com/kitt-technology/protoc-gen-graphql/commits/master).

---

## Release Process

1. Update this CHANGELOG.md with the new version
2. Create a git tag: `git tag -a v0.1.0 -m "Release v0.1.0"`
3. Push the tag: `git push origin v0.1.0`
4. GitHub Actions will automatically create a release with binaries