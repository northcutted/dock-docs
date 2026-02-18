// Test file for YAML mode execution (runYAMLMode with config files).
//
// Globals mutated: dryRun, noMoji, ignoreErrors, verbose, templateName,
// debugTemplate, stdout (via captureOutput).
// All tests use defer resetFlags()() for cleanup.
package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunYAMLMode_ImageSection_DryRun(t *testing.T) {
	defer resetFlags()()

	tmpDir := t.TempDir()

	// Create Dockerfile
	df := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(df, []byte("FROM alpine\nENV MY_VAR=hello\nEXPOSE 3000"), 0644); err != nil {
		t.Fatal(err)
	}

	// Create output file with markers
	readme := filepath.Join(tmpDir, "README.md")
	if err := os.WriteFile(readme, []byte("# Heading\n\n<!-- BEGIN: dock-docs -->\nold\n<!-- END: dock-docs -->\n"), 0644); err != nil {
		t.Fatal(err)
	}

	// Create YAML config
	yamlContent := fmt.Sprintf(`output: %s
sections:
  - type: image
    marker: ""
    source: %s
`, readme, df)
	cfgPath := filepath.Join(tmpDir, "dock-docs.yaml")
	if err := os.WriteFile(cfgPath, []byte(yamlContent), 0644); err != nil {
		t.Fatal(err)
	}

	dryRun = true
	noMoji = false
	ignoreErrors = false
	verbose = false
	templateName = ""
	debugTemplate = false

	output := captureOutput(func() {
		if err := runYAMLMode(context.Background(), cfgPath); err != nil {
			t.Fatalf("runYAMLMode() error: %v", err)
		}
	})

	if !strings.Contains(output, "MY_VAR") {
		t.Errorf("expected MY_VAR in dry-run output, got:\n%s", output)
	}
}

func TestRunYAMLMode_ImageSection_WriteFile(t *testing.T) {
	defer resetFlags()()

	tmpDir := t.TempDir()

	df := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(df, []byte("FROM alpine\nENV WRITE_TEST=yes"), 0644); err != nil {
		t.Fatal(err)
	}

	readme := filepath.Join(tmpDir, "README.md")
	if err := os.WriteFile(readme, []byte("# Doc\n\n<!-- BEGIN: dock-docs -->\nplaceholder\n<!-- END: dock-docs -->\n\nFooter"), 0644); err != nil {
		t.Fatal(err)
	}

	yamlContent := fmt.Sprintf(`output: %s
sections:
  - type: image
    marker: ""
    source: %s
`, readme, df)
	cfgPath := filepath.Join(tmpDir, "dock-docs.yaml")
	if err := os.WriteFile(cfgPath, []byte(yamlContent), 0644); err != nil {
		t.Fatal(err)
	}

	dryRun = false
	noMoji = false
	templateName = ""

	captureOutput(func() {
		if err := runYAMLMode(context.Background(), cfgPath); err != nil {
			t.Fatalf("runYAMLMode() error: %v", err)
		}
	})

	content, err := os.ReadFile(readme)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(content), "WRITE_TEST") {
		t.Error("expected WRITE_TEST in updated README")
	}
	if strings.Contains(string(content), "placeholder") {
		t.Error("expected old placeholder to be replaced")
	}
	if !strings.Contains(string(content), "Footer") {
		t.Error("expected Footer to be preserved")
	}
}

func TestRunYAMLMode_HTMLSection_DryRun(t *testing.T) {
	defer resetFlags()()

	tmpDir := t.TempDir()

	df := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(df, []byte("FROM alpine\nENV HTML_VAR=test"), 0644); err != nil {
		t.Fatal(err)
	}

	readme := filepath.Join(tmpDir, "README.md")
	if err := os.WriteFile(readme, []byte("# Doc"), 0644); err != nil {
		t.Fatal(err)
	}

	yamlContent := fmt.Sprintf(`output: %s
sections:
  - type: image
    marker: main
    source: %s
    template:
      name: html
`, readme, df)
	cfgPath := filepath.Join(tmpDir, "dock-docs.yaml")
	if err := os.WriteFile(cfgPath, []byte(yamlContent), 0644); err != nil {
		t.Fatal(err)
	}

	dryRun = true
	templateName = ""

	output := captureOutput(func() {
		if err := runYAMLMode(context.Background(), cfgPath); err != nil {
			t.Fatalf("runYAMLMode() error: %v", err)
		}
	})

	if !strings.Contains(output, "<") {
		t.Error("expected HTML content in dry-run output")
	}
}

