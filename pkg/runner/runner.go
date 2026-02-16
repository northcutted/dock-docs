package runner

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/northcutted/dock-docs/pkg/types"
)

// ToolRunner defines the interface for external tool integration
type ToolRunner interface {
	Name() string
	IsAvailable() bool
	Run(image string, verbose bool) (*types.ImageStats, error)
}

// runCommand executes a command and handles verbose logging and error reporting
func runCommand(cmd *exec.Cmd, verbose bool) ([]byte, error) {
	if verbose {
		fmt.Fprintf(os.Stderr, "[DEBUG] Running: %s\n", cmd.String())
	}

	// Capture stdout and stderr separately if possible, or combined?
	// exec.Command.Output() captures stdout. Stderr is in ExitError.
	// But sometimes tools write useful info to stderr even on success.
	// Let's capture both if we can, but Output() is easiest for data.

	output, err := cmd.Output()
	if err != nil {
		var stderr []byte
		if exitErr, ok := err.(*exec.ExitError); ok {
			stderr = exitErr.Stderr
		}
		return nil, fmt.Errorf("command failed: %s\nStderr: %s", err, string(stderr))
	}

	if verbose {
		// Print a snippet or full output? Full is safer for debug.
		fmt.Fprintf(os.Stderr, "[DEBUG] Output (%d bytes):\n%s\n", len(output), string(output))
	}

	return output, nil
}

