package analysis

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/northcutted/dock-docs/pkg/runner"
	"github.com/northcutted/dock-docs/pkg/types"
	"golang.org/x/sync/errgroup"
)

// Runner defines the interface for external tool integration used during image analysis.
type Runner interface {
	Name() string
	IsAvailable() bool
	Run(image string, verbose bool) (*types.ImageStats, error)
}

// AnalyzeComparison runs analysis on multiple images in parallel.
func AnalyzeComparison(images []string, runners []Runner, verbose bool) ([]*types.ImageStats, error) {
	var g errgroup.Group

	// Create a slice to hold results
	results := make([]*types.ImageStats, len(images))

	for i, img := range images {
		// Capture loop variables
		i, img := i, img
		g.Go(func() error {
			// Run analysis for this image
			stats, err := AnalyzeImage(img, runners, verbose)
			if err != nil {
				fmt.Printf("Comparison Analysis Failed for %s: %v\n", img, err)
				if stats != nil {
					results[i] = stats
				}
				return nil // Don't fail the group, partial success allowed
			}
			results[i] = stats
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	// Filter out nils if any failed
	var validResults []*types.ImageStats
	for _, res := range results {
		if res != nil {
			validResults = append(validResults, res)
		}
	}

	return validResults, nil
}

// AnalyzeImage runs all available runners and merges results.
// Runners are injected to allow easy testing/mocking or registration.
var ensureImage = runner.EnsureImage

// AnalyzeImage runs all available runners against the given image and merges their results.
func AnalyzeImage(image string, runners []Runner, verbose bool) (*types.ImageStats, error) {
	if image == "" {
		return nil, fmt.Errorf("image tag is required")
	}

	// Ensure the image exists locally before analysis
	if err := ensureImage(image, verbose); err != nil {
		return nil, fmt.Errorf("failed to ensure image %s: %w", image, err)
	}

	finalStats := &types.ImageStats{
		ImageTag:        image,
		VulnSummary:     make(map[string]int),
		Packages:        make([]types.PackageSummary, 0),
		Vulnerabilities: make([]types.Vulnerability, 0),
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	errChan := make(chan error, len(runners))

	for _, r := range runners {
		if !r.IsAvailable() {
			if verbose {
				fmt.Printf("Warning: %s is not installed or not in PATH. Skipping.\n", r.Name())
			}
			continue
		}

		wg.Add(1)
		go func(runner Runner) {
			defer wg.Done()
			stats, err := runner.Run(image, verbose)
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
	var errs []error
	for err := range errChan {
		fmt.Printf("Analysis Warning: %v\n", err)
		errs = append(errs, err)
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

	if len(errs) > 0 {
		// Build a detailed error message listing failing runners
		var errMsgs []string
		for _, e := range errs {
			errMsgs = append(errMsgs, e.Error())
		}
		return finalStats, fmt.Errorf("analysis failed for %d runners: %s", len(errs), strings.Join(errMsgs, "; "))
	}

	return finalStats, nil
}

func mergeStats(dest, src *types.ImageStats) {
	if src == nil {
		return
	}

	if src.Architecture != "" {
		dest.Architecture = src.Architecture
	}
	if src.OS != "" {
		dest.OS = src.OS
	}
	if src.OSDistro != "" {
		dest.OSDistro = src.OSDistro
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
	if len(src.SupportedArchitectures) > 0 {
		dest.SupportedArchitectures = src.SupportedArchitectures
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
	if !src.VulnScanTime.IsZero() {
		dest.VulnScanTime = src.VulnScanTime
	}
}
