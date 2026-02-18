package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/northcutted/dock-docs/pkg/config"
	"github.com/northcutted/dock-docs/pkg/installer"
)

var (
	setupDir   string
	setupForce bool
	setupCheck bool
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Install external tool dependencies (syft, grype, dive)",
	Long: `Downloads and installs the external tools required for image analysis.

Tools installed:
  - syft   (SBOM generation)       from anchore/syft
  - grype  (vulnerability scanning) from anchore/grype
  - dive   (layer efficiency)       from wagoodman/dive

Binaries are downloaded from upstream GitHub Releases and placed in
~/.dock-docs/bin/ by default. System-installed tools (found in PATH)
are always preferred; the install directory serves as a fallback.

Enterprise users can override download URLs and pin versions via
the dock-docs.yaml config file or environment variables:

  Config file (tools section):
    tools:
      syft:
        version: "v1.21.0"
        url: "https://proxy.corp.com/{name}_{version}_{os}_{arch}.tar.gz"

  Environment variables (override config values):
    DOCK_DOCS_SYFT_VERSION, DOCK_DOCS_SYFT_URL
    DOCK_DOCS_GRYPE_VERSION, DOCK_DOCS_GRYPE_URL
    DOCK_DOCS_DIVE_VERSION, DOCK_DOCS_DIVE_URL

URL placeholders: {name}, {version} (no "v"), {tag} (with "v"), {os}, {arch}`,
	Example: `  # Install all missing tools
  dock-docs setup

  # Reinstall even if already present
  dock-docs setup --force

  # Install to a custom directory
  dock-docs setup --dir /usr/local/bin

  # Check status only (no downloads)
  dock-docs setup --check`,
	RunE: runSetup,
}

func init() {
	setupCmd.Flags().StringVar(&setupDir, "dir", "", "Install directory (default: ~/.dock-docs/bin)")
	setupCmd.Flags().BoolVar(&setupForce, "force", false, "Reinstall tools even if already present")
	setupCmd.Flags().BoolVar(&setupCheck, "check", false, "Show tool status without installing")

	rootCmd.AddCommand(setupCmd)
}

// resolveToolOverrides merges tool configuration from the config file and
// environment variables into a map of InstallOptions keyed by tool name.
// Environment variables take precedence over config file values.
var resolveToolOverrides = func(cfgTools map[string]config.ToolConfig) map[string]installer.InstallOptions {
	overrides := make(map[string]installer.InstallOptions)

	for _, t := range installer.Tools {
		name := t.Name
		var opts installer.InstallOptions

		// Layer 1: config file values
		if tc, ok := cfgTools[name]; ok {
			opts.Version = tc.Version
			opts.URL = tc.URL
		}

		// Layer 2: environment variables (override config)
		envPrefix := "DOCK_DOCS_" + strings.ToUpper(name) + "_"
		if v := os.Getenv(envPrefix + "VERSION"); v != "" {
			opts.Version = v
		}
		if v := os.Getenv(envPrefix + "URL"); v != "" {
			opts.URL = v
		}

		if opts.Version != "" || opts.URL != "" {
			overrides[name] = opts
		}
	}

	if len(overrides) > 0 {
		return overrides
	}
	return nil
}

// loadToolConfig attempts to load tool overrides from the dock-docs config
// file. Returns nil (no overrides) if the file doesn't exist or has no tools
// section. Only returns an error for malformed YAML.
var loadToolConfig = func() map[string]config.ToolConfig {
	// Try the default config file location
	const defaultConfig = "dock-docs.yaml"
	cfg, err := config.Load(defaultConfig)
	if err != nil {
		// Config file missing or invalid — not an error for setup
		return nil
	}
	return cfg.Tools
}

func runSetup(cmd *cobra.Command, args []string) error {
	// Resolve install directory
	installDir := setupDir
	if installDir == "" {
		dir, err := installer.DefaultInstallDir()
		if err != nil {
			return fmt.Errorf("failed to determine install directory: %w", err)
		}
		installDir = dir
	}

	// --check: show status and exit
	if setupCheck {
		return printToolStatus(installDir)
	}

	// Create directory if needed
	if err := os.MkdirAll(installDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", installDir, err)
	}

	slog.Info("install directory", "path", installDir)

	// Build tool overrides from config file + env vars
	cfgTools := loadToolConfig()
	overrides := resolveToolOverrides(cfgTools)

	// Validate: URL requires Version
	for name, opts := range overrides {
		if opts.URL != "" && opts.Version == "" {
			return fmt.Errorf("tool %q has a custom URL but no version; set version via config or DOCK_DOCS_%s_VERSION",
				name, strings.ToUpper(name))
		}
	}

	for name, opts := range overrides {
		if opts.URL != "" {
			slog.Info("using custom download URL", "tool", name, "url", opts.URL)
		}
		if opts.Version != "" {
			slog.Info("using pinned version", "tool", name, "version", opts.Version)
		}
	}

	// Install missing (or all if --force)
	if err := installer.InstallAll(installDir, setupForce, overrides); err != nil {
		return err
	}

	// Print final status
	fmt.Fprintln(stdout)
	return printToolStatus(installDir)
}

func printToolStatus(installDir string) error {
	statuses := installer.Status(installDir)

	fmt.Fprintln(stdout, "Tool Status:")
	for _, s := range statuses {
		if s.Installed {
			fmt.Fprintf(stdout, "  [OK]      %-6s  (%s: %s)\n", s.Name, s.Source, s.Path)
		} else {
			fmt.Fprintf(stdout, "  [MISSING] %-6s\n", s.Name)
		}
	}
	return nil
}
