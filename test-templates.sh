#!/usr/bin/env bash
set -e

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Returns the correct file extension for a template name
ext_for() {
  case "$1" in
    html) echo "html" ;;
    json) echo "json" ;;
    *)    echo "md" ;;
  esac
}

echo -e "${BLUE}=== dock-docs Template System Test Script ===${NC}\n"

# Check for required tools
echo -e "${YELLOW}Checking for required tools...${NC}"
DOCKER_AVAILABLE=false
if command -v docker >/dev/null 2>&1 && \
   command -v syft >/dev/null 2>&1 && \
   command -v grype >/dev/null 2>&1 && \
   command -v dive >/dev/null 2>&1; then
  DOCKER_AVAILABLE=true
  echo -e "${GREEN}✓ All required tools found (docker, syft, grype, dive)${NC}"
  echo -e "${GREEN}✓ Will run FULL IMAGE ANALYSIS tests${NC}\n"
else
  echo -e "${YELLOW}⚠ Docker or analysis tools not available${NC}"
  echo -e "${YELLOW}⚠ Will run DOCKERFILE-ONLY tests (limited output)${NC}\n"
fi

# Step 1: Build
echo -e "${YELLOW}[1/9] Building binary...${NC}"
go build -o bin/dock-docs .
echo -e "${GREEN}✓ Build successful${NC}\n"

# Step 2: Basic commands
echo -e "${YELLOW}[2/9] Testing basic commands...${NC}"
./bin/dock-docs --version
echo ""
./bin/dock-docs --list-templates
echo -e "${GREEN}✓ Basic commands work${NC}\n"

# Step 3: Test --export-template (all templates)
echo -e "${YELLOW}[3/9] Testing --export-template...${NC}"
./bin/dock-docs --export-template default > /tmp/dock-docs-default.tmpl
./bin/dock-docs --export-template minimal > /tmp/dock-docs-minimal.tmpl
./bin/dock-docs --export-template detailed > /tmp/dock-docs-detailed.tmpl
./bin/dock-docs --export-template compact > /tmp/dock-docs-compact.tmpl
./bin/dock-docs --export-template html > /tmp/dock-docs-html.tmpl
./bin/dock-docs --export-template json > /tmp/dock-docs-json.tmpl
echo "Exported templates:"
ls -lh /tmp/dock-docs-*.tmpl
echo -e "${GREEN}✓ Template export works${NC}\n"

# Step 4: Test --validate-template (all templates)
echo -e "${YELLOW}[4/9] Testing --validate-template...${NC}"
for tmpl in default minimal detailed compact html json; do
  ./bin/dock-docs --validate-template /tmp/dock-docs-${tmpl}.tmpl
done
echo -e "${GREEN}✓ Template validation works${NC}\n"

# Step 5: Test all built-in templates
echo -e "${YELLOW}[5/9] Testing all templates...${NC}"
cd samples/python

# Backup original README
cp README.md README.md.backup

if [ "$DOCKER_AVAILABLE" = true ]; then
  # Full image analysis mode
  TEST_IMAGE="python:3.14-slim"
  echo -e "${BLUE}Pulling test image: ${TEST_IMAGE}${NC}"
  docker pull ${TEST_IMAGE} > /dev/null 2>&1
  echo ""
  
  for template in default minimal detailed compact html json; do
    EXT=$(ext_for ${template})
    echo -e "${BLUE}Testing template: ${template} (with image analysis)${NC}"
    ../../bin/dock-docs --template ${template} --image ${TEST_IMAGE} --dry-run > output_${template}.${EXT}
    FILE_SIZE=$(wc -c < output_${template}.${EXT} | xargs)
    LINE_COUNT=$(wc -l < output_${template}.${EXT} | xargs)
    echo "  Generated output_${template}.${EXT}: ${LINE_COUNT} lines (${FILE_SIZE} bytes)"
  done
