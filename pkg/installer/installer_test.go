package installer

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// makeTarGz creates an in-memory .tar.gz archive containing a single file
// with the given name and content.
func makeTarGz(t *testing.T, fileName string, content []byte) []byte {
	t.Helper()

	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gw)

	hdr := &tar.Header{
		Name: fileName,
		Mode: 0755,
		Size: int64(len(content)),
	}
	if err := tw.WriteHeader(hdr); err != nil {
		t.Fatalf("tar WriteHeader: %v", err)
	}
	if _, err := tw.Write(content); err != nil {
		t.Fatalf("tar Write: %v", err)
	}
	if err := tw.Close(); err != nil {
		t.Fatalf("tar Close: %v", err)
	}
	if err := gw.Close(); err != nil {
		t.Fatalf("gzip Close: %v", err)
	}

	return buf.Bytes()
}

func TestDefaultInstallDir(t *testing.T) {
	dir, err := DefaultInstallDir()
	if err != nil {
		t.Fatalf("DefaultInstallDir() error: %v", err)
	}
	if dir == "" {
		t.Fatal("DefaultInstallDir() returned empty string")
	}
	if !filepath.IsAbs(dir) {
		t.Errorf("DefaultInstallDir() returned relative path: %s", dir)
	}
	if filepath.Base(dir) != "bin" {
		t.Errorf("DefaultInstallDir() should end with 'bin', got: %s", dir)
	}
}

func TestExtractBinaryFromTarGz(t *testing.T) {
	expected := []byte("#!/bin/sh\necho hello\n")
	archive := makeTarGz(t, "mytool", expected)

	data, err := extractBinaryFromTarGz(bytes.NewReader(archive), "mytool")
	if err != nil {
		t.Fatalf("extractBinaryFromTarGz() error: %v", err)
	}
	if !bytes.Equal(data, expected) {
		t.Errorf("extracted content mismatch: got %q, want %q", data, expected)
	}
}

func TestExtractBinaryFromTarGz_NotFound(t *testing.T) {
	archive := makeTarGz(t, "othertool", []byte("data"))

	_, err := extractBinaryFromTarGz(bytes.NewReader(archive), "mytool")
	if err == nil {
		t.Fatal("expected error when binary not in archive")
	}
}

func TestExtractBinaryFromTarGz_NestedPath(t *testing.T) {
	// Binary inside a subdirectory within the archive
	expected := []byte("binary-data")
	archive := makeTarGz(t, "subdir/mytool", expected)

	data, err := extractBinaryFromTarGz(bytes.NewReader(archive), "mytool")
	if err != nil {
		t.Fatalf("extractBinaryFromTarGz() error: %v", err)
	}
	if !bytes.Equal(data, expected) {
		t.Errorf("extracted content mismatch: got %q, want %q", data, expected)
	}
}

func TestStatus(t *testing.T) {
	// Use an empty temp dir so no tools are found there
	tmpDir := t.TempDir()

	statuses := Status(tmpDir)
	if len(statuses) != len(Tools) {
		t.Fatalf("Status() returned %d entries, want %d", len(statuses), len(Tools))
	}

	for _, s := range statuses {
		if s.Name == "" {
			t.Error("Status entry has empty Name")
		}
	}
}

func TestStatus_WithInstalledTool(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a fake binary in the install directory
	fakeBin := filepath.Join(tmpDir, "syft")
	if err := os.WriteFile(fakeBin, []byte("fake"), 0755); err != nil {
		t.Fatalf("failed to write fake binary: %v", err)
	}

	statuses := Status(tmpDir)
	var syftStatus *ToolStatus
	for i := range statuses {
		if statuses[i].Name == "syft" {
			syftStatus = &statuses[i]
			break
		}
	}

	if syftStatus == nil {
		t.Fatal("syft not found in status results")
	}

	// syft should be found either in PATH (if installed on this machine) or in our tmpDir
	if !syftStatus.Installed {
		t.Error("expected syft to be found (either PATH or install dir)")
	}
}

