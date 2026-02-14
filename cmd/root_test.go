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

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
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

	if !strings.Contains(output, "| APP_PORT | ENV |") {
		t.Errorf("expected dry-run output to contain table row, got:\n%s", output)
	}
}

func TestExecute_Injection(t *testing.T) {
	// Reset flags manually because Cobra holds state in globals/pointers
	dockerfile = ""
	outputFile = ""
	dryRun = false

	tmpDir := t.TempDir()
	dockerfileLoc := filepath.Join(tmpDir, "Dockerfile")
	readme := filepath.Join(tmpDir, "README.md")

	// Dockerfile content
	if err := os.WriteFile(dockerfileLoc, []byte("FROM alpine\nENV FOO=bar"), 0644); err != nil {
		t.Fatalf("failed to write Dockerfile: %v", err)
	}

	// README content with markers
	readmeContent := "# Title\n\n<!-- BEGIN: docker-docs -->\nOLD CONTENT\n<!-- END: docker-docs -->\n\nFooter"
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
	if !strings.Contains(sContent, "| FOO | ENV |") {
		t.Error("expected new table content in README")
	}
	if !strings.Contains(sContent, "<!-- BEGIN: docker-docs -->") {
		t.Error("expected markers to be preserved")
	}
}

func TestExecute_NoMarkers_Stdout(t *testing.T) {
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
	if !strings.Contains(output, "| BAZ | ENV |") {
		t.Errorf("expected stdout output when markers are missing, got: %s", output)
	}

	// File should be unchanged
	content, _ := os.ReadFile(readme)
	if strings.Contains(string(content), "| BAZ | ENV |") {
		t.Error("expected file to remain unchanged when markers are missing")
	}
}
