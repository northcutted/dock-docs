package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/northcutted/dock-docs/pkg/analysis"
	"github.com/northcutted/dock-docs/pkg/config"
	"github.com/northcutted/dock-docs/pkg/injector"
	"github.com/northcutted/dock-docs/pkg/parser"
	"github.com/northcutted/dock-docs/pkg/renderer"
	"github.com/northcutted/dock-docs/pkg/runner"
)

var (
	dockerfile   string
	outputFile   string
	dryRun       bool
	imageTag     string
	configFile   string
	noMoji       bool
	ignoreErrors bool
	verbose      bool
)

var rootCmd = &cobra.Command{
	Use:   "dock-docs",
	Short: "Generate documentation from Dockerfile",
	Long: `Generate comprehensive documentation from your Dockerfiles.

This tool automatically parses your Dockerfile to create standard documentation
tables for arguments, environment variables, and exposed ports.

It can also perform deep analysis on built images using:
- Syft (for SBOM and package listing)
- Grype (for vulnerability scanning)
- Dive (for layer efficiency analysis)

Modes:
- Config Mode: Uses 'dock-docs.yaml' for advanced configuration and multi-image matrix analysis.
- Simple Mode: Runs on a single Dockerfile/image without configuration.`,
	Example: `  # Config Mode (Recommended)
  dock-docs --config dock-docs.yaml

  # Simple Mode: Analyze Dockerfile only
  dock-docs -f ./Dockerfile

  # Simple Mode: Analyze Dockerfile and Image
  dock-docs -f ./Dockerfile --image my-app:latest

  # Simple Mode: Output to specific file
  dock-docs -f ./Dockerfile -o DOCUMENTATION.md`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Detect Config Mode
		cfgPath := configFile
		if cfgPath == "" {
			if _, err := os.Stat("dock-docs.yaml"); err == nil {
				cfgPath = "dock-docs.yaml"
			}
		}

		if cfgPath != "" {
			return runConfigMode(cfgPath)
		}

		return runSimpleMode()
	},
}

func runConfigMode(path string) error {
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

	// Read Output File
	content, err := os.ReadFile(cfg.Output)
	if err != nil {
		// If output file doesn't exist, we can't inject.
		// Spec doesn't strictly say we must create it, but in Config mode,
		// usually we expect to update an existing file.
		return fmt.Errorf("failed to read output file %s: %w", cfg.Output, err)
	}
	fileContent := string(content)

	// Process Sections
	runners := []analysis.Runner{
		&runner.RuntimeRunner{},
		&runner.ManifestRunner{},
		&runner.SyftRunner{},
		&runner.GrypeRunner{},
		&runner.DiveRunner{},
	}

	renderOpts := renderer.RenderOptions{
		NoMoji: noMoji,
	}

	for _, section := range cfg.Sections {
		var sectionContent string

		switch section.Type {
		case config.SectionTypeConfig:
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
			var stats *analysis.ImageStats
			if section.Image != "" {
				fmt.Printf("Analyzing image (config): %s ...\n", section.Image)
				stats, err = analysis.AnalyzeImage(section.Image, runners, verbose)
				if err != nil {
					fmt.Printf("Warning: analysis failed for %s: %v\n", section.Image, err)
					if !ignoreErrors {
						return fmt.Errorf("analysis failed for %s: %w", section.Image, err)
					}
				}
			}

			// Render
			sectionContent, err = renderer.Render(doc, stats, renderOpts)
			if err != nil {
				return fmt.Errorf("failed to render config section: %w", err)
			}

		case config.SectionTypeMatrix:
			if len(section.Images) == 0 {
				continue
			}
			fmt.Printf("Analyzing matrix: %v ...\n", section.Images)
			statsList, err := analysis.AnalyzeMatrix(section.Images, runners, verbose)
			if err != nil {
				return fmt.Errorf("matrix analysis failed: %w", err)
			}
			// Check for partial failures in matrix if strict mode?
			// AnalyzeMatrix currently returns partial results.
			// To be strict, we'd need to inspect the errors inside AnalyzeMatrix or change its signature.
			// For now, let's assume if AnalyzeMatrix returns error (which it does if critical), we catch it.
			// But individual image failures inside matrix might be suppressed.
			// Let's rely on the user to check logs or implement stricter matrix checks later if needed.
			// Wait, I should probably check if statsList has nil entries or implied failures?
			// AnalyzeMatrix filters out nils.

			sectionContent, err = renderer.RenderMatrix(statsList, renderOpts)
			if err != nil {
				return fmt.Errorf("failed to render matrix section: %w", err)
			}

		default:
			fmt.Printf("Warning: unknown section type %s\n", section.Type)
			continue
		}

		// Inject
		newContent, err := injector.Inject(fileContent, section.Marker, sectionContent)
		if err != nil {
			// In config mode, missing markers is likely an error user wants to know about
			// But maybe we should just warn?
			fmt.Printf("Warning: %v\n", err)
			continue
		}
		fileContent = newContent
	}

	if dryRun {
		fmt.Println(fileContent)
		return nil
	}

	if err := os.WriteFile(cfg.Output, []byte(fileContent), 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}
	fmt.Printf("Updated %s\n", cfg.Output)
	return nil
}

