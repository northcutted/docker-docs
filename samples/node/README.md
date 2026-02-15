# Node.js Docker Analysis

<!-- BEGIN: docker-docs:config -->

# 🐳 Docker Image Analysis: node:18-alpine
![Size]() ![Layers]() ![Vulns](https://img.shields.io/static/v1?label=Security&message=49+Vulns+%282+Crit%29&color=red) ![Efficiency]()

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

**Base Image:** ` ()`
**Efficiency Score:** 0.0%

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
| `node:16-alpine` | ![Size]() | ![Vulns](https://img.shields.io/static/v1?label=Security&message=76+Vulns+%282+Crit%29&color=red) | 0.0% | / |
| `node:18-alpine` | ![Size](https://img.shields.io/static/v1?label=Size&message=122.10+MB&color=blue) | ![Vulns](https://img.shields.io/static/v1?label=Security&message=49+Vulns+%282+Crit%29&color=red) | 100.0% | linux/arm64 |
| `node:20-alpine` | ![Size]() | ![Vulns](https://img.shields.io/static/v1?label=Security&message=10+Vulns+%280+Crit%29&color=orange) | 0.0% | / |

<!-- END: docker-docs:comparison -->
