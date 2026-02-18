package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/northcutted/dock-docs/pkg/analysis"
	"github.com/northcutted/dock-docs/pkg/config"
	"github.com/northcutted/dock-docs/pkg/injector"
	"github.com/northcutted/dock-docs/pkg/parser"
	"github.com/northcutted/dock-docs/pkg/renderer"
	"github.com/northcutted/dock-docs/pkg/runner"
	"github.com/northcutted/dock-docs/pkg/templates"
	"github.com/northcutted/dock-docs/pkg/types"
)

// newRunners creates a fresh set of analysis runners.
// Each caller gets its own instances to avoid data races on mutable state
// (e.g., RuntimeRunner.binary is set by IsAvailable).
func newRunners() []analysis.Runner {
	return []analysis.Runner{
		&runner.RuntimeRunner{},
		&runner.ManifestRunner{},
		&runner.SyftRunner{},
		&runner.GrypeRunner{},
		&runner.DiveRunner{},
	}
}

// sectionResult holds the rendered content and metadata for a processed section.
type sectionResult struct {
	index   int    // original section index (for resolveSectionOutput)
	content string // rendered output
	format  string // template format (markdown, html, json)
	marker  string // section marker for injection
}

// indexedSection pairs a config section with its resolved template and index.
type indexedSection struct {
	index   int
	section config.Section
	tmplSel renderer.TemplateSelection
	format  string
}

