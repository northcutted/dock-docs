package runner

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"sort"

	"github.com/northcutted/dock-docs/pkg/types"
)

// ManifestRunner runs 'docker manifest inspect <image>'
type ManifestRunner struct {
	binary string
}

// Name returns the display name for this runner.
func (r *ManifestRunner) Name() string { return "manifest" }

// IsAvailable checks whether a container runtime is installed for manifest inspection.
func (r *ManifestRunner) IsAvailable() bool {
	// Check docker
	if _, err := exec.LookPath("docker"); err == nil {
		r.binary = "docker"
		return true
	}
	// Podman manifest inspect also works
	if _, err := exec.LookPath("podman"); err == nil {
		r.binary = "podman"
		return true
	}
	return false
}

// Run executes 'docker manifest inspect' or 'podman manifest inspect' and parses the result.
// The provided context is used as the parent for the command timeout.
func (r *ManifestRunner) Run(ctx context.Context, image string, verbose bool) (*types.ImageStats, error) {
	if r.binary == "" {
		if !r.IsAvailable() {
			return nil, fmt.Errorf("no container runtime found")
		}
	}

	// Try standard manifest inspect
	// Need DOCKER_CLI_EXPERIMENTAL=enabled for older docker to be safe
	runCtx, cancel := context.WithTimeout(ctx, TimeoutInspect)
	defer cancel()
	cmd := exec.CommandContext(runCtx, r.binary, "manifest", "inspect", image)
	cmd.Env = append(os.Environ(), "DOCKER_CLI_EXPERIMENTAL=enabled")

	output, err := runCommand(cmd, verbose)
	if err != nil {
		// Fallback or just return empty stats (optional feature)
		// We'll return error so analyzer can log warning
		return nil, fmt.Errorf("manifest inspect failed: %w", err)
	}

	return parseManifestInspect(output, image)
}

// parseManifestInspect parses JSON output from 'docker manifest inspect'
// into ImageStats containing supported architectures for multi-arch images.
// The error return is always nil because manifest data is optional and gracefully
// degrades to empty stats when parsing fails or no manifests are found.
func parseManifestInspect(output []byte, image string) (*types.ImageStats, error) { //nolint:unparam // error kept for interface consistency
	type Platform struct {
		Architecture string `json:"architecture"`
		OS           string `json:"os"`
	}
	type Manifest struct {
		Platform Platform `json:"platform"`
	}
	type ManifestIndex struct {
		Manifests []Manifest `json:"manifests"`
	}

	var index ManifestIndex
	if err := json.Unmarshal(output, &index); err == nil && len(index.Manifests) > 0 {
		var archs []string
		seen := make(map[string]bool)
		for _, m := range index.Manifests {
			key := fmt.Sprintf("%s/%s", m.Platform.OS, m.Platform.Architecture)
			if !seen[key] {
				seen[key] = true
				archs = append(archs, key)
			}
		}
		sort.Strings(archs)
		return &types.ImageStats{
			ImageTag:               image,
			SupportedArchitectures: archs,
		}, nil
	}

	// Not a manifest list or unparseable - return empty stats (not error)
	// so the analyzer can merge safely.
	return &types.ImageStats{ImageTag: image}, nil
}
