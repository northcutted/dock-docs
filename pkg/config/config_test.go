package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoad_ValidConfig(t *testing.T) {
	validYAML := `output: "DOCS.md"
badgeBaseURL: "https://custom-badges.example.com/static/v1"
sections:
  - type: "image"
    marker: "main"
    source: "Dockerfile"
    tag: "myapp:latest"
  - type: "comparison"
    marker: "comparison"
    images:
      - tag: "alpine:3.18"
      - tag: "ubuntu:22.04"
`

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "dock-docs.yaml")
	err := os.WriteFile(configPath, []byte(validYAML), 0644)
	if err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.Output != "DOCS.md" {
		t.Errorf("Output = %v, want DOCS.md", cfg.Output)
	}

	if cfg.BadgeBaseURL != "https://custom-badges.example.com/static/v1" {
		t.Errorf("BadgeBaseURL = %v, want custom URL", cfg.BadgeBaseURL)
	}

	if len(cfg.Sections) != 2 {
		t.Fatalf("expected 2 sections, got %d", len(cfg.Sections))
	}

	// Test first section (image type)
	section1 := cfg.Sections[0]
	if section1.Type != SectionTypeImage {
		t.Errorf("Section[0].Type = %v, want %v", section1.Type, SectionTypeImage)
	}
	if section1.Marker != "main" {
		t.Errorf("Section[0].Marker = %v, want main", section1.Marker)
	}
	if section1.Source != "Dockerfile" {
		t.Errorf("Section[0].Source = %v, want Dockerfile", section1.Source)
	}
	if section1.Tag != "myapp:latest" {
		t.Errorf("Section[0].Tag = %v, want myapp:latest", section1.Tag)
	}

	// Test second section (comparison type)
	section2 := cfg.Sections[1]
	if section2.Type != SectionTypeComparison {
		t.Errorf("Section[1].Type = %v, want %v", section2.Type, SectionTypeComparison)
	}
	if section2.Marker != "comparison" {
		t.Errorf("Section[1].Marker = %v, want comparison", section2.Marker)
	}
	if len(section2.Images) != 2 {
		t.Fatalf("expected 2 images in comparison, got %d", len(section2.Images))
	}
	if section2.Images[0].Tag != "alpine:3.18" {
		t.Errorf("Section[1].Images[0].Tag = %v, want alpine:3.18", section2.Images[0].Tag)
	}
	if section2.Images[1].Tag != "ubuntu:22.04" {
		t.Errorf("Section[1].Images[1].Tag = %v, want ubuntu:22.04", section2.Images[1].Tag)
	}
}

func TestLoad_Defaults(t *testing.T) {
	minimalYAML := `sections:
  - type: "image"
    marker: "main"
`

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "dock-docs.yaml")
	err := os.WriteFile(configPath, []byte(minimalYAML), 0644)
	if err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	// Test default values
	if cfg.Output != "README.md" {
		t.Errorf("Output default = %v, want README.md", cfg.Output)
	}

	if cfg.BadgeBaseURL != "https://img.shields.io/static/v1" {
		t.Errorf("BadgeBaseURL default = %v, want shields.io URL", cfg.BadgeBaseURL)
	}
}

func TestLoad_FileNotFound(t *testing.T) {
	_, err := Load("/nonexistent/path/config.yaml")
	if err == nil {
		t.Error("expected error for nonexistent file, got nil")
	}
}

func TestLoad_InvalidYAML(t *testing.T) {
	invalidYAML := `output: "test.md"
sections:
  - type: image
    marker: "main"
    invalid: [unclosed bracket
`

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "invalid.yaml")
	err := os.WriteFile(configPath, []byte(invalidYAML), 0644)
	if err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	_, err = Load(configPath)
	if err == nil {
		t.Error("expected error for invalid YAML, got nil")
	}
}

func TestLoad_EmptyFile(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "empty.yaml")
	err := os.WriteFile(configPath, []byte(""), 0644)
	if err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	_, err = Load(configPath)
	if err == nil {
		t.Fatal("expected error for empty config with no sections")
	}
}

func TestLoad_OnlyOutput(t *testing.T) {
	yamlWithOnlyOutput := `output: "CUSTOM.md"
`

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "output-only.yaml")
	err := os.WriteFile(configPath, []byte(yamlWithOnlyOutput), 0644)
	if err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	_, err = Load(configPath)
	if err == nil {
		t.Fatal("expected error for config with no sections")
	}
}

