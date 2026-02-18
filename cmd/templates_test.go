// Test file for template resolution, listing, export, and validation.
//
// Globals mutated: templateName, stdout (via captureOutput).
// All tests use defer resetFlags()() for cleanup.
package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/northcutted/dock-docs/pkg/config"
	"github.com/northcutted/dock-docs/pkg/renderer"
)

func TestDescribeTemplate(t *testing.T) {
	tests := []struct {
		name     string
		sel      renderer.TemplateSelection
		expected string
	}{
		{"custom path", renderer.TemplateSelection{Path: "/my/template.tmpl"}, "custom file: /my/template.tmpl"},
		{"named builtin", renderer.TemplateSelection{Name: "minimal"}, "built-in: minimal"},
		{"empty defaults", renderer.TemplateSelection{}, "built-in: default"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := describeTemplate(tt.sel)
			if result != tt.expected {
				t.Errorf("describeTemplate() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestResolveTemplateSel(t *testing.T) {
	tests := []struct {
		name         string
		cliTemplate  string // value of templateName global
		cfgTemplate  *config.TemplateConfig
		expectedName string
		expectedPath string
	}{
		{"default empty", "", nil, "", ""},
		{"cli name", "minimal", nil, "minimal", ""},
		{"cli file path with slash", "templates/custom.tmpl", nil, "", "templates/custom.tmpl"},
		{"cli file path with tmpl", "custom.tmpl", nil, "", "custom.tmpl"},
		{"config name", "", &config.TemplateConfig{Name: "detailed"}, "detailed", ""},
		{"config path", "", &config.TemplateConfig{Path: "/path/to/tmpl"}, "", "/path/to/tmpl"},
		{"cli overrides config", "html", &config.TemplateConfig{Name: "minimal"}, "html", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer resetFlags()()
			templateName = tt.cliTemplate

			result := resolveTemplateSel(tt.cfgTemplate)
			if result.Name != tt.expectedName {
				t.Errorf("resolveTemplateSel().Name = %q, want %q", result.Name, tt.expectedName)
			}
			if result.Path != tt.expectedPath {
				t.Errorf("resolveTemplateSel().Path = %q, want %q", result.Path, tt.expectedPath)
			}
		})
	}
}

func TestHandleListTemplates(t *testing.T) {
	output := captureOutput(func() {
		if err := handleListTemplates(); err != nil {
			t.Fatalf("handleListTemplates() error = %v", err)
		}
	})

	if !strings.Contains(output, "Available built-in templates:") {
		t.Error("expected header in output")
	}
	// Should list known templates
	for _, name := range []string{"default", "minimal", "detailed", "compact", "html", "json"} {
		if !strings.Contains(output, name) {
			t.Errorf("expected template %q in output", name)
		}
	}
	if !strings.Contains(output, "dock-docs --template") {
		t.Error("expected usage hint in output")
	}
}

func TestHandleExportTemplate(t *testing.T) {
	// Happy path
	output := captureOutput(func() {
		if err := handleExportTemplate("default"); err != nil {
			t.Fatalf("handleExportTemplate(default) error = %v", err)
		}
	})
	if !strings.Contains(output, "Docker Image Analysis") {
		t.Error("expected template content in output")
	}

	// Error path: unknown template
	err := handleExportTemplate("nonexistent")
	if err == nil {
		t.Fatal("expected error for unknown template")
	}
	if !strings.Contains(err.Error(), "unknown built-in template") {
		t.Errorf("error = %q, want it to contain 'unknown built-in template'", err.Error())
	}
}

func TestHandleValidateTemplate(t *testing.T) {
	tmpDir := t.TempDir()

	// Valid template
	validPath := filepath.Join(tmpDir, "valid.tmpl")
	if err := os.WriteFile(validPath, []byte("# Hello {{ .ImageTag }}"), 0644); err != nil {
		t.Fatalf("failed to write file: %v", err)
	}
	output := captureOutput(func() {
		if err := handleValidateTemplate(validPath); err != nil {
			t.Fatalf("handleValidateTemplate() error = %v for valid template", err)
		}
	})
	if !strings.Contains(output, "is valid") {
		t.Error("expected 'is valid' message")
	}

	// Invalid template
	invalidPath := filepath.Join(tmpDir, "invalid.tmpl")
	if err := os.WriteFile(invalidPath, []byte("{{ .Unclosed"), 0644); err != nil {
		t.Fatalf("failed to write file: %v", err)
	}
	err := handleValidateTemplate(invalidPath)
	if err == nil {
		t.Fatal("expected error for invalid template")
	}

	// Nonexistent file
	err = handleValidateTemplate("/nonexistent/path.tmpl")
	if err == nil {
		t.Fatal("expected error for nonexistent file")
	}
}
