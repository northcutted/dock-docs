package templates

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"
	"time"
)

// --- Loader Tests ---

func TestListBuiltin(t *testing.T) {
	builtins := ListBuiltin()
	if len(builtins) != 6 {
		t.Fatalf("expected 6 built-in templates, got %d", len(builtins))
	}

	expectedNames := []string{"default", "minimal", "detailed", "compact", "html", "json"}
	for i, b := range builtins {
		if b.Name != expectedNames[i] {
			t.Errorf("builtins[%d].Name = %q, want %q", i, b.Name, expectedNames[i])
		}
		if b.Format == "" {
			t.Errorf("builtins[%d].Format is empty", i)
		}
		if b.Description == "" {
			t.Errorf("builtins[%d].Description is empty", i)
		}
	}
}

func TestIsBuiltin(t *testing.T) {
	tests := []struct {
		name     string
		expected bool
	}{
		{"default", true},
		{"minimal", true},
		{"detailed", true},
		{"compact", true},
		{"html", true},
		{"json", true},
		{"nonexistent", false},
		{"", false},
		{"Default", false}, // case-sensitive
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsBuiltin(tt.name)
			if result != tt.expected {
				t.Errorf("IsBuiltin(%q) = %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}

func TestFormatForTemplate(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"default", "markdown"},
		{"minimal", "markdown"},
		{"detailed", "markdown"},
		{"compact", "markdown"},
		{"html", "html"},
		{"json", "json"},
		{"nonexistent", "markdown"}, // default fallback
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatForTemplate(tt.name)
			if result != tt.expected {
				t.Errorf("FormatForTemplate(%q) = %q, want %q", tt.name, result, tt.expected)
			}
		})
	}
}

func TestLoadBuiltin_AllImageTemplates(t *testing.T) {
	loader := NewLoader(false)
	names := []string{"default", "minimal", "detailed", "compact", "html", "json"}

	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			tmpl, err := loader.LoadBuiltin(name, TemplateTypeImage)
			if err != nil {
				t.Fatalf("LoadBuiltin(%q, image) error = %v", name, err)
			}
			if tmpl == nil {
				t.Fatalf("LoadBuiltin(%q, image) returned nil template", name)
			}
		})
	}
}

func TestLoadBuiltin_AllComparisonTemplates(t *testing.T) {
	loader := NewLoader(false)
	names := []string{"default", "minimal", "detailed", "compact", "html", "json"}

	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			tmpl, err := loader.LoadBuiltin(name, TemplateTypeComparison)
			if err != nil {
				t.Fatalf("LoadBuiltin(%q, comparison) error = %v", name, err)
			}
			if tmpl == nil {
				t.Fatalf("LoadBuiltin(%q, comparison) returned nil template", name)
			}
		})
	}
}

func TestLoadBuiltin_UnknownTemplate(t *testing.T) {
	loader := NewLoader(false)
	_, err := loader.LoadBuiltin("nonexistent", TemplateTypeImage)
	if err == nil {
		t.Fatal("expected error for unknown template, got nil")
	}
	if !strings.Contains(err.Error(), "unknown built-in template") {
		t.Errorf("error message = %q, want it to contain 'unknown built-in template'", err.Error())
	}
}

func TestLoadBuiltin_UnknownComparisonTemplate(t *testing.T) {
	loader := NewLoader(false)
	_, err := loader.LoadBuiltin("nonexistent", TemplateTypeComparison)
	if err == nil {
		t.Fatal("expected error for unknown comparison template, got nil")
	}
	if !strings.Contains(err.Error(), "unknown built-in template") {
		t.Errorf("error message = %q, want it to contain 'unknown built-in template'", err.Error())
	}
}

func TestLoadBuiltin_Caching(t *testing.T) {
	loader := NewLoader(false)

	// Load same template twice, should get cached version
	tmpl1, err := loader.LoadBuiltin("default", TemplateTypeImage)
	if err != nil {
		t.Fatalf("first LoadBuiltin error = %v", err)
	}

	tmpl2, err := loader.LoadBuiltin("default", TemplateTypeImage)
	if err != nil {
		t.Fatalf("second LoadBuiltin error = %v", err)
	}

	// Both should be the same pointer (cached)
	if tmpl1 != tmpl2 {
		t.Error("expected cached template to return same pointer")
	}
}

func TestLoadFile(t *testing.T) {
	// Create a temporary template file
	tmpDir := t.TempDir()
	tmplPath := filepath.Join(tmpDir, "custom.tmpl")
	content := `# Custom Template: {{ .ImageTag }}`
	if err := os.WriteFile(tmplPath, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write template file: %v", err)
	}

	loader := NewLoader(false)
	tmpl, err := loader.LoadFile(tmplPath)
	if err != nil {
		t.Fatalf("LoadFile() error = %v", err)
	}
	if tmpl == nil {
		t.Fatal("LoadFile() returned nil template")
	}
}

