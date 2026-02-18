package runner

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// fakeExecHelper is a test helper function used by TestHelperProcess.
// When invoked as a subprocess by exec.Command, it reads the
// GO_TEST_HELPER_CMD env var to decide what to output, simulating
// external tools like docker, syft, grype, dive, etc.
func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	cmd := os.Getenv("GO_TEST_HELPER_CMD")
	switch cmd {
	case "inspect-ok":
		fmt.Print(`[{"Architecture":"amd64","Os":"linux","Size":10485760,"RootFS":{"Layers":["a","b"]}}]`)
	case "inspect-fail":
		os.Exit(1)
	case "manifest-ok":
		fmt.Print(`{"manifests":[{"platform":{"architecture":"amd64","os":"linux"}}]}`)
	case "syft-ok":
		fmt.Print(`{"distro":{"name":"alpine","version":"3.18"},"artifacts":[{"name":"musl","version":"1.2","type":"apk"}]}`)
	case "grype-ok":
		fmt.Print(`{"descriptor":{"timestamp":"2024-01-01T00:00:00Z"},"matches":[{"vulnerability":{"id":"CVE-1","severity":"High"},"artifact":{"name":"pkg","version":"1.0"}}]}`)
	case "dive-ok":
		// Write JSON to the file path specified in args
		args := os.Args
		// Find --json flag and get next arg
		for i, a := range args {
			if a == "--json" && i+1 < len(args) {
				data := `{"image":{"inefficientBytes":1024,"efficiencyScore":0.98}}`
				if err := os.WriteFile(args[i+1], []byte(data), 0644); err != nil {
					fmt.Fprintf(os.Stderr, "failed to write: %v", err)
					os.Exit(1)
				}
				break
			}
		}
	case "pull-ok":
		fmt.Print("pulled successfully")
	case "echo":
		fmt.Print("ok")
	case "error":
		fmt.Fprint(os.Stderr, "something went wrong")
		os.Exit(1)
	default:
		fmt.Fprintf(os.Stderr, "unknown helper cmd: %s\n", cmd)
		os.Exit(2)
	}
	os.Exit(0)
}

// helperCmd builds an exec.Cmd that re-invokes the test binary as a
// subprocess, setting GO_WANT_HELPER_PROCESS=1 so TestHelperProcess
// runs the scenario identified by helperCmd.
func helperCmd(helperType string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--"}
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = append(os.Environ(),
		"GO_WANT_HELPER_PROCESS=1",
		"GO_TEST_HELPER_CMD="+helperType,
	)
	return cmd
}

// createFakeBinary creates a shell script in dir that re-execs the
// test binary as a helper process. Returns the path to the script.
func createFakeBinary(t *testing.T, dir, name, helperType string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	script := fmt.Sprintf("#!/bin/sh\nexec %q -test.run=TestHelperProcess -- \"$@\"\n", os.Args[0])
	if err := os.WriteFile(path, []byte(script), 0755); err != nil {
		t.Fatalf("failed to create fake binary %s: %v", name, err)
	}
	// Set env vars so the subprocess knows what to do
	t.Setenv("GO_WANT_HELPER_PROCESS", "1")
	t.Setenv("GO_TEST_HELPER_CMD", helperType)
	return path
}

// TestRunCommand_Success tests runCommand with a successful command.
func TestRunCommand_Success(t *testing.T) {
	cmd := helperCmd("echo")
	output, err := runCommand(cmd, false)
	if err != nil {
		t.Fatalf("runCommand() unexpected error: %v", err)
	}
	if string(output) != "ok" {
		t.Errorf("runCommand() output = %q, want %q", string(output), "ok")
	}
}

// TestRunCommand_Verbose tests runCommand with verbose output.
func TestRunCommand_Verbose(t *testing.T) {
	cmd := helperCmd("echo")
	output, err := runCommand(cmd, true)
	if err != nil {
		t.Fatalf("runCommand() unexpected error: %v", err)
	}
	if string(output) != "ok" {
		t.Errorf("runCommand() output = %q, want %q", string(output), "ok")
	}
}

// TestRunCommand_Failure tests runCommand with a failing command.
func TestRunCommand_Failure(t *testing.T) {
	cmd := helperCmd("error")
	_, err := runCommand(cmd, false)
	if err == nil {
		t.Fatal("runCommand() expected error, got nil")
	}
	if !strings.Contains(err.Error(), "command failed") {
		t.Errorf("error should contain 'command failed', got: %v", err)
	}
	if !strings.Contains(err.Error(), "something went wrong") {
		t.Errorf("error should contain stderr output, got: %v", err)
	}
}

// TestRunCommand_FailureVerbose tests runCommand error path with verbose.
func TestRunCommand_FailureVerbose(t *testing.T) {
	cmd := helperCmd("error")
	_, err := runCommand(cmd, true)
	if err == nil {
		t.Fatal("runCommand() expected error, got nil")
	}
}