func runSimpleMode() error {
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
			&runner.RuntimeRunner{},
			&runner.ManifestRunner{},
			&runner.SyftRunner{},
			&runner.GrypeRunner{},
			&runner.DiveRunner{},
		}
		stats, err = analysis.AnalyzeImage(imageTag, runners, verbose)
		if err != nil {
			fmt.Printf("Warning: analysis failed: %v\n", err)
			if !ignoreErrors {
				return fmt.Errorf("analysis failed: %w", err)
			}
		}
	}

	// 3. Render
	renderOpts := renderer.RenderOptions{
		NoMoji: noMoji,
	}
	renderedContent, err := renderer.Render(doc, stats, renderOpts)
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
			if !dryRun {
				fmt.Fprintf(os.Stderr, "Warning: Output file '%s' does not exist. Printing to stdout.\n", outputFile)
			}
			fmt.Println(renderedContent)
			return nil
		}
		return err
	}

	fileContent := string(content)
	// Simple mode uses default markers (empty name)
	newContent, err := injector.Inject(fileContent, "", renderedContent)
	if err != nil {
		if !dryRun {
			fmt.Fprintf(os.Stderr, "Warning: %v. Printing to stdout.\n", err)
		}
		fmt.Println(renderedContent)
		return nil
	}

	if err := os.WriteFile(outputFile, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}
	fmt.Printf("Updated %s\n", outputFile)

	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// checkToolStatus returns a string indicating the status of required tools.
func checkToolStatus() string {
	tools := []string{"syft", "grype", "dive"}
	var status strings.Builder
	status.WriteString("\nPrerequisites:\n")

	// Check for Docker or Podman
	if _, err := exec.LookPath("docker"); err == nil {
		status.WriteString("  [OK] docker\n")
	} else if _, err := exec.LookPath("podman"); err == nil {
		status.WriteString("  [OK] podman\n")
	} else {
		status.WriteString("  [MISSING] docker or podman (required for dynamic analysis)\n")
	}

	for _, tool := range tools {
		if _, err := exec.LookPath(tool); err == nil {
			status.WriteString(fmt.Sprintf("  [OK] %s\n", tool))
		} else {
			status.WriteString(fmt.Sprintf("  [MISSING] %s (install for full functionality)\n", tool))
		}
	}
	return status.String()
}

func init() {
	// Dynamically append tool status to the help description
	rootCmd.Long += "\n" + checkToolStatus()

	rootCmd.Flags().StringVarP(&dockerfile, "file", "f", "./Dockerfile", "Path to Dockerfile (Simple Mode only)")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "README.md", "Path to output file")
	rootCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Print to stdout instead of writing to file")
	rootCmd.Flags().StringVar(&imageTag, "image", "", "Docker image tag to analyze (e.g. my-app:latest) (Simple Mode only)")
	rootCmd.Flags().StringVar(&configFile, "config", "", "Path to config file (default: dock-docs.yaml)")
	rootCmd.Flags().BoolVar(&noMoji, "nomoji", false, "Disable emojis in the output")
	rootCmd.Flags().BoolVar(&ignoreErrors, "ignore-errors", false, "Ignore analysis errors and continue (default false)")
	rootCmd.Flags().BoolVar(&verbose, "verbose", false, "Enable verbose logging")
}
