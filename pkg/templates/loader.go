package templates

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"
)

//go:embed builtin/markdown/*.tmpl builtin/html/*.tmpl builtin/json/*.tmpl
var builtinFS embed.FS

// TemplateType describes the kind of template (image vs comparison).
type TemplateType string

const (
	// TemplateTypeImage is for single-image documentation.
	TemplateTypeImage TemplateType = "image"
	// TemplateTypeComparison is for multi-image comparison.
	TemplateTypeComparison TemplateType = "comparison"
)

// BuiltinInfo describes a built-in template.
type BuiltinInfo struct {
	Name        string
	Format      string
	Description string
}

// builtinTemplates is the registry of built-in template names and their metadata.
var builtinTemplates = []BuiltinInfo{
	{Name: "default", Format: "markdown", Description: "Standard format with badges and security analysis"},
	{Name: "minimal", Format: "markdown", Description: "ENV/ARG only, no badges or analysis"},
	{Name: "detailed", Format: "markdown", Description: "Expanded format with full package lists and metadata table"},
	{Name: "compact", Format: "markdown", Description: "Single-line condensed format for quick reference"},
	{Name: "html", Format: "html", Description: "Interactive HTML dashboard with styled tables"},
	{Name: "json", Format: "json", Description: "Machine-readable JSON output for CI/CD integration"},
}

// templatePaths maps template name + type to embedded file path.
var templatePaths = map[string]string{
	// Image templates
	"default:image":  "builtin/markdown/default.tmpl",
	"minimal:image":  "builtin/markdown/minimal.tmpl",
	"detailed:image": "builtin/markdown/detailed.tmpl",
	"compact:image":  "builtin/markdown/compact.tmpl",
	"html:image":     "builtin/html/dashboard.tmpl",
	"json:image":     "builtin/json/summary.tmpl",
	// Comparison templates
	"default:comparison":  "builtin/markdown/default_comparison.tmpl",
	"minimal:comparison":  "builtin/markdown/minimal_comparison.tmpl",
	"detailed:comparison": "builtin/markdown/detailed_comparison.tmpl",
	"compact:comparison":  "builtin/markdown/compact_comparison.tmpl",
	"html:comparison":     "builtin/html/dashboard_comparison.tmpl",
	"json:comparison":     "builtin/json/summary_comparison.tmpl",
}

// Loader handles template loading, caching, and validation.
type Loader struct {
	noMoji bool

	mu    sync.RWMutex
	cache map[string]*template.Template
}

// NewLoader creates a new template Loader.
func NewLoader(noMoji bool) *Loader {
	return &Loader{
		noMoji: noMoji,
		cache:  make(map[string]*template.Template),
	}
}

// ListBuiltin returns metadata for all built-in templates.
func ListBuiltin() []BuiltinInfo {
	result := make([]BuiltinInfo, len(builtinTemplates))
	copy(result, builtinTemplates)
	return result
}

// IsBuiltin returns true if the given name is a built-in template.
func IsBuiltin(name string) bool {
	for _, b := range builtinTemplates {
		if b.Name == name {
			return true
		}
	}
	return false
}

// FormatForTemplate returns the output format for a given template name.
func FormatForTemplate(name string) string {
	for _, b := range builtinTemplates {
		if b.Name == name {
			return b.Format
		}
	}
	return "markdown"
}

// OutputExtension returns the file extension (with dot) for a given template format.
// Returns ".md" for markdown, ".html" for html, ".json" for json.
func OutputExtension(format string) string {
	switch format {
	case "html":
		return ".html"
	case "json":
		return ".json"
	default:
		return ".md"
	}
}

// IsDirectWriteFormat returns true if the format produces a standalone document
// that should be written directly to a file instead of injected into an existing file.
func IsDirectWriteFormat(format string) bool {
	return format == "html" || format == "json"
}

// LoadBuiltin loads a built-in template by name and type.
func (l *Loader) LoadBuiltin(name string, tmplType TemplateType) (*template.Template, error) {
	cacheKey := name + ":" + string(tmplType)

	l.mu.RLock()
	if cached, ok := l.cache[cacheKey]; ok {
		l.mu.RUnlock()
		return cached, nil
	}
	l.mu.RUnlock()

	path, ok := templatePaths[cacheKey]
	if !ok {
		return nil, fmt.Errorf("unknown built-in template: %s (type: %s)", name, tmplType)
	}

	content, err := builtinFS.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read built-in template %s: %w", name, err)
	}

	tmpl, err := template.New(name).Funcs(GetFuncMap(l.noMoji)).Parse(string(content))
	if err != nil {
		return nil, fmt.Errorf("failed to parse built-in template %s: %w", name, err)
	}

	l.mu.Lock()
	l.cache[cacheKey] = tmpl
	l.mu.Unlock()

	return tmpl, nil
}

// LoadFile loads a custom template from a file path.
// It searches for the file in the following order:
// 1. Absolute path (if specified)
// 2. Relative to working directory
func (l *Loader) LoadFile(path string) (*template.Template, error) {
	cacheKey := "file:" + path

	l.mu.RLock()
	if cached, ok := l.cache[cacheKey]; ok {
		l.mu.RUnlock()
		return cached, nil
	}
	l.mu.RUnlock()

	// Security: prevent directory traversal.
	// Resolve to an absolute path and verify it stays within the working directory.
	cleanPath := filepath.Clean(path)
	absPath, err := filepath.Abs(cleanPath)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve template path: %w", err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to determine working directory: %w", err)
	}
	// Allow absolute paths that were explicitly given (they don't need
	// the cwd prefix check). Only relative paths must resolve within cwd.
	if !filepath.IsAbs(path) && !strings.HasPrefix(absPath, cwd+string(filepath.Separator)) && absPath != cwd {
		return nil, fmt.Errorf("template path escapes working directory: %s", path)
	}

	content, err := os.ReadFile(cleanPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read template file %s: %w", path, err)
	}

	// Security: limit file size (1MB)
	if len(content) > 1*1024*1024 {
		return nil, fmt.Errorf("template file too large: %s (max 1MB)", path)
	}

	name := filepath.Base(path)
	tmpl, err := template.New(name).Funcs(GetFuncMap(l.noMoji)).Parse(string(content))
	if err != nil {
		return nil, fmt.Errorf("failed to parse template file %s: %w", path, err)
	}

	l.mu.Lock()
	l.cache[cacheKey] = tmpl
	l.mu.Unlock()

	return tmpl, nil
}

// ExportBuiltin returns the raw content of a built-in template.
func ExportBuiltin(name string, tmplType TemplateType) (string, error) {
	cacheKey := name + ":" + string(tmplType)
	path, ok := templatePaths[cacheKey]
	if !ok {
		return "", fmt.Errorf("unknown built-in template: %s (type: %s)", name, tmplType)
	}

	content, err := builtinFS.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read built-in template %s: %w", name, err)
	}

	return string(content), nil
}

// Validate checks a template for syntax errors by parsing it.
// It returns nil if the template is valid.
func (l *Loader) Validate(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read template file %s: %w", path, err)
	}

	_, err = template.New("validate").Funcs(GetFuncMap(l.noMoji)).Parse(string(content))
	if err != nil {
		return fmt.Errorf("template syntax error in %s: %w", path, err)
	}

	return nil
}
