package renderer

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/northcutted/docker-docs/pkg/analysis"
	"github.com/northcutted/docker-docs/pkg/parser"
)

type RenderOptions struct {
	NoMoji bool
}

// ReportContext holds all data passed to the template
type ReportContext struct {
	Doc      *parser.Documentation
	Stats    *analysis.ImageStats
	ImageTag string
	Options  RenderOptions
}

func (r ReportContext) Emoji(name string) string {
	return getEmoji(name, r.Options.NoMoji)
}

func getEmoji(name string, noMoji bool) string {
	if noMoji {
		switch name {
		case "check":
			return "[YES]"
		case "cross":
			return "[NO]"
		case "critical":
			return "[CRIT]"
		case "high":
			return "[HIGH]"
		case "medium":
			return "[MED]"
		case "low":
			return "[LOW]"
		case "clean":
			return "[OK]"
		default:
			return ""
		}
	}

	switch name {
	case "whale":
		return "🐳 "
	case "gear":
		return "⚙️ "
	case "shield":
		return "🛡️ "
	case "tag":
		return "🏷️ "
	case "search":
		return "🔍 "
	case "down":
		return "👇 "
	case "package":
		return "📦 "
	case "check":
		return "✅"
	case "cross":
		return "❌"
	case "critical":
		return "🛑"
	case "high":
		return "🟠"
	case "medium":
		return "🟡"
	case "low":
		return "🔵"
	case "clean":
		return "🟢"
	default:
		return ""
	}
}

const defaultTemplate = `
# {{ .Emoji "whale" }}Docker Image Analysis: {{ .ImageTag }}

{{- if .Stats }}
![Size]({{ .Stats.SizeBadge }}) ![Layers]({{ .Stats.LayersBadge }}) ![Vulns]({{ .Stats.VulnBadge }}) ![Efficiency]({{ .Stats.EfficiencyBadge }})
{{- end }}

## {{ .Emoji "gear" }}Configuration

{{- if (len (.Doc.FilterByType "ENV")) }}
### Environment Variables
| Name | Description | Default | Required |
|------|-------------|---------|:--------:|
{{- range (.Doc.FilterByType "ENV") }}
| ` + "`{{ .Name }}`" + ` | {{ .Description }} | ` + "`{{ if .Value }}{{ .Value }}{{ else }}\"\"{{ end }}`" + ` | {{ if .Required }}{{ $.Emoji "check" }}{{ else }}{{ $.Emoji "cross" }}{{ end }} |
{{- end }}
{{- end }}

{{- if (len (.Doc.FilterByType "ARG")) }}
### Build Arguments
| Name | Description | Default | Required |
|------|-------------|---------|:--------:|
{{- range (.Doc.FilterByType "ARG") }}
| ` + "`{{ .Name }}`" + ` | {{ .Description }} | ` + "`{{ .Value }}`" + ` | {{ if .Required }}{{ $.Emoji "check" }}{{ else }}{{ $.Emoji "cross" }}{{ end }} |
{{- end }}
{{- end }}

{{- if (len (.Doc.FilterByType "EXPOSE")) }}
### Exposed Ports
| Port | Description |
|------|-------------|
{{- range (.Doc.FilterByType "EXPOSE") }}
| ` + "`{{ .Name }}`" + ` | {{ .Description }} |
{{- end }}
{{- end }}

{{- if (len (.Doc.FilterByType "LABEL")) }}
### Labels
| Key | Value |
|-----|-------|
{{- range (.Doc.FilterByType "LABEL") }}
| ` + "`{{ .Name }}`" + ` | {{ .Value }} |
{{- end }}
{{- end }}

{{- if .Stats }}
---

## {{ .Emoji "shield" }}Security & Efficiency

**Base Image:** ` + "`{{ if .Stats.OSDistro }}{{ .Stats.OSDistro }} ({{ .Stats.OS }}/{{ .Stats.Architecture }}){{ else }}{{ .Stats.OS }} ({{ .Stats.Architecture }}){{ end }}`" + `
{{- if .Stats.SupportedArchitectures }}
**Supported Architectures:** ` + "`{{ join .Stats.SupportedArchitectures \", \" }}`" + `
{{- end }}
**Efficiency Score:** {{ printf "%.1f" .Stats.Efficiency }}%

### Vulnerabilities
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| {{ if gt (index .Stats.VulnSummary "Critical") 0 }}{{ .Emoji "critical" }} {{ else }}{{ .Emoji "clean" }} {{ end }}{{ index .Stats.VulnSummary "Critical" }} | {{ if gt (index .Stats.VulnSummary "High") 0 }}{{ .Emoji "high" }} {{ else }}{{ .Emoji "clean" }} {{ end }}{{ index .Stats.VulnSummary "High" }} | {{ if gt (index .Stats.VulnSummary "Medium") 0 }}{{ .Emoji "medium" }} {{ else }}{{ .Emoji "clean" }} {{ end }}{{ index .Stats.VulnSummary "Medium" }} | {{ if gt (index .Stats.VulnSummary "Low") 0 }}{{ .Emoji "low" }} {{ else }}{{ .Emoji "clean" }} {{ end }}{{ index .Stats.VulnSummary "Low" }} |

<details>
<summary><strong>{{ .Emoji "down" }}Expand Vulnerability Details ({{ .Stats.TotalVulns }} found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
{{- range .Stats.Vulnerabilities }}
| [{{ .ID }}](https://nvd.nist.gov/vuln/detail/{{ .ID }}) | {{ .Severity }} | ` + "`{{ .Package }}`" + ` | ` + "`{{ .Version }}`" + ` |
{{- end }}
</details>

<details>
<summary><strong>{{ .Emoji "package" }}Installed Packages ({{ .Stats.TotalPackages }} total)</strong></summary>

| Package | Version |
|---------|---------|
{{- range .Stats.Packages }}
| {{ .Name }} | {{ .Version }} |
{{- end }}
</details>
{{- end }}
`

