// Test file for CLI mode execution (runCLIMode and rootCmd.Execute with CLI flags).
//
// Globals mutated: dockerfile, outputFile, dryRun, imageTag, ignoreErrors,
// templateName, debugTemplate, stdout (via captureOutput or direct swap).
// All tests use defer resetFlags()() for cleanup.
package cmd

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExecute_DryRun(t *testing.T) {
	defer resetFlags()()

	// Setup dummy Dockerfile
	tmpDir := t.TempDir()
	dockerfile := filepath.Join(tmpDir, "Dockerfile")
	content := "FROM alpine\nENV APP_PORT=8080"
	if err := os.WriteFile(dockerfile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write Dockerfile: %v", err)
	}

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
	defer resetFlags()()

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
	defer resetFlags()()

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
	defer resetFlags()()

	tmpDir := t.TempDir()
	dockerfile := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(dockerfile, []byte("FROM alpine"), 0644); err != nil {
		t.Fatalf("failed to write Dockerfile: %v", err)
	}

	rootCmd.SetArgs([]string{"--file", dockerfile, "--image", "fake-image:latest", "--dry-run"})

	_, logOutput := captureAll(func() {
		if err := rootCmd.Execute(); err == nil {
			t.Fatal("Execute expected to fail, but it succeeded")
		}
	})

	// It should log warning about analysis failure
	if !strings.Contains(logOutput, "analysis failed") {
		t.Errorf("expected analysis warning in log output, got:\n%s", logOutput)
	}
}

func TestExecute_WithImageFlag_IgnoresErrors(t *testing.T) {
	defer resetFlags()()

	tmpDir := t.TempDir()
	dockerfile := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(dockerfile, []byte("FROM alpine"), 0644); err != nil {
		t.Fatalf("failed to write Dockerfile: %v", err)
	}

	// Add --ignore-errors flag
	rootCmd.SetArgs([]string{"--file", dockerfile, "--image", "fake-image:latest", "--dry-run", "--ignore-errors"})

	stdoutOut, logOut := captureAll(func() {
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("Execute failed despite ignore-errors: %v", err)
		}
	})

	// It should log "analyzing image" to slog
	if !strings.Contains(logOut, "analyzing image") {
		t.Errorf("expected analysis log, got:\n%s", logOut)
	}

	// And standard table (because it proceeds to render even if analysis fails)
	if !strings.Contains(stdoutOut, "Configuration") {
		t.Errorf("expected standard table, got:\n%s", stdoutOut)
	}
}

func TestExecute_ListTemplatesFlag(t *testing.T) {
	defer resetFlags()()

	rootCmd.SetArgs([]string{"--list-templates"})
	output := captureOutput(func() {
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("Execute with --list-templates failed: %v", err)
		}
	})
	if !strings.Contains(output, "Available built-in templates:") {
		t.Error("expected template listing output")
	}
}

func TestExecute_DryRunWithHTMLTemplate(t *testing.T) {
	defer resetFlags()()

	tmpDir := t.TempDir()
	df := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(df, []byte("FROM alpine\nENV PORT=8080"), 0644); err != nil {
		t.Fatalf("failed to write Dockerfile: %v", err)
	}

	rootCmd.SetArgs([]string{"--file", df, "--dry-run", "--template", "html"})
	output := captureOutput(func() {
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("Execute with html template failed: %v", err)
		}
	})
	// HTML output should contain HTML tags
	if !strings.Contains(output, "<") {
		t.Error("expected HTML content in dry-run output")
	}
}

