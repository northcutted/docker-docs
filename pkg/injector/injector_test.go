package injector

import (
	"strings"
	"testing"
)

func TestInject_DefaultMarkers(t *testing.T) {
	originalContent := `# My Project

Some text here.

<!-- BEGIN: dock-docs -->
Old content that will be replaced.
<!-- END: dock-docs -->

More text after.
`

	newContent := "## New Documentation\nThis is the new content."

	result, err := Inject(originalContent, "", newContent)
	if err != nil {
		t.Fatalf("Inject() error = %v", err)
	}

	if !strings.Contains(result, "<!-- BEGIN: dock-docs -->") {
		t.Error("Result should contain BEGIN marker")
	}

	if !strings.Contains(result, "<!-- END: dock-docs -->") {
		t.Error("Result should contain END marker")
	}

	if !strings.Contains(result, "## New Documentation") {
		t.Error("Result should contain new content")
	}

	if strings.Contains(result, "Old content that will be replaced") {
		t.Error("Result should not contain old content")
	}

	if !strings.Contains(result, "Some text here") {
		t.Error("Result should preserve content before markers")
	}

	if !strings.Contains(result, "More text after") {
		t.Error("Result should preserve content after markers")
	}
}

func TestInject_NamedMarkers(t *testing.T) {
	originalContent := `# Documentation

<!-- BEGIN: dock-docs:main -->
Old main content
<!-- END: dock-docs:main -->

<!-- BEGIN: dock-docs:matrix -->
Old matrix content
<!-- END: dock-docs:matrix -->
`

	newContent := "New main section content"

	result, err := Inject(originalContent, "main", newContent)
	if err != nil {
		t.Fatalf("Inject() error = %v", err)
	}

	if !strings.Contains(result, "<!-- BEGIN: dock-docs:main -->") {
		t.Error("Result should contain named BEGIN marker")
	}

	if !strings.Contains(result, "New main section content") {
		t.Error("Result should contain new content")
	}

	if strings.Contains(result, "Old main content") {
		t.Error("Result should not contain old main content")
	}

	// Matrix section should remain unchanged
	if !strings.Contains(result, "Old matrix content") {
		t.Error("Result should preserve other sections")
	}
}

func TestInject_MissingBeginMarker(t *testing.T) {
	originalContent := `# Documentation

Some content here
<!-- END: dock-docs -->
`

	newContent := "New content"

	_, err := Inject(originalContent, "", newContent)
	if err == nil {
		t.Error("Expected error for missing BEGIN marker, got nil")
	}

	if !strings.Contains(err.Error(), "not found") {
		t.Errorf("Error message should mention markers not found, got: %v", err)
	}
}

func TestInject_MissingEndMarker(t *testing.T) {
	originalContent := `# Documentation

<!-- BEGIN: dock-docs -->
Some content here
`

	newContent := "New content"

	_, err := Inject(originalContent, "", newContent)
	if err == nil {
		t.Error("Expected error for missing END marker, got nil")
	}

	if !strings.Contains(err.Error(), "not found") {
		t.Errorf("Error message should mention markers not found, got: %v", err)
	}
}

func TestInject_MarkersInWrongOrder(t *testing.T) {
	originalContent := `# Documentation

<!-- END: dock-docs -->
Some content
<!-- BEGIN: dock-docs -->
`

	newContent := "New content"

	_, err := Inject(originalContent, "", newContent)
	if err == nil {
		t.Error("Expected error for wrong marker order, got nil")
	}

	if !strings.Contains(err.Error(), "before start marker") {
		t.Errorf("Error message should mention marker order, got: %v", err)
	}
}

func TestInject_EmptyNewContent(t *testing.T) {
	originalContent := `# Documentation

<!-- BEGIN: dock-docs -->
Old content
<!-- END: dock-docs -->

More text.
`

	newContent := ""

	result, err := Inject(originalContent, "", newContent)
	if err != nil {
		t.Fatalf("Inject() error = %v", err)
	}

	// Should work with empty content
	if !strings.Contains(result, "<!-- BEGIN: dock-docs -->") {
		t.Error("Result should contain BEGIN marker")
	}

	if !strings.Contains(result, "<!-- END: dock-docs -->") {
		t.Error("Result should contain END marker")
	}

	if strings.Contains(result, "Old content") {
		t.Error("Result should not contain old content")
	}
}