// TestRuntimeRunner_IsAvailable tests IsAvailable with system docker/podman.
func TestRuntimeRunner_IsAvailable(t *testing.T) {
	r := &RuntimeRunner{}
	result := r.IsAvailable()
	// On any dev machine with docker or podman, this should be true.
	// We just verify it doesn't panic and sets binary appropriately.
	if result {
		if r.binary != "docker" && r.binary != "podman" {
			t.Errorf("IsAvailable() set binary to %q, expected docker or podman", r.binary)
		}
	}
	// If neither is installed, result is false - that's fine too.
}

// TestRuntimeRunner_Run tests RuntimeRunner.Run with a mock binary.
func TestRuntimeRunner_Run(t *testing.T) {
	dir := t.TempDir()
	fakeBin := createFakeBinary(t, dir, "docker", "inspect-ok")

	r := &RuntimeRunner{binary: fakeBin}
	stats, err := r.Run("test:latest", false)
	if err != nil {
		t.Fatalf("RuntimeRunner.Run() error: %v", err)
	}
	if stats.Architecture != "amd64" {
		t.Errorf("Architecture = %q, want %q", stats.Architecture, "amd64")
	}
	if stats.OS != "linux" {
		t.Errorf("OS = %q, want %q", stats.OS, "linux")
	}
	if stats.TotalLayers != 2 {
		t.Errorf("TotalLayers = %d, want %d", stats.TotalLayers, 2)
	}
}

// TestRuntimeRunner_Run_NoBinary tests RuntimeRunner.Run when binary is empty and IsAvailable fails.
func TestRuntimeRunner_Run_NoBinary(t *testing.T) {
	// Override PATH to ensure neither docker nor podman is found
	t.Setenv("PATH", t.TempDir())
	r := &RuntimeRunner{}
	_, err := r.Run("test:latest", false)
	if err == nil {
		t.Fatal("RuntimeRunner.Run() expected error when no binary, got nil")
	}
	if !strings.Contains(err.Error(), "no container runtime found") {
		t.Errorf("unexpected error: %v", err)
	}
}

// TestManifestRunner_IsAvailable tests ManifestRunner.IsAvailable.
func TestManifestRunner_IsAvailable(t *testing.T) {
	r := &ManifestRunner{}
	result := r.IsAvailable()
	if result {
		if r.binary != "docker" && r.binary != "podman" {
			t.Errorf("IsAvailable() set binary to %q, expected docker or podman", r.binary)
		}
	}
}

// TestManifestRunner_Run tests ManifestRunner.Run with a mock binary.
func TestManifestRunner_Run(t *testing.T) {
	dir := t.TempDir()
	fakeBin := createFakeBinary(t, dir, "docker", "manifest-ok")

	r := &ManifestRunner{binary: fakeBin}
	stats, err := r.Run("test:latest", false)
	if err != nil {
		t.Fatalf("ManifestRunner.Run() error: %v", err)
	}
	if len(stats.SupportedArchitectures) != 1 {
		t.Errorf("SupportedArchitectures count = %d, want 1", len(stats.SupportedArchitectures))
	}
}

// TestManifestRunner_Run_NoBinary tests ManifestRunner.Run without binary.
func TestManifestRunner_Run_NoBinary(t *testing.T) {
	t.Setenv("PATH", t.TempDir())
	r := &ManifestRunner{}
	_, err := r.Run("test:latest", false)
	if err == nil {
		t.Fatal("ManifestRunner.Run() expected error when no binary, got nil")
	}
}

// TestSyftRunner_IsAvailable tests SyftRunner.IsAvailable.
func TestSyftRunner_IsAvailable(t *testing.T) {
	r := &SyftRunner{}
	// Save and restore original
	origLookup := lookupTool
	defer func() { lookupTool = origLookup }()

	// Mock: tool found
	lookupTool = func(name string) (string, error) {
		if name == "syft" {
			return "/usr/local/bin/syft", nil
		}
		return "", fmt.Errorf("not found")
	}
	if !r.IsAvailable() {
		t.Error("SyftRunner.IsAvailable() = false, want true")
	}
	if r.binary != "/usr/local/bin/syft" {
		t.Errorf("binary = %q, want /usr/local/bin/syft", r.binary)
	}

	// Mock: tool not found
	r2 := &SyftRunner{}
	lookupTool = func(name string) (string, error) {
		return "", fmt.Errorf("not found")
	}
	if r2.IsAvailable() {
		t.Error("SyftRunner.IsAvailable() = true, want false")
	}
}

// TestSyftRunner_Run tests SyftRunner.Run with a mock binary.
func TestSyftRunner_Run(t *testing.T) {
	dir := t.TempDir()
	fakeBin := createFakeBinary(t, dir, "syft", "syft-ok")

	r := &SyftRunner{binary: fakeBin}
	stats, err := r.Run("test:latest", false)
	if err != nil {
		t.Fatalf("SyftRunner.Run() error: %v", err)
	}
	if stats.OSDistro != "alpine 3.18" {
		t.Errorf("OSDistro = %q, want %q", stats.OSDistro, "alpine 3.18")
	}
	if stats.TotalPackages != 1 {
		t.Errorf("TotalPackages = %d, want 1", stats.TotalPackages)
	}
}

