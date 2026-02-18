# Project Specification: dock-docs

**Module:** `github.com/northcutted/dock-docs`
**Language:** Go 1.25+
**CLI Framework:** `github.com/spf13/cobra`
**Parser:** `github.com/moby/buildkit/frontend/dockerfile/parser` (AST-based, not regex)
**Templating:** `text/template` (standard library)

## 1. Overview

`dock-docs` is a CLI tool that generates documentation from Dockerfiles and container images. It performs two types of analysis:

- **Static analysis** — Parses a Dockerfile's AST to extract `ARG`, `ENV`, `LABEL`, and `EXPOSE` instructions along with metadata from "magic comments."
- **Dynamic analysis** — Shells out to external tools (`docker`/`podman`, `syft`, `grype`, `dive`) to collect runtime image stats: size, layers, packages, vulnerabilities, and efficiency.

The generated documentation is rendered via Go templates and either injected into an existing Markdown file between markers or written as a standalone HTML/JSON file.

## 2. Operating Modes

### CLI Mode

Runs without a configuration file. Analyzes a single Dockerfile and optionally a single image.

```
dock-docs -f ./Dockerfile --image my-app:latest -o README.md
```

Auto-detection: if no `dock-docs.yaml` exists in the working directory and `--config` is not specified, CLI Mode is used.

### YAML Mode

Uses a `dock-docs.yaml` configuration file for multi-section, multi-image documentation.

```
dock-docs --config dock-docs.yaml
```

Auto-detection: if `dock-docs.yaml` exists in the working directory, YAML Mode is used automatically. The `--config` flag overrides auto-detection.

When a config file path includes a directory component, the working directory is changed to the config file's parent directory so that relative paths in the config resolve correctly.

## 3. CLI Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--file` | `-f` | `./Dockerfile` | Path to Dockerfile (CLI Mode only) |
| `--output` | `-o` | `README.md` | Path to output file |
| `--dry-run` | | `false` | Print to stdout instead of writing |
| `--image` | | `""` | Docker image tag to analyze (CLI Mode only) |
| `--config` | | `""` | Path to config file |
| `--nomoji` | | `false` | Disable emojis in output (text alternatives) |
| `--ignore-errors` | | `false` | Continue on analysis errors |
| `--verbose` | | `false` | Enable verbose/debug logging |
| `--badge-base-url` | | `https://img.shields.io/static/v1` | Base URL for badge generation |
| `--template` | | `""` | Template name or file path |
| `--list-templates` | | `false` | List all built-in templates |
| `--export-template` | | `""` | Export a built-in template to stdout |
| `--validate-template` | | `""` | Validate a custom template file |
| `--debug-template` | | `false` | Print template resolution info |
| `--version` | `-v` | | Print version |

### Subcommands

#### `dock-docs setup`

Downloads and installs external tool dependencies from upstream GitHub Releases.

| Flag | Default | Description |
|------|---------|-------------|
| `--dir` | `~/.dock-docs/bin` | Install directory |
| `--force` | `false` | Reinstall even if present |
| `--check` | `false` | Show status without installing |

## 4. Configuration Schema (YAML Mode)

File: `dock-docs.yaml`

```yaml
output: "README.md"                    # Output file path (default: "README.md")
badgeBaseURL: "https://..."            # Badge URL base (default: shields.io static/v1)

template:                              # Global template override (optional)
  name: "default"                      # Built-in template name
  path: "./custom.tmpl"                # OR path to custom template file

sections:
  - type: "image"                      # Single-image documentation
    marker: "main-docs"                # Injection marker name
    source: "Dockerfile"               # Dockerfile path (default: "Dockerfile")
    tag: "my-app:latest"               # Image tag for dynamic analysis (optional)
    template:                          # Section-level template override (optional)
      name: "detailed"

  - type: "comparison"                 # Multi-image comparison table
    marker: "tag-comparison"           # Injection marker name
    source: "Dockerfile"               # Shared Dockerfile (inherited by entries without source)
    details: false                     # Show per-image detail sections (default: false)
    images:                            # List of images to compare
      - source: "Dockerfile.alpine"    # Per-entry Dockerfile override (optional)
        tag: "my-app:alpine"
      - tag: "my-app:slim"             # Inherits section-level source
    template:
      name: "default"
```

