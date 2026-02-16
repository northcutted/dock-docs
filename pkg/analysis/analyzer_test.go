package analysis

import (
	"errors"
	"strings"
	"testing"

	"github.com/northcutted/dock-docs/pkg/types"
)

// MockRunner implements the Runner interface for testing
type MockRunner struct {
	name        string
	available   bool
	shouldFail  bool
	returnStats *types.ImageStats
}

func (m *MockRunner) Name() string {
	return m.name
}

func (m *MockRunner) IsAvailable() bool {
	return m.available
}

func (m *MockRunner) Run(image string, verbose bool) (*types.ImageStats, error) {
	if m.shouldFail {
		return nil, errors.New("mock runner failed")
	}
	// Return a copy to avoid race conditions if needed, though mostly read-only here
	if m.returnStats == nil {
		return &types.ImageStats{ImageTag: image}, nil
	}
	stats := *m.returnStats
	stats.ImageTag = image // Ensure tag matches input
	return &stats, nil
}

func TestAnalyzeMatrix(t *testing.T) {
	oldEnsureImage := ensureImage
	defer func() { ensureImage = oldEnsureImage }()
	ensureImage = func(image string, verbose bool) error { return nil }
	mockStats := &types.ImageStats{
		SizeMB:     "100MB",
		Efficiency: 95.0,
	}

	runner := &MockRunner{
		name:        "TestRunner",
		available:   true,
		returnStats: mockStats,
	}

	images := []string{"img1:latest", "img2:latest"}
	runners := []Runner{runner}

	results, err := AnalyzeMatrix(images, runners, false)
	if err != nil {
		t.Fatalf("AnalyzeMatrix failed: %v", err)
	}

	if len(results) != 2 {
		t.Errorf("Expected 2 results, got %d", len(results))
	}

	// Verify order is preserved (impl detail: results[i] = stats)
	if results[0].ImageTag != "img1:latest" {
		t.Errorf("Expected first result to be img1:latest, got %s", results[0].ImageTag)
	}
	if results[1].ImageTag != "img2:latest" {
		t.Errorf("Expected second result to be img2:latest, got %s", results[1].ImageTag)
	}
}

func TestAnalyzeMatrix_PartialFailure(t *testing.T) {
	oldEnsureImage := ensureImage
	defer func() { ensureImage = oldEnsureImage }()
	ensureImage = func(image string, verbose bool) error { return nil }
	// Test that if one image analysis fails, others still return (or at least function doesn't crash)
	// The current implementation of AnalyzeMatrix prints error and returns nil for that slot,
	// then filters out nils.

	smartRunner := &SmartMockRunner{
		failOnImage: "bad:image",
	}

	images := []string{"good:image", "bad:image", "good:image2"}
	runners := []Runner{smartRunner}

	results, err := AnalyzeMatrix(images, runners, false)
	if err != nil {
		t.Fatalf("AnalyzeMatrix failed: %v", err)
	}

	// Should have 3 results, even if one failed (it returns partial/empty stats)
	// Wait, check implementation: if AnalyzeImage fails, it returns nil stats but we capture error.
	// In AnalyzeMatrix, if err != nil, we check if stats != nil.
	// But AnalyzeImage returns partial stats even on failure IF it collected partial results?
	// Actually, AnalyzeImage returns "finalStats" even if errs > 0.
	// So results should be 3 valid stats objects.

	if len(results) != 3 {
		t.Errorf("Expected 3 results, got %d", len(results))
	}

	if results[0].ImageTag != "good:image" {
		t.Errorf("Expected first result to be good:image")
	}
	if results[1].ImageTag != "bad:image" {
		t.Errorf("Expected second result to be bad:image")
	}
	if results[2].ImageTag != "good:image2" {
		t.Errorf("Expected third result to be good:image2")
	}
}

type SmartMockRunner struct {
	failOnImage string
}

func (m *SmartMockRunner) Name() string      { return "SmartMock" }
func (m *SmartMockRunner) IsAvailable() bool { return true }
func (m *SmartMockRunner) Run(image string, verbose bool) (*types.ImageStats, error) {
	if image == m.failOnImage {
		return nil, errors.New("simulated failure")
	}
	return &types.ImageStats{ImageTag: image, SizeMB: "10MB"}, nil
}

