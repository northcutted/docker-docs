# Go Docker Analysis

<!-- BEGIN: docker-docs:config -->

# 🐳 Docker Image Analysis: golang:1.24-alpine
![Size](https://img.shields.io/static/v1?label=Size&message=258.71+MB&color=blue) ![Layers](https://img.shields.io/static/v1?label=Layers&message=5&color=blue) ![Vulns](https://img.shields.io/static/v1?label=Security&message=3+Vulns+%280+Crit%29&color=green) ![Efficiency](https://img.shields.io/static/v1?label=Efficiency&message=99.9%&color=green)

## ⚙️ Configuration
### Build Arguments
| Name | Description | Default | Required |
|------|-------------|---------|:--------:|
| `CGO_ENABLED=0` | Enable CGO for building (default false) | `` | ❌ |
### Exposed Ports
| Port | Description |
|------|-------------|
| `8080` | The default port for the Go application |
---

## 🛡️ Security & Efficiency

**Base Image:** `linux (arm64)`
**Efficiency Score:** 99.9%

### Vulnerabilities
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| 🟢 0 | 🟢 0 | 3 | 0 |

<details>
<summary><strong>👇 Expand Vulnerability Details (3 found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox` | `1.37.0-r30` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox-binsh` | `1.37.0-r30` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `ssl_client` | `1.37.0-r30` |
</details>

<details>
<summary><strong>📦 Installed Packages (39 total)</strong></summary>

| Package | Version |
|---------|---------|
| alpine-baselayout | 3.7.1-r8 |
| alpine-baselayout-data | 3.7.1-r8 |
| alpine-keys | 2.6-r0 |
| alpine-release | 3.23.3-r0 |
| apk-tools | 3.0.3-r1 |
| busybox | 1.37.0-r30 |
| busybox-binsh | 1.37.0-r30 |
| ca-certificates | 20251003-r0 |
| ca-certificates-bundle | 20251003-r0 |
| cmd/addr2line | UNKNOWN |
| cmd/asm | UNKNOWN |
| cmd/buildid | UNKNOWN |
| cmd/cgo | UNKNOWN |
| cmd/compile | UNKNOWN |
| cmd/covdata | UNKNOWN |
| cmd/cover | UNKNOWN |
| cmd/doc | UNKNOWN |
| cmd/fix | UNKNOWN |
| cmd/go | UNKNOWN |
| cmd/gofmt | UNKNOWN |
| cmd/link | UNKNOWN |
| cmd/nm | UNKNOWN |
| cmd/objdump | UNKNOWN |
| cmd/pack | UNKNOWN |
| cmd/pprof | UNKNOWN |
| cmd/preprofile | UNKNOWN |
| cmd/test2json | UNKNOWN |
| cmd/trace | UNKNOWN |
| cmd/vet | UNKNOWN |
| go | 1.24.13 |
| libapk | 3.0.3-r1 |
| libcrypto3 | 3.5.5-r0 |
| libssl3 | 3.5.5-r0 |
| musl | 1.2.5-r21 |
| musl-utils | 1.2.5-r21 |
| scanelf | 1.3.8-r2 |
| ssl_client | 1.37.0-r30 |
| stdlib | go1.24.13 |
| zlib | 1.3.1-r2 |
</details>

<!-- END: docker-docs:config -->

## Version Comparison

<!-- BEGIN: docker-docs:comparison -->

### 🏷️ Supported Tags

| Tag | Size | Vulns | Efficiency | OS/Arch |
|-----|------|-------|------------|---------|
| `golang:1.22-alpine` | ![Size]() | ![Vulns](https://img.shields.io/static/v1?label=Security&message=521+Vulns+%2842+Crit%29&color=red) | 0.0% | / |
| `golang:1.23-alpine` | ![Size]() | ![Vulns](https://img.shields.io/static/v1?label=Security&message=417+Vulns+%2823+Crit%29&color=red) | 0.0% | / |
| `golang:1.24-alpine` | ![Size](https://img.shields.io/static/v1?label=Size&message=258.71+MB&color=blue) | ![Vulns](https://img.shields.io/static/v1?label=Security&message=3+Vulns+%280+Crit%29&color=green) | 99.9% | linux/arm64 |

<!-- END: docker-docs:comparison -->