func TestLoad_ComplexComparisonConfig(t *testing.T) {
	complexYAML := `output: "README.md"
sections:
  - type: "comparison"
    marker: "versions"
    images:
      - tag: "python:3.10"
      - tag: "python:3.11"
      - tag: "python:3.12"
      - tag: "python:3.10-alpine"
      - tag: "python:3.11-alpine"
      - tag: "python:3.12-alpine"
`

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "comparison.yaml")
	err := os.WriteFile(configPath, []byte(complexYAML), 0644)
	if err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if len(cfg.Sections) != 1 {
		t.Fatalf("expected 1 section, got %d", len(cfg.Sections))
	}

	section := cfg.Sections[0]
	if section.Type != SectionTypeComparison {
		t.Errorf("Type = %v, want comparison", section.Type)
	}

	if len(section.Images) != 6 {
		t.Errorf("expected 6 images, got %d", len(section.Images))
	}
}

func TestLoad_MultipleImageSections(t *testing.T) {
	multipleImagesYAML := `output: "README.md"
sections:
  - type: "image"
    marker: "production"
    source: "Dockerfile.prod"
    tag: "myapp:prod"
  - type: "image"
    marker: "development"
    source: "Dockerfile.dev"
    tag: "myapp:dev"
  - type: "image"
    marker: "testing"
    source: "Dockerfile.test"
`

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "multi-image.yaml")
	err := os.WriteFile(configPath, []byte(multipleImagesYAML), 0644)
	if err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if len(cfg.Sections) != 3 {
		t.Fatalf("expected 3 sections, got %d", len(cfg.Sections))
	}

	// Verify each section
	expectedMarkers := []string{"production", "development", "testing"}
	expectedSources := []string{"Dockerfile.prod", "Dockerfile.dev", "Dockerfile.test"}
	expectedTags := []string{"myapp:prod", "myapp:dev", ""}

	for i, section := range cfg.Sections {
		if section.Type != SectionTypeImage {
			t.Errorf("Section[%d].Type = %v, want image", i, section.Type)
		}
		if section.Marker != expectedMarkers[i] {
			t.Errorf("Section[%d].Marker = %v, want %v", i, section.Marker, expectedMarkers[i])
		}
		if section.Source != expectedSources[i] {
			t.Errorf("Section[%d].Source = %v, want %v", i, section.Source, expectedSources[i])
		}
		if section.Tag != expectedTags[i] {
			t.Errorf("Section[%d].Tag = %v, want %v", i, section.Tag, expectedTags[i])
		}
	}
}

func TestSectionType_Constants(t *testing.T) {
	if SectionTypeImage != "image" {
		t.Errorf("SectionTypeImage = %v, want 'image'", SectionTypeImage)
	}

	if SectionTypeComparison != "comparison" {
		t.Errorf("SectionTypeComparison = %v, want 'comparison'", SectionTypeComparison)
	}
}

func TestLoad_ComparisonWithSharedSource(t *testing.T) {
	yaml := `output: "README.md"
sections:
  - type: "comparison"
    marker: "compare"
    source: "Dockerfile"
    details: true
    images:
      - tag: "myapp:dev"
      - source: "Dockerfile.prod"
        tag: "myapp:prod"
      - tag: "myapp:staging"
`

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "shared-source.yaml")
	err := os.WriteFile(configPath, []byte(yaml), 0644)
	if err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	section := cfg.Sections[0]
	if section.Source != "Dockerfile" {
		t.Errorf("Section.Source = %v, want Dockerfile", section.Source)
	}
	if !section.Details {
		t.Error("Section.Details = false, want true")
	}

	// Test ResolvedImages
	resolved := section.ResolvedImages()
	if len(resolved) != 3 {
		t.Fatalf("expected 3 resolved images, got %d", len(resolved))
	}

	// First image: inherits shared source
	if resolved[0].Source != "Dockerfile" {
		t.Errorf("resolved[0].Source = %v, want Dockerfile (inherited)", resolved[0].Source)
	}
	if resolved[0].Tag != "myapp:dev" {
		t.Errorf("resolved[0].Tag = %v, want myapp:dev", resolved[0].Tag)
	}

	// Second image: overrides shared source
	if resolved[1].Source != "Dockerfile.prod" {
		t.Errorf("resolved[1].Source = %v, want Dockerfile.prod (override)", resolved[1].Source)
	}
	if resolved[1].Tag != "myapp:prod" {
		t.Errorf("resolved[1].Tag = %v, want myapp:prod", resolved[1].Tag)
	}

	// Third image: inherits shared source
	if resolved[2].Source != "Dockerfile" {
		t.Errorf("resolved[2].Source = %v, want Dockerfile (inherited)", resolved[2].Source)
	}
	if resolved[2].Tag != "myapp:staging" {
		t.Errorf("resolved[2].Tag = %v, want myapp:staging", resolved[2].Tag)
	}
}

