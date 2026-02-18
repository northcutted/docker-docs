#!/usr/bin/env bash

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${BLUE}=== dock-docs Template Test Cleanup Script ===${NC}\n"

# Clean up /tmp files
echo -e "${YELLOW}[1/2] Cleaning up /tmp files...${NC}"
if ls /tmp/dock-docs-*.tmpl 1> /dev/null 2>&1; then
  rm -f /tmp/dock-docs-*.tmpl
  echo -e "${GREEN}✓ Removed exported templates from /tmp${NC}"
else
  echo "No templates found in /tmp"
fi
echo ""

# Clean up samples/python directory
echo -e "${YELLOW}[2/2] Cleaning up samples/python test files...${NC}"
cd samples/python

FILES_TO_CLEAN=(
  "output_default.md"
  "output_minimal.md"
  "output_detailed.md"
  "output_compact.md"
  "output_html.html"
  "output_json.json"
  "output_custom.md"
  "custom_template.tmpl"
  "README-TEST.md"
  "README-MATRIX.md"
  "README-OVERRIDE.md"
  "README-CLI.md"
  "dock-docs-test.yaml"
  "dock-docs-matrix-test.yaml"
  "dock-docs-section-override.yaml"
  "dock-docs-cli-override.yaml"
)

CLEANED=0
for file in "${FILES_TO_CLEAN[@]}"; do
  if [ -f "$file" ]; then
    rm "$file"
    echo "  ✓ Removed $file"
    CLEANED=$((CLEANED + 1))
  fi
done

# Restore original README if backup exists
if [ -f "README.md.backup" ]; then
  mv README.md.backup README.md
  echo "  ✓ Restored README.md from backup"
  CLEANED=$((CLEANED + 1))
fi

cd ../..

if [ $CLEANED -eq 0 ]; then
  echo -e "${YELLOW}No test files found to clean${NC}"
else
  echo -e "${GREEN}✓ Cleaned up $CLEANED file(s) from samples/python${NC}"
fi

echo -e "\n${GREEN}✓ Cleanup complete!${NC}"
