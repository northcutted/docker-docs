// Test file for the setup command (printToolStatus, runSetup, resolveToolOverrides).
//
// Globals mutated: setupCheck, setupDir, setupForce, stdout (via captureOutput),
// loadToolConfig, resolveToolOverrides.
// All tests use defer resetFlags()() for cleanup.
package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/northcutted/dock-docs/pkg/config"
	"github.com/northcutted/dock-docs/pkg/installer"
)

func TestPrintToolStatus(t *testing.T) {
	output := captureOutput(func() {
		if err := printToolStatus(t.TempDir()); err != nil {
			t.Fatalf("printToolStatus() error: %v", err)
		}
	})

	if !strings.Contains(output, "Tool Status:") {
		t.Error("expected 'Tool Status:' header in output")
	}
	// Should mention all three tools
	for _, tool := range []string{"syft", "grype", "dive"} {
		if !strings.Contains(output, tool) {
			t.Errorf("expected %q in tool status output", tool)
		}
	}
}

func TestRunSetup_CheckOnly(t *testing.T) {
	defer resetFlags()()

	setupCheck = true
	setupDir = t.TempDir()
	setupForce = false

	output := captureOutput(func() {
		if err := runSetup(setupCmd, nil); err != nil {
			t.Fatalf("runSetup(--check) error: %v", err)
		}
	})

	if !strings.Contains(output, "Tool Status:") {
		t.Error("expected tool status output from --check")
	}
}

func TestRunSetup_InstallAllPresent(t *testing.T) {
	defer resetFlags()()

	tmpDir := t.TempDir()

	// Pre-create fake tool binaries so InstallAll skips downloads
	for _, tool := range installer.Tools {
		fakePath := filepath.Join(tmpDir, tool.Name)
		if err := os.WriteFile(fakePath, []byte("fake"), 0755); err != nil {
			t.Fatalf("failed to write fake %s: %v", tool.Name, err)
		}
	}

	setupCheck = false
	setupDir = tmpDir
	setupForce = false

	// Stub loadToolConfig to return nil (no config file)
	loadToolConfig = func() map[string]config.ToolConfig { return nil }

	_, logOut := captureAll(func() {
		if err := runSetup(setupCmd, nil); err != nil {
			t.Fatalf("runSetup() error: %v", err)
		}
	})

	// InstallAll should report all tools already installed
	if !strings.Contains(logOut, "all tools are already installed") {
		t.Errorf("expected 'all tools are already installed' in log, got:\n%s", logOut)
	}
}

func TestRunSetup_DefaultDir(t *testing.T) {
	defer resetFlags()()

	// Test that runSetup uses DefaultInstallDir when setupDir is empty
	// and --check mode so it doesn't try to install anything
	setupCheck = true
	setupDir = ""
	setupForce = false

	output := captureOutput(func() {
		if err := runSetup(setupCmd, nil); err != nil {
			t.Fatalf("runSetup() error: %v", err)
		}
	})

	if !strings.Contains(output, "Tool Status:") {
		t.Error("expected tool status output")
	}
}

func TestResolveToolOverrides_ConfigOnly(t *testing.T) {
	defer resetFlags()()

	cfgTools := map[string]config.ToolConfig{
		"syft":  {Version: "v1.21.0", URL: "https://proxy.corp.com/{name}_{version}.tar.gz"},
		"grype": {Version: "v0.87.0"},
	}

	overrides := resolveToolOverrides(cfgTools)
	if overrides == nil {
		t.Fatal("expected non-nil overrides")
	}

	syft, ok := overrides["syft"]
	if !ok {
		t.Fatal("expected 'syft' in overrides")
	}
	if syft.Version != "v1.21.0" {
		t.Errorf("syft.Version = %q, want %q", syft.Version, "v1.21.0")
	}
	if syft.URL != "https://proxy.corp.com/{name}_{version}.tar.gz" {
		t.Errorf("syft.URL = %q, want proxy URL", syft.URL)
	}

	grype, ok := overrides["grype"]
	if !ok {
		t.Fatal("expected 'grype' in overrides")
	}
	if grype.Version != "v0.87.0" {
		t.Errorf("grype.Version = %q, want %q", grype.Version, "v0.87.0")
	}
	if grype.URL != "" {
		t.Errorf("grype.URL = %q, want empty", grype.URL)
	}

	// dive should not appear (no config for it)
	if _, ok := overrides["dive"]; ok {
		t.Error("expected 'dive' NOT in overrides when not configured")
	}
}

