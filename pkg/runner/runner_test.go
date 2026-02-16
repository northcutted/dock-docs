package runner

import (
	"testing"

	"github.com/northcutted/dock-docs/pkg/analysis"
)

// FakeRunner allows mocking external commands
// In real tests, we might not want to exec actual podman/docker unless integration testing.
// However, we can unit test the IsAvailable logic if we could mock exec.LookPath, which is hard.
// Instead, we'll write a test that checks the fallback logic structure if possible,
// or just verify the code compiles and interfaces match.

// Since the user asked for functional implementation, and we can't easily mock exec in Go without structural changes (like CommandContext injection),
// we will rely on manual verification or basic structural tests.

func TestRuntimeRunner_Name(t *testing.T) {
	r := &RuntimeRunner{}
	if r.Name() != "runtime" {
		t.Errorf("expected default name 'runtime', got %s", r.Name())
	}

	r.binary = "podman"
	if r.Name() != "podman" {
		t.Errorf("expected name 'podman', got %s", r.Name())
	}
}

// We can't easily test Run() without a real container runtime.
// But we can verify it implements the interface.
func TestImplementsInterface(t *testing.T) {
	var _ analysis.Runner = &RuntimeRunner{}
	var _ analysis.Runner = &DiveRunner{}
}
