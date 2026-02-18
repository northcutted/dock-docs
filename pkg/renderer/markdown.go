package renderer

import (
	"fmt"
	"text/template"

	"github.com/northcutted/dock-docs/pkg/parser"
	"github.com/northcutted/dock-docs/pkg/templates"
	"github.com/northcutted/dock-docs/pkg/types"
)

// RenderOptions configures rendering behavior.
type RenderOptions struct {
	NoMoji       bool
	BadgeBaseURL string
}

// TemplateSelection specifies which template to use.
// If both Name and Path are empty, the "default" built-in is used.
type TemplateSelection struct {
	// Name is a built-in template name (e.g., "default", "minimal", "detailed", "compact", "html", "json").
	Name string
	// Path is the file path to a custom template (overrides Name).
	Path string
}

// Format returns the output format for this template selection.
// Returns "markdown", "html", or "json". Custom file templates default to "markdown".
func (s TemplateSelection) Format() string {
	if s.Path != "" {
		return "markdown" // custom templates default to markdown
	}
	name := s.Name
	if name == "" {
		name = "default"
	}
	return templates.FormatForTemplate(name)
}

// ReportContext holds all data passed to the image template.
type ReportContext struct {
	Doc      *parser.Documentation
	Stats    *types.ImageStats
	ImageTag string
	Options  RenderOptions
}

// Emoji returns the emoji or text alternative for the given name.
func (r ReportContext) Emoji(name string) string {
	return templates.GetEmoji(name, r.Options.NoMoji)
}

// ComparisonContext holds all data passed to the comparison template.
type ComparisonContext struct {
	Images  []*types.ImageStats
	Options RenderOptions
}

// Emoji returns the emoji or text alternative for the given name.
func (r ComparisonContext) Emoji(name string) string {
	return templates.GetEmoji(name, r.Options.NoMoji)
}

// Render generates documentation using the default built-in template.
// This is the backward-compatible entry point.
func Render(doc *parser.Documentation, stats *types.ImageStats, opts RenderOptions) (string, error) {
	return RenderWithTemplate(doc, stats, opts, TemplateSelection{Name: "default"})
}

// RenderWithTemplate generates documentation using the specified template.
func RenderWithTemplate(doc *parser.Documentation, stats *types.ImageStats, opts RenderOptions, sel TemplateSelection) (string, error) {
	tmpl, err := resolveTemplate(sel, templates.TemplateTypeImage, opts.NoMoji)
	if err != nil {
		return "", fmt.Errorf("failed to load template: %w", err)
	}

	ctx := ReportContext{
		Doc:     doc,
		Stats:   stats,
		Options: opts,
	}
	if stats != nil {
		ctx.ImageTag = stats.ImageTag
	} else {
		if ctx.ImageTag == "" {
			ctx.ImageTag = "Dockerfile"
		}
	}

	sec := templates.DefaultSecurityConfig()
	return templates.ExecuteWithLimits(tmpl, ctx, sec)
}

// RenderComparison generates the comparison table using the default built-in template.
// This is the backward-compatible entry point.
func RenderComparison(stats []*types.ImageStats, opts RenderOptions) (string, error) {
	return RenderComparisonWithTemplate(stats, opts, TemplateSelection{Name: "default"})
}

// RenderComparisonWithTemplate generates the comparison table using the specified template.
func RenderComparisonWithTemplate(stats []*types.ImageStats, opts RenderOptions, sel TemplateSelection) (string, error) {
	tmpl, err := resolveTemplate(sel, templates.TemplateTypeComparison, opts.NoMoji)
	if err != nil {
		return "", fmt.Errorf("failed to load template: %w", err)
	}

	ctx := ComparisonContext{
		Images:  stats,
		Options: opts,
	}

	sec := templates.DefaultSecurityConfig()
	return templates.ExecuteWithLimits(tmpl, ctx, sec)
}

// resolveTemplate loads the appropriate template based on the selection.
func resolveTemplate(sel TemplateSelection, tmplType templates.TemplateType, noMoji bool) (*template.Template, error) {
	loader := templates.NewLoader(noMoji)

	// Custom file path takes precedence
	if sel.Path != "" {
		return loader.LoadFile(sel.Path)
	}

	// Built-in template by name
	name := sel.Name
	if name == "" {
		name = "default"
	}

	// Check if it's a built-in name
	if !templates.IsBuiltin(name) {
		return nil, fmt.Errorf("unknown template %q: not a built-in template name or file path", name)
	}

	return loader.LoadBuiltin(name, tmplType)
}
