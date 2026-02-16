package types

import (
	"fmt"
	"net/url"
)

// PackageSummary represents a simplified view of a package
type PackageSummary struct {
	Name    string
	Version string
}

// Vulnerability represents a security issue
type Vulnerability struct {
	ID       string // e.g., "CVE-2023-1234"
	Severity string // "Critical", "High", "Medium", "Low"
	Package  string // Package name
	Version  string // Installed version
}

// ImageStats holds the dynamic analysis results
type ImageStats struct {
	ImageTag               string
	Architecture           string
	SupportedArchitectures []string // from Manifest Inspect
	OS                     string
	OSDistro               string // from Syft (e.g., "Alpine Linux 3.18")
	SizeMB                 string
	TotalLayers            int
	Efficiency             float64 // from Dive (0-100)
	WastedBytes            string  // from Dive
	TotalPackages          int
	Packages               []PackageSummary // from Syft (Key Frameworks only)
	Vulnerabilities        []Vulnerability  // from Grype (Sorted by severity)
	VulnSummary            map[string]int   // from Grype (Severity -> Count)
}

// Badge Helpers

func (s *ImageStats) SizeBadge() string {
	if s.SizeMB == "" {
		return ""
	}
	// https://img.shields.io/static/v1?label=Size&message=7.6MB&color=blue
	return fmt.Sprintf("https://img.shields.io/static/v1?label=Size&message=%s&color=blue", url.QueryEscape(s.SizeMB))
}

func (s *ImageStats) LayersBadge() string {
	if s.TotalLayers == 0 {
		return ""
	}
	return fmt.Sprintf("https://img.shields.io/static/v1?label=Layers&message=%d&color=blue", s.TotalLayers)
}

func (s *ImageStats) EfficiencyBadge() string {
	if s.Efficiency == 0 {
		return ""
	}
	color := "green"
	if s.Efficiency < 90 {
		color = "yellow"
	}
	if s.Efficiency < 80 {
		color = "red"
	}
	return fmt.Sprintf("https://img.shields.io/static/v1?label=Efficiency&message=%.1f%%&color=%s", s.Efficiency, color)
}

func (s *ImageStats) VulnBadge() string {
	critical := s.VulnSummary["Critical"]
	high := s.VulnSummary["High"]
	total := critical + high + s.VulnSummary["Medium"] + s.VulnSummary["Low"]

	color := "green"
	if critical > 0 {
		color = "red"
	} else if high > 0 {
		color = "orange"
	}

	msg := fmt.Sprintf("%d Vulns (%d Crit)", total, critical)
	return fmt.Sprintf("https://img.shields.io/static/v1?label=Security&message=%s&color=%s", url.QueryEscape(msg), color)
}

func (s *ImageStats) TotalVulns() int {
	return len(s.Vulnerabilities)
}
