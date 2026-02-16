package analysis

import (
	"github.com/northcutted/dock-docs/pkg/types"
	"errors"
	"testing"
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
