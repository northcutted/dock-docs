package runner

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"sort"

	"docker-docs/pkg/analysis"
)

// ToolRunner defines the interface for external tool integration
type ToolRunner interface {
	Name() string
	IsAvailable() bool
	Run(image string) (*analysis.ImageStats, error)
}

// RuntimeRunner runs 'docker inspect' or 'podman inspect'
type RuntimeRunner struct {
	binary string
}

func (r *RuntimeRunner) Name() string {
	if r.binary != "" {
		return r.binary
	}
	return "runtime"
}

func (r *RuntimeRunner) IsAvailable() bool {
	// Check docker first
	if _, err := exec.LookPath("docker"); err == nil {
		r.binary = "docker"
		return true
	}
	// Fallback to podman
	if _, err := exec.LookPath("podman"); err == nil {
		r.binary = "podman"
		return true
	}
	return false
}

func (r *RuntimeRunner) Run(image string) (*analysis.ImageStats, error) {
	// Ensure binary is set if IsAvailable wasn't called (though it should be)
	if r.binary == "" {
		if !r.IsAvailable() {
			return nil, fmt.Errorf("no container runtime found (docker or podman)")
		}
	}

	cmd := exec.Command(r.binary, "inspect", image)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var inspect []struct {
		Architecture string `json:"Architecture"`
		Os           string `json:"Os"`
		Size         int64  `json:"Size"`
		RootFS       struct {
			Layers []string `json:"Layers"`
		} `json:"RootFS"`
	}

	if err := json.Unmarshal(output, &inspect); err != nil {
		return nil, fmt.Errorf("failed to unmarshal %s inspect output: %w", r.binary, err)
	}

	if len(inspect) == 0 {
		return nil, fmt.Errorf("no inspect data returned for image %s", image)
	}

	data := inspect[0]
	stats := &analysis.ImageStats{
		ImageTag:     image,
		Architecture: data.Architecture,
		OS:           data.Os,
		SizeMB:       fmt.Sprintf("%.2f MB", float64(data.Size)/1024/1024),
		TotalLayers:  len(data.RootFS.Layers),
	}

	return stats, nil
}

// SyftRunner runs 'syft <image> -o json'
type SyftRunner struct{}

func (r *SyftRunner) Name() string { return "syft" }

func (r *SyftRunner) IsAvailable() bool {
	_, err := exec.LookPath("syft")
	return err == nil
}

func (r *SyftRunner) Run(image string) (*analysis.ImageStats, error) {
	cmd := exec.Command("syft", image, "-o", "json")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var syftOutput struct {
		Artifacts []struct {
			Name    string `json:"name"`
			Version string `json:"version"`
			Type    string `json:"type"`
		} `json:"artifacts"`
	}

	if err := json.Unmarshal(output, &syftOutput); err != nil {
		return nil, fmt.Errorf("failed to unmarshal syft output: %w", err)
	}

	stats := &analysis.ImageStats{
		TotalPackages: len(syftOutput.Artifacts),
		Packages:      make([]analysis.PackageSummary, 0),
	}

	for _, artifact := range syftOutput.Artifacts {
		stats.Packages = append(stats.Packages, analysis.PackageSummary{
			Name:    artifact.Name,
			Version: artifact.Version,
		})
	}

	// Sort packages alphabetically by name for cleaner output
	sort.Slice(stats.Packages, func(i, j int) bool {
		return stats.Packages[i].Name < stats.Packages[j].Name
	})

	return stats, nil
}

// GrypeRunner runs 'grype <image> -o json'
type GrypeRunner struct{}

func (r *GrypeRunner) Name() string { return "grype" }

func (r *GrypeRunner) IsAvailable() bool {
	_, err := exec.LookPath("grype")
	return err == nil
}