### Section Types

#### `image`

Generates documentation for a single Dockerfile/image combination. Fields:

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `type` | `"image"` | yes | Section type discriminator |
| `marker` | `string` | yes | Named injection marker |
| `source` | `string` | no | Dockerfile path (default: `"Dockerfile"`) |
| `tag` | `string` | no | Image tag for dynamic analysis |
| `template` | `TemplateConfig` | no | Template override |

#### `comparison`

Generates a comparison table across multiple images. Fields:

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `type` | `"comparison"` | yes | Section type discriminator |
| `marker` | `string` | yes | Named injection marker |
| `source` | `string` | no | Shared Dockerfile for entries that omit `source` |
| `images` | `[]ImageEntry` | yes | List of images to compare |
| `details` | `bool` | no | Show per-image collapsed details (default: `false`) |
| `template` | `TemplateConfig` | no | Template override |

Each `ImageEntry` has:

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `source` | `string` | no | Dockerfile path (inherits from section if omitted) |
| `tag` | `string` | yes | Image tag to analyze |

### Template Resolution Order

1. CLI `--template` flag (highest priority)
2. Section-level `template:` in YAML
3. Global `template:` in YAML
4. `"default"` built-in

## 5. Architecture

### Directory Structure

```
.
├── main.go                          # Entry point: cmd.Execute()
├── cmd/
│   ├── root.go                      # CLI definition, runCLIMode(), runYAMLMode()
│   ├── setup.go                     # `dock-docs setup` subcommand
│   ├── version.go                   # Version variable
│   └── root_test.go                 # CLI tests (template resolution, output paths)
├── pkg/
│   ├── analysis/
│   │   ├── analyzer.go              # AnalyzeImage(), AnalyzeComparison(), mergeStats()
│   │   └── analyzer_test.go
│   ├── config/
│   │   ├── config.go                # Config, Section, ImageEntry, Load()
│   │   └── config_test.go
│   ├── injector/
│   │   ├── injector.go              # Inject() — marker-based content replacement
│   │   └── injector_test.go
│   ├── installer/
│   │   ├── installer.go             # Install(), InstallAll(), FindTool(), Status()
│   │   └── installer_test.go
│   ├── parser/
│   │   ├── parser.go                # Parse() — Dockerfile AST walking
│   │   └── parser_test.go
│   ├── renderer/
│   │   ├── markdown.go              # RenderWithTemplate(), RenderComparisonWithTemplate()
│   │   └── markdown_test.go
│   ├── runner/
│   │   ├── runner.go                # RuntimeRunner, ManifestRunner, SyftRunner, GrypeRunner, DiveRunner
│   │   └── runner_test.go
│   ├── templates/
│   │   ├── loader.go                # Loader, LoadBuiltin(), LoadFile(), embed FS
│   │   ├── funcs.go                 # GetFuncMap(), template helper functions
│   │   ├── security.go              # ExecuteWithLimits(), SecurityConfig
│   │   ├── loader_test.go
│   │   └── builtin/
│   │       ├── markdown/            # 8 templates (4 image + 4 comparison)
│   │       ├── html/                # 2 templates (dashboard + dashboard_comparison)
│   │       └── json/                # 2 templates (summary + summary_comparison)
│   └── types/
│       ├── stats.go                 # ImageStats, PackageSummary, Vulnerability, badge helpers
│       └── stats_test.go
├── samples/                         # Sample projects for dogfooding
│   ├── go/
│   ├── node/
│   └── python/
├── .github/workflows/
│   ├── test.yml                     # CI: test + lint + coverage
│   ├── release.yml                  # GoReleaser on v* tags
│   └── dogfood.yml                  # Regenerate sample docs after release
├── action.yml                       # Composite GitHub Action
├── Dockerfile                       # Multi-stage build (golang:1.25-alpine -> distroless)
├── .goreleaser.yaml                 # GoReleaser v2 config
└── dock-docs.yaml                   # (in samples/) Config files
```

### Package Responsibilities

