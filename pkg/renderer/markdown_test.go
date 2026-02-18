package renderer

import (
	"strings"
	"testing"
	"time"

	"github.com/northcutted/dock-docs/pkg/parser"
	"github.com/northcutted/dock-docs/pkg/templates"
	"github.com/northcutted/dock-docs/pkg/types"
)

func TestRender(t *testing.T) {
	doc := &parser.Documentation{
		Items: []parser.DocItem{
			{
				Name:        "PORT",
				Type:        "ENV",
				Description: "Port to listen on",
				Value:       "8080",
				Required:    true,
			},
		},
	}

	// Test Case 1: Without Stats
	output, err := Render(doc, nil, RenderOptions{})
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	if !strings.Contains(output, "### Environment Variables") {
		t.Errorf("expected output to contain Environment Variables section")
	}
	if strings.Contains(output, "🛡️ Security & Efficiency") {
		t.Error("expected output NOT to contain Security section")
	}

	// Test Case 2: With Stats
	scanTime, _ := time.Parse(time.RFC3339, "2024-02-15T14:30:00Z")
	stats := &types.ImageStats{
		ImageTag:     "test:latest",
		SizeBytes:    52428800, // 50 MB
		Architecture: "amd64",
		OS:           "linux",
		Efficiency:   95.5,
		WastedBytes:  "2 MB",
		TotalLayers:  10,
		VulnSummary:  map[string]int{"Critical": 1, "High": 2},
		VulnScanTime: scanTime,
		Vulnerabilities: []types.Vulnerability{
			{ID: "CVE-2023-1234", Severity: "Critical", Package: "openssl", Version: "1.1.1"},
			{ID: "CVE-2023-5678", Severity: "High", Package: "curl", Version: "7.68"},
		},
		TotalPackages: 5,
		Packages: []types.PackageSummary{
			{Name: "python", Version: "3.9"},
		},
	}

	opts := RenderOptions{
		BadgeBaseURL: "https://img.shields.io/static/v1",
	}

	output, err = Render(doc, stats, opts)
	if err != nil {
		t.Fatalf("Render(stats) error = %v", err)
	}

	if !strings.Contains(output, "# 🐳 Docker Image Analysis: test:latest") {
		t.Error("expected output to contain Docker Image Analysis header")
	}
	// Check for badges
	if !strings.Contains(output, "img.shields.io/static/v1?label=Size") {
		t.Error("expected output to contain Size badge")
	}
	// Check for sections
	if !strings.Contains(output, "### Environment Variables") {
		t.Error("expected output to contain Environment Variables section")
	}
	if !strings.Contains(output, "## 🛡️ Security & Efficiency") {
		t.Error("expected output to contain Security section")
	}
	// Check for vulnerability
	if !strings.Contains(output, "[CVE-2023-1234](https://nvd.nist.gov/vuln/detail/CVE-2023-1234)") {
		t.Error("expected output to contain CVE link")
	}
	if !strings.Contains(output, "| python | 3.9 |") {
		t.Error("expected output to contain python package")
	}
}

func TestRender_NoMoji(t *testing.T) {
	doc := &parser.Documentation{
		Items: []parser.DocItem{
			{Name: "PORT", Type: "ENV", Value: "8080", Required: true},
		},
	}

	stats := &types.ImageStats{
		ImageTag:    "test:latest",
		VulnSummary: map[string]int{"Critical": 1},
	}

	opts := RenderOptions{
		NoMoji:       true,
		BadgeBaseURL: "https://img.shields.io/static/v1",
	}

	output, err := Render(doc, stats, opts)
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	// Should not contain emoji
	if strings.Contains(output, "🐳") {
		t.Error("expected output NOT to contain whale emoji with NoMoji=true")
	}
	if strings.Contains(output, "🛡️") {
		t.Error("expected output NOT to contain shield emoji with NoMoji=true")
	}

	// Should contain text alternatives
	if !strings.Contains(output, "[YES]") {
		t.Error("expected output to contain [YES] for required field")
	}
}

