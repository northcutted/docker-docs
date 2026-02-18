package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/northcutted/dock-docs/pkg/analysis"
	"github.com/northcutted/dock-docs/pkg/config"
	"github.com/northcutted/dock-docs/pkg/injector"
	"github.com/northcutted/dock-docs/pkg/installer"
	"github.com/northcutted/dock-docs/pkg/parser"
	"github.com/northcutted/dock-docs/pkg/renderer"
	"github.com/northcutted/dock-docs/pkg/runner"
	"github.com/northcutted/dock-docs/pkg/templates"
	"github.com/northcutted/dock-docs/pkg/types"
)

var (
	dockerfile       string
	outputFile       string
	dryRun           bool
	imageTag         string
	configFile       string
	noMoji           bool
	ignoreErrors     bool
	verbose          bool
	badgeBaseURL     string
	templateName     string
	listTemplates    bool
	exportTemplate   string
	validateTemplate string
	debugTemplate    bool
)

var rootCmd = &cobra.Command{
	Use:   "dock-docs",
	Short: "Generate documentation from Dockerfile",
	Long: `Generate comprehensive documentation from your Dockerfiles.

This tool automatically parses your Dockerfile to create standard documentation
tables for arguments, environment variables, and exposed ports.

It can also perform deep analysis on built images using:
- Syft (for SBOM and package listing)
- Grype (for vulnerability scanning)
- Dive (for layer efficiency analysis)

Modes:
- YAML Mode: Uses 'dock-docs.yaml' for advanced configuration and multi-image comparison analysis.
- CLI Mode: Runs on a single Dockerfile/image without configuration.`,
	Example: `  # YAML Mode (Recommended)
  dock-docs --config dock-docs.yaml

  # CLI Mode: Analyze Dockerfile only
  dock-docs -f ./Dockerfile

  # CLI Mode: Analyze Dockerfile and Image
  dock-docs -f ./Dockerfile --image my-app:latest

  # CLI Mode: Output to specific file
  dock-docs -f ./Dockerfile -o DOCUMENTATION.md`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Handle template developer tool flags first (these exit early)
		if listTemplates {
			return handleListTemplates()
		}
		if exportTemplate != "" {
			return handleExportTemplate(exportTemplate)
		}
		if validateTemplate != "" {
			return handleValidateTemplate(validateTemplate)
		}

		// Detect YAML Mode
		cfgPath := configFile
		if cfgPath == "" {
			if _, err := os.Stat("dock-docs.yaml"); err == nil {
				cfgPath = "dock-docs.yaml"
			}
		}

		if cfgPath != "" {
			return runYAMLMode(cfgPath)
		}

		return runCLIMode()
	},
}

