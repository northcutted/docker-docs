package runner

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
	"time"
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

func TestParseRuntimeInspect(t *testing.T) {
	tests := []struct {
		name       string
		jsonOutput string
		image      string
		binary     string
		wantErr    bool
		wantArch   string
		wantOS     string
		wantLayers int
		wantSizeMB string
	}{
		{
			name: "valid inspect output",
			jsonOutput: `[{
				"Architecture": "amd64",
				"Os": "linux",
				"Size": 67108864,
				"RootFS": {
					"Layers": ["layer1", "layer2", "layer3"]
				}
			}]`,
			image:      "nginx:latest",
			binary:     "docker",
			wantErr:    false,
			wantArch:   "amd64",
			wantOS:     "linux",
			wantLayers: 3,
			wantSizeMB: "64.00 MB",
		},
		{
			name: "arm64 image",
			jsonOutput: `[{
				"Architecture": "arm64",
				"Os": "linux",
				"Size": 52428800,
				"RootFS": {
					"Layers": ["sha256:abc", "sha256:def"]
				}
			}]`,
			image:      "alpine:3.18",
			binary:     "podman",
			wantErr:    false,
			wantArch:   "arm64",
			wantOS:     "linux",
			wantLayers: 2,
			wantSizeMB: "50.00 MB",
		},
		{
			name:       "empty array",
			jsonOutput: `[]`,
			image:      "test:latest",
			binary:     "docker",
			wantErr:    true,
		},
		{
			name:       "invalid json",
			jsonOutput: `{invalid}`,
			image:      "test:latest",
			binary:     "docker",
			wantErr:    true,
		},
		{
			name:       "zero-size image with no layers",
			jsonOutput: `[{"Architecture": "amd64", "Os": "linux", "Size": 0, "RootFS": {"Layers": []}}]`,
			image:      "scratch",
			binary:     "docker",
			wantErr:    false,
			wantArch:   "amd64",
			wantOS:     "linux",
			wantLayers: 0,
			wantSizeMB: "0.00 MB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stats, err := parseRuntimeInspect([]byte(tt.jsonOutput), tt.image, tt.binary)
			if tt.wantErr {
				if err == nil {
					t.Errorf("parseRuntimeInspect() expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("parseRuntimeInspect() unexpected error: %v", err)
			}
			if stats.Architecture != tt.wantArch {
				t.Errorf("Architecture = %v, want %v", stats.Architecture, tt.wantArch)
			}
			if stats.OS != tt.wantOS {
				t.Errorf("OS = %v, want %v", stats.OS, tt.wantOS)
			}
			if stats.TotalLayers != tt.wantLayers {
				t.Errorf("TotalLayers = %v, want %v", stats.TotalLayers, tt.wantLayers)
			}
			if stats.SizeMB != tt.wantSizeMB {
				t.Errorf("SizeMB = %v, want %v", stats.SizeMB, tt.wantSizeMB)
			}
			if stats.ImageTag != tt.image {
				t.Errorf("ImageTag = %v, want %v", stats.ImageTag, tt.image)
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

func TestParseManifestInspect(t *testing.T) {
	tests := []struct {
		name      string
		json      string
		image     string
		wantArchs []string
		wantErr   bool
	}{
		{
			name: "manifest list with duplicates",
			json: `{
				"manifests": [
					{"platform": {"architecture": "amd64", "os": "linux"}},
					{"platform": {"architecture": "arm64", "os": "linux"}},
					{"platform": {"architecture": "arm", "os": "linux"}},
					{"platform": {"architecture": "amd64", "os": "linux"}}
				]
			}`,
			image:     "nginx:latest",
			wantArchs: []string{"linux/amd64", "linux/arm", "linux/arm64"},
		},
		{
			name: "single platform",
			json: `{
				"manifests": [
					{"platform": {"architecture": "amd64", "os": "linux"}}
				]
			}`,
			image:     "myapp:v1",
			wantArchs: []string{"linux/amd64"},
		},
		{
			name: "multi-os manifest",
			json: `{
				"manifests": [
					{"platform": {"architecture": "amd64", "os": "linux"}},
					{"platform": {"architecture": "amd64", "os": "windows"}}
				]
			}`,
			image:     "dotnet:latest",
			wantArchs: []string{"linux/amd64", "windows/amd64"},
		},
		{
			name:      "empty manifests list",
			json:      `{"manifests": []}`,
			image:     "test:latest",
			wantArchs: nil, // falls through to empty stats
		},
		{
			name:      "not a manifest list (single image V2 schema)",
			json:      `{"schemaVersion": 2, "mediaType": "application/vnd.docker.distribution.manifest.v2+json"}`,
			image:     "custom:latest",
			wantArchs: nil,
		},
		{
			name:      "invalid json returns empty stats",
			json:      `{totally broken`,
			image:     "bad:latest",
			wantArchs: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stats, err := parseManifestInspect([]byte(tt.json), tt.image)
			if tt.wantErr {
				if err == nil {
					t.Errorf("parseManifestInspect() expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("parseManifestInspect() unexpected error: %v", err)
			}
			if stats.ImageTag != tt.image {
				t.Errorf("ImageTag = %v, want %v", stats.ImageTag, tt.image)
			}
			if tt.wantArchs == nil {
				if len(stats.SupportedArchitectures) != 0 {
					t.Errorf("SupportedArchitectures = %v, want empty", stats.SupportedArchitectures)
				}
				return
			}
			if len(stats.SupportedArchitectures) != len(tt.wantArchs) {
				t.Fatalf("SupportedArchitectures count = %d, want %d: %v",
					len(stats.SupportedArchitectures), len(tt.wantArchs), stats.SupportedArchitectures)
			}
			for i, arch := range stats.SupportedArchitectures {
				if arch != tt.wantArchs[i] {
					t.Errorf("SupportedArchitectures[%d] = %v, want %v", i, arch, tt.wantArchs[i])
				}
			}
		})
	}
}

func TestSyftRunner_Name(t *testing.T) {
	r := &SyftRunner{}
	if got := r.Name(); got != "syft" {
		t.Errorf("SyftRunner.Name() = %v, want %v", got, "syft")
	}
}

func TestParseSyftOutput(t *testing.T) {
	tests := []struct {
		name         string
		json         string
		wantErr      bool
		wantDistro   string
		wantPkgCount int
		wantFirstPkg string
		wantLastPkg  string
	}{
		{
			name: "standard output with duplicates",
			json: `{
				"distro": {"name": "ubuntu", "version": "22.04"},
				"artifacts": [
					{"name": "openssl", "version": "3.0.2", "type": "deb"},
					{"name": "curl", "version": "7.81.0", "type": "deb"},
					{"name": "openssl", "version": "3.0.2", "type": "deb"}
				]
			}`,
			wantDistro:   "ubuntu 22.04",
			wantPkgCount: 2,
			wantFirstPkg: "curl",
			wantLastPkg:  "openssl",
		},
		{
			name: "distro without version",
			json: `{
				"distro": {"name": "alpine", "version": ""},
				"artifacts": [
					{"name": "musl", "version": "1.2.4", "type": "apk"}
				]
			}`,
			wantDistro:   "alpine",
			wantPkgCount: 1,
			wantFirstPkg: "musl",
			wantLastPkg:  "musl",
		},
		{
			name: "no distro info",
			json: `{
				"distro": {"name": "", "version": ""},
				"artifacts": [
					{"name": "zlib", "version": "1.0", "type": "binary"}
				]
			}`,
			wantDistro:   "",
			wantPkgCount: 1,
			wantFirstPkg: "zlib",
			wantLastPkg:  "zlib",
		},
		{
			name:         "empty artifacts",
			json:         `{"distro": {"name": "", "version": ""}, "artifacts": []}`,
			wantDistro:   "",
			wantPkgCount: 0,
		},
		{
			name:    "invalid json",
			json:    `not json`,
			wantErr: true,
		},
		{
			name: "packages sorted alphabetically",
			json: `{
				"distro": {"name": "debian", "version": "12"},
				"artifacts": [
					{"name": "zlib", "version": "1.2", "type": "deb"},
					{"name": "apt", "version": "2.6", "type": "deb"},
					{"name": "bash", "version": "5.2", "type": "deb"}
				]
			}`,
			wantDistro:   "debian 12",
			wantPkgCount: 3,
			wantFirstPkg: "apt",
			wantLastPkg:  "zlib",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stats, err := parseSyftOutput([]byte(tt.json))
			if tt.wantErr {
				if err == nil {
					t.Errorf("parseSyftOutput() expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("parseSyftOutput() unexpected error: %v", err)
			}
			if stats.OSDistro != tt.wantDistro {
				t.Errorf("OSDistro = %q, want %q", stats.OSDistro, tt.wantDistro)
			}
			if stats.TotalPackages != tt.wantPkgCount {
				t.Errorf("TotalPackages = %d, want %d", stats.TotalPackages, tt.wantPkgCount)
			}
			if len(stats.Packages) != tt.wantPkgCount {
				t.Errorf("len(Packages) = %d, want %d", len(stats.Packages), tt.wantPkgCount)
			}
			if tt.wantPkgCount > 0 {
				if stats.Packages[0].Name != tt.wantFirstPkg {
					t.Errorf("first package = %q, want %q", stats.Packages[0].Name, tt.wantFirstPkg)
				}
				if stats.Packages[len(stats.Packages)-1].Name != tt.wantLastPkg {
					t.Errorf("last package = %q, want %q", stats.Packages[len(stats.Packages)-1].Name, tt.wantLastPkg)
				}
			}
		})
	}
}

func TestGrypeRunner_Name(t *testing.T) {
	r := &GrypeRunner{}
	if got := r.Name(); got != "grype" {
		t.Errorf("GrypeRunner.Name() = %v, want %v", got, "grype")
	}
}

func TestParseGrypeOutput(t *testing.T) {
	tests := []struct {
		name             string
		json             string
		verbose          bool
		wantErr          bool
		wantCritical     int
		wantHigh         int
		wantMedium       int
		wantLow          int
		wantTotalVulns   int
		wantFirstVulnSev string
		wantTimeParsed   bool
		wantTimestamp    string
	}{
		{
			name: "standard output with multiple severities",
			json: `{
				"descriptor": {"timestamp": "2024-02-15T14:30:00Z"},
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
			}`,
			wantCritical:     2,
			wantHigh:         1,
			wantMedium:       1,
			wantTotalVulns:   4,
			wantFirstVulnSev: "Critical",
			wantTimeParsed:   true,
			wantTimestamp:    "2024-02-15T14:30:00Z",
		},
		{
			name: "no vulnerabilities",
			json: `{
				"descriptor": {"timestamp": "2024-01-01T00:00:00Z"},
				"matches": []
			}`,
			wantTotalVulns: 0,
			wantTimeParsed: true,
			wantTimestamp:  "2024-01-01T00:00:00Z",
		},
		{
			name: "missing timestamp defaults to now",
			json: `{
				"descriptor": {},
				"matches": [
					{
						"vulnerability": {"id": "CVE-2024-0001", "severity": "Low"},
						"artifact": {"name": "libc", "version": "2.31"}
					}
				]
			}`,
			wantLow:          1,
			wantTotalVulns:   1,
			wantFirstVulnSev: "Low",
			wantTimeParsed:   false,
		},
		{
			name: "sort order: Critical before High before Medium before Low",
			json: `{
				"descriptor": {"timestamp": "2024-06-01T12:00:00Z"},
				"matches": [
					{
						"vulnerability": {"id": "CVE-LOW", "severity": "Low"},
						"artifact": {"name": "pkg-a", "version": "1.0"}
					},
					{
						"vulnerability": {"id": "CVE-CRIT", "severity": "Critical"},
						"artifact": {"name": "pkg-b", "version": "1.0"}
					},
					{
						"vulnerability": {"id": "CVE-MED", "severity": "Medium"},
						"artifact": {"name": "pkg-c", "version": "1.0"}
					},
					{
						"vulnerability": {"id": "CVE-HIGH", "severity": "High"},
						"artifact": {"name": "pkg-d", "version": "1.0"}
					}
				]
			}`,
			wantCritical:     1,
			wantHigh:         1,
			wantMedium:       1,
			wantLow:          1,
			wantTotalVulns:   4,
			wantFirstVulnSev: "Critical",
			wantTimeParsed:   true,
			wantTimestamp:    "2024-06-01T12:00:00Z",
		},
		{
			name:    "invalid json",
			json:    `broken`,
			wantErr: true,
		},
		{
			name: "invalid timestamp with verbose logs warning",
			json: `{
				"descriptor": {"timestamp": "not-a-timestamp"},
				"matches": []
			}`,
			verbose:        true,
			wantTotalVulns: 0,
			wantTimeParsed: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stats, err := parseGrypeOutput([]byte(tt.json), tt.verbose)
			if tt.wantErr {
				if err == nil {
					t.Errorf("parseGrypeOutput() expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("parseGrypeOutput() unexpected error: %v", err)
			}
			if len(stats.Vulnerabilities) != tt.wantTotalVulns {
				t.Errorf("total vulns = %d, want %d", len(stats.Vulnerabilities), tt.wantTotalVulns)
			}
			if stats.VulnSummary["Critical"] != tt.wantCritical {
				t.Errorf("Critical = %d, want %d", stats.VulnSummary["Critical"], tt.wantCritical)
			}
			if stats.VulnSummary["High"] != tt.wantHigh {
				t.Errorf("High = %d, want %d", stats.VulnSummary["High"], tt.wantHigh)
			}
			if stats.VulnSummary["Medium"] != tt.wantMedium {
				t.Errorf("Medium = %d, want %d", stats.VulnSummary["Medium"], tt.wantMedium)
			}
			if stats.VulnSummary["Low"] != tt.wantLow {
				t.Errorf("Low = %d, want %d", stats.VulnSummary["Low"], tt.wantLow)
			}
			if tt.wantTotalVulns > 0 && tt.wantFirstVulnSev != "" {
				if stats.Vulnerabilities[0].Severity != tt.wantFirstVulnSev {
					t.Errorf("first vuln severity = %q, want %q",
						stats.Vulnerabilities[0].Severity, tt.wantFirstVulnSev)
				}
			}
			if tt.wantTimeParsed {
				expected, _ := time.Parse(time.RFC3339, tt.wantTimestamp)
				if !stats.VulnScanTime.Equal(expected) {
					t.Errorf("VulnScanTime = %v, want %v", stats.VulnScanTime, expected)
				}
			}
		})
	}
}

func TestParseGrypeOutput_SortStability(t *testing.T) {
	// When two vulns have the same severity, they should be sorted by ID ascending
	json := `{
		"descriptor": {"timestamp": "2024-01-01T00:00:00Z"},
		"matches": [
			{
				"vulnerability": {"id": "CVE-2024-0003", "severity": "High"},
				"artifact": {"name": "pkg-c", "version": "1.0"}
			},
			{
				"vulnerability": {"id": "CVE-2024-0001", "severity": "High"},
				"artifact": {"name": "pkg-a", "version": "1.0"}
			},
			{
				"vulnerability": {"id": "CVE-2024-0002", "severity": "High"},
				"artifact": {"name": "pkg-b", "version": "1.0"}
			}
		]
	}`

	stats, err := parseGrypeOutput([]byte(json), false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedOrder := []string{"CVE-2024-0001", "CVE-2024-0002", "CVE-2024-0003"}
	for i, vuln := range stats.Vulnerabilities {
		if vuln.ID != expectedOrder[i] {
			t.Errorf("vuln[%d].ID = %q, want %q", i, vuln.ID, expectedOrder[i])
		}
	}
}

func TestDiveRunner_Name(t *testing.T) {
	r := &DiveRunner{}
	if got := r.Name(); got != "dive" {
		t.Errorf("DiveRunner.Name() = %v, want %v", got, "dive")
	}
}

func TestParseDiveOutput(t *testing.T) {
	tests := []struct {
		name           string
		json           string
		wantErr        bool
		wantEfficiency float64
		wantWasted     string
	}{
		{
			name: "standard output",
			json: `{
				"image": {
					"inefficientBytes": 10485760,
					"efficiencyScore": 0.95
				}
			}`,
			wantEfficiency: 95.0,
			wantWasted:     "10.00 MB",
		},
		{
			name: "perfect efficiency",
			json: `{
				"image": {
					"inefficientBytes": 0,
					"efficiencyScore": 1.0
				}
			}`,
			wantEfficiency: 100.0,
			wantWasted:     "0.00 MB",
		},
		{
			name: "low efficiency",
			json: `{
				"image": {
					"inefficientBytes": 104857600,
					"efficiencyScore": 0.42
				}
			}`,
			wantEfficiency: 42.0,
			wantWasted:     "100.00 MB",
		},
		{
			name:    "invalid json",
			json:    `{broken`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stats, err := parseDiveOutput([]byte(tt.json))
			if tt.wantErr {
				if err == nil {
					t.Errorf("parseDiveOutput() expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("parseDiveOutput() unexpected error: %v", err)
			}
			if stats.Efficiency != tt.wantEfficiency {
				t.Errorf("Efficiency = %v, want %v", stats.Efficiency, tt.wantEfficiency)
			}
			if stats.WastedBytes != tt.wantWasted {
				t.Errorf("WastedBytes = %v, want %v", stats.WastedBytes, tt.wantWasted)
			}
		})
	}
}

func TestPodmanSocketDetection(t *testing.T) {
	// Tests the JSON parsing logic used in DiveRunner.Run() for Podman socket detection.
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

	if err := json.Unmarshal([]byte(testJSON), &machines); err != nil {
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
	tmpFile, err := os.CreateTemp("", "test-dive-output-*.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	tmpPath := tmpFile.Name()
	if err := tmpFile.Close(); err != nil {
		t.Fatalf("failed to close temp file: %v", err)
	}

	if _, err := os.Stat(tmpPath); os.IsNotExist(err) {
		t.Fatal("temp file should exist after creation")
	}

	if err := os.Remove(tmpPath); err != nil {
		t.Fatalf("failed to remove temp file: %v", err)
	}

	if _, err := os.Stat(tmpPath); !os.IsNotExist(err) {
		t.Error("temp file should not exist after removal")
	}
}
