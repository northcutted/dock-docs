package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// SectionType identifies the kind of section in a dock-docs configuration file.
type SectionType string

// Supported section types for dock-docs configuration.
const (
	SectionTypeImage      SectionType = "image"
	SectionTypeComparison SectionType = "comparison"
)

// TemplateConfig specifies the template to use for rendering.
type TemplateConfig struct {
	// Name is the built-in template name (e.g., "default", "minimal", "detailed", "compact", "html", "json").
	Name string `yaml:"name,omitempty"`
	// Path is the file path to a custom template (overrides Name).
	Path string `yaml:"path,omitempty"`
}

// ImageEntry describes a single image to analyze in a comparison section.
type ImageEntry struct {
	// Source is the path to a Dockerfile (optional; inherits from section if omitted).
	Source string `yaml:"source,omitempty"`
	// Tag is the Docker image tag to analyze.
	Tag string `yaml:"tag"`
}

// Section describes a single output section in a dock-docs configuration file.
type Section struct {
	Type   SectionType `yaml:"type"`
	Marker string      `yaml:"marker"`
	// Image section specific
	Source string `yaml:"source,omitempty"` // Dockerfile path
	Tag    string `yaml:"tag,omitempty"`    // Single image tag for image analysis
	// Comparison section specific
	Images  []ImageEntry `yaml:"images,omitempty"`
	Details bool         `yaml:"details,omitempty"` // Show full per-image analysis (collapsed) in comparison
	// Template overrides the global template for this section.
	Template *TemplateConfig `yaml:"template,omitempty"`
}

// ResolvedImages returns the list of image tags for a comparison section,
// resolving the shared Source field as the default for entries that omit it.
func (s Section) ResolvedImages() []ImageEntry {
	resolved := make([]ImageEntry, len(s.Images))
	for i, entry := range s.Images {
		resolved[i] = entry
		if resolved[i].Source == "" && s.Source != "" {
			resolved[i].Source = s.Source
		}
	}
	return resolved
}

// Config is the top-level structure for a dock-docs YAML configuration file.
type Config struct {
	Output       string    `yaml:"output"`
	BadgeBaseURL string    `yaml:"badgeBaseURL,omitempty"`
	Sections     []Section `yaml:"sections"`
	// Template is the global template configuration (can be overridden per-section).
	Template *TemplateConfig `yaml:"template,omitempty"`
}

// Load reads and parses a dock-docs YAML configuration file from the given path.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// Defaults
	if cfg.Output == "" {
		cfg.Output = "README.md"
	}
	if cfg.BadgeBaseURL == "" {
		cfg.BadgeBaseURL = "https://img.shields.io/static/v1"
	}

	return &cfg, nil
}

// ResolveTemplate returns the effective template config for a section,
// considering the section-level override and global default.
// Returns nil if no template is configured (use built-in default).
func (c *Config) ResolveTemplate(section Section) *TemplateConfig {
	if section.Template != nil {
		return section.Template
	}
	return c.Template
}