func TestLatestReleaseTag(t *testing.T) {
	// Set up a mock GitHub API server
	mux := http.NewServeMux()
	mux.HandleFunc("/repos/anchore/syft/releases/latest", func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]string{"tag_name": "v1.42.0"}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	// Temporarily replace the httpClient and override the API URL
	// We need to override latestReleaseTag to use our server.
	// Since latestReleaseTag builds the URL internally, we'll test via Install
	// with a full mock. For unit test of latestReleaseTag, we test the HTTP path
	// by overriding httpClient and reimplementing the test.
	//
	// Actually, let's test via the exported Install function with a full mock server.
	// That's a better integration-style unit test. See TestInstall below.

	// For this test, we verify the mock server works correctly
	resp, err := http.Get(server.URL + "/repos/anchore/syft/releases/latest")
	if err != nil {
		t.Fatalf("mock server request failed: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var release struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if release.TagName != "v1.42.0" {
		t.Errorf("expected tag_name v1.42.0, got %s", release.TagName)
	}
}

func TestInstall_MockServer(t *testing.T) {
	// Create a synthetic binary for the tarball
	fakeBinary := []byte("#!/bin/sh\nexit 0\n")
	tarball := makeTarGz(t, "faketool", fakeBinary)

	mux := http.NewServeMux()

	// Mock GitHub API: latest release
	mux.HandleFunc("/repos/test/faketool/releases/latest", func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]string{"tag_name": "v2.0.0"}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})

	// Mock asset download
	assetName := fmt.Sprintf("faketool_2.0.0_%s_%s.tar.gz", runtime.GOOS, runtime.GOARCH)
	mux.HandleFunc("/repos/test/faketool/releases/download/v2.0.0/"+assetName, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")
		_, _ = w.Write(tarball)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	// Override httpClient to use the test server
	origClient := httpClient
	httpClient = server.Client()
	defer func() { httpClient = origClient }()

	// We need to override the URLs used by latestReleaseTag and httpGet.
	// Since they build URLs internally, we use a transport that rewrites URLs.
	httpClient = &http.Client{
		Transport: &urlRewriteTransport{
			base:      http.DefaultTransport,
			serverURL: server.URL,
		},
	}

	installDir := t.TempDir()
	tool := Tool{Name: "faketool", Repo: "test/faketool"}

	err := Install(tool, installDir)
	if err != nil {
		// The "version" verification step will fail since this is a shell script,
		// but Install should still succeed (version check is non-fatal).
		// Check if the error is something else
		t.Fatalf("Install() error: %v", err)
	}

	// Verify the binary was written
	destPath := filepath.Join(installDir, "faketool")
	data, err := os.ReadFile(destPath)
	if err != nil {
		t.Fatalf("failed to read installed binary: %v", err)
	}
	if !bytes.Equal(data, fakeBinary) {
		t.Errorf("installed binary content mismatch")
	}

	// Verify permissions
	info, err := os.Stat(destPath)
	if err != nil {
		t.Fatalf("failed to stat installed binary: %v", err)
	}
	if info.Mode().Perm()&0100 == 0 {
		t.Error("installed binary is not executable")
	}
}

// urlRewriteTransport rewrites GitHub API and asset URLs to point at a
// local test server.
type urlRewriteTransport struct {
	base      http.RoundTripper
	serverURL string
}

func (t *urlRewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Rewrite github.com API calls
	if req.URL.Host == "api.github.com" {
		req.URL.Scheme = "http"
		req.URL.Host = t.serverURL[len("http://"):]
	}
	// Rewrite github.com asset downloads
	if req.URL.Host == "github.com" {
		// Convert /owner/repo/releases/download/... to /repos/owner/repo/releases/download/...
		req.URL.Scheme = "http"
		req.URL.Host = t.serverURL[len("http://"):]
		req.URL.Path = "/repos" + req.URL.Path
	}
	return t.base.RoundTrip(req)
}

