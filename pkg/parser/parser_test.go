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
