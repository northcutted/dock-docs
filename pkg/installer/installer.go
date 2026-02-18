// Package installer handles downloading and installing external tool
// dependencies (syft, grype, dive) from their upstream GitHub Releases.
package installer

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// Tool represents an external dependency that can be installed.
type Tool struct {
	Name string // binary name: "syft", "grype", "dive"
	Repo string // GitHub owner/repo: "anchore/syft", etc.
}

// Tools is the list of all external dependencies.
var Tools = []Tool{
	{Name: "syft", Repo: "anchore/syft"},
	{Name: "grype", Repo: "anchore/grype"},
	{Name: "dive", Repo: "wagoodman/dive"},
}

// ToolStatus describes the install state of a single tool.
type ToolStatus struct {
	Name      string
	Installed bool
	Path      string // where the binary was found (empty if missing)
	Source    string // "PATH", "dock-docs", or ""
}

// DefaultInstallDir returns ~/.dock-docs/bin, creating nothing.
func DefaultInstallDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot determine home directory: %w", err)
	}
	return filepath.Join(home, ".dock-docs", "bin"), nil
}

// FindTool locates a tool binary by checking PATH first, then the
// dock-docs install directory as a fallback. Returns the resolved
// path and a source label ("PATH" or "dock-docs"), or an error.
func FindTool(name string) (path string, source string, err error) {
	// System PATH first
	if p, err := exec.LookPath(name); err == nil {
		return p, "PATH", nil
	}
	// Fallback: dock-docs install directory
	installDir, err := DefaultInstallDir()
	if err != nil {
		return "", "", fmt.Errorf("%s not found in PATH: %w", name, err)
	}
	candidate := filepath.Join(installDir, name)
	if info, err := os.Stat(candidate); err == nil && !info.IsDir() {
		return candidate, "dock-docs", nil
	}
	return "", "", fmt.Errorf("%s not found in PATH or %s", name, installDir)
}

// Status checks whether each tool is installed and reachable.
func Status(installDir string) []ToolStatus {
	results := make([]ToolStatus, len(Tools))
	for i, t := range Tools {
		results[i] = ToolStatus{Name: t.Name}

		// Check PATH
		if p, err := exec.LookPath(t.Name); err == nil {
			results[i].Installed = true
			results[i].Path = p
			results[i].Source = "PATH"
			continue
		}

		// Check install directory
		candidate := filepath.Join(installDir, t.Name)
		if info, err := os.Stat(candidate); err == nil && !info.IsDir() {
			results[i].Installed = true
			results[i].Path = candidate
			results[i].Source = "dock-docs"
		}
	}
	return results
}

// httpClient is the HTTP client used for all requests. It can be
// replaced in tests to point at a mock server.
var httpClient = http.DefaultClient

// Install downloads and installs a single tool into installDir.
// It fetches the latest release tag from GitHub, downloads the
// platform-appropriate tarball, extracts the binary, and writes
// it to installDir/<name> with mode 0755.
func Install(tool Tool, installDir string) error {
	tag, err := latestReleaseTag(tool.Repo)
	if err != nil {
		return fmt.Errorf("failed to find latest release for %s: %w", tool.Name, err)
	}

	version := strings.TrimPrefix(tag, "v")
	assetName := fmt.Sprintf("%s_%s_%s_%s.tar.gz", tool.Name, version, runtime.GOOS, runtime.GOARCH)
	assetURL := fmt.Sprintf("https://github.com/%s/releases/download/%s/%s", tool.Repo, tag, assetName)

	fmt.Printf("  Downloading %s %s (%s/%s)...\n", tool.Name, tag, runtime.GOOS, runtime.GOARCH)

	// Ensure install directory exists
	if err := os.MkdirAll(installDir, 0755); err != nil {
		return fmt.Errorf("failed to create install directory %s: %w", installDir, err)
	}

	// Download the tarball
	body, err := httpGet(assetURL)
	if err != nil {
		return fmt.Errorf("failed to download %s: %w", assetName, err)
	}
	defer func() { _ = body.Close() }()

	// Extract the binary from the tarball
	binaryData, err := extractBinaryFromTarGz(body, tool.Name)
	if err != nil {
		return fmt.Errorf("failed to extract %s from archive: %w", tool.Name, err)
	}

	// Write to disk
	destPath := filepath.Join(installDir, tool.Name)
	if err := os.WriteFile(destPath, binaryData, 0755); err != nil {
		return fmt.Errorf("failed to write %s: %w", destPath, err)
	}

	// Quick verification: run "<tool> version" to check the binary works
	cmd := exec.CommandContext(context.Background(), destPath, "version")
	if err := cmd.Run(); err != nil {
		// Non-fatal: binary was written but version check failed
		fmt.Printf("  Warning: %s installed but version check failed: %v\n", tool.Name, err)
	}

	fmt.Printf("  Installed %s %s -> %s\n", tool.Name, tag, destPath)
	return nil
}

// InstallAll installs each tool that is missing (or all if force is true).
func InstallAll(installDir string, force bool) error {
	statuses := Status(installDir)
	var toInstall []Tool

	for i, s := range statuses {
		if force || !s.Installed {
			toInstall = append(toInstall, Tools[i])
		} else {
			fmt.Printf("  %s already installed (%s: %s)\n", s.Name, s.Source, s.Path)
		}
	}

	if len(toInstall) == 0 {
		fmt.Println("All tools are already installed.")
		return nil
	}

	var errs []string
	for _, t := range toInstall {
		if err := Install(t, installDir); err != nil {
			errs = append(errs, fmt.Sprintf("%s: %v", t.Name, err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("some tools failed to install:\n  %s", strings.Join(errs, "\n  "))
	}
	return nil
}

// latestReleaseTag queries the GitHub API for the latest release tag.
func latestReleaseTag(repo string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)

	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/vnd.github+json")

	// Use GITHUB_TOKEN if available (avoids rate limiting in CI)
	if token := os.Getenv("GITHUB_TOKEN"); token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GitHub API returned %d for %s", resp.StatusCode, repo)
	}

	var release struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", fmt.Errorf("failed to decode GitHub response: %w", err)
	}
	if release.TagName == "" {
		return "", fmt.Errorf("no tag_name in release response for %s", repo)
	}

	return release.TagName, nil
}

// httpGet performs a GET request and returns the response body.
// The caller must close the returned ReadCloser.
func httpGet(url string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return nil, err
	}

	// GitHub asset downloads may need token for private repos,
	// but these are all public. Include token anyway for rate limits.
	if token := os.Getenv("GITHUB_TOKEN"); token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		_ = resp.Body.Close()
		return nil, fmt.Errorf("HTTP %d for %s", resp.StatusCode, url)
	}

	return resp.Body, nil
}

// extractBinaryFromTarGz reads a .tar.gz stream and extracts the named
// binary. It returns the full file contents in memory.
func extractBinaryFromTarGz(r io.Reader, binaryName string) ([]byte, error) {
	gz, err := gzip.NewReader(r)
	if err != nil {
		return nil, fmt.Errorf("gzip error: %w", err)
	}
	defer func() { _ = gz.Close() }()

	tr := tar.NewReader(gz)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("tar error: %w", err)
		}

		// The binary can be at the root of the archive or in a subdirectory.
		// Match by base name.
		if filepath.Base(hdr.Name) == binaryName && hdr.Typeflag == tar.TypeReg {
			data, err := io.ReadAll(tr)
			if err != nil {
				return nil, fmt.Errorf("failed to read %s from archive: %w", binaryName, err)
			}
			return data, nil
		}
	}

	return nil, fmt.Errorf("binary %q not found in archive", binaryName)
}
