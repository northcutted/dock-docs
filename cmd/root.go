package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"docker-docs/pkg/analysis"
	"docker-docs/pkg/parser"
	"docker-docs/pkg/renderer"
	"docker-docs/pkg/runner"
)

var (
	dockerfile string
	outputFile string
	dryRun     bool
	imageTag   string
)

const (
	markerBegin = "<!-- BEGIN: docker-docs -->"
	markerEnd   = "<!-- END: docker-docs -->"
)

var rootCmd = &cobra.Command{
	Use:   "docker-docs",
	Short: "Generate documentation from Dockerfile",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 1. Parse Dockerfile
		doc, err := parser.Parse(dockerfile)
		if err != nil {
			return fmt.Errorf("failed to parse Dockerfile: %w", err)
		}

		// 2. Dynamic Analysis (if requested)
		var stats *analysis.ImageStats
		if imageTag != "" {
			fmt.Printf("Analyzing image: %s ...\n", imageTag)
			runners := []analysis.Runner{
				&runner.RuntimeRunner{}, // Uses docker or podman
				&runner.SyftRunner{},
				&runner.GrypeRunner{},
				&runner.DiveRunner{},
			}
			stats, err = analysis.AnalyzeImage(imageTag, runners)
			if err != nil {
				// Spec says "log warning but do not fail".
				// AnalyzeImage already logs warnings and skips.
				// But if AnalyzeImage returns error, it might be structural?
				// Actually AnalyzeImage returns error only if input is bad.
				// But let's log and proceed.
				fmt.Printf("Warning: analysis failed: %v\n", err)
			}
		}

		// 3. Render
		renderedContent, err := renderer.Render(doc, stats)
		if err != nil {
			return fmt.Errorf("failed to render documentation: %w", err)
		}

		// 4. Output Strategy
		if dryRun {
			fmt.Println(renderedContent)
			return nil
		}

		// Check if output file exists
		content, err := os.ReadFile(outputFile)
		if err != nil {
			if os.IsNotExist(err) {
				// File doesn't exist -> stdout
				fmt.Println(renderedContent)
				return nil
			}
			return err
		}

		fileContent := string(content)
		startIdx := strings.Index(fileContent, markerBegin)
		endIdx := strings.Index(fileContent, markerEnd)

		if startIdx != -1 && endIdx != -1 && endIdx > startIdx {
			// Markers found, inject
			// We need to keep the markers!
			newContent := fileContent[:startIdx] + markerBegin + "\n" + renderedContent + "\n" + markerEnd + fileContent[endIdx+len(markerEnd):]

			if err := os.WriteFile(outputFile, []byte(newContent), 0644); err != nil {
				return fmt.Errorf("failed to write output file: %w", err)
			}
			fmt.Printf("Updated %s\n", outputFile)
		} else {
			// No markers found -> stdout
			fmt.Println(renderedContent)
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&dockerfile, "file", "f", "./Dockerfile", "Path to Dockerfile")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "README.md", "Path to output file")
	rootCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Print to stdout instead of writing to file")
	rootCmd.Flags().StringVar(&imageTag, "image", "", "Docker image tag to analyze (e.g. my-app:latest)")
}
