package analysis

import (
	"context"
	"fmt"
	"log/slog"
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
	Run(ctx context.Context, image string, verbose bool) (*types.ImageStats, error)
}

// AnalyzeComparison runs analysis on multiple images in parallel.
// The provided context controls the overall deadline for the comparison;
// individual runner timeouts are derived from this parent context.
func AnalyzeComparison(ctx context.Context, images []string, runners []Runner, verbose bool) ([]*types.ImageStats, error) {
	var g errgroup.Group

	// Create a slice to hold results
	results := make([]*types.ImageStats, len(images))

	for i, img := range images {
		// Capture loop variables
		i, img := i, img
		g.Go(func() error {
			// Run analysis for this image
			stats, err := AnalyzeImage(ctx, img, runners, verbose)
			if err != nil {
				slog.Warn("comparison analysis failed", "image", img, "error", err)
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
// The provided context controls the overall deadline; individual runner timeouts
// are derived from this parent context.
func AnalyzeImage(ctx context.Context, image string, runners []Runner, verbose bool) (*types.ImageStats, error) {
	if image == "" {
		return nil, fmt.Errorf("image tag is required")
	}

	// Ensure the image exists locally before analysis
	if err := ensureImage(ctx, image, verbose); err != nil {
		return nil, fmt.Errorf("failed to ensure image %s: %w", image, err)
	}

	finalStats := &types.ImageStats{
		ImageTag:        image,
		VulnSummary:     make(map[string]int),
		Packages:        make([]types.PackageSummary, 0),
		Vulnerabilities: make([]types.Vulnerability, 0),
	}

	var g errgroup.Group
	var mu sync.Mutex
	var errs []error

	for _, r := range runners {
		if !r.IsAvailable() {
			if verbose {
				slog.Debug("tool not available, skipping", "runner", r.Name())
			}
			continue
		}

		r := r // capture loop variable
		g.Go(func() error {
			stats, err := r.Run(ctx, image, verbose)
			if err != nil {
				mu.Lock()
				errs = append(errs, fmt.Errorf("%s failed: %w", r.Name(), err))
				mu.Unlock()
				return nil // Don't fail the group; partial success is allowed
			}

			mu.Lock()
			mergeStats(finalStats, stats)
			mu.Unlock()
			return nil
		})
	}

	// Wait for all goroutines (always returns nil since goroutines never return errors)
	_ = g.Wait()

	// Log collected warnings
	for _, err := range errs {
		slog.Warn("analysis runner failed", "error", err)
	}

	// Final sort of vulnerabilities after merge
	types.SortBySeverity(finalStats.Vulnerabilities)

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
	if src.SizeBytes != 0 {
		dest.SizeBytes = src.SizeBytes
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