func runYAMLMode(path string) error {
	fmt.Printf("Using config file: %s\n", path)
	cfg, err := config.Load(path)
	if err != nil {
		return err
	}

	// Change working directory to config file location
	// This ensures relative paths in config (like output, source) are resolved correctly
	configDir := filepath.Dir(path)
	if configDir != "." {
		if err := os.Chdir(configDir); err != nil {
			return fmt.Errorf("failed to change directory to %s: %w", configDir, err)
		}
		fmt.Printf("Changed working directory to: %s\n", configDir)
	}

	// Read Output File (only needed for markdown injection; read lazily below)
	var fileContent string
	var fileContentLoaded bool

	loadFileContent := func() error {
		if fileContentLoaded {
			return nil
		}
		content, err := os.ReadFile(cfg.Output)
		if err != nil {
			return fmt.Errorf("failed to read output file %s: %w", cfg.Output, err)
		}
		fileContent = string(content)
		fileContentLoaded = true
		return nil
	}

	// Process Sections
	runners := []analysis.Runner{
		&runner.RuntimeRunner{},
		&runner.ManifestRunner{},
		&runner.SyftRunner{},
		&runner.GrypeRunner{},
		&runner.DiveRunner{},
	}

	renderOpts := renderer.RenderOptions{
		NoMoji:       noMoji,
		BadgeBaseURL: cfg.BadgeBaseURL,
	}

	for i, section := range cfg.Sections {
		var sectionContent string

		// Resolve template: CLI flag > section config > global config > default
		tmplSel := resolveTemplateSel(cfg.ResolveTemplate(section))
		format := tmplSel.Format()

		switch section.Type {
		case config.SectionTypeImage:
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
			var stats *types.ImageStats
			if section.Tag != "" {
				fmt.Printf("Analyzing image: %s ...\n", section.Tag)
				stats, err = analysis.AnalyzeImage(section.Tag, runners, verbose)
				if err != nil {
					fmt.Printf("Warning: analysis failed for %s: %v\n", section.Tag, err)
					if !ignoreErrors {
						return fmt.Errorf("analysis failed for %s: %w", section.Tag, err)
					}
				}
			}

			if debugTemplate {
				fmt.Printf("Template: %s (type: image, format: %s)\n", describeTemplate(tmplSel), format)
			}

			// Render
			sectionContent, err = renderer.RenderWithTemplate(doc, stats, renderOpts, tmplSel)
			if err != nil {
				return fmt.Errorf("failed to render image section: %w", err)
			}

		case config.SectionTypeComparison:
			if len(section.Images) == 0 {
				continue
			}

			// Extract tags from ImageEntry structs for analysis
			resolvedImages := section.ResolvedImages()
			tags := make([]string, len(resolvedImages))
			for j, entry := range resolvedImages {
				tags[j] = entry.Tag
			}

			fmt.Printf("Analyzing comparison: %v ...\n", tags)
			statsList, err := analysis.AnalyzeComparison(tags, runners, verbose)
			if err != nil {
				return fmt.Errorf("comparison analysis failed: %w", err)
			}

			if debugTemplate {
				fmt.Printf("Template: %s (type: comparison, format: %s)\n", describeTemplate(tmplSel), format)
			}

			sectionContent, err = renderer.RenderComparisonWithTemplate(statsList, renderOpts, tmplSel)
			if err != nil {
				return fmt.Errorf("failed to render comparison section: %w", err)
			}

		default:
			fmt.Printf("Warning: unknown section type %s\n", section.Type)
			continue
		}

		// Output: direct-write for html/json, inject for markdown
		if templates.IsDirectWriteFormat(format) {
			outPath := resolveSectionOutput(cfg.Output, section.Marker, i, format)

			if dryRun {
				fmt.Printf("--- %s ---\n", outPath)
				fmt.Println(sectionContent)
				continue
			}

			if err := os.WriteFile(outPath, []byte(sectionContent), 0644); err != nil {
				return fmt.Errorf("failed to write output file %s: %w", outPath, err)
			}
			fmt.Printf("Wrote %s\n", outPath)
		} else {
			// Markdown: inject into existing file between markers
			if err := loadFileContent(); err != nil {
				return err
			}
			newContent, err := injector.Inject(fileContent, section.Marker, sectionContent)
			if err != nil {
				fmt.Printf("Warning: %v\n", err)
				continue
			}
			fileContent = newContent
		}
	}

	// Write the markdown output file if we modified it
	if fileContentLoaded {
		if dryRun {
			fmt.Println(fileContent)
			return nil
		}

		if err := os.WriteFile(cfg.Output, []byte(fileContent), 0644); err != nil {
			return fmt.Errorf("failed to write output file: %w", err)
		}
		fmt.Printf("Updated %s\n", cfg.Output)
	}

	return nil
}

