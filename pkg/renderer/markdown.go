package renderer

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/northcutted/docker-docs/pkg/analysis"
	"github.com/northcutted/docker-docs/pkg/parser"
)

// ReportContext holds all data passed to the template
type ReportContext struct {
	Doc      *parser.Documentation
	Stats    *analysis.ImageStats
	ImageTag string
}

const defaultTemplate = `
# 🐳 Docker Image Analysis: {{ .ImageTag }}

{{- if .Stats }}
![Size]({{ .Stats.SizeBadge }}) ![Layers]({{ .Stats.LayersBadge }}) ![Vulns]({{ .Stats.VulnBadge }}) ![Efficiency]({{ .Stats.EfficiencyBadge }})
{{- end }}

## ⚙️ Configuration

{{- if (len (.Doc.FilterByType "ENV")) }}
### Environment Variables
| Name | Description | Default | Required |
|------|-------------|---------|:--------:|
{{- range (.Doc.FilterByType "ENV") }}
| ` + "`{{ .Name }}`" + ` | {{ .Description }} | ` + "`{{ if .Value }}{{ .Value }}{{ else }}\"\"{{ end }}`" + ` | {{ if .Required }}✅{{ else }}❌{{ end }} |
{{- end }}
{{- end }}

{{- if (len (.Doc.FilterByType "ARG")) }}
### Build Arguments
| Name | Description | Default | Required |
|------|-------------|---------|:--------:|
{{- range (.Doc.FilterByType "ARG") }}
| ` + "`{{ .Name }}`" + ` | {{ .Description }} | ` + "`{{ .Value }}`" + ` | {{ if .Required }}✅{{ else }}❌{{ end }} |
{{- end }}
{{- end }}

{{- if (len (.Doc.FilterByType "EXPOSE")) }}
### Exposed Ports
| Port | Description |
|------|-------------|
{{- range (.Doc.FilterByType "EXPOSE") }}
| ` + "`{{ .Name }}`" + ` | {{ .Description }} |
{{- end }}
{{- end }}

{{- if (len (.Doc.FilterByType "LABEL")) }}
### Labels
| Key | Value |
|-----|-------|
{{- range (.Doc.FilterByType "LABEL") }}
| ` + "`{{ .Name }}`" + ` | {{ .Value }} |
{{- end }}
{{- end }}

{{- if .Stats }}
---

## 🛡️ Security & Efficiency

**Base Image:** ` + "`{{ .Stats.OS }} ({{ .Stats.Architecture }})`" + `
{{- if .Stats.SupportedArchitectures }}
**Supported Architectures:** ` + "`{{ join .Stats.SupportedArchitectures \", \" }}`" + `
{{- end }}
**Efficiency Score:** {{ printf "%.1f" .Stats.Efficiency }}%

### Vulnerabilities
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| {{ if gt (index .Stats.VulnSummary "Critical") 0 }}🔴 {{ else }}🟢 {{ end }}{{ index .Stats.VulnSummary "Critical" }} | {{ if gt (index .Stats.VulnSummary "High") 0 }}🟠 {{ else }}🟢 {{ end }}{{ index .Stats.VulnSummary "High" }} | {{ index .Stats.VulnSummary "Medium" }} | {{ index .Stats.VulnSummary "Low" }} |

<details>
<summary><strong>👇 Expand Vulnerability Details ({{ .Stats.TotalVulns }} found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
{{- range .Stats.Vulnerabilities }}
| [{{ .ID }}](https://nvd.nist.gov/vuln/detail/{{ .ID }}) | {{ .Severity }} | ` + "`{{ .Package }}`" + ` | ` + "`{{ .Version }}`" + ` |
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
{{- end }}
`

// Render generates the Markdown table from documentation items.
func Render(doc *parser.Documentation, stats *analysis.ImageStats) (string, error) {
	tmpl, err := template.New("docker-docs").Funcs(template.FuncMap{
		"index": func(m map[string]int, k string) int {
			if v, ok := m[k]; ok {
				return v
			}
			return 0
		},
		"join": strings.Join,
	}).Parse(defaultTemplate)

	if err != nil {
		return "", err
	}

	ctx := ReportContext{
		Doc:   doc,
		Stats: stats,
	}
	if stats != nil {
		ctx.ImageTag = stats.ImageTag
	} else {
		// Fallback if no stats provided
		if ctx.ImageTag == "" {
			ctx.ImageTag = "Dockerfile"
		}
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, ctx); err != nil {
		return "", err
	}

	return buf.String(), nil
}

type MatrixContext struct {
	Matrix []*analysis.ImageStats
}

const matrixTemplate = `
### 🏷️ Supported Tags

| Tag | Size | Vulns | Efficiency | Architectures |
|-----|------|-------|------------|---------------|
{{- range .Matrix }}
| ` + "`{{ .ImageTag }}`" + ` | ![Size]({{ .SizeBadge }}) | ![Vulns]({{ .VulnBadge }}) | {{ printf "%.1f" .Efficiency }}% | {{ if .SupportedArchitectures }}` + "`{{ join .SupportedArchitectures \", \" }}`" + `{{ else }}` + "`{{ .OS }}/{{ .Architecture }}`" + `{{ end }} |
{{- end }}

{{- range .Matrix }}

<details>
<summary><strong>🔍 Full Report: {{ .ImageTag }}</strong></summary>

## 🛡️ Security & Efficiency

**Base Image:** ` + "`{{ .OS }} ({{ .Architecture }})`" + `
{{- if .SupportedArchitectures }}
**Supported Architectures:** ` + "`{{ join .SupportedArchitectures \", \" }}`" + `
{{- end }}
**Efficiency Score:** {{ printf "%.1f" .Efficiency }}%

### Vulnerabilities
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| {{ if gt (index .VulnSummary "Critical") 0 }}🔴 {{ else }}🟢 {{ end }}{{ index .VulnSummary "Critical" }} | {{ if gt (index .VulnSummary "High") 0 }}🟠 {{ else }}🟢 {{ end }}{{ index .VulnSummary "High" }} | {{ index .VulnSummary "Medium" }} | {{ index .VulnSummary "Low" }} |

<details>
<summary><strong>👇 Expand Vulnerability Details ({{ .TotalVulns }} found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
{{- range .Vulnerabilities }}
| [{{ .ID }}](https://nvd.nist.gov/vuln/detail/{{ .ID }}) | {{ .Severity }} | ` + "`{{ .Package }}`" + ` | ` + "`{{ .Version }}`" + ` |
{{- end }}
</details>

<details>
<summary><strong>📦 Installed Packages ({{ .TotalPackages }} total)</strong></summary>

| Package | Version |
|---------|---------|
{{- range .Packages }}
| {{ .Name }} | {{ .Version }} |
{{- end }}
</details>

</details>
{{- end }}
`

// RenderMatrix generates the comparison table for multiple images.
func RenderMatrix(stats []*analysis.ImageStats) (string, error) {
	tmpl, err := template.New("docker-docs-matrix").Funcs(template.FuncMap{
		"index": func(m map[string]int, k string) int {
			if v, ok := m[k]; ok {
				return v
			}
			return 0
		},
		"join": strings.Join,
	}).Parse(matrixTemplate)
	if err != nil {
		return "", err
	}

	ctx := MatrixContext{
		Matrix: stats,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, ctx); err != nil {
		return "", err
	}

	return buf.String(), nil
}