| Package | Responsibility |
|---------|---------------|
| `cmd` | CLI definition (Cobra), flag parsing, orchestration of modes. Thin layer delegating to `pkg/`. |
| `pkg/parser` | Dockerfile AST parsing using Moby BuildKit. Extracts `ARG`, `ENV`, `LABEL`, `EXPOSE` with magic comment metadata. |
| `pkg/analysis` | Orchestrates runners. `AnalyzeImage()` runs all available runners in parallel via goroutines, merges results. `AnalyzeComparison()` runs `AnalyzeImage()` for multiple tags in parallel via `errgroup`. |
| `pkg/runner` | External tool integrations. Each runner implements `ToolRunner` and shells out via `os/exec`. Parses JSON output. |
| `pkg/renderer` | Template loading and execution. Builds context objects and delegates to the template system. |
| `pkg/templates` | Template embedding (`//go:embed`), loading (built-in and custom file), caching, validation, function map, and security-limited execution. |
| `pkg/types` | Shared data types: `ImageStats`, `PackageSummary`, `Vulnerability`. Badge URL generation helpers. |
| `pkg/config` | YAML config file parsing and defaults. Section type constants, template resolution. |
| `pkg/injector` | Marker-based content injection into existing files. |
| `pkg/installer` | Downloads tools from GitHub Releases. Manages `~/.dock-docs/bin/` fallback directory. |

## 6. Data Types

### Parser Types

```go
// parser.DocItem — A single documented instruction from a Dockerfile.
type DocItem struct {
    Name        string // Variable/port name (e.g., "PORT")
    Value       string // Default value inferred from instruction (or @default override)
    Description string // From @description magic comment
    Type        string // "ARG", "ENV", "LABEL", "EXPOSE"
    Required    bool   // From @required magic comment
}

// parser.Documentation — Collection of parsed items with filtering.
type Documentation struct {
    Items []DocItem
}

func (d *Documentation) FilterByType(t string) []DocItem
```

### Image Analysis Types

```go
// types.ImageStats — Aggregated analysis results for a single image.
type ImageStats struct {
    ImageTag               string
    Architecture           string
    SupportedArchitectures []string        // From manifest inspect
    OS                     string
    OSDistro               string          // From Syft (e.g., "Alpine Linux 3.18")
    SizeBytes              int64           // Raw size in bytes; use SizeMB() for formatted string
    TotalLayers            int
    Efficiency             float64         // From Dive (0-100 percentage)
    WastedBytes            string
    TotalPackages          int
    Packages               []PackageSummary
    Vulnerabilities        []Vulnerability // Sorted by severity desc
    VulnSummary            map[string]int  // Severity -> Count
    VulnScanTime           time.Time

    // Badge helper methods:
    // SizeBadge(baseURL) string
    // LayersBadge(baseURL) string
    // EfficiencyBadge(baseURL) string
    // VulnBadge(baseURL) string
    // TotalVulns() int
}

type PackageSummary struct {
    Name    string
    Version string
}

type Vulnerability struct {
    ID       string // e.g., "CVE-2023-1234"
    Severity string // "Critical", "High", "Medium", "Low"
    Package  string
    Version  string
}
```

### Config Types

```go
type SectionType string
const (
    SectionTypeImage      SectionType = "image"
    SectionTypeComparison SectionType = "comparison"
)

type Config struct {
    Output       string          `yaml:"output"`
    BadgeBaseURL string          `yaml:"badgeBaseURL,omitempty"`
    Sections     []Section       `yaml:"sections"`
    Template     *TemplateConfig `yaml:"template,omitempty"`
}

type Section struct {
    Type     SectionType     `yaml:"type"`
    Marker   string          `yaml:"marker"`
    Source   string          `yaml:"source,omitempty"`
    Tag      string          `yaml:"tag,omitempty"`
    Images   []ImageEntry    `yaml:"images,omitempty"`
    Details  bool            `yaml:"details,omitempty"`
    Template *TemplateConfig `yaml:"template,omitempty"`
}

type ImageEntry struct {
    Source string `yaml:"source,omitempty"`
    Tag    string `yaml:"tag"`
}

type TemplateConfig struct {
    Name string `yaml:"name,omitempty"`
    Path string `yaml:"path,omitempty"`
}
```

### Renderer Context Types

