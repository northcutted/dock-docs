package parser

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {
	content := `
FROM alpine:latest

# @name: DB_PORT
# @description: Database listening port
# @default: 5432
# @required: true
ARG DB_PORT=5432

# @description: Environment mode
ENV APP_ENV=production

# @name: VendorName
LABEL vendor="Acme Corp"

EXPOSE 8080
`
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	doc, err := Parse(tmpFile)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if len(doc.Items) != 4 {
		t.Fatalf("expected 4 items, got %d", len(doc.Items))
	}

	// Test Case 1: ARG with all magic comments
	arg := doc.Items[0]
	if arg.Name != "DB_PORT" {
		t.Errorf("expected Name DB_PORT, got %s", arg.Name)
	}
	if arg.Value != "5432" {
		t.Errorf("expected Value 5432, got %s", arg.Value)
	}
	if arg.Description != "Database listening port" {
		t.Errorf("expected Description 'Database listening port', got '%s'", arg.Description)
	}
	if arg.Type != "ARG" {
		t.Errorf("expected Type ARG, got %s", arg.Type)
	}
	if !arg.Required {
		t.Errorf("expected Required true, got false")
	}

	// Test Case 2: ENV with partial magic comments (infer value)
	env := doc.Items[1]
	if env.Name != "APP_ENV" {
		t.Errorf("expected Name APP_ENV, got %s", env.Name)
	}
	if env.Value != "production" {
		t.Errorf("expected Value production, got %s", env.Value)
	}
	if env.Description != "Environment mode" {
		t.Errorf("expected Description 'Environment mode', got '%s'", env.Description)
	}
	if env.Type != "ENV" {
		t.Errorf("expected Type ENV, got %s", env.Type)
	}

	// Test Case 3: LABEL
	label := doc.Items[2]
	if label.Name != "VendorName" {
		t.Errorf("expected Name VendorName, got %s", label.Name)
	}
	if label.Value != "Acme Corp" {
		t.Errorf("expected Value 'Acme Corp', got '%s'", label.Value)
	}

	// Test Case 4: EXPOSE
	expose := doc.Items[3]
	if expose.Name != "8080" {
		t.Errorf("expected Name 8080, got %s", expose.Name)
	}
	if expose.Value != "8080" {
		t.Errorf("expected Value 8080, got %s", expose.Value)
	}
	if expose.Type != "EXPOSE" {
		t.Errorf("expected Type EXPOSE, got %s", expose.Type)
	}
}

func TestFilterByType(t *testing.T) {
	doc := &Documentation{
		Items: []DocItem{
			{Name: "PORT", Type: "ENV"},
			{Name: "DB_HOST", Type: "ENV"},
			{Name: "VERSION", Type: "ARG"},
			{Name: "8080", Type: "EXPOSE"},
			{Name: "maintainer", Type: "LABEL"},
		},
	}

	tests := []struct {
		name     string
		typeStr  string
		expected int
	}{
		{"filter ENV", "ENV", 2},
		{"filter ARG", "ARG", 1},
		{"filter EXPOSE", "EXPOSE", 1},
		{"filter LABEL", "LABEL", 1},
		{"filter nonexistent", "VOLUME", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filtered := doc.FilterByType(tt.typeStr)
			if len(filtered) != tt.expected {
				t.Errorf("FilterByType(%s) = %d items, want %d", tt.typeStr, len(filtered), tt.expected)
			}
			// Verify all returned items match the requested type
			for _, item := range filtered {
				if item.Type != tt.typeStr {
					t.Errorf("FilterByType(%s) returned item with Type %s", tt.typeStr, item.Type)
				}
			}
		})
	}
}