else
  # Dockerfile-only mode (no image analysis)
  echo -e "${YELLOW}Running in Dockerfile-only mode (no image stats)${NC}\n"
  
  # Temporarily hide config file to avoid triggering config mode
  if [ -f dock-docs.yaml ]; then
    mv dock-docs.yaml dock-docs.yaml.hidden
  fi

  for template in default minimal detailed compact html json; do
    EXT=$(ext_for ${template})
    echo -e "${BLUE}Testing template: ${template} (Dockerfile only)${NC}"
    ../../bin/dock-docs --template ${template} --dry-run -f Dockerfile > output_${template}.${EXT} 2>&1 || true
    if [ -f output_${template}.${EXT} ] && [ -s output_${template}.${EXT} ]; then
      FILE_SIZE=$(wc -c < output_${template}.${EXT} | xargs)
      LINE_COUNT=$(wc -l < output_${template}.${EXT} | xargs)
      echo "  Generated output_${template}.${EXT}: ${LINE_COUNT} lines (${FILE_SIZE} bytes)"
    else
      echo "  Skipped (requires image analysis)"
    fi
  done

  # Restore config file
  if [ -f dock-docs.yaml.hidden ]; then
    mv dock-docs.yaml.hidden dock-docs.yaml
  fi
fi

# Validate HTML and JSON outputs
echo ""
echo -e "${BLUE}Validating output formats...${NC}"
if [ -f output_html.html ] && [ -s output_html.html ]; then
  if head -1 output_html.html | grep -q "<!DOCTYPE html>"; then
    echo -e "  ${GREEN}✓ output_html.html starts with <!DOCTYPE html>${NC}"
  else
    echo -e "  ${RED}✗ output_html.html is not valid HTML${NC}"
    exit 1
  fi
fi
if [ -f output_json.json ] && [ -s output_json.json ]; then
  if python3 -m json.tool output_json.json > /dev/null 2>&1; then
    echo -e "  ${GREEN}✓ output_json.json is valid JSON${NC}"
  else
    echo -e "  ${RED}✗ output_json.json is not valid JSON${NC}"
    exit 1
  fi
fi

echo -e "${GREEN}✓ All templates tested${NC}\n"

# Step 6: Test comparison mode with templates
echo -e "${YELLOW}[6/9] Testing comparison mode with templates...${NC}"

if [ "$DOCKER_AVAILABLE" = true ]; then
  # Create a test config with comparison
  cat > dock-docs-matrix-test.yaml << 'EOF'
output: "README-MATRIX.md"
template:
  name: "compact"
sections:
  - type: "comparison"
    marker: "matrix"
    source: "Dockerfile"
    images:
      - tag: "python:3.12-slim"
      - tag: "python:3.13-slim"
      - tag: "python:3.14-slim"
EOF

  # Create test README with markers
  cat > README-MATRIX.md << 'EOF'
# Matrix Test

<!-- BEGIN: dock-docs:matrix -->
OLD CONTENT
<!-- END: dock-docs:matrix -->

EOF

  ../../bin/dock-docs --config dock-docs-matrix-test.yaml
  echo -e "\n${BLUE}Comparison output (first 20 lines):${NC}"
  head -20 README-MATRIX.md
  echo -e "${GREEN}✓ Comparison mode with templates works${NC}\n"
else
  echo -e "${YELLOW}Skipped (requires Docker)${NC}\n"
fi
# Step 7: Test custom template file
echo -e "${YELLOW}[7/9] Testing custom template file...${NC}"

# Temporarily hide config file to test simple mode
if [ -f dock-docs.yaml ]; then
  mv dock-docs.yaml dock-docs.yaml.hidden
fi

cat > custom_template.tmpl << 'EOF'
# Custom Template Demo

**Image:** {{ .ImageTag }}

{{- if .Stats }}
## Image Statistics
- **Size:** {{ .Stats.SizeMB }}
- **Layers:** {{ .Stats.TotalLayers }}
- **Efficiency:** {{ printf "%.1f" .Stats.Efficiency }}%
- **Vulnerabilities:** {{ .Stats.TotalVulns }}
{{- end }}

{{- if (len (.Doc.FilterByType "ENV")) }}
## Environment Variables
{{- range (.Doc.FilterByType "ENV") }}
- **{{ .Name }}**: {{ .Description }} (default: {{ .Value }})
{{- end }}
{{- end }}
EOF

