package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/northcutted/dock-docs/pkg/analysis"
	"github.com/northcutted/dock-docs/pkg/injector"
	"github.com/northcutted/dock-docs/pkg/parser"
	"github.com/northcutted/dock-docs/pkg/renderer"
	"github.com/northcutted/dock-docs/pkg/runner"
	"github.com/northcutted/dock-docs/pkg/templates"
	"github.com/northcutted/dock-docs/pkg/types"
)

func runCLIMode(ctx context.Context) error {
	// 1. Parse Dockerfile
	doc, err := parser.Parse(dockerfile)
	if err != nil {
		return fmt.Errorf("failed to parse Dockerfile: %w", err)
	}

	// 2. Dynamic Analysis (if requested)
	var stats *types.ImageStats
	if imageTag != "" {
		slog.Info("analyzing image", "image", imageTag)
		runners := []analysis.Runner{
			&runner.RuntimeRunner{},
			&runner.ManifestRunner{},
			&runner.SyftRunner{},
			&runner.GrypeRunner{},
			&runner.DiveRunner{},
		}
		stats, err = analysis.AnalyzeImage(ctx, imageTag, runners, verbose)
		if err != nil {
			slog.Warn("analysis failed", "error", err)
			if !ignoreErrors {
				return fmt.Errorf("analysis failed: %w", err)
			}
		}
	}

	// 3. Resolve template selection: CLI flag > default
	tmplSel := resolveTemplateSel(nil)

	if debugTemplate {
		slog.Debug("template resolved", "template", describeTemplate(tmplSel), "type", "image", "format", tmplSel.Format())
	}

	// 4. Render
	renderOpts := renderer.RenderOptions{
		NoMoji:       noMoji,
		BadgeBaseURL: badgeBaseURL,
	}
	renderedContent, err := renderer.RenderWithTemplate(doc, stats, renderOpts, tmplSel)
	if err != nil {
		return fmt.Errorf("failed to render documentation: %w", err)
	}

	// 5. Output Strategy
	if dryRun {
		fmt.Fprintln(stdout, renderedContent)
		return nil
	}

	format := tmplSel.Format()

	// For HTML/JSON: write the complete standalone document directly to a file
	if templates.IsDirectWriteFormat(format) {
		outPath := resolveOutputPath(outputFile, format)
		if err := os.WriteFile(outPath, []byte(renderedContent), 0644); err != nil {
			return fmt.Errorf("failed to write output file: %w", err)
		}
		slog.Info("wrote output file", "path", outPath)
		return nil
	}

	// For Markdown: inject into existing file between markers
	content, err := os.ReadFile(outputFile)
	if err != nil {
		if os.IsNotExist(err) {
			slog.Warn("output file does not exist, printing to stdout", "file", outputFile)
			fmt.Fprintln(stdout, renderedContent)
			return nil
		}
		return err
	}

	fileContent := string(content)
	// Simple mode uses default markers (empty name)
	newContent, err := injector.Inject(fileContent, "", renderedContent)
	if err != nil {
		slog.Warn("injection failed, printing to stdout", "error", err)
		fmt.Fprintln(stdout, renderedContent)
		return nil
	}

	if err := os.WriteFile(outputFile, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}
	slog.Info("updated output file", "path", outputFile)

	return nil
}