// Render generates the Markdown table from documentation items.
func Render(doc *parser.Documentation, stats *analysis.ImageStats, opts RenderOptions) (string, error) {
	tmpl, err := template.New("docker-docs").Funcs(template.FuncMap{
		"index": func(m map[string]int, k string) int {
			if v, ok := m[k]; ok {
				return v
			}
			return 0
		},
		"join": strings.Join,
	}).Parse(defaultTemplate)

	if err != nil {
		return "", err
	}

	ctx := ReportContext{
		Doc:     doc,
		Stats:   stats,
		Options: opts,
	}
	if stats != nil {
		ctx.ImageTag = stats.ImageTag
	} else {
		// Fallback if no stats provided
		if ctx.ImageTag == "" {
			ctx.ImageTag = "Dockerfile"
		}
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, ctx); err != nil {
		return "", err
	}

	return buf.String(), nil
}

type MatrixContext struct {
	Matrix  []*analysis.ImageStats
	Options RenderOptions
}

func (r MatrixContext) Emoji(name string) string {
	return getEmoji(name, r.Options.NoMoji)
}

const matrixTemplate = `
### {{ .Emoji "tag" }}Supported Tags

| Tag | Size | Vulns | Efficiency | Architectures |
|-----|------|-------|------------|---------------|
{{- range .Matrix }}
| ` + "`{{ .ImageTag }}`" + ` | ![Size]({{ .SizeBadge }}) | ![Vulns]({{ .VulnBadge }}) | {{ printf "%.1f" .Efficiency }}% | {{ if .SupportedArchitectures }}` + "`{{ join .SupportedArchitectures \", \" }}`" + `{{ else }}` + "`{{ .OS }}/{{ .Architecture }}`" + `{{ end }} |
{{- end }}

{{- range .Matrix }}

<details>
<summary><strong>{{ $.Emoji "search" }}Full Report: {{ .ImageTag }}</strong></summary>

## {{ $.Emoji "shield" }}Security & Efficiency

**Base Image:** ` + "`{{ if .OSDistro }}{{ .OSDistro }} ({{ .OS }}/{{ .Architecture }}){{ else }}{{ .OS }} ({{ .Architecture }}){{ end }}`" + `
{{- if .SupportedArchitectures }}
**Supported Architectures:** ` + "`{{ join .SupportedArchitectures \", \" }}`" + `
{{- end }}
**Efficiency Score:** {{ printf "%.1f" .Efficiency }}%

### Vulnerabilities
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| {{ if gt (index .VulnSummary "Critical") 0 }}{{ $.Emoji "critical" }} {{ else }}{{ $.Emoji "clean" }} {{ end }}{{ index .VulnSummary "Critical" }} | {{ if gt (index .VulnSummary "High") 0 }}{{ $.Emoji "high" }} {{ else }}{{ $.Emoji "clean" }} {{ end }}{{ index .VulnSummary "High" }} | {{ if gt (index .VulnSummary "Medium") 0 }}{{ $.Emoji "medium" }} {{ else }}{{ $.Emoji "clean" }} {{ end }}{{ index .VulnSummary "Medium" }} | {{ if gt (index .VulnSummary "Low") 0 }}{{ $.Emoji "low" }} {{ else }}{{ $.Emoji "clean" }} {{ end }}{{ index .VulnSummary "Low" }} |

<details>
<summary><strong>{{ $.Emoji "down" }}Expand Vulnerability Details ({{ .TotalVulns }} found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
{{- range .Vulnerabilities }}
| [{{ .ID }}](https://nvd.nist.gov/vuln/detail/{{ .ID }}) | {{ .Severity }} | ` + "`{{ .Package }}`" + ` | ` + "`{{ .Version }}`" + ` |
{{- end }}
</details>

<details>
<summary><strong>{{ $.Emoji "package" }}Installed Packages ({{ .TotalPackages }} total)</strong></summary>

| Package | Version |
|---------|---------|
{{- range .Packages }}
| {{ .Name }} | {{ .Version }} |
{{- end }}
</details>

</details>
{{- end }}
`

// RenderMatrix generates the comparison table for multiple images.
func RenderMatrix(stats []*analysis.ImageStats, opts RenderOptions) (string, error) {
	tmpl, err := template.New("docker-docs-matrix").Funcs(template.FuncMap{
		"index": func(m map[string]int, k string) int {
			if v, ok := m[k]; ok {
				return v
			}
			return 0
		},
		"join": strings.Join,
	}).Parse(matrixTemplate)
	if err != nil {
		return "", err
	}

	ctx := MatrixContext{
		Matrix:  stats,
		Options: opts,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, ctx); err != nil {
		return "", err
	}

	return buf.String(), nil
}
