# Project Specification: Docker-Docs Generator

**Role:** You are a Senior Go Systems Engineer.
**Objective:** Build a CLI tool called `docker-docs` (conceptually similar to `terraform-docs`) that parses a `Dockerfile`, extracts metadata from instructions and "magic comments," and generates a Markdown documentation table.

## 1. Core Tech Stack
* **Language:** Go (Golang) 1.25+
* **Parser:** `github.com/moby/buildkit/frontend/dockerfile/parser` (Do NOT use Regex; use the official AST parser).
* **CLI Framework:** `github.com/spf13/cobra`
* **Templating:** `text/template` (Standard library)

## 2. Functional Requirements

### A. Parsing Logic (The AST Walker)
The tool must read a `Dockerfile` and extract the following instructions into a structured object:
1.  **ARG:** Build-time variables.
2.  **ENV:** Runtime environment variables.
3.  **LABEL:** OCI metadata.
4.  **EXPOSE:** Network ports.

**Critical Constraint:** The parser must handle standard Docker syntax idiosyncrasies:
* `ENV KEY=VALUE` (Equality) vs `ENV KEY VALUE` (Space-separated).
* Multi-line instructions using backslashes (`\`).

### B. "Magic Comment" Extraction
The tool must parse comments *immediately preceding* an instruction to extract metadata.
**Syntax Spec:**

```dockerfile
# @name: DB_HOST
# @description: The hostname of the database.
# @default: localhost
# @required: true
ENV DB_HOST="localhost"
```

* If `@default` is not provided in comments, infer it from the Dockerfile instruction value.
* If `@name` is not provided, use the variable name.

### C. Injection Strategy (Idempotency)
The tool must support updating an existing `README.md` without overwriting the whole file.
* **Markers:** Look for `` and ``.
* **Logic:** Replace only the content between these markers.
* **Fallback:** If no markers are found, print to `stdout`.

## 3. Data Structures
Use this internal schema for the collected data:

```go
type DocItem struct {
    Name        string // e.g., "PORT"
    Value       string // inferred default from the instruction
    Description string // from @description
    Type        string // "ARG", "ENV", "LABEL", "EXPOSE"
    Required    bool   // from @required
}

type Documentation struct {
    Items []DocItem
}
```

## 4. Implementation Plan (Step-by-Step)

**Phase 1: The Parser**
1.  Create `pkg/parser/parser.go`.
2.  Implement `Parse(filename string) (*Documentation, error)`.
3.  Use `parser.Parse()` from Moby BuildKit to get the AST.
4.  Iterate through `Node.Children`.
5.  Implement a helper `extractComments(node *Node)` that reads `node.PrevComment` and scans for `@tag` keys.

**Phase 2: The Renderer**
1.  Create `pkg/renderer/markdown.go`.
2.  Define a default Markdown template containing a table with columns: `| Name | Type | Description | Default | Required |`.
3.  Implement `Render(doc *Documentation) (string, error)`.

**Phase 3: The CLI**
1.  Create `cmd/root.go` using Cobra.
2.  Add flags:
    * `--file, -f`: Path to Dockerfile (default "./Dockerfile").
    * `--output, -o`: Path to output file (default "README.md").
    * `--dry-run`: Print to stdout instead of writing to file.

## 5. Execution
Please start by initializing the Go module and creating the **Phase 1 (Parser)** implementation. Write a test case with a sample Dockerfile containing both `ENV` and `ARG` with magic comments to verify the parsing logic.

## 6. Phase 4: Dynamic Analysis (External Tools)

**Objective:** Extend the tool to inspect *built* container images using external CLI tools.
**Architecture:** Hybrid Analysis. The tool will now shell out to `docker`, `syft`, `grype`, and `dive` to capture runtime data and merge it with the static Dockerfile analysis.

### A. Core Architecture: The "Runner" Pattern
**Constraint:** Do NOT import `syft` or `dive` as Go libraries. This will cause dependency hell.
Instead, use `os/exec` to run them as sub-processes and parse their JSON output.

1.  **New Interface:** Create `pkg/runner/runner.go`.
    ```go
    type ToolRunner interface {
        Name() string
        IsAvailable() bool // e.g. exec.LookPath("syft")
        Run(image string) (interface{}, error)
    }
    ```
2.  **New Data Structs:** Create `pkg/analysis/stats.go`.
    ```go
    type ImageStats struct {
        Architecture string
        OS           string
        SizeMB       string
        TotalLayers  int
        Efficiency   float64 // from Dive (0-100)
        WastedBytes  string  // from Dive
        Packages     []PackageSummary // from Syft
        VulnSummary  map[string]int   // from Grype (Severity -> Count)
    }
    ```

### B. Tool Integrations

#### 1. Docker Inspect (Baseline)
* **Command:** `docker inspect <image>`
* **Data to Extract:**
    * `Architecture` & `Os`
    * `Size` (Convert bytes to human-readable MB/GB)
    * `RootFS.Layers` (Count the length of the array)

#### 2. Syft (SBOM / Packages)
* **Command:** `syft <image> -o json`
* **Data to Extract:**
    * Parse the `artifacts` array.
    * **Logic:** Do NOT list every single package (it's too long).
    * **Action:** Extract the count of total packages.
    * **Action:** Extract a list of "Key Frameworks" if detected (e.g., `python`, `node`, `go`, `glibc`, `openssl`).

#### 3. Grype (Vulnerabilities)
* **Command:** `grype <image> -o json`
* **Data to Extract:**
    * Parse `matches`.
    * Group by `vulnerability.severity` (Critical, High, Medium, Low).
    * Return a simple map: `{"Critical": 2, "High": 5}`.

#### 4. Dive (Image Efficiency)
* **Command:** `dive <image> --json output.json`
    * *Note: Dive writes to a file, not stdout.*
* **Data to Extract:**
    * `image.inefficientBytes`
    * `image.efficiencyScore`

### C. CLI Updates
* Add flag: `--image <tag>` (e.g., `my-app:latest`).
* **Logic:**
    * If `--image` is present, run the active Runners.
    * If a tool (e.g., `dive`) is missing from the system `PATH`, log a warning but **do not fail**. Just skip that section of the report.

### D. Renderer Updates
Update the Markdown template to include a "Runtime Analysis" section if the data is available.

**Proposed Template Addition:**
```markdown
## Image Analysis ({{ .ImageTag }})

| Metric | Value |
|--------|-------|
| Size | {{ .Stats.SizeMB }} |
| Architecture | {{ .Stats.Architecture }}/{{ .Stats.OS }} |
| Efficiency | {{ .Stats.Efficiency }}% ({{ .Stats.WastedBytes }} wasted) |

**Security Summary:**
Critical: {{ .Stats.VulnSummary.Critical }} | High: {{ .Stats.VulnSummary.High }} | Medium: {{ .Stats.VulnSummary.Medium }}

<details>
<summary>Key Packages ({{ .Stats.TotalPackages }} total)</summary>

| Package | Version |
|---------|---------|
{{ range .Stats.Packages }}| {{ .Name }} | {{ .Version }} |
{{ end }}
</details>

## 7. Phase 5: Visual Polish & Grouping

**Objective:** Refine the Markdown output to be "GitHub-native," scannable, and professional.
**Changes:**
1.  **Split Tables:** Instead of one giant table, group items by type (`ENV`, `ARG`, `LABEL`, `EXPOSE`).
2.  **Badges:** Generate Shields.io badge URLs for high-level stats (Size, Layers, Vulns).
3.  **Visual Hierarchy:** Use emojis and specific formatting to highlight security status (e.g., 🔴 for Critical, 🟢 for Safe).

### A. Data Structure Update
Update `pkg/analysis/stats.go` to include helper methods for dynamic badge generation and filtering.

1.  **Badge Helpers:**
    * `GetSizeBadge() string` -> returns `https://img.shields.io/badge/size-7.6MB-blue`
    * `GetVulnBadge() string` -> returns `https://img.shields.io/badge/vulns-0%20Critical-green` (URL encoded).
    * **Logic:** If Critical Vulns > 0, set Badge Color to `red`. Else `green`.

2.  **Filter Helper:**
    * Add a method to the `Documentation` struct: `FilterByType(type string) []DocItem`.
    * This allows the template to call `{{ range .FilterByType "ENV" }}` to only show environment variables in a specific section.

### B. Updated Markdown Template
Replace the existing template in `pkg/renderer/markdown.go` with this "Dashboard Style" layout.

```markdown
# 🐳 Docker Image Analysis: {{ .ImageTag }}

![Size]({{ .Stats.SizeBadge }}) ![Layers]({{ .Stats.LayersBadge }}) ![Vulns]({{ .Stats.VulnBadge }}) ![Efficiency]({{ .Stats.EfficiencyBadge }})

## ⚙️ Configuration

### Environment Variables
| Name | Description | Default | Required |
|------|-------------|---------|:--------:|
{{- range .FilterByType "ENV" }}
| `{{ .Name }}` | {{ .Description }} | `{{ if .Value }}{{ .Value }}{{ else }}""{{ end }}` | {{ if .Required }}✅{{ else }}❌{{ end }} |
{{- end }}

### Build Arguments
| Name | Description | Default | Required |
|------|-------------|---------|:--------:|
{{- range .FilterByType "ARG" }}
| `{{ .Name }}` | {{ .Description }} | `{{ .Value }}` | {{ if .Required }}✅{{ else }}❌{{ end }} |
{{- end }}

### Exposed Ports
| Port | Description |
|------|-------------|
{{- range .FilterByType "EXPOSE" }}
| `{{ .Name }}` | {{ .Description }} |
{{- end }}

---

## 🛡️ Security & Efficiency

**Base Image:** `{{ .Stats.OS }} ({{ .Stats.Architecture }})`
**Efficiency Score:** {{ .Stats.Efficiency }}%

### Vulnerabilities
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| {{ if gt .Stats.VulnSummary.Critical 0 }}🔴 {{ else }}🟢 {{ end }}{{ .Stats.VulnSummary.Critical }} | {{ if gt .Stats.VulnSummary.High 0 }}🟠 {{ else }}🟢 {{ end }}{{ .Stats.VulnSummary.High }} | {{ .Stats.VulnSummary.Medium }} | {{ .Stats.VulnSummary.Low }} |

<details>
<summary><strong>👇 Expand Vulnerability Details ({{ .Stats.TotalVulns }} found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
{{- range .Stats.Vulns }}
| [{{ .ID }}](https://nvd.nist.gov/vuln/detail/{{ .ID }}) | {{ .Severity }} | `{{ .Package }}` | `{{ .Version }}` |
{{- end }}
</details>

<details>
<summary><strong>📦 Installed Packages ({{ .Stats.TotalPackages }} total)</strong></summary>

| Package | Version |
|---------|---------|
{{- range .Stats.Packages }}
| {{ .Name }} | {{ .Version }} |
{{- end }}
</details>

C. Implementation Details

Shields.io Usage: Use the static badge endpoint: https://img.shields.io/static/v1?label=<LABEL>&message=<VALUE>&color=<COLOR>.

Template Logic: Ensure the template handles empty lists gracefully (e.g., if there are no ARGs, the table header should probably be hidden or the section skipped).