func TestLoadFile_DirectoryTraversal(t *testing.T) {
	loader := NewLoader(false)
	_, err := loader.LoadFile("../../../etc/passwd")
	if err == nil {
		t.Fatal("expected error for directory traversal, got nil")
	}
	if !strings.Contains(err.Error(), "escapes working directory") {
		t.Errorf("error = %q, want it to contain 'escapes working directory'", err.Error())
	}
}

func TestLoadFile_FileNotFound(t *testing.T) {
	loader := NewLoader(false)
	_, err := loader.LoadFile("/nonexistent/template.tmpl")
	if err == nil {
		t.Fatal("expected error for nonexistent file, got nil")
	}
}

func TestLoadFile_InvalidSyntax(t *testing.T) {
	tmpDir := t.TempDir()
	tmplPath := filepath.Join(tmpDir, "bad.tmpl")
	content := `{{ .Unclosed`
	if err := os.WriteFile(tmplPath, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write template file: %v", err)
	}

	loader := NewLoader(false)
	_, err := loader.LoadFile(tmplPath)
	if err == nil {
		t.Fatal("expected error for invalid syntax, got nil")
	}
	if !strings.Contains(err.Error(), "failed to parse template") {
		t.Errorf("error = %q, want it to contain 'failed to parse template'", err.Error())
	}
}

func TestLoadFile_TooLarge(t *testing.T) {
	tmpDir := t.TempDir()
	tmplPath := filepath.Join(tmpDir, "large.tmpl")
	// Create a file > 1MB
	content := strings.Repeat("x", 1*1024*1024+1)
	if err := os.WriteFile(tmplPath, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write template file: %v", err)
	}

	loader := NewLoader(false)
	_, err := loader.LoadFile(tmplPath)
	if err == nil {
		t.Fatal("expected error for oversized file, got nil")
	}
	if !strings.Contains(err.Error(), "too large") {
		t.Errorf("error = %q, want it to contain 'too large'", err.Error())
	}
}

func TestExportBuiltin(t *testing.T) {
	content, err := ExportBuiltin("default", TemplateTypeImage)
	if err != nil {
		t.Fatalf("ExportBuiltin() error = %v", err)
	}
	if content == "" {
		t.Fatal("ExportBuiltin() returned empty content")
	}
	// The default template should contain key markers
	if !strings.Contains(content, "Docker Image Analysis") {
		t.Error("exported template should contain 'Docker Image Analysis'")
	}
}

func TestExportBuiltin_Unknown(t *testing.T) {
	_, err := ExportBuiltin("nonexistent", TemplateTypeImage)
	if err == nil {
		t.Fatal("expected error for unknown template, got nil")
	}
}

func TestValidate_ValidTemplate(t *testing.T) {
	tmpDir := t.TempDir()
	tmplPath := filepath.Join(tmpDir, "valid.tmpl")
	content := `# Valid {{ .ImageTag }}`
	if err := os.WriteFile(tmplPath, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write template file: %v", err)
	}

	loader := NewLoader(false)
	err := loader.Validate(tmplPath)
	if err != nil {
		t.Fatalf("Validate() error = %v for valid template", err)
	}
}

func TestValidate_InvalidTemplate(t *testing.T) {
	tmpDir := t.TempDir()
	tmplPath := filepath.Join(tmpDir, "invalid.tmpl")
	content := `{{ .Unclosed`
	if err := os.WriteFile(tmplPath, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write template file: %v", err)
	}

	loader := NewLoader(false)
	err := loader.Validate(tmplPath)
	if err == nil {
		t.Fatal("expected error for invalid template, got nil")
	}
	if !strings.Contains(err.Error(), "syntax error") {
		t.Errorf("error = %q, want it to contain 'syntax error'", err.Error())
	}
}

func TestValidate_FileNotFound(t *testing.T) {
	loader := NewLoader(false)
	err := loader.Validate("/nonexistent/template.tmpl")
	if err == nil {
		t.Fatal("expected error for nonexistent file, got nil")
	}
}

// --- FuncMap Tests ---

func TestGetFuncMap(t *testing.T) {
	fm := GetFuncMap(false)

	// Verify all expected functions are present
	expectedFuncs := []string{
		"index", "join", "lower", "upper", "trim", "contains",
		"hasPrefix", "hasSuffix", "replace", "default", "jsonEscape", "emoji",
	}
	for _, name := range expectedFuncs {
		if _, ok := fm[name]; !ok {
			t.Errorf("missing function %q in FuncMap", name)
		}
	}
}

