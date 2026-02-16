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
