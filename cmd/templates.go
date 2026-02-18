package cmd

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/northcutted/dock-docs/pkg/config"
	"github.com/northcutted/dock-docs/pkg/renderer"
	"github.com/northcutted/dock-docs/pkg/templates"
)

// resolveTemplateSel builds a TemplateSelection from the CLI flag and optional config.
// CLI --template flag takes precedence over config file settings.
func resolveTemplateSel(cfgTemplate *config.TemplateConfig) renderer.TemplateSelection {
	// CLI flag takes precedence
	if templateName != "" {
		// If it looks like a file path (contains / or .tmpl), treat as file
		if strings.Contains(templateName, "/") || strings.HasSuffix(templateName, ".tmpl") {
			return renderer.TemplateSelection{Path: templateName}
		}
		return renderer.TemplateSelection{Name: templateName}
	}

	// Fall back to config file setting
	if cfgTemplate != nil {
		sel := renderer.TemplateSelection{}
		if cfgTemplate.Path != "" {
			sel.Path = cfgTemplate.Path
		} else if cfgTemplate.Name != "" {
			sel.Name = cfgTemplate.Name
		}
		return sel
	}

	// Default: empty selection means "default" built-in
	return renderer.TemplateSelection{}
}

// describeTemplate returns a human-readable description of the template being used.
func describeTemplate(sel renderer.TemplateSelection) string {
	if sel.Path != "" {
		return fmt.Sprintf("custom file: %s", sel.Path)
	}
	if sel.Name != "" {
		return fmt.Sprintf("built-in: %s", sel.Name)
	}
	return "built-in: default"
}

// handleListTemplates prints all available built-in templates.
func handleListTemplates() error { //nolint:unparam // error return is part of RunE handler contract
	builtins := templates.ListBuiltin()
	fmt.Fprintln(stdout, "Available built-in templates:")
	fmt.Fprintln(stdout)
	for _, b := range builtins {
		fmt.Fprintf(stdout, "  %-10s  [%s]  %s\n", b.Name, b.Format, b.Description)
	}
	fmt.Fprintln(stdout)
	fmt.Fprintln(stdout, "Usage:")
	fmt.Fprintln(stdout, "  dock-docs --template <name>")
	fmt.Fprintln(stdout, "  dock-docs --export-template <name> > my-template.tmpl")
	return nil
}

// handleExportTemplate exports a built-in template to stdout.
func handleExportTemplate(name string) error {
	if !templates.IsBuiltin(name) {
		return fmt.Errorf("unknown built-in template: %s (use --list-templates to see available templates)", name)
	}

	// Export image template
	content, err := templates.ExportBuiltin(name, templates.TemplateTypeImage)
	if err != nil {
		return fmt.Errorf("failed to export template: %w", err)
	}
	if _, err = fmt.Fprint(stdout, content); err != nil {
		return fmt.Errorf("failed to write template: %w", err)
	}
	return nil
}

// handleValidateTemplate validates a custom template file for syntax errors.
func handleValidateTemplate(path string) error {
	loader := templates.NewLoader(false)
	if err := loader.Validate(path); err != nil {
		slog.Error("template validation failed", "path", path, "error", err)
		return err
	}
	fmt.Fprintf(stdout, "Template %s is valid.\n", path)
	return nil
}