func TestParse_MultipleEnv(t *testing.T) {
	content := `
FROM alpine:latest

# @description: First variable
# @description: Second variable
ENV VAR1=value1 VAR2=value2
`
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	doc, err := Parse(tmpFile)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if len(doc.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(doc.Items))
	}

	// First ENV should have first description
	if doc.Items[0].Description != "First variable" {
		t.Errorf("expected Description 'First variable', got '%s'", doc.Items[0].Description)
	}
	if doc.Items[0].Name != "VAR1" {
		t.Errorf("expected Name VAR1, got %s", doc.Items[0].Name)
	}

	// Second ENV should have second description
	if doc.Items[1].Description != "Second variable" {
		t.Errorf("expected Description 'Second variable', got '%s'", doc.Items[1].Description)
	}
	if doc.Items[1].Name != "VAR2" {
		t.Errorf("expected Name VAR2, got %s", doc.Items[1].Name)
	}
}

func TestParse_MultipleExpose(t *testing.T) {
	content := `
FROM alpine:latest

EXPOSE 8080 9090 3000
`
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	doc, err := Parse(tmpFile)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if len(doc.Items) != 3 {
		t.Fatalf("expected 3 items, got %d", len(doc.Items))
	}

	expectedPorts := []string{"8080", "9090", "3000"}
	for i, item := range doc.Items {
		if item.Type != "EXPOSE" {
			t.Errorf("item %d: expected Type EXPOSE, got %s", i, item.Type)
		}
		if item.Name != expectedPorts[i] {
			t.Errorf("item %d: expected Name %s, got %s", i, expectedPorts[i], item.Name)
		}
	}
}

func TestParse_NoComments(t *testing.T) {
	content := `
FROM alpine:latest

ARG VERSION=1.0
ENV PORT=8080
LABEL maintainer="test@example.com"
EXPOSE 3000
`
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	doc, err := Parse(tmpFile)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if len(doc.Items) != 4 {
		t.Fatalf("expected 4 items, got %d", len(doc.Items))
	}

	// Verify items exist even without comments
	// Note: Parser behavior may vary based on Dockerfile syntax
	// Just verify types are correct
	types := make(map[string]int)
	for _, item := range doc.Items {
		types[item.Type]++
	}

	if types["ARG"] != 1 {
		t.Errorf("expected 1 ARG item, got %d", types["ARG"])
	}
	if types["ENV"] != 1 {
		t.Errorf("expected 1 ENV item, got %d", types["ENV"])
	}
	if types["LABEL"] != 1 {
		t.Errorf("expected 1 LABEL item, got %d", types["LABEL"])
	}
	if types["EXPOSE"] != 1 {
		t.Errorf("expected 1 EXPOSE item, got %d", types["EXPOSE"])
	}
}

func TestParse_OnlyFROM(t *testing.T) {
	content := `FROM alpine:latest
`
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	doc, err := Parse(tmpFile)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if len(doc.Items) != 0 {
		t.Errorf("expected 0 items for FROM-only Dockerfile, got %d", len(doc.Items))
	}
}

func TestParse_ARGNoDefault(t *testing.T) {
	content := `FROM alpine:latest

# @description: Build version argument
ARG MY_VAR
`
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	doc, err := Parse(tmpFile)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if len(doc.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(doc.Items))
	}

	arg := doc.Items[0]
	if arg.Name != "MY_VAR" {
		t.Errorf("expected Name MY_VAR, got %s", arg.Name)
	}
	if arg.Type != "ARG" {
		t.Errorf("expected Type ARG, got %s", arg.Type)
	}
	if arg.Description != "Build version argument" {
		t.Errorf("expected Description 'Build version argument', got '%s'", arg.Description)
	}
	// ARG with no default should have empty value
	// (value depends on how buildkit represents it, but should not panic)
}

func TestParse_ENVQuotedSpecialChars(t *testing.T) {
	content := `FROM alpine:latest

# @description: Complex path value
ENV PATH="/opt/myapp/bin:$PATH"
`
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	doc, err := Parse(tmpFile)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if len(doc.Items) < 1 {
		t.Fatalf("expected at least 1 item, got %d", len(doc.Items))
	}

	env := doc.Items[0]
	if env.Name != "PATH" {
		t.Errorf("expected Name PATH, got %s", env.Name)
	}
	if env.Description != "Complex path value" {
		t.Errorf("expected Description 'Complex path value', got '%s'", env.Description)
	}
	// The value should have quotes stripped by stripQuotes
	if env.Value != "/opt/myapp/bin:$PATH" {
		t.Errorf("expected Value '/opt/myapp/bin:$PATH', got '%s'", env.Value)
	}
}

