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
**Supported Architectures:** `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown`
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

| Tag | Size | Vulns | Efficiency | Architectures |
|-----|------|-------|------------|---------------|
| `golang:1.22-alpine` | ![Size](https://img.shields.io/static/v1?label=Size&message=229.07+MB&color=blue) | ![Vulns](https://img.shields.io/static/v1?label=Security&message=521+Vulns+%2842+Crit%29&color=red) | 99.9% | `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown` |
| `golang:1.23-alpine` | ![Size](https://img.shields.io/static/v1?label=Size&message=243.88+MB&color=blue) | ![Vulns](https://img.shields.io/static/v1?label=Security&message=417+Vulns+%2823+Crit%29&color=red) | 99.9% | `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown` |
| `golang:1.24-alpine` | ![Size](https://img.shields.io/static/v1?label=Size&message=258.71+MB&color=blue) | ![Vulns](https://img.shields.io/static/v1?label=Security&message=3+Vulns+%280+Crit%29&color=green) | 99.9% | `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown` |

<details>
<summary><strong>🔍 Full Report: golang:1.22-alpine</strong></summary>

## 🛡️ Security & Efficiency

**Base Image:** `linux (arm64)`
**Supported Architectures:** `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown`
**Efficiency Score:** 99.9%

### Vulnerabilities
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| 🔴 42 | 🟠 210 | 243 | 26 |

<details>
<summary><strong>👇 Expand Vulnerability Details (521 found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [CVE-2025-15467](https://nvd.nist.gov/vuln/detail/CVE-2025-15467) | Critical | `libcrypto3` | `3.3.2-r4` |
| [CVE-2025-15467](https://nvd.nist.gov/vuln/detail/CVE-2025-15467) | Critical | `libssl3` | `3.3.2-r4` |
| [CVE-2025-22871](https://nvd.nist.gov/vuln/detail/CVE-2025-22871) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-22871](https://nvd.nist.gov/vuln/detail/CVE-2025-22871) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-22871](https://nvd.nist.gov/vuln/detail/CVE-2025-22871) | Critical | `go` | `1.22.12` |
| [CVE-2025-22871](https://nvd.nist.gov/vuln/detail/CVE-2025-22871) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-22871](https://nvd.nist.gov/vuln/detail/CVE-2025-22871) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-22871](https://nvd.nist.gov/vuln/detail/CVE-2025-22871) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-22871](https://nvd.nist.gov/vuln/detail/CVE-2025-22871) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-22871](https://nvd.nist.gov/vuln/detail/CVE-2025-22871) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-22871](https://nvd.nist.gov/vuln/detail/CVE-2025-22871) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-22871](https://nvd.nist.gov/vuln/detail/CVE-2025-22871) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-22871](https://nvd.nist.gov/vuln/detail/CVE-2025-22871) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-22871](https://nvd.nist.gov/vuln/detail/CVE-2025-22871) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-22871](https://nvd.nist.gov/vuln/detail/CVE-2025-22871) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-22871](https://nvd.nist.gov/vuln/detail/CVE-2025-22871) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-22871](https://nvd.nist.gov/vuln/detail/CVE-2025-22871) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-22871](https://nvd.nist.gov/vuln/detail/CVE-2025-22871) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-22871](https://nvd.nist.gov/vuln/detail/CVE-2025-22871) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-22871](https://nvd.nist.gov/vuln/detail/CVE-2025-22871) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-22871](https://nvd.nist.gov/vuln/detail/CVE-2025-22871) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-22871](https://nvd.nist.gov/vuln/detail/CVE-2025-22871) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `go` | `1.22.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.22.12` |
| [CVE-2025-26519](https://nvd.nist.gov/vuln/detail/CVE-2025-26519) | High | `musl-utils` | `1.2.5-r8` |
| [CVE-2025-26519](https://nvd.nist.gov/vuln/detail/CVE-2025-26519) | High | `musl` | `1.2.5-r8` |
| [CVE-2025-4674](https://nvd.nist.gov/vuln/detail/CVE-2025-4674) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-4674](https://nvd.nist.gov/vuln/detail/CVE-2025-4674) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-4674](https://nvd.nist.gov/vuln/detail/CVE-2025-4674) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-4674](https://nvd.nist.gov/vuln/detail/CVE-2025-4674) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-4674](https://nvd.nist.gov/vuln/detail/CVE-2025-4674) | High | `go` | `1.22.12` |
| [CVE-2025-4674](https://nvd.nist.gov/vuln/detail/CVE-2025-4674) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-4674](https://nvd.nist.gov/vuln/detail/CVE-2025-4674) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-4674](https://nvd.nist.gov/vuln/detail/CVE-2025-4674) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-4674](https://nvd.nist.gov/vuln/detail/CVE-2025-4674) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-4674](https://nvd.nist.gov/vuln/detail/CVE-2025-4674) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-4674](https://nvd.nist.gov/vuln/detail/CVE-2025-4674) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-4674](https://nvd.nist.gov/vuln/detail/CVE-2025-4674) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-4674](https://nvd.nist.gov/vuln/detail/CVE-2025-4674) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-4674](https://nvd.nist.gov/vuln/detail/CVE-2025-4674) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-4674](https://nvd.nist.gov/vuln/detail/CVE-2025-4674) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-4674](https://nvd.nist.gov/vuln/detail/CVE-2025-4674) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-4674](https://nvd.nist.gov/vuln/detail/CVE-2025-4674) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-4674](https://nvd.nist.gov/vuln/detail/CVE-2025-4674) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-4674](https://nvd.nist.gov/vuln/detail/CVE-2025-4674) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-4674](https://nvd.nist.gov/vuln/detail/CVE-2025-4674) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-47907](https://nvd.nist.gov/vuln/detail/CVE-2025-47907) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-47907](https://nvd.nist.gov/vuln/detail/CVE-2025-47907) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-47907](https://nvd.nist.gov/vuln/detail/CVE-2025-47907) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-47907](https://nvd.nist.gov/vuln/detail/CVE-2025-47907) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-47907](https://nvd.nist.gov/vuln/detail/CVE-2025-47907) | High | `go` | `1.22.12` |
| [CVE-2025-47907](https://nvd.nist.gov/vuln/detail/CVE-2025-47907) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-47907](https://nvd.nist.gov/vuln/detail/CVE-2025-47907) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-47907](https://nvd.nist.gov/vuln/detail/CVE-2025-47907) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-47907](https://nvd.nist.gov/vuln/detail/CVE-2025-47907) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-47907](https://nvd.nist.gov/vuln/detail/CVE-2025-47907) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-47907](https://nvd.nist.gov/vuln/detail/CVE-2025-47907) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-47907](https://nvd.nist.gov/vuln/detail/CVE-2025-47907) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-47907](https://nvd.nist.gov/vuln/detail/CVE-2025-47907) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-47907](https://nvd.nist.gov/vuln/detail/CVE-2025-47907) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-47907](https://nvd.nist.gov/vuln/detail/CVE-2025-47907) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-47907](https://nvd.nist.gov/vuln/detail/CVE-2025-47907) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-47907](https://nvd.nist.gov/vuln/detail/CVE-2025-47907) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-47907](https://nvd.nist.gov/vuln/detail/CVE-2025-47907) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-47907](https://nvd.nist.gov/vuln/detail/CVE-2025-47907) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-47907](https://nvd.nist.gov/vuln/detail/CVE-2025-47907) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `go` | `1.22.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `go` | `1.22.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `go` | `1.22.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `go` | `1.22.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `go` | `1.22.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `go` | `1.22.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `go` | `1.22.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `go` | `1.22.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.22.12` |
| [CVE-2025-69419](https://nvd.nist.gov/vuln/detail/CVE-2025-69419) | High | `libssl3` | `3.3.2-r4` |
| [CVE-2025-69419](https://nvd.nist.gov/vuln/detail/CVE-2025-69419) | High | `libcrypto3` | `3.3.2-r4` |
| [CVE-2025-69420](https://nvd.nist.gov/vuln/detail/CVE-2025-69420) | High | `libssl3` | `3.3.2-r4` |
| [CVE-2025-69420](https://nvd.nist.gov/vuln/detail/CVE-2025-69420) | High | `libcrypto3` | `3.3.2-r4` |
| [CVE-2025-69421](https://nvd.nist.gov/vuln/detail/CVE-2025-69421) | High | `libssl3` | `3.3.2-r4` |
| [CVE-2025-69421](https://nvd.nist.gov/vuln/detail/CVE-2025-69421) | High | `libcrypto3` | `3.3.2-r4` |
| [CVE-2025-9230](https://nvd.nist.gov/vuln/detail/CVE-2025-9230) | High | `libssl3` | `3.3.2-r4` |
| [CVE-2025-9230](https://nvd.nist.gov/vuln/detail/CVE-2025-9230) | High | `libcrypto3` | `3.3.2-r4` |
| [CVE-2024-12797](https://nvd.nist.gov/vuln/detail/CVE-2024-12797) | Medium | `libcrypto3` | `3.3.2-r4` |
| [CVE-2024-12797](https://nvd.nist.gov/vuln/detail/CVE-2024-12797) | Medium | `libssl3` | `3.3.2-r4` |
| [CVE-2024-13176](https://nvd.nist.gov/vuln/detail/CVE-2024-13176) | Medium | `libcrypto3` | `3.3.2-r4` |
| [CVE-2024-13176](https://nvd.nist.gov/vuln/detail/CVE-2024-13176) | Medium | `libssl3` | `3.3.2-r4` |
| [CVE-2025-15468](https://nvd.nist.gov/vuln/detail/CVE-2025-15468) | Medium | `libcrypto3` | `3.3.2-r4` |
| [CVE-2025-15468](https://nvd.nist.gov/vuln/detail/CVE-2025-15468) | Medium | `libssl3` | `3.3.2-r4` |
| [CVE-2025-4673](https://nvd.nist.gov/vuln/detail/CVE-2025-4673) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-4673](https://nvd.nist.gov/vuln/detail/CVE-2025-4673) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-4673](https://nvd.nist.gov/vuln/detail/CVE-2025-4673) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-4673](https://nvd.nist.gov/vuln/detail/CVE-2025-4673) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-4673](https://nvd.nist.gov/vuln/detail/CVE-2025-4673) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-4673](https://nvd.nist.gov/vuln/detail/CVE-2025-4673) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-4673](https://nvd.nist.gov/vuln/detail/CVE-2025-4673) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-4673](https://nvd.nist.gov/vuln/detail/CVE-2025-4673) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-4673](https://nvd.nist.gov/vuln/detail/CVE-2025-4673) | Medium | `go` | `1.22.12` |
| [CVE-2025-4673](https://nvd.nist.gov/vuln/detail/CVE-2025-4673) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-4673](https://nvd.nist.gov/vuln/detail/CVE-2025-4673) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-4673](https://nvd.nist.gov/vuln/detail/CVE-2025-4673) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-4673](https://nvd.nist.gov/vuln/detail/CVE-2025-4673) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-4673](https://nvd.nist.gov/vuln/detail/CVE-2025-4673) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-4673](https://nvd.nist.gov/vuln/detail/CVE-2025-4673) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-4673](https://nvd.nist.gov/vuln/detail/CVE-2025-4673) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-4673](https://nvd.nist.gov/vuln/detail/CVE-2025-4673) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-4673](https://nvd.nist.gov/vuln/detail/CVE-2025-4673) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-4673](https://nvd.nist.gov/vuln/detail/CVE-2025-4673) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-4673](https://nvd.nist.gov/vuln/detail/CVE-2025-4673) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47906](https://nvd.nist.gov/vuln/detail/CVE-2025-47906) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47906](https://nvd.nist.gov/vuln/detail/CVE-2025-47906) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47906](https://nvd.nist.gov/vuln/detail/CVE-2025-47906) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47906](https://nvd.nist.gov/vuln/detail/CVE-2025-47906) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47906](https://nvd.nist.gov/vuln/detail/CVE-2025-47906) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47906](https://nvd.nist.gov/vuln/detail/CVE-2025-47906) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47906](https://nvd.nist.gov/vuln/detail/CVE-2025-47906) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47906](https://nvd.nist.gov/vuln/detail/CVE-2025-47906) | Medium | `go` | `1.22.12` |
| [CVE-2025-47906](https://nvd.nist.gov/vuln/detail/CVE-2025-47906) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47906](https://nvd.nist.gov/vuln/detail/CVE-2025-47906) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47906](https://nvd.nist.gov/vuln/detail/CVE-2025-47906) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47906](https://nvd.nist.gov/vuln/detail/CVE-2025-47906) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47906](https://nvd.nist.gov/vuln/detail/CVE-2025-47906) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47906](https://nvd.nist.gov/vuln/detail/CVE-2025-47906) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47906](https://nvd.nist.gov/vuln/detail/CVE-2025-47906) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47906](https://nvd.nist.gov/vuln/detail/CVE-2025-47906) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47906](https://nvd.nist.gov/vuln/detail/CVE-2025-47906) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47906](https://nvd.nist.gov/vuln/detail/CVE-2025-47906) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47906](https://nvd.nist.gov/vuln/detail/CVE-2025-47906) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47906](https://nvd.nist.gov/vuln/detail/CVE-2025-47906) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `go` | `1.22.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `go` | `1.22.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `go` | `1.22.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `go` | `1.22.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `go` | `1.22.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `ssl_client` | `1.37.0-r9` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox-binsh` | `1.37.0-r9` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox` | `1.37.0-r9` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `go` | `1.22.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `go` | `1.22.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `go` | `1.22.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `go` | `1.22.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.22.12` |
| [CVE-2025-66199](https://nvd.nist.gov/vuln/detail/CVE-2025-66199) | Medium | `libssl3` | `3.3.2-r4` |
| [CVE-2025-66199](https://nvd.nist.gov/vuln/detail/CVE-2025-66199) | Medium | `libcrypto3` | `3.3.2-r4` |
| [CVE-2025-68160](https://nvd.nist.gov/vuln/detail/CVE-2025-68160) | Medium | `libssl3` | `3.3.2-r4` |
| [CVE-2025-68160](https://nvd.nist.gov/vuln/detail/CVE-2025-68160) | Medium | `libcrypto3` | `3.3.2-r4` |
| [CVE-2025-69418](https://nvd.nist.gov/vuln/detail/CVE-2025-69418) | Medium | `libcrypto3` | `3.3.2-r4` |
| [CVE-2025-69418](https://nvd.nist.gov/vuln/detail/CVE-2025-69418) | Medium | `libssl3` | `3.3.2-r4` |
| [CVE-2025-9231](https://nvd.nist.gov/vuln/detail/CVE-2025-9231) | Medium | `libssl3` | `3.3.2-r4` |
| [CVE-2025-9231](https://nvd.nist.gov/vuln/detail/CVE-2025-9231) | Medium | `libcrypto3` | `3.3.2-r4` |
| [CVE-2025-9232](https://nvd.nist.gov/vuln/detail/CVE-2025-9232) | Medium | `libcrypto3` | `3.3.2-r4` |
| [CVE-2025-9232](https://nvd.nist.gov/vuln/detail/CVE-2025-9232) | Medium | `libssl3` | `3.3.2-r4` |
| [CVE-2026-22795](https://nvd.nist.gov/vuln/detail/CVE-2026-22795) | Medium | `libcrypto3` | `3.3.2-r4` |
| [CVE-2026-22795](https://nvd.nist.gov/vuln/detail/CVE-2026-22795) | Medium | `libssl3` | `3.3.2-r4` |
| [CVE-2026-22796](https://nvd.nist.gov/vuln/detail/CVE-2026-22796) | Medium | `libcrypto3` | `3.3.2-r4` |
| [CVE-2026-22796](https://nvd.nist.gov/vuln/detail/CVE-2026-22796) | Medium | `libssl3` | `3.3.2-r4` |
| [CVE-2024-58251](https://nvd.nist.gov/vuln/detail/CVE-2024-58251) | Low | `busybox-binsh` | `1.37.0-r9` |
| [CVE-2024-58251](https://nvd.nist.gov/vuln/detail/CVE-2024-58251) | Low | `busybox` | `1.37.0-r9` |
| [CVE-2024-58251](https://nvd.nist.gov/vuln/detail/CVE-2024-58251) | Low | `ssl_client` | `1.37.0-r9` |
| [CVE-2025-22873](https://nvd.nist.gov/vuln/detail/CVE-2025-22873) | Low | `stdlib` | `go1.22.12` |
| [CVE-2025-22873](https://nvd.nist.gov/vuln/detail/CVE-2025-22873) | Low | `stdlib` | `go1.22.12` |
| [CVE-2025-22873](https://nvd.nist.gov/vuln/detail/CVE-2025-22873) | Low | `stdlib` | `go1.22.12` |
| [CVE-2025-22873](https://nvd.nist.gov/vuln/detail/CVE-2025-22873) | Low | `go` | `1.22.12` |
| [CVE-2025-22873](https://nvd.nist.gov/vuln/detail/CVE-2025-22873) | Low | `stdlib` | `go1.22.12` |
| [CVE-2025-22873](https://nvd.nist.gov/vuln/detail/CVE-2025-22873) | Low | `stdlib` | `go1.22.12` |
| [CVE-2025-22873](https://nvd.nist.gov/vuln/detail/CVE-2025-22873) | Low | `stdlib` | `go1.22.12` |
| [CVE-2025-22873](https://nvd.nist.gov/vuln/detail/CVE-2025-22873) | Low | `stdlib` | `go1.22.12` |
| [CVE-2025-22873](https://nvd.nist.gov/vuln/detail/CVE-2025-22873) | Low | `stdlib` | `go1.22.12` |
| [CVE-2025-22873](https://nvd.nist.gov/vuln/detail/CVE-2025-22873) | Low | `stdlib` | `go1.22.12` |
| [CVE-2025-22873](https://nvd.nist.gov/vuln/detail/CVE-2025-22873) | Low | `stdlib` | `go1.22.12` |
| [CVE-2025-22873](https://nvd.nist.gov/vuln/detail/CVE-2025-22873) | Low | `stdlib` | `go1.22.12` |
| [CVE-2025-22873](https://nvd.nist.gov/vuln/detail/CVE-2025-22873) | Low | `stdlib` | `go1.22.12` |
| [CVE-2025-22873](https://nvd.nist.gov/vuln/detail/CVE-2025-22873) | Low | `stdlib` | `go1.22.12` |
| [CVE-2025-22873](https://nvd.nist.gov/vuln/detail/CVE-2025-22873) | Low | `stdlib` | `go1.22.12` |
| [CVE-2025-22873](https://nvd.nist.gov/vuln/detail/CVE-2025-22873) | Low | `stdlib` | `go1.22.12` |
| [CVE-2025-22873](https://nvd.nist.gov/vuln/detail/CVE-2025-22873) | Low | `stdlib` | `go1.22.12` |
| [CVE-2025-22873](https://nvd.nist.gov/vuln/detail/CVE-2025-22873) | Low | `stdlib` | `go1.22.12` |
| [CVE-2025-22873](https://nvd.nist.gov/vuln/detail/CVE-2025-22873) | Low | `stdlib` | `go1.22.12` |
| [CVE-2025-22873](https://nvd.nist.gov/vuln/detail/CVE-2025-22873) | Low | `stdlib` | `go1.22.12` |
| [CVE-2025-46394](https://nvd.nist.gov/vuln/detail/CVE-2025-46394) | Low | `ssl_client` | `1.37.0-r9` |
| [CVE-2025-46394](https://nvd.nist.gov/vuln/detail/CVE-2025-46394) | Low | `busybox-binsh` | `1.37.0-r9` |
| [CVE-2025-46394](https://nvd.nist.gov/vuln/detail/CVE-2025-46394) | Low | `busybox` | `1.37.0-r9` |
</details>

<details>
<summary><strong>📦 Installed Packages (38 total)</strong></summary>

| Package | Version |
|---------|---------|
| alpine-baselayout | 3.6.8-r1 |
| alpine-baselayout-data | 3.6.8-r1 |
| alpine-keys | 2.5-r0 |
| alpine-release | 3.21.2-r0 |
| apk-tools | 2.14.6-r2 |
| busybox | 1.37.0-r9 |
| busybox-binsh | 1.37.0-r9 |
| ca-certificates | 20241121-r1 |
| ca-certificates-bundle | 20241121-r1 |
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
| cmd/test2json | UNKNOWN |
| cmd/trace | UNKNOWN |
| cmd/vet | UNKNOWN |
| d3-pprof | 2.0.0 |
| go | 1.22.12 |
| libcrypto3 | 3.3.2-r4 |
| libssl3 | 3.3.2-r4 |
| musl | 1.2.5-r8 |
| musl-utils | 1.2.5-r8 |
| scanelf | 1.3.8-r1 |
| ssl_client | 1.37.0-r9 |
| stdlib | go1.22.12 |
| zlib | 1.3.1-r2 |
</details>

</details>

<details>
<summary><strong>🔍 Full Report: golang:1.23-alpine</strong></summary>

## 🛡️ Security & Efficiency

**Base Image:** `linux (arm64)`
**Supported Architectures:** `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown`
**Efficiency Score:** 99.9%

### Vulnerabilities
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| 🔴 23 | 🟠 176 | 212 | 6 |

<details>
<summary><strong>👇 Expand Vulnerability Details (417 found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [CVE-2025-15467](https://nvd.nist.gov/vuln/detail/CVE-2025-15467) | Critical | `libcrypto3` | `3.5.1-r0` |
| [CVE-2025-15467](https://nvd.nist.gov/vuln/detail/CVE-2025-15467) | Critical | `libssl3` | `3.5.1-r0` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.23.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.23.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.23.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `go` | `1.23.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.23.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.23.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.23.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.23.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.23.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.23.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.23.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.23.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.23.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.23.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.23.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.23.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.23.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.23.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.23.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.23.12` |
| [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/CVE-2025-68121) | Critical | `stdlib` | `go1.23.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58187](https://nvd.nist.gov/vuln/detail/CVE-2025-58187) | High | `go` | `1.23.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-58188](https://nvd.nist.gov/vuln/detail/CVE-2025-58188) | High | `go` | `1.23.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61723](https://nvd.nist.gov/vuln/detail/CVE-2025-61723) | High | `go` | `1.23.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `go` | `1.23.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61725](https://nvd.nist.gov/vuln/detail/CVE-2025-61725) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `go` | `1.23.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61726](https://nvd.nist.gov/vuln/detail/CVE-2025-61726) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `go` | `1.23.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61729](https://nvd.nist.gov/vuln/detail/CVE-2025-61729) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `go` | `1.23.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61731](https://nvd.nist.gov/vuln/detail/CVE-2025-61731) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `go` | `1.23.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-61732](https://nvd.nist.gov/vuln/detail/CVE-2025-61732) | High | `stdlib` | `go1.23.12` |
| [CVE-2025-69419](https://nvd.nist.gov/vuln/detail/CVE-2025-69419) | High | `libssl3` | `3.5.1-r0` |
| [CVE-2025-69419](https://nvd.nist.gov/vuln/detail/CVE-2025-69419) | High | `libcrypto3` | `3.5.1-r0` |
| [CVE-2025-69420](https://nvd.nist.gov/vuln/detail/CVE-2025-69420) | High | `libcrypto3` | `3.5.1-r0` |
| [CVE-2025-69420](https://nvd.nist.gov/vuln/detail/CVE-2025-69420) | High | `libssl3` | `3.5.1-r0` |
| [CVE-2025-69421](https://nvd.nist.gov/vuln/detail/CVE-2025-69421) | High | `libcrypto3` | `3.5.1-r0` |
| [CVE-2025-69421](https://nvd.nist.gov/vuln/detail/CVE-2025-69421) | High | `libssl3` | `3.5.1-r0` |
| [CVE-2025-9230](https://nvd.nist.gov/vuln/detail/CVE-2025-9230) | High | `libcrypto3` | `3.5.1-r0` |
| [CVE-2025-9230](https://nvd.nist.gov/vuln/detail/CVE-2025-9230) | High | `libssl3` | `3.5.1-r0` |
| [CVE-2025-11187](https://nvd.nist.gov/vuln/detail/CVE-2025-11187) | Medium | `libssl3` | `3.5.1-r0` |
| [CVE-2025-11187](https://nvd.nist.gov/vuln/detail/CVE-2025-11187) | Medium | `libcrypto3` | `3.5.1-r0` |
| [CVE-2025-15468](https://nvd.nist.gov/vuln/detail/CVE-2025-15468) | Medium | `libcrypto3` | `3.5.1-r0` |
| [CVE-2025-15468](https://nvd.nist.gov/vuln/detail/CVE-2025-15468) | Medium | `libssl3` | `3.5.1-r0` |
| [CVE-2025-15469](https://nvd.nist.gov/vuln/detail/CVE-2025-15469) | Medium | `libssl3` | `3.5.1-r0` |
| [CVE-2025-15469](https://nvd.nist.gov/vuln/detail/CVE-2025-15469) | Medium | `libcrypto3` | `3.5.1-r0` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `go` | `1.23.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-47912](https://nvd.nist.gov/vuln/detail/CVE-2025-47912) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `go` | `1.23.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58183](https://nvd.nist.gov/vuln/detail/CVE-2025-58183) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58185](https://nvd.nist.gov/vuln/detail/CVE-2025-58185) | Medium | `go` | `1.23.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58186](https://nvd.nist.gov/vuln/detail/CVE-2025-58186) | Medium | `go` | `1.23.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `go` | `1.23.12` |
| [CVE-2025-58189](https://nvd.nist.gov/vuln/detail/CVE-2025-58189) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `ssl_client` | `1.37.0-r18` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox-binsh` | `1.37.0-r18` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox` | `1.37.0-r18` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `go` | `1.23.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61724](https://nvd.nist.gov/vuln/detail/CVE-2025-61724) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `go` | `1.23.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61727](https://nvd.nist.gov/vuln/detail/CVE-2025-61727) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `go` | `1.23.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61728](https://nvd.nist.gov/vuln/detail/CVE-2025-61728) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `go` | `1.23.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-61730](https://nvd.nist.gov/vuln/detail/CVE-2025-61730) | Medium | `stdlib` | `go1.23.12` |
| [CVE-2025-66199](https://nvd.nist.gov/vuln/detail/CVE-2025-66199) | Medium | `libssl3` | `3.5.1-r0` |
| [CVE-2025-66199](https://nvd.nist.gov/vuln/detail/CVE-2025-66199) | Medium | `libcrypto3` | `3.5.1-r0` |
| [CVE-2025-68160](https://nvd.nist.gov/vuln/detail/CVE-2025-68160) | Medium | `libcrypto3` | `3.5.1-r0` |
| [CVE-2025-68160](https://nvd.nist.gov/vuln/detail/CVE-2025-68160) | Medium | `libssl3` | `3.5.1-r0` |
| [CVE-2025-69418](https://nvd.nist.gov/vuln/detail/CVE-2025-69418) | Medium | `libcrypto3` | `3.5.1-r0` |
| [CVE-2025-69418](https://nvd.nist.gov/vuln/detail/CVE-2025-69418) | Medium | `libssl3` | `3.5.1-r0` |
| [CVE-2025-9231](https://nvd.nist.gov/vuln/detail/CVE-2025-9231) | Medium | `libcrypto3` | `3.5.1-r0` |
| [CVE-2025-9231](https://nvd.nist.gov/vuln/detail/CVE-2025-9231) | Medium | `libssl3` | `3.5.1-r0` |
| [CVE-2025-9232](https://nvd.nist.gov/vuln/detail/CVE-2025-9232) | Medium | `libcrypto3` | `3.5.1-r0` |
| [CVE-2025-9232](https://nvd.nist.gov/vuln/detail/CVE-2025-9232) | Medium | `libssl3` | `3.5.1-r0` |
| [CVE-2026-22795](https://nvd.nist.gov/vuln/detail/CVE-2026-22795) | Medium | `libssl3` | `3.5.1-r0` |
| [CVE-2026-22795](https://nvd.nist.gov/vuln/detail/CVE-2026-22795) | Medium | `libcrypto3` | `3.5.1-r0` |
| [CVE-2026-22796](https://nvd.nist.gov/vuln/detail/CVE-2026-22796) | Medium | `libssl3` | `3.5.1-r0` |
| [CVE-2026-22796](https://nvd.nist.gov/vuln/detail/CVE-2026-22796) | Medium | `libcrypto3` | `3.5.1-r0` |
| [CVE-2024-58251](https://nvd.nist.gov/vuln/detail/CVE-2024-58251) | Low | `busybox` | `1.37.0-r18` |
| [CVE-2024-58251](https://nvd.nist.gov/vuln/detail/CVE-2024-58251) | Low | `busybox-binsh` | `1.37.0-r18` |
| [CVE-2024-58251](https://nvd.nist.gov/vuln/detail/CVE-2024-58251) | Low | `ssl_client` | `1.37.0-r18` |
| [CVE-2025-46394](https://nvd.nist.gov/vuln/detail/CVE-2025-46394) | Low | `busybox-binsh` | `1.37.0-r18` |
| [CVE-2025-46394](https://nvd.nist.gov/vuln/detail/CVE-2025-46394) | Low | `busybox` | `1.37.0-r18` |
| [CVE-2025-46394](https://nvd.nist.gov/vuln/detail/CVE-2025-46394) | Low | `ssl_client` | `1.37.0-r18` |
</details>

<details>
<summary><strong>📦 Installed Packages (39 total)</strong></summary>

| Package | Version |
|---------|---------|
| alpine-baselayout | 3.7.0-r0 |
| alpine-baselayout-data | 3.7.0-r0 |
| alpine-keys | 2.5-r0 |
| alpine-release | 3.22.1-r0 |
| apk-tools | 2.14.9-r2 |
| busybox | 1.37.0-r18 |
| busybox-binsh | 1.37.0-r18 |
| ca-certificates | 20250619-r0 |
| ca-certificates-bundle | 20250619-r0 |
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
| go | 1.23.12 |
| libapk2 | 2.14.9-r2 |
| libcrypto3 | 3.5.1-r0 |
| libssl3 | 3.5.1-r0 |
| musl | 1.2.5-r10 |
| musl-utils | 1.2.5-r10 |
| scanelf | 1.3.8-r1 |
| ssl_client | 1.37.0-r18 |
| stdlib | go1.23.12 |
| zlib | 1.3.1-r2 |
</details>

</details>

<details>
<summary><strong>🔍 Full Report: golang:1.24-alpine</strong></summary>

## 🛡️ Security & Efficiency

**Base Image:** `linux (arm64)`
**Supported Architectures:** `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown`
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

</details>

<!-- END: docker-docs:comparison -->
