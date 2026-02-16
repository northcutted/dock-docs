package runner

import (
	"testing"
)

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
