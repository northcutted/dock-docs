package runner

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"sort"

	"github.com/northcutted/dock-docs/pkg/types"
)

// SyftRunner runs 'syft <image> -o json'
type SyftRunner struct {
	binary string
}

// Name returns the display name for this runner.
func (r *SyftRunner) Name() string { return "syft" }

// IsAvailable checks whether the syft binary is installed.
func (r *SyftRunner) IsAvailable() bool {
	if path, err := lookupTool("syft"); err == nil {
		r.binary = path
		return true
	}
	return false
}

// Run executes 'syft <image> -o json' and parses the result.
// The provided context is used as the parent for the command timeout.
func (r *SyftRunner) Run(ctx context.Context, image string, verbose bool) (*types.ImageStats, error) {
	if r.binary == "" {
		if !r.IsAvailable() {
			return nil, fmt.Errorf("syft not found")
		}
	}
	runCtx, cancel := context.WithTimeout(ctx, TimeoutScan)
	defer cancel()
	cmd := exec.CommandContext(runCtx, r.binary, image, "-o", "json")
	output, err := runCommand(cmd, verbose)
	if err != nil {
		return nil, err
	}

	return parseSyftOutput(output)
}

// parseSyftOutput parses JSON output from 'syft <image> -o json'
// into ImageStats containing OS distro, package count, and deduplicated package list.
func parseSyftOutput(output []byte) (*types.ImageStats, error) {
	var syftOutput struct {
		Distro struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"distro"`
		Artifacts []struct {
			Name    string `json:"name"`
			Version string `json:"version"`
			Type    string `json:"type"`
		} `json:"artifacts"`
	}

	if err := json.Unmarshal(output, &syftOutput); err != nil {
		return nil, fmt.Errorf("failed to unmarshal syft output: %w", err)
	}

	stats := &types.ImageStats{
		TotalPackages: len(syftOutput.Artifacts),
		Packages:      make([]types.PackageSummary, 0),
	}

	if syftOutput.Distro.Name != "" {
		if syftOutput.Distro.Version != "" {
			stats.OSDistro = fmt.Sprintf("%s %s", syftOutput.Distro.Name, syftOutput.Distro.Version)
		} else {
			stats.OSDistro = syftOutput.Distro.Name
		}
	}

	seen := make(map[string]bool)
	for _, artifact := range syftOutput.Artifacts {
		key := artifact.Name + "@" + artifact.Version
		if seen[key] {
			continue
		}
		seen[key] = true
		stats.Packages = append(stats.Packages, types.PackageSummary{
			Name:    artifact.Name,
			Version: artifact.Version,
		})
	}

	stats.TotalPackages = len(stats.Packages)

	sort.Slice(stats.Packages, func(i, j int) bool {
		return stats.Packages[i].Name < stats.Packages[j].Name
	})

	return stats, nil
}
