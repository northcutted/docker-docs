# Project Specification: Docker-Docs Generator

**Role:** You are a Senior Go Systems Engineer.
**Objective:** Build a CLI tool called `docker-docs` (conceptually similar to `terraform-docs`) that parses a `Dockerfile`, extracts metadata from instructions and "magic comments," and generates a Markdown documentation table.

## 1. Core Tech Stack
* **Language:** Go (Golang) 1.25+
* **Parser:** `github.com/moby/buildkit/frontend/dockerfile/parser` (Do NOT use Regex; use the official AST parser).
* **CLI Framework:** `github.com/spf13/cobra`
* **Templating:** `text/template` (Standard library)

## 2. Functional Requirements

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

## 3. Data Structures
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

## 4. Implementation Plan (Step-by-Step)

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

## 5. Execution
Please start by initializing the Go module and creating the **Phase 1 (Parser)** implementation. Write a test case with a sample Dockerfile containing both `ENV` and `ARG` with magic comments to verify the parsing logic.