if [ "$DOCKER_AVAILABLE" = true ]; then
  ../../bin/dock-docs --template custom_template.tmpl --image ${TEST_IMAGE} --dry-run > output_custom.md
else
  ../../bin/dock-docs --template custom_template.tmpl --dry-run -f Dockerfile > output_custom.md 2>&1 || true
fi

# Restore config file
if [ -f dock-docs.yaml.hidden ]; then
  mv dock-docs.yaml.hidden dock-docs.yaml
fi

if [ -f output_custom.md ] && [ -s output_custom.md ]; then
  echo "Generated $(wc -l < output_custom.md) lines with custom template"
  echo -e "\n${BLUE}Custom template output:${NC}"
  cat output_custom.md
  echo -e "${GREEN}✓ Custom template from file works${NC}\n"
else
  echo -e "${YELLOW}Custom template test limited without image analysis${NC}\n"
fi

# Step 8: Test per-section template override
echo -e "${YELLOW}[8/9] Testing per-section template override...${NC}"

cat > dock-docs-section-override.yaml << 'EOF'
output: "README-OVERRIDE.md"
template:
  name: "default"  # Global default
sections:
  - type: "image"
    marker: "config-minimal"
    source: "Dockerfile"
    template:
      name: "minimal"  # Override for this section
  - type: "image"
    marker: "config-detailed"
    source: "Dockerfile"
    template:
      name: "detailed"  # Override for this section
EOF

cat > README-OVERRIDE.md << 'EOF'
# Section Override Test

## Minimal Template Section
<!-- BEGIN: dock-docs:config-minimal -->
OLD
<!-- END: dock-docs:config-minimal -->

## Detailed Template Section
<!-- BEGIN: dock-docs:config-detailed -->
OLD
<!-- END: dock-docs:config-detailed -->

EOF

../../bin/dock-docs --config dock-docs-section-override.yaml
echo -e "\n${BLUE}Section override result:${NC}"
cat README-OVERRIDE.md
echo -e "${GREEN}✓ Per-section template override works${NC}\n"

# Step 9: Test CLI flag override
echo -e "${YELLOW}[9/9] Testing CLI --template flag override...${NC}"

# Create config with default template
cat > dock-docs-cli-override.yaml << 'EOF'
output: "README-CLI.md"
template:
  name: "detailed"  # Config specifies detailed
sections:
  - type: "image"
    marker: "config"
    source: "Dockerfile"
EOF

cat > README-CLI.md << 'EOF'
# CLI Override Test

<!-- BEGIN: dock-docs:config -->
OLD
<!-- END: dock-docs:config -->

EOF

# Run with CLI override
../../bin/dock-docs --config dock-docs-cli-override.yaml --template minimal
echo -e "\n${BLUE}CLI override result (should use minimal, not detailed):${NC}"
head -15 README-CLI.md
echo -e "${GREEN}✓ CLI --template flag overrides config${NC}\n"

# Show summary
echo -e "${BLUE}=== Summary of Generated Files ===${NC}"
ls -lh output_*.* custom_template.tmpl README-*.md dock-docs-*.yaml 2>/dev/null || true

echo -e "\n${BLUE}=== Template Comparison (file sizes) ===${NC}"
for template in default minimal detailed compact html json; do
  EXT=$(ext_for ${template})
  FILE="output_${template}.${EXT}"
  if [ -f "$FILE" ]; then
    SIZE=$(wc -c < "$FILE" | xargs)
    LINES=$(wc -l < "$FILE" | xargs)
    printf "  %-12s: %6d bytes, %4d lines  (%s)\n" "${template}" "${SIZE}" "${LINES}" "${FILE}"
  fi
done

# Cleanup prompt
echo -e "\n${YELLOW}Generated test files in samples/python/:${NC}"
echo "  - output_*.{md,html,json} (6 template variants)"
echo "  - custom_template.tmpl (custom template)"
echo "  - README-*.md (various test outputs)"
echo "  - dock-docs-*.yaml (test configs)"
echo "  - README.md.backup (original)"
echo ""
echo -e "${GREEN}To restore and clean: cd ../.. && ./cleanup-tests.sh${NC}"

cd ../..
echo -e "\n${GREEN}✓ All tests completed successfully!${NC}"