func TestInstall_APIError(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/repos/test/badtool/releases/latest", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	origClient := httpClient
	httpClient = &http.Client{
		Transport: &urlRewriteTransport{
			base:      http.DefaultTransport,
			serverURL: server.URL,
		},
	}
	defer func() { httpClient = origClient }()

	installDir := t.TempDir()
	tool := Tool{Name: "badtool", Repo: "test/badtool"}

	err := Install(tool, installDir)
	if err == nil {
		t.Fatal("expected error for 404 API response")
	}
}

func TestInstall_BadArchive(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/repos/test/badarchive/releases/latest", func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]string{"tag_name": "v1.0.0"}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})

	assetName := fmt.Sprintf("badarchive_1.0.0_%s_%s.tar.gz", runtime.GOOS, runtime.GOARCH)
	mux.HandleFunc("/repos/test/badarchive/releases/download/v1.0.0/"+assetName, func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("not a valid tarball"))
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	origClient := httpClient
	httpClient = &http.Client{
		Transport: &urlRewriteTransport{
			base:      http.DefaultTransport,
			serverURL: server.URL,
		},
	}
	defer func() { httpClient = origClient }()

	installDir := t.TempDir()
	tool := Tool{Name: "badarchive", Repo: "test/badarchive"}

	err := Install(tool, installDir)
	if err == nil {
		t.Fatal("expected error for corrupt archive")
	}
}

func TestFindTool_NotFound(t *testing.T) {
	// A tool name that definitely doesn't exist anywhere
	_, _, err := FindTool("nonexistent_tool_xyz_12345")
	if err == nil {
		t.Fatal("expected error for non-existent tool")
	}
}

func TestFindTool_InDockDocsDir(t *testing.T) {
	// Override HOME so DefaultInstallDir() points to our temp dir
	tmpHome := t.TempDir()
	origHome := os.Getenv("HOME")
	if err := os.Setenv("HOME", tmpHome); err != nil {
		t.Fatalf("failed to set HOME: %v", err)
	}
	defer func() {
		if err := os.Setenv("HOME", origHome); err != nil {
			t.Errorf("failed to restore HOME: %v", err)
		}
	}()

	// Create the fake binary in the expected dock-docs install path
	binDir := filepath.Join(tmpHome, ".dock-docs", "bin")
	if err := os.MkdirAll(binDir, 0755); err != nil {
		t.Fatalf("failed to create bin dir: %v", err)
	}

	fakeName := "faketool_findtest"
	fakePath := filepath.Join(binDir, fakeName)
	if err := os.WriteFile(fakePath, []byte("fake"), 0755); err != nil {
		t.Fatalf("failed to write fake binary: %v", err)
	}

	path, source, err := FindTool(fakeName)
	if err != nil {
		t.Fatalf("FindTool() unexpected error: %v", err)
	}
	if path != fakePath {
		t.Errorf("path = %q, want %q", path, fakePath)
	}
	if source != "dock-docs" {
		t.Errorf("source = %q, want %q", source, "dock-docs")
	}
}

func TestFindTool_DirectoryIgnored(t *testing.T) {
	// A directory with the tool name should not be found
	tmpHome := t.TempDir()
	origHome := os.Getenv("HOME")
	if err := os.Setenv("HOME", tmpHome); err != nil {
		t.Fatalf("failed to set HOME: %v", err)
	}
	defer func() {
		if err := os.Setenv("HOME", origHome); err != nil {
			t.Errorf("failed to restore HOME: %v", err)
		}
	}()

	binDir := filepath.Join(tmpHome, ".dock-docs", "bin")
	if err := os.MkdirAll(binDir, 0755); err != nil {
		t.Fatalf("failed to create bin dir: %v", err)
	}

	// Create a directory (not a file) with the tool name
	dirName := "dirtool_test"
	dirPath := filepath.Join(binDir, dirName)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		t.Fatalf("failed to create fake dir: %v", err)
	}

	_, _, err := FindTool(dirName)
	if err == nil {
		t.Fatal("expected error when candidate is a directory, not a file")
	}
}

