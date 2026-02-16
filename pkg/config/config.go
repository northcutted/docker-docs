package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type SectionType string

const (
	SectionTypeConfig SectionType = "config"
	SectionTypeMatrix SectionType = "matrix"
)

type Section struct {
	Type   SectionType `yaml:"type"`
	Marker string      `yaml:"marker"`
	// Config specific
	Source string `yaml:"source,omitempty"` // Dockerfile path
	Image  string `yaml:"image,omitempty"`  // Single image for config analysis
	// Matrix specific
	Images []string `yaml:"images,omitempty"`
}

type Config struct {
	Output       string    `yaml:"output"`
	BadgeBaseURL string    `yaml:"badgeBaseURL,omitempty"`
	Sections     []Section `yaml:"sections"`
}

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