func TestInject_MultilineNewContent(t *testing.T) {
	originalContent := `# Project

<!-- BEGIN: dock-docs -->
<!-- END: dock-docs -->
`

	newContent := `## Section 1
Content line 1

## Section 2
Content line 2

## Section 3
Content line 3`

	result, err := Inject(originalContent, "", newContent)
	if err != nil {
		t.Fatalf("Inject() error = %v", err)
	}

	if !strings.Contains(result, "## Section 1") {
		t.Error("Result should contain section 1")
	}

	if !strings.Contains(result, "## Section 2") {
		t.Error("Result should contain section 2")
	}

	if !strings.Contains(result, "## Section 3") {
		t.Error("Result should contain section 3")
	}

	// Count newlines to verify structure
	lines := strings.Split(result, "\n")
	if len(lines) < 10 {
		t.Errorf("Result should have multiple lines, got %d", len(lines))
	}
}

func TestInject_PreserveWhitespace(t *testing.T) {
	originalContent := `
Prefix content

<!-- BEGIN: dock-docs -->
Old
<!-- END: dock-docs -->

Suffix content
`

	newContent := "New"

	result, err := Inject(originalContent, "", newContent)
	if err != nil {
		t.Fatalf("Inject() error = %v", err)
	}

	// Verify structure preservation
	if !strings.HasPrefix(result, "\nPrefix content") {
		t.Error("Result should preserve prefix whitespace")
	}

	if !strings.HasSuffix(result, "Suffix content\n") {
		t.Error("Result should preserve suffix whitespace")
	}
}

func TestInject_MultipleNamedSections(t *testing.T) {
	originalContent := `# README

<!-- BEGIN: dock-docs:production -->
Prod old
<!-- END: dock-docs:production -->

<!-- BEGIN: dock-docs:development -->
Dev old
<!-- END: dock-docs:development -->

<!-- BEGIN: dock-docs:testing -->
Test old
<!-- END: dock-docs:testing -->
`

	// Update production section
	result1, err := Inject(originalContent, "production", "Prod new")
	if err != nil {
		t.Fatalf("Inject(production) error = %v", err)
	}

	if !strings.Contains(result1, "Prod new") {
		t.Error("Should contain new production content")
	}

	if strings.Contains(result1, "Prod old") {
		t.Error("Should not contain old production content")
	}

	if !strings.Contains(result1, "Dev old") {
		t.Error("Should preserve dev section")
	}

	// Update development section
	result2, err := Inject(result1, "development", "Dev new")
	if err != nil {
		t.Fatalf("Inject(development) error = %v", err)
	}

	if !strings.Contains(result2, "Prod new") {
		t.Error("Should preserve production changes")
	}

	if !strings.Contains(result2, "Dev new") {
		t.Error("Should contain new development content")
	}

	if strings.Contains(result2, "Dev old") {
		t.Error("Should not contain old development content")
	}
}

func TestInject_SpecialCharacters(t *testing.T) {
	originalContent := `# Docs

<!-- BEGIN: dock-docs -->
Old
<!-- END: dock-docs -->
`

	newContent := `Content with special chars: <>&"'
And some code: if (x > 0) { return true; }
And a table: | Col1 | Col2 |
             |------|------|
             | A    | B    |`

	result, err := Inject(originalContent, "", newContent)
	if err != nil {
		t.Fatalf("Inject() error = %v", err)
	}

	if !strings.Contains(result, "<>&\"'") {
		t.Error("Result should preserve special HTML characters")
	}

	if !strings.Contains(result, "if (x > 0)") {
		t.Error("Result should preserve code snippets")
	}

	if !strings.Contains(result, "| Col1 | Col2 |") {
		t.Error("Result should preserve markdown tables")
	}
}

func TestInject_EdgeCase_EmptyFile(t *testing.T) {
	originalContent := ""
	newContent := "New content"

	_, err := Inject(originalContent, "", newContent)
	if err == nil {
		t.Error("Expected error for empty file, got nil")
	}
}

func TestInject_EdgeCase_MarkersOnly(t *testing.T) {
	originalContent := `<!-- BEGIN: dock-docs -->
<!-- END: dock-docs -->`

	newContent := "Inserted content"

	result, err := Inject(originalContent, "", newContent)
	if err != nil {
		t.Fatalf("Inject() error = %v", err)
	}

	if !strings.Contains(result, "Inserted content") {
		t.Error("Result should contain inserted content")
	}

	// Should have markers and content
	lines := strings.Split(result, "\n")
	if len(lines) < 3 {
		t.Errorf("Result should have at least 3 lines (BEGIN, content, END), got %d", len(lines))
	}
}