func TestRunYAMLMode_HTMLSection_WriteFile(t *testing.T) {
	defer resetFlags()()

	tmpDir := t.TempDir()

	df := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(df, []byte("FROM alpine\nENV DIRECT_WRITE=yes"), 0644); err != nil {
		t.Fatal(err)
	}

	readme := filepath.Join(tmpDir, "README.md")
	if err := os.WriteFile(readme, []byte("# Doc"), 0644); err != nil {
		t.Fatal(err)
	}

	yamlContent := fmt.Sprintf(`output: %s
sections:
  - type: image
    marker: main
    source: %s
    template:
      name: html
`, readme, df)
	cfgPath := filepath.Join(tmpDir, "dock-docs.yaml")
	if err := os.WriteFile(cfgPath, []byte(yamlContent), 0644); err != nil {
		t.Fatal(err)
	}

	dryRun = false
	templateName = ""

	captureOutput(func() {
		if err := runYAMLMode(context.Background(), cfgPath); err != nil {
			t.Fatalf("runYAMLMode() error: %v", err)
		}
	})

	// The HTML section should write to README-main.html
	htmlPath := filepath.Join(tmpDir, "README-main.html")
	content, err := os.ReadFile(htmlPath)
	if err != nil {
		t.Fatalf("expected HTML output file to be created: %v", err)
	}
	if !strings.Contains(string(content), "<") {
		t.Error("expected HTML content in written file")
	}
}

func TestRunYAMLMode_DebugTemplate(t *testing.T) {
	defer resetFlags()()

	tmpDir := t.TempDir()

	df := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(df, []byte("FROM alpine"), 0644); err != nil {
		t.Fatal(err)
	}

	readme := filepath.Join(tmpDir, "README.md")
	if err := os.WriteFile(readme, []byte("<!-- BEGIN: dock-docs -->\n<!-- END: dock-docs -->"), 0644); err != nil {
		t.Fatal(err)
	}

	yamlContent := fmt.Sprintf(`output: %s
sections:
  - type: image
    source: %s
`, readme, df)
	cfgPath := filepath.Join(tmpDir, "dock-docs.yaml")
	if err := os.WriteFile(cfgPath, []byte(yamlContent), 0644); err != nil {
		t.Fatal(err)
	}

	dryRun = true
	debugTemplate = true
	templateName = ""
	verbose = true // needed so slog level is DEBUG

	_, logOut := captureAll(func() {
		if err := runYAMLMode(context.Background(), cfgPath); err != nil {
			t.Fatalf("runYAMLMode() error: %v", err)
		}
	})

	if !strings.Contains(logOut, "template resolved") {
		t.Error("expected debug template info in log output")
	}
}

func TestRunYAMLMode_UnknownSectionType(t *testing.T) {
	defer resetFlags()()

	tmpDir := t.TempDir()

	readme := filepath.Join(tmpDir, "README.md")
	if err := os.WriteFile(readme, []byte("# Doc"), 0644); err != nil {
		t.Fatal(err)
	}

	yamlContent := fmt.Sprintf(`output: %s
sections:
  - type: bogus
    marker: test
`, readme)
	cfgPath := filepath.Join(tmpDir, "dock-docs.yaml")
	if err := os.WriteFile(cfgPath, []byte(yamlContent), 0644); err != nil {
		t.Fatal(err)
	}

	dryRun = true
	templateName = ""

	// Config validation now rejects unknown section types at load time
	err := runYAMLMode(context.Background(), cfgPath)
	if err == nil {
		t.Fatal("expected error for unknown section type")
	}
	if !strings.Contains(err.Error(), "invalid type") {
		t.Errorf("expected 'invalid type' in error, got: %v", err)
	}
}

