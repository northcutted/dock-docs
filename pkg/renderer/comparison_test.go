package renderer

import (
	"strings"
	"testing"

	"github.com/northcutted/dock-docs/pkg/types"
)

func TestRenderComparison(t *testing.T) {
	comparisonStats := []*types.ImageStats{
		{
			ImageTag:     "app:v1",
			SizeBytes:    104857600, // 100 MB
			Architecture: "amd64",
			OS:           "linux",
			Efficiency:   90.0,
			WastedBytes:  "10 MB",
			TotalLayers:  5,
			VulnSummary:  map[string]int{"Critical": 1, "High": 0},
			Vulnerabilities: []types.Vulnerability{
				{ID: "CVE-2023-1111", Severity: "Critical", Package: "libssl", Version: "1.0"},
			},
			TotalPackages: 2,
			Packages: []types.PackageSummary{
				{Name: "bash", Version: "5.0"},
			},
		},
		{
			ImageTag:     "app:v2",
			SizeBytes:    83886080, // 80 MB
			Architecture: "amd64",
			OS:           "linux",
			Efficiency:   98.0,
			WastedBytes:  "2 MB",
			TotalLayers:  4,
			VulnSummary:  map[string]int{"Critical": 0, "High": 1},
			Vulnerabilities: []types.Vulnerability{
				{ID: "CVE-2023-2222", Severity: "High", Package: "curl", Version: "7.88"},
			},
			TotalPackages: 3,
			Packages: []types.PackageSummary{
				{Name: "bash", Version: "5.2"},
			},
			SupportedArchitectures: []string{"linux/amd64", "linux/arm64"},
		},
	}

	output, err := RenderComparison(comparisonStats, RenderOptions{})
	if err != nil {
		t.Fatalf("RenderComparison() error = %v", err)
	}

	// 1. Check for the Summary Table
	if !strings.Contains(output, "### 🏷️ Supported Tags") {
		t.Error("expected output to contain Matrix Comparison header")
	}
	if !strings.Contains(output, "| `app:v1` | ![Size](") {
		t.Error("expected output to contain summary row for app:v1")
	}
	if !strings.Contains(output, "| `app:v2` | ![Size](") {
		t.Error("expected output to contain summary row for app:v2")
	}
	// Check for Multi-Arch support in v2
	if !strings.Contains(output, "linux/amd64, linux/arm64") {
		t.Error("expected output to contain supported architectures for app:v2")
	}

	// 2. Check for Detailed Reports (Collapsible)
	if !strings.Contains(output, "<details>") {
		t.Error("expected output to contain <details> tags for collapsible sections")
	}
	if !strings.Contains(output, "<summary><strong>🔍 Full Report: app:v1</strong></summary>") {
		t.Error("expected output to contain summary tag for app:v1")
	}
	if !strings.Contains(output, "<summary><strong>🔍 Full Report: app:v2</strong></summary>") {
		t.Error("expected output to contain summary tag for app:v2")
	}

	// 3. Check for Specific Details in the Reports
	// v1 details
	if !strings.Contains(output, "CVE-2023-1111") {
		t.Error("expected output to contain v1 vulnerability CVE-2023-1111")
	}
	if !strings.Contains(output, "libssl") {
		t.Error("expected output to contain v1 package libssl")
	}

	// v2 details
	if !strings.Contains(output, "CVE-2023-2222") {
		t.Error("expected output to contain v2 vulnerability CVE-2023-2222")
	}
	if !strings.Contains(output, "curl") {
		t.Error("expected output to contain v2 package curl")
	}
}
