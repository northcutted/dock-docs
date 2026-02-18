package runner

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os/exec"
	"time"

	"github.com/northcutted/dock-docs/pkg/types"
)

// GrypeRunner runs 'grype <image> -o json'
type GrypeRunner struct {
	binary string
}

// Name returns the display name for this runner.
func (r *GrypeRunner) Name() string { return "grype" }

// IsAvailable checks whether the grype binary is installed.
func (r *GrypeRunner) IsAvailable() bool {
	if path, err := lookupTool("grype"); err == nil {
		r.binary = path
		return true
	}
	return false
}

// Run executes 'grype <image> -o json' and parses the result.
// The provided context is used as the parent for the command timeout.
func (r *GrypeRunner) Run(ctx context.Context, image string, verbose bool) (*types.ImageStats, error) {
	if r.binary == "" {
		if !r.IsAvailable() {
			return nil, fmt.Errorf("grype not found")
		}
	}
	runCtx, cancel := context.WithTimeout(ctx, TimeoutScan)
	defer cancel()
	cmd := exec.CommandContext(runCtx, r.binary, image, "-o", "json")
	output, err := runCommand(cmd, verbose)
	if err != nil {
		return nil, err
	}

	return parseGrypeOutput(output, verbose)
}

// parseGrypeOutput parses JSON output from 'grype <image> -o json'
// into ImageStats containing vulnerability summary, details, and scan time.
func parseGrypeOutput(output []byte, verbose bool) (*types.ImageStats, error) {
	var grypeOutput struct {
		Descriptor struct {
			Timestamp string `json:"timestamp"`
		} `json:"descriptor"`
		Matches []struct {
			Vulnerability struct {
				ID       string `json:"id"`
				Severity string `json:"severity"`
			} `json:"vulnerability"`
			Artifact struct {
				Name    string `json:"name"`
				Version string `json:"version"`
			} `json:"artifact"`
		} `json:"matches"`
	}

	if err := json.Unmarshal(output, &grypeOutput); err != nil {
		return nil, fmt.Errorf("failed to unmarshal grype output: %w", err)
	}

	scanTime := time.Now()
	if grypeOutput.Descriptor.Timestamp != "" {
		if parsedTime, err := time.Parse(time.RFC3339, grypeOutput.Descriptor.Timestamp); err == nil {
			scanTime = parsedTime
		} else if verbose {
			slog.Debug("failed to parse grype timestamp", "error", err)
		}
	}

	stats := &types.ImageStats{
		VulnSummary:     make(map[string]int),
		Vulnerabilities: make([]types.Vulnerability, 0),
		VulnScanTime:    scanTime,
	}

	for _, match := range grypeOutput.Matches {
		sev := match.Vulnerability.Severity
		stats.VulnSummary[sev]++

		stats.Vulnerabilities = append(stats.Vulnerabilities, types.Vulnerability{
			ID:       match.Vulnerability.ID,
			Severity: sev,
			Package:  match.Artifact.Name,
			Version:  match.Artifact.Version,
		})
	}

	types.SortBySeverity(stats.Vulnerabilities)

	return stats, nil
}
