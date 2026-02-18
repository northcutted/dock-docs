package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/northcutted/dock-docs/pkg/analysis"
	"github.com/northcutted/dock-docs/pkg/config"
	"github.com/northcutted/dock-docs/pkg/injector"
	"github.com/northcutted/dock-docs/pkg/parser"
	"github.com/northcutted/dock-docs/pkg/renderer"
	"github.com/northcutted/dock-docs/pkg/runner"
	"github.com/northcutted/dock-docs/pkg/templates"
	"github.com/northcutted/dock-docs/pkg/types"
)

func runYAMLMode(path string) error {
	fmt.Printf("Using config file: %s\n", path)
	cfg, err := config.Load(path)
	if err != nil {
		return err
	}

	// Change working directory to config file location
	// This ensures relative paths in config (like output, source) are resolved correctly
	configDir := filepath.Dir(path)
	if configDir != "." {
		if err := os.Chdir(configDir); err != nil {
			return fmt.Errorf("failed to change directory to %s: %w", configDir, err)
		}
		fmt.Printf("Changed working directory to: %s\n", configDir)
	}

	// Read Output File (only needed for markdown injection; read lazily below)
	var fileContent string
	var fileContentLoaded bool

	loadFileContent := func() error {
		if fileContentLoaded {
			return nil
		}
		content, err := os.ReadFile(cfg.Output)
		if err != nil {
			return fmt.Errorf("failed to read output file %s: %w", cfg.Output, err)
		}
		fileContent = string(content)
		fileContentLoaded = true
		return nil
	}

	// Process Sections
	runners := []analysis.Runner{
		&runner.RuntimeRunner{},
		&runner.ManifestRunner{},
		&runner.SyftRunner{},
		&runner.GrypeRunner{},
		&runner.DiveRunner{},
	}

	renderOpts := renderer.RenderOptions{
		NoMoji:       noMoji,
		BadgeBaseURL: cfg.BadgeBaseURL,
	}

	for i, section := range cfg.Sections {
		var sectionContent string

		// Resolve template: CLI flag > section config > global config > default
		tmplSel := resolveTemplateSel(cfg.ResolveTemplate(section))
		format := tmplSel.Format()

		switch section.Type {
		case config.SectionTypeImage:
			// Parse Dockerfile
			dPath := section.Source
			if dPath == "" {
				dPath = "Dockerfile" // Default
			}
			doc, err := parser.Parse(dPath)
			if err != nil {
				return fmt.Errorf("failed to parse Dockerfile %s: %w", dPath, err)
			}

			// Analyze Image (optional)
			var stats *types.ImageStats
			if section.Tag != "" {
				fmt.Printf("Analyzing image: %s ...\n", section.Tag)
				stats, err = analysis.AnalyzeImage(section.Tag, runners, verbose)
				if err != nil {
					fmt.Printf("Warning: analysis failed for %s: %v\n", section.Tag, err)
					if !ignoreErrors {
						return fmt.Errorf("analysis failed for %s: %w", section.Tag, err)
					}
				}
			}

			if debugTemplate {
				fmt.Printf("Template: %s (type: image, format: %s)\n", describeTemplate(tmplSel), format)
			}

			// Render
			sectionContent, err = renderer.RenderWithTemplate(doc, stats, renderOpts, tmplSel)
			if err != nil {
				return fmt.Errorf("failed to render image section: %w", err)
			}

		case config.SectionTypeComparison:
			if len(section.Images) == 0 {
				continue
			}

			// Extract tags from ImageEntry structs for analysis
			resolvedImages := section.ResolvedImages()
			tags := make([]string, len(resolvedImages))
			for j, entry := range resolvedImages {
				tags[j] = entry.Tag
			}

			fmt.Printf("Analyzing comparison: %v ...\n", tags)
			statsList, err := analysis.AnalyzeComparison(tags, runners, verbose)
			if err != nil {
				return fmt.Errorf("comparison analysis failed: %w", err)
			}

			if debugTemplate {
				fmt.Printf("Template: %s (type: comparison, format: %s)\n", describeTemplate(tmplSel), format)
			}

			sectionContent, err = renderer.RenderComparisonWithTemplate(statsList, renderOpts, tmplSel)
			if err != nil {
				return fmt.Errorf("failed to render comparison section: %w", err)
			}

		default:
			fmt.Printf("Warning: unknown section type %s\n", section.Type)
			continue
		}

		// Output: direct-write for html/json, inject for markdown
		if templates.IsDirectWriteFormat(format) {
			outPath := resolveSectionOutput(cfg.Output, section.Marker, i, format)

			if dryRun {
				fmt.Printf("--- %s ---\n", outPath)
				fmt.Println(sectionContent)
				continue
			}

			if err := os.WriteFile(outPath, []byte(sectionContent), 0644); err != nil {
				return fmt.Errorf("failed to write output file %s: %w", outPath, err)
			}
			fmt.Printf("Wrote %s\n", outPath)
		} else {
			// Markdown: inject into existing file between markers
			if err := loadFileContent(); err != nil {
				return err
			}
			newContent, err := injector.Inject(fileContent, section.Marker, sectionContent)
			if err != nil {
				fmt.Printf("Warning: %v\n", err)
				continue
			}
			fileContent = newContent
		}
	}

	// Write the markdown output file if we modified it
	if fileContentLoaded {
		if dryRun {
			fmt.Println(fileContent)
			return nil
		}

		if err := os.WriteFile(cfg.Output, []byte(fileContent), 0644); err != nil {
			return fmt.Errorf("failed to write output file: %w", err)
		}
		fmt.Printf("Updated %s\n", cfg.Output)
	}

	return nil
}
