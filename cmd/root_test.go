package cmd

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/northcutted/dock-docs/pkg/config"
	"github.com/northcutted/dock-docs/pkg/renderer"
)

// Helper to capture stdout
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	// These errors are typically ignored in test helpers, but linter complains
	_ = w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	return buf.String()
}

func TestExecute_DryRun(t *testing.T) {
	// Setup dummy Dockerfile
	tmpDir := t.TempDir()
	dockerfile := filepath.Join(tmpDir, "Dockerfile")
	content := "FROM alpine\nENV APP_PORT=8080"
	if err := os.WriteFile(dockerfile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write Dockerfile: %v", err)
	}

	// Override flags for test context
	// Cobra flags persist, so we need to reset or run in separate processes usually.
	// But since we are testing main execution logic, we can just call the RunE function directly or use SetArgs.

	// Reset flags for safety in this test run context
	rootCmd.SetArgs([]string{"--file", dockerfile, "--dry-run"})

	output := captureOutput(func() {
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("Execute failed: %v", err)
		}
	})

	if !strings.Contains(output, "`APP_PORT`") {
		t.Errorf("expected dry-run output to contain table row, got:\n%s", output)
	}
}

func TestExecute_Injection(t *testing.T) {
	// Reset flags manually because Cobra holds state in globals/pointers
	dockerfile = ""
	outputFile = ""
	dryRun = false
	imageTag = ""

	tmpDir := t.TempDir()
	dockerfileLoc := filepath.Join(tmpDir, "Dockerfile")
	readme := filepath.Join(tmpDir, "README.md")

	// Dockerfile content
	if err := os.WriteFile(dockerfileLoc, []byte("FROM alpine\nENV FOO=bar"), 0644); err != nil {
		t.Fatalf("failed to write Dockerfile: %v", err)
	}

	// README content with markers
	readmeContent := "# Title\n\n<!-- BEGIN: dock-docs -->\nOLD CONTENT\n<!-- END: dock-docs -->\n\nFooter"
	if err := os.WriteFile(readme, []byte(readmeContent), 0644); err != nil {
		t.Fatalf("failed to write README: %v", err)
	}

	// Create a new command for testing to avoid shared state issues or reset flags?
	// Cobra commands are hard to re-use if they bind to globals.
	// We will try setting args explicitly and hoping Execute resets parsing.
	rootCmd.SetArgs([]string{"--file", dockerfileLoc, "--output", readme})

	captureOutput(func() {
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("Execute failed: %v", err)
		}
	})

	// Verify file updated
	newContent, err := os.ReadFile(readme)
	if err != nil {
		t.Fatalf("failed to read updated README: %v", err)
	}

	sContent := string(newContent)
	if strings.Contains(sContent, "OLD CONTENT") {
		t.Error("expected OLD CONTENT to be replaced")
	}
	if !strings.Contains(sContent, "`FOO`") {
		t.Error("expected new table content in README")
	}
	if !strings.Contains(sContent, "<!-- BEGIN: dock-docs -->") {
		t.Error("expected markers to be preserved")
	}
}

func TestExecute_NoMarkers_Stdout(t *testing.T) {
	dockerfile = ""
	outputFile = ""
	dryRun = false
	imageTag = ""

	tmpDir := t.TempDir()
	dockerfile := filepath.Join(tmpDir, "Dockerfile")
	readme := filepath.Join(tmpDir, "README.md")

	if err := os.WriteFile(dockerfile, []byte("FROM alpine\nENV BAZ=qux"), 0644); err != nil {
		t.Fatalf("failed to write Dockerfile: %v", err)
	}

	// README without markers
	if err := os.WriteFile(readme, []byte("# Just a file"), 0644); err != nil {
		t.Fatalf("failed to write README: %v", err)
	}

	rootCmd.SetArgs([]string{"--file", dockerfile, "--output", readme})

	output := captureOutput(func() {
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("Execute failed: %v", err)
		}
	})

	// Should print to stdout because markers are missing
	if !strings.Contains(output, "`BAZ`") {
		t.Errorf("expected stdout output when markers are missing, got: %s", output)
	}

	// File should be unchanged
	content, _ := os.ReadFile(readme)
	if strings.Contains(string(content), "`BAZ`") {
		t.Error("expected file to remain unchanged when markers are missing")
	}
}

func TestExecute_WithImageFlag_FailsWithoutIgnore(t *testing.T) {
	// Smoke test for --image flag failing without ignore-errors
	dockerfile = ""
	outputFile = ""
	dryRun = false
	imageTag = ""
	ignoreErrors = false

	tmpDir := t.TempDir()
	dockerfile := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(dockerfile, []byte("FROM alpine"), 0644); err != nil {
		t.Fatalf("failed to write Dockerfile: %v", err)
	}

	rootCmd.SetArgs([]string{"--file", dockerfile, "--image", "fake-image:latest", "--dry-run"})

	output := captureOutput(func() {
		if err := rootCmd.Execute(); err == nil {
			t.Fatal("Execute expected to fail, but it succeeded")
		}
	})

	// It should print warning about analysis failure
	if !strings.Contains(output, "Warning: analysis failed") {
		t.Errorf("expected analysis warning, got:\n%s", output)
	}
}

