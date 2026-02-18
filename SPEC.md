# Project Specification: Docker-Docs Generator

**Role:** You are a Senior Go Systems Engineer.
**Objective:** Build a CLI tool called `docker-docs` (conceptually similar to `terraform-docs`) that parses a `Dockerfile`, extracts metadata from instructions and "magic comments," and generates a Markdown documentation table.

## 1. Core Tech Stack [DONE]
* **Language:** Go (Golang) 1.25+
* **Parser:** `github.com/moby/buildkit/frontend/dockerfile/parser` (Do NOT use Regex; use the official AST parser).
* **CLI Framework:** `github.com/spf13/cobra`
* **Templating:** `text/template` (Standard library)

## 2. Functional Requirements[DONE]

### A. Parsing Logic (The AST Walker)
The tool must read a `Dockerfile` and extract the following instructions into a structured object:
1.  **ARG:** Build-time variables.
2.  **ENV:** Runtime environment variables.
3.  **LABEL:** OCI metadata.
4.  **EXPOSE:** Network ports.

**Critical Constraint:** The parser must handle standard Docker syntax idiosyncrasies:
* `ENV KEY=VALUE` (Equality) vs `ENV KEY VALUE` (Space-separated).
* Multi-line instructions using backslashes (`\`).

### B. "Magic Comment" Extraction
The tool must parse comments *immediately preceding* an instruction to extract metadata.
**Syntax Spec:**

```dockerfile
# @name: DB_HOST
# @description: The hostname of the database.
# @default: localhost
# @required: true
ENV DB_HOST="localhost"
```

* If `@default` is not provided in comments, infer it from the Dockerfile instruction value.
* If `@name` is not provided, use the variable name.

### C. Injection Strategy (Idempotency)
The tool must support updating an existing `README.md` without overwriting the whole file.
* **Markers:** Look for `` and ``.
* **Logic:** Replace only the content between these markers.
* **Fallback:** If no markers are found, print to `stdout`.

## 3. Data Structures[DONE]
Use this internal schema for the collected data:

```go
type DocItem struct {
    Name        string // e.g., "PORT"
    Value       string // inferred default from the instruction
    Description string // from @description
    Type        string // "ARG", "ENV", "LABEL", "EXPOSE"
    Required    bool   // from @required
}

type Documentation struct {
    Items []DocItem
}
```

## 4. Implementation Plan (Step-by-Step)[DONE]

**Phase 1: The Parser**
1.  Create `pkg/parser/parser.go`.
2.  Implement `Parse(filename string) (*Documentation, error)`.
3.  Use `parser.Parse()` from Moby BuildKit to get the AST.
4.  Iterate through `Node.Children`.
5.  Implement a helper `extractComments(node *Node)` that reads `node.PrevComment` and scans for `@tag` keys.

**Phase 2: The Renderer**
1.  Create `pkg/renderer/markdown.go`.
2.  Define a default Markdown template containing a table with columns: `| Name | Type | Description | Default | Required |`.
3.  Implement `Render(doc *Documentation) (string, error)`.

**Phase 3: The CLI**
1.  Create `cmd/root.go` using Cobra.
2.  Add flags:
    * `--file, -f`: Path to Dockerfile (default "./Dockerfile").
    * `--output, -o`: Path to output file (default "README.md").
    * `--dry-run`: Print to stdout instead of writing to file.

## 5. Execution[DONE]
Please start by initializing the Go module and creating the **Phase 1 (Parser)** implementation. Write a test case with a sample Dockerfile containing both `ENV` and `ARG` with magic comments to verify the parsing logic.

## 6. Phase 4: Dynamic Analysis (External Tools)[DONE]

**Objective:** Extend the tool to inspect *built* container images using external CLI tools.
**Architecture:** Hybrid Analysis. The tool will now shell out to `docker`, `syft`, `grype`, and `dive` to capture runtime data and merge it with the static Dockerfile analysis.

### A. Core Architecture: The "Runner" Pattern
**Constraint:** Do NOT import `syft` or `dive` as Go libraries. This will cause dependency hell.
Instead, use `os/exec` to run them as sub-processes and parse their JSON output.

1.  **New Interface:** Create `pkg/runner/runner.go`.
    ```go
    type ToolRunner interface {
        Name() string
        IsAvailable() bool // e.g. exec.LookPath("syft")
        Run(image string) (interface{}, error)
    }
    ```
2.  **New Data Structs:** Create `pkg/analysis/stats.go`.
    ```go
    type ImageStats struct {
        Architecture string
        OS           string
        SizeMB       string
        TotalLayers  int
        Efficiency   float64 // from Dive (0-100)
        WastedBytes  string  // from Dive
        Packages     []PackageSummary // from Syft
        VulnSummary  map[string]int   // from Grype (Severity -> Count)
    }
    ```

### B. Tool Integrations

#### 1. Docker Inspect (Baseline)
* **Command:** `docker inspect <image>`
* **Data to Extract:**
    * `Architecture` & `Os`
    * `Size` (Convert bytes to human-readable MB/GB)
    * `RootFS.Layers` (Count the length of the array)

#### 2. Syft (SBOM / Packages)
* **Command:** `syft <image> -o json`
* **Data to Extract:**
    * Parse the `artifacts` array.
    * **Logic:** Do NOT list every single package (it's too long).
    * **Action:** Extract the count of total packages.
    * **Action:** Extract a list of "Key Frameworks" if detected (e.g., `python`, `node`, `go`, `glibc`, `openssl`).

#### 3. Grype (Vulnerabilities)
* **Command:** `grype <image> -o json`
* **Data to Extract:**
    * Parse `matches`.
    * Group by `vulnerability.severity` (Critical, High, Medium, Low).
    * Return a simple map: `{"Critical": 2, "High": 5}`.

#### 4. Dive (Image Efficiency)
* **Command:** `dive <image> --json output.json`
    * *Note: Dive writes to a file, not stdout.*
* **Data to Extract:**
    * `image.inefficientBytes`
    * `image.efficiencyScore`

### C. CLI Updates
* Add flag: `--image <tag>` (e.g., `my-app:latest`).
* **Logic:**
    * If `--image` is present, run the active Runners.
    * If a tool (e.g., `dive`) is missing from the system `PATH`, log a warning but **do not fail**. Just skip that section of the report.

### D. Renderer Updates
Update the Markdown template to include a "Runtime Analysis" section if the data is available.

**Proposed Template Addition:**
```markdown
## Image Analysis ({{ .ImageTag }})

| Metric | Value |
|--------|-------|
| Size | {{ .Stats.SizeMB }} |
| Architecture | {{ .Stats.Architecture }}/{{ .Stats.OS }} |
| Efficiency | {{ .Stats.Efficiency }}% ({{ .Stats.WastedBytes }} wasted) |

**Security Summary:**
Critical: {{ .Stats.VulnSummary.Critical }} | High: {{ .Stats.VulnSummary.High }} | Medium: {{ .Stats.VulnSummary.Medium }}

<details>
<summary>Key Packages ({{ .Stats.TotalPackages }} total)</summary>

| Package | Version |
|---------|---------|
{{ range .Stats.Packages }}| {{ .Name }} | {{ .Version }} |
{{ end }}
</details>

## 7. Phase 5: Visual Polish & Grouping [DONE]

**Objective:** Refine the Markdown output to be "GitHub-native," scannable, and professional.
**Changes:**
1.  **Split Tables:** Instead of one giant table, group items by type (`ENV`, `ARG`, `LABEL`, `EXPOSE`).
2.  **Badges:** Generate Shields.io badge URLs for high-level stats (Size, Layers, Vulns).
3.  **Visual Hierarchy:** Use emojis and specific formatting to highlight security status (e.g., 🔴 for Critical, 🟢 for Safe).

### A. Data Structure Update
Update `pkg/analysis/stats.go` to include helper methods for dynamic badge generation and filtering.

1.  **Badge Helpers:**
    * `GetSizeBadge() string` -> returns `https://img.shields.io/badge/size-7.6MB-blue`
    * `GetVulnBadge() string` -> returns `https://img.shields.io/badge/vulns-0%20Critical-green` (URL encoded).
    * **Logic:** If Critical Vulns > 0, set Badge Color to `red`. Else `green`.

2.  **Filter Helper:**
    * Add a method to the `Documentation` struct: `FilterByType(type string) []DocItem`.
    * This allows the template to call `{{ range .FilterByType "ENV" }}` to only show environment variables in a specific section.

### B. Updated Markdown Template
Replace the existing template in `pkg/renderer/markdown.go` with this "Dashboard Style" layout.

```markdown
# 🐳 Docker Image Analysis: {{ .ImageTag }}

![Size]({{ .Stats.SizeBadge }}) ![Layers]({{ .Stats.LayersBadge }}) ![Vulns]({{ .Stats.VulnBadge }}) ![Efficiency]({{ .Stats.EfficiencyBadge }})

## ⚙️ Configuration

### Environment Variables
| Name | Description | Default | Required |
|------|-------------|---------|:--------:|
{{- range .FilterByType "ENV" }}
| `{{ .Name }}` | {{ .Description }} | `{{ if .Value }}{{ .Value }}{{ else }}""{{ end }}` | {{ if .Required }}✅{{ else }}❌{{ end }} |
{{- end }}

### Build Arguments
| Name | Description | Default | Required |
|------|-------------|---------|:--------:|
{{- range .FilterByType "ARG" }}
| `{{ .Name }}` | {{ .Description }} | `{{ .Value }}` | {{ if .Required }}✅{{ else }}❌{{ end }} |
{{- end }}

### Exposed Ports
| Port | Description |
|------|-------------|
{{- range .FilterByType "EXPOSE" }}
| `{{ .Name }}` | {{ .Description }} |
{{- end }}

---

## 🛡️ Security & Efficiency 

**Base Image:** `{{ .Stats.OS }} ({{ .Stats.Architecture }})`
**Efficiency Score:** {{ .Stats.Efficiency }}%

### Vulnerabilities
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| {{ if gt .Stats.VulnSummary.Critical 0 }}🔴 {{ else }}🟢 {{ end }}{{ .Stats.VulnSummary.Critical }} | {{ if gt .Stats.VulnSummary.High 0 }}🟠 {{ else }}🟢 {{ end }}{{ .Stats.VulnSummary.High }} | {{ .Stats.VulnSummary.Medium }} | {{ .Stats.VulnSummary.Low }} |

<details>
<summary><strong>👇 Expand Vulnerability Details ({{ .Stats.TotalVulns }} found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
{{- range .Stats.Vulns }}
| [{{ .ID }}](https://nvd.nist.gov/vuln/detail/{{ .ID }}) | {{ .Severity }} | `{{ .Package }}` | `{{ .Version }}` |
{{- end }}
</details>

<details>
<summary><strong>📦 Installed Packages ({{ .Stats.TotalPackages }} total)</strong></summary>

| Package | Version |
|---------|---------|
{{- range .Stats.Packages }}
| {{ .Name }} | {{ .Version }} |
{{- end }}
</details>
```

C. Implementation Details

Shields.io Usage: Use the static badge endpoint: https://img.shields.io/static/v1?label=<LABEL>&message=<VALUE>&color=<COLOR>.

Template Logic: Ensure the template handles empty lists gracefully (e.g., if there are no ARGs, the table header should probably be hidden or the section skipped).

Here is the Markdown block for **Phase 8**. Append this to your `SPEC.md`.

```markdown
## 8. Phase 6: Comparison Mode & Configuration

**Objective:** Support multi-variant documentation (e.g., `python:3.12`, `python:3.13`, `python:alpine`) by introducing a configuration file and a "Tag Comparison" table.
**Architecture:**
* **New Input:** `docker-docs.yaml` (replaces complex CLI flags).
* **New Logic:** The "Manager" reads the config, runs runners in parallel for each image, and renders multiple templates into named markers.

### A. Configuration Schema
Create `pkg/config/config.go` to parse `docker-docs.yaml`.

```yaml
# docker-docs.yaml
output: "README.md"

sections:
  # 1. Main Configuration (ENV/ARG)
  # Uses the local Dockerfile or a specific image as the "Source of Truth"
  - type: "image"
    marker: "main-config" # source: "Dockerfile"  # or tag: "my-app:latest"

  # 2. Tag Comparison (Dynamic Stats)
  # Scans multiple images to create a comparison table
  - type: "comparison"
    marker: "tag-comparison"
    images:
      - tag: "my-app:3.12"
      - tag: "my-app:3.13"
      - tag: "my-app:3.12-alpine"
      - tag: "my-app:3.13-alpine"

```

### B. The "Tag Comparison" Renderer

Create a new template `pkg/renderer/templates/comparison.md` that renders a row for each image in the `comparison` section.

**Template Logic:**

1. Iterate over the list of `images` defined in the YAML.
2. Run the `ToolRunner` (Phase 4) on each image.
3. Aggregate the results into a `[]ImageStats` slice.
4. Render a comparison table.

**Proposed Template:**

```markdown
### Supported Tags

| Tag | Size | Vulns | Efficiency | OS/Arch |
|-----|------|-------|------------|---------|
{{- range .Images }}
| `{{ .Tag }}` | ![Size]({{ .Stats.SizeBadge }}) | ![Vulns]({{ .Stats.VulnBadge }}) | {{ .Stats.Efficiency }}% | {{ .Stats.OS }}/{{ .Stats.Architecture }} |
{{- end }}

```

### C. Implementation Details

1. **Named Injection Markers:**
* Update the `Injector` (Phase 2) to support **Named Markers**.
* Old: ``
* New: ``
* **Logic:** The tool must parse the README, find the specific marker name defined in the YAML, and replace only that section.


2. **Parallel Execution (Critical):**
* Scanning 5+ images sequentially will be very slow.
* Use `golang.org/x/sync/errgroup` to run the `ToolRunner` for each image in the matrix in parallel.


3. **CLI Update:**
* Modify `cmd/root.go` to look for `docker-docs.yaml` by default.
* Add flag `--config` to specify a custom config file path.

---

### 📋 Phase 7: Open Source Packaging & Distribution

**Objective:** Prepare the repository for public release on GitHub.
**Key Deliverables:**

1. **Automated Releases:** Use `GoReleaser` to build binaries and Docker images on every tag (e.g., `v1.0.0`).
2. **GitHub Action:** Create a `action.yml` so users can use this tool directly in their CI workflows.
3. **Documentation:** Generate a polished `README.md` (using the tool itself!) and standard OSS files.

### A. GoReleaser Configuration

Create a `.goreleaser.yaml` file at the root.

**Requirements:**

* **Builds:** Linux (amd64/arm64), Darwin (amd64/arm64), Windows (amd64).
* **Archives:** `.tar.gz` replacements.
* **Docker:** Build a container image `ghcr.io/OWNER/docker-docs:latest` and push it to GitHub Container Registry.
* *Note:* Use the existing `Dockerfile` but ensure the entrypoint is correct.



### B. GitHub Action (Composite)

Create `action.yml` in the root directory. This allows your tool to be used like `uses: username/docker-docs@v1`.

**Spec:**

```yaml
name: 'Docker Documentation Generator'
description: 'Automatically generate README documentation from Dockerfiles'
inputs:
  config:
    description: 'Path to configuration file'
    default: 'docker-docs.yaml'
  output:
    description: 'Output file path'
    default: 'README.md'
runs:
  using: 'docker'
  image: 'Dockerfile' # Uses the tool's own Dockerfile to run itself
  args:
    - "--config"
    - ${{ inputs.config }}
    - "--output"
    - ${{ inputs.output }}

```

### C. CI/CD Workflows

Create `.github/workflows/` with two files:

1. **`test.yml`**: Runs `go test ./...` and `golangci-lint` on every PR.
2. **`release.yml`**: Runs `goreleaser/goreleaser-action` only on tags starting with `v*`.
* *Permissions:* Needs `contents: write` and `packages: write`.



### D. Documentation & Legal

1. **LICENSE:** Generate an MIT License file.
2. **CONTRIBUTING.md:** Brief guide on how to build locally (`go build`) and run tests.
3. **README.md:**
* Update the root README to be "User Facing" (Installation guide, Usage examples).
* **Dogfooding:** Ensure the `docker-docs` tool is used to generate the "Configuration" section of its *own* README.



---

### 🚀 How to Execute This

1. **Append** the text above to your `SPEC.md`.
2. **Manual Step (Crucial):** Before you run the agent, you must decide on your GitHub path.
* Run this in your terminal: `go mod edit -module github.com/YOUR_USERNAME/docker-docs`
* (Replace `YOUR_USERNAME` with your actual GitHub handle).


3. **Run the Agent:**
* Prompt: *"Implement Phase 9. Start by creating the .goreleaser.yaml file and the action.yml."*



### Why this approach works:

* **GoReleaser** is magic. When you push a tag `v0.1.0`, it automatically builds binaries for Windows, Mac, and Linux, creates a Docker image, pushes it to GHCR, and creates a "Release" page on GitHub with changelogs.
* **The `action.yml**` means users don't even need to install your tool. They can just add 3 lines to their YAML and it works. This is the #1 way to get adoption for this kind of tool.