func runCLIMode() error {
	// 1. Parse Dockerfile
	doc, err := parser.Parse(dockerfile)
	if err != nil {
		return fmt.Errorf("failed to parse Dockerfile: %w", err)
	}

	// 2. Dynamic Analysis (if requested)
	var stats *types.ImageStats
	if imageTag != "" {
		fmt.Printf("Analyzing image: %s ...\n", imageTag)
		runners := []analysis.Runner{
			&runner.RuntimeRunner{},
			&runner.ManifestRunner{},
			&runner.SyftRunner{},
			&runner.GrypeRunner{},
			&runner.DiveRunner{},
		}
		stats, err = analysis.AnalyzeImage(imageTag, runners, verbose)
		if err != nil {
			fmt.Printf("Warning: analysis failed: %v\n", err)
			if !ignoreErrors {
				return fmt.Errorf("analysis failed: %w", err)
			}
		}
	}

	// 3. Resolve template selection: CLI flag > default
	tmplSel := resolveTemplateSel(nil)

	if debugTemplate {
		fmt.Printf("Template: %s (type: image, format: %s)\n", describeTemplate(tmplSel), tmplSel.Format())
	}

	// 4. Render
	renderOpts := renderer.RenderOptions{
		NoMoji:       noMoji,
		BadgeBaseURL: badgeBaseURL,
	}
	renderedContent, err := renderer.RenderWithTemplate(doc, stats, renderOpts, tmplSel)
	if err != nil {
		return fmt.Errorf("failed to render documentation: %w", err)
	}

	// 5. Output Strategy
	if dryRun {
		fmt.Println(renderedContent)
		return nil
	}

	format := tmplSel.Format()

	// For HTML/JSON: write the complete standalone document directly to a file
	if templates.IsDirectWriteFormat(format) {
		outPath := resolveOutputPath(outputFile, format)
		if err := os.WriteFile(outPath, []byte(renderedContent), 0644); err != nil {
			return fmt.Errorf("failed to write output file: %w", err)
		}
		fmt.Printf("Wrote %s\n", outPath)
		return nil
	}

	// For Markdown: inject into existing file between markers
	content, err := os.ReadFile(outputFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Warning: Output file '%s' does not exist. Printing to stdout.\n", outputFile)
			fmt.Println(renderedContent)
			return nil
		}
		return err
	}

	fileContent := string(content)
	// Simple mode uses default markers (empty name)
	newContent, err := injector.Inject(fileContent, "", renderedContent)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: %v. Printing to stdout.\n", err)
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

// checkToolStatus returns a string indicating the status of required tools.
func checkToolStatus() string {
	tools := []string{"syft", "grype", "dive"}
	var status strings.Builder
	status.WriteString("\nPrerequisites:\n")

	// Check for Docker or Podman
	if _, err := exec.LookPath("docker"); err == nil {
		status.WriteString("  [OK] docker\n")
	} else if _, err := exec.LookPath("podman"); err == nil {
		status.WriteString("  [OK] podman\n")
	} else {
		status.WriteString("  [MISSING] docker or podman (required for dynamic analysis)\n")
	}

	for _, tool := range tools {
		if path, source, err := installer.FindTool(tool); err == nil {
			status.WriteString(fmt.Sprintf("  [OK] %s (%s: %s)\n", tool, source, path))
		} else {
			status.WriteString(fmt.Sprintf("  [MISSING] %s (run 'dock-docs setup' to install)\n", tool))
		}
	}
	return status.String()
}

func init() {
	// Dynamically append tool status to the help description
	rootCmd.Long += "\n" + checkToolStatus()

	rootCmd.Flags().StringVarP(&dockerfile, "file", "f", "./Dockerfile", "Path to Dockerfile (CLI Mode only)")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "README.md", "Path to output file")
	rootCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Print to stdout instead of writing to file")
	rootCmd.Flags().StringVar(&imageTag, "image", "", "Docker image tag to analyze (e.g. my-app:latest) (CLI Mode only)")
	rootCmd.Flags().StringVar(&configFile, "config", "", "Path to config file (default: dock-docs.yaml)")
	rootCmd.Flags().BoolVar(&noMoji, "nomoji", false, "Disable emojis in the output")
	rootCmd.Flags().BoolVar(&ignoreErrors, "ignore-errors", false, "Ignore analysis errors and continue (default false)")
	rootCmd.Flags().BoolVar(&verbose, "verbose", false, "Enable verbose logging")
	rootCmd.Flags().StringVar(&badgeBaseURL, "badge-base-url", "https://img.shields.io/static/v1", "Base URL for badge generation (e.g. for self-hosted shields.io)")

	// Template flags
	rootCmd.Flags().StringVar(&templateName, "template", "", "Template to use (built-in name or file path)")
	rootCmd.Flags().BoolVar(&listTemplates, "list-templates", false, "List all available built-in templates")
	rootCmd.Flags().StringVar(&exportTemplate, "export-template", "", "Export a built-in template to stdout (e.g. 'default')")
	rootCmd.Flags().StringVar(&validateTemplate, "validate-template", "", "Validate a custom template file for syntax errors")
	rootCmd.Flags().BoolVar(&debugTemplate, "debug-template", false, "Print template resolution info during rendering")

	// Add version flag as shortcut for "version" command
	rootCmd.Version = Version
	rootCmd.SetVersionTemplate("dock-docs {{.Version}}\n")
}