// TestSyftRunner_Run_NoBinary tests SyftRunner.Run without binary.
func TestSyftRunner_Run_NoBinary(t *testing.T) {
	origLookup := lookupTool
	defer func() { lookupTool = origLookup }()
	lookupTool = func(name string) (string, error) {
		return "", fmt.Errorf("not found")
	}
	r := &SyftRunner{}
	_, err := r.Run("test:latest", false)
	if err == nil {
		t.Fatal("SyftRunner.Run() expected error, got nil")
	}
	if !strings.Contains(err.Error(), "syft not found") {
		t.Errorf("unexpected error: %v", err)
	}
}

// TestGrypeRunner_IsAvailable tests GrypeRunner.IsAvailable.
func TestGrypeRunner_IsAvailable(t *testing.T) {
	origLookup := lookupTool
	defer func() { lookupTool = origLookup }()

	lookupTool = func(name string) (string, error) {
		if name == "grype" {
			return "/usr/local/bin/grype", nil
		}
		return "", fmt.Errorf("not found")
	}
	r := &GrypeRunner{}
	if !r.IsAvailable() {
		t.Error("GrypeRunner.IsAvailable() = false, want true")
	}

	lookupTool = func(name string) (string, error) {
		return "", fmt.Errorf("not found")
	}
	r2 := &GrypeRunner{}
	if r2.IsAvailable() {
		t.Error("GrypeRunner.IsAvailable() = true, want false")
	}
}

// TestGrypeRunner_Run tests GrypeRunner.Run with a mock binary.
func TestGrypeRunner_Run(t *testing.T) {
	dir := t.TempDir()
	fakeBin := createFakeBinary(t, dir, "grype", "grype-ok")

	r := &GrypeRunner{binary: fakeBin}
	stats, err := r.Run("test:latest", false)
	if err != nil {
		t.Fatalf("GrypeRunner.Run() error: %v", err)
	}
	if stats.VulnSummary["High"] != 1 {
		t.Errorf("VulnSummary[High] = %d, want 1", stats.VulnSummary["High"])
	}
	if len(stats.Vulnerabilities) != 1 {
		t.Errorf("Vulnerabilities count = %d, want 1", len(stats.Vulnerabilities))
	}
}

// TestGrypeRunner_Run_NoBinary tests GrypeRunner.Run without binary.
func TestGrypeRunner_Run_NoBinary(t *testing.T) {
	origLookup := lookupTool
	defer func() { lookupTool = origLookup }()
	lookupTool = func(name string) (string, error) {
		return "", fmt.Errorf("not found")
	}
	r := &GrypeRunner{}
	_, err := r.Run("test:latest", false)
	if err == nil {
		t.Fatal("GrypeRunner.Run() expected error, got nil")
	}
	if !strings.Contains(err.Error(), "grype not found") {
		t.Errorf("unexpected error: %v", err)
	}
}

// TestDiveRunner_IsAvailable tests DiveRunner.IsAvailable.
func TestDiveRunner_IsAvailable(t *testing.T) {
	origLookup := lookupTool
	defer func() { lookupTool = origLookup }()

	lookupTool = func(name string) (string, error) {
		if name == "dive" {
			return "/usr/local/bin/dive", nil
		}
		return "", fmt.Errorf("not found")
	}
	r := &DiveRunner{}
	if !r.IsAvailable() {
		t.Error("DiveRunner.IsAvailable() = false, want true")
	}

	lookupTool = func(name string) (string, error) {
		return "", fmt.Errorf("not found")
	}
	r2 := &DiveRunner{}
	if r2.IsAvailable() {
		t.Error("DiveRunner.IsAvailable() = true, want false")
	}
}

// TestDiveRunner_Run tests DiveRunner.Run with a mock binary.
func TestDiveRunner_Run(t *testing.T) {
	dir := t.TempDir()
	fakeBin := createFakeBinary(t, dir, "dive", "dive-ok")

	r := &DiveRunner{binary: fakeBin}
	stats, err := r.Run("test:latest", false)
	if err != nil {
		t.Fatalf("DiveRunner.Run() error: %v", err)
	}
	if stats.Efficiency != 98.0 {
		t.Errorf("Efficiency = %v, want 98.0", stats.Efficiency)
	}
	if stats.WastedBytes != "0.00 MB" {
		t.Errorf("WastedBytes = %q, want %q", stats.WastedBytes, "0.00 MB")
	}
}

// TestDiveRunner_Run_Verbose tests DiveRunner.Run in verbose mode.
func TestDiveRunner_Run_Verbose(t *testing.T) {
	dir := t.TempDir()
	fakeBin := createFakeBinary(t, dir, "dive", "dive-ok")

	r := &DiveRunner{binary: fakeBin}
	stats, err := r.Run("test:latest", true)
	if err != nil {
		t.Fatalf("DiveRunner.Run() error: %v", err)
	}
	if stats.Efficiency != 98.0 {
		t.Errorf("Efficiency = %v, want 98.0", stats.Efficiency)
	}
}

