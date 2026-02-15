# Node.js Docker Analysis

<!-- BEGIN: docker-docs:config -->

# 🐳 Docker Image Analysis: node:18-alpine
![Size](https://img.shields.io/static/v1?label=Size&message=122.10+MB&color=blue) ![Layers](https://img.shields.io/static/v1?label=Layers&message=4&color=blue) ![Vulns](https://img.shields.io/static/v1?label=Security&message=49+Vulns+%282+Crit%29&color=red) ![Efficiency](https://img.shields.io/static/v1?label=Efficiency&message=100.0%&color=green)

## ⚙️ Configuration
### Build Arguments
| Name | Description | Default | Required |
|------|-------------|---------|:--------:|
| `NODE_ENV=development` | Environment (development/production) | `development` | ❌ |
### Exposed Ports
| Port | Description |
|------|-------------|
| `3000` | The default port for the Node.js application |
---

## 🛡️ Security & Efficiency

**Base Image:** `linux (arm64)`
**Efficiency Score:** 100.0%

### Vulnerabilities
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| 🔴 2 | 🟠 18 | 20 | 9 |

<details>
<summary><strong>👇 Expand Vulnerability Details (49 found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [CVE-2025-15467](https://nvd.nist.gov/vuln/detail/CVE-2025-15467) | Critical | `libssl3` | `3.3.3-r0` |
| [CVE-2025-15467](https://nvd.nist.gov/vuln/detail/CVE-2025-15467) | Critical | `libcrypto3` | `3.3.3-r0` |
| [CVE-2025-23166](https://nvd.nist.gov/vuln/detail/CVE-2025-23166) | High | `node` | `18.20.8` |
| [CVE-2025-55131](https://nvd.nist.gov/vuln/detail/CVE-2025-55131) | High | `node` | `18.20.8` |
| [CVE-2025-59465](https://nvd.nist.gov/vuln/detail/CVE-2025-59465) | High | `node` | `18.20.8` |
| [CVE-2025-59466](https://nvd.nist.gov/vuln/detail/CVE-2025-59466) | High | `node` | `18.20.8` |
| [CVE-2025-69419](https://nvd.nist.gov/vuln/detail/CVE-2025-69419) | High | `libcrypto3` | `3.3.3-r0` |
| [CVE-2025-69419](https://nvd.nist.gov/vuln/detail/CVE-2025-69419) | High | `libssl3` | `3.3.3-r0` |
| [CVE-2025-69420](https://nvd.nist.gov/vuln/detail/CVE-2025-69420) | High | `libssl3` | `3.3.3-r0` |
| [CVE-2025-69420](https://nvd.nist.gov/vuln/detail/CVE-2025-69420) | High | `libcrypto3` | `3.3.3-r0` |
| [CVE-2025-69421](https://nvd.nist.gov/vuln/detail/CVE-2025-69421) | High | `libssl3` | `3.3.3-r0` |
| [CVE-2025-69421](https://nvd.nist.gov/vuln/detail/CVE-2025-69421) | High | `libcrypto3` | `3.3.3-r0` |
| [CVE-2025-9230](https://nvd.nist.gov/vuln/detail/CVE-2025-9230) | High | `libcrypto3` | `3.3.3-r0` |
| [CVE-2025-9230](https://nvd.nist.gov/vuln/detail/CVE-2025-9230) | High | `libssl3` | `3.3.3-r0` |
| [CVE-2026-21637](https://nvd.nist.gov/vuln/detail/CVE-2026-21637) | High | `node` | `18.20.8` |
| [GHSA-34x7-hfp2-rc4v](https://nvd.nist.gov/vuln/detail/GHSA-34x7-hfp2-rc4v) | High | `tar` | `6.2.1` |
| [GHSA-3xgq-45jj-v275](https://nvd.nist.gov/vuln/detail/GHSA-3xgq-45jj-v275) | High | `cross-spawn` | `7.0.3` |
| [GHSA-5j98-mcp5-4vw2](https://nvd.nist.gov/vuln/detail/GHSA-5j98-mcp5-4vw2) | High | `glob` | `10.4.2` |
| [GHSA-8qq5-rm4j-mr97](https://nvd.nist.gov/vuln/detail/GHSA-8qq5-rm4j-mr97) | High | `tar` | `6.2.1` |
| [GHSA-r6q2-hw4h-h46w](https://nvd.nist.gov/vuln/detail/GHSA-r6q2-hw4h-h46w) | High | `tar` | `6.2.1` |
| [CVE-2025-15468](https://nvd.nist.gov/vuln/detail/CVE-2025-15468) | Medium | `libcrypto3` | `3.3.3-r0` |
| [CVE-2025-15468](https://nvd.nist.gov/vuln/detail/CVE-2025-15468) | Medium | `libssl3` | `3.3.3-r0` |
| [CVE-2025-23167](https://nvd.nist.gov/vuln/detail/CVE-2025-23167) | Medium | `node` | `18.20.8` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `ssl_client` | `1.37.0-r12` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox-binsh` | `1.37.0-r12` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox` | `1.37.0-r12` |
| [CVE-2025-66199](https://nvd.nist.gov/vuln/detail/CVE-2025-66199) | Medium | `libssl3` | `3.3.3-r0` |
| [CVE-2025-66199](https://nvd.nist.gov/vuln/detail/CVE-2025-66199) | Medium | `libcrypto3` | `3.3.3-r0` |
| [CVE-2025-68160](https://nvd.nist.gov/vuln/detail/CVE-2025-68160) | Medium | `libcrypto3` | `3.3.3-r0` |
| [CVE-2025-68160](https://nvd.nist.gov/vuln/detail/CVE-2025-68160) | Medium | `libssl3` | `3.3.3-r0` |
| [CVE-2025-69418](https://nvd.nist.gov/vuln/detail/CVE-2025-69418) | Medium | `libcrypto3` | `3.3.3-r0` |
| [CVE-2025-69418](https://nvd.nist.gov/vuln/detail/CVE-2025-69418) | Medium | `libssl3` | `3.3.3-r0` |
| [CVE-2025-9231](https://nvd.nist.gov/vuln/detail/CVE-2025-9231) | Medium | `libcrypto3` | `3.3.3-r0` |
| [CVE-2025-9231](https://nvd.nist.gov/vuln/detail/CVE-2025-9231) | Medium | `libssl3` | `3.3.3-r0` |
| [CVE-2025-9232](https://nvd.nist.gov/vuln/detail/CVE-2025-9232) | Medium | `libssl3` | `3.3.3-r0` |
| [CVE-2025-9232](https://nvd.nist.gov/vuln/detail/CVE-2025-9232) | Medium | `libcrypto3` | `3.3.3-r0` |
| [CVE-2026-22795](https://nvd.nist.gov/vuln/detail/CVE-2026-22795) | Medium | `libcrypto3` | `3.3.3-r0` |
| [CVE-2026-22795](https://nvd.nist.gov/vuln/detail/CVE-2026-22795) | Medium | `libssl3` | `3.3.3-r0` |
| [CVE-2026-22796](https://nvd.nist.gov/vuln/detail/CVE-2026-22796) | Medium | `libcrypto3` | `3.3.3-r0` |
| [CVE-2026-22796](https://nvd.nist.gov/vuln/detail/CVE-2026-22796) | Medium | `libssl3` | `3.3.3-r0` |
| [CVE-2024-58251](https://nvd.nist.gov/vuln/detail/CVE-2024-58251) | Low | `busybox` | `1.37.0-r12` |
| [CVE-2024-58251](https://nvd.nist.gov/vuln/detail/CVE-2024-58251) | Low | `busybox-binsh` | `1.37.0-r12` |
| [CVE-2024-58251](https://nvd.nist.gov/vuln/detail/CVE-2024-58251) | Low | `ssl_client` | `1.37.0-r12` |
| [CVE-2025-23165](https://nvd.nist.gov/vuln/detail/CVE-2025-23165) | Low | `node` | `18.20.8` |
| [CVE-2025-46394](https://nvd.nist.gov/vuln/detail/CVE-2025-46394) | Low | `ssl_client` | `1.37.0-r12` |
| [CVE-2025-46394](https://nvd.nist.gov/vuln/detail/CVE-2025-46394) | Low | `busybox-binsh` | `1.37.0-r12` |
| [CVE-2025-46394](https://nvd.nist.gov/vuln/detail/CVE-2025-46394) | Low | `busybox` | `1.37.0-r12` |
| [GHSA-73rr-hh4g-fpgx](https://nvd.nist.gov/vuln/detail/GHSA-73rr-hh4g-fpgx) | Low | `diff` | `5.2.0` |
| [GHSA-v6h2-p8h4-qcjw](https://nvd.nist.gov/vuln/detail/GHSA-v6h2-p8h4-qcjw) | Low | `brace-expansion` | `2.0.1` |
</details>

<details>
<summary><strong>📦 Installed Packages (211 total)</strong></summary>

| Package | Version |
|---------|---------|
| @isaacs/cliui | 8.0.2 |
| @isaacs/string-locale-compare | 1.1.0 |
| @npmcli/agent | 2.2.2 |
| @npmcli/arborist | 7.5.4 |
| @npmcli/config | 8.3.4 |
| @npmcli/fs | 3.1.1 |
| @npmcli/git | 5.0.8 |
| @npmcli/installed-package-contents | 2.1.0 |
| @npmcli/map-workspaces | 3.0.6 |
| @npmcli/metavuln-calculator | 7.1.1 |
| @npmcli/name-from-folder | 2.0.0 |
| @npmcli/node-gyp | 3.0.0 |
| @npmcli/package-json | 5.2.0 |
| @npmcli/promise-spawn | 7.0.2 |
| @npmcli/query | 3.1.0 |
| @npmcli/redact | 2.0.1 |
| @npmcli/run-script | 8.1.0 |
| @pkgjs/parseargs | 0.11.0 |
| @sigstore/bundle | 2.3.2 |
| @sigstore/core | 1.1.0 |
| @sigstore/protobuf-specs | 0.3.2 |
| @sigstore/sign | 2.3.2 |
| @sigstore/tuf | 2.3.4 |
| @sigstore/verify | 1.2.1 |
| @tufjs/canonical-json | 2.0.0 |
| @tufjs/models | 2.0.1 |
| abbrev | 2.0.0 |
| agent-base | 7.1.1 |
| aggregate-error | 3.1.0 |
| alpine-baselayout | 3.6.8-r1 |
| alpine-baselayout-data | 3.6.8-r1 |
| alpine-keys | 2.5-r0 |
| alpine-release | 3.21.3-r0 |
| ansi-regex | 5.0.1 |
| ansi-regex | 6.0.1 |
| ansi-styles | 4.3.0 |
| ansi-styles | 6.2.1 |
| apk-tools | 2.14.6-r3 |
| aproba | 2.0.0 |
| archy | 1.0.0 |
| balanced-match | 1.0.2 |
| bin-links | 4.0.4 |
| binary-extensions | 2.3.0 |
| brace-expansion | 2.0.1 |
| busybox | 1.37.0-r12 |
| busybox-binsh | 1.37.0-r12 |
| ca-certificates-bundle | 20241121-r1 |
| cacache | 18.0.3 |
| chalk | 5.3.0 |
| chownr | 2.0.0 |
| ci-info | 4.0.0 |
| cidr-regex | 4.1.1 |
| clean-stack | 2.2.0 |
| cli-columns | 4.0.0 |
| cmd-shim | 6.0.3 |
| color-convert | 2.0.1 |
| color-name | 1.1.4 |
| common-ancestor-path | 1.0.1 |
| corepack | 0.32.0 |
| cross-spawn | 7.0.3 |
| cssesc | 3.0.0 |
| debug | 4.3.5 |
| diff | 5.2.0 |
| eastasianwidth | 0.2.0 |
| emoji-regex | 8.0.0 |
| emoji-regex | 9.2.2 |
| encoding | 0.1.13 |
| env-paths | 2.2.1 |
| err-code | 2.0.3 |
| exponential-backoff | 3.1.1 |
| fastest-levenshtein | 1.0.16 |
| foreground-child | 3.2.1 |
| fs-minipass | 2.1.0 |
| fs-minipass | 3.0.3 |
| glob | 10.4.2 |
| graceful-fs | 4.2.11 |
| hosted-git-info | 7.0.2 |
| http-cache-semantics | 4.1.1 |
| http-proxy-agent | 7.0.2 |
| https-proxy-agent | 7.0.5 |
| iconv-lite | 0.6.3 |
| ignore-walk | 6.0.5 |
| imurmurhash | 0.1.4 |
| indent-string | 4.0.0 |
| ini | 4.1.3 |
| init-package-json | 6.0.3 |
| ip-address | 9.0.5 |
| ip-regex | 5.0.0 |
| is-cidr | 5.1.0 |
| is-fullwidth-code-point | 3.0.0 |
| is-lambda | 1.0.1 |
| isexe | 2.0.0 |
| isexe | 3.1.1 |
| jackspeak | 3.4.0 |
| jsbn | 1.1.0 |
| json-parse-even-better-errors | 3.0.2 |
| json-stringify-nice | 1.1.4 |
| jsonparse | 1.3.1 |
| just-diff | 6.0.2 |
| just-diff-apply | 5.5.0 |
| libcrypto3 | 3.3.3-r0 |
| libgcc | 14.2.0-r4 |
| libnpmaccess | 8.0.6 |
| libnpmdiff | 6.1.4 |
| libnpmexec | 8.1.3 |
| libnpmfund | 5.0.12 |
| libnpmhook | 10.0.5 |
| libnpmorg | 6.0.6 |
| libnpmpack | 7.0.4 |
| libnpmpublish | 9.0.9 |
| libnpmsearch | 7.0.6 |
| libnpmteam | 6.0.5 |
| libnpmversion | 6.0.3 |
| libssl3 | 3.3.3-r0 |
| libstdc++ | 14.2.0-r4 |
| lru-cache | 10.2.2 |
| make-fetch-happen | 13.0.1 |
| minimatch | 9.0.5 |
| minipass | 3.3.6 |
| minipass | 5.0.0 |
| minipass | 7.1.2 |
| minipass-collect | 2.0.1 |
| minipass-fetch | 3.0.5 |
| minipass-flush | 1.0.5 |
| minipass-pipeline | 1.2.4 |
| minipass-sized | 1.0.3 |
| minizlib | 2.1.2 |
| mkdirp | 1.0.4 |
| ms | 2.1.2 |
| ms | 2.1.3 |
| musl | 1.2.5-r9 |
| musl-utils | 1.2.5-r9 |
| mute-stream | 1.0.0 |
| negotiator | 0.6.3 |
| node | 18.20.8 |
| node-gyp | 10.1.0 |
| nopt | 7.2.1 |
| normalize-package-data | 6.0.2 |
| npm | 10.8.2 |
| npm-audit-report | 5.0.0 |
| npm-bundled | 3.0.1 |
| npm-install-checks | 6.3.0 |
| npm-normalize-package-bin | 3.0.1 |
| npm-package-arg | 11.0.2 |
| npm-packlist | 8.0.2 |
| npm-pick-manifest | 9.1.0 |
| npm-profile | 10.0.0 |
| npm-registry-fetch | 17.1.0 |
| npm-user-validate | 2.0.1 |
| p-map | 4.0.0 |
| package-json-from-dist | 1.0.0 |
| pacote | 18.0.6 |
| parse-conflict-json | 3.0.1 |
| path-key | 3.1.1 |
| path-scurry | 1.11.1 |
| postcss-selector-parser | 6.1.0 |
| proc-log | 3.0.0 |
| proc-log | 4.2.0 |
| proggy | 2.0.0 |
| promise-all-reject-late | 1.0.1 |
| promise-call-limit | 3.0.1 |
| promise-inflight | 1.0.1 |
| promise-retry | 2.0.1 |
| promzard | 1.0.2 |
| qrcode-terminal | 0.12.0 |
| read | 3.0.1 |
| read-cmd-shim | 4.0.0 |
| read-package-json-fast | 3.0.2 |
| retry | 0.12.0 |
| safer-buffer | 2.1.2 |
| scanelf | 1.3.8-r1 |
| semver | 7.6.2 |
| shebang-command | 2.0.0 |
| shebang-regex | 3.0.0 |
| signal-exit | 4.1.0 |
| sigstore | 2.3.1 |
| smart-buffer | 4.2.0 |
| socks | 2.8.3 |
| socks-proxy-agent | 8.0.4 |
| spdx-correct | 3.2.0 |
| spdx-exceptions | 2.5.0 |
| spdx-expression-parse | 3.0.1 |
| spdx-expression-parse | 4.0.0 |
| spdx-license-ids | 3.0.18 |
| sprintf-js | 1.1.3 |
| ssl_client | 1.37.0-r12 |
| ssri | 10.0.6 |
| string-width | 4.2.3 |
| string-width | 5.1.2 |
| strip-ansi | 6.0.1 |
| strip-ansi | 7.1.0 |
| supports-color | 9.4.0 |
| tar | 6.2.1 |
| text-table | 0.2.0 |
| tiny-relative-date | 1.3.0 |
| treeverse | 3.0.0 |
| tuf-js | 2.2.1 |
| unique-filename | 3.0.0 |
| unique-slug | 4.0.0 |
| util-deprecate | 1.0.2 |
| validate-npm-package-license | 3.0.4 |
| validate-npm-package-name | 5.0.1 |
| walk-up-path | 3.0.1 |
| which | 2.0.2 |
| which | 4.0.0 |
| wrap-ansi | 7.0.0 |
| wrap-ansi | 8.1.0 |
| write-file-atomic | 5.0.1 |
| yallist | 4.0.0 |
| yarn | 1.22.22 |
| zlib | 1.3.1-r2 |
</details>

<!-- END: docker-docs:config -->

## Version Comparison

<!-- BEGIN: docker-docs:comparison -->

### 🏷️ Supported Tags

| Tag | Size | Vulns | Efficiency | OS/Arch |
|-----|------|-------|------------|---------|
| `node:16-alpine` | ![Size](https://img.shields.io/static/v1?label=Size&message=113.67+MB&color=blue) | ![Vulns](https://img.shields.io/static/v1?label=Security&message=76+Vulns+%282+Crit%29&color=red) | 99.6% | linux/arm64 |
| `node:18-alpine` | ![Size](https://img.shields.io/static/v1?label=Size&message=122.10+MB&color=blue) | ![Vulns](https://img.shields.io/static/v1?label=Security&message=49+Vulns+%282+Crit%29&color=red) | 100.0% | linux/arm64 |
| `node:20-alpine` | ![Size](https://img.shields.io/static/v1?label=Size&message=130.66+MB&color=blue) | ![Vulns](https://img.shields.io/static/v1?label=Security&message=10+Vulns+%280+Crit%29&color=orange) | 100.0% | linux/arm64 |

<details>
<summary><strong>🔍 Full Report: node:16-alpine</strong></summary>

## 🛡️ Security & Efficiency

**Base Image:** `linux (arm64)`
**Efficiency Score:** 99.6%

### Vulnerabilities
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| 🔴 2 | 🟠 24 | 38 | 12 |

<details>
<summary><strong>👇 Expand Vulnerability Details (76 found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [CVE-2024-5535](https://nvd.nist.gov/vuln/detail/CVE-2024-5535) | Critical | `libcrypto3` | `3.1.2-r0` |
| [CVE-2024-5535](https://nvd.nist.gov/vuln/detail/CVE-2024-5535) | Critical | `libssl3` | `3.1.2-r0` |
| [CVE-2023-38552](https://nvd.nist.gov/vuln/detail/CVE-2023-38552) | High | `node` | `16.20.2` |
| [CVE-2023-44487](https://nvd.nist.gov/vuln/detail/CVE-2023-44487) | High | `node` | `16.20.2` |
| [CVE-2023-46809](https://nvd.nist.gov/vuln/detail/CVE-2023-46809) | High | `node` | `16.20.2` |
| [CVE-2023-5363](https://nvd.nist.gov/vuln/detail/CVE-2023-5363) | High | `libcrypto3` | `3.1.2-r0` |
| [CVE-2023-5363](https://nvd.nist.gov/vuln/detail/CVE-2023-5363) | High | `libssl3` | `3.1.2-r0` |
| [CVE-2024-21892](https://nvd.nist.gov/vuln/detail/CVE-2024-21892) | High | `node` | `16.20.2` |
| [CVE-2024-22019](https://nvd.nist.gov/vuln/detail/CVE-2024-22019) | High | `node` | `16.20.2` |
| [CVE-2024-27983](https://nvd.nist.gov/vuln/detail/CVE-2024-27983) | High | `node` | `16.20.2` |
| [CVE-2024-4741](https://nvd.nist.gov/vuln/detail/CVE-2024-4741) | High | `libssl3` | `3.1.2-r0` |
| [CVE-2024-4741](https://nvd.nist.gov/vuln/detail/CVE-2024-4741) | High | `libcrypto3` | `3.1.2-r0` |
| [CVE-2024-6119](https://nvd.nist.gov/vuln/detail/CVE-2024-6119) | High | `libssl3` | `3.1.2-r0` |
| [CVE-2024-6119](https://nvd.nist.gov/vuln/detail/CVE-2024-6119) | High | `libcrypto3` | `3.1.2-r0` |
| [CVE-2025-23166](https://nvd.nist.gov/vuln/detail/CVE-2025-23166) | High | `node` | `16.20.2` |
| [CVE-2025-26519](https://nvd.nist.gov/vuln/detail/CVE-2025-26519) | High | `musl-utils` | `1.2.4-r1` |
| [CVE-2025-26519](https://nvd.nist.gov/vuln/detail/CVE-2025-26519) | High | `musl` | `1.2.4-r1` |
| [CVE-2025-55131](https://nvd.nist.gov/vuln/detail/CVE-2025-55131) | High | `node` | `16.20.2` |
| [CVE-2025-59465](https://nvd.nist.gov/vuln/detail/CVE-2025-59465) | High | `node` | `16.20.2` |
| [CVE-2025-59466](https://nvd.nist.gov/vuln/detail/CVE-2025-59466) | High | `node` | `16.20.2` |
| [CVE-2026-21637](https://nvd.nist.gov/vuln/detail/CVE-2026-21637) | High | `node` | `16.20.2` |
| [GHSA-2p57-rm9w-gvfp](https://nvd.nist.gov/vuln/detail/GHSA-2p57-rm9w-gvfp) | High | `ip` | `2.0.0` |
| [GHSA-34x7-hfp2-rc4v](https://nvd.nist.gov/vuln/detail/GHSA-34x7-hfp2-rc4v) | High | `tar` | `6.1.11` |
| [GHSA-8qq5-rm4j-mr97](https://nvd.nist.gov/vuln/detail/GHSA-8qq5-rm4j-mr97) | High | `tar` | `6.1.11` |
| [GHSA-c2qf-rxjj-qqgw](https://nvd.nist.gov/vuln/detail/GHSA-c2qf-rxjj-qqgw) | High | `semver` | `7.3.7` |
| [GHSA-r6q2-hw4h-h46w](https://nvd.nist.gov/vuln/detail/GHSA-r6q2-hw4h-h46w) | High | `tar` | `6.1.11` |
| [CVE-2023-39333](https://nvd.nist.gov/vuln/detail/CVE-2023-39333) | Medium | `node` | `16.20.2` |
| [CVE-2023-42363](https://nvd.nist.gov/vuln/detail/CVE-2023-42363) | Medium | `ssl_client` | `1.36.1-r2` |
| [CVE-2023-42363](https://nvd.nist.gov/vuln/detail/CVE-2023-42363) | Medium | `busybox-binsh` | `1.36.1-r2` |
| [CVE-2023-42363](https://nvd.nist.gov/vuln/detail/CVE-2023-42363) | Medium | `busybox` | `1.36.1-r2` |
| [CVE-2023-42364](https://nvd.nist.gov/vuln/detail/CVE-2023-42364) | Medium | `busybox` | `1.36.1-r2` |
| [CVE-2023-42364](https://nvd.nist.gov/vuln/detail/CVE-2023-42364) | Medium | `ssl_client` | `1.36.1-r2` |
| [CVE-2023-42364](https://nvd.nist.gov/vuln/detail/CVE-2023-42364) | Medium | `busybox-binsh` | `1.36.1-r2` |
| [CVE-2023-42365](https://nvd.nist.gov/vuln/detail/CVE-2023-42365) | Medium | `ssl_client` | `1.36.1-r2` |
| [CVE-2023-42365](https://nvd.nist.gov/vuln/detail/CVE-2023-42365) | Medium | `busybox-binsh` | `1.36.1-r2` |
| [CVE-2023-42365](https://nvd.nist.gov/vuln/detail/CVE-2023-42365) | Medium | `busybox` | `1.36.1-r2` |
| [CVE-2023-42366](https://nvd.nist.gov/vuln/detail/CVE-2023-42366) | Medium | `busybox-binsh` | `1.36.1-r2` |
| [CVE-2023-42366](https://nvd.nist.gov/vuln/detail/CVE-2023-42366) | Medium | `ssl_client` | `1.36.1-r2` |
| [CVE-2023-42366](https://nvd.nist.gov/vuln/detail/CVE-2023-42366) | Medium | `busybox` | `1.36.1-r2` |
| [CVE-2023-5678](https://nvd.nist.gov/vuln/detail/CVE-2023-5678) | Medium | `libcrypto3` | `3.1.2-r0` |
| [CVE-2023-5678](https://nvd.nist.gov/vuln/detail/CVE-2023-5678) | Medium | `libssl3` | `3.1.2-r0` |
| [CVE-2023-6129](https://nvd.nist.gov/vuln/detail/CVE-2023-6129) | Medium | `libssl3` | `3.1.2-r0` |
| [CVE-2023-6129](https://nvd.nist.gov/vuln/detail/CVE-2023-6129) | Medium | `libcrypto3` | `3.1.2-r0` |
| [CVE-2023-6237](https://nvd.nist.gov/vuln/detail/CVE-2023-6237) | Medium | `libcrypto3` | `3.1.2-r0` |
| [CVE-2023-6237](https://nvd.nist.gov/vuln/detail/CVE-2023-6237) | Medium | `libssl3` | `3.1.2-r0` |
| [CVE-2024-0727](https://nvd.nist.gov/vuln/detail/CVE-2024-0727) | Medium | `libssl3` | `3.1.2-r0` |
| [CVE-2024-0727](https://nvd.nist.gov/vuln/detail/CVE-2024-0727) | Medium | `libcrypto3` | `3.1.2-r0` |
| [CVE-2024-13176](https://nvd.nist.gov/vuln/detail/CVE-2024-13176) | Medium | `libcrypto3` | `3.1.2-r0` |
| [CVE-2024-13176](https://nvd.nist.gov/vuln/detail/CVE-2024-13176) | Medium | `libssl3` | `3.1.2-r0` |
| [CVE-2024-22020](https://nvd.nist.gov/vuln/detail/CVE-2024-22020) | Medium | `node` | `16.20.2` |
| [CVE-2024-22025](https://nvd.nist.gov/vuln/detail/CVE-2024-22025) | Medium | `node` | `16.20.2` |
| [CVE-2024-2511](https://nvd.nist.gov/vuln/detail/CVE-2024-2511) | Medium | `libssl3` | `3.1.2-r0` |
| [CVE-2024-2511](https://nvd.nist.gov/vuln/detail/CVE-2024-2511) | Medium | `libcrypto3` | `3.1.2-r0` |
| [CVE-2024-27982](https://nvd.nist.gov/vuln/detail/CVE-2024-27982) | Medium | `node` | `16.20.2` |
| [CVE-2024-4603](https://nvd.nist.gov/vuln/detail/CVE-2024-4603) | Medium | `libssl3` | `3.1.2-r0` |
| [CVE-2024-4603](https://nvd.nist.gov/vuln/detail/CVE-2024-4603) | Medium | `libcrypto3` | `3.1.2-r0` |
| [CVE-2024-9143](https://nvd.nist.gov/vuln/detail/CVE-2024-9143) | Medium | `libcrypto3` | `3.1.2-r0` |
| [CVE-2024-9143](https://nvd.nist.gov/vuln/detail/CVE-2024-9143) | Medium | `libssl3` | `3.1.2-r0` |
| [CVE-2025-23085](https://nvd.nist.gov/vuln/detail/CVE-2025-23085) | Medium | `node` | `16.20.2` |
| [CVE-2025-23167](https://nvd.nist.gov/vuln/detail/CVE-2025-23167) | Medium | `node` | `16.20.2` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `ssl_client` | `1.36.1-r2` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox-binsh` | `1.36.1-r2` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox` | `1.36.1-r2` |
| [GHSA-f5x3-32g6-xq36](https://nvd.nist.gov/vuln/detail/GHSA-f5x3-32g6-xq36) | Medium | `tar` | `6.1.11` |
| [CVE-2024-58251](https://nvd.nist.gov/vuln/detail/CVE-2024-58251) | Low | `busybox` | `1.36.1-r2` |
| [CVE-2024-58251](https://nvd.nist.gov/vuln/detail/CVE-2024-58251) | Low | `ssl_client` | `1.36.1-r2` |
| [CVE-2024-58251](https://nvd.nist.gov/vuln/detail/CVE-2024-58251) | Low | `busybox-binsh` | `1.36.1-r2` |
| [CVE-2025-23165](https://nvd.nist.gov/vuln/detail/CVE-2025-23165) | Low | `node` | `16.20.2` |
| [CVE-2025-46394](https://nvd.nist.gov/vuln/detail/CVE-2025-46394) | Low | `ssl_client` | `1.36.1-r2` |
| [CVE-2025-46394](https://nvd.nist.gov/vuln/detail/CVE-2025-46394) | Low | `busybox-binsh` | `1.36.1-r2` |
| [CVE-2025-46394](https://nvd.nist.gov/vuln/detail/CVE-2025-46394) | Low | `busybox` | `1.36.1-r2` |
| [GHSA-73rr-hh4g-fpgx](https://nvd.nist.gov/vuln/detail/GHSA-73rr-hh4g-fpgx) | Low | `diff` | `5.1.0` |
| [GHSA-78xj-cgh5-2h22](https://nvd.nist.gov/vuln/detail/GHSA-78xj-cgh5-2h22) | Low | `ip` | `2.0.0` |
| [GHSA-v6h2-p8h4-qcjw](https://nvd.nist.gov/vuln/detail/GHSA-v6h2-p8h4-qcjw) | Low | `brace-expansion` | `1.1.11` |
| [GHSA-v6h2-p8h4-qcjw](https://nvd.nist.gov/vuln/detail/GHSA-v6h2-p8h4-qcjw) | Low | `brace-expansion` | `1.1.11` |
| [GHSA-v6h2-p8h4-qcjw](https://nvd.nist.gov/vuln/detail/GHSA-v6h2-p8h4-qcjw) | Low | `brace-expansion` | `2.0.1` |
</details>

<details>
<summary><strong>📦 Installed Packages (225 total)</strong></summary>

| Package | Version |
|---------|---------|
| @colors/colors | 1.5.0 |
| @gar/promisify | 1.1.3 |
| @isaacs/string-locale-compare | 1.1.0 |
| @npmcli/arborist | 5.6.3 |
| @npmcli/ci-detect | 2.0.0 |
| @npmcli/config | 4.2.2 |
| @npmcli/disparity-colors | 2.0.0 |
| @npmcli/fs | 2.1.2 |
| @npmcli/git | 3.0.2 |
| @npmcli/installed-package-contents | 1.0.7 |
| @npmcli/map-workspaces | 2.0.4 |
| @npmcli/metavuln-calculator | 3.1.1 |
| @npmcli/move-file | 2.0.1 |
| @npmcli/name-from-folder | 1.0.1 |
| @npmcli/node-gyp | 2.0.0 |
| @npmcli/package-json | 2.0.0 |
| @npmcli/promise-spawn | 3.0.0 |
| @npmcli/query | 1.2.0 |
| @npmcli/run-script | 4.2.1 |
| @tootallnate/once | 2.0.0 |
| abbrev | 1.1.1 |
| agent-base | 6.0.2 |
| agentkeepalive | 4.2.1 |
| aggregate-error | 3.1.0 |
| alpine-baselayout | 3.4.3-r1 |
| alpine-baselayout-data | 3.4.3-r1 |
| alpine-keys | 2.4-r1 |
| ansi-regex | 5.0.1 |
| ansi-styles | 4.3.0 |
| apk-tools | 2.14.0-r2 |
| aproba | 2.0.0 |
| archy | 1.0.0 |
| are-we-there-yet | 3.0.1 |
| asap | 2.0.6 |
| balanced-match | 1.0.2 |
| bin-links | 3.0.3 |
| binary-extensions | 2.2.0 |
| brace-expansion | 1.1.11 |
| brace-expansion | 2.0.1 |
| builtins | 5.0.1 |
| busybox | 1.36.1-r2 |
| busybox-binsh | 1.36.1-r2 |
| ca-certificates-bundle | 20230506-r0 |
| cacache | 16.1.3 |
| chalk | 4.1.2 |
| chownr | 2.0.0 |
| cidr-regex | 3.1.1 |
| clean-stack | 2.2.0 |
| cli-columns | 4.0.0 |
| cli-table3 | 0.6.2 |
| clone | 1.0.4 |
| cmd-shim | 5.0.0 |
| color-convert | 2.0.1 |
| color-name | 1.1.4 |
| color-support | 1.1.3 |
| columnify | 1.6.0 |
| common-ancestor-path | 1.0.1 |
| concat-map | 0.0.1 |
| console-control-strings | 1.1.0 |
| corepack | 0.17.0 |
| cssesc | 3.0.0 |
| debug | 4.3.4 |
| debuglog | 1.0.1 |
| defaults | 1.0.3 |
| delegates | 1.0.0 |
| depd | 1.1.2 |
| dezalgo | 1.0.4 |
| diff | 5.1.0 |
| emoji-regex | 8.0.0 |
| encoding | 0.1.13 |
| env-paths | 2.2.1 |
| err-code | 2.0.3 |
| fastest-levenshtein | 1.0.12 |
| fs-minipass | 2.1.0 |
| fs.realpath | 1.0.0 |
| function-bind | 1.1.1 |
| gauge | 4.0.4 |
| glob | 7.2.3 |
| glob | 8.0.3 |
| graceful-fs | 4.2.10 |
| has | 1.0.3 |
| has-flag | 4.0.0 |
| has-unicode | 2.0.1 |
| hosted-git-info | 5.2.1 |
| http-cache-semantics | 4.1.1 |
| http-proxy-agent | 5.0.0 |
| https-proxy-agent | 5.0.1 |
| humanize-ms | 1.2.1 |
| iconv-lite | 0.6.3 |
| ignore-walk | 5.0.1 |
| imurmurhash | 0.1.4 |
| indent-string | 4.0.0 |
| infer-owner | 1.0.4 |
| inflight | 1.0.6 |
| inherits | 2.0.4 |
| ini | 3.0.1 |
| init-package-json | 3.0.2 |
| ip | 2.0.0 |
| ip-regex | 4.3.0 |
| is-cidr | 4.0.2 |
| is-core-module | 2.10.0 |
| is-fullwidth-code-point | 3.0.0 |
| is-lambda | 1.0.1 |
| isexe | 2.0.0 |
| json-parse-even-better-errors | 2.3.1 |
| json-stringify-nice | 1.1.4 |
| jsonparse | 1.3.1 |
| just-diff | 5.1.1 |
| just-diff-apply | 5.4.1 |
| libc-utils | 0.7.2-r5 |
| libcrypto3 | 3.1.2-r0 |
| libgcc | 12.2.1_git20220924-r10 |
| libnpmaccess | 6.0.4 |
| libnpmdiff | 4.0.5 |
| libnpmexec | 4.0.14 |
| libnpmfund | 3.0.5 |
| libnpmhook | 8.0.4 |
| libnpmorg | 4.0.4 |
| libnpmpack | 4.1.3 |
| libnpmpublish | 6.0.5 |
| libnpmsearch | 5.0.4 |
| libnpmteam | 4.0.4 |
| libnpmversion | 3.0.7 |
| libssl3 | 3.1.2-r0 |
| libstdc++ | 12.2.1_git20220924-r10 |
| lru-cache | 6.0.0 |
| lru-cache | 7.13.2 |
| make-fetch-happen | 10.2.1 |
| minimatch | 3.1.2 |
| minimatch | 5.1.0 |
| minipass | 3.3.4 |
| minipass-collect | 1.0.2 |
| minipass-fetch | 2.1.1 |
| minipass-flush | 1.0.5 |
| minipass-json-stream | 1.0.1 |
| minipass-pipeline | 1.2.4 |
| minipass-sized | 1.0.3 |
| minizlib | 2.1.2 |
| mkdirp | 1.0.4 |
| mkdirp-infer-owner | 2.0.0 |
| ms | 2.1.2 |
| ms | 2.1.3 |
| musl | 1.2.4-r1 |
| musl-utils | 1.2.4-r1 |
| mute-stream | 0.0.8 |
| negotiator | 0.6.3 |
| node | 16.20.2 |
| node-gyp | 9.1.0 |
| nopt | 5.0.0 |
| nopt | 6.0.0 |
| normalize-package-data | 4.0.1 |
| npm | 8.19.4 |
| npm-audit-report | 3.0.0 |
| npm-bundled | 1.1.2 |
| npm-bundled | 2.0.1 |
| npm-init | 0.0.0 |
| npm-install-checks | 5.0.0 |
| npm-normalize-package-bin | 1.0.1 |
| npm-normalize-package-bin | 2.0.0 |
| npm-package-arg | 9.1.0 |
| npm-packlist | 5.1.3 |
| npm-pick-manifest | 7.0.2 |
| npm-profile | 6.2.1 |
| npm-registry-fetch | 13.3.1 |
| npm-user-validate | 1.0.1 |
| npmlog | 6.0.2 |
| once | 1.4.0 |
| opener | 1.5.2 |
| p-map | 4.0.0 |
| pacote | 13.6.2 |
| parse-conflict-json | 2.0.2 |
| path-is-absolute | 1.0.1 |
| postcss-selector-parser | 6.0.10 |
| proc-log | 2.0.1 |
| promise-all-reject-late | 1.0.1 |
| promise-call-limit | 1.0.1 |
| promise-inflight | 1.0.1 |
| promise-retry | 2.0.1 |
| promzard | 0.3.0 |
| qrcode-terminal | 0.12.0 |
| read | 1.0.7 |
| read-cmd-shim | 3.0.0 |
| read-package-json | 5.0.2 |
| read-package-json-fast | 2.0.3 |
| readable-stream | 3.6.0 |
| readdir-scoped-modules | 1.1.0 |
| retry | 0.12.0 |
| rimraf | 3.0.2 |
| safe-buffer | 5.2.1 |
| safer-buffer | 2.1.2 |
| scanelf | 1.3.7-r1 |
| semver | 7.3.7 |
| set-blocking | 2.0.0 |
| signal-exit | 3.0.7 |
| smart-buffer | 4.2.0 |
| socks | 2.7.0 |
| socks-proxy-agent | 7.0.0 |
| spdx-correct | 3.1.1 |
| spdx-exceptions | 2.3.0 |
| spdx-expression-parse | 3.0.1 |
| spdx-license-ids | 3.0.11 |
| ssl_client | 1.36.1-r2 |
| ssri | 9.0.1 |
| string-width | 4.2.3 |
| string_decoder | 1.3.0 |
| strip-ansi | 6.0.1 |
| supports-color | 7.2.0 |
| tar | 6.1.11 |
| text-table | 0.2.0 |
| tiny-relative-date | 1.3.0 |
| treeverse | 2.0.0 |
| unique-filename | 2.0.1 |
| unique-slug | 3.0.0 |
| util-deprecate | 1.0.2 |
| validate-npm-package-license | 3.0.4 |
| validate-npm-package-name | 4.0.0 |
| walk-up-path | 1.0.0 |
| wcwidth | 1.0.1 |
| which | 2.0.2 |
| wide-align | 1.1.5 |
| wrappy | 1.0.2 |
| write-file-atomic | 4.0.2 |
| yallist | 4.0.0 |
| yarn | 1.22.19 |
| zlib | 1.2.13-r1 |
</details>

</details>

<details>
<summary><strong>🔍 Full Report: node:18-alpine</strong></summary>

## 🛡️ Security & Efficiency

**Base Image:** `linux (arm64)`
**Efficiency Score:** 100.0%

### Vulnerabilities
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| 🔴 2 | 🟠 18 | 20 | 9 |

<details>
<summary><strong>👇 Expand Vulnerability Details (49 found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [CVE-2025-15467](https://nvd.nist.gov/vuln/detail/CVE-2025-15467) | Critical | `libssl3` | `3.3.3-r0` |
| [CVE-2025-15467](https://nvd.nist.gov/vuln/detail/CVE-2025-15467) | Critical | `libcrypto3` | `3.3.3-r0` |
| [CVE-2025-23166](https://nvd.nist.gov/vuln/detail/CVE-2025-23166) | High | `node` | `18.20.8` |
| [CVE-2025-55131](https://nvd.nist.gov/vuln/detail/CVE-2025-55131) | High | `node` | `18.20.8` |
| [CVE-2025-59465](https://nvd.nist.gov/vuln/detail/CVE-2025-59465) | High | `node` | `18.20.8` |
| [CVE-2025-59466](https://nvd.nist.gov/vuln/detail/CVE-2025-59466) | High | `node` | `18.20.8` |
| [CVE-2025-69419](https://nvd.nist.gov/vuln/detail/CVE-2025-69419) | High | `libcrypto3` | `3.3.3-r0` |
| [CVE-2025-69419](https://nvd.nist.gov/vuln/detail/CVE-2025-69419) | High | `libssl3` | `3.3.3-r0` |
| [CVE-2025-69420](https://nvd.nist.gov/vuln/detail/CVE-2025-69420) | High | `libssl3` | `3.3.3-r0` |
| [CVE-2025-69420](https://nvd.nist.gov/vuln/detail/CVE-2025-69420) | High | `libcrypto3` | `3.3.3-r0` |
| [CVE-2025-69421](https://nvd.nist.gov/vuln/detail/CVE-2025-69421) | High | `libssl3` | `3.3.3-r0` |
| [CVE-2025-69421](https://nvd.nist.gov/vuln/detail/CVE-2025-69421) | High | `libcrypto3` | `3.3.3-r0` |
| [CVE-2025-9230](https://nvd.nist.gov/vuln/detail/CVE-2025-9230) | High | `libcrypto3` | `3.3.3-r0` |
| [CVE-2025-9230](https://nvd.nist.gov/vuln/detail/CVE-2025-9230) | High | `libssl3` | `3.3.3-r0` |
| [CVE-2026-21637](https://nvd.nist.gov/vuln/detail/CVE-2026-21637) | High | `node` | `18.20.8` |
| [GHSA-34x7-hfp2-rc4v](https://nvd.nist.gov/vuln/detail/GHSA-34x7-hfp2-rc4v) | High | `tar` | `6.2.1` |
| [GHSA-3xgq-45jj-v275](https://nvd.nist.gov/vuln/detail/GHSA-3xgq-45jj-v275) | High | `cross-spawn` | `7.0.3` |
| [GHSA-5j98-mcp5-4vw2](https://nvd.nist.gov/vuln/detail/GHSA-5j98-mcp5-4vw2) | High | `glob` | `10.4.2` |
| [GHSA-8qq5-rm4j-mr97](https://nvd.nist.gov/vuln/detail/GHSA-8qq5-rm4j-mr97) | High | `tar` | `6.2.1` |
| [GHSA-r6q2-hw4h-h46w](https://nvd.nist.gov/vuln/detail/GHSA-r6q2-hw4h-h46w) | High | `tar` | `6.2.1` |
| [CVE-2025-15468](https://nvd.nist.gov/vuln/detail/CVE-2025-15468) | Medium | `libcrypto3` | `3.3.3-r0` |
| [CVE-2025-15468](https://nvd.nist.gov/vuln/detail/CVE-2025-15468) | Medium | `libssl3` | `3.3.3-r0` |
| [CVE-2025-23167](https://nvd.nist.gov/vuln/detail/CVE-2025-23167) | Medium | `node` | `18.20.8` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `ssl_client` | `1.37.0-r12` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox-binsh` | `1.37.0-r12` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox` | `1.37.0-r12` |
| [CVE-2025-66199](https://nvd.nist.gov/vuln/detail/CVE-2025-66199) | Medium | `libssl3` | `3.3.3-r0` |
| [CVE-2025-66199](https://nvd.nist.gov/vuln/detail/CVE-2025-66199) | Medium | `libcrypto3` | `3.3.3-r0` |
| [CVE-2025-68160](https://nvd.nist.gov/vuln/detail/CVE-2025-68160) | Medium | `libcrypto3` | `3.3.3-r0` |
| [CVE-2025-68160](https://nvd.nist.gov/vuln/detail/CVE-2025-68160) | Medium | `libssl3` | `3.3.3-r0` |
| [CVE-2025-69418](https://nvd.nist.gov/vuln/detail/CVE-2025-69418) | Medium | `libcrypto3` | `3.3.3-r0` |
| [CVE-2025-69418](https://nvd.nist.gov/vuln/detail/CVE-2025-69418) | Medium | `libssl3` | `3.3.3-r0` |
| [CVE-2025-9231](https://nvd.nist.gov/vuln/detail/CVE-2025-9231) | Medium | `libcrypto3` | `3.3.3-r0` |
| [CVE-2025-9231](https://nvd.nist.gov/vuln/detail/CVE-2025-9231) | Medium | `libssl3` | `3.3.3-r0` |
| [CVE-2025-9232](https://nvd.nist.gov/vuln/detail/CVE-2025-9232) | Medium | `libssl3` | `3.3.3-r0` |
| [CVE-2025-9232](https://nvd.nist.gov/vuln/detail/CVE-2025-9232) | Medium | `libcrypto3` | `3.3.3-r0` |
| [CVE-2026-22795](https://nvd.nist.gov/vuln/detail/CVE-2026-22795) | Medium | `libcrypto3` | `3.3.3-r0` |
| [CVE-2026-22795](https://nvd.nist.gov/vuln/detail/CVE-2026-22795) | Medium | `libssl3` | `3.3.3-r0` |
| [CVE-2026-22796](https://nvd.nist.gov/vuln/detail/CVE-2026-22796) | Medium | `libcrypto3` | `3.3.3-r0` |
| [CVE-2026-22796](https://nvd.nist.gov/vuln/detail/CVE-2026-22796) | Medium | `libssl3` | `3.3.3-r0` |
| [CVE-2024-58251](https://nvd.nist.gov/vuln/detail/CVE-2024-58251) | Low | `busybox` | `1.37.0-r12` |
| [CVE-2024-58251](https://nvd.nist.gov/vuln/detail/CVE-2024-58251) | Low | `busybox-binsh` | `1.37.0-r12` |
| [CVE-2024-58251](https://nvd.nist.gov/vuln/detail/CVE-2024-58251) | Low | `ssl_client` | `1.37.0-r12` |
| [CVE-2025-23165](https://nvd.nist.gov/vuln/detail/CVE-2025-23165) | Low | `node` | `18.20.8` |
| [CVE-2025-46394](https://nvd.nist.gov/vuln/detail/CVE-2025-46394) | Low | `ssl_client` | `1.37.0-r12` |
| [CVE-2025-46394](https://nvd.nist.gov/vuln/detail/CVE-2025-46394) | Low | `busybox-binsh` | `1.37.0-r12` |
| [CVE-2025-46394](https://nvd.nist.gov/vuln/detail/CVE-2025-46394) | Low | `busybox` | `1.37.0-r12` |
| [GHSA-73rr-hh4g-fpgx](https://nvd.nist.gov/vuln/detail/GHSA-73rr-hh4g-fpgx) | Low | `diff` | `5.2.0` |
| [GHSA-v6h2-p8h4-qcjw](https://nvd.nist.gov/vuln/detail/GHSA-v6h2-p8h4-qcjw) | Low | `brace-expansion` | `2.0.1` |
</details>

<details>
<summary><strong>📦 Installed Packages (211 total)</strong></summary>

| Package | Version |
|---------|---------|
| @isaacs/cliui | 8.0.2 |
| @isaacs/string-locale-compare | 1.1.0 |
| @npmcli/agent | 2.2.2 |
| @npmcli/arborist | 7.5.4 |
| @npmcli/config | 8.3.4 |
| @npmcli/fs | 3.1.1 |
| @npmcli/git | 5.0.8 |
| @npmcli/installed-package-contents | 2.1.0 |
| @npmcli/map-workspaces | 3.0.6 |
| @npmcli/metavuln-calculator | 7.1.1 |
| @npmcli/name-from-folder | 2.0.0 |
| @npmcli/node-gyp | 3.0.0 |
| @npmcli/package-json | 5.2.0 |
| @npmcli/promise-spawn | 7.0.2 |
| @npmcli/query | 3.1.0 |
| @npmcli/redact | 2.0.1 |
| @npmcli/run-script | 8.1.0 |
| @pkgjs/parseargs | 0.11.0 |
| @sigstore/bundle | 2.3.2 |
| @sigstore/core | 1.1.0 |
| @sigstore/protobuf-specs | 0.3.2 |
| @sigstore/sign | 2.3.2 |
| @sigstore/tuf | 2.3.4 |
| @sigstore/verify | 1.2.1 |
| @tufjs/canonical-json | 2.0.0 |
| @tufjs/models | 2.0.1 |
| abbrev | 2.0.0 |
| agent-base | 7.1.1 |
| aggregate-error | 3.1.0 |
| alpine-baselayout | 3.6.8-r1 |
| alpine-baselayout-data | 3.6.8-r1 |
| alpine-keys | 2.5-r0 |
| alpine-release | 3.21.3-r0 |
| ansi-regex | 5.0.1 |
| ansi-regex | 6.0.1 |
| ansi-styles | 4.3.0 |
| ansi-styles | 6.2.1 |
| apk-tools | 2.14.6-r3 |
| aproba | 2.0.0 |
| archy | 1.0.0 |
| balanced-match | 1.0.2 |
| bin-links | 4.0.4 |
| binary-extensions | 2.3.0 |
| brace-expansion | 2.0.1 |
| busybox | 1.37.0-r12 |
| busybox-binsh | 1.37.0-r12 |
| ca-certificates-bundle | 20241121-r1 |
| cacache | 18.0.3 |
| chalk | 5.3.0 |
| chownr | 2.0.0 |
| ci-info | 4.0.0 |
| cidr-regex | 4.1.1 |
| clean-stack | 2.2.0 |
| cli-columns | 4.0.0 |
| cmd-shim | 6.0.3 |
| color-convert | 2.0.1 |
| color-name | 1.1.4 |
| common-ancestor-path | 1.0.1 |
| corepack | 0.32.0 |
| cross-spawn | 7.0.3 |
| cssesc | 3.0.0 |
| debug | 4.3.5 |
| diff | 5.2.0 |
| eastasianwidth | 0.2.0 |
| emoji-regex | 8.0.0 |
| emoji-regex | 9.2.2 |
| encoding | 0.1.13 |
| env-paths | 2.2.1 |
| err-code | 2.0.3 |
| exponential-backoff | 3.1.1 |
| fastest-levenshtein | 1.0.16 |
| foreground-child | 3.2.1 |
| fs-minipass | 2.1.0 |
| fs-minipass | 3.0.3 |
| glob | 10.4.2 |
| graceful-fs | 4.2.11 |
| hosted-git-info | 7.0.2 |
| http-cache-semantics | 4.1.1 |
| http-proxy-agent | 7.0.2 |
| https-proxy-agent | 7.0.5 |
| iconv-lite | 0.6.3 |
| ignore-walk | 6.0.5 |
| imurmurhash | 0.1.4 |
| indent-string | 4.0.0 |
| ini | 4.1.3 |
| init-package-json | 6.0.3 |
| ip-address | 9.0.5 |
| ip-regex | 5.0.0 |
| is-cidr | 5.1.0 |
| is-fullwidth-code-point | 3.0.0 |
| is-lambda | 1.0.1 |
| isexe | 2.0.0 |
| isexe | 3.1.1 |
| jackspeak | 3.4.0 |
| jsbn | 1.1.0 |
| json-parse-even-better-errors | 3.0.2 |
| json-stringify-nice | 1.1.4 |
| jsonparse | 1.3.1 |
| just-diff | 6.0.2 |
| just-diff-apply | 5.5.0 |
| libcrypto3 | 3.3.3-r0 |
| libgcc | 14.2.0-r4 |
| libnpmaccess | 8.0.6 |
| libnpmdiff | 6.1.4 |
| libnpmexec | 8.1.3 |
| libnpmfund | 5.0.12 |
| libnpmhook | 10.0.5 |
| libnpmorg | 6.0.6 |
| libnpmpack | 7.0.4 |
| libnpmpublish | 9.0.9 |
| libnpmsearch | 7.0.6 |
| libnpmteam | 6.0.5 |
| libnpmversion | 6.0.3 |
| libssl3 | 3.3.3-r0 |
| libstdc++ | 14.2.0-r4 |
| lru-cache | 10.2.2 |
| make-fetch-happen | 13.0.1 |
| minimatch | 9.0.5 |
| minipass | 3.3.6 |
| minipass | 5.0.0 |
| minipass | 7.1.2 |
| minipass-collect | 2.0.1 |
| minipass-fetch | 3.0.5 |
| minipass-flush | 1.0.5 |
| minipass-pipeline | 1.2.4 |
| minipass-sized | 1.0.3 |
| minizlib | 2.1.2 |
| mkdirp | 1.0.4 |
| ms | 2.1.2 |
| ms | 2.1.3 |
| musl | 1.2.5-r9 |
| musl-utils | 1.2.5-r9 |
| mute-stream | 1.0.0 |
| negotiator | 0.6.3 |
| node | 18.20.8 |
| node-gyp | 10.1.0 |
| nopt | 7.2.1 |
| normalize-package-data | 6.0.2 |
| npm | 10.8.2 |
| npm-audit-report | 5.0.0 |
| npm-bundled | 3.0.1 |
| npm-install-checks | 6.3.0 |
| npm-normalize-package-bin | 3.0.1 |
| npm-package-arg | 11.0.2 |
| npm-packlist | 8.0.2 |
| npm-pick-manifest | 9.1.0 |
| npm-profile | 10.0.0 |
| npm-registry-fetch | 17.1.0 |
| npm-user-validate | 2.0.1 |
| p-map | 4.0.0 |
| package-json-from-dist | 1.0.0 |
| pacote | 18.0.6 |
| parse-conflict-json | 3.0.1 |
| path-key | 3.1.1 |
| path-scurry | 1.11.1 |
| postcss-selector-parser | 6.1.0 |
| proc-log | 3.0.0 |
| proc-log | 4.2.0 |
| proggy | 2.0.0 |
| promise-all-reject-late | 1.0.1 |
| promise-call-limit | 3.0.1 |
| promise-inflight | 1.0.1 |
| promise-retry | 2.0.1 |
| promzard | 1.0.2 |
| qrcode-terminal | 0.12.0 |
| read | 3.0.1 |
| read-cmd-shim | 4.0.0 |
| read-package-json-fast | 3.0.2 |
| retry | 0.12.0 |
| safer-buffer | 2.1.2 |
| scanelf | 1.3.8-r1 |
| semver | 7.6.2 |
| shebang-command | 2.0.0 |
| shebang-regex | 3.0.0 |
| signal-exit | 4.1.0 |
| sigstore | 2.3.1 |
| smart-buffer | 4.2.0 |
| socks | 2.8.3 |
| socks-proxy-agent | 8.0.4 |
| spdx-correct | 3.2.0 |
| spdx-exceptions | 2.5.0 |
| spdx-expression-parse | 3.0.1 |
| spdx-expression-parse | 4.0.0 |
| spdx-license-ids | 3.0.18 |
| sprintf-js | 1.1.3 |
| ssl_client | 1.37.0-r12 |
| ssri | 10.0.6 |
| string-width | 4.2.3 |
| string-width | 5.1.2 |
| strip-ansi | 6.0.1 |
| strip-ansi | 7.1.0 |
| supports-color | 9.4.0 |
| tar | 6.2.1 |
| text-table | 0.2.0 |
| tiny-relative-date | 1.3.0 |
| treeverse | 3.0.0 |
| tuf-js | 2.2.1 |
| unique-filename | 3.0.0 |
| unique-slug | 4.0.0 |
| util-deprecate | 1.0.2 |
| validate-npm-package-license | 3.0.4 |
| validate-npm-package-name | 5.0.1 |
| walk-up-path | 3.0.1 |
| which | 2.0.2 |
| which | 4.0.0 |
| wrap-ansi | 7.0.0 |
| wrap-ansi | 8.1.0 |
| write-file-atomic | 5.0.1 |
| yallist | 4.0.0 |
| yarn | 1.22.22 |
| zlib | 1.3.1-r2 |
</details>

</details>

<details>
<summary><strong>🔍 Full Report: node:20-alpine</strong></summary>

## 🛡️ Security & Efficiency

**Base Image:** `linux (arm64)`
**Efficiency Score:** 100.0%

### Vulnerabilities
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| 🟢 0 | 🟠 5 | 3 | 2 |

<details>
<summary><strong>👇 Expand Vulnerability Details (10 found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [GHSA-34x7-hfp2-rc4v](https://nvd.nist.gov/vuln/detail/GHSA-34x7-hfp2-rc4v) | High | `tar` | `6.2.1` |
| [GHSA-3xgq-45jj-v275](https://nvd.nist.gov/vuln/detail/GHSA-3xgq-45jj-v275) | High | `cross-spawn` | `7.0.3` |
| [GHSA-5j98-mcp5-4vw2](https://nvd.nist.gov/vuln/detail/GHSA-5j98-mcp5-4vw2) | High | `glob` | `10.4.2` |
| [GHSA-8qq5-rm4j-mr97](https://nvd.nist.gov/vuln/detail/GHSA-8qq5-rm4j-mr97) | High | `tar` | `6.2.1` |
| [GHSA-r6q2-hw4h-h46w](https://nvd.nist.gov/vuln/detail/GHSA-r6q2-hw4h-h46w) | High | `tar` | `6.2.1` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox` | `1.37.0-r30` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox-binsh` | `1.37.0-r30` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `ssl_client` | `1.37.0-r30` |
| [GHSA-73rr-hh4g-fpgx](https://nvd.nist.gov/vuln/detail/GHSA-73rr-hh4g-fpgx) | Low | `diff` | `5.2.0` |
| [GHSA-v6h2-p8h4-qcjw](https://nvd.nist.gov/vuln/detail/GHSA-v6h2-p8h4-qcjw) | Low | `brace-expansion` | `2.0.1` |
</details>

<details>
<summary><strong>📦 Installed Packages (212 total)</strong></summary>

| Package | Version |
|---------|---------|
| @isaacs/cliui | 8.0.2 |
| @isaacs/string-locale-compare | 1.1.0 |
| @npmcli/agent | 2.2.2 |
| @npmcli/arborist | 7.5.4 |
| @npmcli/config | 8.3.4 |
| @npmcli/fs | 3.1.1 |
| @npmcli/git | 5.0.8 |
| @npmcli/installed-package-contents | 2.1.0 |
| @npmcli/map-workspaces | 3.0.6 |
| @npmcli/metavuln-calculator | 7.1.1 |
| @npmcli/name-from-folder | 2.0.0 |
| @npmcli/node-gyp | 3.0.0 |
| @npmcli/package-json | 5.2.0 |
| @npmcli/promise-spawn | 7.0.2 |
| @npmcli/query | 3.1.0 |
| @npmcli/redact | 2.0.1 |
| @npmcli/run-script | 8.1.0 |
| @pkgjs/parseargs | 0.11.0 |
| @sigstore/bundle | 2.3.2 |
| @sigstore/core | 1.1.0 |
| @sigstore/protobuf-specs | 0.3.2 |
| @sigstore/sign | 2.3.2 |
| @sigstore/tuf | 2.3.4 |
| @sigstore/verify | 1.2.1 |
| @tufjs/canonical-json | 2.0.0 |
| @tufjs/models | 2.0.1 |
| abbrev | 2.0.0 |
| agent-base | 7.1.1 |
| aggregate-error | 3.1.0 |
| alpine-baselayout | 3.7.1-r8 |
| alpine-baselayout-data | 3.7.1-r8 |
| alpine-keys | 2.6-r0 |
| alpine-release | 3.23.3-r0 |
| ansi-regex | 5.0.1 |
| ansi-regex | 6.0.1 |
| ansi-styles | 4.3.0 |
| ansi-styles | 6.2.1 |
| apk-tools | 3.0.3-r1 |
| aproba | 2.0.0 |
| archy | 1.0.0 |
| balanced-match | 1.0.2 |
| bin-links | 4.0.4 |
| binary-extensions | 2.3.0 |
| brace-expansion | 2.0.1 |
| busybox | 1.37.0-r30 |
| busybox-binsh | 1.37.0-r30 |
| ca-certificates-bundle | 20251003-r0 |
| cacache | 18.0.3 |
| chalk | 5.3.0 |
| chownr | 2.0.0 |
| ci-info | 4.0.0 |
| cidr-regex | 4.1.1 |
| clean-stack | 2.2.0 |
| cli-columns | 4.0.0 |
| cmd-shim | 6.0.3 |
| color-convert | 2.0.1 |
| color-name | 1.1.4 |
| common-ancestor-path | 1.0.1 |
| corepack | 0.34.1 |
| cross-spawn | 7.0.3 |
| cssesc | 3.0.0 |
| debug | 4.3.5 |
| diff | 5.2.0 |
| eastasianwidth | 0.2.0 |
| emoji-regex | 8.0.0 |
| emoji-regex | 9.2.2 |
| encoding | 0.1.13 |
| env-paths | 2.2.1 |
| err-code | 2.0.3 |
| exponential-backoff | 3.1.1 |
| fastest-levenshtein | 1.0.16 |
| foreground-child | 3.2.1 |
| fs-minipass | 2.1.0 |
| fs-minipass | 3.0.3 |
| glob | 10.4.2 |
| graceful-fs | 4.2.11 |
| hosted-git-info | 7.0.2 |
| http-cache-semantics | 4.1.1 |
| http-proxy-agent | 7.0.2 |
| https-proxy-agent | 7.0.5 |
| iconv-lite | 0.6.3 |
| ignore-walk | 6.0.5 |
| imurmurhash | 0.1.4 |
| indent-string | 4.0.0 |
| ini | 4.1.3 |
| init-package-json | 6.0.3 |
| ip-address | 9.0.5 |
| ip-regex | 5.0.0 |
| is-cidr | 5.1.0 |
| is-fullwidth-code-point | 3.0.0 |
| is-lambda | 1.0.1 |
| isexe | 2.0.0 |
| isexe | 3.1.1 |
| jackspeak | 3.4.0 |
| jsbn | 1.1.0 |
| json-parse-even-better-errors | 3.0.2 |
| json-stringify-nice | 1.1.4 |
| jsonparse | 1.3.1 |
| just-diff | 6.0.2 |
| just-diff-apply | 5.5.0 |
| libapk | 3.0.3-r1 |
| libcrypto3 | 3.5.5-r0 |
| libgcc | 15.2.0-r2 |
| libnpmaccess | 8.0.6 |
| libnpmdiff | 6.1.4 |
| libnpmexec | 8.1.3 |
| libnpmfund | 5.0.12 |
| libnpmhook | 10.0.5 |
| libnpmorg | 6.0.6 |
| libnpmpack | 7.0.4 |
| libnpmpublish | 9.0.9 |
| libnpmsearch | 7.0.6 |
| libnpmteam | 6.0.5 |
| libnpmversion | 6.0.3 |
| libssl3 | 3.5.5-r0 |
| libstdc++ | 15.2.0-r2 |
| lru-cache | 10.2.2 |
| make-fetch-happen | 13.0.1 |
| minimatch | 9.0.5 |
| minipass | 3.3.6 |
| minipass | 5.0.0 |
| minipass | 7.1.2 |
| minipass-collect | 2.0.1 |
| minipass-fetch | 3.0.5 |
| minipass-flush | 1.0.5 |
| minipass-pipeline | 1.2.4 |
| minipass-sized | 1.0.3 |
| minizlib | 2.1.2 |
| mkdirp | 1.0.4 |
| ms | 2.1.2 |
| ms | 2.1.3 |
| musl | 1.2.5-r21 |
| musl-utils | 1.2.5-r21 |
| mute-stream | 1.0.0 |
| negotiator | 0.6.3 |
| node | 20.20.0 |
| node-gyp | 10.1.0 |
| nopt | 7.2.1 |
| normalize-package-data | 6.0.2 |
| npm | 10.8.2 |
| npm-audit-report | 5.0.0 |
| npm-bundled | 3.0.1 |
| npm-install-checks | 6.3.0 |
| npm-normalize-package-bin | 3.0.1 |
| npm-package-arg | 11.0.2 |
| npm-packlist | 8.0.2 |
| npm-pick-manifest | 9.1.0 |
| npm-profile | 10.0.0 |
| npm-registry-fetch | 17.1.0 |
| npm-user-validate | 2.0.1 |
| p-map | 4.0.0 |
| package-json-from-dist | 1.0.0 |
| pacote | 18.0.6 |
| parse-conflict-json | 3.0.1 |
| path-key | 3.1.1 |
| path-scurry | 1.11.1 |
| postcss-selector-parser | 6.1.0 |
| proc-log | 3.0.0 |
| proc-log | 4.2.0 |
| proggy | 2.0.0 |
| promise-all-reject-late | 1.0.1 |
| promise-call-limit | 3.0.1 |
| promise-inflight | 1.0.1 |
| promise-retry | 2.0.1 |
| promzard | 1.0.2 |
| qrcode-terminal | 0.12.0 |
| read | 3.0.1 |
| read-cmd-shim | 4.0.0 |
| read-package-json-fast | 3.0.2 |
| retry | 0.12.0 |
| safer-buffer | 2.1.2 |
| scanelf | 1.3.8-r2 |
| semver | 7.6.2 |
| shebang-command | 2.0.0 |
| shebang-regex | 3.0.0 |
| signal-exit | 4.1.0 |
| sigstore | 2.3.1 |
| smart-buffer | 4.2.0 |
| socks | 2.8.3 |
| socks-proxy-agent | 8.0.4 |
| spdx-correct | 3.2.0 |
| spdx-exceptions | 2.5.0 |
| spdx-expression-parse | 3.0.1 |
| spdx-expression-parse | 4.0.0 |
| spdx-license-ids | 3.0.18 |
| sprintf-js | 1.1.3 |
| ssl_client | 1.37.0-r30 |
| ssri | 10.0.6 |
| string-width | 4.2.3 |
| string-width | 5.1.2 |
| strip-ansi | 6.0.1 |
| strip-ansi | 7.1.0 |
| supports-color | 9.4.0 |
| tar | 6.2.1 |
| text-table | 0.2.0 |
| tiny-relative-date | 1.3.0 |
| treeverse | 3.0.0 |
| tuf-js | 2.2.1 |
| unique-filename | 3.0.0 |
| unique-slug | 4.0.0 |
| util-deprecate | 1.0.2 |
| validate-npm-package-license | 3.0.4 |
| validate-npm-package-name | 5.0.1 |
| walk-up-path | 3.0.1 |
| which | 2.0.2 |
| which | 4.0.0 |
| wrap-ansi | 7.0.0 |
| wrap-ansi | 8.1.0 |
| write-file-atomic | 5.0.1 |
| yallist | 4.0.0 |
| yarn | 1.22.22 |
| zlib | 1.3.1-r2 |
</details>

</details>

<!-- END: docker-docs:comparison -->