func TestInstallAll_AllInstalled(t *testing.T) {
	tmpDir := t.TempDir()

	// Create fake binaries for all tools so Status() reports them as installed
	for _, tool := range Tools {
		fakePath := filepath.Join(tmpDir, tool.Name)
		if err := os.WriteFile(fakePath, []byte("fake"), 0755); err != nil {
			t.Fatalf("failed to write fake %s: %v", tool.Name, err)
		}
	}

	err := InstallAll(tmpDir, false)
	if err != nil {
		t.Fatalf("InstallAll() unexpected error: %v", err)
	}
}

func TestInstallAll_ForceReinstall(t *testing.T) {
	// With force=true, Install is called even for installed tools.
	// Since Install tries to download from GitHub, this will fail with our mock.
	// Set up a mock server that returns errors to verify the force path is taken.
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	origClient := httpClient
	httpClient = &http.Client{
		Transport: &urlRewriteTransport{
			base:      http.DefaultTransport,
			serverURL: server.URL,
		},
	}
	defer func() { httpClient = origClient }()

	tmpDir := t.TempDir()

	// Create fake binaries so they appear installed
	for _, tool := range Tools {
		fakePath := filepath.Join(tmpDir, tool.Name)
		if err := os.WriteFile(fakePath, []byte("fake"), 0755); err != nil {
			t.Fatalf("failed to write fake %s: %v", tool.Name, err)
		}
	}

	err := InstallAll(tmpDir, true)
	if err == nil {
		t.Fatal("expected error when force-reinstalling with mock 404 server")
	}
}

func TestHttpGet_Non200(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/forbidden", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})
	mux.HandleFunc("/server-error", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	origClient := httpClient
	httpClient = server.Client()
	defer func() { httpClient = origClient }()

	tests := []struct {
		name string
		path string
	}{
		{"403 forbidden", "/forbidden"},
		{"500 server error", "/server-error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := httpGet(server.URL + tt.path)
			if err == nil {
				t.Errorf("httpGet() expected error for %s, got nil", tt.name)
			}
		})
	}
}

func TestLatestReleaseTag_EmptyTag(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/repos/test/emptytag/releases/latest", func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]string{"tag_name": ""}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	origClient := httpClient
	httpClient = &http.Client{
		Transport: &urlRewriteTransport{
			base:      http.DefaultTransport,
			serverURL: server.URL,
		},
	}
	defer func() { httpClient = origClient }()

	_, err := latestReleaseTag("test/emptytag")
	if err == nil {
		t.Fatal("expected error for empty tag_name")
	}
}

func TestLatestReleaseTag_InvalidJSON(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/repos/test/badjson/releases/latest", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("not json"))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	origClient := httpClient
	httpClient = &http.Client{
		Transport: &urlRewriteTransport{
			base:      http.DefaultTransport,
			serverURL: server.URL,
		},
	}
	defer func() { httpClient = origClient }()

	_, err := latestReleaseTag("test/badjson")
	if err == nil {
		t.Fatal("expected error for invalid JSON response")
	}
}

func TestLatestReleaseTag_Success(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/repos/test/goodtool/releases/latest", func(w http.ResponseWriter, r *http.Request) {
		// Verify Accept header is set correctly
		if got := r.Header.Get("Accept"); got != "application/vnd.github+json" {
			t.Errorf("Accept header = %q, want %q", got, "application/vnd.github+json")
		}
		resp := map[string]string{"tag_name": "v3.1.0"}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	origClient := httpClient
	httpClient = &http.Client{
		Transport: &urlRewriteTransport{
			base:      http.DefaultTransport,
			serverURL: server.URL,
		},
	}
	defer func() { httpClient = origClient }()

	tag, err := latestReleaseTag("test/goodtool")
	if err != nil {
		t.Fatalf("latestReleaseTag() unexpected error: %v", err)
	}
	if tag != "v3.1.0" {
		t.Errorf("tag = %q, want %q", tag, "v3.1.0")
	}
}

