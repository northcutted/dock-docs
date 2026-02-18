package runner

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"

	"github.com/northcutted/dock-docs/pkg/installer"
	"github.com/northcutted/dock-docs/pkg/types"
)

// lookupTool resolves the path to an external tool binary. It checks
// the system PATH first and falls back to the dock-docs install
// directory (~/.dock-docs/bin/).
var lookupTool = func(name string) (string, error) {
	path, _, err := installer.FindTool(name)
	return path, err
}

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
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			stderr = exitErr.Stderr
		}
		return nil, fmt.Errorf("command failed: %w\nStderr: %s", err, string(stderr))
	}

	if verbose {
		// Print a snippet or full output? Full is safer for debug.
		fmt.Fprintf(os.Stderr, "[DEBUG] Output (%d bytes):\n%s\n", len(output), string(output))
	}

	return output, nil
}

// EnsureImage checks if an image exists locally, and pulls it if not.
func EnsureImage(image string, verbose bool) error {
	// Detect which container runtime is available
	binary := ""
	if _, err := exec.LookPath("docker"); err == nil {
		binary = "docker"
	} else if _, err := exec.LookPath("podman"); err == nil {
		binary = "podman"
	} else {
		return fmt.Errorf("no container runtime found (docker or podman)")
	}

	// Check if image exists
	checkCmd := exec.CommandContext(context.Background(), binary, "inspect", "--type=image", image)
	if err := checkCmd.Run(); err == nil {
		if verbose {
			fmt.Printf("[DEBUG] Image %s found locally\n", image)
		}
		return nil
	}

	// Image not found, pull it
	fmt.Printf("Pulling image: %s ...\n", image)
	pullCmd := exec.CommandContext(context.Background(), binary, "pull", image)
	if _, err := runCommand(pullCmd, verbose); err != nil {
		return fmt.Errorf("failed to pull image %s: %w", image, err)
	}
	return nil
}

// RuntimeRunner runs 'docker inspect' or 'podman inspect'
type RuntimeRunner struct {
	binary string
}

// Name returns the display name for this runner.
func (r *RuntimeRunner) Name() string {
	if r.binary != "" {
		return r.binary
	}
	return "runtime"
}

// IsAvailable checks whether a container runtime (docker or podman) is installed.
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

// Run executes 'docker inspect' or 'podman inspect' and parses the result.
func (r *RuntimeRunner) Run(image string, verbose bool) (*types.ImageStats, error) {
	// Ensure binary is set if IsAvailable wasn't called (though it should be)
	if r.binary == "" {
		if !r.IsAvailable() {
			return nil, fmt.Errorf("no container runtime found (docker or podman)")
		}
	}

	cmd := exec.CommandContext(context.Background(), r.binary, "inspect", image)
	output, err := runCommand(cmd, verbose)
	if err != nil {
		return nil, err
	}

	return parseRuntimeInspect(output, image, r.binary)
}

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
func (r *ManifestRunner) Run(image string, verbose bool) (*types.ImageStats, error) {
	if r.binary == "" {
		if !r.IsAvailable() {
			return nil, fmt.Errorf("no container runtime found")
		}
	}

	// Try standard manifest inspect
	// Need DOCKER_CLI_EXPERIMENTAL=enabled for older docker to be safe
	cmd := exec.CommandContext(context.Background(), r.binary, "manifest", "inspect", image)
	cmd.Env = append(os.Environ(), "DOCKER_CLI_EXPERIMENTAL=enabled")

	output, err := runCommand(cmd, verbose)
	if err != nil {
		// Fallback or just return empty stats (optional feature)
		// We'll return error so analyzer can log warning
		return nil, fmt.Errorf("manifest inspect failed: %w", err)
	}

	return parseManifestInspect(output, image)
}

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
func (r *SyftRunner) Run(image string, verbose bool) (*types.ImageStats, error) {
	if r.binary == "" {
		if !r.IsAvailable() {
			return nil, fmt.Errorf("syft not found")
		}
	}
	cmd := exec.CommandContext(context.Background(), r.binary, image, "-o", "json")
	output, err := runCommand(cmd, verbose)
	if err != nil {
		return nil, err
	}

	return parseSyftOutput(output)
}

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
func (r *GrypeRunner) Run(image string, verbose bool) (*types.ImageStats, error) {
	if r.binary == "" {
		if !r.IsAvailable() {
			return nil, fmt.Errorf("grype not found")
		}
	}
	cmd := exec.CommandContext(context.Background(), r.binary, image, "-o", "json")
	output, err := runCommand(cmd, verbose)
	if err != nil {
		return nil, err
	}

	return parseGrypeOutput(output, verbose)
}

