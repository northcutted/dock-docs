// Test file for the version command.
//
// Globals mutated: Version, Commit, Date, stdout (via captureOutput).
// All tests use defer resetFlags()() for cleanup.
package cmd

import (
	"strings"
	"testing"
)

func TestVersionCommand(t *testing.T) {
	defer resetFlags()()

	Version = "1.2.3"
	Commit = "abc123"
	Date = "2025-01-01"

	rootCmd.SetArgs([]string{"version"})
	output := captureOutput(func() {
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("version command failed: %v", err)
		}
	})

	if !strings.Contains(output, "1.2.3") {
		t.Errorf("expected version in output, got: %s", output)
	}
	if !strings.Contains(output, "abc123") {
		t.Errorf("expected commit in output, got: %s", output)
	}
	if !strings.Contains(output, "2025-01-01") {
		t.Errorf("expected date in output, got: %s", output)
	}
}
