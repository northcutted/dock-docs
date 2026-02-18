package cmd

import (
	"bytes"
	"strings"
	"testing"
)

// resetFlags saves the current values of all package-level mutable globals
// and returns a function that restores them. Tests that mutate CLI flags,
// version vars, setup vars, or the stdout/logOutput writers should call:
//
//	defer resetFlags()()
//
// This eliminates fragile manual save/restore boilerplate and ensures globals
// are always restored, even if a test panics.
func resetFlags() func() {
	// Root flags (root.go)
	savedDockerfile := dockerfile
	savedOutputFile := outputFile
	savedDryRun := dryRun
	savedImageTag := imageTag
	savedConfigFile := configFile
	savedNoMoji := noMoji
	savedIgnoreErrors := ignoreErrors
	savedVerbose := verbose
	savedBadgeBaseURL := badgeBaseURL
	savedTemplateName := templateName
	savedListTemplates := listTemplates
	savedExportTemplate := exportTemplate
	savedValidateTemplate := validateTemplate
	savedDebugTemplate := debugTemplate
	savedAnalysisTimeout := analysisTimeout
	savedStdout := stdout
	savedLogOutput := logOutput

	// Setup flags (setup.go)
	savedSetupDir := setupDir
	savedSetupForce := setupForce
	savedSetupCheck := setupCheck

	// Version vars (version.go)
	savedVersion := Version
	savedCommit := Commit
	savedDate := Date

	return func() {
		dockerfile = savedDockerfile
		outputFile = savedOutputFile
		dryRun = savedDryRun
		imageTag = savedImageTag
		configFile = savedConfigFile
		noMoji = savedNoMoji
		ignoreErrors = savedIgnoreErrors
		verbose = savedVerbose
		badgeBaseURL = savedBadgeBaseURL
		templateName = savedTemplateName
		listTemplates = savedListTemplates
		exportTemplate = savedExportTemplate
		validateTemplate = savedValidateTemplate
		debugTemplate = savedDebugTemplate
		analysisTimeout = savedAnalysisTimeout
		stdout = savedStdout
		logOutput = savedLogOutput

		setupDir = savedSetupDir
		setupForce = savedSetupForce
		setupCheck = savedSetupCheck

		Version = savedVersion
		Commit = savedCommit
		Date = savedDate
	}
}

// captureOutput swaps the package-level stdout var to a buffer, runs f, and
// returns whatever was written. This is safe for parallel tests since it does
// not touch os.Stdout.
func captureOutput(f func()) string {
	var buf bytes.Buffer
	old := stdout
	stdout = &buf

	f()

	stdout = old
	return buf.String()
}

// captureAll swaps both stdout and logOutput to buffers, configures slog via
// initLogger(), runs f, and returns the captured stdout and log output.
// This should be used by tests that need to verify slog messages.
func captureAll(f func()) (stdoutStr, logStr string) {
	var stdoutBuf, logBuf bytes.Buffer
	oldStdout := stdout
	oldLogOutput := logOutput
	stdout = &stdoutBuf
	logOutput = &logBuf

	// Reconfigure slog to write to the captured logOutput
	initLogger()

	f()

	stdout = oldStdout
	logOutput = oldLogOutput
	return stdoutBuf.String(), logBuf.String()
}

func TestCheckToolStatus(t *testing.T) {
	result := checkToolStatus()
	if !strings.Contains(result, "Prerequisites:") {
		t.Errorf("expected 'Prerequisites:' in output, got: %s", result)
	}
	// Should mention docker/podman
	if !strings.Contains(result, "docker") && !strings.Contains(result, "podman") {
		t.Error("expected docker or podman mention in tool status")
	}
	// Should mention each tool
	for _, tool := range []string{"syft", "grype", "dive"} {
		if !strings.Contains(result, tool) {
			t.Errorf("expected %q in tool status", tool)
		}
	}
}
