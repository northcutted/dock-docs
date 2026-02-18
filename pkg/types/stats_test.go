package types

import (
	"net/url"
	"strings"
	"testing"
)

func TestImageStats_SizeBadge(t *testing.T) {
	baseURL := "https://img.shields.io/static/v1"

	tests := []struct {
		name     string
		stats    *ImageStats
		wantURL  bool
		contains []string
	}{
		{
			name:     "valid size",
			stats:    &ImageStats{SizeBytes: 67633152}, // 64.50 MB
			wantURL:  true,
			contains: []string{"label=Size", "message=64.50+MB", "color=blue"},
		},
		{
			name:    "zero size",
			stats:   &ImageStats{SizeBytes: 0},
			wantURL: false,
		},
		{
			name:     "large size",
			stats:    &ImageStats{SizeBytes: 1073741824}, // 1024.00 MB
			wantURL:  true,
			contains: []string{"label=Size", "message=1024.00+MB", "color=blue"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.stats.SizeBadge(baseURL)
			if !tt.wantURL {
				if got != "" {
					t.Errorf("SizeBadge() = %v, want empty string", got)
				}
				return
			}

			if got == "" {
				t.Error("SizeBadge() returned empty string")
			}

			for _, substr := range tt.contains {
				if !strings.Contains(got, substr) {
					t.Errorf("SizeBadge() missing substring %q, got: %v", substr, got)
				}
			}
		})
	}
}

func TestImageStats_LayersBadge(t *testing.T) {
	baseURL := "https://img.shields.io/static/v1"

	tests := []struct {
		name     string
		stats    *ImageStats
		wantURL  bool
		contains []string
	}{
		{
			name:     "5 layers",
			stats:    &ImageStats{TotalLayers: 5},
			wantURL:  true,
			contains: []string{"label=Layers", "message=5", "color=blue"},
		},
		{
			name:    "zero layers",
			stats:   &ImageStats{TotalLayers: 0},
			wantURL: false,
		},
		{
			name:     "many layers",
			stats:    &ImageStats{TotalLayers: 42},
			wantURL:  true,
			contains: []string{"label=Layers", "message=42", "color=blue"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.stats.LayersBadge(baseURL)
			if !tt.wantURL {
				if got != "" {
					t.Errorf("LayersBadge() = %v, want empty string", got)
				}
				return
			}

			if got == "" {
				t.Error("LayersBadge() returned empty string")
			}

			for _, substr := range tt.contains {
				if !strings.Contains(got, substr) {
					t.Errorf("LayersBadge() missing substring %q, got: %v", substr, got)
				}
			}
		})
	}
}

func TestImageStats_EfficiencyBadge(t *testing.T) {
	baseURL := "https://img.shields.io/static/v1"

	tests := []struct {
		name     string
		stats    *ImageStats
		wantURL  bool
		contains []string
	}{
		{
			name:     "high efficiency (green)",
			stats:    &ImageStats{Efficiency: 95.5},
			wantURL:  true,
			contains: []string{"label=Efficiency", "message=95.5%", "color=green"},
		},
		{
			name:     "medium efficiency (yellow)",
			stats:    &ImageStats{Efficiency: 85.0},
			wantURL:  true,
			contains: []string{"label=Efficiency", "message=85.0%", "color=yellow"},
		},
		{
			name:     "low efficiency (red)",
			stats:    &ImageStats{Efficiency: 75.0},
			wantURL:  true,
			contains: []string{"label=Efficiency", "message=75.0%", "color=red"},
		},
		{
			name:     "edge case: exactly 90% (yellow)",
			stats:    &ImageStats{Efficiency: 89.9},
			wantURL:  true,
			contains: []string{"label=Efficiency", "message=89.9%", "color=yellow"},
		},
		{
			name:     "edge case: exactly 80% (red)",
			stats:    &ImageStats{Efficiency: 79.9},
			wantURL:  true,
			contains: []string{"label=Efficiency", "message=79.9%", "color=red"},
		},
		{
			name:    "zero efficiency",
			stats:   &ImageStats{Efficiency: 0},
			wantURL: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.stats.EfficiencyBadge(baseURL)
			if !tt.wantURL {
				if got != "" {
					t.Errorf("EfficiencyBadge() = %v, want empty string", got)
				}
				return
			}

			if got == "" {
				t.Error("EfficiencyBadge() returned empty string")
			}

			for _, substr := range tt.contains {
				if !strings.Contains(got, substr) {
					t.Errorf("EfficiencyBadge() missing substring %q, got: %v", substr, got)
				}
			}
		})
	}
}