func TestSafeMapIndex(t *testing.T) {
	tests := []struct {
		name     string
		m        map[string]int
		key      string
		expected int
	}{
		{"existing key", map[string]int{"a": 1, "b": 2}, "a", 1},
		{"missing key", map[string]int{"a": 1}, "b", 0},
		{"empty map", map[string]int{}, "a", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := safeMapIndex(tt.m, tt.key)
			if result != tt.expected {
				t.Errorf("safeMapIndex(%v, %q) = %d, want %d", tt.m, tt.key, result, tt.expected)
			}
		})
	}
}

func TestDefaultValue(t *testing.T) {
	tests := []struct {
		name     string
		fallback string
		val      string
		expected string
	}{
		{"empty uses fallback", "fallback", "", "fallback"},
		{"non-empty uses val", "fallback", "actual", "actual"},
		{"both empty", "", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := defaultValue(tt.fallback, tt.val)
			if result != tt.expected {
				t.Errorf("defaultValue(%q, %q) = %q, want %q", tt.fallback, tt.val, result, tt.expected)
			}
		})
	}
}

func TestJsonEscape(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"plain text", "hello", "hello"},
		{"quotes", `say "hello"`, `say \"hello\"`},
		{"backslash", `path\to\file`, `path\\to\\file`},
		{"newline", "line1\nline2", `line1\nline2`},
		{"tab", "col1\tcol2", `col1\tcol2`},
		{"carriage return", "text\rmore", `text\rmore`},
		{"combined", "a\"\n\\", `a\"\n\\`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := jsonEscape(tt.input)
			if result != tt.expected {
				t.Errorf("jsonEscape(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGetEmoji(t *testing.T) {
	tests := []struct {
		name     string
		emoji    string
		noMoji   bool
		expected string
	}{
		{"whale emoji", "whale", false, "\U0001F433 "},
		{"whale text", "whale", true, ""},
		{"check emoji", "check", false, "\u2705"},
		{"check text", "check", true, "[YES]"},
		{"cross emoji", "cross", false, "\u274C"},
		{"cross text", "cross", true, "[NO]"},
		{"gear emoji", "gear", false, "\u2699\uFE0F "},
		{"shield emoji", "shield", false, "\U0001F6E1\uFE0F "},
		{"tag emoji", "tag", false, "\U0001F3F7\uFE0F "},
		{"search emoji", "search", false, "\U0001F50D "},
		{"down emoji", "down", false, "\U0001F447 "},
		{"package emoji", "package", false, "\U0001F4E6 "},
		{"critical emoji", "critical", false, "\U0001F6D1"},
		{"critical text", "critical", true, "[CRIT]"},
		{"high emoji", "high", false, "\U0001F7E0"},
		{"high text", "high", true, "[HIGH]"},
		{"medium emoji", "medium", false, "\U0001F7E1"},
		{"medium text", "medium", true, "[MED]"},
		{"low emoji", "low", false, "\U0001F535"},
		{"low text", "low", true, "[LOW]"},
		{"clean emoji", "clean", false, "\U0001F7E2"},
		{"clean text", "clean", true, "[OK]"},
		{"unknown emoji", "unknown", false, ""},
		{"unknown text", "unknown", true, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetEmoji(tt.emoji, tt.noMoji)
			if result != tt.expected {
				t.Errorf("GetEmoji(%q, %v) = %q, want %q", tt.emoji, tt.noMoji, result, tt.expected)
			}
		})
	}
}

func TestGetFuncMap_EmojiFunction(t *testing.T) {
	// Test emoji function with noMoji=false
	fm := GetFuncMap(false)
	emojiFn := fm["emoji"].(func(string) string)
	result := emojiFn("check")
	if result != "\u2705" {
		t.Errorf("emoji('check') = %q, want check mark emoji", result)
	}

	// Test emoji function with noMoji=true
	fm2 := GetFuncMap(true)
	emojiFn2 := fm2["emoji"].(func(string) string)
	result2 := emojiFn2("check")
	if result2 != "[YES]" {
		t.Errorf("emoji('check') with noMoji = %q, want '[YES]'", result2)
	}
}

// --- Security Tests ---

func TestDefaultSecurityConfig(t *testing.T) {
	sec := DefaultSecurityConfig()
	if sec.MaxExecutionTime != 30*time.Second {
		t.Errorf("MaxExecutionTime = %v, want 30s", sec.MaxExecutionTime)
	}
	if sec.MaxOutputSize != 10*1024*1024 {
		t.Errorf("MaxOutputSize = %d, want 10MB", sec.MaxOutputSize)
	}
}

func TestExecuteWithLimits_Success(t *testing.T) {
	tmpl, err := template.New("test").Parse("Hello {{ .Name }}")
	if err != nil {
		t.Fatalf("failed to parse template: %v", err)
	}

	data := struct{ Name string }{"World"}
	sec := DefaultSecurityConfig()

	result, err := ExecuteWithLimits(tmpl, data, sec)
	if err != nil {
		t.Fatalf("ExecuteWithLimits() error = %v", err)
	}
	if result != "Hello World" {
		t.Errorf("result = %q, want 'Hello World'", result)
	}
}

func TestExecuteWithLimits_OutputSizeExceeded(t *testing.T) {
	// Template that generates output larger than limit
	tmpl, err := template.New("test").Parse("{{ range .Items }}x{{ end }}")
	if err != nil {
		t.Fatalf("failed to parse template: %v", err)
	}

	// Create data that will generate >100 bytes of output
	items := make([]int, 200)
	data := struct{ Items []int }{items}
	sec := SecurityConfig{
		MaxExecutionTime: 30 * time.Second,
		MaxOutputSize:    100, // Very small limit
	}

	_, err = ExecuteWithLimits(tmpl, data, sec)
	if err == nil {
		t.Fatal("expected error for output size exceeded, got nil")
	}
	if !strings.Contains(err.Error(), "exceeded limit") {
		t.Errorf("error = %q, want it to contain 'exceeded limit'", err.Error())
	}
}

func TestExecuteWithLimits_TemplateError(t *testing.T) {
	// Template that references a non-existent field
	tmpl, err := template.New("test").Parse("{{ .NonExistent.Method }}")
	if err != nil {
		t.Fatalf("failed to parse template: %v", err)
	}

	data := struct{ Name string }{"World"}
	sec := DefaultSecurityConfig()

	_, err = ExecuteWithLimits(tmpl, data, sec)
	if err == nil {
		t.Fatal("expected error for template execution failure, got nil")
	}
	if !strings.Contains(err.Error(), "template execution failed") {
		t.Errorf("error = %q, want it to contain 'template execution failed'", err.Error())
	}
}

func TestLimitedWriter(t *testing.T) {
	w := &limitedWriter{limit: 10}

	// Write within limit
	n, err := w.Write([]byte("hello"))
	if err != nil {
		t.Fatalf("Write() error = %v", err)
	}
	if n != 5 {
		t.Errorf("Write() = %d, want 5", n)
	}

	// Write that exceeds limit
	_, err = w.Write([]byte("world!"))
	if err == nil {
		t.Fatal("expected error when exceeding limit, got nil")
	}

	// Verify content from successful write
	if w.String() != "hello" {
		t.Errorf("String() = %q, want 'hello'", w.String())
	}
}

func TestOutputExtension(t *testing.T) {
	tests := []struct {
		format   string
		expected string
	}{
		{"markdown", ".md"},
		{"html", ".html"},
		{"json", ".json"},
		{"", ".md"},
		{"unknown", ".md"},
	}
	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			result := OutputExtension(tt.format)
			if result != tt.expected {
				t.Errorf("OutputExtension(%q) = %q, want %q", tt.format, result, tt.expected)
			}
		})
	}
}