func TestAnalyzeImage_EmptyImage(t *testing.T) {
	oldEnsureImage := ensureImage
	defer func() { ensureImage = oldEnsureImage }()
	ensureImage = func(image string, verbose bool) error { return nil }

	runner := &MockRunner{name: "test", available: true}
	_, err := AnalyzeImage("", []Runner{runner}, false)
	if err == nil {
		t.Error("Expected error for empty image tag")
	}
	if !strings.Contains(err.Error(), "image tag is required") {
		t.Errorf("Expected 'image tag is required' error, got: %v", err)
	}
}

func TestAnalyzeImage_EnsureImageFailure(t *testing.T) {
	oldEnsureImage := ensureImage
	defer func() { ensureImage = oldEnsureImage }()
	ensureImage = func(image string, verbose bool) error {
		return errors.New("pull failed")
	}

	runner := &MockRunner{name: "test", available: true}
	_, err := AnalyzeImage("test:latest", []Runner{runner}, false)
	if err == nil {
		t.Error("Expected error when ensureImage fails")
	}
	if !strings.Contains(err.Error(), "failed to ensure image") {
		t.Errorf("Expected 'failed to ensure image' error, got: %v", err)
	}
}

func TestAnalyzeImage_UnavailableRunner(t *testing.T) {
	oldEnsureImage := ensureImage
	defer func() { ensureImage = oldEnsureImage }()
	ensureImage = func(image string, verbose bool) error { return nil }

	unavailableRunner := &MockRunner{name: "unavailable", available: false}
	stats, err := AnalyzeImage("test:latest", []Runner{unavailableRunner}, false)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if stats == nil {
		t.Fatal("Expected non-nil stats")
	}
	if stats.ImageTag != "test:latest" {
		t.Errorf("Expected ImageTag test:latest, got %s", stats.ImageTag)
	}
}

func TestAnalyzeImage_MultipleRunners(t *testing.T) {
	oldEnsureImage := ensureImage
	defer func() { ensureImage = oldEnsureImage }()
	ensureImage = func(image string, verbose bool) error { return nil }

	runner1 := &MockRunner{
		name:      "runner1",
		available: true,
		returnStats: &types.ImageStats{
			Architecture: "amd64",
			OS:           "linux",
			SizeMB:       "100MB",
		},
	}
	runner2 := &MockRunner{
		name:      "runner2",
		available: true,
		returnStats: &types.ImageStats{
			Efficiency:    95.0,
			TotalPackages: 50,
			Packages: []types.PackageSummary{
				{Name: "pkg1", Version: "1.0"},
			},
		},
	}

	stats, err := AnalyzeImage("test:latest", []Runner{runner1, runner2}, false)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if stats == nil {
		t.Fatal("Expected non-nil stats")
	}

	// Verify merged stats from both runners
	if stats.Architecture != "amd64" {
		t.Errorf("Expected Architecture amd64, got %s", stats.Architecture)
	}
	if stats.SizeMB != "100MB" {
		t.Errorf("Expected SizeMB 100MB, got %s", stats.SizeMB)
	}
	if stats.Efficiency != 95.0 {
		t.Errorf("Expected Efficiency 95.0, got %f", stats.Efficiency)
	}
	if stats.TotalPackages != 50 {
		t.Errorf("Expected TotalPackages 50, got %d", stats.TotalPackages)
	}
	if len(stats.Packages) != 1 {
		t.Errorf("Expected 1 package, got %d", len(stats.Packages))
	}
}