func TestImageStats_VulnBadge(t *testing.T) {
	baseURL := "https://img.shields.io/static/v1"

	tests := []struct {
		name     string
		stats    *ImageStats
		contains []string
		color    string
	}{
		{
			name: "critical vulnerabilities (red)",
			stats: &ImageStats{
				Vulnerabilities: []Vulnerability{
					{Severity: "Critical"}, {Severity: "Critical"},
					{Severity: "High"}, {Severity: "High"}, {Severity: "High"}, {Severity: "High"}, {Severity: "High"},
					{Severity: "Medium"}, {Severity: "Medium"}, {Severity: "Medium"}, {Severity: "Medium"}, {Severity: "Medium"},
					{Severity: "Medium"}, {Severity: "Medium"}, {Severity: "Medium"}, {Severity: "Medium"}, {Severity: "Medium"},
					{Severity: "Low"}, {Severity: "Low"}, {Severity: "Low"},
				},
				VulnSummary: map[string]int{
					"Critical": 2,
					"High":     5,
					"Medium":   10,
					"Low":      3,
				},
			},
			contains: []string{"label=Security", "color=red"},
			color:    "red",
		},
		{
			name: "high vulnerabilities only (orange)",
			stats: &ImageStats{
				Vulnerabilities: []Vulnerability{
					{Severity: "High"}, {Severity: "High"}, {Severity: "High"},
					{Severity: "Medium"}, {Severity: "Medium"}, {Severity: "Medium"}, {Severity: "Medium"}, {Severity: "Medium"},
					{Severity: "Low"}, {Severity: "Low"},
				},
				VulnSummary: map[string]int{
					"Critical": 0,
					"High":     3,
					"Medium":   5,
					"Low":      2,
				},
			},
			contains: []string{"label=Security", "color=orange"},
			color:    "orange",
		},
		{
			name: "no critical or high (green)",
			stats: &ImageStats{
				Vulnerabilities: []Vulnerability{
					{Severity: "Medium"}, {Severity: "Medium"}, {Severity: "Medium"},
					{Severity: "Low"}, {Severity: "Low"}, {Severity: "Low"}, {Severity: "Low"}, {Severity: "Low"},
				},
				VulnSummary: map[string]int{
					"Critical": 0,
					"High":     0,
					"Medium":   3,
					"Low":      5,
				},
			},
			contains: []string{"label=Security", "color=green"},
			color:    "green",
		},
		{
			name: "no vulnerabilities (green)",
			stats: &ImageStats{
				VulnSummary: map[string]int{},
			},
			contains: []string{"label=Security", "color=green", "0+Crit"},
			color:    "green",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.stats.VulnBadge(baseURL)

			if got == "" {
				t.Error("VulnBadge() returned empty string")
			}

			for _, substr := range tt.contains {
				if !strings.Contains(got, substr) {
					t.Errorf("VulnBadge() missing substring %q, got: %v", substr, got)
				}
			}

			// Verify correct color
			if !strings.Contains(got, "color="+tt.color) {
				t.Errorf("VulnBadge() wrong color, want %s, got: %v", tt.color, got)
			}

			// Verify URL encoding is valid
			_, err := url.Parse(got)
			if err != nil {
				t.Errorf("VulnBadge() produced invalid URL: %v", err)
			}
		})
	}
}