func TestIsDirectWriteFormat(t *testing.T) {
	tests := []struct {
		format   string
		expected bool
	}{
		{"html", true},
		{"json", true},
		{"markdown", false},
		{"", false},
		{"md", false},
	}
	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			result := IsDirectWriteFormat(tt.format)
			if result != tt.expected {
				t.Errorf("IsDirectWriteFormat(%q) = %v, want %v", tt.format, result, tt.expected)
			}
		})
	}
}

// --- Integration Tests ---

func TestLoadBuiltin_DefaultProducesValidOutput(t *testing.T) {
	loader := NewLoader(false)
	tmpl, err := loader.LoadBuiltin("default", TemplateTypeImage)
	if err != nil {
		t.Fatalf("LoadBuiltin error = %v", err)
	}

	// Verify the template can parse without error
	if tmpl.Name() != "default" {
		t.Errorf("template name = %q, want 'default'", tmpl.Name())
	}
}

func TestNewLoader_NoMojiFlag(t *testing.T) {
	loader := NewLoader(true)
	if loader == nil {
		t.Fatal("NewLoader returned nil")
	}
	if !loader.noMoji {
		t.Error("expected noMoji = true")
	}
}

func TestLoadFile_Caching(t *testing.T) {
	tmpDir := t.TempDir()
	tmplPath := filepath.Join(tmpDir, "cached.tmpl")
	content := `Cached: {{ .Name }}`
	if err := os.WriteFile(tmplPath, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write template file: %v", err)
	}

	loader := NewLoader(false)
	tmpl1, err := loader.LoadFile(tmplPath)
	if err != nil {
		t.Fatalf("first LoadFile error = %v", err)
	}

	tmpl2, err := loader.LoadFile(tmplPath)
	if err != nil {
		t.Fatalf("second LoadFile error = %v", err)
	}

	if tmpl1 != tmpl2 {
		t.Error("expected cached template to return same pointer")
	}
}