func TestExecute_WithImageFlag_IgnoresErrors(t *testing.T) {
	// Smoke test for --image flag with --ignore-errors
	dockerfile = ""
	outputFile = ""
	dryRun = false
	imageTag = ""
	ignoreErrors = false

	tmpDir := t.TempDir()
	dockerfile := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(dockerfile, []byte("FROM alpine"), 0644); err != nil {
		t.Fatalf("failed to write Dockerfile: %v", err)
	}

	// Add --ignore-errors flag
	rootCmd.SetArgs([]string{"--file", dockerfile, "--image", "fake-image:latest", "--dry-run", "--ignore-errors"})

	output := captureOutput(func() {
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("Execute failed despite ignore-errors: %v", err)
		}
	})

	// It should print "Analyzing image: fake-image:latest ..."
	if !strings.Contains(output, "Analyzing image: fake-image:latest") {
		t.Errorf("expected analysis log, got:\n%s", output)
	}

	// And standard table (because it proceeds to render even if analysis fails)
	if !strings.Contains(output, "Configuration") {
		t.Errorf("expected standard table, got:\n%s", output)
	}
}

// --- Pure function unit tests ---

func TestResolveOutputPath(t *testing.T) {
	tests := []struct {
		name          string
		currentOutput string
		format        string
		expected      string
	}{
		{"default markdown", "README.md", "markdown", "README.md"},
		{"default to html", "README.md", "html", "README.html"},
		{"default to json", "README.md", "json", "README.json"},
		{"custom output kept for html", "DOCS.md", "html", "DOCS.md"},
		{"custom output kept for json", "output.txt", "json", "output.txt"},
		{"custom output kept for markdown", "CUSTOM.md", "markdown", "CUSTOM.md"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := resolveOutputPath(tt.currentOutput, tt.format)
			if result != tt.expected {
				t.Errorf("resolveOutputPath(%q, %q) = %q, want %q", tt.currentOutput, tt.format, result, tt.expected)
			}
		})
	}
}

func TestResolveSectionOutput(t *testing.T) {
	tests := []struct {
		name         string
		baseOutput   string
		marker       string
		sectionIndex int
		format       string
		expected     string
	}{
		{"marker html", "README.md", "main", 0, "html", "README-main.html"},
		{"marker json", "README.md", "compare", 1, "json", "README-compare.json"},
		{"no marker uses index", "README.md", "", 2, "html", "README-section2.html"},
		{"subdirectory", "docs/README.md", "main", 0, "json", filepath.Join("docs", "README-main.json")},
		{"custom base name", "DOCS.md", "overview", 0, "html", "DOCS-overview.html"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := resolveSectionOutput(tt.baseOutput, tt.marker, tt.sectionIndex, tt.format)
			if result != tt.expected {
				t.Errorf("resolveSectionOutput(%q, %q, %d, %q) = %q, want %q",
					tt.baseOutput, tt.marker, tt.sectionIndex, tt.format, result, tt.expected)
			}
		})
	}
}