func TestRunYAMLMode_BadConfigPath(t *testing.T) {
	err := runYAMLMode(context.Background(), "/nonexistent/config.yaml")
	if err == nil {
		t.Fatal("expected error for nonexistent config")
	}
}

func TestRunYAMLMode_ImageWithAnalysis_IgnoreErrors(t *testing.T) {
	defer resetFlags()()

	tmpDir := t.TempDir()

	df := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(df, []byte("FROM alpine\nENV ANALYZED=true"), 0644); err != nil {
		t.Fatal(err)
	}

	readme := filepath.Join(tmpDir, "README.md")
	if err := os.WriteFile(readme, []byte("<!-- BEGIN: dock-docs -->\nold\n<!-- END: dock-docs -->"), 0644); err != nil {
		t.Fatal(err)
	}

	yamlContent := fmt.Sprintf(`output: %s
sections:
  - type: image
    source: %s
    tag: "fake-nonexistent-image:latest"
`, readme, df)
	cfgPath := filepath.Join(tmpDir, "dock-docs.yaml")
	if err := os.WriteFile(cfgPath, []byte(yamlContent), 0644); err != nil {
		t.Fatal(err)
	}

	dryRun = true
	ignoreErrors = true
	templateName = ""
	verbose = false

	_, logOut := captureAll(func() {
		if err := runYAMLMode(context.Background(), cfgPath); err != nil {
			t.Fatalf("runYAMLMode() error: %v", err)
		}
	})

	if !strings.Contains(logOut, "analyzing image") {
		t.Errorf("expected analysis log in slog output, got:\n%s", logOut)
	}
}

func TestRunYAMLMode_ComparisonSection_EmptyImages(t *testing.T) {
	defer resetFlags()()

	tmpDir := t.TempDir()

	readme := filepath.Join(tmpDir, "README.md")
	if err := os.WriteFile(readme, []byte("<!-- BEGIN: comp -->\n<!-- END: comp -->"), 0644); err != nil {
		t.Fatal(err)
	}

	yamlContent := fmt.Sprintf(`output: %s
sections:
  - type: comparison
    marker: comp
    images: []
`, readme)
	cfgPath := filepath.Join(tmpDir, "dock-docs.yaml")
	if err := os.WriteFile(cfgPath, []byte(yamlContent), 0644); err != nil {
		t.Fatal(err)
	}

	dryRun = true
	templateName = ""

	// Config validation now rejects comparison sections with no images
	err := runYAMLMode(context.Background(), cfgPath)
	if err == nil {
		t.Fatal("expected error for comparison section with empty images")
	}
	if !strings.Contains(err.Error(), "at least one image") {
		t.Errorf("expected 'at least one image' in error, got: %v", err)
	}
}

func TestRunYAMLMode_NoMarkersInOutput(t *testing.T) {
	defer resetFlags()()

	tmpDir := t.TempDir()

	df := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(df, []byte("FROM alpine\nENV NO_MARKERS=yes"), 0644); err != nil {
		t.Fatal(err)
	}

	readme := filepath.Join(tmpDir, "README.md")
	// File without markers
	if err := os.WriteFile(readme, []byte("# Just a readme"), 0644); err != nil {
		t.Fatal(err)
	}

	yamlContent := fmt.Sprintf(`output: %s
sections:
  - type: image
    marker: missing
    source: %s
`, readme, df)
	cfgPath := filepath.Join(tmpDir, "dock-docs.yaml")
	if err := os.WriteFile(cfgPath, []byte(yamlContent), 0644); err != nil {
		t.Fatal(err)
	}

	dryRun = false
	templateName = ""

	_, logOut := captureAll(func() {
		// Should not error - just warn
		if err := runYAMLMode(context.Background(), cfgPath); err != nil {
			t.Fatalf("runYAMLMode() error: %v", err)
		}
	})

	if !strings.Contains(logOut, "injection failed") {
		t.Errorf("expected warning about missing markers in slog output, got:\n%s", logOut)
	}
}
