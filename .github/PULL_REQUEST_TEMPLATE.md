## Description

<!-- Provide a clear and concise description of your changes -->

## Type of Change

<!-- Mark the relevant option with an 'x' -->

- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] Documentation update
- [ ] Code refactoring (no functional changes)
- [ ] Test improvements
- [ ] CI/CD improvements
- [ ] Other (please describe):

## Related Issues

<!-- Link to related issues using keywords like "Fixes #123" or "Relates to #456" -->

- Fixes #
- Relates to #

## Changes Made

<!-- Provide a detailed list of changes -->

-
-
-

## Testing

<!-- Describe how you tested your changes -->

### Test Commands

```bash
# Commands you ran to test
go test ./...
go build -o dock-docs .
./dock-docs --help
```

### Test Results

- [ ] All existing tests pass
- [ ] Added new tests for new functionality
- [ ] Tested manually with example Dockerfiles
- [ ] Tested in CLI Mode
- [ ] Tested in YAML Mode
- [ ] Tested with `--verbose` flag
- [ ] No race conditions (`go test -race ./...`)

### Test Environment

- **OS**: (e.g., Ubuntu 22.04, macOS 14.1)
- **Go version**: (run `go version`)
- **Docker/Podman version**:

## Documentation

<!-- Check all that apply -->

- [ ] Updated README.md (if needed)
- [ ] Updated CONTRIBUTING.md (if needed)
- [ ] Added/updated code comments
- [ ] Added/updated examples in `samples/`
- [ ] Updated CLI help text
- [ ] No documentation changes needed

## Checklist

<!-- Ensure you've completed these steps -->

- [ ] My code follows the project's code style (see AGENTS.md)
- [ ] I have performed a self-review of my code
- [ ] I have commented my code, particularly in hard-to-understand areas
- [ ] My changes generate no new warnings or errors
- [ ] I have added tests that prove my fix is effective or that my feature works
- [ ] New and existing unit tests pass locally with my changes
- [ ] I have run `go fmt ./...` to format my code
- [ ] I have run `go mod tidy` to clean up dependencies
- [ ] My commit messages follow [Conventional Commits](https://www.conventionalcommits.org/)
- [ ] I have updated the documentation accordingly

## Breaking Changes

<!-- If this is a breaking change, describe the impact and migration path -->

**Impact**:
<!-- What will break? -->

**Migration Path**:
<!-- How should users update their code/config? -->

## Screenshots / Examples

<!-- If applicable, add screenshots or example output -->

### Before

```
Paste output or screenshot
```

### After

```
Paste output or screenshot
```

## Performance Impact

<!-- Does this change affect performance? -->

- [ ] No performance impact
- [ ] Improves performance
- [ ] May impact performance (please explain below)

**Details**:

## Additional Context

<!-- Add any other context about the pull request here -->

## Reviewer Notes

<!-- Optional: specific areas you'd like reviewers to focus on -->

---

**For Maintainers**:
- [ ] Squash commits before merging
- [ ] Update CHANGELOG (if using one)
- [ ] Add release notes (if needed)
