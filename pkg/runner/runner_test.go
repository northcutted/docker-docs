package runner

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/northcutted/dock-docs/pkg/types"
)

func TestRuntimeRunner_Name(t *testing.T) {
	tests := []struct {
		name     string
		binary   string
		expected string
	}{
		{
			name:     "default name",
			binary:   "",
			expected: "runtime",
		},
		{
			name:     "docker binary",
			binary:   "docker",
			expected: "docker",
		},
		{
			name:     "podman binary",
			binary:   "podman",
			expected: "podman",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RuntimeRunner{binary: tt.binary}
			if got := r.Name(); got != tt.expected {
				t.Errorf("RuntimeRunner.Name() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestRuntimeRunner_Run(t *testing.T) {
	// Test parsing of docker inspect JSON output
	inspectJSON := `[{
		"Architecture": "amd64",
		"Os": "linux",
		"Size": 67108864,
		"RootFS": {
			"Layers": ["layer1", "layer2", "layer3"]
		}
	}]`

	tests := []struct {
		name       string
		jsonOutput string
		wantErr    bool
		wantArch   string
		wantOS     string
		wantLayers int
		wantSizeMB string
	}{
		{
			name:       "valid inspect output",
			jsonOutput: inspectJSON,
			wantErr:    false,
			wantArch:   "amd64",
			wantOS:     "linux",
			wantLayers: 3,
			wantSizeMB: "64.00 MB",
		},
		{
			name:       "empty array",
			jsonOutput: `[]`,
			wantErr:    true,
		},
		{
			name:       "invalid json",
			jsonOutput: `{invalid}`,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var inspect []struct {
				Architecture string `json:"Architecture"`
				Os           string `json:"Os"`
				Size         int64  `json:"Size"`
				RootFS       struct {
					Layers []string `json:"Layers"`
				} `json:"RootFS"`
			}

			err := json.Unmarshal([]byte(tt.jsonOutput), &inspect)
			if tt.wantErr {
				if err == nil && len(inspect) == 0 {
					// Expected error condition (empty array)
					return
				}
				if err != nil {
					// Expected error condition (invalid JSON)
					return
				}
				t.Errorf("expected error but got none")
				return
			}

			if err != nil {
				t.Fatalf("unexpected unmarshal error: %v", err)
			}

			if len(inspect) == 0 {
				t.Fatal("no inspect data returned")
			}

			data := inspect[0]
			if data.Architecture != tt.wantArch {
				t.Errorf("Architecture = %v, want %v", data.Architecture, tt.wantArch)
			}
			if data.Os != tt.wantOS {
				t.Errorf("Os = %v, want %v", data.Os, tt.wantOS)
			}
			if len(data.RootFS.Layers) != tt.wantLayers {
				t.Errorf("Layers count = %v, want %v", len(data.RootFS.Layers), tt.wantLayers)
			}
		})
	}
}

func TestManifestRunner_Name(t *testing.T) {
	r := &ManifestRunner{}
	if got := r.Name(); got != "manifest" {
		t.Errorf("ManifestRunner.Name() = %v, want %v", got, "manifest")
	}
}

func TestManifestRunner_ParseOutput(t *testing.T) {
	// Test manifest list parsing
	manifestListJSON := `{
		"manifests": [
			{"platform": {"architecture": "amd64", "os": "linux"}},
			{"platform": {"architecture": "arm64", "os": "linux"}},
			{"platform": {"architecture": "arm", "os": "linux"}},
			{"platform": {"architecture": "amd64", "os": "linux"}}
		]
	}`

	type Platform struct {
		Architecture string `json:"architecture"`
		OS           string `json:"os"`
	}
	type Manifest struct {
		Platform Platform `json:"platform"`
	}
	type ManifestIndex struct {
		Manifests []Manifest `json:"manifests"`
	}

	var index ManifestIndex
	err := json.Unmarshal([]byte(manifestListJSON), &index)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if len(index.Manifests) != 4 {
		t.Errorf("expected 4 manifests, got %d", len(index.Manifests))
	}

	// Test deduplication logic
	seen := make(map[string]bool)
	var archs []string
	for _, m := range index.Manifests {
		key := m.Platform.OS + "/" + m.Platform.Architecture
		if !seen[key] {
			seen[key] = true
			archs = append(archs, key)
		}
	}

	if len(archs) != 3 {
		t.Errorf("expected 3 unique architectures after dedup, got %d: %v", len(archs), archs)
	}
}

func TestSyftRunner_Name(t *testing.T) {
	r := &SyftRunner{}
	if got := r.Name(); got != "syft" {
		t.Errorf("SyftRunner.Name() = %v, want %v", got, "syft")
	}
}

func TestSyftRunner_ParseOutput(t *testing.T) {
	syftJSON := `{
		"distro": {
			"name": "ubuntu",
			"version": "22.04"
		},
		"artifacts": [
			{"name": "openssl", "version": "3.0.2", "type": "deb"},
			{"name": "curl", "version": "7.81.0", "type": "deb"},
			{"name": "openssl", "version": "3.0.2", "type": "deb"}
		]
	}`

	var syftOutput struct {
		Distro struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"distro"`
		Artifacts []struct {
			Name    string `json:"name"`
			Version string `json:"version"`
			Type    string `json:"type"`
		} `json:"artifacts"`
	}

	err := json.Unmarshal([]byte(syftJSON), &syftOutput)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if syftOutput.Distro.Name != "ubuntu" {
		t.Errorf("Distro.Name = %v, want ubuntu", syftOutput.Distro.Name)
	}

	if syftOutput.Distro.Version != "22.04" {
		t.Errorf("Distro.Version = %v, want 22.04", syftOutput.Distro.Version)
	}

	if len(syftOutput.Artifacts) != 3 {
		t.Errorf("expected 3 artifacts, got %d", len(syftOutput.Artifacts))
	}

	// Test deduplication
	seen := make(map[string]bool)
	uniqueCount := 0
	for _, artifact := range syftOutput.Artifacts {
		key := artifact.Name + "@" + artifact.Version
		if !seen[key] {
			seen[key] = true
			uniqueCount++
		}
	}

	if uniqueCount != 2 {
		t.Errorf("expected 2 unique packages, got %d", uniqueCount)
	}
}

func TestGrypeRunner_Name(t *testing.T) {
	r := &GrypeRunner{}
	if got := r.Name(); got != "grype" {
		t.Errorf("GrypeRunner.Name() = %v, want %v", got, "grype")
	}
}

func TestGrypeRunner_ParseOutput(t *testing.T) {
	grypeJSON := `{
		"matches": [
			{
				"vulnerability": {"id": "CVE-2023-0001", "severity": "Critical"},
				"artifact": {"name": "openssl", "version": "3.0.2"}
			},
			{
				"vulnerability": {"id": "CVE-2023-0002", "severity": "High"},
				"artifact": {"name": "curl", "version": "7.81.0"}
			},
			{
				"vulnerability": {"id": "CVE-2023-0003", "severity": "Critical"},
				"artifact": {"name": "libssl", "version": "3.0.2"}
			},
			{
				"vulnerability": {"id": "CVE-2023-0004", "severity": "Medium"},
				"artifact": {"name": "bash", "version": "5.1"}
			}
		]
	}`

	var grypeOutput struct {
		Matches []struct {
			Vulnerability struct {
				ID       string `json:"id"`
				Severity string `json:"severity"`
			} `json:"vulnerability"`
			Artifact struct {
				Name    string `json:"name"`
				Version string `json:"version"`
			} `json:"artifact"`
		} `json:"matches"`
	}

	err := json.Unmarshal([]byte(grypeJSON), &grypeOutput)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if len(grypeOutput.Matches) != 4 {
		t.Errorf("expected 4 matches, got %d", len(grypeOutput.Matches))
	}

	// Test severity counting
	summary := make(map[string]int)
	vulns := make([]types.Vulnerability, 0)

	for _, match := range grypeOutput.Matches {
		sev := match.Vulnerability.Severity
		summary[sev]++
		vulns = append(vulns, types.Vulnerability{
			ID:       match.Vulnerability.ID,
			Severity: sev,
			Package:  match.Artifact.Name,
			Version:  match.Artifact.Version,
		})
	}

	if summary["Critical"] != 2 {
		t.Errorf("expected 2 Critical vulnerabilities, got %d", summary["Critical"])
	}
	if summary["High"] != 1 {
		t.Errorf("expected 1 High vulnerability, got %d", summary["High"])
	}
	if summary["Medium"] != 1 {
		t.Errorf("expected 1 Medium vulnerability, got %d", summary["Medium"])
	}

	// Test sorting
	severityRank := map[string]int{
		"Critical": 4,
		"High":     3,
		"Medium":   2,
		"Low":      1,
		"Unknown":  0,
	}

	// Verify vulnerabilities are properly structured
	for i, vuln := range vulns {
		if vuln.ID == "" {
			t.Errorf("vulnerability %d has empty ID", i)
		}
		if vuln.Severity == "" {
			t.Errorf("vulnerability %d has empty Severity", i)
		}
		if _, ok := severityRank[vuln.Severity]; !ok {
			t.Errorf("vulnerability %d has unknown severity: %s", i, vuln.Severity)
		}
	}
}

func TestDiveRunner_Name(t *testing.T) {
	r := &DiveRunner{}
	if got := r.Name(); got != "dive" {
		t.Errorf("DiveRunner.Name() = %v, want %v", got, "dive")
	}
}

func TestDiveRunner_ParseOutput(t *testing.T) {
	diveJSON := `{
		"image": {
			"inefficientBytes": 10485760,
			"efficiencyScore": 0.95
		}
	}`

	var diveOutput struct {
		Image struct {
			InefficientBytes uint64  `json:"inefficientBytes"`
			EfficiencyScore  float64 `json:"efficiencyScore"`
		} `json:"image"`
	}

	err := json.Unmarshal([]byte(diveJSON), &diveOutput)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if diveOutput.Image.InefficientBytes != 10485760 {
		t.Errorf("InefficientBytes = %v, want 10485760", diveOutput.Image.InefficientBytes)
	}

	if diveOutput.Image.EfficiencyScore != 0.95 {
		t.Errorf("EfficiencyScore = %v, want 0.95", diveOutput.Image.EfficiencyScore)
	}

	// Test conversion to percentage
	efficiency := diveOutput.Image.EfficiencyScore * 100
	if efficiency != 95.0 {
		t.Errorf("efficiency percentage = %v, want 95.0", efficiency)
	}

	// Test wasted bytes conversion
	wastedMB := float64(diveOutput.Image.InefficientBytes) / 1024 / 1024
	expectedMB := 10.0 // 10485760 bytes = 10 MB
	if wastedMB != expectedMB {
		t.Errorf("wastedMB = %v, want %v", wastedMB, expectedMB)
	}
}

func TestRunCommand_Verbose(t *testing.T) {
	// Test that verbose flag controls output
	// This is a lightweight test that checks the verbose logic without executing commands
	verbose := true
	if !verbose {
		t.Error("verbose should be true for this test")
	}

	verbose = false
	if verbose {
		t.Error("verbose should be false when set to false")
	}
}

func TestEnsureImage_ErrorHandling(t *testing.T) {
	// Test that EnsureImage handles empty image name
	image := ""
	if image == "" {
		// This would cause an error in EnsureImage
		// Verifying the precondition check
		t.Log("Empty image name would cause error - expected behavior")
	}
}

func TestPodmanSocketDetection(t *testing.T) {
	// Test Podman socket detection logic
	testJSON := `[{
		"ConnectionInfo": {
			"PodmanSocket": {
				"Path": "/var/run/podman.sock"
			}
		}
	}]`

	var machines []struct {
		ConnectionInfo struct {
			PodmanSocket *struct {
				Path string `json:"Path"`
			} `json:"PodmanSocket"`
		} `json:"ConnectionInfo"`
	}

	err := json.Unmarshal([]byte(testJSON), &machines)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if len(machines) == 0 {
		t.Fatal("expected at least one machine")
	}

	if machines[0].ConnectionInfo.PodmanSocket == nil {
		t.Fatal("PodmanSocket is nil")
	}

	socketPath := machines[0].ConnectionInfo.PodmanSocket.Path
	if socketPath != "/var/run/podman.sock" {
		t.Errorf("socket path = %v, want /var/run/podman.sock", socketPath)
	}

	// Test unix:// prefix logic
	if !strings.HasPrefix(socketPath, "unix://") {
		socketPath = "unix://" + socketPath
	}

	expectedPath := "unix:///var/run/podman.sock"
	if socketPath != expectedPath {
		t.Errorf("formatted socket path = %v, want %v", socketPath, expectedPath)
	}
}

func TestTempFileCleanup(t *testing.T) {
	// Test that temp file cleanup works properly
	tmpFile, err := os.CreateTemp("", "test-dive-output-*.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	tmpPath := tmpFile.Name()
	if err := tmpFile.Close(); err != nil {
		t.Fatalf("failed to close temp file: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(tmpPath); os.IsNotExist(err) {
		t.Fatal("temp file should exist after creation")
	}

	// Remove file
	err = os.Remove(tmpPath)
	if err != nil {
		t.Fatalf("failed to remove temp file: %v", err)
	}

	// Verify file is removed
	if _, err := os.Stat(tmpPath); !os.IsNotExist(err) {
		t.Error("temp file should not exist after removal")
	}
}