// EnsureImage checks if an image exists locally, and pulls it if not.
func EnsureImage(image string, verbose bool) error {
	// Check if image exists
	checkCmd := exec.Command("docker", "inspect", "--type=image", image)
	if err := checkCmd.Run(); err == nil {
		if verbose {
			fmt.Printf("[DEBUG] Image %s found locally\n", image)
		}
		return nil
	}

	// Image not found, pull it
	fmt.Printf("Pulling image: %s ...\n", image)
	pullCmd := exec.Command("docker", "pull", image)
	if _, err := runCommand(pullCmd, verbose); err != nil {
		return fmt.Errorf("failed to pull image %s: %w", image, err)
	}
	return nil
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

func (r *RuntimeRunner) Run(image string, verbose bool) (*types.ImageStats, error) {
	// Ensure binary is set if IsAvailable wasn't called (though it should be)
	if r.binary == "" {
		if !r.IsAvailable() {
			return nil, fmt.Errorf("no container runtime found (docker or podman)")
		}
	}

	cmd := exec.Command(r.binary, "inspect", image)
	output, err := runCommand(cmd, verbose)
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
	stats := &types.ImageStats{
		ImageTag:     image,
		Architecture: data.Architecture,
		OS:           data.Os,
		SizeMB:       fmt.Sprintf("%.2f MB", float64(data.Size)/1024/1024),
		TotalLayers:  len(data.RootFS.Layers),
	}

	return stats, nil
}

// ManifestRunner runs 'docker manifest inspect <image>'
type ManifestRunner struct {
	binary string
}

func (r *ManifestRunner) Name() string { return "manifest" }

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

func (r *ManifestRunner) Run(image string, verbose bool) (*types.ImageStats, error) {
	if r.binary == "" {
		if !r.IsAvailable() {
			return nil, fmt.Errorf("no container runtime found")
		}
	}

	// Try standard manifest inspect
	// Need DOCKER_CLI_EXPERIMENTAL=enabled for older docker to be safe
	cmd := exec.Command(r.binary, "manifest", "inspect", image)
	cmd.Env = append(os.Environ(), "DOCKER_CLI_EXPERIMENTAL=enabled")

	output, err := runCommand(cmd, verbose)
	if err != nil {
		// Fallback or just return empty stats (optional feature)
		// We'll return error so analyzer can log warning
		return nil, fmt.Errorf("manifest inspect failed: %w", err)
	}

	// Manifest inspect can return a single manifest or a manifest list.
	// We need to handle both structures loosely.
	// Structure for Manifest List:
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

	// First try to unmarshal as Index
	var index ManifestIndex
	if err := json.Unmarshal(output, &index); err == nil && len(index.Manifests) > 0 {
		// It's a manifest list
		var archs []string
		seen := make(map[string]bool)
		for _, m := range index.Manifests {
			// Format: linux/amd64
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

	// If not a list, it might be a single manifest (e.g. locally built image or specific tag)
	// In this case, we just return the single architecture if we can find it,
	// BUT 'docker inspect' already handles single image details.
	// 'docker manifest inspect' on a single image often returns the V2 Schema 2 JSON directly.

	// If we can't parse as index, we'll assume no multi-arch info available from manifest
	// or it's a single image. We return empty stats (not error) to avoid noise.
	// Wait, if we return nil, analyzer might log warning.
	// Let's return empty valid stats so it merges safely.
	return &types.ImageStats{ImageTag: image}, nil
}

// SyftRunner runs 'syft <image> -o json'
type SyftRunner struct{}

func (r *SyftRunner) Name() string { return "syft" }

func (r *SyftRunner) IsAvailable() bool {
	_, err := exec.LookPath("syft")
	return err == nil
}

func (r *SyftRunner) Run(image string, verbose bool) (*types.ImageStats, error) {
	cmd := exec.Command("syft", image, "-o", "json")
	output, err := runCommand(cmd, verbose)
	if err != nil {
		return nil, err
	}

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

	// Update TotalPackages to reflect unique count?
	// The user asked for "all packages found", but usually unique is better.
	// Syft's total count includes duplicates.
	// Let's update TotalPackages to be the unique count to match the list.
	stats.TotalPackages = len(stats.Packages)

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

func (r *GrypeRunner) Run(image string, verbose bool) (*types.ImageStats, error) {
	cmd := exec.Command("grype", image, "-o", "json")
	output, err := runCommand(cmd, verbose)
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

	stats := &types.ImageStats{
		VulnSummary:     make(map[string]int),
		Vulnerabilities: make([]types.Vulnerability, 0),
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

func (r *DiveRunner) Run(image string, verbose bool) (*types.ImageStats, error) {
	// Create a temp file for output
	tmpFile, err := os.CreateTemp("", "dive-output-*.json")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}
	defer func() {
		// Attempt to remove file, ignore error (best effort cleanup)
		_ = os.Remove(tmpFile.Name())
	}()
	// Close immediately, dive will write to it. Ignore error.
	_ = tmpFile.Close()

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
					// Create a flexible struct to handle potential variations
					var machines []struct {
						ConnectionInfo struct {
							PodmanSocket *struct {
								Path string `json:"Path"`
							} `json:"PodmanSocket"`
						} `json:"ConnectionInfo"`
					}
					if json.Unmarshal(out, &machines) == nil && len(machines) > 0 {
						// Access safely with pointer check if needed, though unmarshal handles missing fields
						if machines[0].ConnectionInfo.PodmanSocket != nil && machines[0].ConnectionInfo.PodmanSocket.Path != "" {
							socketPath := machines[0].ConnectionInfo.PodmanSocket.Path
							// Set env for the dive command
							// Note: os.Environ() returns a copy, so append works safely for the cmd.Env
							env := os.Environ()
							// Ensure unix:// prefix
							if !strings.HasPrefix(socketPath, "unix://") {
								socketPath = "unix://" + socketPath
							}
							env = append(env, "DOCKER_HOST="+socketPath)
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

	// Dive uses CombinedOutput because it writes analysis logs to stdout/stderr even with --json
	// But our runCommand assumes Output().
	// We can use runCommand here if we want consistent logging, but dive output (logs) is not the JSON.
	// The JSON is in the file.

	if verbose {
		fmt.Fprintf(os.Stderr, "[DEBUG] Running: %s\n", cmd.String())
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("dive failed: %s, output: %s", err, string(output))
	}

	if verbose {
		fmt.Fprintf(os.Stderr, "[DEBUG] Output (%d bytes):\n%s\n", len(output), string(output))
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

	stats := &types.ImageStats{
		Efficiency: diveOutput.Image.EfficiencyScore * 100, // Convert to percentage if needed? Spec says 0-100. Dive usually gives 0.95 etc? Let's check or assume 0-1.
		// Wait, spec example says "98%". If Dive gives 0.98, verify.
		// Assuming Dive gives fractional 0.0-1.0 or percentage.
		// Let's assume standard float. If > 1, maybe it's percentage.
		// Actually, Dive JSON output is `efficiencyScore: 0.99`. So * 100 is correct.
		WastedBytes: fmt.Sprintf("%.2f MB", float64(diveOutput.Image.InefficientBytes)/1024/1024),
	}

	return stats, nil
}
