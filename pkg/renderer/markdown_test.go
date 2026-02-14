package renderer

import (
	"strings"
	"testing"

	"docker-docs/pkg/parser"
)

func TestRender(t *testing.T) {
	doc := &parser.Documentation{
		Items: []parser.DocItem{
			{
				Name:        "PORT",
				Type:        "ENV",
				Description: "Port to listen on",
				Value:       "8080",
				Required:    true,
			},
			{
				Name:        "DB_HOST",
				Type:        "ARG",
				Description: "Database host",
				Value:       "",
				Required:    false,
			},
		},
	}

	output, err := Render(doc)
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	expectedHeader := "| Name | Type | Description | Default | Required |"
	if !strings.Contains(output, expectedHeader) {
		t.Errorf("expected output to contain header %q", expectedHeader)
	}

	// Verify row content
	if !strings.Contains(output, "| PORT | ENV | Port to listen on | 8080 | true |") {
		t.Errorf("expected output to contain PORT row")
	}
	if !strings.Contains(output, "| DB_HOST | ARG | Database host |  | false |") {
		t.Errorf("expected output to contain DB_HOST row")
	}
}
