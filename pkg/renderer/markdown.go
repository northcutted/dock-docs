package renderer

import (
	"bytes"
	"text/template"

	"docker-docs/pkg/analysis"
	"docker-docs/pkg/parser"
)

// ReportContext holds all data passed to the template
type ReportContext struct {
	Items    []parser.DocItem
	Stats    *analysis.ImageStats
	ImageTag string
}

const defaultTemplate = `
| Name | Type | Description | Default | Required |
|------|------|-------------|---------|----------|
{{- range .Items }}
| {{ .Name }} | {{ .Type }} | {{ .Description }} | {{ .Value }} | {{ .Required }} |
{{- end }}

{{- if .Stats }}

## Image Analysis ({{ .ImageTag }})

| Metric | Value |
|--------|-------|
| Size | {{ .Stats.SizeMB }} |
| Architecture | {{ .Stats.Architecture }}/{{ .Stats.OS }} |
| Efficiency | {{ printf "%.1f" .Stats.Efficiency }}% ({{ .Stats.WastedBytes }} wasted) |
| Total Layers | {{ .Stats.TotalLayers }} |

### Security Summary
Critical: {{ index .Stats.VulnSummary "Critical" }} | High: {{ index .Stats.VulnSummary "High" }} | Medium: {{ index .Stats.VulnSummary "Medium" }}

{{- if .Stats.Vulnerabilities }}
<details>
<summary>Vulnerabilities Details ({{ len .Stats.Vulnerabilities }} found)</summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
{{- range .Stats.Vulnerabilities }}
| {{ .ID }} | {{ .Severity }} | {{ .Package }} | {{ .Version }} |
{{- end }}
</details>
{{- end }}

{{- if .Stats.Packages }}
<details>
<summary>Packages ({{ .Stats.TotalPackages }} total)</summary>

| Package | Version |
|---------|---------|
{{- range .Stats.Packages }}
| {{ .Name }} | {{ .Version }} |
{{- end }}
</details>
{{- else }}
*No packages detected.*
{{- end }}
{{- end }}
`

// Render generates the Markdown table from documentation items.
func Render(doc *parser.Documentation, stats *analysis.ImageStats) (string, error) {
	tmpl, err := template.New("docker-docs").Funcs(template.FuncMap{
		"index": func(m map[string]int, k string) int {
			if v, ok := m[k]; ok {
				return v
			}
			return 0
		},
	}).Parse(defaultTemplate)

	if err != nil {
		return "", err
	}

	ctx := ReportContext{
		Items: doc.Items,
		Stats: stats,
	}
	if stats != nil {
		ctx.ImageTag = stats.ImageTag
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, ctx); err != nil {
		return "", err
	}

	return buf.String(), nil
}