func (r *GrypeRunner) Run(image string) (*analysis.ImageStats, error) {
	cmd := exec.Command("grype", image, "-o", "json")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var grypeOutput struct {
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

	stats := &analysis.ImageStats{
		VulnSummary:     make(map[string]int),
		Vulnerabilities: make([]analysis.Vulnerability, 0),
	}

	for _, match := range grypeOutput.Matches {
		sev := match.Vulnerability.Severity
		stats.VulnSummary[sev]++

		stats.Vulnerabilities = append(stats.Vulnerabilities, analysis.Vulnerability{
			ID:       match.Vulnerability.ID,
			Severity: sev,
			Package:  match.Artifact.Name,
			Version:  match.Artifact.Version,
		})
	}

	// Sort vulnerabilities by severity (Critical > High > Medium > Low)
	severityRank := map[string]int{
		"Critical": 4,
		"High":     3,
		"Medium":   2,
		"Low":      1,
		"Unknown":  0,
	}

	sort.Slice(stats.Vulnerabilities, func(i, j int) bool {
		// Sort by severity (descending)
		rankI := severityRank[stats.Vulnerabilities[i].Severity]
		rankJ := severityRank[stats.Vulnerabilities[j].Severity]
		if rankI != rankJ {
			return rankI > rankJ
		}
		// Then by ID (ascending) for stability
		return stats.Vulnerabilities[i].ID < stats.Vulnerabilities[j].ID
	})

	return stats, nil
}

// DiveRunner runs 'dive <image> --json output.json'
type DiveRunner struct{}

func (r *DiveRunner) Name() string { return "dive" }

func (r *DiveRunner) IsAvailable() bool {
	_, err := exec.LookPath("dive")
	return err == nil
}

func (r *DiveRunner) Run(image string) (*analysis.ImageStats, error) {
	// Create a temp file for output
	tmpFile, err := os.CreateTemp("", "dive-output-*.json")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close() // Close immediately, dive will write to it

	// Note: CI=true suppresses interactive UI which might be default for dive?
	// The --json flag should suffice.
	cmd := exec.Command("dive", image, "--json", tmpFile.Name())

	// Podman Support: Check if running via Podman and if DOCKER_HOST is missing
	// If 'docker' command is missing but 'podman' is present, we likely need to help dive connect.
	if _, err := exec.LookPath("docker"); err != nil {
		if _, err := exec.LookPath("podman"); err == nil {
			// Podman detected. Check DOCKER_HOST
			if os.Getenv("DOCKER_HOST") == "" {
				// Try to detect Podman socket via 'podman machine inspect' (macOS/Windows)
				// Note: On Linux, the socket is usually at a standard path, but 'machine' command might fail.
				// We'll wrap this in a helper or just try.

				// Using JSON approach for robustness
				pCmd := exec.Command("podman", "machine", "inspect")
				if out, err := pCmd.Output(); err == nil {
					var machines []struct {
						ConnectionInfo struct {
							PodmanSocket struct {
								Path string `json:"Path"`
							} `json:"PodmanSocket"`
						} `json:"ConnectionInfo"`
					}
					if json.Unmarshal(out, &machines) == nil && len(machines) > 0 {
						socketPath := machines[0].ConnectionInfo.PodmanSocket.Path
						if socketPath != "" {
							// Set env for the dive command
							// Note: os.Environ() returns a copy, so append works safely for the cmd.Env
							env := os.Environ()
							env = append(env, "DOCKER_HOST=unix://"+socketPath)
							cmd.Env = env
						}
					}
				}
				// If machine inspect fails (e.g. Linux native podman, not machine),
				// we might check `podman info` or standard paths.
				// But let's stick to the specific request logic first.
			}
		}
	}

	// Dive writes to file, but might output logs to stdout/stderr. capture or ignore?
	// cmd.CombinedOutput() might be useful for debugging if it fails.
	if output, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("dive failed: %s, output: %s", err, string(output))
	}

	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to read dive output: %w", err)
	}

	var diveOutput struct {
		Image struct {
			InefficientBytes uint64  `json:"inefficientBytes"`
			EfficiencyScore  float64 `json:"efficiencyScore"`
		} `json:"image"`
	}

	if err := json.Unmarshal(content, &diveOutput); err != nil {
		return nil, fmt.Errorf("failed to unmarshal dive output: %w", err)
	}

	stats := &analysis.ImageStats{
		Efficiency: diveOutput.Image.EfficiencyScore * 100, // Convert to percentage if needed? Spec says 0-100. Dive usually gives 0.95 etc? Let's check or assume 0-1.
		// Wait, spec example says "98%". If Dive gives 0.98, verify.
		// Assuming Dive gives fractional 0.0-1.0 or percentage.
		// Let's assume standard float. If > 1, maybe it's percentage.
		// Actually, Dive JSON output is `efficiencyScore: 0.99`. So * 100 is correct.
		WastedBytes: fmt.Sprintf("%.2f MB", float64(diveOutput.Image.InefficientBytes)/1024/1024),
	}

	return stats, nil
}
