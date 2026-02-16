package cmd

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Helper to capture stdout
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	// These errors are typically ignored in test helpers, but linter complains
	_ = w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	return buf.String()
}

func TestExecute_DryRun(t *testing.T) {
	// Setup dummy Dockerfile
	tmpDir := t.TempDir()
	dockerfile := filepath.Join(tmpDir, "Dockerfile")
	content := "FROM alpine\nENV APP_PORT=8080"
	if err := os.WriteFile(dockerfile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write Dockerfile: %v", err)
	}

	// Override flags for test context
	// Cobra flags persist, so we need to reset or run in separate processes usually.
	// But since we are testing main execution logic, we can just call the RunE function directly or use SetArgs.

	// Reset flags for safety in this test run context
	rootCmd.SetArgs([]string{"--file", dockerfile, "--dry-run"})

	output := captureOutput(func() {
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("Execute failed: %v", err)
		}
	})

	if !strings.Contains(output, "`APP_PORT`") {
		t.Errorf("expected dry-run output to contain table row, got:\n%s", output)
	}
}

func TestExecute_Injection(t *testing.T) {
	// Reset flags manually because Cobra holds state in globals/pointers
	dockerfile = ""
	outputFile = ""
	dryRun = false
	imageTag = ""

	tmpDir := t.TempDir()
	dockerfileLoc := filepath.Join(tmpDir, "Dockerfile")
	readme := filepath.Join(tmpDir, "README.md")

	// Dockerfile content
	if err := os.WriteFile(dockerfileLoc, []byte("FROM alpine\nENV FOO=bar"), 0644); err != nil {
		t.Fatalf("failed to write Dockerfile: %v", err)
	}

	// README content with markers
	readmeContent := "# Title\n\n<!-- BEGIN: dock-docs -->\nOLD CONTENT\n<!-- END: dock-docs -->\n\nFooter"
	if err := os.WriteFile(readme, []byte(readmeContent), 0644); err != nil {
		t.Fatalf("failed to write README: %v", err)
	}

	// Create a new command for testing to avoid shared state issues or reset flags?
	// Cobra commands are hard to re-use if they bind to globals.
	// We will try setting args explicitly and hoping Execute resets parsing.
	rootCmd.SetArgs([]string{"--file", dockerfileLoc, "--output", readme})

	captureOutput(func() {
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("Execute failed: %v", err)
		}
	})

	// Verify file updated
	newContent, err := os.ReadFile(readme)
	if err != nil {
		t.Fatalf("failed to read updated README: %v", err)
	}

	sContent := string(newContent)
	if strings.Contains(sContent, "OLD CONTENT") {
		t.Error("expected OLD CONTENT to be replaced")
	}
	if !strings.Contains(sContent, "`FOO`") {
		t.Error("expected new table content in README")
	}
	if !strings.Contains(sContent, "<!-- BEGIN: dock-docs -->") {
		t.Error("expected markers to be preserved")
	}
}

func TestExecute_NoMarkers_Stdout(t *testing.T) {
	dockerfile = ""
	outputFile = ""
	dryRun = false
	imageTag = ""

	tmpDir := t.TempDir()
	dockerfile := filepath.Join(tmpDir, "Dockerfile")
	readme := filepath.Join(tmpDir, "README.md")

	if err := os.WriteFile(dockerfile, []byte("FROM alpine\nENV BAZ=qux"), 0644); err != nil {
		t.Fatalf("failed to write Dockerfile: %v", err)
	}

	// README without markers
	if err := os.WriteFile(readme, []byte("# Just a file"), 0644); err != nil {
		t.Fatalf("failed to write README: %v", err)
	}

	rootCmd.SetArgs([]string{"--file", dockerfile, "--output", readme})

	output := captureOutput(func() {
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("Execute failed: %v", err)
		}
	})

	// Should print to stdout because markers are missing
	if !strings.Contains(output, "`BAZ`") {
		t.Errorf("expected stdout output when markers are missing, got: %s", output)
	}

	// File should be unchanged
	content, _ := os.ReadFile(readme)
	if strings.Contains(string(content), "`BAZ`") {
		t.Error("expected file to remain unchanged when markers are missing")
	}
}

func TestExecute_WithImageFlag_FailsWithoutIgnore(t *testing.T) {
	// Smoke test for --image flag failing without ignore-errors
	dockerfile = ""
	outputFile = ""
	dryRun = false
	imageTag = ""
	ignoreErrors = false

	tmpDir := t.TempDir()
	dockerfile := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(dockerfile, []byte("FROM alpine"), 0644); err != nil {
		t.Fatalf("failed to write Dockerfile: %v", err)
	}

	rootCmd.SetArgs([]string{"--file", dockerfile, "--image", "fake-image:latest", "--dry-run"})

	output := captureOutput(func() {
		if err := rootCmd.Execute(); err == nil {
			t.Fatal("Execute expected to fail, but it succeeded")
		}
	})

	// It should print warning about analysis failure
	if !strings.Contains(output, "Warning: analysis failed") {
		t.Errorf("expected analysis warning, got:\n%s", output)
	}
}

func TestExecute_WithImageFlag_IgnoresErrors(t *testing.T) {
	// Smoke test for --image flag with --ignore-errors
	dockerfile = ""
	outputFile = ""
	dryRun = false
	imageTag = ""
	ignoreErrors = false

	tmpDir := t.TempDir()
	dockerfile := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(dockerfile, []byte("FROM alpine"), 0644); err != nil {
		t.Fatalf("failed to write Dockerfile: %v", err)
	}

	// Add --ignore-errors flag
	rootCmd.SetArgs([]string{"--file", dockerfile, "--image", "fake-image:latest", "--dry-run", "--ignore-errors"})

	output := captureOutput(func() {
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("Execute failed despite ignore-errors: %v", err)
		}
	})

	// It should print "Analyzing image: fake-image:latest ..."
	if !strings.Contains(output, "Analyzing image: fake-image:latest") {
		t.Errorf("expected analysis log, got:\n%s", output)
	}

	// And standard table (because it proceeds to render even if analysis fails)
	if !strings.Contains(output, "Configuration") {
		t.Errorf("expected standard table, got:\n%s", output)
	}
}