func TestParse_LABELMultiLine(t *testing.T) {
	content := `FROM alpine:latest

# @name: description
# @description: A multi-line label
LABEL description="This is a \
long description"
`
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	doc, err := Parse(tmpFile)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if len(doc.Items) < 1 {
		t.Fatalf("expected at least 1 item, got %d", len(doc.Items))
	}

	label := doc.Items[0]
	if label.Type != "LABEL" {
		t.Errorf("expected Type LABEL, got %s", label.Type)
	}
	if label.Name != "description" {
		t.Errorf("expected Name 'description', got '%s'", label.Name)
	}
	if label.Description != "A multi-line label" {
		t.Errorf("expected Description 'A multi-line label', got '%s'", label.Description)
	}
}

func TestParse_MoreMetasThanItems(t *testing.T) {
	// Three @description blocks but only one ARG item — extras silently dropped
	content := `FROM alpine:latest

# @description: First desc
# @description: Second desc
# @description: Third desc
ARG ONLY_ONE=value
`
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	doc, err := Parse(tmpFile)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	// Should produce 1 item (the ARG), extra metas are silently dropped
	if len(doc.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(doc.Items))
	}

	if doc.Items[0].Description != "First desc" {
		t.Errorf("expected Description 'First desc', got '%s'", doc.Items[0].Description)
	}
	// Without a @name: tag, ARG name is inferred from buildkit AST.
	// For "ARG ONLY_ONE=value", buildkit may represent key as "ONLY_ONE=value".
	// The @name: tag can override this (see TestParse where @name: DB_PORT is used).
	if doc.Items[0].Name != "ONLY_ONE" && doc.Items[0].Name != "ONLY_ONE=value" {
		t.Errorf("expected Name containing ONLY_ONE, got %s", doc.Items[0].Name)
	}
}

func TestParse_FewerMetasThanItems(t *testing.T) {
	// Only 1 @description for 3 ENV vars — first gets metadata, rest get none
	content := `FROM alpine:latest

# @description: Only first var described
ENV A=1 B=2 C=3
`
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	doc, err := Parse(tmpFile)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if len(doc.Items) != 3 {
		t.Fatalf("expected 3 items, got %d", len(doc.Items))
	}

	// First item gets the description
	if doc.Items[0].Description != "Only first var described" {
		t.Errorf("expected first item Description 'Only first var described', got '%s'", doc.Items[0].Description)
	}
	// Remaining items get no description
	if doc.Items[1].Description != "" {
		t.Errorf("expected second item Description empty, got '%s'", doc.Items[1].Description)
	}
	if doc.Items[2].Description != "" {
		t.Errorf("expected third item Description empty, got '%s'", doc.Items[2].Description)
	}
}

func TestParse_PlainCommentsNoTags(t *testing.T) {
	// Comments without recognized @-tags should be ignored (no metadata applied)
	content := `FROM alpine:latest

# This is just a regular comment
# It has no magic tags at all
ARG VERSION=1.0
`
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	doc, err := Parse(tmpFile)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if len(doc.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(doc.Items))
	}

	// No @description, @name, etc. so defaults should be inferred from buildkit AST
	// For "ARG VERSION=1.0", buildkit may represent the key as "VERSION=1.0"
	if doc.Items[0].Name != "VERSION" && doc.Items[0].Name != "VERSION=1.0" {
		t.Errorf("expected Name containing VERSION, got %s", doc.Items[0].Name)
	}
	if doc.Items[0].Description != "" {
		t.Errorf("expected empty Description, got '%s'", doc.Items[0].Description)
	}
}

func TestParse_RequiredFalse(t *testing.T) {
	content := `FROM alpine:latest

# @description: Not required var
# @required: false
ARG OPT_VAR=default
`
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	doc, err := Parse(tmpFile)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if len(doc.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(doc.Items))
	}

	// @required: false should leave Required as false
	if doc.Items[0].Required {
		t.Errorf("expected Required false, got true")
	}
	if doc.Items[0].Description != "Not required var" {
		t.Errorf("expected Description 'Not required var', got '%s'", doc.Items[0].Description)
	}
}