// TestDiveRunner_Run_NoBinary tests DiveRunner.Run without binary.
func TestDiveRunner_Run_NoBinary(t *testing.T) {
	origLookup := lookupTool
	defer func() { lookupTool = origLookup }()
	lookupTool = func(name string) (string, error) {
		return "", fmt.Errorf("not found")
	}
	r := &DiveRunner{}
	_, err := r.Run("test:latest", false)
	if err == nil {
		t.Fatal("DiveRunner.Run() expected error, got nil")
	}
	if !strings.Contains(err.Error(), "dive not found") {
		t.Errorf("unexpected error: %v", err)
	}
}

// TestEnsureImage_AlreadyExists tests EnsureImage when image is already local.
func TestEnsureImage_AlreadyExists(t *testing.T) {
	// Create a fake docker/podman that succeeds on inspect
	dir := t.TempDir()
	fakeBin := filepath.Join(dir, "docker")
	script := fmt.Sprintf("#!/bin/sh\nexec %q -test.run=TestHelperProcess -- \"$@\"\n", os.Args[0])
	if err := os.WriteFile(fakeBin, []byte(script), 0755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("GO_WANT_HELPER_PROCESS", "1")
	t.Setenv("GO_TEST_HELPER_CMD", "echo")
	t.Setenv("PATH", dir+":"+os.Getenv("PATH"))

	err := EnsureImage("test:latest", false)
	if err != nil {
		t.Fatalf("EnsureImage() unexpected error: %v", err)
	}
}

// TestEnsureImage_Verbose tests EnsureImage with verbose flag when image exists.
func TestEnsureImage_Verbose(t *testing.T) {
	dir := t.TempDir()
	fakeBin := filepath.Join(dir, "docker")
	script := fmt.Sprintf("#!/bin/sh\nexec %q -test.run=TestHelperProcess -- \"$@\"\n", os.Args[0])
	if err := os.WriteFile(fakeBin, []byte(script), 0755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("GO_WANT_HELPER_PROCESS", "1")
	t.Setenv("GO_TEST_HELPER_CMD", "echo")
	t.Setenv("PATH", dir+":"+os.Getenv("PATH"))

	err := EnsureImage("test:latest", true)
	if err != nil {
		t.Fatalf("EnsureImage() unexpected error: %v", err)
	}
}

// TestEnsureImage_NoRuntime tests EnsureImage when no docker/podman is available.
func TestEnsureImage_NoRuntime(t *testing.T) {
	t.Setenv("PATH", t.TempDir())
	err := EnsureImage("test:latest", false)
	if err == nil {
		t.Fatal("EnsureImage() expected error, got nil")
	}
	if !strings.Contains(err.Error(), "no container runtime found") {
		t.Errorf("unexpected error: %v", err)
	}
}

// TestDetectPodmanSocket tests the detectPodmanSocket helper function.
func TestDetectPodmanSocket(t *testing.T) {
	// When podman is not in PATH, should return empty string.
	t.Setenv("PATH", t.TempDir())
	result := detectPodmanSocket()
	if result != "" {
		t.Errorf("detectPodmanSocket() = %q, want empty string", result)
	}
}

// TestLookupTool_Swappable verifies that lookupTool is a swappable function variable.
func TestLookupTool_Swappable(t *testing.T) {
	orig := lookupTool
	defer func() { lookupTool = orig }()

	lookupTool = func(name string) (string, error) {
		return "/mock/path/" + name, nil
	}

	path, err := lookupTool("test-tool")
	if err != nil {
		t.Fatalf("lookupTool() unexpected error: %v", err)
	}
	if path != "/mock/path/test-tool" {
		t.Errorf("lookupTool() = %q, want /mock/path/test-tool", path)
	}
}

func TestRuntimeRunner_Name(t *testing.T) {
	tests := []struct {
		name     string
		binary   string
		expected string
	}{
		{
			name:     "default name",
			binary:   "",
			expected: "runtime",
		},
		{
			name:     "docker binary",
			binary:   "docker",
			expected: "docker",
		},
		{
			name:     "podman binary",
			binary:   "podman",
			expected: "podman",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RuntimeRunner{binary: tt.binary}
			if got := r.Name(); got != tt.expected {
				t.Errorf("RuntimeRunner.Name() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestParseRuntimeInspect(t *testing.T) {
	tests := []struct {
		name       string
		jsonOutput string
		image      string
		binary     string
		wantErr    bool
		wantArch   string
		wantOS     string
		wantLayers int
		wantSizeMB string
	}{
		{
			name: "valid inspect output",
			jsonOutput: `[{
				"Architecture": "amd64",
				"Os": "linux",
				"Size": 67108864,
				"RootFS": {
					"Layers": ["layer1", "layer2", "layer3"]
				}
			}]`,
			image:      "nginx:latest",
			binary:     "docker",
			wantErr:    false,
			wantArch:   "amd64",
			wantOS:     "linux",
			wantLayers: 3,
			wantSizeMB: "64.00 MB",
		},
		{
			name: "arm64 image",
			jsonOutput: `[{
				"Architecture": "arm64",
				"Os": "linux",
				"Size": 52428800,
				"RootFS": {
					"Layers": ["sha256:abc", "sha256:def"]
				}
			}]`,
			image:      "alpine:3.18",
			binary:     "podman",
			wantErr:    false,
			wantArch:   "arm64",
			wantOS:     "linux",
			wantLayers: 2,
			wantSizeMB: "50.00 MB",
		},
		{
			name:       "empty array",
			jsonOutput: `[]`,
			image:      "test:latest",
			binary:     "docker",
			wantErr:    true,
		},
		{
			name:       "invalid json",
			jsonOutput: `{invalid}`,
			image:      "test:latest",
			binary:     "docker",
			wantErr:    true,
		},
		{
			name:       "zero-size image with no layers",
			jsonOutput: `[{"Architecture": "amd64", "Os": "linux", "Size": 0, "RootFS": {"Layers": []}}]`,
			image:      "scratch",
			binary:     "docker",
			wantErr:    false,
			wantArch:   "amd64",
			wantOS:     "linux",
			wantLayers: 0,
			wantSizeMB: "0.00 MB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stats, err := parseRuntimeInspect([]byte(tt.jsonOutput), tt.image, tt.binary)
			if tt.wantErr {
				if err == nil {
					t.Errorf("parseRuntimeInspect() expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("parseRuntimeInspect() unexpected error: %v", err)
			}
			if stats.Architecture != tt.wantArch {
				t.Errorf("Architecture = %v, want %v", stats.Architecture, tt.wantArch)
			}
			if stats.OS != tt.wantOS {
				t.Errorf("OS = %v, want %v", stats.OS, tt.wantOS)
			}
			if stats.TotalLayers != tt.wantLayers {
				t.Errorf("TotalLayers = %v, want %v", stats.TotalLayers, tt.wantLayers)
			}
			if stats.SizeMB != tt.wantSizeMB {
				t.Errorf("SizeMB = %v, want %v", stats.SizeMB, tt.wantSizeMB)
			}
			if stats.ImageTag != tt.image {
				t.Errorf("ImageTag = %v, want %v", stats.ImageTag, tt.image)
			}
		})
	}
}

func TestManifestRunner_Name(t *testing.T) {
	r := &ManifestRunner{}
	if got := r.Name(); got != "manifest" {
		t.Errorf("ManifestRunner.Name() = %v, want %v", got, "manifest")
	}
}

func TestParseManifestInspect(t *testing.T) {
	tests := []struct {
		name      string
		json      string
		image     string
		wantArchs []string
		wantErr   bool
	}{
		{
			name: "manifest list with duplicates",
			json: `{
				"manifests": [
					{"platform": {"architecture": "amd64", "os": "linux"}},
					{"platform": {"architecture": "arm64", "os": "linux"}},
					{"platform": {"architecture": "arm", "os": "linux"}},
					{"platform": {"architecture": "amd64", "os": "linux"}}
				]
			}`,
			image:     "nginx:latest",
			wantArchs: []string{"linux/amd64", "linux/arm", "linux/arm64"},
		},
		{
			name: "single platform",
			json: `{
				"manifests": [
					{"platform": {"architecture": "amd64", "os": "linux"}}
				]
			}`,
			image:     "myapp:v1",
			wantArchs: []string{"linux/amd64"},
		},
		{
			name: "multi-os manifest",
			json: `{
				"manifests": [
					{"platform": {"architecture": "amd64", "os": "linux"}},
					{"platform": {"architecture": "amd64", "os": "windows"}}
				]
			}`,
			image:     "dotnet:latest",
			wantArchs: []string{"linux/amd64", "windows/amd64"},
		},
		{
			name:      "empty manifests list",
			json:      `{"manifests": []}`,
			image:     "test:latest",
			wantArchs: nil, // falls through to empty stats
		},
		{
			name:      "not a manifest list (single image V2 schema)",
			json:      `{"schemaVersion": 2, "mediaType": "application/vnd.docker.distribution.manifest.v2+json"}`,
			image:     "custom:latest",
			wantArchs: nil,
		},
		{
			name:      "invalid json returns empty stats",
			json:      `{totally broken`,
			image:     "bad:latest",
			wantArchs: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stats, err := parseManifestInspect([]byte(tt.json), tt.image)
			if tt.wantErr {
				if err == nil {
					t.Errorf("parseManifestInspect() expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("parseManifestInspect() unexpected error: %v", err)
			}
			if stats.ImageTag != tt.image {
				t.Errorf("ImageTag = %v, want %v", stats.ImageTag, tt.image)
			}
			if tt.wantArchs == nil {
				if len(stats.SupportedArchitectures) != 0 {
					t.Errorf("SupportedArchitectures = %v, want empty", stats.SupportedArchitectures)
				}
				return
			}
			if len(stats.SupportedArchitectures) != len(tt.wantArchs) {
				t.Fatalf("SupportedArchitectures count = %d, want %d: %v",
					len(stats.SupportedArchitectures), len(tt.wantArchs), stats.SupportedArchitectures)
			}
			for i, arch := range stats.SupportedArchitectures {
				if arch != tt.wantArchs[i] {
					t.Errorf("SupportedArchitectures[%d] = %v, want %v", i, arch, tt.wantArchs[i])
				}
			}
		})
	}
}

func TestSyftRunner_Name(t *testing.T) {
	r := &SyftRunner{}
	if got := r.Name(); got != "syft" {
		t.Errorf("SyftRunner.Name() = %v, want %v", got, "syft")
	}
}

func TestParseSyftOutput(t *testing.T) {
	tests := []struct {
		name         string
		json         string
		wantErr      bool
		wantDistro   string
		wantPkgCount int
		wantFirstPkg string
		wantLastPkg  string
	}{
		{
			name: "standard output with duplicates",
			json: `{
				"distro": {"name": "ubuntu", "version": "22.04"},
				"artifacts": [
					{"name": "openssl", "version": "3.0.2", "type": "deb"},
					{"name": "curl", "version": "7.81.0", "type": "deb"},
					{"name": "openssl", "version": "3.0.2", "type": "deb"}
				]
			}`,
			wantDistro:   "ubuntu 22.04",
			wantPkgCount: 2,
			wantFirstPkg: "curl",
			wantLastPkg:  "openssl",
		},
		{
			name: "distro without version",
			json: `{
				"distro": {"name": "alpine", "version": ""},
				"artifacts": [
					{"name": "musl", "version": "1.2.4", "type": "apk"}
				]
			}`,
			wantDistro:   "alpine",
			wantPkgCount: 1,
			wantFirstPkg: "musl",
			wantLastPkg:  "musl",
		},
		{
			name: "no distro info",
			json: `{
				"distro": {"name": "", "version": ""},
				"artifacts": [
					{"name": "zlib", "version": "1.0", "type": "binary"}
				]
			}`,
			wantDistro:   "",
			wantPkgCount: 1,
			wantFirstPkg: "zlib",
			wantLastPkg:  "zlib",
		},
		{
			name:         "empty artifacts",
			json:         `{"distro": {"name": "", "version": ""}, "artifacts": []}`,
			wantDistro:   "",
			wantPkgCount: 0,
		},
		{
			name:    "invalid json",
			json:    `not json`,
			wantErr: true,
		},
		{
			name: "packages sorted alphabetically",
			json: `{
				"distro": {"name": "debian", "version": "12"},
				"artifacts": [
					{"name": "zlib", "version": "1.2", "type": "deb"},
					{"name": "apt", "version": "2.6", "type": "deb"},
					{"name": "bash", "version": "5.2", "type": "deb"}
				]
			}`,
			wantDistro:   "debian 12",
			wantPkgCount: 3,
			wantFirstPkg: "apt",
			wantLastPkg:  "zlib",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stats, err := parseSyftOutput([]byte(tt.json))
			if tt.wantErr {
				if err == nil {
					t.Errorf("parseSyftOutput() expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("parseSyftOutput() unexpected error: %v", err)
			}
			if stats.OSDistro != tt.wantDistro {
				t.Errorf("OSDistro = %q, want %q", stats.OSDistro, tt.wantDistro)
			}
			if stats.TotalPackages != tt.wantPkgCount {
				t.Errorf("TotalPackages = %d, want %d", stats.TotalPackages, tt.wantPkgCount)
			}
			if len(stats.Packages) != tt.wantPkgCount {
				t.Errorf("len(Packages) = %d, want %d", len(stats.Packages), tt.wantPkgCount)
			}
			if tt.wantPkgCount > 0 {
				if stats.Packages[0].Name != tt.wantFirstPkg {
					t.Errorf("first package = %q, want %q", stats.Packages[0].Name, tt.wantFirstPkg)
				}
				if stats.Packages[len(stats.Packages)-1].Name != tt.wantLastPkg {
					t.Errorf("last package = %q, want %q", stats.Packages[len(stats.Packages)-1].Name, tt.wantLastPkg)
				}
			}
		})
	}
}

func TestGrypeRunner_Name(t *testing.T) {
	r := &GrypeRunner{}
	if got := r.Name(); got != "grype" {
		t.Errorf("GrypeRunner.Name() = %v, want %v", got, "grype")
	}
}

func TestParseGrypeOutput(t *testing.T) {
	tests := []struct {
		name             string
		json             string
		verbose          bool
		wantErr          bool
		wantCritical     int
		wantHigh         int
		wantMedium       int
		wantLow          int
		wantTotalVulns   int
		wantFirstVulnSev string
		wantTimeParsed   bool
		wantTimestamp    string
	}{
		{
			name: "standard output with multiple severities",
			json: `{
				"descriptor": {"timestamp": "2024-02-15T14:30:00Z"},
				"matches": [
					{
						"vulnerability": {"id": "CVE-2023-0001", "severity": "Critical"},
						"artifact": {"name": "openssl", "version": "3.0.2"}
					},
					{
						"vulnerability": {"id": "CVE-2023-0002", "severity": "High"},
						"artifact": {"name": "curl", "version": "7.81.0"}
					},
					{
						"vulnerability": {"id": "CVE-2023-0003", "severity": "Critical"},
						"artifact": {"name": "libssl", "version": "3.0.2"}
					},
					{
						"vulnerability": {"id": "CVE-2023-0004", "severity": "Medium"},
						"artifact": {"name": "bash", "version": "5.1"}
					}
				]
			}`,
			wantCritical:     2,
			wantHigh:         1,
			wantMedium:       1,
			wantTotalVulns:   4,
			wantFirstVulnSev: "Critical",
			wantTimeParsed:   true,
			wantTimestamp:    "2024-02-15T14:30:00Z",
		},
		{
			name: "no vulnerabilities",
			json: `{
				"descriptor": {"timestamp": "2024-01-01T00:00:00Z"},
				"matches": []
			}`,
			wantTotalVulns: 0,
			wantTimeParsed: true,
			wantTimestamp:  "2024-01-01T00:00:00Z",
		},
		{
			name: "missing timestamp defaults to now",
			json: `{
				"descriptor": {},
				"matches": [
					{
						"vulnerability": {"id": "CVE-2024-0001", "severity": "Low"},
						"artifact": {"name": "libc", "version": "2.31"}
					}
				]
			}`,
			wantLow:          1,
			wantTotalVulns:   1,
			wantFirstVulnSev: "Low",
			wantTimeParsed:   false,
		},
		{
			name: "sort order: Critical before High before Medium before Low",
			json: `{
				"descriptor": {"timestamp": "2024-06-01T12:00:00Z"},
				"matches": [
					{
						"vulnerability": {"id": "CVE-LOW", "severity": "Low"},
						"artifact": {"name": "pkg-a", "version": "1.0"}
					},
					{
						"vulnerability": {"id": "CVE-CRIT", "severity": "Critical"},
						"artifact": {"name": "pkg-b", "version": "1.0"}
					},
					{
						"vulnerability": {"id": "CVE-MED", "severity": "Medium"},
						"artifact": {"name": "pkg-c", "version": "1.0"}
					},
					{
						"vulnerability": {"id": "CVE-HIGH", "severity": "High"},
						"artifact": {"name": "pkg-d", "version": "1.0"}
					}
				]
			}`,
			wantCritical:     1,
			wantHigh:         1,
			wantMedium:       1,
			wantLow:          1,
			wantTotalVulns:   4,
			wantFirstVulnSev: "Critical",
			wantTimeParsed:   true,
			wantTimestamp:    "2024-06-01T12:00:00Z",
		},
		{
			name:    "invalid json",
			json:    `broken`,
			wantErr: true,
		},
		{
			name: "invalid timestamp with verbose logs warning",
			json: `{
				"descriptor": {"timestamp": "not-a-timestamp"},
				"matches": []
			}`,
			verbose:        true,
			wantTotalVulns: 0,
			wantTimeParsed: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stats, err := parseGrypeOutput([]byte(tt.json), tt.verbose)
			if tt.wantErr {
				if err == nil {
					t.Errorf("parseGrypeOutput() expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("parseGrypeOutput() unexpected error: %v", err)
			}
			if len(stats.Vulnerabilities) != tt.wantTotalVulns {
				t.Errorf("total vulns = %d, want %d", len(stats.Vulnerabilities), tt.wantTotalVulns)
			}
			if stats.VulnSummary["Critical"] != tt.wantCritical {
				t.Errorf("Critical = %d, want %d", stats.VulnSummary["Critical"], tt.wantCritical)
			}
			if stats.VulnSummary["High"] != tt.wantHigh {
				t.Errorf("High = %d, want %d", stats.VulnSummary["High"], tt.wantHigh)
			}
			if stats.VulnSummary["Medium"] != tt.wantMedium {
				t.Errorf("Medium = %d, want %d", stats.VulnSummary["Medium"], tt.wantMedium)
			}
			if stats.VulnSummary["Low"] != tt.wantLow {
				t.Errorf("Low = %d, want %d", stats.VulnSummary["Low"], tt.wantLow)
			}
			if tt.wantTotalVulns > 0 && tt.wantFirstVulnSev != "" {
				if stats.Vulnerabilities[0].Severity != tt.wantFirstVulnSev {
					t.Errorf("first vuln severity = %q, want %q",
						stats.Vulnerabilities[0].Severity, tt.wantFirstVulnSev)
				}
			}
			if tt.wantTimeParsed {
				expected, _ := time.Parse(time.RFC3339, tt.wantTimestamp)
				if !stats.VulnScanTime.Equal(expected) {
					t.Errorf("VulnScanTime = %v, want %v", stats.VulnScanTime, expected)
				}
			}
		})
	}
}

func TestParseGrypeOutput_SortStability(t *testing.T) {
	// When two vulns have the same severity, they should be sorted by ID ascending
	json := `{
		"descriptor": {"timestamp": "2024-01-01T00:00:00Z"},
		"matches": [
			{
				"vulnerability": {"id": "CVE-2024-0003", "severity": "High"},
				"artifact": {"name": "pkg-c", "version": "1.0"}
			},
			{
				"vulnerability": {"id": "CVE-2024-0001", "severity": "High"},
				"artifact": {"name": "pkg-a", "version": "1.0"}
			},
			{
				"vulnerability": {"id": "CVE-2024-0002", "severity": "High"},
				"artifact": {"name": "pkg-b", "version": "1.0"}
			}
		]
	}`

	stats, err := parseGrypeOutput([]byte(json), false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedOrder := []string{"CVE-2024-0001", "CVE-2024-0002", "CVE-2024-0003"}
	for i, vuln := range stats.Vulnerabilities {
		if vuln.ID != expectedOrder[i] {
			t.Errorf("vuln[%d].ID = %q, want %q", i, vuln.ID, expectedOrder[i])
		}
	}
}

func TestDiveRunner_Name(t *testing.T) {
	r := &DiveRunner{}
	if got := r.Name(); got != "dive" {
		t.Errorf("DiveRunner.Name() = %v, want %v", got, "dive")
	}
}

func TestParseDiveOutput(t *testing.T) {
	tests := []struct {
		name           string
		json           string
		wantErr        bool
		wantEfficiency float64
		wantWasted     string
	}{
		{
			name: "standard output",
			json: `{
				"image": {
					"inefficientBytes": 10485760,
					"efficiencyScore": 0.95
				}
			}`,
			wantEfficiency: 95.0,
			wantWasted:     "10.00 MB",
		},
		{
			name: "perfect efficiency",
			json: `{
				"image": {
					"inefficientBytes": 0,
					"efficiencyScore": 1.0
				}
			}`,
			wantEfficiency: 100.0,
			wantWasted:     "0.00 MB",
		},
		{
			name: "low efficiency",
			json: `{
				"image": {
					"inefficientBytes": 104857600,
					"efficiencyScore": 0.42
				}
			}`,
			wantEfficiency: 42.0,
			wantWasted:     "100.00 MB",
		},
		{
			name:    "invalid json",
			json:    `{broken`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stats, err := parseDiveOutput([]byte(tt.json))
			if tt.wantErr {
				if err == nil {
					t.Errorf("parseDiveOutput() expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("parseDiveOutput() unexpected error: %v", err)
			}
			if stats.Efficiency != tt.wantEfficiency {
				t.Errorf("Efficiency = %v, want %v", stats.Efficiency, tt.wantEfficiency)
			}
			if stats.WastedBytes != tt.wantWasted {
				t.Errorf("WastedBytes = %v, want %v", stats.WastedBytes, tt.wantWasted)
			}
		})
	}
}

func TestPodmanSocketDetection(t *testing.T) {
	// Tests the JSON parsing logic used in DiveRunner.Run() for Podman socket detection.
	testJSON := `[{
		"ConnectionInfo": {
			"PodmanSocket": {
				"Path": "/var/run/podman.sock"
			}
		}
	}]`

	var machines []struct {
		ConnectionInfo struct {
			PodmanSocket *struct {
				Path string `json:"Path"`
			} `json:"PodmanSocket"`
		} `json:"ConnectionInfo"`
	}

	if err := json.Unmarshal([]byte(testJSON), &machines); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if len(machines) == 0 {
		t.Fatal("expected at least one machine")
	}

	if machines[0].ConnectionInfo.PodmanSocket == nil {
		t.Fatal("PodmanSocket is nil")
	}

	socketPath := machines[0].ConnectionInfo.PodmanSocket.Path
	if socketPath != "/var/run/podman.sock" {
		t.Errorf("socket path = %v, want /var/run/podman.sock", socketPath)
	}

	// Test unix:// prefix logic
	if !strings.HasPrefix(socketPath, "unix://") {
		socketPath = "unix://" + socketPath
	}

	expectedPath := "unix:///var/run/podman.sock"
	if socketPath != expectedPath {
		t.Errorf("formatted socket path = %v, want %v", socketPath, expectedPath)
	}
}

func TestTempFileCleanup(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test-dive-output-*.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	tmpPath := tmpFile.Name()
	if err := tmpFile.Close(); err != nil {
		t.Fatalf("failed to close temp file: %v", err)
	}

	if _, err := os.Stat(tmpPath); os.IsNotExist(err) {
		t.Fatal("temp file should exist after creation")
	}

	if err := os.Remove(tmpPath); err != nil {
		t.Fatalf("failed to remove temp file: %v", err)
	}

	if _, err := os.Stat(tmpPath); !os.IsNotExist(err) {
		t.Error("temp file should not exist after removal")
	}
}
