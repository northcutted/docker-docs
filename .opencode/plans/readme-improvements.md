# README Improvements Plan

## Summary

Complete rewrite of the project README implementing all 15 identified improvements. The new README restructures the document for better onboarding flow and addresses factual gaps (missing prerequisites, Podman support, binary downloads).

## Changes

### New Sections Added
1. **Quick Start** - 3-step getting started (install prereqs, install dock-docs, run)
2. **How It Works** - Brief 2-phase workflow explanation (static analysis + dynamic analysis)
3. **Prerequisites** - Table of syft/grype/dive with install commands and links
4. **Troubleshooting** - 5 common issues with solutions

### Sections Rewritten/Expanded
5. **Title + Description** - Added tagline ("Like terraform-docs, but for Dockerfiles"), mentions HTML/JSON/Podman
6. **Installation** - Added pre-built binary download section, fixed Docker example to mount socket, added Podman note
7. **Usage** - Replaced confusing CLI example, split CLI flags into 3 logical tables (General, CLI Mode only, Template tools)
8. **GitHub Action** - Updated checkout@v3->v4, git-auto-commit@v4->v5, expanded file_pattern, linked to action.yml for full input list
9. **Samples** - Added per-sample descriptions of image tags and comparison configurations
10. **License** - Added link to LICENSE file and copyright holder

### Sections Removed
11. **"What You Get"** sub-section under Demo - redundant with Features and Quick Start

### Minor Tweaks
12. **Features** - Added Docker & Podman auto-detection bullet, added Magic Comments bullet
13. **Config Reference** - Enriched comparison example with `source`, `details: true`, per-image source override
14. **TOC** - Fully rebuilt to match new structure

### Sections Kept As-Is
- Badges/logo header
- Magic Comments
- Templates (entire section)
- Credits
- AI Disclaimer

## Estimated Size
~530 lines (currently 418)