```go
// renderer.ReportContext — Passed to image templates.
type ReportContext struct {
    Doc      *parser.Documentation
    Stats    *types.ImageStats      // nil if no --image / no tag
    ImageTag string
    Options  RenderOptions
}

// renderer.ComparisonContext — Passed to comparison templates.
type ComparisonContext struct {
    Images  []*types.ImageStats
    Options RenderOptions
}

// Both have: Emoji(name string) string

type RenderOptions struct {
    NoMoji       bool
    BadgeBaseURL string
}

type TemplateSelection struct {
    Name string // Built-in template name
    Path string // Custom template file path (overrides Name)
}
```

## 7. Dockerfile Parsing

The parser uses `github.com/moby/buildkit/frontend/dockerfile/parser` to walk the Dockerfile AST. It extracts:

| Instruction | Extracted Fields |
|-------------|-----------------|
| `ARG` | Name, default value |
| `ENV` | Name, default value (handles both `KEY=VALUE` and `KEY VALUE` syntax) |
| `LABEL` | Name, value |
| `EXPOSE` | Port number/protocol |

### Magic Comments

Comments immediately preceding an instruction can provide metadata:

```dockerfile
# @name: DB_HOST
# @description: The hostname of the database
# @default: localhost
# @required: true
ENV DB_HOST="localhost"
```

- `@name` overrides the variable name.
- `@description` sets human-readable description.
- `@default` overrides the inferred default value.
- `@required` marks the variable as required (`true`/`false`).
- If `@default` is omitted, the value is inferred from the instruction.
- If `@name` is omitted, the variable name from the instruction is used.
- Multiple magic comment blocks before a multi-value instruction are mapped 1:1 to the items.

## 8. Dynamic Analysis (Runners)

External tools are invoked via `os/exec`. Each runner implements:

```go
type ToolRunner interface {
    Name() string
    IsAvailable() bool
    Run(image string, verbose bool) (*types.ImageStats, error)
}
```

If a tool is not installed, it is silently skipped (warning logged in verbose mode). Results from all runners are merged into a single `ImageStats` struct.

### Runner: RuntimeRunner

- **Binary:** `docker` or `podman` (auto-detected)
- **Command:** `<binary> inspect <image>`
- **Extracts:** Architecture, OS, Size (bytes -> MB), TotalLayers (from RootFS.Layers)

### Runner: ManifestRunner

- **Binary:** `docker` or `podman`
- **Command:** `<binary> manifest inspect <image>`
- **Extracts:** SupportedArchitectures (from manifest list platforms, e.g., `linux/amd64`)
- Sets `DOCKER_CLI_EXPERIMENTAL=enabled` for older Docker versions.

### Runner: SyftRunner

- **Binary:** `syft` (resolved via `installer.FindTool()`)
- **Command:** `syft <image> -o json`
- **Extracts:** TotalPackages (unique count), Packages (deduplicated, sorted alphabetically), OSDistro (from distro field)

### Runner: GrypeRunner

- **Binary:** `grype` (resolved via `installer.FindTool()`)
- **Command:** `grype <image> -o json`
- **Extracts:** VulnSummary (severity -> count map), Vulnerabilities (sorted by severity desc, then ID asc), VulnScanTime (from descriptor timestamp)

### Runner: DiveRunner

- **Binary:** `dive` (resolved via `installer.FindTool()`)
- **Command:** `dive <image> --json <tmpfile>`
- **Extracts:** Efficiency (efficiencyScore * 100), WastedBytes (inefficientBytes -> MB)
- **Podman support:** When Docker is absent but Podman is available, auto-detects the Podman socket via `podman machine inspect` and sets `DOCKER_HOST` for Dive.

### Tool Resolution

`installer.FindTool(name)` checks:
1. System `PATH` via `exec.LookPath`
2. Fallback: `~/.dock-docs/bin/<name>`

### Image Pre-Pull

Before analysis, `runner.EnsureImage()` checks if the image exists locally (`<binary> inspect --type=image`). If not found, it pulls the image. Container runtime auto-detection: `docker` preferred, `podman` as fallback.

### Parallelism

- **Single image:** All available runners execute in parallel goroutines with `sync.WaitGroup`. Results are merged under a `sync.Mutex`. Individual runner failures are logged as warnings; partial results are returned.
- **Comparison:** `AnalyzeComparison()` uses `golang.org/x/sync/errgroup` to analyze all images in the comparison list concurrently. Individual image failures are non-fatal.

