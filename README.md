# Docker Documentation Generator

A powerful CLI tool that automatically generates comprehensive Markdown documentation from your Dockerfiles. It goes beyond static analysis by building your image and running deep inspections using industry-standard tools like `syft`, `grype`, and `dive`.

## Features

- 📝 **Automatic Documentation**: Parses standard Dockerfile instructions (`FROM`, `ENV`, `RUN`, `EXPOSE`, etc.) into clean Markdown.
- 🔍 **Deep Analysis**:
  - **Syft**: Generates a Software Bill of Materials (SBOM) to list all installed packages.
  - **Grype**: Scans the image for vulnerabilities.
  - **Dive**: Analyzes layer efficiency and wasted space.
- 🏗️ **Build & Inspect**: Automatically builds the container image to perform dynamic analysis.
- 🧩 **Matrix Support**: Scan multiple Dockerfiles in a single run (e.g., `dev.Dockerfile`, `prod.Dockerfile`).
- 🎨 **Customizable Templates**: Uses Go templates for flexible output formatting.

## Installation

### Go Install

```bash
go install github.com/northcutted/docker-docs@latest
```

### Docker

```bash
docker run --rm -v $(pwd):/app -w /app ghcr.io/northcutted/docker-docs:latest
```

## Usage

### Basic Usage

Run the tool in your project directory:

```bash
docker-docs
```

This will look for a `docker-docs.yaml` configuration file by default.

### CLI Arguments

```bash
docker-docs --config my-config.yaml --output README.md
```

- `--config`, `-c`: Path to configuration file (default: `docker-docs.yaml`). If found, this enables **Config Mode**.
- `--output`, `-o`: Output file path (default: `README.md`).
- `--file`, `-f`: Path to Dockerfile (Legacy Mode only).
- `--image`: Docker image tag to analyze (Legacy Mode only).
- `--dry-run`: Print output to stdout instead of modifying files.

## Configuration Reference (`docker-docs.yaml`)

The `docker-docs.yaml` file is the heart of the tool. It allows you to define multiple sections of documentation that will be injected into your output file (e.g., `README.md`).

### Structure

```yaml
# The file where documentation will be injected (default: README.md)
output: "README.md"

# A list of documentation sections to generate
sections:
  - type: "config"      # Generates docs for a single Dockerfile
    marker: "main"      # Maps to <!-- BEGIN: docker-docs:main -->
    source: "Dockerfile" # Path to the Dockerfile
    image: "myapp:latest" # (Optional) Image to analyze for deep inspection

  - type: "matrix"      # Generates a comparison table for multiple images
    marker: "comparison" # Maps to <!-- BEGIN: docker-docs:comparison -->
    images:             # List of images to compare
      - "myapp:dev"
      - "myapp:prod"
      - "alpine:latest"
```

### Markers

To tell `docker-docs` where to insert the generated content, you must add **HTML comments** (markers) to your target file (e.g., `README.md`).

For a section with `marker: "main"`, add this to your README:

```markdown
<!-- BEGIN: docker-docs:main -->
<!-- END: docker-docs:main -->
```

Everything between these two lines will be overwritten by the tool.

### Section Types

#### 1. `config`
Parses a Dockerfile and optionally analyzes a built image.

- **`marker`** (Required): unique string to identify the injection point.
- **`source`** (Optional): Path to the `Dockerfile`. Defaults to `Dockerfile`.
- **`image`** (Optional): If provided, the tool will pull/build and analyze this image using Syft, Grype, and Dive.

#### 2. `matrix`
Generates a comparison table for multiple images, showing size, vulnerability counts, and efficiency scores.

- **`marker`** (Required): Unique string to identify the injection point.
- **`images`** (Required): A list of image tags to analyze and compare.

## GitHub Action

You can easily integrate this tool into your CI/CD pipeline using the GitHub Action:

```yaml
name: Generate Docker Docs
on:
  push:
    paths:
      - 'Dockerfile'
      - 'docker-docs.yaml'

jobs:
  docs:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Generate Documentation
        uses: northcutted/docker-docs@main
        with:
          config: 'docker-docs.yaml'
          output: 'README.md'
      
      - name: Commit changes
        uses: stefanzweifel/git-auto-commit-action@v4
        with:
          commit_message: "docs: update docker documentation"
          file_pattern: README.md
```

## Legacy Mode

If no `docker-docs.yaml` is found, the tool runs in **Legacy Mode**:

1. Parses `./Dockerfile` (or file specified by `-f`).
2. (Optional) Analyzes the image specified by `--image`.
3. Injects the output into `README.md` using the **default markers**:

```markdown
<!-- BEGIN: docker-docs -->
<!-- END: docker-docs -->
```

## License

MIT
