# Security Policy

## Supported Versions

We release patches for security vulnerabilities in the following versions:

| Version | Supported          |
| ------- | ------------------ |
| latest  | :white_check_mark: |
| < 1.0   | :x:                |

**Note:** We recommend always using the latest release for the most up-to-date security patches.

## Reporting a Vulnerability

We take the security of dock-docs seriously. If you believe you have found a security vulnerability, please report it to us as described below.

### How to Report

**Please do NOT report security vulnerabilities through public GitHub issues.**

Instead, please report them via one of the following methods:

1. **GitHub Security Advisory (Preferred)**
   - Navigate to the [Security Advisories](https://github.com/northcutted/dock-docs/security/advisories) page
   - Click "Report a vulnerability"
   - Fill out the form with details about the vulnerability

2. **Email**
   - Send an email to the repository maintainers
   - Include the word "SECURITY" in the subject line
   - Provide detailed information about the vulnerability

### What to Include

Please include the following information in your report:

- **Type of vulnerability** (e.g., command injection, arbitrary file access, etc.)
- **Full paths of source file(s)** related to the vulnerability
- **Location of the affected source code** (tag/branch/commit or direct URL)
- **Step-by-step instructions** to reproduce the issue
- **Proof-of-concept or exploit code** (if possible)
- **Impact of the vulnerability** - what can an attacker do?
- **Any suggested fixes** (optional but appreciated)

### What to Expect

- **Acknowledgment**: We will acknowledge receipt of your vulnerability report within 48 hours
- **Updates**: We will send you regular updates about our progress (at least every 7 days)
- **Timeline**: We aim to patch critical vulnerabilities within 7-14 days
- **Disclosure**: We will work with you to understand the vulnerability and coordinate disclosure
- **Credit**: If you wish, we will publicly acknowledge your responsible disclosure

## Security Considerations

### Docker Socket Access

Dock-docs requires access to the Docker/Podman socket to analyze images. Be aware that:

- Socket access grants significant privileges
- Only run dock-docs with trusted Dockerfiles and images
- Consider using read-only bind mounts when possible
- Review generated documentation before committing to your repository

### External Tool Dependencies

Dock-docs executes external CLI tools (syft, grype, dive):

- These tools are invoked as subprocesses
- Ensure you download these tools from official sources
- Keep these tools updated to their latest versions
- Review the security policies of these upstream projects:
  - [Syft Security](https://github.com/anchore/syft/security)
  - [Grype Security](https://github.com/anchore/grype/security)
  - [Dive](https://github.com/wagoodman/dive)

### CI/CD Usage

When using dock-docs in CI/CD pipelines:

- Use pinned versions (not `latest` tag) for reproducibility
- Review the [action.yml](action.yml) to understand what the GitHub Action does
- Use GitHub Actions security features (e.g., `permissions:` blocks)
- Consider using Docker Content Trust for image verification
- Implement branch protection rules to prevent unauthorized changes

### Configuration Files

When using `dock-docs.yaml`:

- Store configuration in version control
- Review changes to configuration in pull requests
- Avoid putting sensitive information in config files
- Use environment variables for secrets (though dock-docs doesn't currently use secrets)

## Security Updates

We publish security updates through:

- **GitHub Releases** - with detailed changelogs
- **GitHub Security Advisories** - for critical vulnerabilities
- **CVE Database** - when applicable

To stay informed:
- Watch the repository for releases
- Subscribe to GitHub Security Advisories
- Review the [CHANGELOG](https://github.com/northcutted/dock-docs/releases)

## Known Limitations

Current security considerations and limitations:

1. **Command Injection**: Dock-docs executes external commands. We sanitize inputs, but always review Dockerfiles from untrusted sources
2. **File System Access**: The tool reads Dockerfiles and writes documentation. Ensure proper file permissions
3. **Network Access**: External tools (syft, grype) may make network requests to fetch vulnerability databases

## Security Best Practices

When using dock-docs:

1. **Use specific image tags** instead of `latest` in your `dock-docs.yaml`
2. **Review generated documentation** before committing to ensure no sensitive data is exposed
3. **Run in isolated environments** when analyzing untrusted images
4. **Keep dock-docs updated** to the latest version
5. **Audit external tool versions** and update them regularly
6. **Use minimal privilege** - don't run dock-docs as root unless necessary

## Hall of Fame

We appreciate security researchers who help keep dock-docs safe. Responsible disclosures will be acknowledged here:

*No security vulnerabilities have been publicly disclosed yet.*

---

**Last Updated**: February 2026

For general questions about security, please open a discussion in the [Discussions](https://github.com/northcutted/dock-docs/discussions) tab.