func TestDescribeTemplate(t *testing.T) {
	tests := []struct {
		name     string
		sel      renderer.TemplateSelection
		expected string
	}{
		{"custom path", renderer.TemplateSelection{Path: "/my/template.tmpl"}, "custom file: /my/template.tmpl"},
		{"named builtin", renderer.TemplateSelection{Name: "minimal"}, "built-in: minimal"},
		{"empty defaults", renderer.TemplateSelection{}, "built-in: default"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := describeTemplate(tt.sel)
			if result != tt.expected {
				t.Errorf("describeTemplate() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestResolveTemplateSel(t *testing.T) {
	tests := []struct {
		name         string
		cliTemplate  string // value of templateName global
		cfgTemplate  *config.TemplateConfig
		expectedName string
		expectedPath string
	}{
		{"default empty", "", nil, "", ""},
		{"cli name", "minimal", nil, "minimal", ""},
		{"cli file path with slash", "templates/custom.tmpl", nil, "", "templates/custom.tmpl"},
		{"cli file path with tmpl", "custom.tmpl", nil, "", "custom.tmpl"},
		{"config name", "", &config.TemplateConfig{Name: "detailed"}, "detailed", ""},
		{"config path", "", &config.TemplateConfig{Path: "/path/to/tmpl"}, "", "/path/to/tmpl"},
		{"cli overrides config", "html", &config.TemplateConfig{Name: "minimal"}, "html", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set the global
			oldTemplateName := templateName
			templateName = tt.cliTemplate
			defer func() { templateName = oldTemplateName }()

			result := resolveTemplateSel(tt.cfgTemplate)
			if result.Name != tt.expectedName {
				t.Errorf("resolveTemplateSel().Name = %q, want %q", result.Name, tt.expectedName)
			}
			if result.Path != tt.expectedPath {
				t.Errorf("resolveTemplateSel().Path = %q, want %q", result.Path, tt.expectedPath)
			}
		})
	}
}

func TestHandleListTemplates(t *testing.T) {
	output := captureOutput(func() {
		if err := handleListTemplates(); err != nil {
			t.Fatalf("handleListTemplates() error = %v", err)
		}
	})

	if !strings.Contains(output, "Available built-in templates:") {
		t.Error("expected header in output")
	}
	// Should list known templates
	for _, name := range []string{"default", "minimal", "detailed", "compact", "html", "json"} {
		if !strings.Contains(output, name) {
			t.Errorf("expected template %q in output", name)
		}
	}
	if !strings.Contains(output, "dock-docs --template") {
		t.Error("expected usage hint in output")
	}
}

func TestHandleExportTemplate(t *testing.T) {
	// Happy path
	output := captureOutput(func() {
		if err := handleExportTemplate("default"); err != nil {
			t.Fatalf("handleExportTemplate(default) error = %v", err)
		}
	})
	if !strings.Contains(output, "Docker Image Analysis") {
		t.Error("expected template content in output")
	}

	// Error path: unknown template
	err := handleExportTemplate("nonexistent")
	if err == nil {
		t.Fatal("expected error for unknown template")
	}
	if !strings.Contains(err.Error(), "unknown built-in template") {
		t.Errorf("error = %q, want it to contain 'unknown built-in template'", err.Error())
	}
}

func TestHandleValidateTemplate(t *testing.T) {
	tmpDir := t.TempDir()

	// Valid template
	validPath := filepath.Join(tmpDir, "valid.tmpl")
	if err := os.WriteFile(validPath, []byte("# Hello {{ .ImageTag }}"), 0644); err != nil {
		t.Fatalf("failed to write file: %v", err)
	}
	output := captureOutput(func() {
		if err := handleValidateTemplate(validPath); err != nil {
			t.Fatalf("handleValidateTemplate() error = %v for valid template", err)
		}
	})
	if !strings.Contains(output, "is valid") {
		t.Error("expected 'is valid' message")
	}

	// Invalid template
	invalidPath := filepath.Join(tmpDir, "invalid.tmpl")
	if err := os.WriteFile(invalidPath, []byte("{{ .Unclosed"), 0644); err != nil {
		t.Fatalf("failed to write file: %v", err)
	}
	err := handleValidateTemplate(invalidPath)
	if err == nil {
		t.Fatal("expected error for invalid template")
	}

	// Nonexistent file
	err = handleValidateTemplate("/nonexistent/path.tmpl")
	if err == nil {
		t.Fatal("expected error for nonexistent file")
	}
}

func TestVersionCommand(t *testing.T) {
	// Set version globals for test
	oldVersion, oldCommit, oldDate := Version, Commit, Date
	Version = "1.2.3"
	Commit = "abc123"
	Date = "2025-01-01"
	defer func() {
		Version, Commit, Date = oldVersion, oldCommit, oldDate
	}()

	rootCmd.SetArgs([]string{"version"})
	output := captureOutput(func() {
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("version command failed: %v", err)
		}
	})

	if !strings.Contains(output, "1.2.3") {
		t.Errorf("expected version in output, got: %s", output)
	}
	if !strings.Contains(output, "abc123") {
		t.Errorf("expected commit in output, got: %s", output)
	}
	if !strings.Contains(output, "2025-01-01") {
		t.Errorf("expected date in output, got: %s", output)
	}
}

func TestCheckToolStatus(t *testing.T) {
	result := checkToolStatus()
	if !strings.Contains(result, "Prerequisites:") {
		t.Errorf("expected 'Prerequisites:' in output, got: %s", result)
	}
	// Should mention docker/podman
	if !strings.Contains(result, "docker") && !strings.Contains(result, "podman") {
		t.Error("expected docker or podman mention in tool status")
	}
	// Should mention each tool
	for _, tool := range []string{"syft", "grype", "dive"} {
		if !strings.Contains(result, tool) {
			t.Errorf("expected %q in tool status", tool)
		}
	}
}

func TestExecute_ListTemplatesFlag(t *testing.T) {
	rootCmd.SetArgs([]string{"--list-templates"})
	output := captureOutput(func() {
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("Execute with --list-templates failed: %v", err)
		}
	})
	if !strings.Contains(output, "Available built-in templates:") {
		t.Error("expected template listing output")
	}
}

func TestExecute_DryRunWithHTMLTemplate(t *testing.T) {
	tmpDir := t.TempDir()
	df := filepath.Join(tmpDir, "Dockerfile")
	if err := os.WriteFile(df, []byte("FROM alpine\nENV PORT=8080"), 0644); err != nil {
		t.Fatalf("failed to write Dockerfile: %v", err)
	}

	rootCmd.SetArgs([]string{"--file", df, "--dry-run", "--template", "html"})
	output := captureOutput(func() {
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("Execute with html template failed: %v", err)
		}
	})
	// HTML output should contain HTML tags
	if !strings.Contains(output, "<") {
		t.Error("expected HTML content in dry-run output")
	}
}
