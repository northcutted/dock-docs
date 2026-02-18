// Test file for output path resolution helpers (resolveOutputPath, resolveSectionOutput).
//
// No globals are mutated by these tests — all functions are pure.
package cmd

import (
	"path/filepath"
	"testing"
)

func TestResolveOutputPath(t *testing.T) {
	tests := []struct {
		name          string
		currentOutput string
		format        string
		expected      string
	}{
		{"default markdown", "README.md", "markdown", "README.md"},
		{"default to html", "README.md", "html", "README.html"},
		{"default to json", "README.md", "json", "README.json"},
		{"custom output kept for html", "DOCS.md", "html", "DOCS.md"},
		{"custom output kept for json", "output.txt", "json", "output.txt"},
		{"custom output kept for markdown", "CUSTOM.md", "markdown", "CUSTOM.md"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := resolveOutputPath(tt.currentOutput, tt.format)
			if result != tt.expected {
				t.Errorf("resolveOutputPath(%q, %q) = %q, want %q", tt.currentOutput, tt.format, result, tt.expected)
			}
		})
	}
}

func TestResolveSectionOutput(t *testing.T) {
	tests := []struct {
		name         string
		baseOutput   string
		marker       string
		sectionIndex int
		format       string
		expected     string
	}{
		{"marker html", "README.md", "main", 0, "html", "README-main.html"},
		{"marker json", "README.md", "compare", 1, "json", "README-compare.json"},
		{"no marker uses index", "README.md", "", 2, "html", "README-section2.html"},
		{"subdirectory", "docs/README.md", "main", 0, "json", filepath.Join("docs", "README-main.json")},
		{"custom base name", "DOCS.md", "overview", 0, "html", "DOCS-overview.html"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := resolveSectionOutput(tt.baseOutput, tt.marker, tt.sectionIndex, tt.format)
			if result != tt.expected {
				t.Errorf("resolveSectionOutput(%q, %q, %d, %q) = %q, want %q",
					tt.baseOutput, tt.marker, tt.sectionIndex, tt.format, result, tt.expected)
			}
		})
	}
}