func TestResolveToolOverrides_EnvVarsOnly(t *testing.T) {
	defer resetFlags()()

	// Set env vars for syft
	t.Setenv("DOCK_DOCS_SYFT_VERSION", "v2.0.0")
	t.Setenv("DOCK_DOCS_SYFT_URL", "https://env.example.com/syft.tar.gz")

	overrides := resolveToolOverrides(nil)
	if overrides == nil {
		t.Fatal("expected non-nil overrides from env vars")
	}

	syft, ok := overrides["syft"]
	if !ok {
		t.Fatal("expected 'syft' in overrides")
	}
	if syft.Version != "v2.0.0" {
		t.Errorf("syft.Version = %q, want %q", syft.Version, "v2.0.0")
	}
	if syft.URL != "https://env.example.com/syft.tar.gz" {
		t.Errorf("syft.URL = %q, want env URL", syft.URL)
	}
}

func TestResolveToolOverrides_EnvPrecedence(t *testing.T) {
	defer resetFlags()()

	// Config sets syft version and URL
	cfgTools := map[string]config.ToolConfig{
		"syft": {Version: "v1.0.0", URL: "https://config.example.com/{name}.tar.gz"},
	}

	// Env vars override both
	t.Setenv("DOCK_DOCS_SYFT_VERSION", "v9.9.9")
	t.Setenv("DOCK_DOCS_SYFT_URL", "https://env.example.com/{name}.tar.gz")

	overrides := resolveToolOverrides(cfgTools)
	if overrides == nil {
		t.Fatal("expected non-nil overrides")
	}

	syft := overrides["syft"]
	if syft.Version != "v9.9.9" {
		t.Errorf("syft.Version = %q, want env value %q", syft.Version, "v9.9.9")
	}
	if syft.URL != "https://env.example.com/{name}.tar.gz" {
		t.Errorf("syft.URL = %q, want env URL", syft.URL)
	}
}

func TestResolveToolOverrides_EnvPartialOverride(t *testing.T) {
	defer resetFlags()()

	// Config sets version and URL for syft
	cfgTools := map[string]config.ToolConfig{
		"syft": {Version: "v1.0.0", URL: "https://config.example.com/{name}.tar.gz"},
	}

	// Env overrides only the version (URL stays from config)
	t.Setenv("DOCK_DOCS_SYFT_VERSION", "v5.0.0")

	overrides := resolveToolOverrides(cfgTools)
	if overrides == nil {
		t.Fatal("expected non-nil overrides")
	}

	syft := overrides["syft"]
	if syft.Version != "v5.0.0" {
		t.Errorf("syft.Version = %q, want env value %q", syft.Version, "v5.0.0")
	}
	if syft.URL != "https://config.example.com/{name}.tar.gz" {
		t.Errorf("syft.URL = %q, want config URL (not overridden by env)", syft.URL)
	}
}

func TestResolveToolOverrides_NoConfig(t *testing.T) {
	defer resetFlags()()

	overrides := resolveToolOverrides(nil)
	if overrides != nil {
		t.Errorf("expected nil overrides when no config and no env vars, got %v", overrides)
	}
}

func TestRunSetup_URLWithoutVersionError(t *testing.T) {
	defer resetFlags()()

	tmpDir := t.TempDir()
	setupCheck = false
	setupDir = tmpDir
	setupForce = false

	// Stub loadToolConfig to return nil
	loadToolConfig = func() map[string]config.ToolConfig { return nil }

	// Stub resolveToolOverrides to return URL without version
	resolveToolOverrides = func(_ map[string]config.ToolConfig) map[string]installer.InstallOptions {
		return map[string]installer.InstallOptions{
			"syft": {URL: "https://proxy.corp.com/syft.tar.gz"},
		}
	}

	_, _ = captureAll(func() {
		err := runSetup(setupCmd, nil)
		if err == nil {
			t.Fatal("expected error for URL without version")
		}
		if !strings.Contains(err.Error(), "no version") {
			t.Errorf("error = %q, want substring 'no version'", err.Error())
		}
	})
}

func TestRunSetup_WithOverridesLogged(t *testing.T) {
	defer resetFlags()()

	tmpDir := t.TempDir()

	// Pre-create fake tool binaries so InstallAll skips downloads
	for _, tool := range installer.Tools {
		fakePath := filepath.Join(tmpDir, tool.Name)
		if err := os.WriteFile(fakePath, []byte("fake"), 0755); err != nil {
			t.Fatalf("failed to write fake %s: %v", tool.Name, err)
		}
	}

	setupCheck = false
	setupDir = tmpDir
	setupForce = false

	// Stub loadToolConfig to return a config with syft pinned
	loadToolConfig = func() map[string]config.ToolConfig {
		return map[string]config.ToolConfig{
			"syft": {Version: "v1.21.0"},
		}
	}

	_, logOut := captureAll(func() {
		if err := runSetup(setupCmd, nil); err != nil {
			t.Fatalf("runSetup() error: %v", err)
		}
	})

	// Should log the pinned version
	if !strings.Contains(logOut, "pinned version") {
		t.Errorf("expected 'pinned version' in log, got:\n%s", logOut)
	}
	if !strings.Contains(logOut, "v1.21.0") {
		t.Errorf("expected 'v1.21.0' in log, got:\n%s", logOut)
	}
}
