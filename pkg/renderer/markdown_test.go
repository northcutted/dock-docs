package renderer

import (
	"strings"
	"testing"

	"github.com/northcutted/dock-docs/pkg/types"
	"github.com/northcutted/dock-docs/pkg/parser"
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
	stats := &types.ImageStats{
		ImageTag:     "test:latest",
		SizeMB:       "50 MB",
		Architecture: "amd64",
		OS:           "linux",
		Efficiency:   95.5,
		WastedBytes:  "2 MB",
		TotalLayers:  10,
		VulnSummary:  map[string]int{"Critical": 1, "High": 2},
		Vulnerabilities: []types.Vulnerability{
			{ID: "CVE-2023-1234", Severity: "Critical", Package: "openssl", Version: "1.1.1"},
			{ID: "CVE-2023-5678", Severity: "High", Package: "curl", Version: "7.68"},
		},
		TotalPackages: 5,
		Packages: []types.PackageSummary{
			{Name: "python", Version: "3.9"},
		},
	}

	output, err = Render(doc, stats, RenderOptions{})
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
