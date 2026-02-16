# Contributing to dock-docs

Thank you for your interest in contributing to dock-docs! We welcome contributions of all kinds, including bug reports, feature requests, documentation improvements, and code contributions.

## Table of Contents

- [Contributing to dock-docs](#contributing-to-dock-docs)
  - [Table of Contents](#table-of-contents)
  - [Code of Conduct](#code-of-conduct)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation Scripts](#installation-scripts)
  - [Development Setup](#development-setup)
  - [Making Changes](#making-changes)
  - [Testing](#testing)
    - [Running Tests](#running-tests)
    - [Writing Tests](#writing-tests)
  - [Code Style](#code-style)
    - [General Guidelines](#general-guidelines)
    - [Error Handling](#error-handling)
    - [Documentation](#documentation)
  - [Commit Messages](#commit-messages)
  - [Pull Request Process](#pull-request-process)
    - [PR Review Checklist](#pr-review-checklist)
  - [Issue Templates](#issue-templates)
  - [Getting Help](#getting-help)
  - [License](#license)

## Code of Conduct

This project adheres to the [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. Please report unacceptable behavior by opening an issue or contacting the maintainers.

## Getting Started

### Prerequisites

- **Go** (>= 1.25)
- **Git**
- **External tools** (for integration testing):
  - `syft` - SBOM generation
  - `grype` - Vulnerability scanning
  - `dive` - Image analysis

Note: Unit tests with mocks do not require external tools.

### Installation Scripts

See `.github/workflows/action.yml` for installation scripts for the external tools.

## Development Setup

1. **Fork and clone the repository:**
   ```bash
   git clone https://github.com/YOUR_USERNAME/dock-docs.git
   cd dock-docs
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Build the binary:**
   ```bash
   go build -o bin/dock-docs .
   ```

4. **Run the binary:**
   ```bash
   ./bin/dock-docs --help
   ```

5. **Verify your setup:**
   ```bash
   go test -v ./...
   ```

## Making Changes

1. **Create a new branch** from `main`:
   ```bash
   git checkout -b feature/my-feature
   ```
   
   Use branch naming conventions:
   - `feature/` - New features
   - `fix/` - Bug fixes
   - `docs/` - Documentation changes
   - `test/` - Test improvements
   - `chore/` - Maintenance tasks

2. **Make your changes** following the [Code Style](#code-style) guidelines.

3. **Add or update tests** for your changes.

4. **Run tests locally** to ensure everything passes:
   ```bash
   go test -v ./...
   go test -race -v ./...  # Check for race conditions
   ```

5. **Format your code:**
   ```bash
   go fmt ./...
   ```

6. **Run the linter:**
   ```bash
   golangci-lint run
   golangci-lint run --fix  # Auto-fix issues
   ```

## Testing

### Running Tests

- **All tests:**
  ```bash
  go test -v ./...
  ```

- **Specific package:**
  ```bash
  go test -v ./pkg/analysis
  ```

- **Specific test case:**
  ```bash
  go test -v ./pkg/analysis -run TestAnalyzeDockerfile
  ```

- **With race detection:**
  ```bash
  go test -race -v ./...
  ```

### Writing Tests

- Use **table-driven tests** for multiple scenarios
- Use the standard `testing` package
- For mocks:
  - Use **interface injection** for dependencies
  - Use **function variable swapping** for standalone functions
  - Avoid introducing new mocking libraries

Example:
```go
func TestMyFunction(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    string
        wantErr bool
    }{
        {"valid input", "test", "result", false},
        {"empty input", "", "", true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := MyFunction(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("MyFunction() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("MyFunction() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

## Code Style

### General Guidelines

- **Formatting:** Strictly follow `gofmt` standards
- **Imports:** Organize in three groups:
  1. Standard library
  2. Third-party packages
  3. Local packages

  ```go
  import (
      "fmt"
      "os"

      "github.com/spf13/cobra"

      "github.com/northcutted/dock-docs/pkg/analyzer"
  )
  ```

- **Naming:**
  - Variables/Functions: `MixedCaps` or `camelCase`
  - Exported identifiers: Start with uppercase
  - Constants: `MixedCaps` (e.g., `DefaultConfigPath`)
  - Files: `snake_case.go`
  - Test files: `_test.go`

### Error Handling

- Always return errors as the last return value
- Wrap errors with context using `fmt.Errorf`:
  ```go
  if err := doSomething(); err != nil {
      return fmt.Errorf("failed to do something: %w", err)
  }
  ```
- Always check returned errors (don't use `_` unless strictly necessary)

### Documentation

- **All exported functions and types must have comments**
- Comments should start with the function/type name
- Ensure comments are compatible with `godoc`

Example:
```go
// AnalyzeDockerfile parses a Dockerfile and returns analysis results.
func AnalyzeDockerfile(path string) (*Analysis, error) {
    // ...
}
```

## Commit Messages

Follow [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` - New features
- `fix:` - Bug fixes
- `docs:` - Documentation changes
- `test:` - Test additions or modifications
- `refactor:` - Code refactoring
- `chore:` - Maintenance tasks
- `style:` - Formatting changes

Examples:
```
feat: add support for multi-stage dockerfiles
fix: handle empty dockerfile edge case
docs: update installation instructions
test: add coverage for config package
```

## Pull Request Process

1. **Ensure all tests pass** and the code builds successfully:
   ```bash
   go build -o bin/dock-docs .
   go test -v ./...
   ```

2. **Update documentation** if needed (README.md, SPEC.md, etc.)

3. **Push your branch** to your fork:
   ```bash
   git push origin feature/my-feature
   ```

4. **Open a Pull Request** on GitHub with:
   - Clear title following commit message conventions
   - Description of changes and motivation
   - Reference to related issues (e.g., "Fixes #123")
   - Screenshots/examples if applicable

5. **Respond to review feedback** and update your PR as needed

6. **Squash commits** if requested by maintainers

### PR Review Checklist

Before submitting, ensure your PR:
- [ ] Builds successfully (`go build`)
- [ ] Passes all tests (`go test -v ./...`)
- [ ] Has no linting errors (`golangci-lint run`)
- [ ] Includes tests for new functionality
- [ ] Updates relevant documentation
- [ ] Follows code style guidelines
- [ ] Uses conventional commit messages

## Issue Templates

When opening an issue, please use the appropriate template:

- **Bug Report** - Report a defect or unexpected behavior
- **Feature Request** - Suggest a new feature or enhancement
- **Question** - Ask for help or clarification

This helps us understand and address your concern more efficiently.

## Getting Help

If you have questions or need help:

1. Check existing [issues](https://github.com/northcutted/dock-docs/issues)
2. Read the [documentation](README.md) and [specification](SPEC.md)
3. Open a [question issue](https://github.com/northcutted/dock-docs/issues/new?template=question.md)

## License

By contributing to dock-docs, you agree that your contributions will be licensed under the MIT License.