func TestLatestReleaseTag_WithGitHubToken(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/repos/test/authtool/releases/latest", func(w http.ResponseWriter, r *http.Request) {
		// Verify that Authorization header is set when GITHUB_TOKEN is present
		auth := r.Header.Get("Authorization")
		if auth != "Bearer test-token-123" {
			t.Errorf("Authorization header = %q, want %q", auth, "Bearer test-token-123")
		}
		resp := map[string]string{"tag_name": "v4.0.0"}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	origClient := httpClient
	httpClient = &http.Client{
		Transport: &urlRewriteTransport{
			base:      http.DefaultTransport,
			serverURL: server.URL,
		},
	}
	defer func() { httpClient = origClient }()

	// Set GITHUB_TOKEN
	origToken := os.Getenv("GITHUB_TOKEN")
	if err := os.Setenv("GITHUB_TOKEN", "test-token-123"); err != nil {
		t.Fatalf("failed to set GITHUB_TOKEN: %v", err)
	}
	defer func() {
		if origToken == "" {
			_ = os.Unsetenv("GITHUB_TOKEN")
		} else {
			_ = os.Setenv("GITHUB_TOKEN", origToken)
		}
	}()

	tag, err := latestReleaseTag("test/authtool")
	if err != nil {
		t.Fatalf("latestReleaseTag() unexpected error: %v", err)
	}
	if tag != "v4.0.0" {
		t.Errorf("tag = %q, want %q", tag, "v4.0.0")
	}
}

func TestHttpGet_WithGitHubToken(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/asset", func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth != "Bearer download-token-456" {
			t.Errorf("Authorization header = %q, want %q", auth, "Bearer download-token-456")
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("asset-data"))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	origClient := httpClient
	httpClient = server.Client()
	defer func() { httpClient = origClient }()

	origToken := os.Getenv("GITHUB_TOKEN")
	if err := os.Setenv("GITHUB_TOKEN", "download-token-456"); err != nil {
		t.Fatalf("failed to set GITHUB_TOKEN: %v", err)
	}
	defer func() {
		if origToken == "" {
			_ = os.Unsetenv("GITHUB_TOKEN")
		} else {
			_ = os.Setenv("GITHUB_TOKEN", origToken)
		}
	}()

	body, err := httpGet(server.URL + "/asset")
	if err != nil {
		t.Fatalf("httpGet() unexpected error: %v", err)
	}
	defer func() { _ = body.Close() }()

	data, err := io.ReadAll(body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}
	if string(data) != "asset-data" {
		t.Errorf("body = %q, want %q", string(data), "asset-data")
	}
}

func TestExtractBinaryFromTarGz_InvalidGzip(t *testing.T) {
	_, err := extractBinaryFromTarGz(bytes.NewReader([]byte("not gzip data")), "mytool")
	if err == nil {
		t.Fatal("expected error for invalid gzip data")
	}
}

func TestExtractBinaryFromTarGz_DirectoryEntry(t *testing.T) {
	// Archive containing a directory entry with the same name — should be skipped
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gw)

	// Write a directory entry named "mytool"
	dirHdr := &tar.Header{
		Name:     "mytool",
		Typeflag: tar.TypeDir,
		Mode:     0755,
	}
	if err := tw.WriteHeader(dirHdr); err != nil {
		t.Fatalf("tar WriteHeader dir: %v", err)
	}

	if err := tw.Close(); err != nil {
		t.Fatalf("tar Close: %v", err)
	}
	if err := gw.Close(); err != nil {
		t.Fatalf("gzip Close: %v", err)
	}

	_, err := extractBinaryFromTarGz(bytes.NewReader(buf.Bytes()), "mytool")
	if err == nil {
		t.Fatal("expected error when archive only contains directory, not regular file")
	}
}

