package types

import (
	"fmt"
	"net/url"
	"sort"
	"time"
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
	SizeBytes              int64
	TotalLayers            int
	Efficiency             float64 // from Dive (0-100)
	WastedBytes            string  // from Dive
	TotalPackages          int
	Packages               []PackageSummary // from Syft (Key Frameworks only)
	Vulnerabilities        []Vulnerability  // from Grype (Sorted by severity)
	VulnSummary            map[string]int   // from Grype (Severity -> Count)
	VulnScanTime           time.Time        // from Grype (When vulnerability scan was performed)
}

// SizeMB returns the image size formatted as a human-readable string (e.g., "7.60 MB").
// Templates can call {{ .SizeMB }} or {{ .Stats.SizeMB }} and get the same result
// as the old string field.
func (s *ImageStats) SizeMB() string {
	if s.SizeBytes == 0 {
		return ""
	}
	return fmt.Sprintf("%.2f MB", float64(s.SizeBytes)/1024/1024)
}

// Badge Helpers

// SizeBadge returns a shields.io badge URL for the image size.
func (s *ImageStats) SizeBadge(baseURL string) string {
	sizeMB := s.SizeMB()
	if sizeMB == "" {
		return ""
	}
	// https://img.shields.io/static/v1?label=Size&message=7.6MB&color=blue
	return fmt.Sprintf("%s?label=Size&message=%s&color=blue", baseURL, url.QueryEscape(sizeMB))
}

// LayersBadge returns a shields.io badge URL for the layer count.
func (s *ImageStats) LayersBadge(baseURL string) string {
	if s.TotalLayers == 0 {
		return ""
	}
	return fmt.Sprintf("%s?label=Layers&message=%d&color=blue", baseURL, s.TotalLayers)
}

// EfficiencyBadge returns a shields.io badge URL for the image efficiency score.
func (s *ImageStats) EfficiencyBadge(baseURL string) string {
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
	return fmt.Sprintf("%s?label=Efficiency&message=%.1f%%&color=%s", baseURL, s.Efficiency, color)
}

// VulnBadge returns a shields.io badge URL summarizing vulnerability findings.
// The total uses TotalVulns() (which includes all severities, including Unknown)
// for consistency with the Vulnerabilities slice.
func (s *ImageStats) VulnBadge(baseURL string) string {
	critical := s.VulnSummary["Critical"]
	high := s.VulnSummary["High"]
	total := s.TotalVulns()

	color := "green"
	if critical > 0 {
		color = "red"
	} else if high > 0 {
		color = "orange"
	}

	msg := fmt.Sprintf("%d Vulns (%d Crit)", total, critical)
	return fmt.Sprintf("%s?label=Security&message=%s&color=%s", baseURL, url.QueryEscape(msg), color)
}

// TotalVulns returns the total number of vulnerabilities found.
func (s *ImageStats) TotalVulns() int {
	return len(s.Vulnerabilities)
}

// severityRank maps vulnerability severity strings to numeric ranks for sorting.
// Higher values indicate more severe vulnerabilities.
var severityRank = map[string]int{
	"Critical": 4,
	"High":     3,
	"Medium":   2,
	"Low":      1,
	"Unknown":  0,
}

// SortBySeverity sorts a slice of vulnerabilities by severity (Critical first)
// with a secondary sort by ID for deterministic ordering.
func SortBySeverity(vulns []Vulnerability) {
	sort.Slice(vulns, func(i, j int) bool {
		rankI := severityRank[vulns[i].Severity]
		rankJ := severityRank[vulns[j].Severity]
		if rankI != rankJ {
			return rankI > rankJ
		}
		return vulns[i].ID < vulns[j].ID
	})
}
