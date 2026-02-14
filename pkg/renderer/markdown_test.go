package renderer

import (
	"strings"
	"testing"

	"docker-docs/pkg/analysis"
	"docker-docs/pkg/parser"
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
	output, err := Render(doc, nil)
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	expectedHeader := "| Name | Type | Description | Default | Required |"
	if !strings.Contains(output, expectedHeader) {
		t.Errorf("expected output to contain header %q", expectedHeader)
	}
	if strings.Contains(output, "Image Analysis") {
		t.Error("expected output NOT to contain Image Analysis")
	}

	// Test Case 2: With Stats
	stats := &analysis.ImageStats{
		ImageTag:     "test:latest",
		SizeMB:       "50 MB",
		Architecture: "amd64",
		OS:           "linux",
		Efficiency:   95.5,
		WastedBytes:  "2 MB",
		TotalLayers:  10,
		VulnSummary:  map[string]int{"Critical": 1, "High": 2},
		Vulnerabilities: []analysis.Vulnerability{
			{ID: "CVE-2023-1234", Severity: "Critical", Package: "openssl", Version: "1.1.1"},
			{ID: "CVE-2023-5678", Severity: "High", Package: "curl", Version: "7.68"},
		},
		TotalPackages: 5,
		Packages: []analysis.PackageSummary{
			{Name: "python", Version: "3.9"},
		},
	}

	output, err = Render(doc, stats)
	if err != nil {
		t.Fatalf("Render(stats) error = %v", err)
	}

	if !strings.Contains(output, "## Image Analysis (test:latest)") {
		t.Error("expected output to contain Image Analysis header")
	}
	if !strings.Contains(output, "| Size | 50 MB |") {
		t.Error("expected output to contain Size")
	}
	if !strings.Contains(output, "Critical: 1") {
		t.Error("expected output to contain Critical count")
	}
	if !strings.Contains(output, "| CVE-2023-1234 | Critical | openssl | 1.1.1 |") {
		t.Error("expected output to contain CVE table row")
	}
	if !strings.Contains(output, "| python | 3.9 |") {
		t.Error("expected output to contain python package")
	}
}