// resolveOutputPath determines the output file path for a given template format.
// If the user explicitly set a non-default output file, that is used as-is.
// Otherwise, the extension is derived from the template format (e.g. README.html, README.json).
func resolveOutputPath(currentOutput string, format string) string {
	ext := templates.OutputExtension(format)
	// If user explicitly set a custom output path (not the default), respect it
	if currentOutput != "README.md" {
		return currentOutput
	}
	// Replace the default .md extension with the format-appropriate extension
	return strings.TrimSuffix(currentOutput, filepath.Ext(currentOutput)) + ext
}

// resolveSectionOutput determines the output file path for a config section
// that uses a direct-write format (html/json).
// It derives the filename from the marker name or section index.
func resolveSectionOutput(baseOutput string, marker string, sectionIndex int, format string) string {
	ext := templates.OutputExtension(format)
	dir := filepath.Dir(baseOutput)
	base := strings.TrimSuffix(filepath.Base(baseOutput), filepath.Ext(baseOutput))

	// Use marker name as suffix if available, otherwise use section index
	var suffix string
	if marker != "" {
		suffix = "-" + marker
	} else {
		suffix = fmt.Sprintf("-section%d", sectionIndex)
	}

	return filepath.Join(dir, base+suffix+ext)
}

// resolveTemplateSel builds a TemplateSelection from the CLI flag and optional config.
// CLI --template flag takes precedence over config file settings.
func resolveTemplateSel(cfgTemplate *config.TemplateConfig) renderer.TemplateSelection {
	// CLI flag takes precedence
	if templateName != "" {
		// If it looks like a file path (contains / or .tmpl), treat as file
		if strings.Contains(templateName, "/") || strings.HasSuffix(templateName, ".tmpl") {
			return renderer.TemplateSelection{Path: templateName}
		}
		return renderer.TemplateSelection{Name: templateName}
	}

	// Fall back to config file setting
	if cfgTemplate != nil {
		sel := renderer.TemplateSelection{}
		if cfgTemplate.Path != "" {
			sel.Path = cfgTemplate.Path
		} else if cfgTemplate.Name != "" {
			sel.Name = cfgTemplate.Name
		}
		return sel
	}

	// Default: empty selection means "default" built-in
	return renderer.TemplateSelection{}
}

// describeTemplate returns a human-readable description of the template being used.
func describeTemplate(sel renderer.TemplateSelection) string {
	if sel.Path != "" {
		return fmt.Sprintf("custom file: %s", sel.Path)
	}
	if sel.Name != "" {
		return fmt.Sprintf("built-in: %s", sel.Name)
	}
	return "built-in: default"
}

// handleListTemplates prints all available built-in templates.
func handleListTemplates() error {
	builtins := templates.ListBuiltin()
	fmt.Println("Available built-in templates:")
	fmt.Println()
	for _, b := range builtins {
		fmt.Printf("  %-10s  [%s]  %s\n", b.Name, b.Format, b.Description)
	}
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  dock-docs --template <name>")
	fmt.Println("  dock-docs --export-template <name> > my-template.tmpl")
	return nil
}

// handleExportTemplate exports a built-in template to stdout.
func handleExportTemplate(name string) error {
	if !templates.IsBuiltin(name) {
		return fmt.Errorf("unknown built-in template: %s (use --list-templates to see available templates)", name)
	}

	// Export image template
	content, err := templates.ExportBuiltin(name, templates.TemplateTypeImage)
	if err != nil {
		return fmt.Errorf("failed to export template: %w", err)
	}
	fmt.Print(content)
	return nil
}

// handleValidateTemplate validates a custom template file for syntax errors.
func handleValidateTemplate(path string) error {
	loader := templates.NewLoader(false)
	if err := loader.Validate(path); err != nil {
		fmt.Fprintf(os.Stderr, "Validation failed: %v\n", err)
		return err
	}
	fmt.Printf("Template %s is valid.\n", path)
	return nil
}
