package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/northcutted/dock-docs/pkg/templates"
)

// resolveOutputPath determines the output file path for a given template format.
// If the user explicitly set a non-default output file, that is used as-is.
// Otherwise, the extension is derived from the template format (e.g. README.html, README.json).
func resolveOutputPath(currentOutput string, format string) string {
	ext := templates.OutputExtension(format)
	// If user explicitly set a custom output path (not the default), respect it
	if currentOutput != "README.md" {
		return currentOutput
	}
	// Replace the default .md extension with the format-appropriate extension
	return strings.TrimSuffix(currentOutput, filepath.Ext(currentOutput)) + ext
}

// resolveSectionOutput determines the output file path for a config section
// that uses a direct-write format (html/json).
// It derives the filename from the marker name or section index.
func resolveSectionOutput(baseOutput string, marker string, sectionIndex int, format string) string {
	ext := templates.OutputExtension(format)
	dir := filepath.Dir(baseOutput)
	base := strings.TrimSuffix(filepath.Base(baseOutput), filepath.Ext(baseOutput))

	// Use marker name as suffix if available, otherwise use section index
	var suffix string
	if marker != "" {
		suffix = "-" + marker
	} else {
		suffix = fmt.Sprintf("-section%d", sectionIndex)
	}

	return filepath.Join(dir, base+suffix+ext)
}