func TestResolveTemplate(t *testing.T) {
	globalTmpl := &TemplateConfig{Name: "minimal"}
	sectionTmpl := &TemplateConfig{Name: "detailed"}

	tests := []struct {
		name     string
		cfg      *Config
		section  Section
		expected *TemplateConfig
	}{
		{
			name:     "section override wins",
			cfg:      &Config{Template: globalTmpl},
			section:  Section{Template: sectionTmpl},
			expected: sectionTmpl,
		},
		{
			name:     "falls back to global",
			cfg:      &Config{Template: globalTmpl},
			section:  Section{},
			expected: globalTmpl,
		},
		{
			name:     "both nil returns nil",
			cfg:      &Config{},
			section:  Section{},
			expected: nil,
		},
		{
			name:     "section set global nil",
			cfg:      &Config{},
			section:  Section{Template: sectionTmpl},
			expected: sectionTmpl,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.cfg.ResolveTemplate(tt.section)
			if result != tt.expected {
				t.Errorf("ResolveTemplate() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestLoad_ComparisonDetailsDefault(t *testing.T) {
	yaml := `output: "README.md"
sections:
  - type: "comparison"
    marker: "compare"
    images:
      - tag: "myapp:v1"
`

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "details-default.yaml")
	err := os.WriteFile(configPath, []byte(yaml), 0644)
	if err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	section := cfg.Sections[0]
	if section.Details {
		t.Error("Section.Details should default to false")
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     Config
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid image section",
			cfg: Config{
				Sections: []Section{
					{Type: SectionTypeImage, Marker: "main"},
				},
			},
			wantErr: false,
		},
		{
			name: "valid comparison section",
			cfg: Config{
				Sections: []Section{
					{Type: SectionTypeComparison, Marker: "compare", Images: []ImageEntry{{Tag: "a:1"}}},
				},
			},
			wantErr: false,
		},
		{
			name: "valid mixed sections",
			cfg: Config{
				Sections: []Section{
					{Type: SectionTypeImage, Marker: "img"},
					{Type: SectionTypeComparison, Marker: "cmp", Images: []ImageEntry{{Tag: "a:1"}, {Tag: "b:2"}}},
				},
			},
			wantErr: false,
		},
		{
			name:    "no sections",
			cfg:     Config{},
			wantErr: true,
			errMsg:  "at least one section",
		},
		{
			name: "invalid section type",
			cfg: Config{
				Sections: []Section{
					{Type: "bogus", Marker: "main"},
				},
			},
			wantErr: true,
			errMsg:  "invalid type",
		},
		{
			name: "empty section type",
			cfg: Config{
				Sections: []Section{
					{Type: "", Marker: "main"},
				},
			},
			wantErr: true,
			errMsg:  "invalid type",
		},
		{
			name: "comparison with no images",
			cfg: Config{
				Sections: []Section{
					{Type: SectionTypeComparison, Marker: "compare"},
				},
			},
			wantErr: true,
			errMsg:  "at least one image",
		},
		{
			name: "first section valid second invalid",
			cfg: Config{
				Sections: []Section{
					{Type: SectionTypeImage, Marker: "ok"},
					{Type: "invalid", Marker: "bad"},
				},
			},
			wantErr: true,
			errMsg:  "section 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.Validate()
			if tt.wantErr {
				if err == nil {
					t.Fatalf("Validate() expected error containing %q, got nil", tt.errMsg)
				}
				if tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("Validate() error = %q, want substring %q", err.Error(), tt.errMsg)
				}
			} else {
				if err != nil {
					t.Fatalf("Validate() unexpected error: %v", err)
				}
			}
		})
	}
}