## 9. Template System

### Built-in Templates

12 templates total: 6 names x 2 types (image + comparison).

| Name | Format | Description |
|------|--------|-------------|
| `default` | markdown | Standard format with badges and security analysis |
| `minimal` | markdown | ENV/ARG only, no badges or analysis |
| `detailed` | markdown | Expanded format with full package lists and metadata table |
| `compact` | markdown | Single-line condensed format for quick reference |
| `html` | html | Interactive HTML dashboard with styled tables |
| `json` | json | Machine-readable JSON output for CI/CD integration |

### Template Type Constants

```go
type TemplateType string
const (
    TemplateTypeImage      TemplateType = "image"
    TemplateTypeComparison TemplateType = "comparison"
)
```

### Template Map Keys

Templates are registered as `"name:type"` keys mapping to embedded file paths:

```
"default:image"       -> builtin/markdown/default.tmpl
"default:comparison"  -> builtin/markdown/default_comparison.tmpl
"minimal:image"       -> builtin/markdown/minimal.tmpl
"minimal:comparison"  -> builtin/markdown/minimal_comparison.tmpl
...
"html:image"          -> builtin/html/dashboard.tmpl
"html:comparison"     -> builtin/html/dashboard_comparison.tmpl
"json:image"          -> builtin/json/summary.tmpl
"json:comparison"     -> builtin/json/summary_comparison.tmpl
```

### Template Embedding

Templates are embedded at compile time via `//go:embed`:

```go
//go:embed builtin/markdown/*.tmpl builtin/html/*.tmpl builtin/json/*.tmpl
var builtinFS embed.FS
```

### Template Loading & Caching

The `Loader` struct handles loading and caches parsed templates in a `sync.RWMutex`-protected map:

- `LoadBuiltin(name, type)` — loads from embedded FS.
- `LoadFile(path)` — loads a custom template from disk.
- `Validate(path)` — parses a custom template to check for syntax errors.

### Custom Templates

Custom templates are specified by file path (via `--template` flag or `template.path` in YAML). Security controls:

- **Max file size:** 1 MB
- **Directory traversal prevention:** Rejects paths containing `..`
- Custom file templates default to `"markdown"` output format.

### Template Functions

Available in all templates via `GetFuncMap()`:

| Function | Signature | Description |
|----------|-----------|-------------|
| `join` | `([]string, string) string` | `strings.Join` |
| `upper` | `(string) string` | Uppercase |
| `lower` | `(string) string` | Lowercase |
| `trim` | `(string) string` | Trim whitespace |
| `contains` | `(string, string) bool` | Substring check |
| `hasPrefix` | `(string, string) bool` | Prefix check |
| `hasSuffix` | `(string, string) bool` | Suffix check |
| `replace` | `(string, string, string) string` | Replace all occurrences |
| `default` | `(fallback, val string) string` | Returns fallback if val is empty |
| `jsonEscape` | `(string) string` | Escapes for JSON string values |
| `index` | `(map[string]int, string) int` | Nil-safe map access (returns 0 for missing keys) |
| `emoji` | `(string) string` | Returns emoji or text alternative based on `--nomoji` |

### Emoji System

The `emoji` function and the `.Emoji(name)` method on context objects return emoji characters by default, or plain-text alternatives when `--nomoji` is set:

| Name | Emoji | NoMoji |
|------|-------|--------|
| `check` | `✅` | `[YES]` |
| `cross` | `❌` | `[NO]` |
| `critical` | `🛑` | `[CRIT]` |
| `high` | `🟠` | `[HIGH]` |
| `medium` | `🟡` | `[MED]` |
| `low` | `🔵` | `[LOW]` |
| `clean` | `🟢` | `[OK]` |
| `whale` | `🐳` | (empty) |
| `gear` | `⚙️` | (empty) |
| `shield` | `🛡️` | (empty) |
| `tag` | `🏷️` | (empty) |
| `search` | `🔍` | (empty) |
| `down` | `👇` | (empty) |
| `package` | `📦` | (empty) |

### Template Execution Security