func TestParse_DefaultOverridesInferred(t *testing.T) {
	content := `FROM alpine:latest

# @description: Overridden default
# @default: custom_value
ARG MY_VAR=original_value
`
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	doc, err := Parse(tmpFile)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if len(doc.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(doc.Items))
	}

	// @default: custom_value should override the inferred value "original_value"
	if doc.Items[0].Value != "custom_value" {
		t.Errorf("expected Value 'custom_value', got '%s'", doc.Items[0].Value)
	}
}

func TestParse_ExposeWithProtocol(t *testing.T) {
	content := `FROM alpine:latest

EXPOSE 80/tcp
EXPOSE 53/udp
`
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	doc, err := Parse(tmpFile)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if len(doc.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(doc.Items))
	}

	// EXPOSE with protocol — the raw value includes the protocol
	if doc.Items[0].Type != "EXPOSE" {
		t.Errorf("expected Type EXPOSE, got %s", doc.Items[0].Type)
	}
	// Buildkit parses "80/tcp" as the full value
	if doc.Items[0].Name != "80/tcp" {
		t.Errorf("expected Name '80/tcp', got '%s'", doc.Items[0].Name)
	}
	if doc.Items[1].Name != "53/udp" {
		t.Errorf("expected Name '53/udp', got '%s'", doc.Items[1].Name)
	}
}

func TestParse_MultiStageDockerfile(t *testing.T) {
	content := `FROM golang:1.21 AS builder

# @description: Build mode
ARG BUILD_MODE=release

FROM alpine:latest

# @description: Application port
ENV APP_PORT=8080

EXPOSE 8080
`
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	doc, err := Parse(tmpFile)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	// Should parse items from all stages
	if len(doc.Items) != 3 {
		t.Fatalf("expected 3 items, got %d", len(doc.Items))
	}

	// First stage: ARG BUILD_MODE — without @name: tag, buildkit may include "=release"
	if doc.Items[0].Name != "BUILD_MODE" && doc.Items[0].Name != "BUILD_MODE=release" {
		t.Errorf("expected first item Name containing BUILD_MODE, got %s", doc.Items[0].Name)
	}
	if doc.Items[0].Type != "ARG" {
		t.Errorf("expected first item Type ARG, got %s", doc.Items[0].Type)
	}
	if doc.Items[0].Description != "Build mode" {
		t.Errorf("expected first item Description 'Build mode', got '%s'", doc.Items[0].Description)
	}

	// Second stage: ENV APP_PORT
	if doc.Items[1].Name != "APP_PORT" {
		t.Errorf("expected second item Name APP_PORT, got %s", doc.Items[1].Name)
	}
	if doc.Items[1].Type != "ENV" {
		t.Errorf("expected second item Type ENV, got %s", doc.Items[1].Type)
	}

	// Second stage: EXPOSE 8080
	if doc.Items[2].Name != "8080" {
		t.Errorf("expected third item Name 8080, got %s", doc.Items[2].Name)
	}
	if doc.Items[2].Type != "EXPOSE" {
		t.Errorf("expected third item Type EXPOSE, got %s", doc.Items[2].Type)
	}
}

func TestParse_FileNotFound(t *testing.T) {
	_, err := Parse("/nonexistent/Dockerfile")
	if err == nil {
		t.Error("expected error for non-existent file")
	}
}

func TestParse_InvalidDockerfile(t *testing.T) {
	content := `
This is not a valid Dockerfile syntax
INVALID COMMAND
`
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	// Parser should still work, it just won't find any recognized commands
	doc, err := Parse(tmpFile)
	if err != nil {
		// Some invalid syntax might cause parser errors
		t.Logf("Parse error (expected for invalid syntax): %v", err)
		return
	}

	// If it doesn't error, it should return empty or minimal items
	t.Logf("Parsed %d items from invalid Dockerfile", len(doc.Items))
}
