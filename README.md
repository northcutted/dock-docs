# Dock-Docks - A powerful automated documentaion generator for docker images.
<p align="center">
  <img src="./assets/dock-docs-logo.png" alt="Logo Alt Text" >
</p>

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

## Magic Comments

You can enhance the generated documentation by adding special comments directly in your `Dockerfile`. These comments allow you to provide descriptions, default values, and requirement status for `ARG`, `ENV`, and `EXPOSE` instructions.

Supported tags:
- `@description:` A description of the variable or instruction.
- `@default:` The default value (overrides the value in the instruction).
- `@required:` (true/false) Whether the variable is mandatory.

### Examples

**Documenting an Environment Variable:**
```dockerfile
# @description: The API key for the service
# @required: true
ENV API_KEY=""
```

**Documenting a Build Argument:**
```dockerfile
# @description: The version of the base image
# @default: latest
ARG VERSION=latest
```

**Documenting an Exposed Port:**
```dockerfile
# @description: The main HTTP port
EXPOSE 8080
```

## Installation

### Go Install

```bash
go install github.com/northcutted/dock-docs@latest
```

### Docker

```bash
docker run --rm -v $(pwd):/app -w /app ghcr.io/northcutted/dock-docs:latest
```

## Usage

### Basic Usage

Run the tool in your project directory:

```bash
dock-docs
```

This will look for a `dock-docs.yaml` configuration file by default. If no `dock-docs.yaml` is found, the tool runs in **Simple Mode**.


#### Simple Mode

If no `dock-docs.yaml` is found, the tool runs in **Simple Mode**. This mode is useful for documentation generation without creating a configuration file, but it offers less customization.

1. Parses `./Dockerfile` (or file specified by `-f`).
2. (Optional) Analyzes the image specified by `--image`.
3. Injects the output into `README.md` using the **default markers**:

```markdown
<!-- BEGIN: dock-docs -->

# 🐳 Docker Image Analysis: Dockerfile

## ⚙️ Configuration
### Environment Variables
| Name | Description | Default | Required |
|------|-------------|---------|:--------:|
| `PATH` |  | `/usr/local/bin:${PATH}` | ❌ |
### Build Arguments
| Name | Description | Default | Required |
|------|-------------|---------|:--------:|
| `TARGETARCH` |  | `` | ❌ |
| `DIVE_VERSION=0.13.1` |  | `` | ❌ |
| `DOCKER_VERSION=29.2.1` |  | `` | ❌ |
| `TARGETPLATFORM` |  | `` | ❌ |

<!-- END: dock-docs -->
```

### CLI Arguments

```bash
dock-docs --config my-config.yaml --output README.md
```

- `--config`, `-c`: Path to configuration file (default: `dock-docs.yaml`). If found, this enables **Config Mode**.
- `--output`, `-o`: Output file path (default: `README.md`).
- `--dry-run`: Print output to stdout instead of modifying files.
- `--nomoji`: Disable emojis in the output.
- `--ignore-errors`: Ignore analysis errors and continue generation.
- `--verbose`: Enable verbose logging for debugging.
- `--file`, `-f`: Path to Dockerfile (Simple Mode only).
- `--image`: Docker image tag to analyze (Simple Mode only).

## Configuration Reference (`dock-docs.yaml`)

The `dock-docs.yaml` file is the heart of the tool. It allows you to define multiple sections of documentation that will be injected into your output file (e.g., `README.md`).

### Structure

```yaml
# The file where documentation will be injected (default: README.md)
output: "README.md"

# A list of documentation sections to generate
sections:
  - type: "config"      # Generates docs for a single Dockerfile
    marker: "main"      # Maps to <!-- BEGIN: dock-docs:main -->
    source: "Dockerfile" # Path to the Dockerfile
    image: "myapp:latest" # (Optional) Image to analyze for deep inspection

  - type: "matrix"      # Generates a comparison table for multiple images
    marker: "comparison" # Maps to <!-- BEGIN: dock-docs:comparison -->
    images:             # List of images to compare
      - "myapp:dev"
      - "myapp:prod"
      - "alpine:latest"
```

### Markers

To tell `dock-docs` where to insert the generated content, you must add **HTML comments** (markers) to your target file (e.g., `README.md`).

For a section with `marker: "main"`, add this to your `README`:

```markdown
<!-- BEGIN: dock-docs:main -->
<!-- END: dock-docs:main -->
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
      - 'dock-docs.yaml'

jobs:
  docs:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Generate Documentation
        uses: northcutted/dock-docs@main
        with:
          config: 'dock-docs.yaml'
          output: 'README.md'
      
      - name: Commit changes
        uses: stefanzweifel/git-auto-commit-action@v4
        with:
          commit_message: "docs: update docker documentation"
          file_pattern: README.md
```

## Samples

Check out the sample projects to see `dock-docs` in action:

- [Go Sample](samples/go)
- [Node.js Sample](samples/node)
- [Python Sample](samples/python)

## Credits

This tool relies on the following amazing open-source projects for deep analysis:

- [Syft](https://github.com/anchore/syft) by Anchore (SBOM generation)
- [Grype](https://github.com/anchore/grype) by Anchore (Vulnerability scanning)
- [Dive](https://github.com/wagoodman/dive) by Alex Goodman (Layer efficiency analysis)

## License

MIT
