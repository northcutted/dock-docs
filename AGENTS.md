# Docker Docs Agent Guidelines

This document provides instructions and guidelines for AI agents operating within the `docker-docs` repository.

## 1. Build, Lint, and Test

### Build
The project is a Go CLI application.
- **Install dependencies:**
  ```bash
  go mod download
  ```
- **Build the binary:**
  ```bash
  go build -o bin/dock-docs .
  ```
- **Run the binary (example):**
  ```bash
  ./bin/dock-docs --help
  ```

### Test
Tests are written using the standard `testing` package.
- **Run all tests:**
  ```bash
  go test -v ./...
  ```
- **Run a specific package:**
  ```bash
  go test -v ./pkg/analysis
  ```
- **Run a specific test case:**
  ```bash
  go test -v ./pkg/analysis -run TestAnalyzeDockerfile
  ```
- **Run tests with race detection (recommended for concurrency):**
  ```bash
  go test -race -v ./...
  ```

### Lint
The project uses `golangci-lint`.
- **Run linter:**
  ```bash
  golangci-lint run
  ```
- **Fix linting issues (auto-fixable):**
  ```bash
  golangci-lint run --fix
  ```
- **Format code:**
  Always ensure code is formatted with `gofmt` and imports are organized with `goimports` before committing.
  ```bash
  go fmt ./...
  ```

## 2. Code Style & Conventions

### General Go Guidelines
- **Formatting:** Strictly follow `gofmt` standards.
- **Idioms:** Write idiomatic Go code (Effective Go). Avoid trying to write Java/Python in Go.
- **Complexity:** Keep functions small and focused. 

### Naming
- **Variables/Functions:** Use `MixedCaps` or `camelCase`. Exported identifiers must start with an uppercase letter.
- **Constants:** Use `MixedCaps` (e.g., `DefaultConfigPath`), not `ALL_CAPS`.
- **Files:** Use `snake_case.go` (e.g., `docker_analyzer.go`).
- **Tests:** Test files must end in `_test.go`.

### Imports
Group imports in the following order:
1.  Standard library
2.  Third-party packages (e.g., `github.com/...`)
3.  Local packages (e.g., `github.com/northcutted/dock-docs/pkg/...`)

```go
import (
    "fmt"
    "os"

    "github.com/spf13/cobra"

    "github.com/northcutted/dock-docs/pkg/analyzer"
)
```

### Error Handling
- **Return Errors:** Functions that can fail should return `error` as the last return value.
- **Wrap Errors:** Use `fmt.Errorf("...: %w", err)` to wrap errors with context. Do not just return the error raw if you can add context.
- **Check Errors:** Always check returned errors. Do not ignore them using `_` unless strictly necessary and commented.

```go
// Good
if err := doSomething(); err != nil {
    return fmt.Errorf("failed to do something: %w", err)
}
```

### Testing
- **Table-Driven Tests:** Prefer table-driven tests for covering multiple scenarios.
- **Assertions:** Use standard `testing` package. If a simpler assertion library is present in `go.mod` (e.g., `testify`), use it consistently.
- **Mocks:**
  - Use **interface injection** for dependencies (e.g., passing `Runner` interfaces to functions).
  - Use **function variable swapping** for standalone side-effect functions (e.g., `var ensureImage = runner.EnsureImage`).
  - Avoid introducing new mocking libraries (like `gomock` or `testify`) unless necessary; prefer simple struct mocks.

### CLI Structure (Cobra)
- The application uses `spf13/cobra`.
- **Commands:** Defined in `cmd/`.
- **Flags:** logic for flags should be handled in the `init()` function of the command file.
- **Logic:** Keep `cmd/` files thin. Move business logic to packages in `pkg/`.

### Documentation
- **Comments:** Exported functions and types **must** have a comment starting with the function/type name.
- **Godoc:** Ensure comments are compatible with `godoc`.

## 3. Project Specifics

### Architecture
- **`cmd/`**: CLI entry points.
- **`pkg/`**: Core logic libraries.
- **`bin/`**: Output directory for binaries.
- **`samples/`**: Sample projects for dogfooding/testing.

### External Dependencies
This tool wraps external CLIs. To run the binary locally or execute integration tests, you must have the following installed (see `action.yml` for installation scripts):
- `syft` (SBOM generation)
- `grype` (Vulnerability scanning)
- `dive` (Image analysis)

*Note: Unit tests using mocks do not require these tools.*

### CI/CD
- **Workflows:** Located in `.github/workflows/`.
- **Dogfooding:** The `dogfood.yml` workflow generates docs for sample projects to verify functionality.
- **Release:** `goreleaser` is used for builds. Configuration is in `.goreleaser.yaml`.

### Commit Messages
Follow Conventional Commits:
- `feat: ...` for new features
- `fix: ...` for bug fixes
- `docs: ...` for documentation
- `test: ...` for tests
- `chore: ...` for maintenance

Example:
```text
feat: add support for multi-stage dockerfiles
```

## 4. Agent Behavior Rules

- **Safety First:** Never run destructive commands (like `rm -rf /`) without explicit confirmation or very narrow scoping.
- **Verification:** After writing code, ALWAYS run `go build` and `go test` to verify your changes.
- **Context:** Read related files before modifying code to ensure consistency with existing patterns.
- **Minimal Changes:** Only modify what is necessary to complete the task. Avoid formatting changes in unrelated files.
- **Feedback:** If a task is ambiguous, ask the user for clarification before proceeding.
- **Git Push Forbidden:** You may create commits and tags locally if requested, but **NEVER** push to the remote repository without explicit, direct permission from the user for that specific action.