func TestLoad_InvalidSectionType(t *testing.T) {
	yamlContent := `output: "README.md"
sections:
  - type: "foobar"
    marker: "test"
`

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "invalid-type.yaml")
	if err := os.WriteFile(configPath, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	_, err := Load(configPath)
	if err == nil {
		t.Fatal("expected error for invalid section type")
	}
}

func TestLoad_ComparisonNoImages(t *testing.T) {
	yamlContent := `output: "README.md"
sections:
  - type: "comparison"
    marker: "compare"
`

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "no-images.yaml")
	if err := os.WriteFile(configPath, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	_, err := Load(configPath)
	if err == nil {
		t.Fatal("expected error for comparison section with no images")
	}
}

func TestResolveRelativePaths(t *testing.T) {
	tests := []struct {
		name           string
		baseDir        string
		cfg            Config
		expectedOutput string
		expectedSource string
		expectedTmpl   string
	}{
		{
			name:    "resolves relative output and source",
			baseDir: "/projects/myapp",
			cfg: Config{
				Output: "README.md",
				Sections: []Section{
					{
						Type:   SectionTypeImage,
						Source: "Dockerfile",
					},
				},
			},
			expectedOutput: "/projects/myapp/README.md",
			expectedSource: "/projects/myapp/Dockerfile",
		},
		{
			name:    "preserves absolute paths",
			baseDir: "/projects/myapp",
			cfg: Config{
				Output: "/absolute/README.md",
				Sections: []Section{
					{
						Type:   SectionTypeImage,
						Source: "/absolute/Dockerfile",
					},
				},
			},
			expectedOutput: "/absolute/README.md",
			expectedSource: "/absolute/Dockerfile",
		},
		{
			name:    "empty baseDir is no-op",
			baseDir: "",
			cfg: Config{
				Output: "README.md",
				Sections: []Section{
					{
						Type:   SectionTypeImage,
						Source: "Dockerfile",
					},
				},
			},
			expectedOutput: "README.md",
			expectedSource: "Dockerfile",
		},
		{
			name:    "dot baseDir is no-op",
			baseDir: ".",
			cfg: Config{
				Output: "README.md",
				Sections: []Section{
					{
						Type:   SectionTypeImage,
						Source: "Dockerfile",
					},
				},
			},
			expectedOutput: "README.md",
			expectedSource: "Dockerfile",
		},
		{
			name:    "resolves template path",
			baseDir: "/projects/myapp",
			cfg: Config{
				Output: "README.md",
				Sections: []Section{
					{
						Type:     SectionTypeImage,
						Source:   "Dockerfile",
						Template: &TemplateConfig{Path: "templates/custom.tmpl"},
					},
				},
			},
			expectedOutput: "/projects/myapp/README.md",
			expectedSource: "/projects/myapp/Dockerfile",
			expectedTmpl:   "/projects/myapp/templates/custom.tmpl",
		},
		{
			name:    "empty source stays empty",
			baseDir: "/projects/myapp",
			cfg: Config{
				Output: "out.md",
				Sections: []Section{
					{
						Type: SectionTypeImage,
					},
				},
			},
			expectedOutput: "/projects/myapp/out.md",
			expectedSource: "",
		},
		{
			name:    "nested relative paths",
			baseDir: "/projects/myapp",
			cfg: Config{
				Output: "docs/output/README.md",
				Sections: []Section{
					{
						Type:   SectionTypeImage,
						Source: "docker/Dockerfile.prod",
					},
				},
			},
			expectedOutput: "/projects/myapp/docs/output/README.md",
			expectedSource: "/projects/myapp/docker/Dockerfile.prod",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := tt.cfg
			cfg.ResolveRelativePaths(tt.baseDir)

			if cfg.Output != tt.expectedOutput {
				t.Errorf("Output = %q, want %q", cfg.Output, tt.expectedOutput)
			}
			if len(cfg.Sections) > 0 && cfg.Sections[0].Source != tt.expectedSource {
				t.Errorf("Source = %q, want %q", cfg.Sections[0].Source, tt.expectedSource)
			}
			if tt.expectedTmpl != "" && cfg.Sections[0].Template != nil {
				if cfg.Sections[0].Template.Path != tt.expectedTmpl {
					t.Errorf("Template.Path = %q, want %q", cfg.Sections[0].Template.Path, tt.expectedTmpl)
				}
			}
		})
	}
}
