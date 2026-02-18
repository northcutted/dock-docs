// Test file for the setup command (printToolStatus, runSetup).
//
// Globals mutated: setupCheck, setupDir, setupForce, stdout (via captureOutput).
// All tests use defer resetFlags()() for cleanup.
package cmd

import (
	"strings"
	"testing"
)

func TestPrintToolStatus(t *testing.T) {
	output := captureOutput(func() {
		if err := printToolStatus(t.TempDir()); err != nil {
			t.Fatalf("printToolStatus() error: %v", err)
		}
	})

	if !strings.Contains(output, "Tool Status:") {
		t.Error("expected 'Tool Status:' header in output")
	}
	// Should mention all three tools
	for _, tool := range []string{"syft", "grype", "dive"} {
		if !strings.Contains(output, tool) {
			t.Errorf("expected %q in tool status output", tool)
		}
	}
}

func TestRunSetup_CheckOnly(t *testing.T) {
	defer resetFlags()()

	setupCheck = true
	setupDir = t.TempDir()
	setupForce = false

	output := captureOutput(func() {
		if err := runSetup(setupCmd, nil); err != nil {
			t.Fatalf("runSetup(--check) error: %v", err)
		}
	})

	if !strings.Contains(output, "Tool Status:") {
		t.Error("expected tool status output from --check")
	}
}