func runYAMLMode(ctx context.Context, path string) error {
	slog.Info("using config file", "path", path)
	cfg, err := config.Load(path)
	if err != nil {
		return err
	}

	// Resolve relative paths in config against the config file's directory.
	// This avoids mutating global process state with os.Chdir while still
	// allowing relative paths in the YAML config to work correctly.
	configDir := filepath.Dir(path)
	cfg.ResolveRelativePaths(configDir)

	renderOpts := renderer.RenderOptions{
		NoMoji:       noMoji,
		BadgeBaseURL: cfg.BadgeBaseURL,
	}

	// Partition sections into direct-write (html/json) and markdown-inject groups.
	// Direct-write sections write to independent files and can run in parallel.
	// Markdown sections inject into a shared output file and must be sequential.

	var directWriteSections []indexedSection
	var markdownSections []indexedSection

	for i, section := range cfg.Sections {
		tmplSel := resolveTemplateSel(cfg.ResolveTemplate(section))
		format := tmplSel.Format()

		is := indexedSection{index: i, section: section, tmplSel: tmplSel, format: format}
		if templates.IsDirectWriteFormat(format) {
			directWriteSections = append(directWriteSections, is)
		} else {
			markdownSections = append(markdownSections, is)
		}
	}

	// Process direct-write sections in parallel using errgroup.
	var directResults []sectionResult
	var mu sync.Mutex

	g, gctx := errgroup.WithContext(ctx)
	for _, is := range directWriteSections {
		is := is // capture loop variable
		g.Go(func() error {
			content, err := processSection(gctx, is.section, is.tmplSel, is.format, renderOpts)
			if err != nil {
				return err
			}
			if content == "" {
				return nil
			}
			mu.Lock()
			directResults = append(directResults, sectionResult{
				index:   is.index,
				content: content,
				format:  is.format,
				marker:  is.section.Marker,
			})
			mu.Unlock()
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	// Write direct-write results to independent files
	for _, res := range directResults {
		outPath := resolveSectionOutput(cfg.Output, res.marker, res.index, res.format)

		if dryRun {
			fmt.Fprintf(stdout, "--- %s ---\n", outPath)
			fmt.Fprintln(stdout, res.content)
			continue
		}

		if err := os.WriteFile(outPath, []byte(res.content), 0644); err != nil {
			return fmt.Errorf("failed to write output file %s: %w", outPath, err)
		}
		slog.Info("wrote output file", "path", outPath)
	}

	// Process markdown sections sequentially (they share the output file).
	if len(markdownSections) > 0 {
		if err := processMarkdownSections(ctx, cfg.Output, markdownSections, renderOpts); err != nil {
			return err
		}
	}

	return nil
}

// processMarkdownSections renders markdown sections sequentially and injects
// them into the shared output file. Markdown sections cannot be parallelized
// because they all inject into the same document.
func processMarkdownSections(ctx context.Context, outputPath string, sections []indexedSection, renderOpts renderer.RenderOptions) error {
	content, err := os.ReadFile(outputPath)
	if err != nil {
		return fmt.Errorf("failed to read output file %s: %w", outputPath, err)
	}
	fileContent := string(content)

	for _, is := range sections {
		sectionContent, err := processSection(ctx, is.section, is.tmplSel, is.format, renderOpts)
		if err != nil {
			return err
		}
		if sectionContent == "" {
			continue
		}

		newContent, err := injector.Inject(fileContent, is.section.Marker, sectionContent)
		if err != nil {
			slog.Warn("injection failed for section", "marker", is.section.Marker, "error", err)
			continue
		}
		fileContent = newContent
	}

	if dryRun {
		fmt.Fprintln(stdout, fileContent)
		return nil
	}

	if err := os.WriteFile(outputPath, []byte(fileContent), 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}
	slog.Info("updated output file", "path", outputPath)
	return nil
}

// processSection renders a single config section and returns the rendered content.
// Returns empty string for sections that should be skipped (e.g., empty comparison, unknown type).
func processSection(ctx context.Context, section config.Section, tmplSel renderer.TemplateSelection, format string, renderOpts renderer.RenderOptions) (string, error) {
	runners := newRunners()

	switch section.Type {
	case config.SectionTypeImage:
		// Parse Dockerfile
		dPath := section.Source
		if dPath == "" {
			dPath = "Dockerfile" // Default
		}
		doc, err := parser.Parse(dPath)
		if err != nil {
			return "", fmt.Errorf("failed to parse Dockerfile %s: %w", dPath, err)
		}

		// Analyze Image (optional)
		var stats *types.ImageStats
		if section.Tag != "" {
			slog.Info("analyzing image", "image", section.Tag)
			stats, err = analysis.AnalyzeImage(ctx, section.Tag, runners, verbose)
			if err != nil {
				slog.Warn("analysis failed", "image", section.Tag, "error", err)
				if !ignoreErrors {
					return "", fmt.Errorf("analysis failed for %s: %w", section.Tag, err)
				}
			}
		}

		if debugTemplate {
			slog.Debug("template resolved", "template", describeTemplate(tmplSel), "type", "image", "format", format)
		}

		content, err := renderer.RenderWithTemplate(doc, stats, renderOpts, tmplSel)
		if err != nil {
			return "", fmt.Errorf("failed to render image section: %w", err)
		}
		return content, nil

	case config.SectionTypeComparison:
		if len(section.Images) == 0 {
			return "", nil
		}

		// Extract tags from ImageEntry structs for analysis
		resolvedImages := section.ResolvedImages()
		tags := make([]string, len(resolvedImages))
		for j, entry := range resolvedImages {
			tags[j] = entry.Tag
		}

		slog.Info("analyzing comparison", "images", tags)
		statsList, err := analysis.AnalyzeComparison(ctx, tags, runners, verbose)
		if err != nil {
			return "", fmt.Errorf("comparison analysis failed: %w", err)
		}

		if debugTemplate {
			slog.Debug("template resolved", "template", describeTemplate(tmplSel), "type", "comparison", "format", format)
		}

		content, err := renderer.RenderComparisonWithTemplate(statsList, renderOpts, tmplSel)
		if err != nil {
			return "", fmt.Errorf("failed to render comparison section: %w", err)
		}
		return content, nil

	default:
		slog.Warn("unknown section type", "type", section.Type)
		return "", nil
	}
}