Templates execute within `ExecuteWithLimits()`:

- **Max execution time:** 30 seconds
- **Max output size:** 10 MB
- Uses `context.WithTimeout` and a goroutine to enforce the time limit.
- Uses a `limitedWriter` that returns an error if output exceeds the byte limit.

### Output Formats

| Format | Behavior |
|--------|----------|
| `markdown` | Injected into existing file between markers |
| `html` | Written as standalone file (direct-write) |
| `json` | Written as standalone file (direct-write) |

For direct-write formats:
- In CLI Mode: output path is derived from `--output` by replacing the `.md` extension (e.g., `README.html`, `README.json`). If `--output` is not the default, it is used as-is.
- In YAML Mode: output path is derived from the section marker name (e.g., `README-html-image.html`).

## 10. Injection System

Content is inserted into existing files using HTML comment markers.

### Default Markers

```html
<!-- BEGIN: dock-docs -->
...generated content...
<!-- END: dock-docs -->
```

### Named Markers

Used in YAML Mode to support multiple sections in one file:

```html
<!-- BEGIN: dock-docs:main-docs -->
...generated content...
<!-- END: dock-docs:main-docs -->
```

### Behavior

- Content between markers is replaced entirely; markers themselves are preserved.
- Both start and end markers must be present; otherwise an error is returned.
- End marker must appear after start marker.
- In CLI Mode, default (unnamed) markers are used.

## 11. Installer

The `dock-docs setup` command manages external tool dependencies.

### Tools

| Tool | Repository | Purpose |
|------|-----------|---------|
| `syft` | `anchore/syft` | SBOM / package listing |
| `grype` | `anchore/grype` | Vulnerability scanning |
| `dive` | `wagoodman/dive` | Layer efficiency analysis |

### Install Process

1. Queries GitHub API for latest release tag (`GET /repos/:owner/:repo/releases/latest`).
2. Constructs platform-specific asset name: `<tool>_<version>_<GOOS>_<GOARCH>.tar.gz`.
3. Downloads the tarball and extracts the binary by base name.
4. Writes to `<installDir>/<tool>` with mode `0755`.
5. Runs `<tool> version` as verification (non-fatal if it fails).

### Authentication

Uses `GITHUB_TOKEN` environment variable if set (for both API requests and asset downloads) to avoid rate limiting.

### Default Install Directory

`~/.dock-docs/bin/`

## 12. Badge Generation

`ImageStats` includes badge URL helper methods that generate Shields.io static badge URLs:

| Method | Label | Color Logic |
|--------|-------|-------------|
| `SizeBadge(baseURL)` | `Size` | Always `blue` |
| `LayersBadge(baseURL)` | `Layers` | Always `blue` |
| `EfficiencyBadge(baseURL)` | `Efficiency` | `green` (>=90%), `yellow` (80-90%), `red` (<80%) |
| `VulnBadge(baseURL)` | `Security` | `green` (no critical/high), `orange` (high only), `red` (critical) |

URL format: `<baseURL>?label=<LABEL>&message=<VALUE>&color=<COLOR>`

The base URL defaults to `https://img.shields.io/static/v1` and is configurable via `--badge-base-url` (CLI Mode) or `badgeBaseURL` (YAML Mode).

## 13. Build & Distribution

### Binary Build

```bash
go build -o bin/dock-docs .
```

### GoReleaser

Configuration: `.goreleaser.yaml` (GoReleaser v2 format with `version: 2`)

- **Platforms:** Linux (amd64, arm64), Darwin (amd64, arm64), Windows (amd64)
- **Archives:** `.tar.gz` with README.md and LICENSE
- **Docker images:** Pushed to `ghcr.io/northcutted/dock-docs` (latest + version tag)
- **Changelog:** Auto-generated, grouped by commit prefix

### Dockerfile

Multi-stage build:

1. **Builder stage:** `golang:1.25-alpine` — builds the Go binary.
2. **Tools stage:** Downloads `syft`, `grype`, `dive`, and Docker CLI into a staging area.
3. **Final stage:** `gcr.io/distroless/static-debian12` — copies binary and tools. Entry point: `/dock-docs`.

### GitHub Action

File: `action.yml` — Composite action for use in CI workflows.

