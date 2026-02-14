package parser

import (
	"bufio"
	"os"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

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

// Parse reads a Dockerfile and extracts documentation metadata.
func Parse(filename string) (*Documentation, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	result, err := parser.Parse(f)
	if err != nil {
		return nil, err
	}

	doc := &Documentation{
		Items: make([]DocItem, 0),
	}

	for _, node := range result.AST.Children {
		item := DocItem{}
		// Look for comments immediately preceding the instruction
		parseComments(node, &item)

		switch strings.ToUpper(node.Value) {
		case "ARG":
			item.Type = "ARG"
			parseArgEnv(node, &item)
		case "ENV":
			item.Type = "ENV"
			parseArgEnv(node, &item)
		case "LABEL":
			item.Type = "LABEL"
			parseLabel(node, &item)
		case "EXPOSE":
			item.Type = "EXPOSE"
			item.Name = node.Next.Value // Usually just the port number
			if item.Value == "" {
				item.Value = node.Next.Value
			}
		default:
			continue
		}

		// If we found something relevant, add it
		if item.Name != "" || item.Type == "EXPOSE" {
			// For EXPOSE, name might be ambiguous, usually it's the port.
			// If name wasn't set by comment, use the port as name?
			// The spec says: Name string // e.g., "PORT".
			// For EXPOSE 8080, Name might be "8080" effectively if not named.
			if item.Name == "" && item.Type == "EXPOSE" {
				item.Name = item.Value
			}
			doc.Items = append(doc.Items, item)
		}
	}

	return doc, nil
}

func parseComments(node *parser.Node, item *DocItem) {
	if node.PrevComment == nil {
		return
	}
	scanner := bufio.NewScanner(strings.NewReader(strings.Join(node.PrevComment, "\n")))
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		// Magic comments start with #
		// but parser already gives us the comment lines.
		// However, buildkit parser might include the # or not depending on version?
		// Checking source: PrevComment is slice of strings, usually includes the #.

		cleanLine := strings.TrimPrefix(line, "#")
		cleanLine = strings.TrimSpace(cleanLine)

		if strings.HasPrefix(cleanLine, "@name:") {
			item.Name = strings.TrimSpace(strings.TrimPrefix(cleanLine, "@name:"))
		} else if strings.HasPrefix(cleanLine, "@description:") {
			item.Description = strings.TrimSpace(strings.TrimPrefix(cleanLine, "@description:"))
		} else if strings.HasPrefix(cleanLine, "@default:") {
			item.Value = strings.TrimSpace(strings.TrimPrefix(cleanLine, "@default:"))
		} else if strings.HasPrefix(cleanLine, "@required:") {
			val := strings.TrimSpace(strings.TrimPrefix(cleanLine, "@required:"))
			if val == "true" {
				item.Required = true
			}
		}
	}
}

// Helper to strip surrounding quotes
func stripQuotes(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}

func parseArgEnv(node *parser.Node, item *DocItem) {
	// Format: ENV KEY=VALUE or ENV KEY VALUE
	// Node structure for ENV/ARG usually has children or Next pointer chain.
	// Buildkit parser:
	// "ENV KEY=VALUE" -> Value="ENV", Next -> Value="KEY=VALUE"
	// "ENV KEY VALUE" -> Value="ENV", Next -> Value="KEY", Next -> Value="VALUE"

	if node.Next == nil {
		return
	}

	// Helper to handle KEY=VALUE vs KEY VALUE
	// In newer buildkit parser, it might normalize, but let's check basic structure.

	firstArg := node.Next.Value

	if strings.Contains(firstArg, "=") {
		// KEY=VALUE format
		parts := strings.SplitN(firstArg, "=", 2)
		if item.Name == "" {
			item.Name = parts[0]
		}
		// Only set value (default) if not overridden by comment
		if item.Value == "" && len(parts) > 1 {
			item.Value = stripQuotes(parts[1])
		}
	} else {
		// KEY VALUE format
		if item.Name == "" {
			item.Name = firstArg
		}
		if item.Value == "" && node.Next.Next != nil {
			item.Value = stripQuotes(node.Next.Next.Value)
		}
	}
}

func parseLabel(node *parser.Node, item *DocItem) {
	// LABEL usually is KEY=VALUE
	if node.Next == nil {
		return
	}

	first := node.Next.Value

	if strings.Contains(first, "=") {
		parts := strings.SplitN(first, "=", 2)
		if item.Name == "" {
			item.Name = parts[0]
		}
		item.Value = stripQuotes(parts[1])
	} else {
		// Might be separated nodes
		if item.Name == "" {
			item.Name = first
		}
		if node.Next.Next != nil {
			item.Value = stripQuotes(node.Next.Next.Value)
		}
	}
}
