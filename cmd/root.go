package cmd

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// stdout is the writer used for normal program output. Tests can swap this
// to capture output without touching os.Stdout, which makes it safe to run
// tests that capture output without global process-state mutation.
var stdout io.Writer = os.Stdout

// logOutput is the writer used for structured log output (slog). Tests can
// swap this to capture log messages. Defaults to os.Stderr.
var logOutput io.Writer = os.Stderr

var (
	dockerfile       string
	outputFile       string
	dryRun           bool
	imageTag         string
	configFile       string
	noMoji           bool
	ignoreErrors     bool
	verbose          bool
	badgeBaseURL     string
	templateName     string
	listTemplates    bool
	exportTemplate   string
	validateTemplate string
	debugTemplate    bool
	analysisTimeout  time.Duration
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
- YAML Mode: Uses 'dock-docs.yaml' for advanced configuration and multi-image comparison analysis.
- CLI Mode: Runs on a single Dockerfile/image without configuration.`,
	Example: `  # YAML Mode (Recommended)
  dock-docs --config dock-docs.yaml

  # CLI Mode: Analyze Dockerfile only
  dock-docs -f ./Dockerfile

  # CLI Mode: Analyze Dockerfile and Image
  dock-docs -f ./Dockerfile --image my-app:latest

  # CLI Mode: Output to specific file
  dock-docs -f ./Dockerfile -o DOCUMENTATION.md`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initLogger()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// Handle template developer tool flags first (these exit early)
		if listTemplates {
			return handleListTemplates()
		}
		if exportTemplate != "" {
			return handleExportTemplate(exportTemplate)
		}
		if validateTemplate != "" {
			return handleValidateTemplate(validateTemplate)
		}

		ctx := cmd.Context()

		// Apply analysis timeout if set
		if analysisTimeout > 0 {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, analysisTimeout)
			defer cancel()
		}

		// Detect YAML Mode
		cfgPath := configFile
		if cfgPath == "" {
			if _, err := os.Stat("dock-docs.yaml"); err == nil {
				cfgPath = "dock-docs.yaml"
			}
		}

		if cfgPath != "" {
			return runYAMLMode(ctx, cfgPath)
		}

		return runCLIMode(ctx)
	},
}

// Execute runs the root cobra command and exits on error.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// initLogger configures the global slog logger based on the --verbose flag.
// DEBUG level is enabled when verbose is true; otherwise INFO is the minimum.
// Output goes to logOutput (defaults to stderr) so it doesn't interfere with
// program content on stdout.
func initLogger() {
	level := slog.LevelInfo
	if verbose {
		level = slog.LevelDebug
	}
	handler := slog.NewTextHandler(logOutput, &slog.HandlerOptions{
		Level: level,
	})
	slog.SetDefault(slog.New(handler))
}

func init() {
	// Dynamically append tool status to the help description
	rootCmd.Long += "\n" + checkToolStatus()

	rootCmd.Flags().StringVarP(&dockerfile, "file", "f", "./Dockerfile", "Path to Dockerfile (CLI Mode only)")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "README.md", "Path to output file")
	rootCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Print to stdout instead of writing to file")
	rootCmd.Flags().StringVar(&imageTag, "image", "", "Docker image tag to analyze (e.g. my-app:latest) (CLI Mode only)")
	rootCmd.Flags().StringVar(&configFile, "config", "", "Path to config file (default: dock-docs.yaml)")
	rootCmd.Flags().BoolVar(&noMoji, "nomoji", false, "Disable emojis in the output")
	rootCmd.Flags().BoolVar(&ignoreErrors, "ignore-errors", false, "Ignore analysis errors and continue (default false)")
	rootCmd.Flags().BoolVar(&verbose, "verbose", false, "Enable verbose logging")
	rootCmd.Flags().StringVar(&badgeBaseURL, "badge-base-url", "https://img.shields.io/static/v1", "Base URL for badge generation (e.g. for self-hosted shields.io)")

	// Template flags
	rootCmd.Flags().StringVar(&templateName, "template", "", "Template to use (built-in name or file path)")
	rootCmd.Flags().BoolVar(&listTemplates, "list-templates", false, "List all available built-in templates")
	rootCmd.Flags().StringVar(&exportTemplate, "export-template", "", "Export a built-in template to stdout (e.g. 'default')")
	rootCmd.Flags().StringVar(&validateTemplate, "validate-template", "", "Validate a custom template file for syntax errors")
	rootCmd.Flags().BoolVar(&debugTemplate, "debug-template", false, "Print template resolution info during rendering")
	rootCmd.Flags().DurationVar(&analysisTimeout, "timeout", 10*time.Minute, "Overall timeout for all analysis operations (e.g. 5m, 30s)")

	// Add version flag as shortcut for "version" command
	rootCmd.Version = Version
	rootCmd.SetVersionTemplate("dock-docs {{.Version}}\n")
}