**Inputs:** `version`, `config`, `file`, `output`, `image`, `template`, `dry_run`, `nomoji`, `ignore_errors`, `verbose`, `badge_base_url`

**Steps:**
1. Downloads `dock-docs` binary from GitHub Releases (respects `version` input).
2. Installs external tools via `dock-docs setup`.
3. Runs `dock-docs` with inputs mapped to environment variables.

All inputs are passed via environment variables (not interpolated into shell commands) to prevent script injection.

## 14. CI/CD Workflows

### `test.yml`

Triggers on PR and push to `main`. Steps:
- Checkout, setup Go 1.25
- `go test -race -v ./...`
- `golangci-lint run`
- Coverage reporting

### `release.yml`

Triggers on tags matching `v*`. Steps:
- Checkout, setup Go 1.25
- Login to GHCR
- Run GoReleaser

### `dogfood.yml`

Triggers after a release is published. Steps:
- Checkout, setup Go 1.25
- Build `dock-docs` from source
- Run `dock-docs setup`
- Run `dock-docs` against each sample project (`samples/go/`, `samples/node/`, `samples/python/`)
- Commits and pushes regenerated sample docs

## 15. Testing

### Test Strategy

- **77 tests** across 10 test files, all passing with `-race`.
- Standard `testing` package (no external assertion libraries).
- **Table-driven tests** for multi-scenario coverage.
- **Interface injection** for runner mocking (pass mock `Runner` implementations to `AnalyzeImage`).
- **Function variable swapping** for standalone side-effect functions (e.g., `var ensureImage = runner.EnsureImage` in `analysis/analyzer.go`).

### Test Files

| File | Coverage Area |
|------|--------------|
| `cmd/root_test.go` | Template resolution, output path derivation, template selection |
| `pkg/analysis/analyzer_test.go` | Image analysis orchestration, runner merging, error handling |
| `pkg/config/config_test.go` | YAML parsing, defaults, section resolution |
| `pkg/injector/injector_test.go` | Marker injection, edge cases |
| `pkg/installer/installer_test.go` | Tool finding, install process (with HTTP mock) |
| `pkg/parser/parser_test.go` | Dockerfile parsing, magic comments, multi-value instructions |
| `pkg/renderer/markdown_test.go` | Template rendering, context building |
| `pkg/runner/runner_test.go` | Runner availability, JSON parsing |
| `pkg/templates/loader_test.go` | Template loading, caching, validation, security |
| `pkg/types/stats_test.go` | Badge URL generation, TotalVulns |

### Running Tests

```bash
# All tests
go test -v ./...

# With race detection
go test -race -v ./...

# Specific package
go test -v ./pkg/analysis

# Specific test
go test -v ./pkg/analysis -run TestAnalyzeImage
```

### Linting

```bash
golangci-lint run
```

Zero issues expected.

## 16. Dependencies

### Direct

| Module | Version | Purpose |
|--------|---------|---------|
| `github.com/moby/buildkit` | `v0.27.1` | Dockerfile AST parser |
| `github.com/spf13/cobra` | `v1.10.2` | CLI framework |
| `golang.org/x/sync` | `v0.19.0` | `errgroup` for parallel comparison analysis |
| `gopkg.in/yaml.v3` | `v3.0.1` | YAML config parsing |

### External Tool Dependencies (Runtime)

| Tool | Purpose | Required For |
|------|---------|-------------|
| `docker` or `podman` | Container runtime | Image inspect, pull, manifest inspect |
| `syft` | SBOM generation | Package listing, distro detection |
| `grype` | Vulnerability scanning | Vulnerability summary and details |
| `dive` | Image efficiency | Efficiency score, wasted bytes |

None of the external tools are required for basic Dockerfile parsing. They are only needed for dynamic image analysis. Unit tests do not require these tools.

## 17. Error Handling

- Functions that can fail return `error` as the last return value.
- Errors are wrapped with context: `fmt.Errorf("description: %w", err)`.
- Runner failures are non-fatal by default: logged as warnings, partial results returned.
- `--ignore-errors` allows continuing even when image analysis fails entirely.
- Missing external tools are silently skipped (warning in `--verbose` mode).
- Comparison analysis uses partial success: individual image failures don't block the entire comparison.
