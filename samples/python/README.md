# Python Docker Analysis

<!-- BEGIN: docker-docs:config -->

# 🐳 Docker Image Analysis: python:3.9-slim
![Size](https://img.shields.io/static/v1?label=Size&message=144.71+MB&color=blue) ![Layers](https://img.shields.io/static/v1?label=Layers&message=4&color=blue) ![Vulns](https://img.shields.io/static/v1?label=Security&message=77+Vulns+%283+Crit%29&color=red) ![Efficiency](https://img.shields.io/static/v1?label=Efficiency&message=97.3%&color=green)

## ⚙️ Configuration
### Environment Variables
| Name | Description | Default | Required |
|------|-------------|---------|:--------:|
| `NAME` | The name to greet | `World` | ❌ |
### Exposed Ports
| Port | Description |
|------|-------------|
| `80` | The port where the FastAPI application listens |
---

## 🛡️ Security & Efficiency

**Base Image:** `linux (arm64)`
**Supported Architectures:** `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown`
**Efficiency Score:** 97.3%

### Vulnerabilities
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| 🔴 3 | 🟠 20 | 51 | 3 |

<details>
<summary><strong>👇 Expand Vulnerability Details (117 found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [CVE-2025-15467](https://nvd.nist.gov/vuln/detail/CVE-2025-15467) | Critical | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2025-15467](https://nvd.nist.gov/vuln/detail/CVE-2025-15467) | Critical | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2025-15467](https://nvd.nist.gov/vuln/detail/CVE-2025-15467) | Critical | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [CVE-2023-36632](https://nvd.nist.gov/vuln/detail/CVE-2023-36632) | High | `python` | `3.9.25` |
| [CVE-2025-13836](https://nvd.nist.gov/vuln/detail/CVE-2025-13836) | High | `python` | `3.9.25` |
| [CVE-2025-15281](https://nvd.nist.gov/vuln/detail/CVE-2025-15281) | High | `libc-bin` | `2.41-12` |
| [CVE-2025-15281](https://nvd.nist.gov/vuln/detail/CVE-2025-15281) | High | `libc6` | `2.41-12` |
| [CVE-2025-69419](https://nvd.nist.gov/vuln/detail/CVE-2025-69419) | High | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [CVE-2025-69419](https://nvd.nist.gov/vuln/detail/CVE-2025-69419) | High | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2025-69419](https://nvd.nist.gov/vuln/detail/CVE-2025-69419) | High | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2025-69420](https://nvd.nist.gov/vuln/detail/CVE-2025-69420) | High | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2025-69420](https://nvd.nist.gov/vuln/detail/CVE-2025-69420) | High | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2025-69420](https://nvd.nist.gov/vuln/detail/CVE-2025-69420) | High | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [CVE-2025-69421](https://nvd.nist.gov/vuln/detail/CVE-2025-69421) | High | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2025-69421](https://nvd.nist.gov/vuln/detail/CVE-2025-69421) | High | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [CVE-2025-69421](https://nvd.nist.gov/vuln/detail/CVE-2025-69421) | High | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2026-0861](https://nvd.nist.gov/vuln/detail/CVE-2026-0861) | High | `libc-bin` | `2.41-12` |
| [CVE-2026-0861](https://nvd.nist.gov/vuln/detail/CVE-2026-0861) | High | `libc6` | `2.41-12` |
| [CVE-2026-0915](https://nvd.nist.gov/vuln/detail/CVE-2026-0915) | High | `libc6` | `2.41-12` |
| [CVE-2026-0915](https://nvd.nist.gov/vuln/detail/CVE-2026-0915) | High | `libc-bin` | `2.41-12` |
| [GHSA-58pv-8j8x-9vj2](https://nvd.nist.gov/vuln/detail/GHSA-58pv-8j8x-9vj2) | High | `jaraco-context` | `5.3.0` |
| [GHSA-8rrh-rw8j-w5fx](https://nvd.nist.gov/vuln/detail/GHSA-8rrh-rw8j-w5fx) | High | `wheel` | `0.45.1` |
| [GHSA-8rrh-rw8j-w5fx](https://nvd.nist.gov/vuln/detail/GHSA-8rrh-rw8j-w5fx) | High | `wheel` | `0.45.1` |
| [CVE-2025-11187](https://nvd.nist.gov/vuln/detail/CVE-2025-11187) | Medium | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2025-11187](https://nvd.nist.gov/vuln/detail/CVE-2025-11187) | Medium | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [CVE-2025-11187](https://nvd.nist.gov/vuln/detail/CVE-2025-11187) | Medium | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2025-11468](https://nvd.nist.gov/vuln/detail/CVE-2025-11468) | Medium | `python` | `3.9.25` |
| [CVE-2025-12084](https://nvd.nist.gov/vuln/detail/CVE-2025-12084) | Medium | `python` | `3.9.25` |
| [CVE-2025-12781](https://nvd.nist.gov/vuln/detail/CVE-2025-12781) | Medium | `python` | `3.9.25` |
| [CVE-2025-13837](https://nvd.nist.gov/vuln/detail/CVE-2025-13837) | Medium | `python` | `3.9.25` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `util-linux` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `login` | `1:4.16.0-2+really2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `liblastlog2-2` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libblkid1` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `bsdutils` | `1:2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libsmartcols1` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libuuid1` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libmount1` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `mount` | `2.41-5` |
| [CVE-2025-15282](https://nvd.nist.gov/vuln/detail/CVE-2025-15282) | Medium | `python` | `3.9.25` |
| [CVE-2025-15366](https://nvd.nist.gov/vuln/detail/CVE-2025-15366) | Medium | `python` | `3.9.25` |
| [CVE-2025-15367](https://nvd.nist.gov/vuln/detail/CVE-2025-15367) | Medium | `python` | `3.9.25` |
| [CVE-2025-15468](https://nvd.nist.gov/vuln/detail/CVE-2025-15468) | Medium | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2025-15468](https://nvd.nist.gov/vuln/detail/CVE-2025-15468) | Medium | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2025-15468](https://nvd.nist.gov/vuln/detail/CVE-2025-15468) | Medium | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [CVE-2025-15469](https://nvd.nist.gov/vuln/detail/CVE-2025-15469) | Medium | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2025-15469](https://nvd.nist.gov/vuln/detail/CVE-2025-15469) | Medium | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [CVE-2025-15469](https://nvd.nist.gov/vuln/detail/CVE-2025-15469) | Medium | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2025-6075](https://nvd.nist.gov/vuln/detail/CVE-2025-6075) | Medium | `python` | `3.9.25` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `ncurses-bin` | `6.5+20250216-2` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libncursesw6` | `6.5+20250216-2` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libtinfo6` | `6.5+20250216-2` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `ncurses-base` | `6.5+20250216-2` |
| [CVE-2025-66199](https://nvd.nist.gov/vuln/detail/CVE-2025-66199) | Medium | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2025-66199](https://nvd.nist.gov/vuln/detail/CVE-2025-66199) | Medium | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2025-66199](https://nvd.nist.gov/vuln/detail/CVE-2025-66199) | Medium | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [CVE-2025-68160](https://nvd.nist.gov/vuln/detail/CVE-2025-68160) | Medium | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2025-68160](https://nvd.nist.gov/vuln/detail/CVE-2025-68160) | Medium | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2025-68160](https://nvd.nist.gov/vuln/detail/CVE-2025-68160) | Medium | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [CVE-2025-69418](https://nvd.nist.gov/vuln/detail/CVE-2025-69418) | Medium | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [CVE-2025-69418](https://nvd.nist.gov/vuln/detail/CVE-2025-69418) | Medium | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2025-69418](https://nvd.nist.gov/vuln/detail/CVE-2025-69418) | Medium | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2025-7709](https://nvd.nist.gov/vuln/detail/CVE-2025-7709) | Medium | `libsqlite3-0` | `3.46.1-7` |
| [CVE-2026-0672](https://nvd.nist.gov/vuln/detail/CVE-2026-0672) | Medium | `python` | `3.9.25` |
| [CVE-2026-0865](https://nvd.nist.gov/vuln/detail/CVE-2026-0865) | Medium | `python` | `3.9.25` |
| [CVE-2026-1299](https://nvd.nist.gov/vuln/detail/CVE-2026-1299) | Medium | `python` | `3.9.25` |
| [CVE-2026-22795](https://nvd.nist.gov/vuln/detail/CVE-2026-22795) | Medium | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2026-22795](https://nvd.nist.gov/vuln/detail/CVE-2026-22795) | Medium | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2026-22795](https://nvd.nist.gov/vuln/detail/CVE-2026-22795) | Medium | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [CVE-2026-22796](https://nvd.nist.gov/vuln/detail/CVE-2026-22796) | Medium | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2026-22796](https://nvd.nist.gov/vuln/detail/CVE-2026-22796) | Medium | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2026-22796](https://nvd.nist.gov/vuln/detail/CVE-2026-22796) | Medium | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [GHSA-4xh5-x5gv-qwph](https://nvd.nist.gov/vuln/detail/GHSA-4xh5-x5gv-qwph) | Medium | `pip` | `23.0.1` |
| [GHSA-mq26-g339-26xf](https://nvd.nist.gov/vuln/detail/GHSA-mq26-g339-26xf) | Medium | `pip` | `23.0.1` |
| [CVE-2024-56433](https://nvd.nist.gov/vuln/detail/CVE-2024-56433) | Low | `login.defs` | `1:4.17.4-2` |
| [CVE-2024-56433](https://nvd.nist.gov/vuln/detail/CVE-2024-56433) | Low | `passwd` | `1:4.17.4-2` |
| [GHSA-6vgw-5pg2-w6jp](https://nvd.nist.gov/vuln/detail/GHSA-6vgw-5pg2-w6jp) | Low | `pip` | `23.0.1` |
| [CVE-2005-2541](https://nvd.nist.gov/vuln/detail/CVE-2005-2541) | Negligible | `tar` | `1.35+dfsg-3.1` |
| [CVE-2007-5686](https://nvd.nist.gov/vuln/detail/CVE-2007-5686) | Negligible | `login.defs` | `1:4.17.4-2` |
| [CVE-2007-5686](https://nvd.nist.gov/vuln/detail/CVE-2007-5686) | Negligible | `passwd` | `1:4.17.4-2` |
| [CVE-2010-4756](https://nvd.nist.gov/vuln/detail/CVE-2010-4756) | Negligible | `libc6` | `2.41-12` |
| [CVE-2010-4756](https://nvd.nist.gov/vuln/detail/CVE-2010-4756) | Negligible | `libc-bin` | `2.41-12` |
| [CVE-2011-3374](https://nvd.nist.gov/vuln/detail/CVE-2011-3374) | Negligible | `libapt-pkg7.0` | `3.0.3` |
| [CVE-2011-3374](https://nvd.nist.gov/vuln/detail/CVE-2011-3374) | Negligible | `apt` | `3.0.3` |
| [CVE-2011-4116](https://nvd.nist.gov/vuln/detail/CVE-2011-4116) | Negligible | `perl-base` | `5.40.1-6` |
| [CVE-2013-4392](https://nvd.nist.gov/vuln/detail/CVE-2013-4392) | Negligible | `libsystemd0` | `257.8-1~deb13u2` |
| [CVE-2013-4392](https://nvd.nist.gov/vuln/detail/CVE-2013-4392) | Negligible | `libudev1` | `257.8-1~deb13u2` |
| [CVE-2017-18018](https://nvd.nist.gov/vuln/detail/CVE-2017-18018) | Negligible | `coreutils` | `9.7-3` |
| [CVE-2018-20796](https://nvd.nist.gov/vuln/detail/CVE-2018-20796) | Negligible | `libc6` | `2.41-12` |
| [CVE-2018-20796](https://nvd.nist.gov/vuln/detail/CVE-2018-20796) | Negligible | `libc-bin` | `2.41-12` |
| [CVE-2019-1010022](https://nvd.nist.gov/vuln/detail/CVE-2019-1010022) | Negligible | `libc6` | `2.41-12` |
| [CVE-2019-1010022](https://nvd.nist.gov/vuln/detail/CVE-2019-1010022) | Negligible | `libc-bin` | `2.41-12` |
| [CVE-2019-1010023](https://nvd.nist.gov/vuln/detail/CVE-2019-1010023) | Negligible | `libc6` | `2.41-12` |
| [CVE-2019-1010023](https://nvd.nist.gov/vuln/detail/CVE-2019-1010023) | Negligible | `libc-bin` | `2.41-12` |
| [CVE-2019-1010024](https://nvd.nist.gov/vuln/detail/CVE-2019-1010024) | Negligible | `libc6` | `2.41-12` |
| [CVE-2019-1010024](https://nvd.nist.gov/vuln/detail/CVE-2019-1010024) | Negligible | `libc-bin` | `2.41-12` |
| [CVE-2019-1010025](https://nvd.nist.gov/vuln/detail/CVE-2019-1010025) | Negligible | `libc-bin` | `2.41-12` |
| [CVE-2019-1010025](https://nvd.nist.gov/vuln/detail/CVE-2019-1010025) | Negligible | `libc6` | `2.41-12` |
| [CVE-2019-9192](https://nvd.nist.gov/vuln/detail/CVE-2019-9192) | Negligible | `libc-bin` | `2.41-12` |
| [CVE-2019-9192](https://nvd.nist.gov/vuln/detail/CVE-2019-9192) | Negligible | `libc6` | `2.41-12` |
| [CVE-2021-45346](https://nvd.nist.gov/vuln/detail/CVE-2021-45346) | Negligible | `libsqlite3-0` | `3.46.1-7` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libblkid1` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libmount1` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `util-linux` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `mount` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `login` | `1:4.16.0-2+really2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libuuid1` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libsmartcols1` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `bsdutils` | `1:2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `liblastlog2-2` | `2.41-5` |
| [CVE-2023-31437](https://nvd.nist.gov/vuln/detail/CVE-2023-31437) | Negligible | `libudev1` | `257.8-1~deb13u2` |
| [CVE-2023-31437](https://nvd.nist.gov/vuln/detail/CVE-2023-31437) | Negligible | `libsystemd0` | `257.8-1~deb13u2` |
| [CVE-2023-31438](https://nvd.nist.gov/vuln/detail/CVE-2023-31438) | Negligible | `libsystemd0` | `257.8-1~deb13u2` |
| [CVE-2023-31438](https://nvd.nist.gov/vuln/detail/CVE-2023-31438) | Negligible | `libudev1` | `257.8-1~deb13u2` |
| [CVE-2023-31439](https://nvd.nist.gov/vuln/detail/CVE-2023-31439) | Negligible | `libudev1` | `257.8-1~deb13u2` |
| [CVE-2023-31439](https://nvd.nist.gov/vuln/detail/CVE-2023-31439) | Negligible | `libsystemd0` | `257.8-1~deb13u2` |
| [CVE-2025-5278](https://nvd.nist.gov/vuln/detail/CVE-2025-5278) | Negligible | `coreutils` | `9.7-3` |
</details>

<details>
<summary><strong>📦 Installed Packages (115 total)</strong></summary>

| Package | Version |
|---------|---------|
| Simple Launcher | 1.1.0.14 |
| adduser | 3.152 |
| apt | 3.0.3 |
| autocommand | 2.2.2 |
| backports-tarfile | 1.2.0 |
| base-files | 13.8+deb13u1 |
| base-passwd | 3.6.7 |
| bash | 5.2.37-2+b5 |
| bsdutils | 1:2.41-5 |
| ca-certificates | 20250419 |
| cli | UNKNOWN |
| cli-32 | UNKNOWN |
| cli-64 | UNKNOWN |
| cli-arm64 | UNKNOWN |
| coreutils | 9.7-3 |
| dash | 0.5.12-12 |
| debconf | 1.5.91 |
| debian-archive-keyring | 2025.1 |
| debianutils | 5.23.2 |
| diffutils | 1:3.10-4 |
| dpkg | 1.22.21 |
| findutils | 4.10.0-3 |
| gcc-14-base | 14.2.0-19 |
| grep | 3.11-4+b1 |
| gui | UNKNOWN |
| gui-32 | UNKNOWN |
| gui-64 | UNKNOWN |
| gui-arm64 | UNKNOWN |
| gzip | 1.13-1 |
| hostname | 3.25 |
| importlib-metadata | 8.0.0 |
| inflect | 7.3.1 |
| init-system-helpers | 1.69~deb13u1 |
| jaraco-collections | 5.1.0 |
| jaraco-context | 5.3.0 |
| jaraco-functools | 4.0.1 |
| jaraco-text | 3.12.1 |
| libacl1 | 2.3.2-2+b1 |
| libapt-pkg7.0 | 3.0.3 |
| libattr1 | 1:2.5.2-3 |
| libaudit-common | 1:4.0.2-2 |
| libaudit1 | 1:4.0.2-2+b2 |
| libblkid1 | 2.41-5 |
| libbsd0 | 0.12.2-2 |
| libbz2-1.0 | 1.0.8-6 |
| libc-bin | 2.41-12 |
| libc6 | 2.41-12 |
| libcap-ng0 | 0.8.5-4+b1 |
| libcap2 | 1:2.75-10+b1 |
| libcrypt1 | 1:4.4.38-1 |
| libdb5.3t64 | 5.3.28+dfsg2-9 |
| libdebconfclient0 | 0.280 |
| libffi8 | 3.4.8-2 |
| libgcc-s1 | 14.2.0-19 |
| libgdbm6t64 | 1.24-2 |
| libgmp10 | 2:6.3.0+dfsg-3 |
| libhogweed6t64 | 3.10.1-1 |
| liblastlog2-2 | 2.41-5 |
| liblz4-1 | 1.10.0-4 |
| liblzma5 | 5.8.1-1 |
| libmd0 | 1.1.0-2+b1 |
| libmount1 | 2.41-5 |
| libncursesw6 | 6.5+20250216-2 |
| libnettle8t64 | 3.10.1-1 |
| libpam-modules | 1.7.0-5 |
| libpam-modules-bin | 1.7.0-5 |
| libpam-runtime | 1.7.0-5 |
| libpam0g | 1.7.0-5 |
| libpcre2-8-0 | 10.46-1~deb13u1 |
| libreadline8t64 | 8.2-6 |
| libseccomp2 | 2.6.0-2 |
| libselinux1 | 3.8.1-1 |
| libsemanage-common | 3.8.1-1 |
| libsemanage2 | 3.8.1-1 |
| libsepol2 | 3.8.1-1 |
| libsmartcols1 | 2.41-5 |
| libsqlite3-0 | 3.46.1-7 |
| libssl3t64 | 3.5.1-1+deb13u1 |
| libstdc++6 | 14.2.0-19 |
| libsystemd0 | 257.8-1~deb13u2 |
| libtinfo6 | 6.5+20250216-2 |
| libudev1 | 257.8-1~deb13u2 |
| libuuid1 | 2.41-5 |
| libxxhash0 | 0.8.3-2 |
| libzstd1 | 1.5.7+dfsg-1 |
| login | 1:4.16.0-2+really2.41-5 |
| login.defs | 1:4.17.4-2 |
| mawk | 1.3.4.20250131-1 |
| more-itertools | 10.3.0 |
| mount | 2.41-5 |
| ncurses-base | 6.5+20250216-2 |
| ncurses-bin | 6.5+20250216-2 |
| netbase | 6.5 |
| openssl | 3.5.1-1+deb13u1 |
| openssl-provider-legacy | 3.5.1-1+deb13u1 |
| packaging | 24.2 |
| passwd | 1:4.17.4-2 |
| perl-base | 5.40.1-6 |
| pip | 23.0.1 |
| platformdirs | 4.2.2 |
| python | 3.9.25 |
| readline-common | 8.2-6 |
| sed | 4.9-2+b1 |
| setuptools | 79.0.1 |
| sqv | 1.3.0-3 |
| sysvinit-utils | 3.14-4 |
| tar | 1.35+dfsg-3.1 |
| tomli | 2.0.1 |
| typeguard | 4.3.0 |
| typing-extensions | 4.12.2 |
| tzdata | 2025b-4+deb13u1 |
| util-linux | 2.41-5 |
| wheel | 0.45.1 |
| zipp | 3.19.2 |
| zlib1g | 1:1.3.dfsg+really1.3.1-1+b1 |
</details>

<!-- END: docker-docs:config -->

## Version Comparison

<!-- BEGIN: docker-docs:comparison -->

### 🏷️ Supported Tags

| Tag | Size | Vulns | Efficiency | Architectures |
|-----|------|-------|------------|---------------|
| `python:3.9-slim` | ![Size](https://img.shields.io/static/v1?label=Size&message=144.71+MB&color=blue) | ![Vulns](https://img.shields.io/static/v1?label=Security&message=77+Vulns+%283+Crit%29&color=red) | 97.3% | `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown` |
| `python:3.10-slim` | ![Size](https://img.shields.io/static/v1?label=Size&message=144.62+MB&color=blue) | ![Vulns](https://img.shields.io/static/v1?label=Security&message=41+Vulns+%280+Crit%29&color=orange) | 97.3% | `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown` |
| `python:3.11-slim` | ![Size](https://img.shields.io/static/v1?label=Size&message=147.15+MB&color=blue) | ![Vulns](https://img.shields.io/static/v1?label=Security&message=39+Vulns+%280+Crit%29&color=orange) | 97.4% | `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown` |

<details>
<summary><strong>🔍 Full Report: python:3.9-slim</strong></summary>

## 🛡️ Security & Efficiency

**Base Image:** `linux (arm64)`
**Supported Architectures:** `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown`
**Efficiency Score:** 97.3%

### Vulnerabilities
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| 🔴 3 | 🟠 20 | 51 | 3 |

<details>
<summary><strong>👇 Expand Vulnerability Details (117 found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [CVE-2025-15467](https://nvd.nist.gov/vuln/detail/CVE-2025-15467) | Critical | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2025-15467](https://nvd.nist.gov/vuln/detail/CVE-2025-15467) | Critical | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2025-15467](https://nvd.nist.gov/vuln/detail/CVE-2025-15467) | Critical | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [CVE-2023-36632](https://nvd.nist.gov/vuln/detail/CVE-2023-36632) | High | `python` | `3.9.25` |
| [CVE-2025-13836](https://nvd.nist.gov/vuln/detail/CVE-2025-13836) | High | `python` | `3.9.25` |
| [CVE-2025-15281](https://nvd.nist.gov/vuln/detail/CVE-2025-15281) | High | `libc-bin` | `2.41-12` |
| [CVE-2025-15281](https://nvd.nist.gov/vuln/detail/CVE-2025-15281) | High | `libc6` | `2.41-12` |
| [CVE-2025-69419](https://nvd.nist.gov/vuln/detail/CVE-2025-69419) | High | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [CVE-2025-69419](https://nvd.nist.gov/vuln/detail/CVE-2025-69419) | High | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2025-69419](https://nvd.nist.gov/vuln/detail/CVE-2025-69419) | High | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2025-69420](https://nvd.nist.gov/vuln/detail/CVE-2025-69420) | High | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2025-69420](https://nvd.nist.gov/vuln/detail/CVE-2025-69420) | High | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2025-69420](https://nvd.nist.gov/vuln/detail/CVE-2025-69420) | High | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [CVE-2025-69421](https://nvd.nist.gov/vuln/detail/CVE-2025-69421) | High | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2025-69421](https://nvd.nist.gov/vuln/detail/CVE-2025-69421) | High | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [CVE-2025-69421](https://nvd.nist.gov/vuln/detail/CVE-2025-69421) | High | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2026-0861](https://nvd.nist.gov/vuln/detail/CVE-2026-0861) | High | `libc-bin` | `2.41-12` |
| [CVE-2026-0861](https://nvd.nist.gov/vuln/detail/CVE-2026-0861) | High | `libc6` | `2.41-12` |
| [CVE-2026-0915](https://nvd.nist.gov/vuln/detail/CVE-2026-0915) | High | `libc6` | `2.41-12` |
| [CVE-2026-0915](https://nvd.nist.gov/vuln/detail/CVE-2026-0915) | High | `libc-bin` | `2.41-12` |
| [GHSA-58pv-8j8x-9vj2](https://nvd.nist.gov/vuln/detail/GHSA-58pv-8j8x-9vj2) | High | `jaraco-context` | `5.3.0` |
| [GHSA-8rrh-rw8j-w5fx](https://nvd.nist.gov/vuln/detail/GHSA-8rrh-rw8j-w5fx) | High | `wheel` | `0.45.1` |
| [GHSA-8rrh-rw8j-w5fx](https://nvd.nist.gov/vuln/detail/GHSA-8rrh-rw8j-w5fx) | High | `wheel` | `0.45.1` |
| [CVE-2025-11187](https://nvd.nist.gov/vuln/detail/CVE-2025-11187) | Medium | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2025-11187](https://nvd.nist.gov/vuln/detail/CVE-2025-11187) | Medium | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [CVE-2025-11187](https://nvd.nist.gov/vuln/detail/CVE-2025-11187) | Medium | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2025-11468](https://nvd.nist.gov/vuln/detail/CVE-2025-11468) | Medium | `python` | `3.9.25` |
| [CVE-2025-12084](https://nvd.nist.gov/vuln/detail/CVE-2025-12084) | Medium | `python` | `3.9.25` |
| [CVE-2025-12781](https://nvd.nist.gov/vuln/detail/CVE-2025-12781) | Medium | `python` | `3.9.25` |
| [CVE-2025-13837](https://nvd.nist.gov/vuln/detail/CVE-2025-13837) | Medium | `python` | `3.9.25` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `util-linux` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `login` | `1:4.16.0-2+really2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `liblastlog2-2` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libblkid1` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `bsdutils` | `1:2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libsmartcols1` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libuuid1` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libmount1` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `mount` | `2.41-5` |
| [CVE-2025-15282](https://nvd.nist.gov/vuln/detail/CVE-2025-15282) | Medium | `python` | `3.9.25` |
| [CVE-2025-15366](https://nvd.nist.gov/vuln/detail/CVE-2025-15366) | Medium | `python` | `3.9.25` |
| [CVE-2025-15367](https://nvd.nist.gov/vuln/detail/CVE-2025-15367) | Medium | `python` | `3.9.25` |
| [CVE-2025-15468](https://nvd.nist.gov/vuln/detail/CVE-2025-15468) | Medium | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2025-15468](https://nvd.nist.gov/vuln/detail/CVE-2025-15468) | Medium | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2025-15468](https://nvd.nist.gov/vuln/detail/CVE-2025-15468) | Medium | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [CVE-2025-15469](https://nvd.nist.gov/vuln/detail/CVE-2025-15469) | Medium | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2025-15469](https://nvd.nist.gov/vuln/detail/CVE-2025-15469) | Medium | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [CVE-2025-15469](https://nvd.nist.gov/vuln/detail/CVE-2025-15469) | Medium | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2025-6075](https://nvd.nist.gov/vuln/detail/CVE-2025-6075) | Medium | `python` | `3.9.25` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `ncurses-bin` | `6.5+20250216-2` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libncursesw6` | `6.5+20250216-2` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libtinfo6` | `6.5+20250216-2` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `ncurses-base` | `6.5+20250216-2` |
| [CVE-2025-66199](https://nvd.nist.gov/vuln/detail/CVE-2025-66199) | Medium | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2025-66199](https://nvd.nist.gov/vuln/detail/CVE-2025-66199) | Medium | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2025-66199](https://nvd.nist.gov/vuln/detail/CVE-2025-66199) | Medium | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [CVE-2025-68160](https://nvd.nist.gov/vuln/detail/CVE-2025-68160) | Medium | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2025-68160](https://nvd.nist.gov/vuln/detail/CVE-2025-68160) | Medium | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2025-68160](https://nvd.nist.gov/vuln/detail/CVE-2025-68160) | Medium | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [CVE-2025-69418](https://nvd.nist.gov/vuln/detail/CVE-2025-69418) | Medium | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [CVE-2025-69418](https://nvd.nist.gov/vuln/detail/CVE-2025-69418) | Medium | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2025-69418](https://nvd.nist.gov/vuln/detail/CVE-2025-69418) | Medium | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2025-7709](https://nvd.nist.gov/vuln/detail/CVE-2025-7709) | Medium | `libsqlite3-0` | `3.46.1-7` |
| [CVE-2026-0672](https://nvd.nist.gov/vuln/detail/CVE-2026-0672) | Medium | `python` | `3.9.25` |
| [CVE-2026-0865](https://nvd.nist.gov/vuln/detail/CVE-2026-0865) | Medium | `python` | `3.9.25` |
| [CVE-2026-1299](https://nvd.nist.gov/vuln/detail/CVE-2026-1299) | Medium | `python` | `3.9.25` |
| [CVE-2026-22795](https://nvd.nist.gov/vuln/detail/CVE-2026-22795) | Medium | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2026-22795](https://nvd.nist.gov/vuln/detail/CVE-2026-22795) | Medium | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2026-22795](https://nvd.nist.gov/vuln/detail/CVE-2026-22795) | Medium | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [CVE-2026-22796](https://nvd.nist.gov/vuln/detail/CVE-2026-22796) | Medium | `openssl` | `3.5.1-1+deb13u1` |
| [CVE-2026-22796](https://nvd.nist.gov/vuln/detail/CVE-2026-22796) | Medium | `libssl3t64` | `3.5.1-1+deb13u1` |
| [CVE-2026-22796](https://nvd.nist.gov/vuln/detail/CVE-2026-22796) | Medium | `openssl-provider-legacy` | `3.5.1-1+deb13u1` |
| [GHSA-4xh5-x5gv-qwph](https://nvd.nist.gov/vuln/detail/GHSA-4xh5-x5gv-qwph) | Medium | `pip` | `23.0.1` |
| [GHSA-mq26-g339-26xf](https://nvd.nist.gov/vuln/detail/GHSA-mq26-g339-26xf) | Medium | `pip` | `23.0.1` |
| [CVE-2024-56433](https://nvd.nist.gov/vuln/detail/CVE-2024-56433) | Low | `login.defs` | `1:4.17.4-2` |
| [CVE-2024-56433](https://nvd.nist.gov/vuln/detail/CVE-2024-56433) | Low | `passwd` | `1:4.17.4-2` |
| [GHSA-6vgw-5pg2-w6jp](https://nvd.nist.gov/vuln/detail/GHSA-6vgw-5pg2-w6jp) | Low | `pip` | `23.0.1` |
| [CVE-2005-2541](https://nvd.nist.gov/vuln/detail/CVE-2005-2541) | Negligible | `tar` | `1.35+dfsg-3.1` |
| [CVE-2007-5686](https://nvd.nist.gov/vuln/detail/CVE-2007-5686) | Negligible | `login.defs` | `1:4.17.4-2` |
| [CVE-2007-5686](https://nvd.nist.gov/vuln/detail/CVE-2007-5686) | Negligible | `passwd` | `1:4.17.4-2` |
| [CVE-2010-4756](https://nvd.nist.gov/vuln/detail/CVE-2010-4756) | Negligible | `libc6` | `2.41-12` |
| [CVE-2010-4756](https://nvd.nist.gov/vuln/detail/CVE-2010-4756) | Negligible | `libc-bin` | `2.41-12` |
| [CVE-2011-3374](https://nvd.nist.gov/vuln/detail/CVE-2011-3374) | Negligible | `libapt-pkg7.0` | `3.0.3` |
| [CVE-2011-3374](https://nvd.nist.gov/vuln/detail/CVE-2011-3374) | Negligible | `apt` | `3.0.3` |
| [CVE-2011-4116](https://nvd.nist.gov/vuln/detail/CVE-2011-4116) | Negligible | `perl-base` | `5.40.1-6` |
| [CVE-2013-4392](https://nvd.nist.gov/vuln/detail/CVE-2013-4392) | Negligible | `libsystemd0` | `257.8-1~deb13u2` |
| [CVE-2013-4392](https://nvd.nist.gov/vuln/detail/CVE-2013-4392) | Negligible | `libudev1` | `257.8-1~deb13u2` |
| [CVE-2017-18018](https://nvd.nist.gov/vuln/detail/CVE-2017-18018) | Negligible | `coreutils` | `9.7-3` |
| [CVE-2018-20796](https://nvd.nist.gov/vuln/detail/CVE-2018-20796) | Negligible | `libc6` | `2.41-12` |
| [CVE-2018-20796](https://nvd.nist.gov/vuln/detail/CVE-2018-20796) | Negligible | `libc-bin` | `2.41-12` |
| [CVE-2019-1010022](https://nvd.nist.gov/vuln/detail/CVE-2019-1010022) | Negligible | `libc6` | `2.41-12` |
| [CVE-2019-1010022](https://nvd.nist.gov/vuln/detail/CVE-2019-1010022) | Negligible | `libc-bin` | `2.41-12` |
| [CVE-2019-1010023](https://nvd.nist.gov/vuln/detail/CVE-2019-1010023) | Negligible | `libc6` | `2.41-12` |
| [CVE-2019-1010023](https://nvd.nist.gov/vuln/detail/CVE-2019-1010023) | Negligible | `libc-bin` | `2.41-12` |
| [CVE-2019-1010024](https://nvd.nist.gov/vuln/detail/CVE-2019-1010024) | Negligible | `libc6` | `2.41-12` |
| [CVE-2019-1010024](https://nvd.nist.gov/vuln/detail/CVE-2019-1010024) | Negligible | `libc-bin` | `2.41-12` |
| [CVE-2019-1010025](https://nvd.nist.gov/vuln/detail/CVE-2019-1010025) | Negligible | `libc-bin` | `2.41-12` |
| [CVE-2019-1010025](https://nvd.nist.gov/vuln/detail/CVE-2019-1010025) | Negligible | `libc6` | `2.41-12` |
| [CVE-2019-9192](https://nvd.nist.gov/vuln/detail/CVE-2019-9192) | Negligible | `libc-bin` | `2.41-12` |
| [CVE-2019-9192](https://nvd.nist.gov/vuln/detail/CVE-2019-9192) | Negligible | `libc6` | `2.41-12` |
| [CVE-2021-45346](https://nvd.nist.gov/vuln/detail/CVE-2021-45346) | Negligible | `libsqlite3-0` | `3.46.1-7` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libblkid1` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libmount1` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `util-linux` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `mount` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `login` | `1:4.16.0-2+really2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libuuid1` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libsmartcols1` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `bsdutils` | `1:2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `liblastlog2-2` | `2.41-5` |
| [CVE-2023-31437](https://nvd.nist.gov/vuln/detail/CVE-2023-31437) | Negligible | `libudev1` | `257.8-1~deb13u2` |
| [CVE-2023-31437](https://nvd.nist.gov/vuln/detail/CVE-2023-31437) | Negligible | `libsystemd0` | `257.8-1~deb13u2` |
| [CVE-2023-31438](https://nvd.nist.gov/vuln/detail/CVE-2023-31438) | Negligible | `libsystemd0` | `257.8-1~deb13u2` |
| [CVE-2023-31438](https://nvd.nist.gov/vuln/detail/CVE-2023-31438) | Negligible | `libudev1` | `257.8-1~deb13u2` |
| [CVE-2023-31439](https://nvd.nist.gov/vuln/detail/CVE-2023-31439) | Negligible | `libudev1` | `257.8-1~deb13u2` |
| [CVE-2023-31439](https://nvd.nist.gov/vuln/detail/CVE-2023-31439) | Negligible | `libsystemd0` | `257.8-1~deb13u2` |
| [CVE-2025-5278](https://nvd.nist.gov/vuln/detail/CVE-2025-5278) | Negligible | `coreutils` | `9.7-3` |
</details>

<details>
<summary><strong>📦 Installed Packages (115 total)</strong></summary>

| Package | Version |
|---------|---------|
| Simple Launcher | 1.1.0.14 |
| adduser | 3.152 |
| apt | 3.0.3 |
| autocommand | 2.2.2 |
| backports-tarfile | 1.2.0 |
| base-files | 13.8+deb13u1 |
| base-passwd | 3.6.7 |
| bash | 5.2.37-2+b5 |
| bsdutils | 1:2.41-5 |
| ca-certificates | 20250419 |
| cli | UNKNOWN |
| cli-32 | UNKNOWN |
| cli-64 | UNKNOWN |
| cli-arm64 | UNKNOWN |
| coreutils | 9.7-3 |
| dash | 0.5.12-12 |
| debconf | 1.5.91 |
| debian-archive-keyring | 2025.1 |
| debianutils | 5.23.2 |
| diffutils | 1:3.10-4 |
| dpkg | 1.22.21 |
| findutils | 4.10.0-3 |
| gcc-14-base | 14.2.0-19 |
| grep | 3.11-4+b1 |
| gui | UNKNOWN |
| gui-32 | UNKNOWN |
| gui-64 | UNKNOWN |
| gui-arm64 | UNKNOWN |
| gzip | 1.13-1 |
| hostname | 3.25 |
| importlib-metadata | 8.0.0 |
| inflect | 7.3.1 |
| init-system-helpers | 1.69~deb13u1 |
| jaraco-collections | 5.1.0 |
| jaraco-context | 5.3.0 |
| jaraco-functools | 4.0.1 |
| jaraco-text | 3.12.1 |
| libacl1 | 2.3.2-2+b1 |
| libapt-pkg7.0 | 3.0.3 |
| libattr1 | 1:2.5.2-3 |
| libaudit-common | 1:4.0.2-2 |
| libaudit1 | 1:4.0.2-2+b2 |
| libblkid1 | 2.41-5 |
| libbsd0 | 0.12.2-2 |
| libbz2-1.0 | 1.0.8-6 |
| libc-bin | 2.41-12 |
| libc6 | 2.41-12 |
| libcap-ng0 | 0.8.5-4+b1 |
| libcap2 | 1:2.75-10+b1 |
| libcrypt1 | 1:4.4.38-1 |
| libdb5.3t64 | 5.3.28+dfsg2-9 |
| libdebconfclient0 | 0.280 |
| libffi8 | 3.4.8-2 |
| libgcc-s1 | 14.2.0-19 |
| libgdbm6t64 | 1.24-2 |
| libgmp10 | 2:6.3.0+dfsg-3 |
| libhogweed6t64 | 3.10.1-1 |
| liblastlog2-2 | 2.41-5 |
| liblz4-1 | 1.10.0-4 |
| liblzma5 | 5.8.1-1 |
| libmd0 | 1.1.0-2+b1 |
| libmount1 | 2.41-5 |
| libncursesw6 | 6.5+20250216-2 |
| libnettle8t64 | 3.10.1-1 |
| libpam-modules | 1.7.0-5 |
| libpam-modules-bin | 1.7.0-5 |
| libpam-runtime | 1.7.0-5 |
| libpam0g | 1.7.0-5 |
| libpcre2-8-0 | 10.46-1~deb13u1 |
| libreadline8t64 | 8.2-6 |
| libseccomp2 | 2.6.0-2 |
| libselinux1 | 3.8.1-1 |
| libsemanage-common | 3.8.1-1 |
| libsemanage2 | 3.8.1-1 |
| libsepol2 | 3.8.1-1 |
| libsmartcols1 | 2.41-5 |
| libsqlite3-0 | 3.46.1-7 |
| libssl3t64 | 3.5.1-1+deb13u1 |
| libstdc++6 | 14.2.0-19 |
| libsystemd0 | 257.8-1~deb13u2 |
| libtinfo6 | 6.5+20250216-2 |
| libudev1 | 257.8-1~deb13u2 |
| libuuid1 | 2.41-5 |
| libxxhash0 | 0.8.3-2 |
| libzstd1 | 1.5.7+dfsg-1 |
| login | 1:4.16.0-2+really2.41-5 |
| login.defs | 1:4.17.4-2 |
| mawk | 1.3.4.20250131-1 |
| more-itertools | 10.3.0 |
| mount | 2.41-5 |
| ncurses-base | 6.5+20250216-2 |
| ncurses-bin | 6.5+20250216-2 |
| netbase | 6.5 |
| openssl | 3.5.1-1+deb13u1 |
| openssl-provider-legacy | 3.5.1-1+deb13u1 |
| packaging | 24.2 |
| passwd | 1:4.17.4-2 |
| perl-base | 5.40.1-6 |
| pip | 23.0.1 |
| platformdirs | 4.2.2 |
| python | 3.9.25 |
| readline-common | 8.2-6 |
| sed | 4.9-2+b1 |
| setuptools | 79.0.1 |
| sqv | 1.3.0-3 |
| sysvinit-utils | 3.14-4 |
| tar | 1.35+dfsg-3.1 |
| tomli | 2.0.1 |
| typeguard | 4.3.0 |
| typing-extensions | 4.12.2 |
| tzdata | 2025b-4+deb13u1 |
| util-linux | 2.41-5 |
| wheel | 0.45.1 |
| zipp | 3.19.2 |
| zlib1g | 1:1.3.dfsg+really1.3.1-1+b1 |
</details>

</details>

<details>
<summary><strong>🔍 Full Report: python:3.10-slim</strong></summary>

## 🛡️ Security & Efficiency

**Base Image:** `linux (arm64)`
**Supported Architectures:** `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown`
**Efficiency Score:** 97.3%

### Vulnerabilities
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| 🟢 0 | 🟠 11 | 27 | 3 |

<details>
<summary><strong>👇 Expand Vulnerability Details (81 found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [CVE-2023-36632](https://nvd.nist.gov/vuln/detail/CVE-2023-36632) | High | `python` | `3.10.19` |
| [CVE-2025-13836](https://nvd.nist.gov/vuln/detail/CVE-2025-13836) | High | `python` | `3.10.19` |
| [CVE-2025-15281](https://nvd.nist.gov/vuln/detail/CVE-2025-15281) | High | `libc-bin` | `2.41-12+deb13u1` |
| [CVE-2025-15281](https://nvd.nist.gov/vuln/detail/CVE-2025-15281) | High | `libc6` | `2.41-12+deb13u1` |
| [CVE-2026-0861](https://nvd.nist.gov/vuln/detail/CVE-2026-0861) | High | `libc6` | `2.41-12+deb13u1` |
| [CVE-2026-0861](https://nvd.nist.gov/vuln/detail/CVE-2026-0861) | High | `libc-bin` | `2.41-12+deb13u1` |
| [CVE-2026-0915](https://nvd.nist.gov/vuln/detail/CVE-2026-0915) | High | `libc6` | `2.41-12+deb13u1` |
| [CVE-2026-0915](https://nvd.nist.gov/vuln/detail/CVE-2026-0915) | High | `libc-bin` | `2.41-12+deb13u1` |
| [GHSA-58pv-8j8x-9vj2](https://nvd.nist.gov/vuln/detail/GHSA-58pv-8j8x-9vj2) | High | `jaraco-context` | `5.3.0` |
| [GHSA-8rrh-rw8j-w5fx](https://nvd.nist.gov/vuln/detail/GHSA-8rrh-rw8j-w5fx) | High | `wheel` | `0.45.1` |
| [GHSA-8rrh-rw8j-w5fx](https://nvd.nist.gov/vuln/detail/GHSA-8rrh-rw8j-w5fx) | High | `wheel` | `0.45.1` |
| [CVE-2025-11468](https://nvd.nist.gov/vuln/detail/CVE-2025-11468) | Medium | `python` | `3.10.19` |
| [CVE-2025-12084](https://nvd.nist.gov/vuln/detail/CVE-2025-12084) | Medium | `python` | `3.10.19` |
| [CVE-2025-12781](https://nvd.nist.gov/vuln/detail/CVE-2025-12781) | Medium | `python` | `3.10.19` |
| [CVE-2025-13837](https://nvd.nist.gov/vuln/detail/CVE-2025-13837) | Medium | `python` | `3.10.19` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libsmartcols1` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `login` | `1:4.16.0-2+really2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libuuid1` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libblkid1` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libmount1` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `mount` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `bsdutils` | `1:2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `liblastlog2-2` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `util-linux` | `2.41-5` |
| [CVE-2025-15282](https://nvd.nist.gov/vuln/detail/CVE-2025-15282) | Medium | `python` | `3.10.19` |
| [CVE-2025-15366](https://nvd.nist.gov/vuln/detail/CVE-2025-15366) | Medium | `python` | `3.10.19` |
| [CVE-2025-15367](https://nvd.nist.gov/vuln/detail/CVE-2025-15367) | Medium | `python` | `3.10.19` |
| [CVE-2025-6075](https://nvd.nist.gov/vuln/detail/CVE-2025-6075) | Medium | `python` | `3.10.19` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libtinfo6` | `6.5+20250216-2` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `ncurses-base` | `6.5+20250216-2` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `ncurses-bin` | `6.5+20250216-2` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libncursesw6` | `6.5+20250216-2` |
| [CVE-2025-7709](https://nvd.nist.gov/vuln/detail/CVE-2025-7709) | Medium | `libsqlite3-0` | `3.46.1-7` |
| [CVE-2026-0672](https://nvd.nist.gov/vuln/detail/CVE-2026-0672) | Medium | `python` | `3.10.19` |
| [CVE-2026-0865](https://nvd.nist.gov/vuln/detail/CVE-2026-0865) | Medium | `python` | `3.10.19` |
| [CVE-2026-1299](https://nvd.nist.gov/vuln/detail/CVE-2026-1299) | Medium | `python` | `3.10.19` |
| [GHSA-4xh5-x5gv-qwph](https://nvd.nist.gov/vuln/detail/GHSA-4xh5-x5gv-qwph) | Medium | `pip` | `23.0.1` |
| [GHSA-mq26-g339-26xf](https://nvd.nist.gov/vuln/detail/GHSA-mq26-g339-26xf) | Medium | `pip` | `23.0.1` |
| [CVE-2024-56433](https://nvd.nist.gov/vuln/detail/CVE-2024-56433) | Low | `passwd` | `1:4.17.4-2` |
| [CVE-2024-56433](https://nvd.nist.gov/vuln/detail/CVE-2024-56433) | Low | `login.defs` | `1:4.17.4-2` |
| [GHSA-6vgw-5pg2-w6jp](https://nvd.nist.gov/vuln/detail/GHSA-6vgw-5pg2-w6jp) | Low | `pip` | `23.0.1` |
| [CVE-2005-2541](https://nvd.nist.gov/vuln/detail/CVE-2005-2541) | Negligible | `tar` | `1.35+dfsg-3.1` |
| [CVE-2007-5686](https://nvd.nist.gov/vuln/detail/CVE-2007-5686) | Negligible | `passwd` | `1:4.17.4-2` |
| [CVE-2007-5686](https://nvd.nist.gov/vuln/detail/CVE-2007-5686) | Negligible | `login.defs` | `1:4.17.4-2` |
| [CVE-2010-4756](https://nvd.nist.gov/vuln/detail/CVE-2010-4756) | Negligible | `libc6` | `2.41-12+deb13u1` |
| [CVE-2010-4756](https://nvd.nist.gov/vuln/detail/CVE-2010-4756) | Negligible | `libc-bin` | `2.41-12+deb13u1` |
| [CVE-2011-3374](https://nvd.nist.gov/vuln/detail/CVE-2011-3374) | Negligible | `apt` | `3.0.3` |
| [CVE-2011-3374](https://nvd.nist.gov/vuln/detail/CVE-2011-3374) | Negligible | `libapt-pkg7.0` | `3.0.3` |
| [CVE-2011-4116](https://nvd.nist.gov/vuln/detail/CVE-2011-4116) | Negligible | `perl-base` | `5.40.1-6` |
| [CVE-2013-4392](https://nvd.nist.gov/vuln/detail/CVE-2013-4392) | Negligible | `libudev1` | `257.9-1~deb13u1` |
| [CVE-2013-4392](https://nvd.nist.gov/vuln/detail/CVE-2013-4392) | Negligible | `libsystemd0` | `257.9-1~deb13u1` |
| [CVE-2017-18018](https://nvd.nist.gov/vuln/detail/CVE-2017-18018) | Negligible | `coreutils` | `9.7-3` |
| [CVE-2018-20796](https://nvd.nist.gov/vuln/detail/CVE-2018-20796) | Negligible | `libc6` | `2.41-12+deb13u1` |
| [CVE-2018-20796](https://nvd.nist.gov/vuln/detail/CVE-2018-20796) | Negligible | `libc-bin` | `2.41-12+deb13u1` |
| [CVE-2019-1010022](https://nvd.nist.gov/vuln/detail/CVE-2019-1010022) | Negligible | `libc-bin` | `2.41-12+deb13u1` |
| [CVE-2019-1010022](https://nvd.nist.gov/vuln/detail/CVE-2019-1010022) | Negligible | `libc6` | `2.41-12+deb13u1` |
| [CVE-2019-1010023](https://nvd.nist.gov/vuln/detail/CVE-2019-1010023) | Negligible | `libc-bin` | `2.41-12+deb13u1` |
| [CVE-2019-1010023](https://nvd.nist.gov/vuln/detail/CVE-2019-1010023) | Negligible | `libc6` | `2.41-12+deb13u1` |
| [CVE-2019-1010024](https://nvd.nist.gov/vuln/detail/CVE-2019-1010024) | Negligible | `libc-bin` | `2.41-12+deb13u1` |
| [CVE-2019-1010024](https://nvd.nist.gov/vuln/detail/CVE-2019-1010024) | Negligible | `libc6` | `2.41-12+deb13u1` |
| [CVE-2019-1010025](https://nvd.nist.gov/vuln/detail/CVE-2019-1010025) | Negligible | `libc-bin` | `2.41-12+deb13u1` |
| [CVE-2019-1010025](https://nvd.nist.gov/vuln/detail/CVE-2019-1010025) | Negligible | `libc6` | `2.41-12+deb13u1` |
| [CVE-2019-9192](https://nvd.nist.gov/vuln/detail/CVE-2019-9192) | Negligible | `libc6` | `2.41-12+deb13u1` |
| [CVE-2019-9192](https://nvd.nist.gov/vuln/detail/CVE-2019-9192) | Negligible | `libc-bin` | `2.41-12+deb13u1` |
| [CVE-2021-45346](https://nvd.nist.gov/vuln/detail/CVE-2021-45346) | Negligible | `libsqlite3-0` | `3.46.1-7` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `mount` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `login` | `1:4.16.0-2+really2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libsmartcols1` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `util-linux` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libuuid1` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `bsdutils` | `1:2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libblkid1` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `liblastlog2-2` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libmount1` | `2.41-5` |
| [CVE-2023-31437](https://nvd.nist.gov/vuln/detail/CVE-2023-31437) | Negligible | `libudev1` | `257.9-1~deb13u1` |
| [CVE-2023-31437](https://nvd.nist.gov/vuln/detail/CVE-2023-31437) | Negligible | `libsystemd0` | `257.9-1~deb13u1` |
| [CVE-2023-31438](https://nvd.nist.gov/vuln/detail/CVE-2023-31438) | Negligible | `libsystemd0` | `257.9-1~deb13u1` |
| [CVE-2023-31438](https://nvd.nist.gov/vuln/detail/CVE-2023-31438) | Negligible | `libudev1` | `257.9-1~deb13u1` |
| [CVE-2023-31439](https://nvd.nist.gov/vuln/detail/CVE-2023-31439) | Negligible | `libudev1` | `257.9-1~deb13u1` |
| [CVE-2023-31439](https://nvd.nist.gov/vuln/detail/CVE-2023-31439) | Negligible | `libsystemd0` | `257.9-1~deb13u1` |
| [CVE-2025-5278](https://nvd.nist.gov/vuln/detail/CVE-2025-5278) | Negligible | `coreutils` | `9.7-3` |
</details>

<details>
<summary><strong>📦 Installed Packages (115 total)</strong></summary>

| Package | Version |
|---------|---------|
| Simple Launcher | 1.1.0.14 |
| adduser | 3.152 |
| apt | 3.0.3 |
| autocommand | 2.2.2 |
| backports-tarfile | 1.2.0 |
| base-files | 13.8+deb13u3 |
| base-passwd | 3.6.7 |
| bash | 5.2.37-2+b7 |
| bsdutils | 1:2.41-5 |
| ca-certificates | 20250419 |
| cli | UNKNOWN |
| cli-32 | UNKNOWN |
| cli-64 | UNKNOWN |
| cli-arm64 | UNKNOWN |
| coreutils | 9.7-3 |
| dash | 0.5.12-12 |
| debconf | 1.5.91 |
| debian-archive-keyring | 2025.1 |
| debianutils | 5.23.2 |
| diffutils | 1:3.10-4 |
| dpkg | 1.22.21 |
| findutils | 4.10.0-3 |
| gcc-14-base | 14.2.0-19 |
| grep | 3.11-4+b1 |
| gui | UNKNOWN |
| gui-32 | UNKNOWN |
| gui-64 | UNKNOWN |
| gui-arm64 | UNKNOWN |
| gzip | 1.13-1 |
| hostname | 3.25 |
| importlib-metadata | 8.0.0 |
| inflect | 7.3.1 |
| init-system-helpers | 1.69~deb13u1 |
| jaraco-collections | 5.1.0 |
| jaraco-context | 5.3.0 |
| jaraco-functools | 4.0.1 |
| jaraco-text | 3.12.1 |
| libacl1 | 2.3.2-2+b1 |
| libapt-pkg7.0 | 3.0.3 |
| libattr1 | 1:2.5.2-3 |
| libaudit-common | 1:4.0.2-2 |
| libaudit1 | 1:4.0.2-2+b2 |
| libblkid1 | 2.41-5 |
| libbsd0 | 0.12.2-2 |
| libbz2-1.0 | 1.0.8-6 |
| libc-bin | 2.41-12+deb13u1 |
| libc6 | 2.41-12+deb13u1 |
| libcap-ng0 | 0.8.5-4+b1 |
| libcap2 | 1:2.75-10+b3 |
| libcrypt1 | 1:4.4.38-1 |
| libdb5.3t64 | 5.3.28+dfsg2-9 |
| libdebconfclient0 | 0.280 |
| libffi8 | 3.4.8-2 |
| libgcc-s1 | 14.2.0-19 |
| libgdbm6t64 | 1.24-2 |
| libgmp10 | 2:6.3.0+dfsg-3 |
| libhogweed6t64 | 3.10.1-1 |
| liblastlog2-2 | 2.41-5 |
| liblz4-1 | 1.10.0-4 |
| liblzma5 | 5.8.1-1 |
| libmd0 | 1.1.0-2+b1 |
| libmount1 | 2.41-5 |
| libncursesw6 | 6.5+20250216-2 |
| libnettle8t64 | 3.10.1-1 |
| libpam-modules | 1.7.0-5 |
| libpam-modules-bin | 1.7.0-5 |
| libpam-runtime | 1.7.0-5 |
| libpam0g | 1.7.0-5 |
| libpcre2-8-0 | 10.46-1~deb13u1 |
| libreadline8t64 | 8.2-6 |
| libseccomp2 | 2.6.0-2 |
| libselinux1 | 3.8.1-1 |
| libsemanage-common | 3.8.1-1 |
| libsemanage2 | 3.8.1-1 |
| libsepol2 | 3.8.1-1 |
| libsmartcols1 | 2.41-5 |
| libsqlite3-0 | 3.46.1-7 |
| libssl3t64 | 3.5.4-1~deb13u2 |
| libstdc++6 | 14.2.0-19 |
| libsystemd0 | 257.9-1~deb13u1 |
| libtinfo6 | 6.5+20250216-2 |
| libudev1 | 257.9-1~deb13u1 |
| libuuid1 | 2.41-5 |
| libxxhash0 | 0.8.3-2 |
| libzstd1 | 1.5.7+dfsg-1 |
| login | 1:4.16.0-2+really2.41-5 |
| login.defs | 1:4.17.4-2 |
| mawk | 1.3.4.20250131-1 |
| more-itertools | 10.3.0 |
| mount | 2.41-5 |
| ncurses-base | 6.5+20250216-2 |
| ncurses-bin | 6.5+20250216-2 |
| netbase | 6.5 |
| openssl | 3.5.4-1~deb13u2 |
| openssl-provider-legacy | 3.5.4-1~deb13u2 |
| packaging | 24.2 |
| passwd | 1:4.17.4-2 |
| perl-base | 5.40.1-6 |
| pip | 23.0.1 |
| platformdirs | 4.2.2 |
| python | 3.10.19 |
| readline-common | 8.2-6 |
| sed | 4.9-2+b1 |
| setuptools | 79.0.1 |
| sqv | 1.3.0-3+b2 |
| sysvinit-utils | 3.14-4 |
| tar | 1.35+dfsg-3.1 |
| tomli | 2.0.1 |
| typeguard | 4.3.0 |
| typing-extensions | 4.12.2 |
| tzdata | 2025b-4+deb13u1 |
| util-linux | 2.41-5 |
| wheel | 0.45.1 |
| zipp | 3.19.2 |
| zlib1g | 1:1.3.dfsg+really1.3.1-1+b1 |
</details>

</details>

<details>
<summary><strong>🔍 Full Report: python:3.11-slim</strong></summary>

## 🛡️ Security & Efficiency

**Base Image:** `linux (arm64)`
**Supported Architectures:** `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown`
**Efficiency Score:** 97.4%

### Vulnerabilities
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| 🟢 0 | 🟠 10 | 26 | 3 |

<details>
<summary><strong>👇 Expand Vulnerability Details (79 found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [CVE-2025-13836](https://nvd.nist.gov/vuln/detail/CVE-2025-13836) | High | `python` | `3.11.14` |
| [CVE-2025-15281](https://nvd.nist.gov/vuln/detail/CVE-2025-15281) | High | `libc6` | `2.41-12+deb13u1` |
| [CVE-2025-15281](https://nvd.nist.gov/vuln/detail/CVE-2025-15281) | High | `libc-bin` | `2.41-12+deb13u1` |
| [CVE-2026-0861](https://nvd.nist.gov/vuln/detail/CVE-2026-0861) | High | `libc6` | `2.41-12+deb13u1` |
| [CVE-2026-0861](https://nvd.nist.gov/vuln/detail/CVE-2026-0861) | High | `libc-bin` | `2.41-12+deb13u1` |
| [CVE-2026-0915](https://nvd.nist.gov/vuln/detail/CVE-2026-0915) | High | `libc6` | `2.41-12+deb13u1` |
| [CVE-2026-0915](https://nvd.nist.gov/vuln/detail/CVE-2026-0915) | High | `libc-bin` | `2.41-12+deb13u1` |
| [GHSA-58pv-8j8x-9vj2](https://nvd.nist.gov/vuln/detail/GHSA-58pv-8j8x-9vj2) | High | `jaraco-context` | `5.3.0` |
| [GHSA-8rrh-rw8j-w5fx](https://nvd.nist.gov/vuln/detail/GHSA-8rrh-rw8j-w5fx) | High | `wheel` | `0.45.1` |
| [GHSA-8rrh-rw8j-w5fx](https://nvd.nist.gov/vuln/detail/GHSA-8rrh-rw8j-w5fx) | High | `wheel` | `0.45.1` |
| [CVE-2025-11468](https://nvd.nist.gov/vuln/detail/CVE-2025-11468) | Medium | `python` | `3.11.14` |
| [CVE-2025-12084](https://nvd.nist.gov/vuln/detail/CVE-2025-12084) | Medium | `python` | `3.11.14` |
| [CVE-2025-12781](https://nvd.nist.gov/vuln/detail/CVE-2025-12781) | Medium | `python` | `3.11.14` |
| [CVE-2025-13837](https://nvd.nist.gov/vuln/detail/CVE-2025-13837) | Medium | `python` | `3.11.14` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `bsdutils` | `1:2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libblkid1` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `util-linux` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `mount` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `login` | `1:4.16.0-2+really2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libuuid1` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libsmartcols1` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libmount1` | `2.41-5` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `liblastlog2-2` | `2.41-5` |
| [CVE-2025-15282](https://nvd.nist.gov/vuln/detail/CVE-2025-15282) | Medium | `python` | `3.11.14` |
| [CVE-2025-15366](https://nvd.nist.gov/vuln/detail/CVE-2025-15366) | Medium | `python` | `3.11.14` |
| [CVE-2025-15367](https://nvd.nist.gov/vuln/detail/CVE-2025-15367) | Medium | `python` | `3.11.14` |
| [CVE-2025-6075](https://nvd.nist.gov/vuln/detail/CVE-2025-6075) | Medium | `python` | `3.11.14` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libtinfo6` | `6.5+20250216-2` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libncursesw6` | `6.5+20250216-2` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `ncurses-base` | `6.5+20250216-2` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `ncurses-bin` | `6.5+20250216-2` |
| [CVE-2025-7709](https://nvd.nist.gov/vuln/detail/CVE-2025-7709) | Medium | `libsqlite3-0` | `3.46.1-7` |
| [CVE-2026-0672](https://nvd.nist.gov/vuln/detail/CVE-2026-0672) | Medium | `python` | `3.11.14` |
| [CVE-2026-0865](https://nvd.nist.gov/vuln/detail/CVE-2026-0865) | Medium | `python` | `3.11.14` |
| [CVE-2026-1299](https://nvd.nist.gov/vuln/detail/CVE-2026-1299) | Medium | `python` | `3.11.14` |
| [GHSA-4xh5-x5gv-qwph](https://nvd.nist.gov/vuln/detail/GHSA-4xh5-x5gv-qwph) | Medium | `pip` | `24.0` |
| [CVE-2024-56433](https://nvd.nist.gov/vuln/detail/CVE-2024-56433) | Low | `login.defs` | `1:4.17.4-2` |
| [CVE-2024-56433](https://nvd.nist.gov/vuln/detail/CVE-2024-56433) | Low | `passwd` | `1:4.17.4-2` |
| [GHSA-6vgw-5pg2-w6jp](https://nvd.nist.gov/vuln/detail/GHSA-6vgw-5pg2-w6jp) | Low | `pip` | `24.0` |
| [CVE-2005-2541](https://nvd.nist.gov/vuln/detail/CVE-2005-2541) | Negligible | `tar` | `1.35+dfsg-3.1` |
| [CVE-2007-5686](https://nvd.nist.gov/vuln/detail/CVE-2007-5686) | Negligible | `login.defs` | `1:4.17.4-2` |
| [CVE-2007-5686](https://nvd.nist.gov/vuln/detail/CVE-2007-5686) | Negligible | `passwd` | `1:4.17.4-2` |
| [CVE-2010-4756](https://nvd.nist.gov/vuln/detail/CVE-2010-4756) | Negligible | `libc6` | `2.41-12+deb13u1` |
| [CVE-2010-4756](https://nvd.nist.gov/vuln/detail/CVE-2010-4756) | Negligible | `libc-bin` | `2.41-12+deb13u1` |
| [CVE-2011-3374](https://nvd.nist.gov/vuln/detail/CVE-2011-3374) | Negligible | `apt` | `3.0.3` |
| [CVE-2011-3374](https://nvd.nist.gov/vuln/detail/CVE-2011-3374) | Negligible | `libapt-pkg7.0` | `3.0.3` |
| [CVE-2011-4116](https://nvd.nist.gov/vuln/detail/CVE-2011-4116) | Negligible | `perl-base` | `5.40.1-6` |
| [CVE-2013-4392](https://nvd.nist.gov/vuln/detail/CVE-2013-4392) | Negligible | `libudev1` | `257.9-1~deb13u1` |
| [CVE-2013-4392](https://nvd.nist.gov/vuln/detail/CVE-2013-4392) | Negligible | `libsystemd0` | `257.9-1~deb13u1` |
| [CVE-2017-18018](https://nvd.nist.gov/vuln/detail/CVE-2017-18018) | Negligible | `coreutils` | `9.7-3` |
| [CVE-2018-20796](https://nvd.nist.gov/vuln/detail/CVE-2018-20796) | Negligible | `libc-bin` | `2.41-12+deb13u1` |
| [CVE-2018-20796](https://nvd.nist.gov/vuln/detail/CVE-2018-20796) | Negligible | `libc6` | `2.41-12+deb13u1` |
| [CVE-2019-1010022](https://nvd.nist.gov/vuln/detail/CVE-2019-1010022) | Negligible | `libc-bin` | `2.41-12+deb13u1` |
| [CVE-2019-1010022](https://nvd.nist.gov/vuln/detail/CVE-2019-1010022) | Negligible | `libc6` | `2.41-12+deb13u1` |
| [CVE-2019-1010023](https://nvd.nist.gov/vuln/detail/CVE-2019-1010023) | Negligible | `libc-bin` | `2.41-12+deb13u1` |
| [CVE-2019-1010023](https://nvd.nist.gov/vuln/detail/CVE-2019-1010023) | Negligible | `libc6` | `2.41-12+deb13u1` |
| [CVE-2019-1010024](https://nvd.nist.gov/vuln/detail/CVE-2019-1010024) | Negligible | `libc6` | `2.41-12+deb13u1` |
| [CVE-2019-1010024](https://nvd.nist.gov/vuln/detail/CVE-2019-1010024) | Negligible | `libc-bin` | `2.41-12+deb13u1` |
| [CVE-2019-1010025](https://nvd.nist.gov/vuln/detail/CVE-2019-1010025) | Negligible | `libc-bin` | `2.41-12+deb13u1` |
| [CVE-2019-1010025](https://nvd.nist.gov/vuln/detail/CVE-2019-1010025) | Negligible | `libc6` | `2.41-12+deb13u1` |
| [CVE-2019-9192](https://nvd.nist.gov/vuln/detail/CVE-2019-9192) | Negligible | `libc6` | `2.41-12+deb13u1` |
| [CVE-2019-9192](https://nvd.nist.gov/vuln/detail/CVE-2019-9192) | Negligible | `libc-bin` | `2.41-12+deb13u1` |
| [CVE-2021-45346](https://nvd.nist.gov/vuln/detail/CVE-2021-45346) | Negligible | `libsqlite3-0` | `3.46.1-7` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `mount` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libuuid1` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libsmartcols1` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `login` | `1:4.16.0-2+really2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `liblastlog2-2` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `util-linux` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libmount1` | `2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `bsdutils` | `1:2.41-5` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libblkid1` | `2.41-5` |
| [CVE-2023-31437](https://nvd.nist.gov/vuln/detail/CVE-2023-31437) | Negligible | `libsystemd0` | `257.9-1~deb13u1` |
| [CVE-2023-31437](https://nvd.nist.gov/vuln/detail/CVE-2023-31437) | Negligible | `libudev1` | `257.9-1~deb13u1` |
| [CVE-2023-31438](https://nvd.nist.gov/vuln/detail/CVE-2023-31438) | Negligible | `libudev1` | `257.9-1~deb13u1` |
| [CVE-2023-31438](https://nvd.nist.gov/vuln/detail/CVE-2023-31438) | Negligible | `libsystemd0` | `257.9-1~deb13u1` |
| [CVE-2023-31439](https://nvd.nist.gov/vuln/detail/CVE-2023-31439) | Negligible | `libudev1` | `257.9-1~deb13u1` |
| [CVE-2023-31439](https://nvd.nist.gov/vuln/detail/CVE-2023-31439) | Negligible | `libsystemd0` | `257.9-1~deb13u1` |
| [CVE-2025-5278](https://nvd.nist.gov/vuln/detail/CVE-2025-5278) | Negligible | `coreutils` | `9.7-3` |
</details>

<details>
<summary><strong>📦 Installed Packages (115 total)</strong></summary>

| Package | Version |
|---------|---------|
| Simple Launcher | 1.1.0.14 |
| adduser | 3.152 |
| apt | 3.0.3 |
| autocommand | 2.2.2 |
| backports-tarfile | 1.2.0 |
| base-files | 13.8+deb13u3 |
| base-passwd | 3.6.7 |
| bash | 5.2.37-2+b7 |
| bsdutils | 1:2.41-5 |
| ca-certificates | 20250419 |
| cli | UNKNOWN |
| cli-32 | UNKNOWN |
| cli-64 | UNKNOWN |
| cli-arm64 | UNKNOWN |
| coreutils | 9.7-3 |
| dash | 0.5.12-12 |
| debconf | 1.5.91 |
| debian-archive-keyring | 2025.1 |
| debianutils | 5.23.2 |
| diffutils | 1:3.10-4 |
| dpkg | 1.22.21 |
| findutils | 4.10.0-3 |
| gcc-14-base | 14.2.0-19 |
| grep | 3.11-4+b1 |
| gui | UNKNOWN |
| gui-32 | UNKNOWN |
| gui-64 | UNKNOWN |
| gui-arm64 | UNKNOWN |
| gzip | 1.13-1 |
| hostname | 3.25 |
| importlib-metadata | 8.0.0 |
| inflect | 7.3.1 |
| init-system-helpers | 1.69~deb13u1 |
| jaraco-collections | 5.1.0 |
| jaraco-context | 5.3.0 |
| jaraco-functools | 4.0.1 |
| jaraco-text | 3.12.1 |
| libacl1 | 2.3.2-2+b1 |
| libapt-pkg7.0 | 3.0.3 |
| libattr1 | 1:2.5.2-3 |
| libaudit-common | 1:4.0.2-2 |
| libaudit1 | 1:4.0.2-2+b2 |
| libblkid1 | 2.41-5 |
| libbsd0 | 0.12.2-2 |
| libbz2-1.0 | 1.0.8-6 |
| libc-bin | 2.41-12+deb13u1 |
| libc6 | 2.41-12+deb13u1 |
| libcap-ng0 | 0.8.5-4+b1 |
| libcap2 | 1:2.75-10+b3 |
| libcrypt1 | 1:4.4.38-1 |
| libdb5.3t64 | 5.3.28+dfsg2-9 |
| libdebconfclient0 | 0.280 |
| libffi8 | 3.4.8-2 |
| libgcc-s1 | 14.2.0-19 |
| libgdbm6t64 | 1.24-2 |
| libgmp10 | 2:6.3.0+dfsg-3 |
| libhogweed6t64 | 3.10.1-1 |
| liblastlog2-2 | 2.41-5 |
| liblz4-1 | 1.10.0-4 |
| liblzma5 | 5.8.1-1 |
| libmd0 | 1.1.0-2+b1 |
| libmount1 | 2.41-5 |
| libncursesw6 | 6.5+20250216-2 |
| libnettle8t64 | 3.10.1-1 |
| libpam-modules | 1.7.0-5 |
| libpam-modules-bin | 1.7.0-5 |
| libpam-runtime | 1.7.0-5 |
| libpam0g | 1.7.0-5 |
| libpcre2-8-0 | 10.46-1~deb13u1 |
| libreadline8t64 | 8.2-6 |
| libseccomp2 | 2.6.0-2 |
| libselinux1 | 3.8.1-1 |
| libsemanage-common | 3.8.1-1 |
| libsemanage2 | 3.8.1-1 |
| libsepol2 | 3.8.1-1 |
| libsmartcols1 | 2.41-5 |
| libsqlite3-0 | 3.46.1-7 |
| libssl3t64 | 3.5.4-1~deb13u2 |
| libstdc++6 | 14.2.0-19 |
| libsystemd0 | 257.9-1~deb13u1 |
| libtinfo6 | 6.5+20250216-2 |
| libudev1 | 257.9-1~deb13u1 |
| libuuid1 | 2.41-5 |
| libxxhash0 | 0.8.3-2 |
| libzstd1 | 1.5.7+dfsg-1 |
| login | 1:4.16.0-2+really2.41-5 |
| login.defs | 1:4.17.4-2 |
| mawk | 1.3.4.20250131-1 |
| more-itertools | 10.3.0 |
| mount | 2.41-5 |
| ncurses-base | 6.5+20250216-2 |
| ncurses-bin | 6.5+20250216-2 |
| netbase | 6.5 |
| openssl | 3.5.4-1~deb13u2 |
| openssl-provider-legacy | 3.5.4-1~deb13u2 |
| packaging | 24.2 |
| passwd | 1:4.17.4-2 |
| perl-base | 5.40.1-6 |
| pip | 24.0 |
| platformdirs | 4.2.2 |
| python | 3.11.14 |
| readline-common | 8.2-6 |
| sed | 4.9-2+b1 |
| setuptools | 79.0.1 |
| sqv | 1.3.0-3+b2 |
| sysvinit-utils | 3.14-4 |
| tar | 1.35+dfsg-3.1 |
| tomli | 2.0.1 |
| typeguard | 4.3.0 |
| typing-extensions | 4.12.2 |
| tzdata | 2025b-4+deb13u1 |
| util-linux | 2.41-5 |
| wheel | 0.45.1 |
| zipp | 3.19.2 |
| zlib1g | 1:1.3.dfsg+really1.3.1-1+b1 |
</details>

</details>

<!-- END: docker-docs:comparison -->