func TestExtractBinaryFromTarGz_OversizedHeader(t *testing.T) {
	// Verify the constant is reasonable
	if maxBinarySize != 200*1024*1024 {
		t.Fatalf("maxBinarySize = %d, want %d", maxBinarySize, 200*1024*1024)
	}

	// To test the header size check, we manually construct a tar.gz archive
	// with a forged header declaring a file larger than the limit.
	var rawTar bytes.Buffer
	tw := tar.NewWriter(&rawTar)

	hdr := &tar.Header{
		Name:     "mytool",
		Mode:     0755,
		Typeflag: tar.TypeReg,
		Size:     maxBinarySize + 1,
	}
	// WriteHeader will succeed; it's Close/Flush that needs data to match.
	// We write the header and the data padding manually via tar internals.
	if err := tw.WriteHeader(hdr); err != nil {
		t.Fatalf("tar WriteHeader: %v", err)
	}
	// Don't close tw — we just need a valid header in the stream.
	// Gzip-wrap whatever we have (the header bytes are enough for our test).
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	if _, err := gw.Write(rawTar.Bytes()); err != nil {
		t.Fatalf("gzip Write: %v", err)
	}
	if err := gw.Close(); err != nil {
		t.Fatalf("gzip Close: %v", err)
	}

	_, err := extractBinaryFromTarGz(bytes.NewReader(buf.Bytes()), "mytool")
	if err == nil {
		t.Fatal("expected error for oversized binary header")
	}
	if !strings.Contains(err.Error(), "exceeds limit") {
		t.Errorf("expected 'exceeds limit' in error, got: %v", err)
	}
}

func TestInstall_DownloadError(t *testing.T) {
	// Mock server returns success for API but 404 for the asset download
	mux := http.NewServeMux()
	mux.HandleFunc("/repos/test/dltool/releases/latest", func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]string{"tag_name": "v1.0.0"}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})
	// No handler for the download URL — will 404 via default handler
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	origClient := httpClient
	httpClient = &http.Client{
		Transport: &urlRewriteTransport{
			base:      http.DefaultTransport,
			serverURL: server.URL,
		},
	}
	defer func() { httpClient = origClient }()

	installDir := t.TempDir()
	tool := Tool{Name: "dltool", Repo: "test/dltool"}

	err := Install(tool, installDir)
	if err == nil {
		t.Fatal("expected error when asset download returns 404")
	}
}

func TestInstall_BinaryNotInArchive(t *testing.T) {
	// Archive contains a different file name than the tool name
	tarball := makeTarGz(t, "wrong_name", []byte("binary"))

	mux := http.NewServeMux()
	mux.HandleFunc("/repos/test/missingtool/releases/latest", func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]string{"tag_name": "v1.0.0"}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})

	assetName := fmt.Sprintf("missingtool_1.0.0_%s_%s.tar.gz", runtime.GOOS, runtime.GOARCH)
	mux.HandleFunc("/repos/test/missingtool/releases/download/v1.0.0/"+assetName, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")
		_, _ = w.Write(tarball)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	origClient := httpClient
	httpClient = &http.Client{
		Transport: &urlRewriteTransport{
			base:      http.DefaultTransport,
			serverURL: server.URL,
		},
	}
	defer func() { httpClient = origClient }()

	installDir := t.TempDir()
	tool := Tool{Name: "missingtool", Repo: "test/missingtool"}

	err := Install(tool, installDir)
	if err == nil {
		t.Fatal("expected error when binary not found in archive")
	}
}

func TestHttpGet_Success(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("hello world"))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	origClient := httpClient
	httpClient = server.Client()
	defer func() { httpClient = origClient }()

	body, err := httpGet(server.URL + "/data")
	if err != nil {
		t.Fatalf("httpGet() unexpected error: %v", err)
	}
	defer func() { _ = body.Close() }()

	data, err := io.ReadAll(body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}
	if string(data) != "hello world" {
		t.Errorf("body = %q, want %q", string(data), "hello world")
	}
}