func TestImageStats_TotalVulns(t *testing.T) {
	tests := []struct {
		name  string
		stats *ImageStats
		want  int
	}{
		{
			name: "multiple vulnerabilities",
			stats: &ImageStats{
				Vulnerabilities: []Vulnerability{
					{ID: "CVE-2023-0001", Severity: "Critical"},
					{ID: "CVE-2023-0002", Severity: "High"},
					{ID: "CVE-2023-0003", Severity: "Medium"},
				},
			},
			want: 3,
		},
		{
			name: "no vulnerabilities",
			stats: &ImageStats{
				Vulnerabilities: []Vulnerability{},
			},
			want: 0,
		},
		{
			name:  "nil vulnerabilities",
			stats: &ImageStats{},
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stats.TotalVulns(); got != tt.want {
				t.Errorf("TotalVulns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackageSummary(t *testing.T) {
	pkg := PackageSummary{
		Name:    "openssl",
		Version: "3.0.2",
	}

	if pkg.Name != "openssl" {
		t.Errorf("Name = %v, want openssl", pkg.Name)
	}

	if pkg.Version != "3.0.2" {
		t.Errorf("Version = %v, want 3.0.2", pkg.Version)
	}
}

func TestVulnerability(t *testing.T) {
	vuln := Vulnerability{
		ID:       "CVE-2023-1234",
		Severity: "Critical",
		Package:  "curl",
		Version:  "7.81.0",
	}

	if vuln.ID != "CVE-2023-1234" {
		t.Errorf("ID = %v, want CVE-2023-1234", vuln.ID)
	}

	if vuln.Severity != "Critical" {
		t.Errorf("Severity = %v, want Critical", vuln.Severity)
	}

	if vuln.Package != "curl" {
		t.Errorf("Package = %v, want curl", vuln.Package)
	}

	if vuln.Version != "7.81.0" {
		t.Errorf("Version = %v, want 7.81.0", vuln.Version)
	}
}

func TestImageStats_URLEncoding(t *testing.T) {
	baseURL := "https://img.shields.io/static/v1"

	stats := &ImageStats{
		SizeBytes: 105381888, // 100.50 MB
		VulnSummary: map[string]int{
			"Critical": 5,
			"High":     10,
		},
	}

	// Test that special characters are properly URL encoded
	sizeBadge := stats.SizeBadge(baseURL)
	vulnBadge := stats.VulnBadge(baseURL)

	// Verify valid URLs
	if _, err := url.Parse(sizeBadge); err != nil {
		t.Errorf("SizeBadge produced invalid URL: %v", err)
	}

	if _, err := url.Parse(vulnBadge); err != nil {
		t.Errorf("VulnBadge produced invalid URL: %v", err)
	}

	// Verify URL encoding for spaces
	if !strings.Contains(sizeBadge, "+") && !strings.Contains(sizeBadge, "%20") {
		t.Error("SizeBadge should properly encode spaces")
	}
}

func TestImageStats_CustomBaseURL(t *testing.T) {
	customBaseURL := "https://custom-badges.example.com/badge"

	stats := &ImageStats{
		SizeBytes:   52428800, // 50.00 MB
		TotalLayers: 10,
		Efficiency:  95.0,
		VulnSummary: map[string]int{"Critical": 0},
	}

	sizeBadge := stats.SizeBadge(customBaseURL)
	layersBadge := stats.LayersBadge(customBaseURL)
	effBadge := stats.EfficiencyBadge(customBaseURL)
	vulnBadge := stats.VulnBadge(customBaseURL)

	// All badges should use custom base URL
	badges := []string{sizeBadge, layersBadge, effBadge, vulnBadge}
	for i, badge := range badges {
		if !strings.HasPrefix(badge, customBaseURL) {
			t.Errorf("Badge %d should use custom base URL, got: %v", i, badge)
		}
	}
}

func TestImageStats_EmptyStats(t *testing.T) {
	baseURL := "https://img.shields.io/static/v1"
	stats := &ImageStats{}

	// All badge methods should return empty strings for empty stats
	if got := stats.SizeBadge(baseURL); got != "" {
		t.Errorf("SizeBadge() for empty stats = %v, want empty", got)
	}

	if got := stats.LayersBadge(baseURL); got != "" {
		t.Errorf("LayersBadge() for empty stats = %v, want empty", got)
	}

	if got := stats.EfficiencyBadge(baseURL); got != "" {
		t.Errorf("EfficiencyBadge() for empty stats = %v, want empty", got)
	}

	// VulnBadge should always return a badge (even for 0 vulns)
	vulnBadge := stats.VulnBadge(baseURL)
	if vulnBadge == "" {
		t.Error("VulnBadge() should return badge even for 0 vulnerabilities")
	}

	if !strings.Contains(vulnBadge, "0+Crit") {
		t.Errorf("VulnBadge() for empty stats should show 0 Crit, got: %v", vulnBadge)
	}
}

func TestImageStats_VulnBadge_IncludesUnknown(t *testing.T) {
	baseURL := "https://img.shields.io/static/v1"

	stats := &ImageStats{
		Vulnerabilities: []Vulnerability{
			{ID: "CVE-1", Severity: "Unknown"},
			{ID: "CVE-2", Severity: "Unknown"},
			{ID: "CVE-3", Severity: "Low"},
		},
		VulnSummary: map[string]int{
			"Unknown": 2,
			"Low":     1,
		},
	}

	badge := stats.VulnBadge(baseURL)

	// VulnBadge should count all 3 vulnerabilities (including Unknown severity)
	if !strings.Contains(badge, "3+Vulns") {
		t.Errorf("VulnBadge() should count Unknown severity vulns, got: %v", badge)
	}
	if !strings.Contains(badge, "color=green") {
		t.Errorf("VulnBadge() should be green (no Critical/High), got: %v", badge)
	}
}