func TestMergeStats_Comprehensive(t *testing.T) {
	dest := &types.ImageStats{
		ImageTag:    "test:latest",
		VulnSummary: map[string]int{"Critical": 1},
		Packages:    []types.PackageSummary{{Name: "pkg1", Version: "1.0"}},
		Vulnerabilities: []types.Vulnerability{
			{ID: "CVE-001", Severity: "Critical"},
		},
	}

	src := &types.ImageStats{
		Architecture:           "arm64",
		OS:                     "linux",
		OSDistro:               "alpine",
		SizeMB:                 "50MB",
		TotalLayers:            10,
		Efficiency:             92.5,
		WastedBytes:            "5MB",
		TotalPackages:          20,
		SupportedArchitectures: []string{"amd64", "arm64"},
		VulnSummary:            map[string]int{"High": 2, "Critical": 1},
		Packages: []types.PackageSummary{
			{Name: "pkg2", Version: "2.0"},
		},
		Vulnerabilities: []types.Vulnerability{
			{ID: "CVE-002", Severity: "High"},
		},
	}

	mergeStats(dest, src)

	// Verify all fields were merged
	if dest.Architecture != "arm64" {
		t.Errorf("Expected Architecture arm64, got %s", dest.Architecture)
	}
	if dest.OS != "linux" {
		t.Errorf("Expected OS linux, got %s", dest.OS)
	}
	if dest.OSDistro != "alpine" {
		t.Errorf("Expected OSDistro alpine, got %s", dest.OSDistro)
	}
	if dest.SizeMB != "50MB" {
		t.Errorf("Expected SizeMB 50MB, got %s", dest.SizeMB)
	}
	if dest.TotalLayers != 10 {
		t.Errorf("Expected TotalLayers 10, got %d", dest.TotalLayers)
	}
	if dest.Efficiency != 92.5 {
		t.Errorf("Expected Efficiency 92.5, got %f", dest.Efficiency)
	}
	if dest.WastedBytes != "5MB" {
		t.Errorf("Expected WastedBytes 5MB, got %s", dest.WastedBytes)
	}
	if dest.TotalPackages != 20 {
		t.Errorf("Expected TotalPackages 20, got %d", dest.TotalPackages)
	}
	if len(dest.SupportedArchitectures) != 2 {
		t.Errorf("Expected 2 supported architectures, got %d", len(dest.SupportedArchitectures))
	}
	if len(dest.Packages) != 2 {
		t.Errorf("Expected 2 packages, got %d", len(dest.Packages))
	}
	if len(dest.Vulnerabilities) != 2 {
		t.Errorf("Expected 2 vulnerabilities, got %d", len(dest.Vulnerabilities))
	}
	if dest.VulnSummary["Critical"] != 2 {
		t.Errorf("Expected 2 Critical vulns, got %d", dest.VulnSummary["Critical"])
	}
	if dest.VulnSummary["High"] != 2 {
		t.Errorf("Expected 2 High vulns, got %d", dest.VulnSummary["High"])
	}
}

func TestMergeStats_NilSource(t *testing.T) {
	dest := &types.ImageStats{
		ImageTag:     "test:latest",
		Architecture: "amd64",
	}
	original := dest.Architecture

	mergeStats(dest, nil)

	// Verify dest is unchanged
	if dest.Architecture != original {
		t.Error("mergeStats modified dest when src is nil")
	}
}

func TestAnalyzeImage_VulnerabilitySorting(t *testing.T) {
	oldEnsureImage := ensureImage
	defer func() { ensureImage = oldEnsureImage }()
	ensureImage = func(image string, verbose bool) error { return nil }

	runner := &MockRunner{
		name:      "vuln-runner",
		available: true,
		returnStats: &types.ImageStats{
			Vulnerabilities: []types.Vulnerability{
				{ID: "CVE-003", Severity: "Low"},
				{ID: "CVE-001", Severity: "Critical"},
				{ID: "CVE-004", Severity: "Medium"},
				{ID: "CVE-002", Severity: "High"},
			},
		},
	}

	stats, err := AnalyzeImage("test:latest", []Runner{runner}, false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Verify vulnerabilities are sorted by severity (Critical > High > Medium > Low)
	if len(stats.Vulnerabilities) != 4 {
		t.Fatalf("Expected 4 vulnerabilities, got %d", len(stats.Vulnerabilities))
	}

	expected := []string{"Critical", "High", "Medium", "Low"}
	for i, vuln := range stats.Vulnerabilities {
		if vuln.Severity != expected[i] {
			t.Errorf("Vulnerability %d: expected severity %s, got %s", i, expected[i], vuln.Severity)
		}
	}
}
