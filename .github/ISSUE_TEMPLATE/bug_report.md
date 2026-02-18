---
name: Bug Report
about: Report a bug to help us improve dock-docs
title: '[BUG] '
labels: bug
assignees: ''
---

## Bug Description

A clear and concise description of what the bug is.

## Steps to Reproduce

1. Run command: `dock-docs ...`
2. With config file: `...`
3. Using Dockerfile: `...`
4. See error

## Expected Behavior

A clear description of what you expected to happen.

## Actual Behavior

What actually happened. Include error messages, stack traces, or unexpected output.

```
Paste error output here
```

## Environment

- **dock-docs version**: (run `dock-docs version`)
- **OS**: (e.g., Ubuntu 22.04, macOS 14.1, Windows 11)
- **Architecture**: (e.g., amd64, arm64)
- **Docker/Podman version**: (run `docker version` or `podman version`)
- **External tools installed**:
  - [ ] syft (version: )
  - [ ] grype (version: )
  - [ ] dive (version: )

## Dockerfile

If applicable, provide a minimal Dockerfile that reproduces the issue:

```dockerfile
# Paste Dockerfile here
```

## Configuration File

If using YAML mode, provide your `dock-docs.yaml`:

```yaml
# Paste config here (remove any sensitive data)
```

## Logs

If you ran with `--verbose`, paste the verbose output:

```
Paste verbose output here
```

## Additional Context

Add any other context about the problem here:
- Screenshots
- Links to similar issues
- Workarounds you've tried
- Related documentation

## Possible Solution

If you have ideas on how to fix this, please share!
