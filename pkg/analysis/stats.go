package analysis

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
	ImageTag        string
	Architecture    string
	OS              string
	SizeMB          string
	TotalLayers     int
	Efficiency      float64 // from Dive (0-100)
	WastedBytes     string  // from Dive
	TotalPackages   int
	Packages        []PackageSummary // from Syft (Key Frameworks only)
	Vulnerabilities []Vulnerability  // from Grype (Sorted by severity)
	VulnSummary     map[string]int   // from Grype (Severity -> Count)
}