// DiveRunner runs 'dive <image> --json output.json'
type DiveRunner struct {
	binary string
}

// Name returns the display name for this runner.
func (r *DiveRunner) Name() string { return "dive" }

// IsAvailable checks whether the dive binary is installed.
func (r *DiveRunner) IsAvailable() bool {
	if path, err := lookupTool("dive"); err == nil {
		r.binary = path
		return true
	}
	return false
}

// detectPodmanSocket attempts to detect the Podman machine socket path.
// It returns a DOCKER_HOST value (e.g. "unix:///path/to/socket") if found,
// or an empty string if detection fails.
func detectPodmanSocket() string {
	pCmd := exec.CommandContext(context.Background(), "podman", "machine", "inspect")
	out, err := pCmd.Output()
	if err != nil {
		return ""
	}

	var machines []struct {
		ConnectionInfo struct {
			PodmanSocket *struct {
				Path string `json:"Path"`
			} `json:"PodmanSocket"`
		} `json:"ConnectionInfo"`
	}
	if json.Unmarshal(out, &machines) != nil || len(machines) == 0 {
		return ""
	}

	if machines[0].ConnectionInfo.PodmanSocket == nil || machines[0].ConnectionInfo.PodmanSocket.Path == "" {
		return ""
	}

	socketPath := machines[0].ConnectionInfo.PodmanSocket.Path
	if !strings.HasPrefix(socketPath, "unix://") {
		socketPath = "unix://" + socketPath
	}
	return socketPath
}

// Run executes dive against the given image and parses the efficiency results.
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

	if r.binary == "" {
		if !r.IsAvailable() {
			return nil, fmt.Errorf("dive not found")
		}
	}

	cmd := exec.CommandContext(context.Background(), r.binary, image, "--json", tmpFile.Name())

	// Podman Support: If docker is missing but podman is present and DOCKER_HOST
	// is not set, try to detect the Podman machine socket automatically.
	if _, err := exec.LookPath("docker"); err != nil {
		if _, err := exec.LookPath("podman"); err == nil && os.Getenv("DOCKER_HOST") == "" {
			if socketPath := detectPodmanSocket(); socketPath != "" {
				env := os.Environ()
				env = append(env, "DOCKER_HOST="+socketPath)
				cmd.Env = env
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
		return nil, fmt.Errorf("dive failed: %w, output: %s", err, string(output))
	}

	if verbose {
		fmt.Fprintf(os.Stderr, "[DEBUG] Output (%d bytes):\n%s\n", len(output), string(output))
	}

	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to read dive output: %w", err)
	}

	return parseDiveOutput(content)
}

// parseRuntimeInspect parses JSON output from 'docker inspect' or 'podman inspect'
// into ImageStats containing architecture, OS, size, and layer count.
func parseRuntimeInspect(output []byte, image string, binary string) (*types.ImageStats, error) {
	var inspect []struct {
		Architecture string `json:"Architecture"`
		Os           string `json:"Os"`
		Size         int64  `json:"Size"`
		RootFS       struct {
			Layers []string `json:"Layers"`
		} `json:"RootFS"`
	}

	if err := json.Unmarshal(output, &inspect); err != nil {
		return nil, fmt.Errorf("failed to unmarshal %s inspect output: %w", binary, err)
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

// parseManifestInspect parses JSON output from 'docker manifest inspect'
// into ImageStats containing supported architectures for multi-arch images.
func parseManifestInspect(output []byte, image string) (*types.ImageStats, error) {
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
			fmt.Fprintf(os.Stderr, "[DEBUG] Failed to parse grype timestamp: %v\n", err)
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

	severityRank := map[string]int{
		"Critical": 4,
		"High":     3,
		"Medium":   2,
		"Low":      1,
		"Unknown":  0,
	}

	sort.Slice(stats.Vulnerabilities, func(i, j int) bool {
		rankI := severityRank[stats.Vulnerabilities[i].Severity]
		rankJ := severityRank[stats.Vulnerabilities[j].Severity]
		if rankI != rankJ {
			return rankI > rankJ
		}
		return stats.Vulnerabilities[i].ID < stats.Vulnerabilities[j].ID
	})

	return stats, nil
}

// parseDiveOutput parses JSON output from dive's analysis file
// into ImageStats containing efficiency score and wasted bytes.
func parseDiveOutput(content []byte) (*types.ImageStats, error) {
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
		Efficiency:  diveOutput.Image.EfficiencyScore * 100,
		WastedBytes: fmt.Sprintf("%.2f MB", float64(diveOutput.Image.InefficientBytes)/1024/1024),
	}

	return stats, nil
}
