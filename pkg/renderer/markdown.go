package renderer

import (
	"bytes"
	"text/template"

	"docker-docs/pkg/parser"
)

const defaultTemplate = `| Name | Type | Description | Default | Required |
|------|------|-------------|---------|----------|
{{- range .Items }}
| {{ .Name }} | {{ .Type }} | {{ .Description }} | {{ .Value }} | {{ .Required }} |
{{- end }}
`

// Render generates the Markdown table from documentation items.
func Render(doc *parser.Documentation) (string, error) {
	tmpl, err := template.New("docker-docs").Parse(defaultTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, doc); err != nil {
		return "", err
	}

	return buf.String(), nil
}
