package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/northcutted/docker-docs/pkg/analysis"
	"github.com/northcutted/docker-docs/pkg/config"
	"github.com/northcutted/docker-docs/pkg/injector"
	"github.com/northcutted/docker-docs/pkg/parser"
	"github.com/northcutted/docker-docs/pkg/renderer"
	"github.com/northcutted/docker-docs/pkg/runner"
)

var (
	dockerfile string
	outputFile string
	dryRun     bool
	imageTag   string
	configFile string
)

var rootCmd = &cobra.Command{
	Use:   "docker-docs",
	Short: "Generate documentation from Dockerfile",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Detect Config Mode
		cfgPath := configFile
		if cfgPath == "" {
			if _, err := os.Stat("docker-docs.yaml"); err == nil {
				cfgPath = "docker-docs.yaml"
			}
		}

		if cfgPath != "" {
			return runConfigMode(cfgPath)
		}

		return runLegacyMode()
	},
}

func runConfigMode(path string) error {
	fmt.Printf("Using config file: %s\n", path)
	cfg, err := config.Load(path)
	if err != nil {
		return err
	}

	// Read Output File
	content, err := os.ReadFile(cfg.Output)
	if err != nil {
		// If output file doesn't exist, we can't inject.
		// Spec doesn't strictly say we must create it, but in Config mode,
		// usually we expect to update an existing file.
		return fmt.Errorf("failed to read output file %s: %w", cfg.Output, err)
	}
	fileContent := string(content)

	// Process Sections
	runners := []analysis.Runner{
		&runner.RuntimeRunner{},
		&runner.SyftRunner{},
		&runner.GrypeRunner{},
		&runner.DiveRunner{},
	}

	for _, section := range cfg.Sections {
		var sectionContent string

		switch section.Type {
		case config.SectionTypeConfig:
			// Parse Dockerfile
			dPath := section.Source
			if dPath == "" {
				dPath = "Dockerfile" // Default
			}
			doc, err := parser.Parse(dPath)
			if err != nil {
				return fmt.Errorf("failed to parse Dockerfile %s: %w", dPath, err)
			}

			// Analyze Image (optional)
			var stats *analysis.ImageStats
			if section.Image != "" {
				fmt.Printf("Analyzing image (config): %s ...\n", section.Image)
				stats, err = analysis.AnalyzeImage(section.Image, runners)
				if err != nil {
					fmt.Printf("Warning: analysis failed for %s: %v\n", section.Image, err)
				}
			}

			// Render
			sectionContent, err = renderer.Render(doc, stats)
			if err != nil {
				return fmt.Errorf("failed to render config section: %w", err)
			}

		case config.SectionTypeMatrix:
			if len(section.Images) == 0 {
				continue
			}
			fmt.Printf("Analyzing matrix: %v ...\n", section.Images)
			statsList, err := analysis.AnalyzeMatrix(section.Images, runners)
			if err != nil {
				return fmt.Errorf("matrix analysis failed: %w", err)
			}

			sectionContent, err = renderer.RenderMatrix(statsList)
			if err != nil {
				return fmt.Errorf("failed to render matrix section: %w", err)
			}

		default:
			fmt.Printf("Warning: unknown section type %s\n", section.Type)
			continue
		}

		// Inject
		newContent, err := injector.Inject(fileContent, section.Marker, sectionContent)
		if err != nil {
			// In config mode, missing markers is likely an error user wants to know about
			// But maybe we should just warn?
			fmt.Printf("Warning: %v\n", err)
			continue
		}
		fileContent = newContent
	}

	if dryRun {
		fmt.Println(fileContent)
		return nil
	}

	if err := os.WriteFile(cfg.Output, []byte(fileContent), 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}
	fmt.Printf("Updated %s\n", cfg.Output)
	return nil
}

func runLegacyMode() error {
	// 1. Parse Dockerfile
	doc, err := parser.Parse(dockerfile)
	if err != nil {
		return fmt.Errorf("failed to parse Dockerfile: %w", err)
	}

	// 2. Dynamic Analysis (if requested)
	var stats *analysis.ImageStats
	if imageTag != "" {
		fmt.Printf("Analyzing image: %s ...\n", imageTag)
		runners := []analysis.Runner{
			&runner.RuntimeRunner{},
			&runner.SyftRunner{},
			&runner.GrypeRunner{},
			&runner.DiveRunner{},
		}
		stats, err = analysis.AnalyzeImage(imageTag, runners)
		if err != nil {
			fmt.Printf("Warning: analysis failed: %v\n", err)
		}
	}

	// 3. Render
	renderedContent, err := renderer.Render(doc, stats)
	if err != nil {
		return fmt.Errorf("failed to render documentation: %w", err)
	}

	// 4. Output Strategy
	if dryRun {
		fmt.Println(renderedContent)
		return nil
	}

	// Check if output file exists
	content, err := os.ReadFile(outputFile)
	if err != nil {
		if os.IsNotExist(err) {
			if !dryRun {
				fmt.Fprintf(os.Stderr, "Warning: Output file '%s' does not exist. Printing to stdout.\n", outputFile)
			}
			fmt.Println(renderedContent)
			return nil
		}
		return err
	}

	fileContent := string(content)
	// Legacy mode uses default markers (empty name)
	newContent, err := injector.Inject(fileContent, "", renderedContent)
	if err != nil {
		if !dryRun {
			fmt.Fprintf(os.Stderr, "Warning: %v. Printing to stdout.\n", err)
		}
		fmt.Println(renderedContent)
		return nil
	}

	if err := os.WriteFile(outputFile, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}
	fmt.Printf("Updated %s\n", outputFile)

	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&dockerfile, "file", "f", "./Dockerfile", "Path to Dockerfile")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "README.md", "Path to output file")
	rootCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Print to stdout instead of writing to file")
	rootCmd.Flags().StringVar(&imageTag, "image", "", "Docker image tag to analyze (e.g. my-app:latest)")
	rootCmd.Flags().StringVar(&configFile, "config", "", "Path to config file (default: docker-docs.yaml)")
}