func TestGetEmoji(t *testing.T) {
	tests := []struct {
		name      string
		emojiName string
		noMoji    bool
		expected  string
	}{
		{"whale emoji", "whale", false, "🐳 "},
		{"whale text", "whale", true, ""},
		{"check emoji", "check", false, "✅"},
		{"check text", "check", true, "[YES]"},
		{"cross emoji", "cross", false, "❌"},
		{"cross text", "cross", true, "[NO]"},
		{"critical emoji", "critical", false, "🛑"},
		{"critical text", "critical", true, "[CRIT]"},
		{"high emoji", "high", false, "🟠"},
		{"high text", "high", true, "[HIGH]"},
		{"medium emoji", "medium", false, "🟡"},
		{"medium text", "medium", true, "[MED]"},
		{"low emoji", "low", false, "🔵"},
		{"low text", "low", true, "[LOW]"},
		{"clean emoji", "clean", false, "🟢"},
		{"clean text", "clean", true, "[OK]"},
		{"unknown emoji", "unknown", false, ""},
		{"unknown text", "unknown", true, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := templates.GetEmoji(tt.emojiName, tt.noMoji)
			if result != tt.expected {
				t.Errorf("GetEmoji(%s, %v) = %q, want %q", tt.emojiName, tt.noMoji, result, tt.expected)
			}
		})
	}
}

func TestTemplateSelection_Format(t *testing.T) {
	tests := []struct {
		name     string
		sel      TemplateSelection
		expected string
	}{
		{"default name", TemplateSelection{Name: "default"}, "markdown"},
		{"empty name defaults to markdown", TemplateSelection{}, "markdown"},
		{"html name", TemplateSelection{Name: "html"}, "html"},
		{"json name", TemplateSelection{Name: "json"}, "json"},
		{"minimal name", TemplateSelection{Name: "minimal"}, "markdown"},
		{"compact name", TemplateSelection{Name: "compact"}, "markdown"},
		{"detailed name", TemplateSelection{Name: "detailed"}, "markdown"},
		{"custom path defaults to markdown", TemplateSelection{Path: "/some/file.tmpl"}, "markdown"},
		{"path overrides name", TemplateSelection{Name: "html", Path: "/some/file.tmpl"}, "markdown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.sel.Format()
			if result != tt.expected {
				t.Errorf("Format() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestRender_AllSections(t *testing.T) {
	doc := &parser.Documentation{
		Items: []parser.DocItem{
			{Name: "PORT", Type: "ENV", Value: "8080", Description: "Server port"},
			{Name: "VERSION", Type: "ARG", Value: "1.0", Description: "Build version"},
			{Name: "8080", Type: "EXPOSE", Value: "8080", Description: "HTTP port"},
			{Name: "maintainer", Type: "LABEL", Value: "test@example.com"},
		},
	}

	output, err := Render(doc, nil, RenderOptions{})
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	// Verify all sections are present
	if !strings.Contains(output, "### Environment Variables") {
		t.Error("missing Environment Variables section")
	}
	if !strings.Contains(output, "### Build Arguments") {
		t.Error("missing Build Arguments section")
	}
	if !strings.Contains(output, "### Exposed Ports") {
		t.Error("missing Exposed Ports section")
	}
	if !strings.Contains(output, "### Labels") {
		t.Error("missing Labels section")
	}
}

func TestRender_WithOSDistro(t *testing.T) {
	doc := &parser.Documentation{
		Items: []parser.DocItem{},
	}

	stats := &types.ImageStats{
		ImageTag:     "test:latest",
		OS:           "linux",
		Architecture: "amd64",
		OSDistro:     "alpine",
		VulnSummary:  map[string]int{},
	}

	output, err := Render(doc, stats, RenderOptions{})
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	// Verify OS distro is shown
	if !strings.Contains(output, "alpine (linux/amd64)") {
		t.Error("expected OS distro to be displayed")
	}
}

func TestRender_EmptyDocument(t *testing.T) {
	doc := &parser.Documentation{
		Items: []parser.DocItem{},
	}

	output, err := Render(doc, nil, RenderOptions{})
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	// Should still generate output with header
	if !strings.Contains(output, "Docker Image Analysis") {
		t.Error("expected header even for empty document")
	}

	// Should not contain section headers for empty sections
	if strings.Contains(output, "### Environment Variables") {
		t.Error("should not contain Environment Variables section when empty")
	}
}