func TestRunCLIMode_HTMLTemplate_WriteFile(t *testing.T) {
	defer resetFlags()()

	tmpDir := t.TempDir()

	df := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(df, []byte("FROM alpine\nENV CLI_HTML=yes"), 0644); err != nil {
		t.Fatal(err)
	}

	dockerfile = df
	outputFile = filepath.Join(tmpDir, "output.html")
	dryRun = false
	imageTag = ""
	templateName = "html"

	captureOutput(func() {
		if err := runCLIMode(context.Background()); err != nil {
			t.Fatalf("runCLIMode(context.Background()) error: %v", err)
		}
	})

	content, err := os.ReadFile(filepath.Join(tmpDir, "output.html"))
	if err != nil {
		t.Fatalf("expected output file to be created: %v", err)
	}
	if !strings.Contains(string(content), "<") {
		t.Error("expected HTML content in output file")
	}
}

func TestRunCLIMode_NonexistentOutputFile(t *testing.T) {
	defer resetFlags()()

	tmpDir := t.TempDir()

	df := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(df, []byte("FROM alpine\nENV NOFILE=test"), 0644); err != nil {
		t.Fatal(err)
	}

	dockerfile = df
	outputFile = filepath.Join(tmpDir, "nonexistent.md")
	dryRun = false
	imageTag = ""
	templateName = ""

	var runErr error
	stdoutOut, logOut := captureAll(func() {
		runErr = runCLIMode(context.Background())
	})

	if runErr != nil {
		t.Fatalf("runCLIMode(context.Background()) error: %v", runErr)
	}

	// Should log warning about nonexistent file via slog
	if !strings.Contains(logOut, "does not exist") {
		t.Errorf("expected 'does not exist' warning in log output, got:\nstdout: %s\nlog: %s", stdoutOut, logOut)
	}
	// Content should be printed to stdout
	if !strings.Contains(stdoutOut, "NOFILE") {
		t.Errorf("expected NOFILE in stdout output, got:\n%s", stdoutOut)
	}
}

func TestRunCLIMode_DebugTemplate(t *testing.T) {
	defer resetFlags()()

	tmpDir := t.TempDir()

	df := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(df, []byte("FROM alpine"), 0644); err != nil {
		t.Fatal(err)
	}

	dockerfile = df
	dryRun = true
	imageTag = ""
	templateName = ""
	debugTemplate = true
	verbose = true // needed so slog level is DEBUG

	_, logOut := captureAll(func() {
		if err := runCLIMode(context.Background()); err != nil {
			t.Fatalf("runCLIMode(context.Background()) error: %v", err)
		}
	})

	if !strings.Contains(logOut, "template resolved") {
		t.Error("expected debug template info in log output")
	}
}

func TestRunCLIMode_BadDockerfile(t *testing.T) {
	defer resetFlags()()

	dockerfile = "/nonexistent/Dockerfile"
	err := runCLIMode(context.Background())
	if err == nil {
		t.Fatal("expected error for nonexistent Dockerfile")
	}
	if !strings.Contains(err.Error(), "failed to parse Dockerfile") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestRunCLIMode_TimeoutFlag(t *testing.T) {
	defer resetFlags()()

	tmpDir := t.TempDir()
	df := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(df, []byte("FROM alpine"), 0644); err != nil {
		t.Fatal(err)
	}

	// Use --timeout flag via rootCmd to verify it's wired through
	rootCmd.SetArgs([]string{"--file", df, "--dry-run", "--timeout", "30s"})
	output := captureOutput(func() {
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("Execute with --timeout failed: %v", err)
		}
	})

	// Should succeed (no image analysis, so timeout doesn't fire)
	if !strings.Contains(output, "Configuration") {
		t.Errorf("expected standard output with --timeout flag, got:\n%s", output)
	}
}

func TestRunCLIMode_AlreadyCancelledContext(t *testing.T) {
	defer resetFlags()()

	tmpDir := t.TempDir()
	df := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(df, []byte("FROM alpine\nENV PORT=8080"), 0644); err != nil {
		t.Fatal(err)
	}

	dockerfile = df
	dryRun = true
	imageTag = "fake-image:latest"

	// Use an already-cancelled context to verify context propagation
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // immediately cancel

	err := runCLIMode(ctx)
	if err == nil {
		t.Fatal("expected error with cancelled context when analyzing image")
	}
}
