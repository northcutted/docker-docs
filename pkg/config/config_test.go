package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_ValidConfig(t *testing.T) {
	validYAML := `output: "DOCS.md"
badgeBaseURL: "https://custom-badges.example.com/static/v1"
sections:
  - type: "config"
    marker: "main"
    source: "Dockerfile"
    image: "myapp:latest"
  - type: "matrix"
    marker: "comparison"
    images:
      - "alpine:3.18"
      - "ubuntu:22.04"
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

	// Test first section (config type)
	section1 := cfg.Sections[0]
	if section1.Type != SectionTypeConfig {
		t.Errorf("Section[0].Type = %v, want %v", section1.Type, SectionTypeConfig)
	}
	if section1.Marker != "main" {
		t.Errorf("Section[0].Marker = %v, want main", section1.Marker)
	}
	if section1.Source != "Dockerfile" {
		t.Errorf("Section[0].Source = %v, want Dockerfile", section1.Source)
	}
	if section1.Image != "myapp:latest" {
		t.Errorf("Section[0].Image = %v, want myapp:latest", section1.Image)
	}

	// Test second section (matrix type)
	section2 := cfg.Sections[1]
	if section2.Type != SectionTypeMatrix {
		t.Errorf("Section[1].Type = %v, want %v", section2.Type, SectionTypeMatrix)
	}
	if section2.Marker != "comparison" {
		t.Errorf("Section[1].Marker = %v, want comparison", section2.Marker)
	}
	if len(section2.Images) != 2 {
		t.Fatalf("expected 2 images in matrix, got %d", len(section2.Images))
	}
	if section2.Images[0] != "alpine:3.18" {
		t.Errorf("Section[1].Images[0] = %v, want alpine:3.18", section2.Images[0])
	}
	if section2.Images[1] != "ubuntu:22.04" {
		t.Errorf("Section[1].Images[1] = %v, want ubuntu:22.04", section2.Images[1])
	}
}

func TestLoad_Defaults(t *testing.T) {
	minimalYAML := `sections:
  - type: "config"
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
  - type: config
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

	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	// Empty file should load with defaults
	if cfg.Output != "README.md" {
		t.Errorf("Output = %v, want README.md", cfg.Output)
	}

	if len(cfg.Sections) != 0 {
		t.Errorf("expected 0 sections, got %d", len(cfg.Sections))
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

	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.Output != "CUSTOM.md" {
		t.Errorf("Output = %v, want CUSTOM.md", cfg.Output)
	}

	if cfg.BadgeBaseURL != "https://img.shields.io/static/v1" {
		t.Errorf("BadgeBaseURL should use default")
	}
}

func TestLoad_ComplexMatrixConfig(t *testing.T) {
	complexYAML := `output: "README.md"
sections:
  - type: "matrix"
    marker: "versions"
    images:
      - "python:3.10"
      - "python:3.11"
      - "python:3.12"
      - "python:3.10-alpine"
      - "python:3.11-alpine"
      - "python:3.12-alpine"
`

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "matrix.yaml")
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
	if section.Type != SectionTypeMatrix {
		t.Errorf("Type = %v, want matrix", section.Type)
	}

	if len(section.Images) != 6 {
		t.Errorf("expected 6 images, got %d", len(section.Images))
	}
}

func TestLoad_MultipleConfigSections(t *testing.T) {
	multipleConfigsYAML := `output: "README.md"
sections:
  - type: "config"
    marker: "production"
    source: "Dockerfile.prod"
    image: "myapp:prod"
  - type: "config"
    marker: "development"
    source: "Dockerfile.dev"
    image: "myapp:dev"
  - type: "config"
    marker: "testing"
    source: "Dockerfile.test"
`

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "multi-config.yaml")
	err := os.WriteFile(configPath, []byte(multipleConfigsYAML), 0644)
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
	expectedImages := []string{"myapp:prod", "myapp:dev", ""}

	for i, section := range cfg.Sections {
		if section.Type != SectionTypeConfig {
			t.Errorf("Section[%d].Type = %v, want config", i, section.Type)
		}
		if section.Marker != expectedMarkers[i] {
			t.Errorf("Section[%d].Marker = %v, want %v", i, section.Marker, expectedMarkers[i])
		}
		if section.Source != expectedSources[i] {
			t.Errorf("Section[%d].Source = %v, want %v", i, section.Source, expectedSources[i])
		}
		if section.Image != expectedImages[i] {
			t.Errorf("Section[%d].Image = %v, want %v", i, section.Image, expectedImages[i])
		}
	}
}

func TestSectionType_Constants(t *testing.T) {
	if SectionTypeConfig != "config" {
		t.Errorf("SectionTypeConfig = %v, want 'config'", SectionTypeConfig)
	}

	if SectionTypeMatrix != "matrix" {
		t.Errorf("SectionTypeMatrix = %v, want 'matrix'", SectionTypeMatrix)
	}
}
