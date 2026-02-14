package analysis

import (
	"fmt"
	"sort"
	"sync"
)

type Runner interface {
	Name() string
	IsAvailable() bool
	Run(image string) (*ImageStats, error)
}

// AnalyzeImage runs all available runners and merges results.
// Runners are injected to allow easy testing/mocking or registration.
func AnalyzeImage(image string, runners []Runner) (*ImageStats, error) {
	if image == "" {
		return nil, fmt.Errorf("image tag is required")
	}

	finalStats := &ImageStats{
		ImageTag:        image,
		VulnSummary:     make(map[string]int),
		Packages:        make([]PackageSummary, 0),
		Vulnerabilities: make([]Vulnerability, 0),
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	errChan := make(chan error, len(runners))

	for _, r := range runners {
		if !r.IsAvailable() {
			fmt.Printf("Warning: %s is not installed or not in PATH. Skipping.\n", r.Name())
			continue
		}

		wg.Add(1)
		go func(runner Runner) {
			defer wg.Done()
			stats, err := runner.Run(image)
			if err != nil {
				errChan <- fmt.Errorf("%s failed: %w", runner.Name(), err)
				return
			}

			mu.Lock()
			defer mu.Unlock()
			mergeStats(finalStats, stats)
		}(r)
	}

	wg.Wait()
	close(errChan)

	// Collect errors if any (logging or returning partial success?)
	// Spec implies "log a warning but do not fail".
	for err := range errChan {
		fmt.Printf("Analysis Warning: %v\n", err)
	}

	// Final sort of vulnerabilities after merge
	severityRank := map[string]int{
		"Critical": 4,
		"High":     3,
		"Medium":   2,
		"Low":      1,
		"Unknown":  0,
	}
	sort.Slice(finalStats.Vulnerabilities, func(i, j int) bool {
		rankI := severityRank[finalStats.Vulnerabilities[i].Severity]
		rankJ := severityRank[finalStats.Vulnerabilities[j].Severity]
		if rankI != rankJ {
			return rankI > rankJ
		}
		return finalStats.Vulnerabilities[i].ID < finalStats.Vulnerabilities[j].ID
	})

	return finalStats, nil
}

func mergeStats(dest, src *ImageStats) {
	if src == nil {
		return
	}

	if src.Architecture != "" {
		dest.Architecture = src.Architecture
	}
	if src.OS != "" {
		dest.OS = src.OS
	}
	if src.SizeMB != "" {
		dest.SizeMB = src.SizeMB
	}
	if src.TotalLayers != 0 {
		dest.TotalLayers = src.TotalLayers
	}
	if src.Efficiency != 0 {
		dest.Efficiency = src.Efficiency
	}
	if src.WastedBytes != "" {
		dest.WastedBytes = src.WastedBytes
	}
	if src.TotalPackages != 0 {
		dest.TotalPackages = src.TotalPackages
	}
	if len(src.Packages) > 0 {
		dest.Packages = append(dest.Packages, src.Packages...)
	}
	if len(src.Vulnerabilities) > 0 {
		dest.Vulnerabilities = append(dest.Vulnerabilities, src.Vulnerabilities...)
	}
	if len(src.VulnSummary) > 0 {
		for k, v := range src.VulnSummary {
			dest.VulnSummary[k] += v
		}
	}
}
