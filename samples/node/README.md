# Node.js Docker Analysis

<!-- BEGIN: dock-docs:config -->

# 🐳 Docker Image Analysis: node:24
![Size](https://img.shields.io/static/v1?label=Size&message=1098.28+MB&color=blue) ![Layers](https://img.shields.io/static/v1?label=Layers&message=8&color=blue) ![Vulns](https://img.shields.io/static/v1?label=Security&message=290+Vulns+%286+Crit%29&color=red) ![Efficiency](https://img.shields.io/static/v1?label=Efficiency&message=99.4%&color=green)

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

**Base Image:** `Debian GNU/Linux 12 (bookworm) (linux/arm64)`
**Supported Architectures:** `linux/amd64, linux/arm64, linux/ppc64le, linux/s390x, unknown/unknown`
**Efficiency Score:** 99.4%

### Vulnerabilities
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| 🛑 6 | 🟠 66 | 🟡 161 | 🔵 57 |

<details>
<summary><strong>👇 Expand Vulnerability Details (1020 found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [CVE-2023-45853](https://nvd.nist.gov/vuln/detail/CVE-2023-45853) | Critical | `zlib1g-dev` | `1:1.2.13.dfsg-1` |
| [CVE-2023-5841](https://nvd.nist.gov/vuln/detail/CVE-2023-5841) | Critical | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2023-5841](https://nvd.nist.gov/vuln/detail/CVE-2023-5841) | Critical | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2023-6879](https://nvd.nist.gov/vuln/detail/CVE-2023-6879) | Critical | `libaom3` | `3.6.0-1+deb12u2` |
| [CVE-2025-7458](https://nvd.nist.gov/vuln/detail/CVE-2025-7458) | Critical | `libsqlite3-dev` | `3.40.1-2+deb12u2` |
| [CVE-2025-7458](https://nvd.nist.gov/vuln/detail/CVE-2025-7458) | Critical | `libsqlite3-0` | `3.40.1-2+deb12u2` |
| [CVE-2023-25193](https://nvd.nist.gov/vuln/detail/CVE-2023-25193) | High | `libharfbuzz0b` | `6.0.0+dfsg-3` |
| [CVE-2023-2953](https://nvd.nist.gov/vuln/detail/CVE-2023-2953) | High | `libldap-2.5-0` | `2.5.13+dfsg-5` |
| [CVE-2023-39616](https://nvd.nist.gov/vuln/detail/CVE-2023-39616) | High | `libaom3` | `3.6.0-1+deb12u2` |
| [CVE-2023-52355](https://nvd.nist.gov/vuln/detail/CVE-2023-52355) | High | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2023-52355](https://nvd.nist.gov/vuln/detail/CVE-2023-52355) | High | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2023-52355](https://nvd.nist.gov/vuln/detail/CVE-2023-52355) | High | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2025-12495](https://nvd.nist.gov/vuln/detail/CVE-2025-12495) | High | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2025-12495](https://nvd.nist.gov/vuln/detail/CVE-2025-12495) | High | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2025-12839](https://nvd.nist.gov/vuln/detail/CVE-2025-12839) | High | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2025-12839](https://nvd.nist.gov/vuln/detail/CVE-2025-12839) | High | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2025-12840](https://nvd.nist.gov/vuln/detail/CVE-2025-12840) | High | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2025-12840](https://nvd.nist.gov/vuln/detail/CVE-2025-12840) | High | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2025-13151](https://nvd.nist.gov/vuln/detail/CVE-2025-13151) | High | `libtasn1-6` | `4.19.0-2+deb12u1` |
| [CVE-2025-13699](https://nvd.nist.gov/vuln/detail/CVE-2025-13699) | High | `libmariadb-dev-compat` | `1:10.11.14-0+deb12u2` |
| [CVE-2025-13699](https://nvd.nist.gov/vuln/detail/CVE-2025-13699) | High | `libmariadb-dev` | `1:10.11.14-0+deb12u2` |
| [CVE-2025-13699](https://nvd.nist.gov/vuln/detail/CVE-2025-13699) | High | `libmariadb3` | `1:10.11.14-0+deb12u2` |
| [CVE-2025-13699](https://nvd.nist.gov/vuln/detail/CVE-2025-13699) | High | `mariadb-common` | `1:10.11.14-0+deb12u2` |
| [CVE-2025-13836](https://nvd.nist.gov/vuln/detail/CVE-2025-13836) | High | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-13836](https://nvd.nist.gov/vuln/detail/CVE-2025-13836) | High | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-13836](https://nvd.nist.gov/vuln/detail/CVE-2025-13836) | High | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-13836](https://nvd.nist.gov/vuln/detail/CVE-2025-13836) | High | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-15281](https://nvd.nist.gov/vuln/detail/CVE-2025-15281) | High | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2025-15281](https://nvd.nist.gov/vuln/detail/CVE-2025-15281) | High | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2025-15281](https://nvd.nist.gov/vuln/detail/CVE-2025-15281) | High | `libc6` | `2.36-9+deb12u13` |
| [CVE-2025-15281](https://nvd.nist.gov/vuln/detail/CVE-2025-15281) | High | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2025-46835](https://nvd.nist.gov/vuln/detail/CVE-2025-46835) | High | `git-man` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-46835](https://nvd.nist.gov/vuln/detail/CVE-2025-46835) | High | `git` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-48384](https://nvd.nist.gov/vuln/detail/CVE-2025-48384) | High | `git-man` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-48384](https://nvd.nist.gov/vuln/detail/CVE-2025-48384) | High | `git` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-48385](https://nvd.nist.gov/vuln/detail/CVE-2025-48385) | High | `git` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-48385](https://nvd.nist.gov/vuln/detail/CVE-2025-48385) | High | `git-man` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-59375](https://nvd.nist.gov/vuln/detail/CVE-2025-59375) | High | `libexpat1` | `2.5.0-1+deb12u2` |
| [CVE-2025-59375](https://nvd.nist.gov/vuln/detail/CVE-2025-59375) | High | `libexpat1-dev` | `2.5.0-1+deb12u2` |
| [CVE-2025-6297](https://nvd.nist.gov/vuln/detail/CVE-2025-6297) | High | `libdpkg-perl` | `1.21.22` |
| [CVE-2025-6297](https://nvd.nist.gov/vuln/detail/CVE-2025-6297) | High | `dpkg` | `1.21.22` |
| [CVE-2025-6297](https://nvd.nist.gov/vuln/detail/CVE-2025-6297) | High | `dpkg-dev` | `1.21.22` |
| [CVE-2025-64181](https://nvd.nist.gov/vuln/detail/CVE-2025-64181) | High | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2025-64181](https://nvd.nist.gov/vuln/detail/CVE-2025-64181) | High | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2025-7425](https://nvd.nist.gov/vuln/detail/CVE-2025-7425) | High | `libxslt1-dev` | `1.1.35-1+deb12u3` |
| [CVE-2025-7425](https://nvd.nist.gov/vuln/detail/CVE-2025-7425) | High | `libxslt1.1` | `1.1.35-1+deb12u3` |
| [CVE-2025-8194](https://nvd.nist.gov/vuln/detail/CVE-2025-8194) | High | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-8194](https://nvd.nist.gov/vuln/detail/CVE-2025-8194) | High | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-8194](https://nvd.nist.gov/vuln/detail/CVE-2025-8194) | High | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-8194](https://nvd.nist.gov/vuln/detail/CVE-2025-8194) | High | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2026-0861](https://nvd.nist.gov/vuln/detail/CVE-2026-0861) | High | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2026-0861](https://nvd.nist.gov/vuln/detail/CVE-2026-0861) | High | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2026-0861](https://nvd.nist.gov/vuln/detail/CVE-2026-0861) | High | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2026-0861](https://nvd.nist.gov/vuln/detail/CVE-2026-0861) | High | `libc6` | `2.36-9+deb12u13` |
| [CVE-2026-0915](https://nvd.nist.gov/vuln/detail/CVE-2026-0915) | High | `libc6` | `2.36-9+deb12u13` |
| [CVE-2026-0915](https://nvd.nist.gov/vuln/detail/CVE-2026-0915) | High | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2026-0915](https://nvd.nist.gov/vuln/detail/CVE-2026-0915) | High | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2026-0915](https://nvd.nist.gov/vuln/detail/CVE-2026-0915) | High | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2026-2004](https://nvd.nist.gov/vuln/detail/CVE-2026-2004) | High | `libpq5` | `15.15-0+deb12u1` |
| [CVE-2026-2004](https://nvd.nist.gov/vuln/detail/CVE-2026-2004) | High | `libpq-dev` | `15.15-0+deb12u1` |
| [CVE-2026-2005](https://nvd.nist.gov/vuln/detail/CVE-2026-2005) | High | `libpq-dev` | `15.15-0+deb12u1` |
| [CVE-2026-2005](https://nvd.nist.gov/vuln/detail/CVE-2026-2005) | High | `libpq5` | `15.15-0+deb12u1` |
| [CVE-2026-2006](https://nvd.nist.gov/vuln/detail/CVE-2026-2006) | High | `libpq5` | `15.15-0+deb12u1` |
| [CVE-2026-2006](https://nvd.nist.gov/vuln/detail/CVE-2026-2006) | High | `libpq-dev` | `15.15-0+deb12u1` |
| [CVE-2026-22695](https://nvd.nist.gov/vuln/detail/CVE-2026-22695) | High | `libpng16-16` | `1.6.39-2+deb12u1` |
| [CVE-2026-22695](https://nvd.nist.gov/vuln/detail/CVE-2026-22695) | High | `libpng-dev` | `1.6.39-2+deb12u1` |
| [CVE-2026-22801](https://nvd.nist.gov/vuln/detail/CVE-2026-22801) | High | `libpng16-16` | `1.6.39-2+deb12u1` |
| [CVE-2026-22801](https://nvd.nist.gov/vuln/detail/CVE-2026-22801) | High | `libpng-dev` | `1.6.39-2+deb12u1` |
| [CVE-2026-25646](https://nvd.nist.gov/vuln/detail/CVE-2026-25646) | High | `libpng16-16` | `1.6.39-2+deb12u1` |
| [CVE-2026-25646](https://nvd.nist.gov/vuln/detail/CVE-2026-25646) | High | `libpng-dev` | `1.6.39-2+deb12u1` |
| [GHSA-34x7-hfp2-rc4v](https://nvd.nist.gov/vuln/detail/GHSA-34x7-hfp2-rc4v) | High | `tar` | `7.5.4` |
| [GHSA-7h2j-956f-4vf2](https://nvd.nist.gov/vuln/detail/GHSA-7h2j-956f-4vf2) | High | `@isaacs/brace-expansion` | `5.0.0` |
| [CVE-2007-3996](https://nvd.nist.gov/vuln/detail/CVE-2007-3996) | Medium | `libwmf-dev` | `0.2.12-5.1` |
| [CVE-2007-3996](https://nvd.nist.gov/vuln/detail/CVE-2007-3996) | Medium | `libwmf-0.2-7` | `0.2.12-5.1` |
| [CVE-2007-3996](https://nvd.nist.gov/vuln/detail/CVE-2007-3996) | Medium | `libwmflite-0.2-7` | `0.2.12-5.1` |
| [CVE-2009-3546](https://nvd.nist.gov/vuln/detail/CVE-2009-3546) | Medium | `libwmf-0.2-7` | `0.2.12-5.1` |
| [CVE-2009-3546](https://nvd.nist.gov/vuln/detail/CVE-2009-3546) | Medium | `libwmf-dev` | `0.2.12-5.1` |
| [CVE-2009-3546](https://nvd.nist.gov/vuln/detail/CVE-2009-3546) | Medium | `libwmflite-0.2-7` | `0.2.12-5.1` |
| [CVE-2021-31879](https://nvd.nist.gov/vuln/detail/CVE-2021-31879) | Medium | `wget` | `1.21.3-1+deb12u1` |
| [CVE-2023-32570](https://nvd.nist.gov/vuln/detail/CVE-2023-32570) | Medium | `libdav1d6` | `1.0.0-2+deb12u1` |
| [CVE-2023-39327](https://nvd.nist.gov/vuln/detail/CVE-2023-39327) | Medium | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2023-39327](https://nvd.nist.gov/vuln/detail/CVE-2023-39327) | Medium | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2023-39328](https://nvd.nist.gov/vuln/detail/CVE-2023-39328) | Medium | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2023-39328](https://nvd.nist.gov/vuln/detail/CVE-2023-39328) | Medium | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2023-39329](https://nvd.nist.gov/vuln/detail/CVE-2023-39329) | Medium | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2023-39329](https://nvd.nist.gov/vuln/detail/CVE-2023-39329) | Medium | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `libncursesw6` | `6.4-4` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `libncurses6` | `6.4-4` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `ncurses-base` | `6.4-4` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `libtinfo6` | `6.4-4` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `libncurses5-dev` | `6.4-4` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `libncurses-dev` | `6.4-4` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `libncursesw5-dev` | `6.4-4` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `ncurses-bin` | `6.4-4` |
| [CVE-2023-6277](https://nvd.nist.gov/vuln/detail/CVE-2023-6277) | Medium | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2023-6277](https://nvd.nist.gov/vuln/detail/CVE-2023-6277) | Medium | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2023-6277](https://nvd.nist.gov/vuln/detail/CVE-2023-6277) | Medium | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2024-10041](https://nvd.nist.gov/vuln/detail/CVE-2024-10041) | Medium | `libpam-modules-bin` | `1.5.2-6+deb12u2` |
| [CVE-2024-10041](https://nvd.nist.gov/vuln/detail/CVE-2024-10041) | Medium | `libpam-runtime` | `1.5.2-6+deb12u2` |
| [CVE-2024-10041](https://nvd.nist.gov/vuln/detail/CVE-2024-10041) | Medium | `libpam-modules` | `1.5.2-6+deb12u2` |
| [CVE-2024-10041](https://nvd.nist.gov/vuln/detail/CVE-2024-10041) | Medium | `libpam0g` | `1.5.2-6+deb12u2` |
| [CVE-2024-10524](https://nvd.nist.gov/vuln/detail/CVE-2024-10524) | Medium | `wget` | `1.21.3-1+deb12u1` |
| [CVE-2024-38949](https://nvd.nist.gov/vuln/detail/CVE-2024-38949) | Medium | `libde265-0` | `1.0.11-1+deb12u2` |
| [CVE-2024-38950](https://nvd.nist.gov/vuln/detail/CVE-2024-38950) | Medium | `libde265-0` | `1.0.11-1+deb12u2` |
| [CVE-2025-10148](https://nvd.nist.gov/vuln/detail/CVE-2025-10148) | Medium | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-10148](https://nvd.nist.gov/vuln/detail/CVE-2025-10148) | Medium | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-10148](https://nvd.nist.gov/vuln/detail/CVE-2025-10148) | Medium | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-10148](https://nvd.nist.gov/vuln/detail/CVE-2025-10148) | Medium | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-10911](https://nvd.nist.gov/vuln/detail/CVE-2025-10911) | Medium | `libxslt1.1` | `1.1.35-1+deb12u3` |
| [CVE-2025-10911](https://nvd.nist.gov/vuln/detail/CVE-2025-10911) | Medium | `libxslt1-dev` | `1.1.35-1+deb12u3` |
| [CVE-2025-12084](https://nvd.nist.gov/vuln/detail/CVE-2025-12084) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-12084](https://nvd.nist.gov/vuln/detail/CVE-2025-12084) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-12084](https://nvd.nist.gov/vuln/detail/CVE-2025-12084) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-12084](https://nvd.nist.gov/vuln/detail/CVE-2025-12084) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-12781](https://nvd.nist.gov/vuln/detail/CVE-2025-12781) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-12781](https://nvd.nist.gov/vuln/detail/CVE-2025-12781) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-12781](https://nvd.nist.gov/vuln/detail/CVE-2025-12781) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-12781](https://nvd.nist.gov/vuln/detail/CVE-2025-12781) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-13837](https://nvd.nist.gov/vuln/detail/CVE-2025-13837) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-13837](https://nvd.nist.gov/vuln/detail/CVE-2025-13837) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-13837](https://nvd.nist.gov/vuln/detail/CVE-2025-13837) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-13837](https://nvd.nist.gov/vuln/detail/CVE-2025-13837) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `bsdutils` | `1:2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libuuid1` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libmount1` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `util-linux` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `util-linux-extra` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `uuid-dev` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libmount-dev` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libblkid1` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libblkid-dev` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `mount` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libsmartcols1` | `2.38.1-5+deb12u3` |
| [CVE-2025-14524](https://nvd.nist.gov/vuln/detail/CVE-2025-14524) | Medium | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-14524](https://nvd.nist.gov/vuln/detail/CVE-2025-14524) | Medium | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-14524](https://nvd.nist.gov/vuln/detail/CVE-2025-14524) | Medium | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-14524](https://nvd.nist.gov/vuln/detail/CVE-2025-14524) | Medium | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-14819](https://nvd.nist.gov/vuln/detail/CVE-2025-14819) | Medium | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-14819](https://nvd.nist.gov/vuln/detail/CVE-2025-14819) | Medium | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-14819](https://nvd.nist.gov/vuln/detail/CVE-2025-14819) | Medium | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-14819](https://nvd.nist.gov/vuln/detail/CVE-2025-14819) | Medium | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-14831](https://nvd.nist.gov/vuln/detail/CVE-2025-14831) | Medium | `libgnutls30` | `3.7.9-2+deb12u5` |
| [CVE-2025-15282](https://nvd.nist.gov/vuln/detail/CVE-2025-15282) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-15282](https://nvd.nist.gov/vuln/detail/CVE-2025-15282) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-15282](https://nvd.nist.gov/vuln/detail/CVE-2025-15282) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-15282](https://nvd.nist.gov/vuln/detail/CVE-2025-15282) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-15366](https://nvd.nist.gov/vuln/detail/CVE-2025-15366) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-15366](https://nvd.nist.gov/vuln/detail/CVE-2025-15366) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-15366](https://nvd.nist.gov/vuln/detail/CVE-2025-15366) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-15366](https://nvd.nist.gov/vuln/detail/CVE-2025-15366) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-15367](https://nvd.nist.gov/vuln/detail/CVE-2025-15367) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-15367](https://nvd.nist.gov/vuln/detail/CVE-2025-15367) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-15367](https://nvd.nist.gov/vuln/detail/CVE-2025-15367) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-15367](https://nvd.nist.gov/vuln/detail/CVE-2025-15367) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gpg-wks-client` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gnupg-utils` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gpgconf` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gpgsm` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `dirmngr` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gnupg` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gnupg-l10n` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gpgv` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gpg-wks-server` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gpg-agent` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gpg` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-48074](https://nvd.nist.gov/vuln/detail/CVE-2025-48074) | Medium | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2025-48074](https://nvd.nist.gov/vuln/detail/CVE-2025-48074) | Medium | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2025-6069](https://nvd.nist.gov/vuln/detail/CVE-2025-6069) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-6069](https://nvd.nist.gov/vuln/detail/CVE-2025-6069) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-6069](https://nvd.nist.gov/vuln/detail/CVE-2025-6069) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-6069](https://nvd.nist.gov/vuln/detail/CVE-2025-6069) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-6075](https://nvd.nist.gov/vuln/detail/CVE-2025-6075) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-6075](https://nvd.nist.gov/vuln/detail/CVE-2025-6075) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-6075](https://nvd.nist.gov/vuln/detail/CVE-2025-6075) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-6075](https://nvd.nist.gov/vuln/detail/CVE-2025-6075) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libncurses5-dev` | `6.4-4` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libncurses-dev` | `6.4-4` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libncursesw5-dev` | `6.4-4` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `ncurses-bin` | `6.4-4` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libtinfo6` | `6.4-4` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libncursesw6` | `6.4-4` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libncurses6` | `6.4-4` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `ncurses-base` | `6.4-4` |
| [CVE-2025-66382](https://nvd.nist.gov/vuln/detail/CVE-2025-66382) | Medium | `libexpat1-dev` | `2.5.0-1+deb12u2` |
| [CVE-2025-66382](https://nvd.nist.gov/vuln/detail/CVE-2025-66382) | Medium | `libexpat1` | `2.5.0-1+deb12u2` |
| [CVE-2025-68431](https://nvd.nist.gov/vuln/detail/CVE-2025-68431) | Medium | `libheif1` | `1.15.1-1+deb12u1` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gpgconf` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gnupg-l10n` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `dirmngr` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gnupg-utils` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gnupg` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gpgv` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gpgsm` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gpg` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gpg-wks-server` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gpg-agent` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gpg-wks-client` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-7709](https://nvd.nist.gov/vuln/detail/CVE-2025-7709) | Medium | `libsqlite3-dev` | `3.40.1-2+deb12u2` |
| [CVE-2025-7709](https://nvd.nist.gov/vuln/detail/CVE-2025-7709) | Medium | `libsqlite3-0` | `3.40.1-2+deb12u2` |
| [CVE-2025-8291](https://nvd.nist.gov/vuln/detail/CVE-2025-8291) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-8291](https://nvd.nist.gov/vuln/detail/CVE-2025-8291) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-8291](https://nvd.nist.gov/vuln/detail/CVE-2025-8291) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-8291](https://nvd.nist.gov/vuln/detail/CVE-2025-8291) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-9820](https://nvd.nist.gov/vuln/detail/CVE-2025-9820) | Medium | `libgnutls30` | `3.7.9-2+deb12u5` |
| [CVE-2026-0672](https://nvd.nist.gov/vuln/detail/CVE-2026-0672) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2026-0672](https://nvd.nist.gov/vuln/detail/CVE-2026-0672) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2026-0672](https://nvd.nist.gov/vuln/detail/CVE-2026-0672) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2026-0672](https://nvd.nist.gov/vuln/detail/CVE-2026-0672) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2026-0865](https://nvd.nist.gov/vuln/detail/CVE-2026-0865) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2026-0865](https://nvd.nist.gov/vuln/detail/CVE-2026-0865) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2026-0865](https://nvd.nist.gov/vuln/detail/CVE-2026-0865) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2026-0865](https://nvd.nist.gov/vuln/detail/CVE-2026-0865) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2026-0990](https://nvd.nist.gov/vuln/detail/CVE-2026-0990) | Medium | `libxml2` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-0990](https://nvd.nist.gov/vuln/detail/CVE-2026-0990) | Medium | `libxml2-dev` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-1299](https://nvd.nist.gov/vuln/detail/CVE-2026-1299) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2026-1299](https://nvd.nist.gov/vuln/detail/CVE-2026-1299) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2026-1299](https://nvd.nist.gov/vuln/detail/CVE-2026-1299) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2026-1299](https://nvd.nist.gov/vuln/detail/CVE-2026-1299) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2026-1484](https://nvd.nist.gov/vuln/detail/CVE-2026-1484) | Medium | `libglib2.0-dev` | `2.74.6-2+deb12u8` |
| [CVE-2026-1484](https://nvd.nist.gov/vuln/detail/CVE-2026-1484) | Medium | `libglib2.0-0` | `2.74.6-2+deb12u8` |
| [CVE-2026-1484](https://nvd.nist.gov/vuln/detail/CVE-2026-1484) | Medium | `libglib2.0-dev-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-1484](https://nvd.nist.gov/vuln/detail/CVE-2026-1484) | Medium | `libglib2.0-data` | `2.74.6-2+deb12u8` |
| [CVE-2026-1484](https://nvd.nist.gov/vuln/detail/CVE-2026-1484) | Medium | `libglib2.0-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-1489](https://nvd.nist.gov/vuln/detail/CVE-2026-1489) | Medium | `libglib2.0-dev` | `2.74.6-2+deb12u8` |
| [CVE-2026-1489](https://nvd.nist.gov/vuln/detail/CVE-2026-1489) | Medium | `libglib2.0-dev-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-1489](https://nvd.nist.gov/vuln/detail/CVE-2026-1489) | Medium | `libglib2.0-data` | `2.74.6-2+deb12u8` |
| [CVE-2026-1489](https://nvd.nist.gov/vuln/detail/CVE-2026-1489) | Medium | `libglib2.0-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-1489](https://nvd.nist.gov/vuln/detail/CVE-2026-1489) | Medium | `libglib2.0-0` | `2.74.6-2+deb12u8` |
| [CVE-2026-2003](https://nvd.nist.gov/vuln/detail/CVE-2026-2003) | Medium | `libpq5` | `15.15-0+deb12u1` |
| [CVE-2026-2003](https://nvd.nist.gov/vuln/detail/CVE-2026-2003) | Medium | `libpq-dev` | `15.15-0+deb12u1` |
| [CVE-2026-22693](https://nvd.nist.gov/vuln/detail/CVE-2026-22693) | Medium | `libharfbuzz0b` | `6.0.0+dfsg-3` |
| [CVE-2026-25210](https://nvd.nist.gov/vuln/detail/CVE-2026-25210) | Medium | `libexpat1-dev` | `2.5.0-1+deb12u2` |
| [CVE-2026-25210](https://nvd.nist.gov/vuln/detail/CVE-2026-25210) | Medium | `libexpat1` | `2.5.0-1+deb12u2` |
| [CVE-2007-3476](https://nvd.nist.gov/vuln/detail/CVE-2007-3476) | Low | `libwmf-dev` | `0.2.12-5.1` |
| [CVE-2007-3476](https://nvd.nist.gov/vuln/detail/CVE-2007-3476) | Low | `libwmflite-0.2-7` | `0.2.12-5.1` |
| [CVE-2007-3476](https://nvd.nist.gov/vuln/detail/CVE-2007-3476) | Low | `libwmf-0.2-7` | `0.2.12-5.1` |
| [CVE-2007-3477](https://nvd.nist.gov/vuln/detail/CVE-2007-3477) | Low | `libwmflite-0.2-7` | `0.2.12-5.1` |
| [CVE-2007-3477](https://nvd.nist.gov/vuln/detail/CVE-2007-3477) | Low | `libwmf-dev` | `0.2.12-5.1` |
| [CVE-2007-3477](https://nvd.nist.gov/vuln/detail/CVE-2007-3477) | Low | `libwmf-0.2-7` | `0.2.12-5.1` |
| [CVE-2016-2781](https://nvd.nist.gov/vuln/detail/CVE-2016-2781) | Low | `coreutils` | `9.1-1` |
| [CVE-2017-7475](https://nvd.nist.gov/vuln/detail/CVE-2017-7475) | Low | `libcairo2` | `1.16.0-7` |
| [CVE-2017-7475](https://nvd.nist.gov/vuln/detail/CVE-2017-7475) | Low | `libcairo2-dev` | `1.16.0-7` |
| [CVE-2017-7475](https://nvd.nist.gov/vuln/detail/CVE-2017-7475) | Low | `libcairo-gobject2` | `1.16.0-7` |
| [CVE-2017-7475](https://nvd.nist.gov/vuln/detail/CVE-2017-7475) | Low | `libcairo-script-interpreter2` | `1.16.0-7` |
| [CVE-2019-6461](https://nvd.nist.gov/vuln/detail/CVE-2019-6461) | Low | `libcairo2` | `1.16.0-7` |
| [CVE-2019-6461](https://nvd.nist.gov/vuln/detail/CVE-2019-6461) | Low | `libcairo-script-interpreter2` | `1.16.0-7` |
| [CVE-2019-6461](https://nvd.nist.gov/vuln/detail/CVE-2019-6461) | Low | `libcairo-gobject2` | `1.16.0-7` |
| [CVE-2019-6461](https://nvd.nist.gov/vuln/detail/CVE-2019-6461) | Low | `libcairo2-dev` | `1.16.0-7` |
| [CVE-2019-6462](https://nvd.nist.gov/vuln/detail/CVE-2019-6462) | Low | `libcairo2` | `1.16.0-7` |
| [CVE-2019-6462](https://nvd.nist.gov/vuln/detail/CVE-2019-6462) | Low | `libcairo-script-interpreter2` | `1.16.0-7` |
| [CVE-2019-6462](https://nvd.nist.gov/vuln/detail/CVE-2019-6462) | Low | `libcairo-gobject2` | `1.16.0-7` |
| [CVE-2019-6462](https://nvd.nist.gov/vuln/detail/CVE-2019-6462) | Low | `libcairo2-dev` | `1.16.0-7` |
| [CVE-2019-6988](https://nvd.nist.gov/vuln/detail/CVE-2019-6988) | Low | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2019-6988](https://nvd.nist.gov/vuln/detail/CVE-2019-6988) | Low | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2023-4016](https://nvd.nist.gov/vuln/detail/CVE-2023-4016) | Low | `procps` | `2:4.0.2-3` |
| [CVE-2023-4016](https://nvd.nist.gov/vuln/detail/CVE-2023-4016) | Low | `libproc2-0` | `2:4.0.2-3` |
| [CVE-2023-51792](https://nvd.nist.gov/vuln/detail/CVE-2023-51792) | Low | `libde265-0` | `1.0.11-1+deb12u2` |
| [CVE-2024-13978](https://nvd.nist.gov/vuln/detail/CVE-2024-13978) | Low | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2024-13978](https://nvd.nist.gov/vuln/detail/CVE-2024-13978) | Low | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2024-13978](https://nvd.nist.gov/vuln/detail/CVE-2024-13978) | Low | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2024-31047](https://nvd.nist.gov/vuln/detail/CVE-2024-31047) | Low | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2024-31047](https://nvd.nist.gov/vuln/detail/CVE-2024-31047) | Low | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2024-56433](https://nvd.nist.gov/vuln/detail/CVE-2024-56433) | Low | `passwd` | `1:4.13+dfsg1-1+deb12u2` |
| [CVE-2024-56433](https://nvd.nist.gov/vuln/detail/CVE-2024-56433) | Low | `login` | `1:4.13+dfsg1-1+deb12u2` |
| [CVE-2025-11731](https://nvd.nist.gov/vuln/detail/CVE-2025-11731) | Low | `libxslt1-dev` | `1.1.35-1+deb12u3` |
| [CVE-2025-11731](https://nvd.nist.gov/vuln/detail/CVE-2025-11731) | Low | `libxslt1.1` | `1.1.35-1+deb12u3` |
| [CVE-2025-27613](https://nvd.nist.gov/vuln/detail/CVE-2025-27613) | Low | `git` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-27613](https://nvd.nist.gov/vuln/detail/CVE-2025-27613) | Low | `git-man` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-50422](https://nvd.nist.gov/vuln/detail/CVE-2025-50422) | Low | `libcairo-gobject2` | `1.16.0-7` |
| [CVE-2025-50422](https://nvd.nist.gov/vuln/detail/CVE-2025-50422) | Low | `libcairo-script-interpreter2` | `1.16.0-7` |
| [CVE-2025-50422](https://nvd.nist.gov/vuln/detail/CVE-2025-50422) | Low | `libcairo2` | `1.16.0-7` |
| [CVE-2025-50422](https://nvd.nist.gov/vuln/detail/CVE-2025-50422) | Low | `libcairo2-dev` | `1.16.0-7` |
| [CVE-2025-61984](https://nvd.nist.gov/vuln/detail/CVE-2025-61984) | Low | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2025-61985](https://nvd.nist.gov/vuln/detail/CVE-2025-61985) | Low | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2026-0988](https://nvd.nist.gov/vuln/detail/CVE-2026-0988) | Low | `libglib2.0-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-0988](https://nvd.nist.gov/vuln/detail/CVE-2026-0988) | Low | `libglib2.0-dev-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-0988](https://nvd.nist.gov/vuln/detail/CVE-2026-0988) | Low | `libglib2.0-0` | `2.74.6-2+deb12u8` |
| [CVE-2026-0988](https://nvd.nist.gov/vuln/detail/CVE-2026-0988) | Low | `libglib2.0-dev` | `2.74.6-2+deb12u8` |
| [CVE-2026-0988](https://nvd.nist.gov/vuln/detail/CVE-2026-0988) | Low | `libglib2.0-data` | `2.74.6-2+deb12u8` |
| [CVE-2026-0989](https://nvd.nist.gov/vuln/detail/CVE-2026-0989) | Low | `libxml2-dev` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-0989](https://nvd.nist.gov/vuln/detail/CVE-2026-0989) | Low | `libxml2` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-0992](https://nvd.nist.gov/vuln/detail/CVE-2026-0992) | Low | `libxml2` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-0992](https://nvd.nist.gov/vuln/detail/CVE-2026-0992) | Low | `libxml2-dev` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-1485](https://nvd.nist.gov/vuln/detail/CVE-2026-1485) | Low | `libglib2.0-0` | `2.74.6-2+deb12u8` |
| [CVE-2026-1485](https://nvd.nist.gov/vuln/detail/CVE-2026-1485) | Low | `libglib2.0-dev-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-1485](https://nvd.nist.gov/vuln/detail/CVE-2026-1485) | Low | `libglib2.0-dev` | `2.74.6-2+deb12u8` |
| [CVE-2026-1485](https://nvd.nist.gov/vuln/detail/CVE-2026-1485) | Low | `libglib2.0-data` | `2.74.6-2+deb12u8` |
| [CVE-2026-1485](https://nvd.nist.gov/vuln/detail/CVE-2026-1485) | Low | `libglib2.0-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-24515](https://nvd.nist.gov/vuln/detail/CVE-2026-24515) | Low | `libexpat1-dev` | `2.5.0-1+deb12u2` |
| [CVE-2026-24515](https://nvd.nist.gov/vuln/detail/CVE-2026-24515) | Low | `libexpat1` | `2.5.0-1+deb12u2` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-2541](https://nvd.nist.gov/vuln/detail/CVE-2005-2541) | Negligible | `tar` | `1.34+dfsg-1.2+deb12u1` |
| [CVE-2007-2243](https://nvd.nist.gov/vuln/detail/CVE-2007-2243) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2007-2768](https://nvd.nist.gov/vuln/detail/CVE-2007-2768) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2007-5686](https://nvd.nist.gov/vuln/detail/CVE-2007-5686) | Negligible | `passwd` | `1:4.13+dfsg1-1+deb12u2` |
| [CVE-2007-5686](https://nvd.nist.gov/vuln/detail/CVE-2007-5686) | Negligible | `login` | `1:4.13+dfsg1-1+deb12u2` |
| [CVE-2008-1687](https://nvd.nist.gov/vuln/detail/CVE-2008-1687) | Negligible | `m4` | `1.4.19-3` |
| [CVE-2008-1688](https://nvd.nist.gov/vuln/detail/CVE-2008-1688) | Negligible | `m4` | `1.4.19-3` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3234](https://nvd.nist.gov/vuln/detail/CVE-2008-3234) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2010-4651](https://nvd.nist.gov/vuln/detail/CVE-2010-4651) | Negligible | `patch` | `2.7.6-7` |
| [CVE-2010-4756](https://nvd.nist.gov/vuln/detail/CVE-2010-4756) | Negligible | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2010-4756](https://nvd.nist.gov/vuln/detail/CVE-2010-4756) | Negligible | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2010-4756](https://nvd.nist.gov/vuln/detail/CVE-2010-4756) | Negligible | `libc6` | `2.36-9+deb12u13` |
| [CVE-2010-4756](https://nvd.nist.gov/vuln/detail/CVE-2010-4756) | Negligible | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2011-3374](https://nvd.nist.gov/vuln/detail/CVE-2011-3374) | Negligible | `libapt-pkg6.0` | `2.6.1` |
| [CVE-2011-3374](https://nvd.nist.gov/vuln/detail/CVE-2011-3374) | Negligible | `apt` | `2.6.1` |
| [CVE-2011-3389](https://nvd.nist.gov/vuln/detail/CVE-2011-3389) | Negligible | `libgnutls30` | `3.7.9-2+deb12u5` |
| [CVE-2011-4116](https://nvd.nist.gov/vuln/detail/CVE-2011-4116) | Negligible | `libperl5.36` | `5.36.0-7+deb12u3` |
| [CVE-2011-4116](https://nvd.nist.gov/vuln/detail/CVE-2011-4116) | Negligible | `perl-modules-5.36` | `5.36.0-7+deb12u3` |
| [CVE-2011-4116](https://nvd.nist.gov/vuln/detail/CVE-2011-4116) | Negligible | `perl-base` | `5.36.0-7+deb12u3` |
| [CVE-2011-4116](https://nvd.nist.gov/vuln/detail/CVE-2011-4116) | Negligible | `perl` | `5.36.0-7+deb12u3` |
| [CVE-2012-0039](https://nvd.nist.gov/vuln/detail/CVE-2012-0039) | Negligible | `libglib2.0-data` | `2.74.6-2+deb12u8` |
| [CVE-2012-0039](https://nvd.nist.gov/vuln/detail/CVE-2012-0039) | Negligible | `libglib2.0-dev-bin` | `2.74.6-2+deb12u8` |
| [CVE-2012-0039](https://nvd.nist.gov/vuln/detail/CVE-2012-0039) | Negligible | `libglib2.0-dev` | `2.74.6-2+deb12u8` |
| [CVE-2012-0039](https://nvd.nist.gov/vuln/detail/CVE-2012-0039) | Negligible | `libglib2.0-bin` | `2.74.6-2+deb12u8` |
| [CVE-2012-0039](https://nvd.nist.gov/vuln/detail/CVE-2012-0039) | Negligible | `libglib2.0-0` | `2.74.6-2+deb12u8` |
| [CVE-2013-4392](https://nvd.nist.gov/vuln/detail/CVE-2013-4392) | Negligible | `libudev1` | `252.39-1~deb12u1` |
| [CVE-2013-4392](https://nvd.nist.gov/vuln/detail/CVE-2013-4392) | Negligible | `libsystemd0` | `252.39-1~deb12u1` |
| [CVE-2015-3276](https://nvd.nist.gov/vuln/detail/CVE-2015-3276) | Negligible | `libldap-2.5-0` | `2.5.13+dfsg-5` |
| [CVE-2015-9019](https://nvd.nist.gov/vuln/detail/CVE-2015-9019) | Negligible | `libxslt1.1` | `1.1.35-1+deb12u3` |
| [CVE-2015-9019](https://nvd.nist.gov/vuln/detail/CVE-2015-9019) | Negligible | `libxslt1-dev` | `1.1.35-1+deb12u3` |
| [CVE-2016-10505](https://nvd.nist.gov/vuln/detail/CVE-2016-10505) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2016-10505](https://nvd.nist.gov/vuln/detail/CVE-2016-10505) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-20012](https://nvd.nist.gov/vuln/detail/CVE-2016-20012) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-9113](https://nvd.nist.gov/vuln/detail/CVE-2016-9113) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-9113](https://nvd.nist.gov/vuln/detail/CVE-2016-9113) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2016-9114](https://nvd.nist.gov/vuln/detail/CVE-2016-9114) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2016-9114](https://nvd.nist.gov/vuln/detail/CVE-2016-9114) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-9115](https://nvd.nist.gov/vuln/detail/CVE-2016-9115) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2016-9115](https://nvd.nist.gov/vuln/detail/CVE-2016-9115) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-9116](https://nvd.nist.gov/vuln/detail/CVE-2016-9116) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2016-9116](https://nvd.nist.gov/vuln/detail/CVE-2016-9116) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-9117](https://nvd.nist.gov/vuln/detail/CVE-2016-9117) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2016-9117](https://nvd.nist.gov/vuln/detail/CVE-2016-9117) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-9580](https://nvd.nist.gov/vuln/detail/CVE-2016-9580) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-9580](https://nvd.nist.gov/vuln/detail/CVE-2016-9580) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2016-9581](https://nvd.nist.gov/vuln/detail/CVE-2016-9581) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-9581](https://nvd.nist.gov/vuln/detail/CVE-2016-9581) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-13716](https://nvd.nist.gov/vuln/detail/CVE-2017-13716) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2017-13716](https://nvd.nist.gov/vuln/detail/CVE-2017-13716) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2017-13716](https://nvd.nist.gov/vuln/detail/CVE-2017-13716) | Negligible | `binutils` | `2.40-2` |
| [CVE-2017-13716](https://nvd.nist.gov/vuln/detail/CVE-2017-13716) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2017-13716](https://nvd.nist.gov/vuln/detail/CVE-2017-13716) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2017-13716](https://nvd.nist.gov/vuln/detail/CVE-2017-13716) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2017-13716](https://nvd.nist.gov/vuln/detail/CVE-2017-13716) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2017-14159](https://nvd.nist.gov/vuln/detail/CVE-2017-14159) | Negligible | `libldap-2.5-0` | `2.5.13+dfsg-5` |
| [CVE-2017-14988](https://nvd.nist.gov/vuln/detail/CVE-2017-14988) | Negligible | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2017-14988](https://nvd.nist.gov/vuln/detail/CVE-2017-14988) | Negligible | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2017-16232](https://nvd.nist.gov/vuln/detail/CVE-2017-16232) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2017-16232](https://nvd.nist.gov/vuln/detail/CVE-2017-16232) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2017-16232](https://nvd.nist.gov/vuln/detail/CVE-2017-16232) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2017-17740](https://nvd.nist.gov/vuln/detail/CVE-2017-17740) | Negligible | `libldap-2.5-0` | `2.5.13+dfsg-5` |
| [CVE-2017-18018](https://nvd.nist.gov/vuln/detail/CVE-2017-18018) | Negligible | `coreutils` | `9.1-1` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-9937](https://nvd.nist.gov/vuln/detail/CVE-2017-9937) | Negligible | `libjbig-dev` | `2.1-6.1` |
| [CVE-2017-9937](https://nvd.nist.gov/vuln/detail/CVE-2017-9937) | Negligible | `libjbig0` | `2.1-6.1` |
| [CVE-2018-1000021](https://nvd.nist.gov/vuln/detail/CVE-2018-1000021) | Negligible | `git-man` | `1:2.39.5-0+deb12u3` |
| [CVE-2018-1000021](https://nvd.nist.gov/vuln/detail/CVE-2018-1000021) | Negligible | `git` | `1:2.39.5-0+deb12u3` |
| [CVE-2018-10126](https://nvd.nist.gov/vuln/detail/CVE-2018-10126) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2018-10126](https://nvd.nist.gov/vuln/detail/CVE-2018-10126) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2018-10126](https://nvd.nist.gov/vuln/detail/CVE-2018-10126) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15919](https://nvd.nist.gov/vuln/detail/CVE-2018-15919) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2018-16376](https://nvd.nist.gov/vuln/detail/CVE-2018-16376) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2018-16376](https://nvd.nist.gov/vuln/detail/CVE-2018-16376) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2018-18064](https://nvd.nist.gov/vuln/detail/CVE-2018-18064) | Negligible | `libcairo-script-interpreter2` | `1.16.0-7` |
| [CVE-2018-18064](https://nvd.nist.gov/vuln/detail/CVE-2018-18064) | Negligible | `libcairo2-dev` | `1.16.0-7` |
| [CVE-2018-18064](https://nvd.nist.gov/vuln/detail/CVE-2018-18064) | Negligible | `libcairo2` | `1.16.0-7` |
| [CVE-2018-18064](https://nvd.nist.gov/vuln/detail/CVE-2018-18064) | Negligible | `libcairo-gobject2` | `1.16.0-7` |
| [CVE-2018-20673](https://nvd.nist.gov/vuln/detail/CVE-2018-20673) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2018-20673](https://nvd.nist.gov/vuln/detail/CVE-2018-20673) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2018-20673](https://nvd.nist.gov/vuln/detail/CVE-2018-20673) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2018-20673](https://nvd.nist.gov/vuln/detail/CVE-2018-20673) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2018-20673](https://nvd.nist.gov/vuln/detail/CVE-2018-20673) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2018-20673](https://nvd.nist.gov/vuln/detail/CVE-2018-20673) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2018-20673](https://nvd.nist.gov/vuln/detail/CVE-2018-20673) | Negligible | `binutils` | `2.40-2` |
| [CVE-2018-20712](https://nvd.nist.gov/vuln/detail/CVE-2018-20712) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2018-20712](https://nvd.nist.gov/vuln/detail/CVE-2018-20712) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2018-20712](https://nvd.nist.gov/vuln/detail/CVE-2018-20712) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2018-20712](https://nvd.nist.gov/vuln/detail/CVE-2018-20712) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2018-20712](https://nvd.nist.gov/vuln/detail/CVE-2018-20712) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2018-20712](https://nvd.nist.gov/vuln/detail/CVE-2018-20712) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2018-20712](https://nvd.nist.gov/vuln/detail/CVE-2018-20712) | Negligible | `binutils` | `2.40-2` |
| [CVE-2018-20796](https://nvd.nist.gov/vuln/detail/CVE-2018-20796) | Negligible | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2018-20796](https://nvd.nist.gov/vuln/detail/CVE-2018-20796) | Negligible | `libc6` | `2.36-9+deb12u13` |
| [CVE-2018-20796](https://nvd.nist.gov/vuln/detail/CVE-2018-20796) | Negligible | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2018-20796](https://nvd.nist.gov/vuln/detail/CVE-2018-20796) | Negligible | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libkdb5-10` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libkrb5-dev` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libkrb5support0` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libkrb5-3` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libkadm5srv-mit12` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libkadm5clnt-mit12` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libk5crypto3` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libgssrpc4` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libgssapi-krb5-2` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `krb5-multidev` | `1.20.1-2+deb12u4` |
| [CVE-2018-6829](https://nvd.nist.gov/vuln/detail/CVE-2018-6829) | Negligible | `libgcrypt20` | `1.10.1-3` |
| [CVE-2018-6951](https://nvd.nist.gov/vuln/detail/CVE-2018-6951) | Negligible | `patch` | `2.7.6-7` |
| [CVE-2018-6952](https://nvd.nist.gov/vuln/detail/CVE-2018-6952) | Negligible | `patch` | `2.7.6-7` |
| [CVE-2018-9996](https://nvd.nist.gov/vuln/detail/CVE-2018-9996) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2018-9996](https://nvd.nist.gov/vuln/detail/CVE-2018-9996) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2018-9996](https://nvd.nist.gov/vuln/detail/CVE-2018-9996) | Negligible | `binutils` | `2.40-2` |
| [CVE-2018-9996](https://nvd.nist.gov/vuln/detail/CVE-2018-9996) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2018-9996](https://nvd.nist.gov/vuln/detail/CVE-2018-9996) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2018-9996](https://nvd.nist.gov/vuln/detail/CVE-2018-9996) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2018-9996](https://nvd.nist.gov/vuln/detail/CVE-2018-9996) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2019-1010022](https://nvd.nist.gov/vuln/detail/CVE-2019-1010022) | Negligible | `libc6` | `2.36-9+deb12u13` |
| [CVE-2019-1010022](https://nvd.nist.gov/vuln/detail/CVE-2019-1010022) | Negligible | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2019-1010022](https://nvd.nist.gov/vuln/detail/CVE-2019-1010022) | Negligible | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2019-1010022](https://nvd.nist.gov/vuln/detail/CVE-2019-1010022) | Negligible | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2019-1010023](https://nvd.nist.gov/vuln/detail/CVE-2019-1010023) | Negligible | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2019-1010023](https://nvd.nist.gov/vuln/detail/CVE-2019-1010023) | Negligible | `libc6` | `2.36-9+deb12u13` |
| [CVE-2019-1010023](https://nvd.nist.gov/vuln/detail/CVE-2019-1010023) | Negligible | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2019-1010023](https://nvd.nist.gov/vuln/detail/CVE-2019-1010023) | Negligible | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2019-1010024](https://nvd.nist.gov/vuln/detail/CVE-2019-1010024) | Negligible | `libc6` | `2.36-9+deb12u13` |
| [CVE-2019-1010024](https://nvd.nist.gov/vuln/detail/CVE-2019-1010024) | Negligible | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2019-1010024](https://nvd.nist.gov/vuln/detail/CVE-2019-1010024) | Negligible | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2019-1010024](https://nvd.nist.gov/vuln/detail/CVE-2019-1010024) | Negligible | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2019-1010025](https://nvd.nist.gov/vuln/detail/CVE-2019-1010025) | Negligible | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2019-1010025](https://nvd.nist.gov/vuln/detail/CVE-2019-1010025) | Negligible | `libc6` | `2.36-9+deb12u13` |
| [CVE-2019-1010025](https://nvd.nist.gov/vuln/detail/CVE-2019-1010025) | Negligible | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2019-1010025](https://nvd.nist.gov/vuln/detail/CVE-2019-1010025) | Negligible | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2019-6110](https://nvd.nist.gov/vuln/detail/CVE-2019-6110) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2019-9192](https://nvd.nist.gov/vuln/detail/CVE-2019-9192) | Negligible | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2019-9192](https://nvd.nist.gov/vuln/detail/CVE-2019-9192) | Negligible | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2019-9192](https://nvd.nist.gov/vuln/detail/CVE-2019-9192) | Negligible | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2019-9192](https://nvd.nist.gov/vuln/detail/CVE-2019-9192) | Negligible | `libc6` | `2.36-9+deb12u13` |
| [CVE-2020-14145](https://nvd.nist.gov/vuln/detail/CVE-2020-14145) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2020-15719](https://nvd.nist.gov/vuln/detail/CVE-2020-15719) | Negligible | `libldap-2.5-0` | `2.5.13+dfsg-5` |
| [CVE-2020-15778](https://nvd.nist.gov/vuln/detail/CVE-2020-15778) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2020-36325](https://nvd.nist.gov/vuln/detail/CVE-2020-36325) | Negligible | `libjansson4` | `2.14-2` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-32256](https://nvd.nist.gov/vuln/detail/CVE-2021-32256) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2021-32256](https://nvd.nist.gov/vuln/detail/CVE-2021-32256) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2021-32256](https://nvd.nist.gov/vuln/detail/CVE-2021-32256) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2021-32256](https://nvd.nist.gov/vuln/detail/CVE-2021-32256) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2021-32256](https://nvd.nist.gov/vuln/detail/CVE-2021-32256) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2021-32256](https://nvd.nist.gov/vuln/detail/CVE-2021-32256) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2021-32256](https://nvd.nist.gov/vuln/detail/CVE-2021-32256) | Negligible | `binutils` | `2.40-2` |
| [CVE-2021-4214](https://nvd.nist.gov/vuln/detail/CVE-2021-4214) | Negligible | `libpng16-16` | `1.6.39-2+deb12u1` |
| [CVE-2021-4214](https://nvd.nist.gov/vuln/detail/CVE-2021-4214) | Negligible | `libpng-dev` | `1.6.39-2+deb12u1` |
| [CVE-2021-4217](https://nvd.nist.gov/vuln/detail/CVE-2021-4217) | Negligible | `unzip` | `6.0-28` |
| [CVE-2021-45261](https://nvd.nist.gov/vuln/detail/CVE-2021-45261) | Negligible | `patch` | `2.7.6-7` |
| [CVE-2021-45346](https://nvd.nist.gov/vuln/detail/CVE-2021-45346) | Negligible | `libsqlite3-0` | `3.40.1-2+deb12u2` |
| [CVE-2021-45346](https://nvd.nist.gov/vuln/detail/CVE-2021-45346) | Negligible | `libsqlite3-dev` | `3.40.1-2+deb12u2` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libmount1` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `util-linux-extra` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `bsdutils` | `1:2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libsmartcols1` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libuuid1` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `mount` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `util-linux` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libmount-dev` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libblkid-dev` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `uuid-dev` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libblkid1` | `2.38.1-5+deb12u3` |
| [CVE-2022-1210](https://nvd.nist.gov/vuln/detail/CVE-2022-1210) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2022-1210](https://nvd.nist.gov/vuln/detail/CVE-2022-1210) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2022-1210](https://nvd.nist.gov/vuln/detail/CVE-2022-1210) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2022-24975](https://nvd.nist.gov/vuln/detail/CVE-2022-24975) | Negligible | `git-man` | `1:2.39.5-0+deb12u3` |
| [CVE-2022-24975](https://nvd.nist.gov/vuln/detail/CVE-2022-24975) | Negligible | `git` | `1:2.39.5-0+deb12u3` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libhwasan0` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libgcc-s1` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libtsan2` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libstdc++6` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libstdc++-12-dev` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `liblsan0` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libitm1` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libgomp1` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libubsan1` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `cpp-12` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `g++-12` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `gcc-12` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `gcc-12-base` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libgcc-12-dev` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libcc1-0` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libatomic1` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libasan8` | `12.2.0-14+deb12u1` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `dirmngr` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gpg-agent` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gpgv` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gpg-wks-client` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gnupg-utils` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gnupg-l10n` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gnupg` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gpgsm` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gpg` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gpg-wks-server` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gpgconf` | `2.2.40-1.1+deb12u2` |
| [CVE-2023-1916](https://nvd.nist.gov/vuln/detail/CVE-2023-1916) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2023-1916](https://nvd.nist.gov/vuln/detail/CVE-2023-1916) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2023-1916](https://nvd.nist.gov/vuln/detail/CVE-2023-1916) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2023-1972](https://nvd.nist.gov/vuln/detail/CVE-2023-1972) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2023-1972](https://nvd.nist.gov/vuln/detail/CVE-2023-1972) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2023-1972](https://nvd.nist.gov/vuln/detail/CVE-2023-1972) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2023-1972](https://nvd.nist.gov/vuln/detail/CVE-2023-1972) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2023-1972](https://nvd.nist.gov/vuln/detail/CVE-2023-1972) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2023-1972](https://nvd.nist.gov/vuln/detail/CVE-2023-1972) | Negligible | `binutils` | `2.40-2` |
| [CVE-2023-1972](https://nvd.nist.gov/vuln/detail/CVE-2023-1972) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2023-31437](https://nvd.nist.gov/vuln/detail/CVE-2023-31437) | Negligible | `libsystemd0` | `252.39-1~deb12u1` |
| [CVE-2023-31437](https://nvd.nist.gov/vuln/detail/CVE-2023-31437) | Negligible | `libudev1` | `252.39-1~deb12u1` |
| [CVE-2023-31438](https://nvd.nist.gov/vuln/detail/CVE-2023-31438) | Negligible | `libudev1` | `252.39-1~deb12u1` |
| [CVE-2023-31438](https://nvd.nist.gov/vuln/detail/CVE-2023-31438) | Negligible | `libsystemd0` | `252.39-1~deb12u1` |
| [CVE-2023-31439](https://nvd.nist.gov/vuln/detail/CVE-2023-31439) | Negligible | `libudev1` | `252.39-1~deb12u1` |
| [CVE-2023-31439](https://nvd.nist.gov/vuln/detail/CVE-2023-31439) | Negligible | `libsystemd0` | `252.39-1~deb12u1` |
| [CVE-2023-31486](https://nvd.nist.gov/vuln/detail/CVE-2023-31486) | Negligible | `perl-modules-5.36` | `5.36.0-7+deb12u3` |
| [CVE-2023-31486](https://nvd.nist.gov/vuln/detail/CVE-2023-31486) | Negligible | `perl-base` | `5.36.0-7+deb12u3` |
| [CVE-2023-31486](https://nvd.nist.gov/vuln/detail/CVE-2023-31486) | Negligible | `perl` | `5.36.0-7+deb12u3` |
| [CVE-2023-31486](https://nvd.nist.gov/vuln/detail/CVE-2023-31486) | Negligible | `libperl5.36` | `5.36.0-7+deb12u3` |
| [CVE-2023-3164](https://nvd.nist.gov/vuln/detail/CVE-2023-3164) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2023-3164](https://nvd.nist.gov/vuln/detail/CVE-2023-3164) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2023-3164](https://nvd.nist.gov/vuln/detail/CVE-2023-3164) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-37769](https://nvd.nist.gov/vuln/detail/CVE-2023-37769) | Negligible | `libpixman-1-0` | `0.42.2-1` |
| [CVE-2023-37769](https://nvd.nist.gov/vuln/detail/CVE-2023-37769) | Negligible | `libpixman-1-dev` | `0.42.2-1` |
| [CVE-2023-49463](https://nvd.nist.gov/vuln/detail/CVE-2023-49463) | Negligible | `libheif1` | `1.15.1-1+deb12u1` |
| [CVE-2023-52426](https://nvd.nist.gov/vuln/detail/CVE-2023-52426) | Negligible | `libexpat1-dev` | `2.5.0-1+deb12u2` |
| [CVE-2023-52426](https://nvd.nist.gov/vuln/detail/CVE-2023-52426) | Negligible | `libexpat1` | `2.5.0-1+deb12u2` |
| [CVE-2023-6228](https://nvd.nist.gov/vuln/detail/CVE-2023-6228) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2023-6228](https://nvd.nist.gov/vuln/detail/CVE-2023-6228) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2023-6228](https://nvd.nist.gov/vuln/detail/CVE-2023-6228) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2024-2236](https://nvd.nist.gov/vuln/detail/CVE-2024-2236) | Negligible | `libgcrypt20` | `1.10.1-3` |
| [CVE-2024-2379](https://nvd.nist.gov/vuln/detail/CVE-2024-2379) | Negligible | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2024-2379](https://nvd.nist.gov/vuln/detail/CVE-2024-2379) | Negligible | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2024-2379](https://nvd.nist.gov/vuln/detail/CVE-2024-2379) | Negligible | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2024-2379](https://nvd.nist.gov/vuln/detail/CVE-2024-2379) | Negligible | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2024-25260](https://nvd.nist.gov/vuln/detail/CVE-2024-25260) | Negligible | `libelf1` | `0.188-2.1` |
| [CVE-2024-25269](https://nvd.nist.gov/vuln/detail/CVE-2024-25269) | Negligible | `libheif1` | `1.15.1-1+deb12u1` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libgssrpc4` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libkadm5clnt-mit12` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libk5crypto3` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libgssapi-krb5-2` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libkrb5support0` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libkrb5-dev` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libkrb5-3` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libkdb5-10` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `krb5-multidev` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libkadm5srv-mit12` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libkdb5-10` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libkrb5support0` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `krb5-multidev` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libgssrpc4` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libk5crypto3` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libkadm5clnt-mit12` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libkadm5srv-mit12` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libgssapi-krb5-2` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libkrb5-3` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libkrb5-dev` | `1.20.1-2+deb12u4` |
| [CVE-2024-28757](https://nvd.nist.gov/vuln/detail/CVE-2024-28757) | Negligible | `libexpat1` | `2.5.0-1+deb12u2` |
| [CVE-2024-28757](https://nvd.nist.gov/vuln/detail/CVE-2024-28757) | Negligible | `libexpat1-dev` | `2.5.0-1+deb12u2` |
| [CVE-2024-52005](https://nvd.nist.gov/vuln/detail/CVE-2024-52005) | Negligible | `git-man` | `1:2.39.5-0+deb12u3` |
| [CVE-2024-52005](https://nvd.nist.gov/vuln/detail/CVE-2024-52005) | Negligible | `git` | `1:2.39.5-0+deb12u3` |
| [CVE-2024-53589](https://nvd.nist.gov/vuln/detail/CVE-2024-53589) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2024-53589](https://nvd.nist.gov/vuln/detail/CVE-2024-53589) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2024-53589](https://nvd.nist.gov/vuln/detail/CVE-2024-53589) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2024-53589](https://nvd.nist.gov/vuln/detail/CVE-2024-53589) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2024-53589](https://nvd.nist.gov/vuln/detail/CVE-2024-53589) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2024-53589](https://nvd.nist.gov/vuln/detail/CVE-2024-53589) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2024-53589](https://nvd.nist.gov/vuln/detail/CVE-2024-53589) | Negligible | `binutils` | `2.40-2` |
| [CVE-2024-57360](https://nvd.nist.gov/vuln/detail/CVE-2024-57360) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2024-57360](https://nvd.nist.gov/vuln/detail/CVE-2024-57360) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2024-57360](https://nvd.nist.gov/vuln/detail/CVE-2024-57360) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2024-57360](https://nvd.nist.gov/vuln/detail/CVE-2024-57360) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2024-57360](https://nvd.nist.gov/vuln/detail/CVE-2024-57360) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2024-57360](https://nvd.nist.gov/vuln/detail/CVE-2024-57360) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2024-57360](https://nvd.nist.gov/vuln/detail/CVE-2024-57360) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-0725](https://nvd.nist.gov/vuln/detail/CVE-2025-0725) | Negligible | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-0725](https://nvd.nist.gov/vuln/detail/CVE-2025-0725) | Negligible | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-0725](https://nvd.nist.gov/vuln/detail/CVE-2025-0725) | Negligible | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-0725](https://nvd.nist.gov/vuln/detail/CVE-2025-0725) | Negligible | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-0840](https://nvd.nist.gov/vuln/detail/CVE-2025-0840) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-0840](https://nvd.nist.gov/vuln/detail/CVE-2025-0840) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-0840](https://nvd.nist.gov/vuln/detail/CVE-2025-0840) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-0840](https://nvd.nist.gov/vuln/detail/CVE-2025-0840) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-0840](https://nvd.nist.gov/vuln/detail/CVE-2025-0840) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-0840](https://nvd.nist.gov/vuln/detail/CVE-2025-0840) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-0840](https://nvd.nist.gov/vuln/detail/CVE-2025-0840) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-10966](https://nvd.nist.gov/vuln/detail/CVE-2025-10966) | Negligible | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-10966](https://nvd.nist.gov/vuln/detail/CVE-2025-10966) | Negligible | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-10966](https://nvd.nist.gov/vuln/detail/CVE-2025-10966) | Negligible | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-10966](https://nvd.nist.gov/vuln/detail/CVE-2025-10966) | Negligible | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-11081](https://nvd.nist.gov/vuln/detail/CVE-2025-11081) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11081](https://nvd.nist.gov/vuln/detail/CVE-2025-11081) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11081](https://nvd.nist.gov/vuln/detail/CVE-2025-11081) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11081](https://nvd.nist.gov/vuln/detail/CVE-2025-11081) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11081](https://nvd.nist.gov/vuln/detail/CVE-2025-11081) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11081](https://nvd.nist.gov/vuln/detail/CVE-2025-11081) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11081](https://nvd.nist.gov/vuln/detail/CVE-2025-11081) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11082](https://nvd.nist.gov/vuln/detail/CVE-2025-11082) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11082](https://nvd.nist.gov/vuln/detail/CVE-2025-11082) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11082](https://nvd.nist.gov/vuln/detail/CVE-2025-11082) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11082](https://nvd.nist.gov/vuln/detail/CVE-2025-11082) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11082](https://nvd.nist.gov/vuln/detail/CVE-2025-11082) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11082](https://nvd.nist.gov/vuln/detail/CVE-2025-11082) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11082](https://nvd.nist.gov/vuln/detail/CVE-2025-11082) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11083](https://nvd.nist.gov/vuln/detail/CVE-2025-11083) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11083](https://nvd.nist.gov/vuln/detail/CVE-2025-11083) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11083](https://nvd.nist.gov/vuln/detail/CVE-2025-11083) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11083](https://nvd.nist.gov/vuln/detail/CVE-2025-11083) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11083](https://nvd.nist.gov/vuln/detail/CVE-2025-11083) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11083](https://nvd.nist.gov/vuln/detail/CVE-2025-11083) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11083](https://nvd.nist.gov/vuln/detail/CVE-2025-11083) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11412](https://nvd.nist.gov/vuln/detail/CVE-2025-11412) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11412](https://nvd.nist.gov/vuln/detail/CVE-2025-11412) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11412](https://nvd.nist.gov/vuln/detail/CVE-2025-11412) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11412](https://nvd.nist.gov/vuln/detail/CVE-2025-11412) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11412](https://nvd.nist.gov/vuln/detail/CVE-2025-11412) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11412](https://nvd.nist.gov/vuln/detail/CVE-2025-11412) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11412](https://nvd.nist.gov/vuln/detail/CVE-2025-11412) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11413](https://nvd.nist.gov/vuln/detail/CVE-2025-11413) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11413](https://nvd.nist.gov/vuln/detail/CVE-2025-11413) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11413](https://nvd.nist.gov/vuln/detail/CVE-2025-11413) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11413](https://nvd.nist.gov/vuln/detail/CVE-2025-11413) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11413](https://nvd.nist.gov/vuln/detail/CVE-2025-11413) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11413](https://nvd.nist.gov/vuln/detail/CVE-2025-11413) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11413](https://nvd.nist.gov/vuln/detail/CVE-2025-11413) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11414](https://nvd.nist.gov/vuln/detail/CVE-2025-11414) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11414](https://nvd.nist.gov/vuln/detail/CVE-2025-11414) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11414](https://nvd.nist.gov/vuln/detail/CVE-2025-11414) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11414](https://nvd.nist.gov/vuln/detail/CVE-2025-11414) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11414](https://nvd.nist.gov/vuln/detail/CVE-2025-11414) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11414](https://nvd.nist.gov/vuln/detail/CVE-2025-11414) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11414](https://nvd.nist.gov/vuln/detail/CVE-2025-11414) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11468](https://nvd.nist.gov/vuln/detail/CVE-2025-11468) | Negligible | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-11468](https://nvd.nist.gov/vuln/detail/CVE-2025-11468) | Negligible | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-11468](https://nvd.nist.gov/vuln/detail/CVE-2025-11468) | Negligible | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-11468](https://nvd.nist.gov/vuln/detail/CVE-2025-11468) | Negligible | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-1147](https://nvd.nist.gov/vuln/detail/CVE-2025-1147) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1147](https://nvd.nist.gov/vuln/detail/CVE-2025-1147) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1147](https://nvd.nist.gov/vuln/detail/CVE-2025-1147) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1147](https://nvd.nist.gov/vuln/detail/CVE-2025-1147) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1147](https://nvd.nist.gov/vuln/detail/CVE-2025-1147) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1147](https://nvd.nist.gov/vuln/detail/CVE-2025-1147) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1147](https://nvd.nist.gov/vuln/detail/CVE-2025-1147) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1148](https://nvd.nist.gov/vuln/detail/CVE-2025-1148) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1148](https://nvd.nist.gov/vuln/detail/CVE-2025-1148) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1148](https://nvd.nist.gov/vuln/detail/CVE-2025-1148) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1148](https://nvd.nist.gov/vuln/detail/CVE-2025-1148) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1148](https://nvd.nist.gov/vuln/detail/CVE-2025-1148) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1148](https://nvd.nist.gov/vuln/detail/CVE-2025-1148) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1148](https://nvd.nist.gov/vuln/detail/CVE-2025-1148) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1149](https://nvd.nist.gov/vuln/detail/CVE-2025-1149) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1149](https://nvd.nist.gov/vuln/detail/CVE-2025-1149) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1149](https://nvd.nist.gov/vuln/detail/CVE-2025-1149) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1149](https://nvd.nist.gov/vuln/detail/CVE-2025-1149) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1149](https://nvd.nist.gov/vuln/detail/CVE-2025-1149) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1149](https://nvd.nist.gov/vuln/detail/CVE-2025-1149) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1149](https://nvd.nist.gov/vuln/detail/CVE-2025-1149) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11494](https://nvd.nist.gov/vuln/detail/CVE-2025-11494) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11494](https://nvd.nist.gov/vuln/detail/CVE-2025-11494) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11494](https://nvd.nist.gov/vuln/detail/CVE-2025-11494) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11494](https://nvd.nist.gov/vuln/detail/CVE-2025-11494) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11494](https://nvd.nist.gov/vuln/detail/CVE-2025-11494) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11494](https://nvd.nist.gov/vuln/detail/CVE-2025-11494) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11494](https://nvd.nist.gov/vuln/detail/CVE-2025-11494) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11495](https://nvd.nist.gov/vuln/detail/CVE-2025-11495) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11495](https://nvd.nist.gov/vuln/detail/CVE-2025-11495) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11495](https://nvd.nist.gov/vuln/detail/CVE-2025-11495) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11495](https://nvd.nist.gov/vuln/detail/CVE-2025-11495) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11495](https://nvd.nist.gov/vuln/detail/CVE-2025-11495) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11495](https://nvd.nist.gov/vuln/detail/CVE-2025-11495) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11495](https://nvd.nist.gov/vuln/detail/CVE-2025-11495) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1150](https://nvd.nist.gov/vuln/detail/CVE-2025-1150) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1150](https://nvd.nist.gov/vuln/detail/CVE-2025-1150) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1150](https://nvd.nist.gov/vuln/detail/CVE-2025-1150) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1150](https://nvd.nist.gov/vuln/detail/CVE-2025-1150) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1150](https://nvd.nist.gov/vuln/detail/CVE-2025-1150) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1150](https://nvd.nist.gov/vuln/detail/CVE-2025-1150) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1150](https://nvd.nist.gov/vuln/detail/CVE-2025-1150) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1151](https://nvd.nist.gov/vuln/detail/CVE-2025-1151) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1151](https://nvd.nist.gov/vuln/detail/CVE-2025-1151) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1151](https://nvd.nist.gov/vuln/detail/CVE-2025-1151) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1151](https://nvd.nist.gov/vuln/detail/CVE-2025-1151) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1151](https://nvd.nist.gov/vuln/detail/CVE-2025-1151) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1151](https://nvd.nist.gov/vuln/detail/CVE-2025-1151) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1151](https://nvd.nist.gov/vuln/detail/CVE-2025-1151) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1152](https://nvd.nist.gov/vuln/detail/CVE-2025-1152) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1152](https://nvd.nist.gov/vuln/detail/CVE-2025-1152) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1152](https://nvd.nist.gov/vuln/detail/CVE-2025-1152) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1152](https://nvd.nist.gov/vuln/detail/CVE-2025-1152) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1152](https://nvd.nist.gov/vuln/detail/CVE-2025-1152) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1152](https://nvd.nist.gov/vuln/detail/CVE-2025-1152) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1152](https://nvd.nist.gov/vuln/detail/CVE-2025-1152) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1153](https://nvd.nist.gov/vuln/detail/CVE-2025-1153) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1153](https://nvd.nist.gov/vuln/detail/CVE-2025-1153) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1153](https://nvd.nist.gov/vuln/detail/CVE-2025-1153) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1153](https://nvd.nist.gov/vuln/detail/CVE-2025-1153) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1153](https://nvd.nist.gov/vuln/detail/CVE-2025-1153) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1153](https://nvd.nist.gov/vuln/detail/CVE-2025-1153) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1153](https://nvd.nist.gov/vuln/detail/CVE-2025-1153) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1176](https://nvd.nist.gov/vuln/detail/CVE-2025-1176) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1176](https://nvd.nist.gov/vuln/detail/CVE-2025-1176) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1176](https://nvd.nist.gov/vuln/detail/CVE-2025-1176) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1176](https://nvd.nist.gov/vuln/detail/CVE-2025-1176) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1176](https://nvd.nist.gov/vuln/detail/CVE-2025-1176) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1176](https://nvd.nist.gov/vuln/detail/CVE-2025-1176) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1176](https://nvd.nist.gov/vuln/detail/CVE-2025-1176) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1178](https://nvd.nist.gov/vuln/detail/CVE-2025-1178) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1178](https://nvd.nist.gov/vuln/detail/CVE-2025-1178) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1178](https://nvd.nist.gov/vuln/detail/CVE-2025-1178) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1178](https://nvd.nist.gov/vuln/detail/CVE-2025-1178) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1178](https://nvd.nist.gov/vuln/detail/CVE-2025-1178) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1178](https://nvd.nist.gov/vuln/detail/CVE-2025-1178) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1178](https://nvd.nist.gov/vuln/detail/CVE-2025-1178) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1179](https://nvd.nist.gov/vuln/detail/CVE-2025-1179) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1179](https://nvd.nist.gov/vuln/detail/CVE-2025-1179) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1179](https://nvd.nist.gov/vuln/detail/CVE-2025-1179) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1179](https://nvd.nist.gov/vuln/detail/CVE-2025-1179) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1179](https://nvd.nist.gov/vuln/detail/CVE-2025-1179) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1179](https://nvd.nist.gov/vuln/detail/CVE-2025-1179) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1179](https://nvd.nist.gov/vuln/detail/CVE-2025-1179) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1180](https://nvd.nist.gov/vuln/detail/CVE-2025-1180) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1180](https://nvd.nist.gov/vuln/detail/CVE-2025-1180) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1180](https://nvd.nist.gov/vuln/detail/CVE-2025-1180) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1180](https://nvd.nist.gov/vuln/detail/CVE-2025-1180) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1180](https://nvd.nist.gov/vuln/detail/CVE-2025-1180) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1180](https://nvd.nist.gov/vuln/detail/CVE-2025-1180) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1180](https://nvd.nist.gov/vuln/detail/CVE-2025-1180) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1181](https://nvd.nist.gov/vuln/detail/CVE-2025-1181) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1181](https://nvd.nist.gov/vuln/detail/CVE-2025-1181) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1181](https://nvd.nist.gov/vuln/detail/CVE-2025-1181) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1181](https://nvd.nist.gov/vuln/detail/CVE-2025-1181) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1181](https://nvd.nist.gov/vuln/detail/CVE-2025-1181) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1181](https://nvd.nist.gov/vuln/detail/CVE-2025-1181) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1181](https://nvd.nist.gov/vuln/detail/CVE-2025-1181) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1182](https://nvd.nist.gov/vuln/detail/CVE-2025-1182) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1182](https://nvd.nist.gov/vuln/detail/CVE-2025-1182) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1182](https://nvd.nist.gov/vuln/detail/CVE-2025-1182) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1182](https://nvd.nist.gov/vuln/detail/CVE-2025-1182) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1182](https://nvd.nist.gov/vuln/detail/CVE-2025-1182) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1182](https://nvd.nist.gov/vuln/detail/CVE-2025-1182) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1182](https://nvd.nist.gov/vuln/detail/CVE-2025-1182) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11839](https://nvd.nist.gov/vuln/detail/CVE-2025-11839) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11839](https://nvd.nist.gov/vuln/detail/CVE-2025-11839) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11839](https://nvd.nist.gov/vuln/detail/CVE-2025-11839) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11839](https://nvd.nist.gov/vuln/detail/CVE-2025-11839) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11839](https://nvd.nist.gov/vuln/detail/CVE-2025-11839) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11839](https://nvd.nist.gov/vuln/detail/CVE-2025-11839) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11839](https://nvd.nist.gov/vuln/detail/CVE-2025-11839) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11840](https://nvd.nist.gov/vuln/detail/CVE-2025-11840) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11840](https://nvd.nist.gov/vuln/detail/CVE-2025-11840) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11840](https://nvd.nist.gov/vuln/detail/CVE-2025-11840) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11840](https://nvd.nist.gov/vuln/detail/CVE-2025-11840) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11840](https://nvd.nist.gov/vuln/detail/CVE-2025-11840) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11840](https://nvd.nist.gov/vuln/detail/CVE-2025-11840) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11840](https://nvd.nist.gov/vuln/detail/CVE-2025-11840) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1352](https://nvd.nist.gov/vuln/detail/CVE-2025-1352) | Negligible | `libelf1` | `0.188-2.1` |
| [CVE-2025-1365](https://nvd.nist.gov/vuln/detail/CVE-2025-1365) | Negligible | `libelf1` | `0.188-2.1` |
| [CVE-2025-1371](https://nvd.nist.gov/vuln/detail/CVE-2025-1371) | Negligible | `libelf1` | `0.188-2.1` |
| [CVE-2025-1372](https://nvd.nist.gov/vuln/detail/CVE-2025-1372) | Negligible | `libelf1` | `0.188-2.1` |
| [CVE-2025-1376](https://nvd.nist.gov/vuln/detail/CVE-2025-1376) | Negligible | `libelf1` | `0.188-2.1` |
| [CVE-2025-1377](https://nvd.nist.gov/vuln/detail/CVE-2025-1377) | Negligible | `libelf1` | `0.188-2.1` |
| [CVE-2025-14017](https://nvd.nist.gov/vuln/detail/CVE-2025-14017) | Negligible | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-14017](https://nvd.nist.gov/vuln/detail/CVE-2025-14017) | Negligible | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-14017](https://nvd.nist.gov/vuln/detail/CVE-2025-14017) | Negligible | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-14017](https://nvd.nist.gov/vuln/detail/CVE-2025-14017) | Negligible | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-15079](https://nvd.nist.gov/vuln/detail/CVE-2025-15079) | Negligible | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-15079](https://nvd.nist.gov/vuln/detail/CVE-2025-15079) | Negligible | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-15079](https://nvd.nist.gov/vuln/detail/CVE-2025-15079) | Negligible | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-15079](https://nvd.nist.gov/vuln/detail/CVE-2025-15079) | Negligible | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-15224](https://nvd.nist.gov/vuln/detail/CVE-2025-15224) | Negligible | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-15224](https://nvd.nist.gov/vuln/detail/CVE-2025-15224) | Negligible | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-15224](https://nvd.nist.gov/vuln/detail/CVE-2025-15224) | Negligible | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-15224](https://nvd.nist.gov/vuln/detail/CVE-2025-15224) | Negligible | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-27587](https://nvd.nist.gov/vuln/detail/CVE-2025-27587) | Negligible | `libssl-dev` | `3.0.18-1~deb12u2` |
| [CVE-2025-27587](https://nvd.nist.gov/vuln/detail/CVE-2025-27587) | Negligible | `libssl3` | `3.0.18-1~deb12u2` |
| [CVE-2025-27587](https://nvd.nist.gov/vuln/detail/CVE-2025-27587) | Negligible | `openssl` | `3.0.18-1~deb12u2` |
| [CVE-2025-28162](https://nvd.nist.gov/vuln/detail/CVE-2025-28162) | Negligible | `libpng16-16` | `1.6.39-2+deb12u1` |
| [CVE-2025-28162](https://nvd.nist.gov/vuln/detail/CVE-2025-28162) | Negligible | `libpng-dev` | `1.6.39-2+deb12u1` |
| [CVE-2025-28164](https://nvd.nist.gov/vuln/detail/CVE-2025-28164) | Negligible | `libpng-dev` | `1.6.39-2+deb12u1` |
| [CVE-2025-28164](https://nvd.nist.gov/vuln/detail/CVE-2025-28164) | Negligible | `libpng16-16` | `1.6.39-2+deb12u1` |
| [CVE-2025-29070](https://nvd.nist.gov/vuln/detail/CVE-2025-29070) | Negligible | `liblcms2-2` | `2.14-2` |
| [CVE-2025-29070](https://nvd.nist.gov/vuln/detail/CVE-2025-29070) | Negligible | `liblcms2-dev` | `2.14-2` |
| [CVE-2025-29088](https://nvd.nist.gov/vuln/detail/CVE-2025-29088) | Negligible | `libsqlite3-dev` | `3.40.1-2+deb12u2` |
| [CVE-2025-29088](https://nvd.nist.gov/vuln/detail/CVE-2025-29088) | Negligible | `libsqlite3-0` | `3.40.1-2+deb12u2` |
| [CVE-2025-3198](https://nvd.nist.gov/vuln/detail/CVE-2025-3198) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-3198](https://nvd.nist.gov/vuln/detail/CVE-2025-3198) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-3198](https://nvd.nist.gov/vuln/detail/CVE-2025-3198) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-3198](https://nvd.nist.gov/vuln/detail/CVE-2025-3198) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-3198](https://nvd.nist.gov/vuln/detail/CVE-2025-3198) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-3198](https://nvd.nist.gov/vuln/detail/CVE-2025-3198) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-3198](https://nvd.nist.gov/vuln/detail/CVE-2025-3198) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-4516](https://nvd.nist.gov/vuln/detail/CVE-2025-4516) | Negligible | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-4516](https://nvd.nist.gov/vuln/detail/CVE-2025-4516) | Negligible | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-4516](https://nvd.nist.gov/vuln/detail/CVE-2025-4516) | Negligible | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-4516](https://nvd.nist.gov/vuln/detail/CVE-2025-4516) | Negligible | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-5244](https://nvd.nist.gov/vuln/detail/CVE-2025-5244) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-5244](https://nvd.nist.gov/vuln/detail/CVE-2025-5244) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-5244](https://nvd.nist.gov/vuln/detail/CVE-2025-5244) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-5244](https://nvd.nist.gov/vuln/detail/CVE-2025-5244) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-5244](https://nvd.nist.gov/vuln/detail/CVE-2025-5244) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-5244](https://nvd.nist.gov/vuln/detail/CVE-2025-5244) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-5244](https://nvd.nist.gov/vuln/detail/CVE-2025-5244) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-5245](https://nvd.nist.gov/vuln/detail/CVE-2025-5245) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-5245](https://nvd.nist.gov/vuln/detail/CVE-2025-5245) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-5245](https://nvd.nist.gov/vuln/detail/CVE-2025-5245) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-5245](https://nvd.nist.gov/vuln/detail/CVE-2025-5245) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-5245](https://nvd.nist.gov/vuln/detail/CVE-2025-5245) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-5245](https://nvd.nist.gov/vuln/detail/CVE-2025-5245) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-5245](https://nvd.nist.gov/vuln/detail/CVE-2025-5245) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-5278](https://nvd.nist.gov/vuln/detail/CVE-2025-5278) | Negligible | `coreutils` | `9.1-1` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-66861](https://nvd.nist.gov/vuln/detail/CVE-2025-66861) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-66861](https://nvd.nist.gov/vuln/detail/CVE-2025-66861) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-66861](https://nvd.nist.gov/vuln/detail/CVE-2025-66861) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-66861](https://nvd.nist.gov/vuln/detail/CVE-2025-66861) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-66861](https://nvd.nist.gov/vuln/detail/CVE-2025-66861) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-66861](https://nvd.nist.gov/vuln/detail/CVE-2025-66861) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-66861](https://nvd.nist.gov/vuln/detail/CVE-2025-66861) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-66862](https://nvd.nist.gov/vuln/detail/CVE-2025-66862) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-66862](https://nvd.nist.gov/vuln/detail/CVE-2025-66862) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-66862](https://nvd.nist.gov/vuln/detail/CVE-2025-66862) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-66862](https://nvd.nist.gov/vuln/detail/CVE-2025-66862) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-66862](https://nvd.nist.gov/vuln/detail/CVE-2025-66862) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-66862](https://nvd.nist.gov/vuln/detail/CVE-2025-66862) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-66862](https://nvd.nist.gov/vuln/detail/CVE-2025-66862) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-66863](https://nvd.nist.gov/vuln/detail/CVE-2025-66863) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-66863](https://nvd.nist.gov/vuln/detail/CVE-2025-66863) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-66863](https://nvd.nist.gov/vuln/detail/CVE-2025-66863) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-66863](https://nvd.nist.gov/vuln/detail/CVE-2025-66863) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-66863](https://nvd.nist.gov/vuln/detail/CVE-2025-66863) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-66863](https://nvd.nist.gov/vuln/detail/CVE-2025-66863) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-66863](https://nvd.nist.gov/vuln/detail/CVE-2025-66863) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-66864](https://nvd.nist.gov/vuln/detail/CVE-2025-66864) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-66864](https://nvd.nist.gov/vuln/detail/CVE-2025-66864) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-66864](https://nvd.nist.gov/vuln/detail/CVE-2025-66864) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-66864](https://nvd.nist.gov/vuln/detail/CVE-2025-66864) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-66864](https://nvd.nist.gov/vuln/detail/CVE-2025-66864) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-66864](https://nvd.nist.gov/vuln/detail/CVE-2025-66864) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-66864](https://nvd.nist.gov/vuln/detail/CVE-2025-66864) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-66865](https://nvd.nist.gov/vuln/detail/CVE-2025-66865) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-66865](https://nvd.nist.gov/vuln/detail/CVE-2025-66865) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-66865](https://nvd.nist.gov/vuln/detail/CVE-2025-66865) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-66865](https://nvd.nist.gov/vuln/detail/CVE-2025-66865) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-66865](https://nvd.nist.gov/vuln/detail/CVE-2025-66865) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-66865](https://nvd.nist.gov/vuln/detail/CVE-2025-66865) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-66865](https://nvd.nist.gov/vuln/detail/CVE-2025-66865) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-66866](https://nvd.nist.gov/vuln/detail/CVE-2025-66866) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-66866](https://nvd.nist.gov/vuln/detail/CVE-2025-66866) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-66866](https://nvd.nist.gov/vuln/detail/CVE-2025-66866) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-66866](https://nvd.nist.gov/vuln/detail/CVE-2025-66866) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-66866](https://nvd.nist.gov/vuln/detail/CVE-2025-66866) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-66866](https://nvd.nist.gov/vuln/detail/CVE-2025-66866) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-66866](https://nvd.nist.gov/vuln/detail/CVE-2025-66866) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-7545](https://nvd.nist.gov/vuln/detail/CVE-2025-7545) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-7545](https://nvd.nist.gov/vuln/detail/CVE-2025-7545) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-7545](https://nvd.nist.gov/vuln/detail/CVE-2025-7545) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-7545](https://nvd.nist.gov/vuln/detail/CVE-2025-7545) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-7545](https://nvd.nist.gov/vuln/detail/CVE-2025-7545) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-7545](https://nvd.nist.gov/vuln/detail/CVE-2025-7545) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-7545](https://nvd.nist.gov/vuln/detail/CVE-2025-7545) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-7546](https://nvd.nist.gov/vuln/detail/CVE-2025-7546) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-7546](https://nvd.nist.gov/vuln/detail/CVE-2025-7546) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-7546](https://nvd.nist.gov/vuln/detail/CVE-2025-7546) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-7546](https://nvd.nist.gov/vuln/detail/CVE-2025-7546) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-7546](https://nvd.nist.gov/vuln/detail/CVE-2025-7546) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-7546](https://nvd.nist.gov/vuln/detail/CVE-2025-7546) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-7546](https://nvd.nist.gov/vuln/detail/CVE-2025-7546) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-8176](https://nvd.nist.gov/vuln/detail/CVE-2025-8176) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8176](https://nvd.nist.gov/vuln/detail/CVE-2025-8176) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2025-8176](https://nvd.nist.gov/vuln/detail/CVE-2025-8176) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8177](https://nvd.nist.gov/vuln/detail/CVE-2025-8177) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8177](https://nvd.nist.gov/vuln/detail/CVE-2025-8177) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8177](https://nvd.nist.gov/vuln/detail/CVE-2025-8177) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2025-8224](https://nvd.nist.gov/vuln/detail/CVE-2025-8224) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-8224](https://nvd.nist.gov/vuln/detail/CVE-2025-8224) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-8224](https://nvd.nist.gov/vuln/detail/CVE-2025-8224) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-8224](https://nvd.nist.gov/vuln/detail/CVE-2025-8224) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-8224](https://nvd.nist.gov/vuln/detail/CVE-2025-8224) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-8224](https://nvd.nist.gov/vuln/detail/CVE-2025-8224) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-8224](https://nvd.nist.gov/vuln/detail/CVE-2025-8224) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-8225](https://nvd.nist.gov/vuln/detail/CVE-2025-8225) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-8225](https://nvd.nist.gov/vuln/detail/CVE-2025-8225) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-8225](https://nvd.nist.gov/vuln/detail/CVE-2025-8225) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-8225](https://nvd.nist.gov/vuln/detail/CVE-2025-8225) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-8225](https://nvd.nist.gov/vuln/detail/CVE-2025-8225) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-8225](https://nvd.nist.gov/vuln/detail/CVE-2025-8225) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-8225](https://nvd.nist.gov/vuln/detail/CVE-2025-8225) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-8534](https://nvd.nist.gov/vuln/detail/CVE-2025-8534) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8534](https://nvd.nist.gov/vuln/detail/CVE-2025-8534) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2025-8534](https://nvd.nist.gov/vuln/detail/CVE-2025-8534) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8732](https://nvd.nist.gov/vuln/detail/CVE-2025-8732) | Negligible | `libxml2-dev` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2025-8732](https://nvd.nist.gov/vuln/detail/CVE-2025-8732) | Negligible | `libxml2` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2025-8851](https://nvd.nist.gov/vuln/detail/CVE-2025-8851) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2025-8851](https://nvd.nist.gov/vuln/detail/CVE-2025-8851) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8851](https://nvd.nist.gov/vuln/detail/CVE-2025-8851) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8961](https://nvd.nist.gov/vuln/detail/CVE-2025-8961) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8961](https://nvd.nist.gov/vuln/detail/CVE-2025-8961) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8961](https://nvd.nist.gov/vuln/detail/CVE-2025-8961) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2025-9165](https://nvd.nist.gov/vuln/detail/CVE-2025-9165) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2025-9165](https://nvd.nist.gov/vuln/detail/CVE-2025-9165) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2025-9165](https://nvd.nist.gov/vuln/detail/CVE-2025-9165) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2026-1757](https://nvd.nist.gov/vuln/detail/CVE-2026-1757) | Negligible | `libxml2-dev` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-1757](https://nvd.nist.gov/vuln/detail/CVE-2026-1757) | Negligible | `libxml2` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-22185](https://nvd.nist.gov/vuln/detail/CVE-2026-22185) | Negligible | `libldap-2.5-0` | `2.5.13+dfsg-5` |
</details>

<details>
<summary><strong>📦 Installed Packages (576 total)</strong></summary>

| Package | Version |
|---------|---------|
| @isaacs/balanced-match | 4.0.1 |
| @isaacs/brace-expansion | 5.0.0 |
| @isaacs/fs-minipass | 4.0.1 |
| @isaacs/string-locale-compare | 1.1.0 |
| @npmcli/agent | 4.0.0 |
| @npmcli/arborist | 9.1.10 |
| @npmcli/config | 10.5.0 |
| @npmcli/fs | 5.0.0 |
| @npmcli/git | 7.0.1 |
| @npmcli/installed-package-contents | 4.0.0 |
| @npmcli/map-workspaces | 5.0.3 |
| @npmcli/metavuln-calculator | 9.0.3 |
| @npmcli/name-from-folder | 4.0.0 |
| @npmcli/node-gyp | 5.0.0 |
| @npmcli/package-json | 7.0.4 |
| @npmcli/promise-spawn | 9.0.1 |
| @npmcli/query | 5.0.0 |
| @npmcli/redact | 4.0.0 |
| @npmcli/run-script | 10.0.3 |
| @sigstore/bundle | 4.0.0 |
| @sigstore/core | 3.1.0 |
| @sigstore/protobuf-specs | 0.5.0 |
| @sigstore/sign | 4.1.0 |
| @sigstore/tuf | 4.0.1 |
| @sigstore/verify | 3.1.0 |
| @tufjs/canonical-json | 2.0.0 |
| @tufjs/models | 4.1.0 |
| abbrev | 4.0.0 |
| adduser | 3.134 |
| agent-base | 7.1.4 |
| ansi-regex | 5.0.1 |
| aproba | 2.1.0 |
| apt | 2.6.1 |
| archy | 1.0.0 |
| autoconf | 2.71-3 |
| automake | 1:1.16.5-1.3 |
| autotools-dev | 20220109.1 |
| base-files | 12.4+deb12u13 |
| base-passwd | 3.6.1 |
| bash | 5.2.15-2+b10 |
| bin-links | 6.0.0 |
| binary-extensions | 3.1.0 |
| binutils | 2.40-2 |
| binutils-aarch64-linux-gnu | 2.40-2 |
| binutils-common | 2.40-2 |
| bsdutils | 1:2.38.1-5+deb12u3 |
| bzip2 | 1.0.8-5+b1 |
| ca-certificates | 20230311+deb12u1 |
| cacache | 20.0.3 |
| chalk | 5.6.2 |
| chownr | 3.0.0 |
| ci-info | 4.3.1 |
| cidr-regex | 5.0.1 |
| cli-columns | 4.0.0 |
| cmd-shim | 8.0.0 |
| comerr-dev | 2.1-1.47.0-2+b2 |
| common-ancestor-path | 2.0.0 |
| corepack | 0.34.6 |
| coreutils | 9.1-1 |
| cpp | 4:12.2.0-3 |
| cpp-12 | 12.2.0-14+deb12u1 |
| cssesc | 3.0.0 |
| curl | 7.88.1-10+deb12u14 |
| dash | 0.5.12-2 |
| debconf | 1.5.82 |
| debian-archive-keyring | 2023.3+deb12u2 |
| debianutils | 5.7-0.5~deb12u1 |
| debug | 4.4.3 |
| default-libmysqlclient-dev | 1.1.0 |
| diff | 8.0.3 |
| diffutils | 1:3.8-4 |
| dirmngr | 2.2.40-1.1+deb12u2 |
| dpkg | 1.21.22 |
| dpkg-dev | 1.21.22 |
| e2fsprogs | 1.47.0-2+b2 |
| emoji-regex | 8.0.0 |
| encoding | 0.1.13 |
| env-paths | 2.2.1 |
| err-code | 2.0.3 |
| exponential-backoff | 3.1.3 |
| fastest-levenshtein | 1.0.16 |
| fdir | 6.5.0 |
| file | 1:5.44-3 |
| findutils | 4.9.0-4 |
| fontconfig | 2.14.1-4 |
| fontconfig-config | 2.14.1-4 |
| fonts-dejavu-core | 2.37-6 |
| fs-minipass | 3.0.3 |
| g++ | 4:12.2.0-3 |
| g++-12 | 12.2.0-14+deb12u1 |
| gcc | 4:12.2.0-3 |
| gcc-12 | 12.2.0-14+deb12u1 |
| gcc-12-base | 12.2.0-14+deb12u1 |
| gir1.2-freedesktop | 1.74.0-3 |
| gir1.2-gdkpixbuf-2.0 | 2.42.10+dfsg-1+deb12u3 |
| gir1.2-glib-2.0 | 1.74.0-3 |
| gir1.2-rsvg-2.0 | 2.54.7+dfsg-1~deb12u1 |
| git | 1:2.39.5-0+deb12u3 |
| git-man | 1:2.39.5-0+deb12u3 |
| glob | 13.0.0 |
| gnupg | 2.2.40-1.1+deb12u2 |
| gnupg-l10n | 2.2.40-1.1+deb12u2 |
| gnupg-utils | 2.2.40-1.1+deb12u2 |
| gpg | 2.2.40-1.1+deb12u2 |
| gpg-agent | 2.2.40-1.1+deb12u2 |
| gpg-wks-client | 2.2.40-1.1+deb12u2 |
| gpg-wks-server | 2.2.40-1.1+deb12u2 |
| gpgconf | 2.2.40-1.1+deb12u2 |
| gpgsm | 2.2.40-1.1+deb12u2 |
| gpgv | 2.2.40-1.1+deb12u2 |
| graceful-fs | 4.2.11 |
| grep | 3.8-5 |
| gzip | 1.12-1 |
| hicolor-icon-theme | 0.17-2 |
| hosted-git-info | 9.0.2 |
| hostname | 3.23+nmu1 |
| http-cache-semantics | 4.2.0 |
| http-proxy-agent | 7.0.2 |
| https-proxy-agent | 7.0.6 |
| iconv-lite | 0.6.3 |
| icu-devtools | 72.1-3+deb12u1 |
| ignore-walk | 8.0.0 |
| imagemagick | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| imagemagick-6-common | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| imagemagick-6.q16 | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| imurmurhash | 0.1.4 |
| ini | 6.0.0 |
| init-package-json | 8.2.4 |
| init-system-helpers | 1.65.2+deb12u1 |
| ip-address | 10.1.0 |
| ip-regex | 5.0.0 |
| is-cidr | 6.0.1 |
| is-fullwidth-code-point | 3.0.0 |
| isexe | 3.1.1 |
| json-parse-even-better-errors | 5.0.0 |
| json-stringify-nice | 1.1.4 |
| jsonparse | 1.3.1 |
| just-diff | 6.0.2 |
| just-diff-apply | 5.5.0 |
| krb5-multidev | 1.20.1-2+deb12u4 |
| libacl1 | 2.3.1-3 |
| libaom3 | 3.6.0-1+deb12u2 |
| libapr1 | 1.7.2-3+deb12u1 |
| libaprutil1 | 1.6.3-1 |
| libapt-pkg6.0 | 2.6.1 |
| libasan8 | 12.2.0-14+deb12u1 |
| libassuan0 | 2.5.5-5 |
| libatomic1 | 12.2.0-14+deb12u1 |
| libattr1 | 1:2.5.1-4 |
| libaudit-common | 1:3.0.9-1 |
| libaudit1 | 1:3.0.9-1 |
| libbinutils | 2.40-2 |
| libblkid-dev | 2.38.1-5+deb12u3 |
| libblkid1 | 2.38.1-5+deb12u3 |
| libbrotli-dev | 1.0.9-2+b6 |
| libbrotli1 | 1.0.9-2+b6 |
| libbsd0 | 0.11.7-2 |
| libbz2-1.0 | 1.0.8-5+b1 |
| libbz2-dev | 1.0.8-5+b1 |
| libc-bin | 2.36-9+deb12u13 |
| libc-dev-bin | 2.36-9+deb12u13 |
| libc6 | 2.36-9+deb12u13 |
| libc6-dev | 2.36-9+deb12u13 |
| libcairo-gobject2 | 1.16.0-7 |
| libcairo-script-interpreter2 | 1.16.0-7 |
| libcairo2 | 1.16.0-7 |
| libcairo2-dev | 1.16.0-7 |
| libcap-ng0 | 0.8.3-1+b3 |
| libcap2 | 1:2.66-4+deb12u2+b2 |
| libcbor0.8 | 0.8.0-2+b1 |
| libcc1-0 | 12.2.0-14+deb12u1 |
| libcom-err2 | 1.47.0-2+b2 |
| libcrypt-dev | 1:4.4.33-2 |
| libcrypt1 | 1:4.4.33-2 |
| libctf-nobfd0 | 2.40-2 |
| libctf0 | 2.40-2 |
| libcurl3-gnutls | 7.88.1-10+deb12u14 |
| libcurl4 | 7.88.1-10+deb12u14 |
| libcurl4-openssl-dev | 7.88.1-10+deb12u14 |
| libdatrie1 | 0.2.13-2+b1 |
| libdav1d6 | 1.0.0-2+deb12u1 |
| libdb-dev | 5.3.2 |
| libdb5.3 | 5.3.28+dfsg2-1 |
| libdb5.3-dev | 5.3.28+dfsg2-1 |
| libde265-0 | 1.0.11-1+deb12u2 |
| libdebconfclient0 | 0.270 |
| libdeflate-dev | 1.14-1 |
| libdeflate0 | 1.14-1 |
| libdjvulibre-dev | 3.5.28-2.2~deb12u1 |
| libdjvulibre-text | 3.5.28-2.2~deb12u1 |
| libdjvulibre21 | 3.5.28-2.2~deb12u1 |
| libdpkg-perl | 1.21.22 |
| libedit2 | 3.1-20221030-2 |
| libelf1 | 0.188-2.1 |
| liberror-perl | 0.17029-2 |
| libevent-2.1-7 | 2.1.12-stable-8 |
| libevent-core-2.1-7 | 2.1.12-stable-8 |
| libevent-dev | 2.1.12-stable-8 |
| libevent-extra-2.1-7 | 2.1.12-stable-8 |
| libevent-openssl-2.1-7 | 2.1.12-stable-8 |
| libevent-pthreads-2.1-7 | 2.1.12-stable-8 |
| libexif-dev | 0.6.24-1+b1 |
| libexif12 | 0.6.24-1+b1 |
| libexpat1 | 2.5.0-1+deb12u2 |
| libexpat1-dev | 2.5.0-1+deb12u2 |
| libext2fs2 | 1.47.0-2+b2 |
| libffi-dev | 3.4.4-1 |
| libffi8 | 3.4.4-1 |
| libfftw3-double3 | 3.3.10-1 |
| libfido2-1 | 1.12.0-2+b1 |
| libfontconfig-dev | 2.14.1-4 |
| libfontconfig1 | 2.14.1-4 |
| libfreetype-dev | 2.12.1+dfsg-5+deb12u4 |
| libfreetype6 | 2.12.1+dfsg-5+deb12u4 |
| libfreetype6-dev | 2.12.1+dfsg-5+deb12u4 |
| libfribidi0 | 1.0.8-2.1 |
| libgcc-12-dev | 12.2.0-14+deb12u1 |
| libgcc-s1 | 12.2.0-14+deb12u1 |
| libgcrypt20 | 1.10.1-3 |
| libgdbm-compat4 | 1.23-3 |
| libgdbm-dev | 1.23-3 |
| libgdbm6 | 1.23-3 |
| libgdk-pixbuf-2.0-0 | 2.42.10+dfsg-1+deb12u3 |
| libgdk-pixbuf-2.0-dev | 2.42.10+dfsg-1+deb12u3 |
| libgdk-pixbuf2.0-bin | 2.42.10+dfsg-1+deb12u3 |
| libgdk-pixbuf2.0-common | 2.42.10+dfsg-1+deb12u3 |
| libgirepository-1.0-1 | 1.74.0-3 |
| libglib2.0-0 | 2.74.6-2+deb12u8 |
| libglib2.0-bin | 2.74.6-2+deb12u8 |
| libglib2.0-data | 2.74.6-2+deb12u8 |
| libglib2.0-dev | 2.74.6-2+deb12u8 |
| libglib2.0-dev-bin | 2.74.6-2+deb12u8 |
| libgmp-dev | 2:6.2.1+dfsg1-1.1 |
| libgmp10 | 2:6.2.1+dfsg1-1.1 |
| libgmpxx4ldbl | 2:6.2.1+dfsg1-1.1 |
| libgnutls30 | 3.7.9-2+deb12u5 |
| libgomp1 | 12.2.0-14+deb12u1 |
| libgpg-error0 | 1.46-1 |
| libgprofng0 | 2.40-2 |
| libgraphite2-3 | 1.3.14-1 |
| libgssapi-krb5-2 | 1.20.1-2+deb12u4 |
| libgssrpc4 | 1.20.1-2+deb12u4 |
| libharfbuzz0b | 6.0.0+dfsg-3 |
| libheif1 | 1.15.1-1+deb12u1 |
| libhogweed6 | 3.8.1-2 |
| libhwasan0 | 12.2.0-14+deb12u1 |
| libice-dev | 2:1.0.10-1 |
| libice6 | 2:1.0.10-1 |
| libicu-dev | 72.1-3+deb12u1 |
| libicu72 | 72.1-3+deb12u1 |
| libidn2-0 | 2.3.3-1+b1 |
| libimath-3-1-29 | 3.1.6-1 |
| libimath-dev | 3.1.6-1 |
| libisl23 | 0.25-1.1 |
| libitm1 | 12.2.0-14+deb12u1 |
| libjansson4 | 2.14-2 |
| libjbig-dev | 2.1-6.1 |
| libjbig0 | 2.1-6.1 |
| libjpeg-dev | 1:2.1.5-2 |
| libjpeg62-turbo | 1:2.1.5-2 |
| libjpeg62-turbo-dev | 1:2.1.5-2 |
| libk5crypto3 | 1.20.1-2+deb12u4 |
| libkadm5clnt-mit12 | 1.20.1-2+deb12u4 |
| libkadm5srv-mit12 | 1.20.1-2+deb12u4 |
| libkdb5-10 | 1.20.1-2+deb12u4 |
| libkeyutils1 | 1.6.3-2 |
| libkrb5-3 | 1.20.1-2+deb12u4 |
| libkrb5-dev | 1.20.1-2+deb12u4 |
| libkrb5support0 | 1.20.1-2+deb12u4 |
| libksba8 | 1.6.3-2 |
| liblcms2-2 | 2.14-2 |
| liblcms2-dev | 2.14-2 |
| libldap-2.5-0 | 2.5.13+dfsg-5 |
| liblerc-dev | 4.0.0+ds-2 |
| liblerc4 | 4.0.0+ds-2 |
| liblqr-1-0 | 0.4.2-2.1 |
| liblqr-1-0-dev | 0.4.2-2.1 |
| liblsan0 | 12.2.0-14+deb12u1 |
| libltdl-dev | 2.4.7-7~deb12u1 |
| libltdl7 | 2.4.7-7~deb12u1 |
| liblz4-1 | 1.9.4-1 |
| liblzma-dev | 5.4.1-1 |
| liblzma5 | 5.4.1-1 |
| liblzo2-2 | 2.10-2 |
| libmagic-mgc | 1:5.44-3 |
| libmagic1 | 1:5.44-3 |
| libmagickcore-6-arch-config | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickcore-6-headers | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickcore-6.q16-6 | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickcore-6.q16-6-extra | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickcore-6.q16-dev | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickcore-dev | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickwand-6-headers | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickwand-6.q16-6 | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickwand-6.q16-dev | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickwand-dev | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmariadb-dev | 1:10.11.14-0+deb12u2 |
| libmariadb-dev-compat | 1:10.11.14-0+deb12u2 |
| libmariadb3 | 1:10.11.14-0+deb12u2 |
| libmaxminddb-dev | 1.7.1-1 |
| libmaxminddb0 | 1.7.1-1 |
| libmd0 | 1.0.4-2 |
| libmount-dev | 2.38.1-5+deb12u3 |
| libmount1 | 2.38.1-5+deb12u3 |
| libmpc3 | 1.3.1-1 |
| libmpfr6 | 4.2.0-1 |
| libncurses-dev | 6.4-4 |
| libncurses5-dev | 6.4-4 |
| libncurses6 | 6.4-4 |
| libncursesw5-dev | 6.4-4 |
| libncursesw6 | 6.4-4 |
| libnettle8 | 3.8.1-2 |
| libnghttp2-14 | 1.52.0-1+deb12u2 |
| libnpmaccess | 10.0.3 |
| libnpmdiff | 8.0.13 |
| libnpmexec | 10.1.12 |
| libnpmfund | 7.0.13 |
| libnpmorg | 8.0.1 |
| libnpmpack | 9.0.13 |
| libnpmpublish | 11.1.3 |
| libnpmsearch | 9.0.1 |
| libnpmteam | 8.0.2 |
| libnpmversion | 8.0.3 |
| libnpth0 | 1.6-3 |
| libnsl-dev | 1.3.0-2 |
| libnsl2 | 1.3.0-2 |
| libnuma1 | 2.0.16-1 |
| libopenexr-3-1-30 | 3.1.5-5 |
| libopenexr-dev | 3.1.5-5 |
| libopenjp2-7 | 2.5.0-2+deb12u2 |
| libopenjp2-7-dev | 2.5.0-2+deb12u2 |
| libp11-kit0 | 0.24.1-2 |
| libpam-modules | 1.5.2-6+deb12u2 |
| libpam-modules-bin | 1.5.2-6+deb12u2 |
| libpam-runtime | 1.5.2-6+deb12u2 |
| libpam0g | 1.5.2-6+deb12u2 |
| libpango-1.0-0 | 1.50.12+ds-1 |
| libpangocairo-1.0-0 | 1.50.12+ds-1 |
| libpangoft2-1.0-0 | 1.50.12+ds-1 |
| libpcre2-16-0 | 10.42-1 |
| libpcre2-32-0 | 10.42-1 |
| libpcre2-8-0 | 10.42-1 |
| libpcre2-dev | 10.42-1 |
| libpcre2-posix3 | 10.42-1 |
| libperl5.36 | 5.36.0-7+deb12u3 |
| libpixman-1-0 | 0.42.2-1 |
| libpixman-1-dev | 0.42.2-1 |
| libpkgconf3 | 1.8.1-1 |
| libpng-dev | 1.6.39-2+deb12u1 |
| libpng16-16 | 1.6.39-2+deb12u1 |
| libpq-dev | 15.15-0+deb12u1 |
| libpq5 | 15.15-0+deb12u1 |
| libproc2-0 | 2:4.0.2-3 |
| libpsl5 | 0.21.2-1 |
| libpthread-stubs0-dev | 0.4-1 |
| libpython3-stdlib | 3.11.2-1+b1 |
| libpython3.11-minimal | 3.11.2-6+deb12u6 |
| libpython3.11-stdlib | 3.11.2-6+deb12u6 |
| libreadline-dev | 8.2-1.3 |
| libreadline8 | 8.2-1.3 |
| librsvg2-2 | 2.54.7+dfsg-1~deb12u1 |
| librsvg2-common | 2.54.7+dfsg-1~deb12u1 |
| librsvg2-dev | 2.54.7+dfsg-1~deb12u1 |
| librtmp1 | 2.4+20151223.gitfa8646d.1-2+b2 |
| libsasl2-2 | 2.1.28+dfsg-10 |
| libsasl2-modules-db | 2.1.28+dfsg-10 |
| libseccomp2 | 2.5.4-1+deb12u1 |
| libselinux1 | 3.4-1+b6 |
| libselinux1-dev | 3.4-1+b6 |
| libsemanage-common | 3.4-1 |
| libsemanage2 | 3.4-1+b5 |
| libsepol-dev | 3.4-2.1 |
| libsepol2 | 3.4-2.1 |
| libserf-1-1 | 1.3.9-11 |
| libsm-dev | 2:1.2.3-1 |
| libsm6 | 2:1.2.3-1 |
| libsmartcols1 | 2.38.1-5+deb12u3 |
| libsqlite3-0 | 3.40.1-2+deb12u2 |
| libsqlite3-dev | 3.40.1-2+deb12u2 |
| libss2 | 1.47.0-2+b2 |
| libssh2-1 | 1.10.0-3+b1 |
| libssl-dev | 3.0.18-1~deb12u2 |
| libssl3 | 3.0.18-1~deb12u2 |
| libstdc++-12-dev | 12.2.0-14+deb12u1 |
| libstdc++6 | 12.2.0-14+deb12u1 |
| libsvn1 | 1.14.2-4+deb12u1 |
| libsystemd0 | 252.39-1~deb12u1 |
| libtasn1-6 | 4.19.0-2+deb12u1 |
| libthai-data | 0.1.29-1 |
| libthai0 | 0.1.29-1 |
| libtiff-dev | 4.5.0-6+deb12u3 |
| libtiff6 | 4.5.0-6+deb12u3 |
| libtiffxx6 | 4.5.0-6+deb12u3 |
| libtinfo6 | 6.4-4 |
| libtirpc-common | 1.3.3+ds-1 |
| libtirpc-dev | 1.3.3+ds-1 |
| libtirpc3 | 1.3.3+ds-1 |
| libtool | 2.4.7-7~deb12u1 |
| libtsan2 | 12.2.0-14+deb12u1 |
| libubsan1 | 12.2.0-14+deb12u1 |
| libudev1 | 252.39-1~deb12u1 |
| libunistring2 | 1.0-2 |
| libutf8proc2 | 2.8.0-1 |
| libuuid1 | 2.38.1-5+deb12u3 |
| libwebp-dev | 1.2.4-0.2+deb12u1 |
| libwebp7 | 1.2.4-0.2+deb12u1 |
| libwebpdemux2 | 1.2.4-0.2+deb12u1 |
| libwebpmux3 | 1.2.4-0.2+deb12u1 |
| libwmf-0.2-7 | 0.2.12-5.1 |
| libwmf-dev | 0.2.12-5.1 |
| libwmflite-0.2-7 | 0.2.12-5.1 |
| libx11-6 | 2:1.8.4-2+deb12u2 |
| libx11-data | 2:1.8.4-2+deb12u2 |
| libx11-dev | 2:1.8.4-2+deb12u2 |
| libx265-199 | 3.5-2+b1 |
| libxau-dev | 1:1.0.9-1 |
| libxau6 | 1:1.0.9-1 |
| libxcb-render0 | 1.15-1 |
| libxcb-render0-dev | 1.15-1 |
| libxcb-shm0 | 1.15-1 |
| libxcb-shm0-dev | 1.15-1 |
| libxcb1 | 1.15-1 |
| libxcb1-dev | 1.15-1 |
| libxdmcp-dev | 1:1.1.2-3 |
| libxdmcp6 | 1:1.1.2-3 |
| libxext-dev | 2:1.3.4-1+b1 |
| libxext6 | 2:1.3.4-1+b1 |
| libxml2 | 2.9.14+dfsg-1.3~deb12u5 |
| libxml2-dev | 2.9.14+dfsg-1.3~deb12u5 |
| libxrender-dev | 1:0.9.10-1.1 |
| libxrender1 | 1:0.9.10-1.1 |
| libxslt1-dev | 1.1.35-1+deb12u3 |
| libxslt1.1 | 1.1.35-1+deb12u3 |
| libxt-dev | 1:1.2.1-1.1 |
| libxt6 | 1:1.2.1-1.1 |
| libxxhash0 | 0.8.1-1 |
| libyaml-0-2 | 0.2.5-1 |
| libyaml-dev | 0.2.5-1 |
| libzstd-dev | 1.5.4+dfsg2-5 |
| libzstd1 | 1.5.4+dfsg2-5 |
| linux-libc-dev | 6.1.159-1 |
| login | 1:4.13+dfsg1-1+deb12u2 |
| logsave | 1.47.0-2+b2 |
| lru-cache | 11.2.4 |
| m4 | 1.4.19-3 |
| make | 4.3-4.1 |
| make-fetch-happen | 15.0.3 |
| mariadb-common | 1:10.11.14-0+deb12u2 |
| mawk | 1.3.4.20200120-3.1 |
| media-types | 10.0.0 |
| mercurial | 6.3.2 |
| mercurial | 6.3.2-1+deb12u1 |
| mercurial-common | 6.3.2-1+deb12u1 |
| minimatch | 10.1.1 |
| minipass | 3.3.6 |
| minipass | 7.1.2 |
| minipass-collect | 2.0.1 |
| minipass-fetch | 5.0.0 |
| minipass-flush | 1.0.5 |
| minipass-pipeline | 1.2.4 |
| minipass-sized | 1.0.3 |
| minizlib | 3.1.0 |
| mount | 2.38.1-5+deb12u3 |
| ms | 2.1.3 |
| mute-stream | 3.0.0 |
| mysql-common | 5.8+1.1.0 |
| ncurses-base | 6.4-4 |
| ncurses-bin | 6.4-4 |
| negotiator | 1.0.0 |
| netbase | 6.4 |
| node | 24.13.1 |
| node-gyp | 12.1.0 |
| nopt | 9.0.0 |
| npm | 11.8.0 |
| npm-audit-report | 7.0.0 |
| npm-bundled | 5.0.0 |
| npm-install-checks | 8.0.0 |
| npm-normalize-package-bin | 5.0.0 |
| npm-package-arg | 13.0.2 |
| npm-packlist | 10.0.3 |
| npm-pick-manifest | 11.0.3 |
| npm-profile | 12.0.1 |
| npm-registry-fetch | 19.1.1 |
| npm-user-validate | 4.0.0 |
| openssh-client | 1:9.2p1-2+deb12u7 |
| openssl | 3.0.18-1~deb12u2 |
| p-map | 7.0.4 |
| pacote | 21.0.4 |
| parse-conflict-json | 5.0.1 |
| passwd | 1:4.13+dfsg1-1+deb12u2 |
| patch | 2.7.6-7 |
| path-scurry | 2.0.1 |
| perl | 5.36.0-7+deb12u3 |
| perl-base | 5.36.0-7+deb12u3 |
| perl-modules-5.36 | 5.36.0-7+deb12u3 |
| picomatch | 4.0.3 |
| pinentry-curses | 1.2.1-1 |
| pkg-config | 1.8.1-1 |
| pkgconf | 1.8.1-1 |
| pkgconf-bin | 1.8.1-1 |
| postcss-selector-parser | 7.1.1 |
| proc-log | 6.1.0 |
| procps | 2:4.0.2-3 |
| proggy | 4.0.0 |
| promise-all-reject-late | 1.0.1 |
| promise-call-limit | 3.0.2 |
| promise-retry | 2.0.1 |
| promzard | 3.0.1 |
| python3 | 3.11.2-1+b1 |
| python3-distutils | 3.11.2-3 |
| python3-lib2to3 | 3.11.2-3 |
| python3-minimal | 3.11.2-1+b1 |
| python3.11 | 3.11.2-6+deb12u6 |
| python3.11-minimal | 3.11.2-6+deb12u6 |
| qrcode-terminal | 0.12.0 |
| read | 5.0.1 |
| read-cmd-shim | 6.0.0 |
| readline-common | 8.2-1.3 |
| retry | 0.12.0 |
| rpcsvc-proto | 1.4.3-1 |
| safer-buffer | 2.1.2 |
| sed | 4.9-1 |
| semver | 7.7.3 |
| sensible-utils | 0.0.17+nmu1 |
| shared-mime-info | 2.2-1 |
| signal-exit | 4.1.0 |
| sigstore | 4.1.0 |
| smart-buffer | 4.2.0 |
| socks | 2.8.7 |
| socks-proxy-agent | 8.0.5 |
| spdx-correct | 3.2.0 |
| spdx-exceptions | 2.5.0 |
| spdx-expression-parse | 3.0.1 |
| spdx-expression-parse | 4.0.0 |
| spdx-license-ids | 3.0.22 |
| sq | 0.27.0-2+b1 |
| ssri | 13.0.0 |
| string-width | 4.2.3 |
| strip-ansi | 6.0.1 |
| subversion | 1.14.2-4+deb12u1 |
| supports-color | 10.2.2 |
| sysvinit-utils | 3.06-4 |
| tar | 1.34+dfsg-1.2+deb12u1 |
| tar | 7.5.4 |
| text-table | 0.2.0 |
| tiny-relative-date | 2.0.2 |
| tinyglobby | 0.2.15 |
| treeverse | 3.0.0 |
| tuf-js | 4.1.0 |
| tzdata | 2025b-0+deb12u2 |
| ucf | 3.0043+nmu1+deb12u1 |
| unique-filename | 5.0.0 |
| unique-slug | 6.0.0 |
| unzip | 6.0-28 |
| usr-is-merged | 37~deb12u1 |
| util-deprecate | 1.0.2 |
| util-linux | 2.38.1-5+deb12u3 |
| util-linux-extra | 2.38.1-5+deb12u3 |
| uuid-dev | 2.38.1-5+deb12u3 |
| validate-npm-package-license | 3.0.4 |
| validate-npm-package-name | 7.0.2 |
| walk-up-path | 4.0.0 |
| wget | 1.21.3-1+deb12u1 |
| which | 6.0.0 |
| write-file-atomic | 7.0.0 |
| x11-common | 1:7.7+23 |
| x11proto-core-dev | 2022.1-1 |
| x11proto-dev | 2022.1-1 |
| xorg-sgml-doctools | 1:1.11-1.1 |
| xtrans-dev | 1.4.0-1 |
| xz-utils | 5.4.1-1 |
| yallist | 4.0.0 |
| yallist | 5.0.0 |
| yarn | 1.22.22 |
| zlib1g | 1:1.2.13.dfsg-1 |
| zlib1g-dev | 1:1.2.13.dfsg-1 |
</details>

<!-- END: dock-docs:config -->

## Version Comparison

<!-- BEGIN: dock-docs:comparison -->

### 🏷️ Supported Tags

| Tag | Size | Vulns | Efficiency | Architectures |
|-----|------|-------|------------|---------------|
| `node:22-alpine` | ![Size](https://img.shields.io/static/v1?label=Size&message=155.69+MB&color=blue) | ![Vulns](https://img.shields.io/static/v1?label=Security&message=14+Vulns+%280+Crit%29&color=orange) | 100.0% | `linux/amd64, linux/arm, linux/arm64, linux/s390x, unknown/unknown` |
| `node:24-alpine` | ![Size](https://img.shields.io/static/v1?label=Size&message=153.10+MB&color=blue) | ![Vulns](https://img.shields.io/static/v1?label=Security&message=5+Vulns+%280+Crit%29&color=orange) | 100.0% | `linux/amd64, linux/arm64, linux/s390x, unknown/unknown` |
| `node:22` | ![Size](https://img.shields.io/static/v1?label=Size&message=1103.11+MB&color=blue) | ![Vulns](https://img.shields.io/static/v1?label=Security&message=299+Vulns+%286+Crit%29&color=red) | 99.4% | `linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/s390x, unknown/unknown` |
| `node:24` | ![Size](https://img.shields.io/static/v1?label=Size&message=1098.28+MB&color=blue) | ![Vulns](https://img.shields.io/static/v1?label=Security&message=290+Vulns+%286+Crit%29&color=red) | 99.4% | `linux/amd64, linux/arm64, linux/ppc64le, linux/s390x, unknown/unknown` |

<details>
<summary><strong>🔍 Full Report: node:22-alpine</strong></summary>

## 🛡️ Security & Efficiency

**Base Image:** `Alpine Linux (linux/arm64)`
**Supported Architectures:** `linux/amd64, linux/arm, linux/arm64, linux/s390x, unknown/unknown`
**Efficiency Score:** 100.0%

### Vulnerabilities
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| 🟢 0 | 🟠 10 | 🟡 3 | 🔵 1 |

<details>
<summary><strong>👇 Expand Vulnerability Details (14 found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [GHSA-34x7-hfp2-rc4v](https://nvd.nist.gov/vuln/detail/GHSA-34x7-hfp2-rc4v) | High | `tar` | `6.2.1` |
| [GHSA-34x7-hfp2-rc4v](https://nvd.nist.gov/vuln/detail/GHSA-34x7-hfp2-rc4v) | High | `tar` | `7.4.3` |
| [GHSA-34x7-hfp2-rc4v](https://nvd.nist.gov/vuln/detail/GHSA-34x7-hfp2-rc4v) | High | `tar` | `7.4.3` |
| [GHSA-5j98-mcp5-4vw2](https://nvd.nist.gov/vuln/detail/GHSA-5j98-mcp5-4vw2) | High | `glob` | `10.4.5` |
| [GHSA-8qq5-rm4j-mr97](https://nvd.nist.gov/vuln/detail/GHSA-8qq5-rm4j-mr97) | High | `tar` | `6.2.1` |
| [GHSA-8qq5-rm4j-mr97](https://nvd.nist.gov/vuln/detail/GHSA-8qq5-rm4j-mr97) | High | `tar` | `7.4.3` |
| [GHSA-8qq5-rm4j-mr97](https://nvd.nist.gov/vuln/detail/GHSA-8qq5-rm4j-mr97) | High | `tar` | `7.4.3` |
| [GHSA-r6q2-hw4h-h46w](https://nvd.nist.gov/vuln/detail/GHSA-r6q2-hw4h-h46w) | High | `tar` | `6.2.1` |
| [GHSA-r6q2-hw4h-h46w](https://nvd.nist.gov/vuln/detail/GHSA-r6q2-hw4h-h46w) | High | `tar` | `7.4.3` |
| [GHSA-r6q2-hw4h-h46w](https://nvd.nist.gov/vuln/detail/GHSA-r6q2-hw4h-h46w) | High | `tar` | `7.4.3` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox` | `1.37.0-r30` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `ssl_client` | `1.37.0-r30` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox-binsh` | `1.37.0-r30` |
| [GHSA-73rr-hh4g-fpgx](https://nvd.nist.gov/vuln/detail/GHSA-73rr-hh4g-fpgx) | Low | `diff` | `5.2.0` |
</details>

<details>
<summary><strong>📦 Installed Packages (215 total)</strong></summary>

| Package | Version |
|---------|---------|
| @isaacs/cliui | 8.0.2 |
| @isaacs/fs-minipass | 4.0.1 |
| @isaacs/string-locale-compare | 1.1.0 |
| @npmcli/agent | 3.0.0 |
| @npmcli/arborist | 8.0.1 |
| @npmcli/config | 9.0.0 |
| @npmcli/fs | 4.0.0 |
| @npmcli/git | 6.0.3 |
| @npmcli/installed-package-contents | 3.0.0 |
| @npmcli/map-workspaces | 4.0.2 |
| @npmcli/metavuln-calculator | 8.0.1 |
| @npmcli/name-from-folder | 3.0.0 |
| @npmcli/node-gyp | 4.0.0 |
| @npmcli/package-json | 6.2.0 |
| @npmcli/promise-spawn | 8.0.2 |
| @npmcli/query | 4.0.1 |
| @npmcli/redact | 3.2.2 |
| @npmcli/run-script | 9.1.0 |
| @pkgjs/parseargs | 0.11.0 |
| @sigstore/bundle | 3.1.0 |
| @sigstore/core | 2.0.0 |
| @sigstore/protobuf-specs | 0.4.3 |
| @sigstore/sign | 3.1.0 |
| @sigstore/tuf | 3.1.1 |
| @sigstore/verify | 2.1.1 |
| @tufjs/canonical-json | 2.0.0 |
| @tufjs/models | 3.0.1 |
| abbrev | 3.0.1 |
| agent-base | 7.1.3 |
| alpine-baselayout | 3.7.1-r8 |
| alpine-baselayout-data | 3.7.1-r8 |
| alpine-keys | 2.6-r0 |
| alpine-release | 3.23.3-r0 |
| ansi-regex | 5.0.1 |
| ansi-regex | 6.1.0 |
| ansi-styles | 4.3.0 |
| ansi-styles | 6.2.1 |
| apk-tools | 3.0.3-r1 |
| aproba | 2.0.0 |
| archy | 1.0.0 |
| balanced-match | 1.0.2 |
| bin-links | 5.0.0 |
| binary-extensions | 2.3.0 |
| brace-expansion | 2.0.2 |
| busybox | 1.37.0-r30 |
| busybox-binsh | 1.37.0-r30 |
| ca-certificates-bundle | 20251003-r0 |
| cacache | 19.0.1 |
| chalk | 5.4.1 |
| chownr | 2.0.0 |
| chownr | 3.0.0 |
| ci-info | 4.2.0 |
| cidr-regex | 4.1.3 |
| cli-columns | 4.0.0 |
| cmd-shim | 7.0.0 |
| color-convert | 2.0.1 |
| color-name | 1.1.4 |
| common-ancestor-path | 1.0.1 |
| corepack | 0.34.0 |
| cross-spawn | 7.0.6 |
| cssesc | 3.0.0 |
| debug | 4.4.1 |
| diff | 5.2.0 |
| eastasianwidth | 0.2.0 |
| emoji-regex | 8.0.0 |
| emoji-regex | 9.2.2 |
| encoding | 0.1.13 |
| env-paths | 2.2.1 |
| err-code | 2.0.3 |
| exponential-backoff | 3.1.2 |
| fastest-levenshtein | 1.0.16 |
| fdir | 6.4.6 |
| foreground-child | 3.3.1 |
| fs-minipass | 2.1.0 |
| fs-minipass | 3.0.3 |
| glob | 10.4.5 |
| graceful-fs | 4.2.11 |
| hosted-git-info | 8.1.0 |
| http-cache-semantics | 4.2.0 |
| http-proxy-agent | 7.0.2 |
| https-proxy-agent | 7.0.6 |
| iconv-lite | 0.6.3 |
| ignore-walk | 7.0.0 |
| imurmurhash | 0.1.4 |
| ini | 5.0.0 |
| init-package-json | 7.0.2 |
| ip-address | 9.0.5 |
| ip-regex | 5.0.0 |
| is-cidr | 5.1.1 |
| is-fullwidth-code-point | 3.0.0 |
| isexe | 2.0.0 |
| isexe | 3.1.1 |
| jackspeak | 3.4.3 |
| jsbn | 1.1.0 |
| json-parse-even-better-errors | 4.0.0 |
| json-stringify-nice | 1.1.4 |
| jsonparse | 1.3.1 |
| just-diff | 6.0.2 |
| just-diff-apply | 5.5.0 |
| libapk | 3.0.3-r1 |
| libcrypto3 | 3.5.5-r0 |
| libgcc | 15.2.0-r2 |
| libnpmaccess | 9.0.0 |
| libnpmdiff | 7.0.1 |
| libnpmexec | 9.0.1 |
| libnpmfund | 6.0.1 |
| libnpmhook | 11.0.0 |
| libnpmorg | 7.0.0 |
| libnpmpack | 8.0.1 |
| libnpmpublish | 10.0.1 |
| libnpmsearch | 8.0.0 |
| libnpmteam | 7.0.0 |
| libnpmversion | 7.0.0 |
| libssl3 | 3.5.5-r0 |
| libstdc++ | 15.2.0-r2 |
| lru-cache | 10.4.3 |
| make-fetch-happen | 14.0.3 |
| minimatch | 9.0.5 |
| minipass | 3.3.6 |
| minipass | 5.0.0 |
| minipass | 7.1.2 |
| minipass-collect | 2.0.1 |
| minipass-fetch | 4.0.1 |
| minipass-flush | 1.0.5 |
| minipass-pipeline | 1.2.4 |
| minipass-sized | 1.0.3 |
| minizlib | 2.1.2 |
| minizlib | 3.0.2 |
| mkdirp | 1.0.4 |
| mkdirp | 3.0.1 |
| ms | 2.1.3 |
| musl | 1.2.5-r21 |
| musl-utils | 1.2.5-r21 |
| mute-stream | 2.0.0 |
| negotiator | 1.0.0 |
| node | 22.22.0 |
| node-gyp | 11.2.0 |
| nopt | 8.1.0 |
| normalize-package-data | 7.0.0 |
| npm | 10.9.4 |
| npm-audit-report | 6.0.0 |
| npm-bundled | 4.0.0 |
| npm-install-checks | 7.1.1 |
| npm-normalize-package-bin | 4.0.0 |
| npm-package-arg | 12.0.2 |
| npm-packlist | 9.0.0 |
| npm-pick-manifest | 10.0.0 |
| npm-profile | 11.0.1 |
| npm-registry-fetch | 18.0.2 |
| npm-user-validate | 3.0.0 |
| p-map | 7.0.3 |
| package-json-from-dist | 1.0.1 |
| pacote | 19.0.1 |
| pacote | 20.0.0 |
| parse-conflict-json | 4.0.0 |
| path-key | 3.1.1 |
| path-scurry | 1.11.1 |
| picomatch | 4.0.2 |
| postcss-selector-parser | 7.1.0 |
| proc-log | 5.0.0 |
| proggy | 3.0.0 |
| promise-all-reject-late | 1.0.1 |
| promise-call-limit | 3.0.2 |
| promise-retry | 2.0.1 |
| promzard | 2.0.0 |
| qrcode-terminal | 0.12.0 |
| read | 4.1.0 |
| read-cmd-shim | 5.0.0 |
| read-package-json-fast | 4.0.0 |
| retry | 0.12.0 |
| safer-buffer | 2.1.2 |
| scanelf | 1.3.8-r2 |
| semver | 7.7.2 |
| shebang-command | 2.0.0 |
| shebang-regex | 3.0.0 |
| signal-exit | 4.1.0 |
| sigstore | 3.1.0 |
| smart-buffer | 4.2.0 |
| socks | 2.8.5 |
| socks-proxy-agent | 8.0.5 |
| spdx-correct | 3.2.0 |
| spdx-exceptions | 2.5.0 |
| spdx-expression-parse | 3.0.1 |
| spdx-expression-parse | 4.0.0 |
| spdx-license-ids | 3.0.21 |
| sprintf-js | 1.1.3 |
| ssl_client | 1.37.0-r30 |
| ssri | 12.0.0 |
| string-width | 4.2.3 |
| string-width | 5.1.2 |
| strip-ansi | 6.0.1 |
| strip-ansi | 7.1.0 |
| supports-color | 9.4.0 |
| tar | 6.2.1 |
| tar | 7.4.3 |
| text-table | 0.2.0 |
| tiny-relative-date | 1.3.0 |
| tinyglobby | 0.2.14 |
| treeverse | 3.0.0 |
| tuf-js | 3.0.1 |
| unique-filename | 4.0.0 |
| unique-slug | 5.0.0 |
| util-deprecate | 1.0.2 |
| validate-npm-package-license | 3.0.4 |
| validate-npm-package-name | 6.0.1 |
| walk-up-path | 3.0.1 |
| which | 2.0.2 |
| which | 5.0.0 |
| wrap-ansi | 7.0.0 |
| wrap-ansi | 8.1.0 |
| write-file-atomic | 6.0.0 |
| yallist | 4.0.0 |
| yallist | 5.0.0 |
| yarn | 1.22.22 |
| zlib | 1.3.1-r2 |
</details>

</details>

<details>
<summary><strong>🔍 Full Report: node:24-alpine</strong></summary>

## 🛡️ Security & Efficiency

**Base Image:** `Alpine Linux (linux/arm64)`
**Supported Architectures:** `linux/amd64, linux/arm64, linux/s390x, unknown/unknown`
**Efficiency Score:** 100.0%

### Vulnerabilities
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| 🟢 0 | 🟠 2 | 🟡 3 | 🟢 0 |

<details>
<summary><strong>👇 Expand Vulnerability Details (5 found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [GHSA-34x7-hfp2-rc4v](https://nvd.nist.gov/vuln/detail/GHSA-34x7-hfp2-rc4v) | High | `tar` | `7.5.4` |
| [GHSA-7h2j-956f-4vf2](https://nvd.nist.gov/vuln/detail/GHSA-7h2j-956f-4vf2) | High | `@isaacs/brace-expansion` | `5.0.0` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox` | `1.37.0-r30` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox-binsh` | `1.37.0-r30` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `ssl_client` | `1.37.0-r30` |
</details>

<details>
<summary><strong>📦 Installed Packages (180 total)</strong></summary>

| Package | Version |
|---------|---------|
| @isaacs/balanced-match | 4.0.1 |
| @isaacs/brace-expansion | 5.0.0 |
| @isaacs/fs-minipass | 4.0.1 |
| @isaacs/string-locale-compare | 1.1.0 |
| @npmcli/agent | 4.0.0 |
| @npmcli/arborist | 9.1.10 |
| @npmcli/config | 10.5.0 |
| @npmcli/fs | 5.0.0 |
| @npmcli/git | 7.0.1 |
| @npmcli/installed-package-contents | 4.0.0 |
| @npmcli/map-workspaces | 5.0.3 |
| @npmcli/metavuln-calculator | 9.0.3 |
| @npmcli/name-from-folder | 4.0.0 |
| @npmcli/node-gyp | 5.0.0 |
| @npmcli/package-json | 7.0.4 |
| @npmcli/promise-spawn | 9.0.1 |
| @npmcli/query | 5.0.0 |
| @npmcli/redact | 4.0.0 |
| @npmcli/run-script | 10.0.3 |
| @sigstore/bundle | 4.0.0 |
| @sigstore/core | 3.1.0 |
| @sigstore/protobuf-specs | 0.5.0 |
| @sigstore/sign | 4.1.0 |
| @sigstore/tuf | 4.0.1 |
| @sigstore/verify | 3.1.0 |
| @tufjs/canonical-json | 2.0.0 |
| @tufjs/models | 4.1.0 |
| abbrev | 4.0.0 |
| agent-base | 7.1.4 |
| alpine-baselayout | 3.7.1-r8 |
| alpine-baselayout-data | 3.7.1-r8 |
| alpine-keys | 2.6-r0 |
| alpine-release | 3.23.3-r0 |
| ansi-regex | 5.0.1 |
| apk-tools | 3.0.3-r1 |
| aproba | 2.1.0 |
| archy | 1.0.0 |
| bin-links | 6.0.0 |
| binary-extensions | 3.1.0 |
| busybox | 1.37.0-r30 |
| busybox-binsh | 1.37.0-r30 |
| ca-certificates-bundle | 20251003-r0 |
| cacache | 20.0.3 |
| chalk | 5.6.2 |
| chownr | 3.0.0 |
| ci-info | 4.3.1 |
| cidr-regex | 5.0.1 |
| cli-columns | 4.0.0 |
| cmd-shim | 8.0.0 |
| common-ancestor-path | 2.0.0 |
| corepack | 0.34.6 |
| cssesc | 3.0.0 |
| debug | 4.4.3 |
| diff | 8.0.3 |
| emoji-regex | 8.0.0 |
| encoding | 0.1.13 |
| env-paths | 2.2.1 |
| err-code | 2.0.3 |
| exponential-backoff | 3.1.3 |
| fastest-levenshtein | 1.0.16 |
| fdir | 6.5.0 |
| fs-minipass | 3.0.3 |
| glob | 13.0.0 |
| graceful-fs | 4.2.11 |
| hosted-git-info | 9.0.2 |
| http-cache-semantics | 4.2.0 |
| http-proxy-agent | 7.0.2 |
| https-proxy-agent | 7.0.6 |
| iconv-lite | 0.6.3 |
| ignore-walk | 8.0.0 |
| imurmurhash | 0.1.4 |
| ini | 6.0.0 |
| init-package-json | 8.2.4 |
| ip-address | 10.1.0 |
| ip-regex | 5.0.0 |
| is-cidr | 6.0.1 |
| is-fullwidth-code-point | 3.0.0 |
| isexe | 3.1.1 |
| json-parse-even-better-errors | 5.0.0 |
| json-stringify-nice | 1.1.4 |
| jsonparse | 1.3.1 |
| just-diff | 6.0.2 |
| just-diff-apply | 5.5.0 |
| libapk | 3.0.3-r1 |
| libcrypto3 | 3.5.5-r0 |
| libgcc | 15.2.0-r2 |
| libnpmaccess | 10.0.3 |
| libnpmdiff | 8.0.13 |
| libnpmexec | 10.1.12 |
| libnpmfund | 7.0.13 |
| libnpmorg | 8.0.1 |
| libnpmpack | 9.0.13 |
| libnpmpublish | 11.1.3 |
| libnpmsearch | 9.0.1 |
| libnpmteam | 8.0.2 |
| libnpmversion | 8.0.3 |
| libssl3 | 3.5.5-r0 |
| libstdc++ | 15.2.0-r2 |
| lru-cache | 11.2.4 |
| make-fetch-happen | 15.0.3 |
| minimatch | 10.1.1 |
| minipass | 3.3.6 |
| minipass | 7.1.2 |
| minipass-collect | 2.0.1 |
| minipass-fetch | 5.0.0 |
| minipass-flush | 1.0.5 |
| minipass-pipeline | 1.2.4 |
| minipass-sized | 1.0.3 |
| minizlib | 3.1.0 |
| ms | 2.1.3 |
| musl | 1.2.5-r21 |
| musl-utils | 1.2.5-r21 |
| mute-stream | 3.0.0 |
| negotiator | 1.0.0 |
| node | 24.13.1 |
| node-gyp | 12.1.0 |
| nopt | 9.0.0 |
| npm | 11.8.0 |
| npm-audit-report | 7.0.0 |
| npm-bundled | 5.0.0 |
| npm-install-checks | 8.0.0 |
| npm-normalize-package-bin | 5.0.0 |
| npm-package-arg | 13.0.2 |
| npm-packlist | 10.0.3 |
| npm-pick-manifest | 11.0.3 |
| npm-profile | 12.0.1 |
| npm-registry-fetch | 19.1.1 |
| npm-user-validate | 4.0.0 |
| p-map | 7.0.4 |
| pacote | 21.0.4 |
| parse-conflict-json | 5.0.1 |
| path-scurry | 2.0.1 |
| picomatch | 4.0.3 |
| postcss-selector-parser | 7.1.1 |
| proc-log | 6.1.0 |
| proggy | 4.0.0 |
| promise-all-reject-late | 1.0.1 |
| promise-call-limit | 3.0.2 |
| promise-retry | 2.0.1 |
| promzard | 3.0.1 |
| qrcode-terminal | 0.12.0 |
| read | 5.0.1 |
| read-cmd-shim | 6.0.0 |
| retry | 0.12.0 |
| safer-buffer | 2.1.2 |
| scanelf | 1.3.8-r2 |
| semver | 7.7.3 |
| signal-exit | 4.1.0 |
| sigstore | 4.1.0 |
| smart-buffer | 4.2.0 |
| socks | 2.8.7 |
| socks-proxy-agent | 8.0.5 |
| spdx-correct | 3.2.0 |
| spdx-exceptions | 2.5.0 |
| spdx-expression-parse | 3.0.1 |
| spdx-expression-parse | 4.0.0 |
| spdx-license-ids | 3.0.22 |
| ssl_client | 1.37.0-r30 |
| ssri | 13.0.0 |
| string-width | 4.2.3 |
| strip-ansi | 6.0.1 |
| supports-color | 10.2.2 |
| tar | 7.5.4 |
| text-table | 0.2.0 |
| tiny-relative-date | 2.0.2 |
| tinyglobby | 0.2.15 |
| treeverse | 3.0.0 |
| tuf-js | 4.1.0 |
| unique-filename | 5.0.0 |
| unique-slug | 6.0.0 |
| util-deprecate | 1.0.2 |
| validate-npm-package-license | 3.0.4 |
| validate-npm-package-name | 7.0.2 |
| walk-up-path | 4.0.0 |
| which | 6.0.0 |
| write-file-atomic | 7.0.0 |
| yallist | 4.0.0 |
| yallist | 5.0.0 |
| yarn | 1.22.22 |
| zlib | 1.3.1-r2 |
</details>

</details>

<details>
<summary><strong>🔍 Full Report: node:22</strong></summary>

## 🛡️ Security & Efficiency

**Base Image:** `Debian GNU/Linux 12 (bookworm) (linux/arm64)`
**Supported Architectures:** `linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/s390x, unknown/unknown`
**Efficiency Score:** 99.4%

### Vulnerabilities
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| 🛑 6 | 🟠 74 | 🟡 161 | 🔵 58 |

<details>
<summary><strong>👇 Expand Vulnerability Details (1029 found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [CVE-2023-45853](https://nvd.nist.gov/vuln/detail/CVE-2023-45853) | Critical | `zlib1g-dev` | `1:1.2.13.dfsg-1` |
| [CVE-2023-5841](https://nvd.nist.gov/vuln/detail/CVE-2023-5841) | Critical | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2023-5841](https://nvd.nist.gov/vuln/detail/CVE-2023-5841) | Critical | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2023-6879](https://nvd.nist.gov/vuln/detail/CVE-2023-6879) | Critical | `libaom3` | `3.6.0-1+deb12u2` |
| [CVE-2025-7458](https://nvd.nist.gov/vuln/detail/CVE-2025-7458) | Critical | `libsqlite3-0` | `3.40.1-2+deb12u2` |
| [CVE-2025-7458](https://nvd.nist.gov/vuln/detail/CVE-2025-7458) | Critical | `libsqlite3-dev` | `3.40.1-2+deb12u2` |
| [CVE-2023-25193](https://nvd.nist.gov/vuln/detail/CVE-2023-25193) | High | `libharfbuzz0b` | `6.0.0+dfsg-3` |
| [CVE-2023-2953](https://nvd.nist.gov/vuln/detail/CVE-2023-2953) | High | `libldap-2.5-0` | `2.5.13+dfsg-5` |
| [CVE-2023-39616](https://nvd.nist.gov/vuln/detail/CVE-2023-39616) | High | `libaom3` | `3.6.0-1+deb12u2` |
| [CVE-2023-52355](https://nvd.nist.gov/vuln/detail/CVE-2023-52355) | High | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2023-52355](https://nvd.nist.gov/vuln/detail/CVE-2023-52355) | High | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2023-52355](https://nvd.nist.gov/vuln/detail/CVE-2023-52355) | High | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2025-12495](https://nvd.nist.gov/vuln/detail/CVE-2025-12495) | High | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2025-12495](https://nvd.nist.gov/vuln/detail/CVE-2025-12495) | High | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2025-12839](https://nvd.nist.gov/vuln/detail/CVE-2025-12839) | High | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2025-12839](https://nvd.nist.gov/vuln/detail/CVE-2025-12839) | High | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2025-12840](https://nvd.nist.gov/vuln/detail/CVE-2025-12840) | High | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2025-12840](https://nvd.nist.gov/vuln/detail/CVE-2025-12840) | High | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2025-13151](https://nvd.nist.gov/vuln/detail/CVE-2025-13151) | High | `libtasn1-6` | `4.19.0-2+deb12u1` |
| [CVE-2025-13699](https://nvd.nist.gov/vuln/detail/CVE-2025-13699) | High | `libmariadb-dev-compat` | `1:10.11.14-0+deb12u2` |
| [CVE-2025-13699](https://nvd.nist.gov/vuln/detail/CVE-2025-13699) | High | `libmariadb-dev` | `1:10.11.14-0+deb12u2` |
| [CVE-2025-13699](https://nvd.nist.gov/vuln/detail/CVE-2025-13699) | High | `libmariadb3` | `1:10.11.14-0+deb12u2` |
| [CVE-2025-13699](https://nvd.nist.gov/vuln/detail/CVE-2025-13699) | High | `mariadb-common` | `1:10.11.14-0+deb12u2` |
| [CVE-2025-13836](https://nvd.nist.gov/vuln/detail/CVE-2025-13836) | High | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-13836](https://nvd.nist.gov/vuln/detail/CVE-2025-13836) | High | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-13836](https://nvd.nist.gov/vuln/detail/CVE-2025-13836) | High | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-13836](https://nvd.nist.gov/vuln/detail/CVE-2025-13836) | High | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-15281](https://nvd.nist.gov/vuln/detail/CVE-2025-15281) | High | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2025-15281](https://nvd.nist.gov/vuln/detail/CVE-2025-15281) | High | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2025-15281](https://nvd.nist.gov/vuln/detail/CVE-2025-15281) | High | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2025-15281](https://nvd.nist.gov/vuln/detail/CVE-2025-15281) | High | `libc6` | `2.36-9+deb12u13` |
| [CVE-2025-46835](https://nvd.nist.gov/vuln/detail/CVE-2025-46835) | High | `git-man` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-46835](https://nvd.nist.gov/vuln/detail/CVE-2025-46835) | High | `git` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-48384](https://nvd.nist.gov/vuln/detail/CVE-2025-48384) | High | `git` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-48384](https://nvd.nist.gov/vuln/detail/CVE-2025-48384) | High | `git-man` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-48385](https://nvd.nist.gov/vuln/detail/CVE-2025-48385) | High | `git` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-48385](https://nvd.nist.gov/vuln/detail/CVE-2025-48385) | High | `git-man` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-59375](https://nvd.nist.gov/vuln/detail/CVE-2025-59375) | High | `libexpat1-dev` | `2.5.0-1+deb12u2` |
| [CVE-2025-59375](https://nvd.nist.gov/vuln/detail/CVE-2025-59375) | High | `libexpat1` | `2.5.0-1+deb12u2` |
| [CVE-2025-6297](https://nvd.nist.gov/vuln/detail/CVE-2025-6297) | High | `dpkg` | `1.21.22` |
| [CVE-2025-6297](https://nvd.nist.gov/vuln/detail/CVE-2025-6297) | High | `libdpkg-perl` | `1.21.22` |
| [CVE-2025-6297](https://nvd.nist.gov/vuln/detail/CVE-2025-6297) | High | `dpkg-dev` | `1.21.22` |
| [CVE-2025-64181](https://nvd.nist.gov/vuln/detail/CVE-2025-64181) | High | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2025-64181](https://nvd.nist.gov/vuln/detail/CVE-2025-64181) | High | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2025-7425](https://nvd.nist.gov/vuln/detail/CVE-2025-7425) | High | `libxslt1-dev` | `1.1.35-1+deb12u3` |
| [CVE-2025-7425](https://nvd.nist.gov/vuln/detail/CVE-2025-7425) | High | `libxslt1.1` | `1.1.35-1+deb12u3` |
| [CVE-2025-8194](https://nvd.nist.gov/vuln/detail/CVE-2025-8194) | High | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-8194](https://nvd.nist.gov/vuln/detail/CVE-2025-8194) | High | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-8194](https://nvd.nist.gov/vuln/detail/CVE-2025-8194) | High | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-8194](https://nvd.nist.gov/vuln/detail/CVE-2025-8194) | High | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2026-0861](https://nvd.nist.gov/vuln/detail/CVE-2026-0861) | High | `libc6` | `2.36-9+deb12u13` |
| [CVE-2026-0861](https://nvd.nist.gov/vuln/detail/CVE-2026-0861) | High | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2026-0861](https://nvd.nist.gov/vuln/detail/CVE-2026-0861) | High | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2026-0861](https://nvd.nist.gov/vuln/detail/CVE-2026-0861) | High | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2026-0915](https://nvd.nist.gov/vuln/detail/CVE-2026-0915) | High | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2026-0915](https://nvd.nist.gov/vuln/detail/CVE-2026-0915) | High | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2026-0915](https://nvd.nist.gov/vuln/detail/CVE-2026-0915) | High | `libc6` | `2.36-9+deb12u13` |
| [CVE-2026-0915](https://nvd.nist.gov/vuln/detail/CVE-2026-0915) | High | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2026-2004](https://nvd.nist.gov/vuln/detail/CVE-2026-2004) | High | `libpq-dev` | `15.15-0+deb12u1` |
| [CVE-2026-2004](https://nvd.nist.gov/vuln/detail/CVE-2026-2004) | High | `libpq5` | `15.15-0+deb12u1` |
| [CVE-2026-2005](https://nvd.nist.gov/vuln/detail/CVE-2026-2005) | High | `libpq-dev` | `15.15-0+deb12u1` |
| [CVE-2026-2005](https://nvd.nist.gov/vuln/detail/CVE-2026-2005) | High | `libpq5` | `15.15-0+deb12u1` |
| [CVE-2026-2006](https://nvd.nist.gov/vuln/detail/CVE-2026-2006) | High | `libpq-dev` | `15.15-0+deb12u1` |
| [CVE-2026-2006](https://nvd.nist.gov/vuln/detail/CVE-2026-2006) | High | `libpq5` | `15.15-0+deb12u1` |
| [CVE-2026-22695](https://nvd.nist.gov/vuln/detail/CVE-2026-22695) | High | `libpng16-16` | `1.6.39-2+deb12u1` |
| [CVE-2026-22695](https://nvd.nist.gov/vuln/detail/CVE-2026-22695) | High | `libpng-dev` | `1.6.39-2+deb12u1` |
| [CVE-2026-22801](https://nvd.nist.gov/vuln/detail/CVE-2026-22801) | High | `libpng16-16` | `1.6.39-2+deb12u1` |
| [CVE-2026-22801](https://nvd.nist.gov/vuln/detail/CVE-2026-22801) | High | `libpng-dev` | `1.6.39-2+deb12u1` |
| [CVE-2026-25646](https://nvd.nist.gov/vuln/detail/CVE-2026-25646) | High | `libpng-dev` | `1.6.39-2+deb12u1` |
| [CVE-2026-25646](https://nvd.nist.gov/vuln/detail/CVE-2026-25646) | High | `libpng16-16` | `1.6.39-2+deb12u1` |
| [GHSA-34x7-hfp2-rc4v](https://nvd.nist.gov/vuln/detail/GHSA-34x7-hfp2-rc4v) | High | `tar` | `7.4.3` |
| [GHSA-34x7-hfp2-rc4v](https://nvd.nist.gov/vuln/detail/GHSA-34x7-hfp2-rc4v) | High | `tar` | `7.4.3` |
| [GHSA-34x7-hfp2-rc4v](https://nvd.nist.gov/vuln/detail/GHSA-34x7-hfp2-rc4v) | High | `tar` | `6.2.1` |
| [GHSA-5j98-mcp5-4vw2](https://nvd.nist.gov/vuln/detail/GHSA-5j98-mcp5-4vw2) | High | `glob` | `10.4.5` |
| [GHSA-8qq5-rm4j-mr97](https://nvd.nist.gov/vuln/detail/GHSA-8qq5-rm4j-mr97) | High | `tar` | `7.4.3` |
| [GHSA-8qq5-rm4j-mr97](https://nvd.nist.gov/vuln/detail/GHSA-8qq5-rm4j-mr97) | High | `tar` | `7.4.3` |
| [GHSA-8qq5-rm4j-mr97](https://nvd.nist.gov/vuln/detail/GHSA-8qq5-rm4j-mr97) | High | `tar` | `6.2.1` |
| [GHSA-r6q2-hw4h-h46w](https://nvd.nist.gov/vuln/detail/GHSA-r6q2-hw4h-h46w) | High | `tar` | `7.4.3` |
| [GHSA-r6q2-hw4h-h46w](https://nvd.nist.gov/vuln/detail/GHSA-r6q2-hw4h-h46w) | High | `tar` | `6.2.1` |
| [GHSA-r6q2-hw4h-h46w](https://nvd.nist.gov/vuln/detail/GHSA-r6q2-hw4h-h46w) | High | `tar` | `7.4.3` |
| [CVE-2007-3996](https://nvd.nist.gov/vuln/detail/CVE-2007-3996) | Medium | `libwmflite-0.2-7` | `0.2.12-5.1` |
| [CVE-2007-3996](https://nvd.nist.gov/vuln/detail/CVE-2007-3996) | Medium | `libwmf-dev` | `0.2.12-5.1` |
| [CVE-2007-3996](https://nvd.nist.gov/vuln/detail/CVE-2007-3996) | Medium | `libwmf-0.2-7` | `0.2.12-5.1` |
| [CVE-2009-3546](https://nvd.nist.gov/vuln/detail/CVE-2009-3546) | Medium | `libwmflite-0.2-7` | `0.2.12-5.1` |
| [CVE-2009-3546](https://nvd.nist.gov/vuln/detail/CVE-2009-3546) | Medium | `libwmf-0.2-7` | `0.2.12-5.1` |
| [CVE-2009-3546](https://nvd.nist.gov/vuln/detail/CVE-2009-3546) | Medium | `libwmf-dev` | `0.2.12-5.1` |
| [CVE-2021-31879](https://nvd.nist.gov/vuln/detail/CVE-2021-31879) | Medium | `wget` | `1.21.3-1+deb12u1` |
| [CVE-2023-32570](https://nvd.nist.gov/vuln/detail/CVE-2023-32570) | Medium | `libdav1d6` | `1.0.0-2+deb12u1` |
| [CVE-2023-39327](https://nvd.nist.gov/vuln/detail/CVE-2023-39327) | Medium | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2023-39327](https://nvd.nist.gov/vuln/detail/CVE-2023-39327) | Medium | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2023-39328](https://nvd.nist.gov/vuln/detail/CVE-2023-39328) | Medium | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2023-39328](https://nvd.nist.gov/vuln/detail/CVE-2023-39328) | Medium | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2023-39329](https://nvd.nist.gov/vuln/detail/CVE-2023-39329) | Medium | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2023-39329](https://nvd.nist.gov/vuln/detail/CVE-2023-39329) | Medium | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `libncurses6` | `6.4-4` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `ncurses-bin` | `6.4-4` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `libncurses-dev` | `6.4-4` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `libncurses5-dev` | `6.4-4` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `libncursesw5-dev` | `6.4-4` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `libncursesw6` | `6.4-4` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `libtinfo6` | `6.4-4` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `ncurses-base` | `6.4-4` |
| [CVE-2023-6277](https://nvd.nist.gov/vuln/detail/CVE-2023-6277) | Medium | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2023-6277](https://nvd.nist.gov/vuln/detail/CVE-2023-6277) | Medium | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2023-6277](https://nvd.nist.gov/vuln/detail/CVE-2023-6277) | Medium | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2024-10041](https://nvd.nist.gov/vuln/detail/CVE-2024-10041) | Medium | `libpam-modules` | `1.5.2-6+deb12u2` |
| [CVE-2024-10041](https://nvd.nist.gov/vuln/detail/CVE-2024-10041) | Medium | `libpam0g` | `1.5.2-6+deb12u2` |
| [CVE-2024-10041](https://nvd.nist.gov/vuln/detail/CVE-2024-10041) | Medium | `libpam-runtime` | `1.5.2-6+deb12u2` |
| [CVE-2024-10041](https://nvd.nist.gov/vuln/detail/CVE-2024-10041) | Medium | `libpam-modules-bin` | `1.5.2-6+deb12u2` |
| [CVE-2024-10524](https://nvd.nist.gov/vuln/detail/CVE-2024-10524) | Medium | `wget` | `1.21.3-1+deb12u1` |
| [CVE-2024-38949](https://nvd.nist.gov/vuln/detail/CVE-2024-38949) | Medium | `libde265-0` | `1.0.11-1+deb12u2` |
| [CVE-2024-38950](https://nvd.nist.gov/vuln/detail/CVE-2024-38950) | Medium | `libde265-0` | `1.0.11-1+deb12u2` |
| [CVE-2025-10148](https://nvd.nist.gov/vuln/detail/CVE-2025-10148) | Medium | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-10148](https://nvd.nist.gov/vuln/detail/CVE-2025-10148) | Medium | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-10148](https://nvd.nist.gov/vuln/detail/CVE-2025-10148) | Medium | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-10148](https://nvd.nist.gov/vuln/detail/CVE-2025-10148) | Medium | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-10911](https://nvd.nist.gov/vuln/detail/CVE-2025-10911) | Medium | `libxslt1.1` | `1.1.35-1+deb12u3` |
| [CVE-2025-10911](https://nvd.nist.gov/vuln/detail/CVE-2025-10911) | Medium | `libxslt1-dev` | `1.1.35-1+deb12u3` |
| [CVE-2025-12084](https://nvd.nist.gov/vuln/detail/CVE-2025-12084) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-12084](https://nvd.nist.gov/vuln/detail/CVE-2025-12084) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-12084](https://nvd.nist.gov/vuln/detail/CVE-2025-12084) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-12084](https://nvd.nist.gov/vuln/detail/CVE-2025-12084) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-12781](https://nvd.nist.gov/vuln/detail/CVE-2025-12781) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-12781](https://nvd.nist.gov/vuln/detail/CVE-2025-12781) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-12781](https://nvd.nist.gov/vuln/detail/CVE-2025-12781) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-12781](https://nvd.nist.gov/vuln/detail/CVE-2025-12781) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-13837](https://nvd.nist.gov/vuln/detail/CVE-2025-13837) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-13837](https://nvd.nist.gov/vuln/detail/CVE-2025-13837) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-13837](https://nvd.nist.gov/vuln/detail/CVE-2025-13837) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-13837](https://nvd.nist.gov/vuln/detail/CVE-2025-13837) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libmount1` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `util-linux` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libmount-dev` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libsmartcols1` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `util-linux-extra` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `mount` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libuuid1` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libblkid1` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libblkid-dev` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `bsdutils` | `1:2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `uuid-dev` | `2.38.1-5+deb12u3` |
| [CVE-2025-14524](https://nvd.nist.gov/vuln/detail/CVE-2025-14524) | Medium | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-14524](https://nvd.nist.gov/vuln/detail/CVE-2025-14524) | Medium | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-14524](https://nvd.nist.gov/vuln/detail/CVE-2025-14524) | Medium | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-14524](https://nvd.nist.gov/vuln/detail/CVE-2025-14524) | Medium | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-14819](https://nvd.nist.gov/vuln/detail/CVE-2025-14819) | Medium | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-14819](https://nvd.nist.gov/vuln/detail/CVE-2025-14819) | Medium | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-14819](https://nvd.nist.gov/vuln/detail/CVE-2025-14819) | Medium | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-14819](https://nvd.nist.gov/vuln/detail/CVE-2025-14819) | Medium | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-14831](https://nvd.nist.gov/vuln/detail/CVE-2025-14831) | Medium | `libgnutls30` | `3.7.9-2+deb12u5` |
| [CVE-2025-15282](https://nvd.nist.gov/vuln/detail/CVE-2025-15282) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-15282](https://nvd.nist.gov/vuln/detail/CVE-2025-15282) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-15282](https://nvd.nist.gov/vuln/detail/CVE-2025-15282) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-15282](https://nvd.nist.gov/vuln/detail/CVE-2025-15282) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-15366](https://nvd.nist.gov/vuln/detail/CVE-2025-15366) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-15366](https://nvd.nist.gov/vuln/detail/CVE-2025-15366) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-15366](https://nvd.nist.gov/vuln/detail/CVE-2025-15366) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-15366](https://nvd.nist.gov/vuln/detail/CVE-2025-15366) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-15367](https://nvd.nist.gov/vuln/detail/CVE-2025-15367) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-15367](https://nvd.nist.gov/vuln/detail/CVE-2025-15367) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-15367](https://nvd.nist.gov/vuln/detail/CVE-2025-15367) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-15367](https://nvd.nist.gov/vuln/detail/CVE-2025-15367) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gpg-agent` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gpg` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gpgv` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gpg-wks-client` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gpgsm` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gpg-wks-server` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gpgconf` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gnupg` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gnupg-l10n` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gnupg-utils` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `dirmngr` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-48074](https://nvd.nist.gov/vuln/detail/CVE-2025-48074) | Medium | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2025-48074](https://nvd.nist.gov/vuln/detail/CVE-2025-48074) | Medium | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2025-6069](https://nvd.nist.gov/vuln/detail/CVE-2025-6069) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-6069](https://nvd.nist.gov/vuln/detail/CVE-2025-6069) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-6069](https://nvd.nist.gov/vuln/detail/CVE-2025-6069) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-6069](https://nvd.nist.gov/vuln/detail/CVE-2025-6069) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-6075](https://nvd.nist.gov/vuln/detail/CVE-2025-6075) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-6075](https://nvd.nist.gov/vuln/detail/CVE-2025-6075) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-6075](https://nvd.nist.gov/vuln/detail/CVE-2025-6075) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-6075](https://nvd.nist.gov/vuln/detail/CVE-2025-6075) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libncurses-dev` | `6.4-4` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libncurses5-dev` | `6.4-4` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libncursesw5-dev` | `6.4-4` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libncursesw6` | `6.4-4` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `ncurses-bin` | `6.4-4` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `ncurses-base` | `6.4-4` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libtinfo6` | `6.4-4` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libncurses6` | `6.4-4` |
| [CVE-2025-66382](https://nvd.nist.gov/vuln/detail/CVE-2025-66382) | Medium | `libexpat1-dev` | `2.5.0-1+deb12u2` |
| [CVE-2025-66382](https://nvd.nist.gov/vuln/detail/CVE-2025-66382) | Medium | `libexpat1` | `2.5.0-1+deb12u2` |
| [CVE-2025-68431](https://nvd.nist.gov/vuln/detail/CVE-2025-68431) | Medium | `libheif1` | `1.15.1-1+deb12u1` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gpg-agent` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gpgv` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gpg-wks-client` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gnupg-utils` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gnupg-l10n` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gpg` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gpgsm` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `dirmngr` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gnupg` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gpgconf` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gpg-wks-server` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-7709](https://nvd.nist.gov/vuln/detail/CVE-2025-7709) | Medium | `libsqlite3-dev` | `3.40.1-2+deb12u2` |
| [CVE-2025-7709](https://nvd.nist.gov/vuln/detail/CVE-2025-7709) | Medium | `libsqlite3-0` | `3.40.1-2+deb12u2` |
| [CVE-2025-8291](https://nvd.nist.gov/vuln/detail/CVE-2025-8291) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-8291](https://nvd.nist.gov/vuln/detail/CVE-2025-8291) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-8291](https://nvd.nist.gov/vuln/detail/CVE-2025-8291) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-8291](https://nvd.nist.gov/vuln/detail/CVE-2025-8291) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-9820](https://nvd.nist.gov/vuln/detail/CVE-2025-9820) | Medium | `libgnutls30` | `3.7.9-2+deb12u5` |
| [CVE-2026-0672](https://nvd.nist.gov/vuln/detail/CVE-2026-0672) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2026-0672](https://nvd.nist.gov/vuln/detail/CVE-2026-0672) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2026-0672](https://nvd.nist.gov/vuln/detail/CVE-2026-0672) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2026-0672](https://nvd.nist.gov/vuln/detail/CVE-2026-0672) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2026-0865](https://nvd.nist.gov/vuln/detail/CVE-2026-0865) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2026-0865](https://nvd.nist.gov/vuln/detail/CVE-2026-0865) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2026-0865](https://nvd.nist.gov/vuln/detail/CVE-2026-0865) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2026-0865](https://nvd.nist.gov/vuln/detail/CVE-2026-0865) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2026-0990](https://nvd.nist.gov/vuln/detail/CVE-2026-0990) | Medium | `libxml2` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-0990](https://nvd.nist.gov/vuln/detail/CVE-2026-0990) | Medium | `libxml2-dev` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-1299](https://nvd.nist.gov/vuln/detail/CVE-2026-1299) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2026-1299](https://nvd.nist.gov/vuln/detail/CVE-2026-1299) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2026-1299](https://nvd.nist.gov/vuln/detail/CVE-2026-1299) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2026-1299](https://nvd.nist.gov/vuln/detail/CVE-2026-1299) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2026-1484](https://nvd.nist.gov/vuln/detail/CVE-2026-1484) | Medium | `libglib2.0-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-1484](https://nvd.nist.gov/vuln/detail/CVE-2026-1484) | Medium | `libglib2.0-dev-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-1484](https://nvd.nist.gov/vuln/detail/CVE-2026-1484) | Medium | `libglib2.0-0` | `2.74.6-2+deb12u8` |
| [CVE-2026-1484](https://nvd.nist.gov/vuln/detail/CVE-2026-1484) | Medium | `libglib2.0-data` | `2.74.6-2+deb12u8` |
| [CVE-2026-1484](https://nvd.nist.gov/vuln/detail/CVE-2026-1484) | Medium | `libglib2.0-dev` | `2.74.6-2+deb12u8` |
| [CVE-2026-1489](https://nvd.nist.gov/vuln/detail/CVE-2026-1489) | Medium | `libglib2.0-0` | `2.74.6-2+deb12u8` |
| [CVE-2026-1489](https://nvd.nist.gov/vuln/detail/CVE-2026-1489) | Medium | `libglib2.0-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-1489](https://nvd.nist.gov/vuln/detail/CVE-2026-1489) | Medium | `libglib2.0-data` | `2.74.6-2+deb12u8` |
| [CVE-2026-1489](https://nvd.nist.gov/vuln/detail/CVE-2026-1489) | Medium | `libglib2.0-dev` | `2.74.6-2+deb12u8` |
| [CVE-2026-1489](https://nvd.nist.gov/vuln/detail/CVE-2026-1489) | Medium | `libglib2.0-dev-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-2003](https://nvd.nist.gov/vuln/detail/CVE-2026-2003) | Medium | `libpq5` | `15.15-0+deb12u1` |
| [CVE-2026-2003](https://nvd.nist.gov/vuln/detail/CVE-2026-2003) | Medium | `libpq-dev` | `15.15-0+deb12u1` |
| [CVE-2026-22693](https://nvd.nist.gov/vuln/detail/CVE-2026-22693) | Medium | `libharfbuzz0b` | `6.0.0+dfsg-3` |
| [CVE-2026-25210](https://nvd.nist.gov/vuln/detail/CVE-2026-25210) | Medium | `libexpat1` | `2.5.0-1+deb12u2` |
| [CVE-2026-25210](https://nvd.nist.gov/vuln/detail/CVE-2026-25210) | Medium | `libexpat1-dev` | `2.5.0-1+deb12u2` |
| [CVE-2007-3476](https://nvd.nist.gov/vuln/detail/CVE-2007-3476) | Low | `libwmflite-0.2-7` | `0.2.12-5.1` |
| [CVE-2007-3476](https://nvd.nist.gov/vuln/detail/CVE-2007-3476) | Low | `libwmf-dev` | `0.2.12-5.1` |
| [CVE-2007-3476](https://nvd.nist.gov/vuln/detail/CVE-2007-3476) | Low | `libwmf-0.2-7` | `0.2.12-5.1` |
| [CVE-2007-3477](https://nvd.nist.gov/vuln/detail/CVE-2007-3477) | Low | `libwmf-0.2-7` | `0.2.12-5.1` |
| [CVE-2007-3477](https://nvd.nist.gov/vuln/detail/CVE-2007-3477) | Low | `libwmflite-0.2-7` | `0.2.12-5.1` |
| [CVE-2007-3477](https://nvd.nist.gov/vuln/detail/CVE-2007-3477) | Low | `libwmf-dev` | `0.2.12-5.1` |
| [CVE-2016-2781](https://nvd.nist.gov/vuln/detail/CVE-2016-2781) | Low | `coreutils` | `9.1-1` |
| [CVE-2017-7475](https://nvd.nist.gov/vuln/detail/CVE-2017-7475) | Low | `libcairo2-dev` | `1.16.0-7` |
| [CVE-2017-7475](https://nvd.nist.gov/vuln/detail/CVE-2017-7475) | Low | `libcairo2` | `1.16.0-7` |
| [CVE-2017-7475](https://nvd.nist.gov/vuln/detail/CVE-2017-7475) | Low | `libcairo-script-interpreter2` | `1.16.0-7` |
| [CVE-2017-7475](https://nvd.nist.gov/vuln/detail/CVE-2017-7475) | Low | `libcairo-gobject2` | `1.16.0-7` |
| [CVE-2019-6461](https://nvd.nist.gov/vuln/detail/CVE-2019-6461) | Low | `libcairo-gobject2` | `1.16.0-7` |
| [CVE-2019-6461](https://nvd.nist.gov/vuln/detail/CVE-2019-6461) | Low | `libcairo-script-interpreter2` | `1.16.0-7` |
| [CVE-2019-6461](https://nvd.nist.gov/vuln/detail/CVE-2019-6461) | Low | `libcairo2-dev` | `1.16.0-7` |
| [CVE-2019-6461](https://nvd.nist.gov/vuln/detail/CVE-2019-6461) | Low | `libcairo2` | `1.16.0-7` |
| [CVE-2019-6462](https://nvd.nist.gov/vuln/detail/CVE-2019-6462) | Low | `libcairo2-dev` | `1.16.0-7` |
| [CVE-2019-6462](https://nvd.nist.gov/vuln/detail/CVE-2019-6462) | Low | `libcairo-gobject2` | `1.16.0-7` |
| [CVE-2019-6462](https://nvd.nist.gov/vuln/detail/CVE-2019-6462) | Low | `libcairo-script-interpreter2` | `1.16.0-7` |
| [CVE-2019-6462](https://nvd.nist.gov/vuln/detail/CVE-2019-6462) | Low | `libcairo2` | `1.16.0-7` |
| [CVE-2019-6988](https://nvd.nist.gov/vuln/detail/CVE-2019-6988) | Low | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2019-6988](https://nvd.nist.gov/vuln/detail/CVE-2019-6988) | Low | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2023-4016](https://nvd.nist.gov/vuln/detail/CVE-2023-4016) | Low | `libproc2-0` | `2:4.0.2-3` |
| [CVE-2023-4016](https://nvd.nist.gov/vuln/detail/CVE-2023-4016) | Low | `procps` | `2:4.0.2-3` |
| [CVE-2023-51792](https://nvd.nist.gov/vuln/detail/CVE-2023-51792) | Low | `libde265-0` | `1.0.11-1+deb12u2` |
| [CVE-2024-13978](https://nvd.nist.gov/vuln/detail/CVE-2024-13978) | Low | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2024-13978](https://nvd.nist.gov/vuln/detail/CVE-2024-13978) | Low | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2024-13978](https://nvd.nist.gov/vuln/detail/CVE-2024-13978) | Low | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2024-31047](https://nvd.nist.gov/vuln/detail/CVE-2024-31047) | Low | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2024-31047](https://nvd.nist.gov/vuln/detail/CVE-2024-31047) | Low | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2024-56433](https://nvd.nist.gov/vuln/detail/CVE-2024-56433) | Low | `passwd` | `1:4.13+dfsg1-1+deb12u2` |
| [CVE-2024-56433](https://nvd.nist.gov/vuln/detail/CVE-2024-56433) | Low | `login` | `1:4.13+dfsg1-1+deb12u2` |
| [CVE-2025-11731](https://nvd.nist.gov/vuln/detail/CVE-2025-11731) | Low | `libxslt1.1` | `1.1.35-1+deb12u3` |
| [CVE-2025-11731](https://nvd.nist.gov/vuln/detail/CVE-2025-11731) | Low | `libxslt1-dev` | `1.1.35-1+deb12u3` |
| [CVE-2025-27613](https://nvd.nist.gov/vuln/detail/CVE-2025-27613) | Low | `git-man` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-27613](https://nvd.nist.gov/vuln/detail/CVE-2025-27613) | Low | `git` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-50422](https://nvd.nist.gov/vuln/detail/CVE-2025-50422) | Low | `libcairo2` | `1.16.0-7` |
| [CVE-2025-50422](https://nvd.nist.gov/vuln/detail/CVE-2025-50422) | Low | `libcairo2-dev` | `1.16.0-7` |
| [CVE-2025-50422](https://nvd.nist.gov/vuln/detail/CVE-2025-50422) | Low | `libcairo-script-interpreter2` | `1.16.0-7` |
| [CVE-2025-50422](https://nvd.nist.gov/vuln/detail/CVE-2025-50422) | Low | `libcairo-gobject2` | `1.16.0-7` |
| [CVE-2025-61984](https://nvd.nist.gov/vuln/detail/CVE-2025-61984) | Low | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2025-61985](https://nvd.nist.gov/vuln/detail/CVE-2025-61985) | Low | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2026-0988](https://nvd.nist.gov/vuln/detail/CVE-2026-0988) | Low | `libglib2.0-dev-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-0988](https://nvd.nist.gov/vuln/detail/CVE-2026-0988) | Low | `libglib2.0-dev` | `2.74.6-2+deb12u8` |
| [CVE-2026-0988](https://nvd.nist.gov/vuln/detail/CVE-2026-0988) | Low | `libglib2.0-data` | `2.74.6-2+deb12u8` |
| [CVE-2026-0988](https://nvd.nist.gov/vuln/detail/CVE-2026-0988) | Low | `libglib2.0-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-0988](https://nvd.nist.gov/vuln/detail/CVE-2026-0988) | Low | `libglib2.0-0` | `2.74.6-2+deb12u8` |
| [CVE-2026-0989](https://nvd.nist.gov/vuln/detail/CVE-2026-0989) | Low | `libxml2` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-0989](https://nvd.nist.gov/vuln/detail/CVE-2026-0989) | Low | `libxml2-dev` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-0992](https://nvd.nist.gov/vuln/detail/CVE-2026-0992) | Low | `libxml2` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-0992](https://nvd.nist.gov/vuln/detail/CVE-2026-0992) | Low | `libxml2-dev` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-1485](https://nvd.nist.gov/vuln/detail/CVE-2026-1485) | Low | `libglib2.0-0` | `2.74.6-2+deb12u8` |
| [CVE-2026-1485](https://nvd.nist.gov/vuln/detail/CVE-2026-1485) | Low | `libglib2.0-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-1485](https://nvd.nist.gov/vuln/detail/CVE-2026-1485) | Low | `libglib2.0-data` | `2.74.6-2+deb12u8` |
| [CVE-2026-1485](https://nvd.nist.gov/vuln/detail/CVE-2026-1485) | Low | `libglib2.0-dev` | `2.74.6-2+deb12u8` |
| [CVE-2026-1485](https://nvd.nist.gov/vuln/detail/CVE-2026-1485) | Low | `libglib2.0-dev-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-24515](https://nvd.nist.gov/vuln/detail/CVE-2026-24515) | Low | `libexpat1` | `2.5.0-1+deb12u2` |
| [CVE-2026-24515](https://nvd.nist.gov/vuln/detail/CVE-2026-24515) | Low | `libexpat1-dev` | `2.5.0-1+deb12u2` |
| [GHSA-73rr-hh4g-fpgx](https://nvd.nist.gov/vuln/detail/GHSA-73rr-hh4g-fpgx) | Low | `diff` | `5.2.0` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-2541](https://nvd.nist.gov/vuln/detail/CVE-2005-2541) | Negligible | `tar` | `1.34+dfsg-1.2+deb12u1` |
| [CVE-2007-2243](https://nvd.nist.gov/vuln/detail/CVE-2007-2243) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2007-2768](https://nvd.nist.gov/vuln/detail/CVE-2007-2768) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2007-5686](https://nvd.nist.gov/vuln/detail/CVE-2007-5686) | Negligible | `login` | `1:4.13+dfsg1-1+deb12u2` |
| [CVE-2007-5686](https://nvd.nist.gov/vuln/detail/CVE-2007-5686) | Negligible | `passwd` | `1:4.13+dfsg1-1+deb12u2` |
| [CVE-2008-1687](https://nvd.nist.gov/vuln/detail/CVE-2008-1687) | Negligible | `m4` | `1.4.19-3` |
| [CVE-2008-1688](https://nvd.nist.gov/vuln/detail/CVE-2008-1688) | Negligible | `m4` | `1.4.19-3` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3234](https://nvd.nist.gov/vuln/detail/CVE-2008-3234) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2010-4651](https://nvd.nist.gov/vuln/detail/CVE-2010-4651) | Negligible | `patch` | `2.7.6-7` |
| [CVE-2010-4756](https://nvd.nist.gov/vuln/detail/CVE-2010-4756) | Negligible | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2010-4756](https://nvd.nist.gov/vuln/detail/CVE-2010-4756) | Negligible | `libc6` | `2.36-9+deb12u13` |
| [CVE-2010-4756](https://nvd.nist.gov/vuln/detail/CVE-2010-4756) | Negligible | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2010-4756](https://nvd.nist.gov/vuln/detail/CVE-2010-4756) | Negligible | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2011-3374](https://nvd.nist.gov/vuln/detail/CVE-2011-3374) | Negligible | `apt` | `2.6.1` |
| [CVE-2011-3374](https://nvd.nist.gov/vuln/detail/CVE-2011-3374) | Negligible | `libapt-pkg6.0` | `2.6.1` |
| [CVE-2011-3389](https://nvd.nist.gov/vuln/detail/CVE-2011-3389) | Negligible | `libgnutls30` | `3.7.9-2+deb12u5` |
| [CVE-2011-4116](https://nvd.nist.gov/vuln/detail/CVE-2011-4116) | Negligible | `perl-base` | `5.36.0-7+deb12u3` |
| [CVE-2011-4116](https://nvd.nist.gov/vuln/detail/CVE-2011-4116) | Negligible | `perl-modules-5.36` | `5.36.0-7+deb12u3` |
| [CVE-2011-4116](https://nvd.nist.gov/vuln/detail/CVE-2011-4116) | Negligible | `perl` | `5.36.0-7+deb12u3` |
| [CVE-2011-4116](https://nvd.nist.gov/vuln/detail/CVE-2011-4116) | Negligible | `libperl5.36` | `5.36.0-7+deb12u3` |
| [CVE-2012-0039](https://nvd.nist.gov/vuln/detail/CVE-2012-0039) | Negligible | `libglib2.0-0` | `2.74.6-2+deb12u8` |
| [CVE-2012-0039](https://nvd.nist.gov/vuln/detail/CVE-2012-0039) | Negligible | `libglib2.0-dev` | `2.74.6-2+deb12u8` |
| [CVE-2012-0039](https://nvd.nist.gov/vuln/detail/CVE-2012-0039) | Negligible | `libglib2.0-data` | `2.74.6-2+deb12u8` |
| [CVE-2012-0039](https://nvd.nist.gov/vuln/detail/CVE-2012-0039) | Negligible | `libglib2.0-bin` | `2.74.6-2+deb12u8` |
| [CVE-2012-0039](https://nvd.nist.gov/vuln/detail/CVE-2012-0039) | Negligible | `libglib2.0-dev-bin` | `2.74.6-2+deb12u8` |
| [CVE-2013-4392](https://nvd.nist.gov/vuln/detail/CVE-2013-4392) | Negligible | `libudev1` | `252.39-1~deb12u1` |
| [CVE-2013-4392](https://nvd.nist.gov/vuln/detail/CVE-2013-4392) | Negligible | `libsystemd0` | `252.39-1~deb12u1` |
| [CVE-2015-3276](https://nvd.nist.gov/vuln/detail/CVE-2015-3276) | Negligible | `libldap-2.5-0` | `2.5.13+dfsg-5` |
| [CVE-2015-9019](https://nvd.nist.gov/vuln/detail/CVE-2015-9019) | Negligible | `libxslt1-dev` | `1.1.35-1+deb12u3` |
| [CVE-2015-9019](https://nvd.nist.gov/vuln/detail/CVE-2015-9019) | Negligible | `libxslt1.1` | `1.1.35-1+deb12u3` |
| [CVE-2016-10505](https://nvd.nist.gov/vuln/detail/CVE-2016-10505) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-10505](https://nvd.nist.gov/vuln/detail/CVE-2016-10505) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2016-20012](https://nvd.nist.gov/vuln/detail/CVE-2016-20012) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-9113](https://nvd.nist.gov/vuln/detail/CVE-2016-9113) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2016-9113](https://nvd.nist.gov/vuln/detail/CVE-2016-9113) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-9114](https://nvd.nist.gov/vuln/detail/CVE-2016-9114) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2016-9114](https://nvd.nist.gov/vuln/detail/CVE-2016-9114) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-9115](https://nvd.nist.gov/vuln/detail/CVE-2016-9115) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-9115](https://nvd.nist.gov/vuln/detail/CVE-2016-9115) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2016-9116](https://nvd.nist.gov/vuln/detail/CVE-2016-9116) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2016-9116](https://nvd.nist.gov/vuln/detail/CVE-2016-9116) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-9117](https://nvd.nist.gov/vuln/detail/CVE-2016-9117) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2016-9117](https://nvd.nist.gov/vuln/detail/CVE-2016-9117) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-9580](https://nvd.nist.gov/vuln/detail/CVE-2016-9580) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2016-9580](https://nvd.nist.gov/vuln/detail/CVE-2016-9580) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-9581](https://nvd.nist.gov/vuln/detail/CVE-2016-9581) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-9581](https://nvd.nist.gov/vuln/detail/CVE-2016-9581) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-13716](https://nvd.nist.gov/vuln/detail/CVE-2017-13716) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2017-13716](https://nvd.nist.gov/vuln/detail/CVE-2017-13716) | Negligible | `binutils` | `2.40-2` |
| [CVE-2017-13716](https://nvd.nist.gov/vuln/detail/CVE-2017-13716) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2017-13716](https://nvd.nist.gov/vuln/detail/CVE-2017-13716) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2017-13716](https://nvd.nist.gov/vuln/detail/CVE-2017-13716) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2017-13716](https://nvd.nist.gov/vuln/detail/CVE-2017-13716) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2017-13716](https://nvd.nist.gov/vuln/detail/CVE-2017-13716) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2017-14159](https://nvd.nist.gov/vuln/detail/CVE-2017-14159) | Negligible | `libldap-2.5-0` | `2.5.13+dfsg-5` |
| [CVE-2017-14988](https://nvd.nist.gov/vuln/detail/CVE-2017-14988) | Negligible | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2017-14988](https://nvd.nist.gov/vuln/detail/CVE-2017-14988) | Negligible | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2017-16232](https://nvd.nist.gov/vuln/detail/CVE-2017-16232) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2017-16232](https://nvd.nist.gov/vuln/detail/CVE-2017-16232) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2017-16232](https://nvd.nist.gov/vuln/detail/CVE-2017-16232) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2017-17740](https://nvd.nist.gov/vuln/detail/CVE-2017-17740) | Negligible | `libldap-2.5-0` | `2.5.13+dfsg-5` |
| [CVE-2017-18018](https://nvd.nist.gov/vuln/detail/CVE-2017-18018) | Negligible | `coreutils` | `9.1-1` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-9937](https://nvd.nist.gov/vuln/detail/CVE-2017-9937) | Negligible | `libjbig0` | `2.1-6.1` |
| [CVE-2017-9937](https://nvd.nist.gov/vuln/detail/CVE-2017-9937) | Negligible | `libjbig-dev` | `2.1-6.1` |
| [CVE-2018-1000021](https://nvd.nist.gov/vuln/detail/CVE-2018-1000021) | Negligible | `git-man` | `1:2.39.5-0+deb12u3` |
| [CVE-2018-1000021](https://nvd.nist.gov/vuln/detail/CVE-2018-1000021) | Negligible | `git` | `1:2.39.5-0+deb12u3` |
| [CVE-2018-10126](https://nvd.nist.gov/vuln/detail/CVE-2018-10126) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2018-10126](https://nvd.nist.gov/vuln/detail/CVE-2018-10126) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2018-10126](https://nvd.nist.gov/vuln/detail/CVE-2018-10126) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15919](https://nvd.nist.gov/vuln/detail/CVE-2018-15919) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2018-16376](https://nvd.nist.gov/vuln/detail/CVE-2018-16376) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2018-16376](https://nvd.nist.gov/vuln/detail/CVE-2018-16376) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2018-18064](https://nvd.nist.gov/vuln/detail/CVE-2018-18064) | Negligible | `libcairo2-dev` | `1.16.0-7` |
| [CVE-2018-18064](https://nvd.nist.gov/vuln/detail/CVE-2018-18064) | Negligible | `libcairo-gobject2` | `1.16.0-7` |
| [CVE-2018-18064](https://nvd.nist.gov/vuln/detail/CVE-2018-18064) | Negligible | `libcairo-script-interpreter2` | `1.16.0-7` |
| [CVE-2018-18064](https://nvd.nist.gov/vuln/detail/CVE-2018-18064) | Negligible | `libcairo2` | `1.16.0-7` |
| [CVE-2018-20673](https://nvd.nist.gov/vuln/detail/CVE-2018-20673) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2018-20673](https://nvd.nist.gov/vuln/detail/CVE-2018-20673) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2018-20673](https://nvd.nist.gov/vuln/detail/CVE-2018-20673) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2018-20673](https://nvd.nist.gov/vuln/detail/CVE-2018-20673) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2018-20673](https://nvd.nist.gov/vuln/detail/CVE-2018-20673) | Negligible | `binutils` | `2.40-2` |
| [CVE-2018-20673](https://nvd.nist.gov/vuln/detail/CVE-2018-20673) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2018-20673](https://nvd.nist.gov/vuln/detail/CVE-2018-20673) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2018-20712](https://nvd.nist.gov/vuln/detail/CVE-2018-20712) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2018-20712](https://nvd.nist.gov/vuln/detail/CVE-2018-20712) | Negligible | `binutils` | `2.40-2` |
| [CVE-2018-20712](https://nvd.nist.gov/vuln/detail/CVE-2018-20712) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2018-20712](https://nvd.nist.gov/vuln/detail/CVE-2018-20712) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2018-20712](https://nvd.nist.gov/vuln/detail/CVE-2018-20712) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2018-20712](https://nvd.nist.gov/vuln/detail/CVE-2018-20712) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2018-20712](https://nvd.nist.gov/vuln/detail/CVE-2018-20712) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2018-20796](https://nvd.nist.gov/vuln/detail/CVE-2018-20796) | Negligible | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2018-20796](https://nvd.nist.gov/vuln/detail/CVE-2018-20796) | Negligible | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2018-20796](https://nvd.nist.gov/vuln/detail/CVE-2018-20796) | Negligible | `libc6` | `2.36-9+deb12u13` |
| [CVE-2018-20796](https://nvd.nist.gov/vuln/detail/CVE-2018-20796) | Negligible | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libkrb5-dev` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `krb5-multidev` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libkrb5-3` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libgssrpc4` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libkdb5-10` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libkadm5srv-mit12` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libkadm5clnt-mit12` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libk5crypto3` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libgssapi-krb5-2` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libkrb5support0` | `1.20.1-2+deb12u4` |
| [CVE-2018-6829](https://nvd.nist.gov/vuln/detail/CVE-2018-6829) | Negligible | `libgcrypt20` | `1.10.1-3` |
| [CVE-2018-6951](https://nvd.nist.gov/vuln/detail/CVE-2018-6951) | Negligible | `patch` | `2.7.6-7` |
| [CVE-2018-6952](https://nvd.nist.gov/vuln/detail/CVE-2018-6952) | Negligible | `patch` | `2.7.6-7` |
| [CVE-2018-9996](https://nvd.nist.gov/vuln/detail/CVE-2018-9996) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2018-9996](https://nvd.nist.gov/vuln/detail/CVE-2018-9996) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2018-9996](https://nvd.nist.gov/vuln/detail/CVE-2018-9996) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2018-9996](https://nvd.nist.gov/vuln/detail/CVE-2018-9996) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2018-9996](https://nvd.nist.gov/vuln/detail/CVE-2018-9996) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2018-9996](https://nvd.nist.gov/vuln/detail/CVE-2018-9996) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2018-9996](https://nvd.nist.gov/vuln/detail/CVE-2018-9996) | Negligible | `binutils` | `2.40-2` |
| [CVE-2019-1010022](https://nvd.nist.gov/vuln/detail/CVE-2019-1010022) | Negligible | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2019-1010022](https://nvd.nist.gov/vuln/detail/CVE-2019-1010022) | Negligible | `libc6` | `2.36-9+deb12u13` |
| [CVE-2019-1010022](https://nvd.nist.gov/vuln/detail/CVE-2019-1010022) | Negligible | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2019-1010022](https://nvd.nist.gov/vuln/detail/CVE-2019-1010022) | Negligible | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2019-1010023](https://nvd.nist.gov/vuln/detail/CVE-2019-1010023) | Negligible | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2019-1010023](https://nvd.nist.gov/vuln/detail/CVE-2019-1010023) | Negligible | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2019-1010023](https://nvd.nist.gov/vuln/detail/CVE-2019-1010023) | Negligible | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2019-1010023](https://nvd.nist.gov/vuln/detail/CVE-2019-1010023) | Negligible | `libc6` | `2.36-9+deb12u13` |
| [CVE-2019-1010024](https://nvd.nist.gov/vuln/detail/CVE-2019-1010024) | Negligible | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2019-1010024](https://nvd.nist.gov/vuln/detail/CVE-2019-1010024) | Negligible | `libc6` | `2.36-9+deb12u13` |
| [CVE-2019-1010024](https://nvd.nist.gov/vuln/detail/CVE-2019-1010024) | Negligible | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2019-1010024](https://nvd.nist.gov/vuln/detail/CVE-2019-1010024) | Negligible | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2019-1010025](https://nvd.nist.gov/vuln/detail/CVE-2019-1010025) | Negligible | `libc6` | `2.36-9+deb12u13` |
| [CVE-2019-1010025](https://nvd.nist.gov/vuln/detail/CVE-2019-1010025) | Negligible | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2019-1010025](https://nvd.nist.gov/vuln/detail/CVE-2019-1010025) | Negligible | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2019-1010025](https://nvd.nist.gov/vuln/detail/CVE-2019-1010025) | Negligible | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2019-6110](https://nvd.nist.gov/vuln/detail/CVE-2019-6110) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2019-9192](https://nvd.nist.gov/vuln/detail/CVE-2019-9192) | Negligible | `libc6` | `2.36-9+deb12u13` |
| [CVE-2019-9192](https://nvd.nist.gov/vuln/detail/CVE-2019-9192) | Negligible | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2019-9192](https://nvd.nist.gov/vuln/detail/CVE-2019-9192) | Negligible | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2019-9192](https://nvd.nist.gov/vuln/detail/CVE-2019-9192) | Negligible | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2020-14145](https://nvd.nist.gov/vuln/detail/CVE-2020-14145) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2020-15719](https://nvd.nist.gov/vuln/detail/CVE-2020-15719) | Negligible | `libldap-2.5-0` | `2.5.13+dfsg-5` |
| [CVE-2020-15778](https://nvd.nist.gov/vuln/detail/CVE-2020-15778) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2020-36325](https://nvd.nist.gov/vuln/detail/CVE-2020-36325) | Negligible | `libjansson4` | `2.14-2` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-32256](https://nvd.nist.gov/vuln/detail/CVE-2021-32256) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2021-32256](https://nvd.nist.gov/vuln/detail/CVE-2021-32256) | Negligible | `binutils` | `2.40-2` |
| [CVE-2021-32256](https://nvd.nist.gov/vuln/detail/CVE-2021-32256) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2021-32256](https://nvd.nist.gov/vuln/detail/CVE-2021-32256) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2021-32256](https://nvd.nist.gov/vuln/detail/CVE-2021-32256) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2021-32256](https://nvd.nist.gov/vuln/detail/CVE-2021-32256) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2021-32256](https://nvd.nist.gov/vuln/detail/CVE-2021-32256) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2021-4214](https://nvd.nist.gov/vuln/detail/CVE-2021-4214) | Negligible | `libpng-dev` | `1.6.39-2+deb12u1` |
| [CVE-2021-4214](https://nvd.nist.gov/vuln/detail/CVE-2021-4214) | Negligible | `libpng16-16` | `1.6.39-2+deb12u1` |
| [CVE-2021-4217](https://nvd.nist.gov/vuln/detail/CVE-2021-4217) | Negligible | `unzip` | `6.0-28` |
| [CVE-2021-45261](https://nvd.nist.gov/vuln/detail/CVE-2021-45261) | Negligible | `patch` | `2.7.6-7` |
| [CVE-2021-45346](https://nvd.nist.gov/vuln/detail/CVE-2021-45346) | Negligible | `libsqlite3-dev` | `3.40.1-2+deb12u2` |
| [CVE-2021-45346](https://nvd.nist.gov/vuln/detail/CVE-2021-45346) | Negligible | `libsqlite3-0` | `3.40.1-2+deb12u2` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libuuid1` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libmount1` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `uuid-dev` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `util-linux-extra` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `bsdutils` | `1:2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libblkid-dev` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `util-linux` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libblkid1` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libmount-dev` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `mount` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libsmartcols1` | `2.38.1-5+deb12u3` |
| [CVE-2022-1210](https://nvd.nist.gov/vuln/detail/CVE-2022-1210) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2022-1210](https://nvd.nist.gov/vuln/detail/CVE-2022-1210) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2022-1210](https://nvd.nist.gov/vuln/detail/CVE-2022-1210) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2022-24975](https://nvd.nist.gov/vuln/detail/CVE-2022-24975) | Negligible | `git-man` | `1:2.39.5-0+deb12u3` |
| [CVE-2022-24975](https://nvd.nist.gov/vuln/detail/CVE-2022-24975) | Negligible | `git` | `1:2.39.5-0+deb12u3` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libgcc-12-dev` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libasan8` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libubsan1` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libstdc++6` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libstdc++-12-dev` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `liblsan0` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libitm1` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libhwasan0` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libgomp1` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `cpp-12` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `g++-12` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `gcc-12` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `gcc-12-base` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libgcc-s1` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libcc1-0` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libatomic1` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libtsan2` | `12.2.0-14+deb12u1` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gpg-wks-client` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gpg-wks-server` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gpgv` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gpgsm` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gnupg-l10n` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gnupg` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `dirmngr` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gpgconf` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gpg` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gpg-agent` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gnupg-utils` | `2.2.40-1.1+deb12u2` |
| [CVE-2023-1916](https://nvd.nist.gov/vuln/detail/CVE-2023-1916) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2023-1916](https://nvd.nist.gov/vuln/detail/CVE-2023-1916) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2023-1916](https://nvd.nist.gov/vuln/detail/CVE-2023-1916) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2023-1972](https://nvd.nist.gov/vuln/detail/CVE-2023-1972) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2023-1972](https://nvd.nist.gov/vuln/detail/CVE-2023-1972) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2023-1972](https://nvd.nist.gov/vuln/detail/CVE-2023-1972) | Negligible | `binutils` | `2.40-2` |
| [CVE-2023-1972](https://nvd.nist.gov/vuln/detail/CVE-2023-1972) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2023-1972](https://nvd.nist.gov/vuln/detail/CVE-2023-1972) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2023-1972](https://nvd.nist.gov/vuln/detail/CVE-2023-1972) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2023-1972](https://nvd.nist.gov/vuln/detail/CVE-2023-1972) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2023-31437](https://nvd.nist.gov/vuln/detail/CVE-2023-31437) | Negligible | `libsystemd0` | `252.39-1~deb12u1` |
| [CVE-2023-31437](https://nvd.nist.gov/vuln/detail/CVE-2023-31437) | Negligible | `libudev1` | `252.39-1~deb12u1` |
| [CVE-2023-31438](https://nvd.nist.gov/vuln/detail/CVE-2023-31438) | Negligible | `libudev1` | `252.39-1~deb12u1` |
| [CVE-2023-31438](https://nvd.nist.gov/vuln/detail/CVE-2023-31438) | Negligible | `libsystemd0` | `252.39-1~deb12u1` |
| [CVE-2023-31439](https://nvd.nist.gov/vuln/detail/CVE-2023-31439) | Negligible | `libudev1` | `252.39-1~deb12u1` |
| [CVE-2023-31439](https://nvd.nist.gov/vuln/detail/CVE-2023-31439) | Negligible | `libsystemd0` | `252.39-1~deb12u1` |
| [CVE-2023-31486](https://nvd.nist.gov/vuln/detail/CVE-2023-31486) | Negligible | `perl-modules-5.36` | `5.36.0-7+deb12u3` |
| [CVE-2023-31486](https://nvd.nist.gov/vuln/detail/CVE-2023-31486) | Negligible | `perl-base` | `5.36.0-7+deb12u3` |
| [CVE-2023-31486](https://nvd.nist.gov/vuln/detail/CVE-2023-31486) | Negligible | `perl` | `5.36.0-7+deb12u3` |
| [CVE-2023-31486](https://nvd.nist.gov/vuln/detail/CVE-2023-31486) | Negligible | `libperl5.36` | `5.36.0-7+deb12u3` |
| [CVE-2023-3164](https://nvd.nist.gov/vuln/detail/CVE-2023-3164) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2023-3164](https://nvd.nist.gov/vuln/detail/CVE-2023-3164) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2023-3164](https://nvd.nist.gov/vuln/detail/CVE-2023-3164) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-37769](https://nvd.nist.gov/vuln/detail/CVE-2023-37769) | Negligible | `libpixman-1-dev` | `0.42.2-1` |
| [CVE-2023-37769](https://nvd.nist.gov/vuln/detail/CVE-2023-37769) | Negligible | `libpixman-1-0` | `0.42.2-1` |
| [CVE-2023-49463](https://nvd.nist.gov/vuln/detail/CVE-2023-49463) | Negligible | `libheif1` | `1.15.1-1+deb12u1` |
| [CVE-2023-52426](https://nvd.nist.gov/vuln/detail/CVE-2023-52426) | Negligible | `libexpat1-dev` | `2.5.0-1+deb12u2` |
| [CVE-2023-52426](https://nvd.nist.gov/vuln/detail/CVE-2023-52426) | Negligible | `libexpat1` | `2.5.0-1+deb12u2` |
| [CVE-2023-6228](https://nvd.nist.gov/vuln/detail/CVE-2023-6228) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2023-6228](https://nvd.nist.gov/vuln/detail/CVE-2023-6228) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2023-6228](https://nvd.nist.gov/vuln/detail/CVE-2023-6228) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2024-2236](https://nvd.nist.gov/vuln/detail/CVE-2024-2236) | Negligible | `libgcrypt20` | `1.10.1-3` |
| [CVE-2024-2379](https://nvd.nist.gov/vuln/detail/CVE-2024-2379) | Negligible | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2024-2379](https://nvd.nist.gov/vuln/detail/CVE-2024-2379) | Negligible | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2024-2379](https://nvd.nist.gov/vuln/detail/CVE-2024-2379) | Negligible | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2024-2379](https://nvd.nist.gov/vuln/detail/CVE-2024-2379) | Negligible | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2024-25260](https://nvd.nist.gov/vuln/detail/CVE-2024-25260) | Negligible | `libelf1` | `0.188-2.1` |
| [CVE-2024-25269](https://nvd.nist.gov/vuln/detail/CVE-2024-25269) | Negligible | `libheif1` | `1.15.1-1+deb12u1` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libgssrpc4` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libkadm5clnt-mit12` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libk5crypto3` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libgssapi-krb5-2` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libkrb5support0` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `krb5-multidev` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libkrb5-dev` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libkrb5-3` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libkdb5-10` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libkadm5srv-mit12` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libkdb5-10` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libkrb5support0` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `krb5-multidev` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libgssrpc4` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libk5crypto3` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libkadm5clnt-mit12` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libkadm5srv-mit12` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libgssapi-krb5-2` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libkrb5-3` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libkrb5-dev` | `1.20.1-2+deb12u4` |
| [CVE-2024-28757](https://nvd.nist.gov/vuln/detail/CVE-2024-28757) | Negligible | `libexpat1` | `2.5.0-1+deb12u2` |
| [CVE-2024-28757](https://nvd.nist.gov/vuln/detail/CVE-2024-28757) | Negligible | `libexpat1-dev` | `2.5.0-1+deb12u2` |
| [CVE-2024-52005](https://nvd.nist.gov/vuln/detail/CVE-2024-52005) | Negligible | `git-man` | `1:2.39.5-0+deb12u3` |
| [CVE-2024-52005](https://nvd.nist.gov/vuln/detail/CVE-2024-52005) | Negligible | `git` | `1:2.39.5-0+deb12u3` |
| [CVE-2024-53589](https://nvd.nist.gov/vuln/detail/CVE-2024-53589) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2024-53589](https://nvd.nist.gov/vuln/detail/CVE-2024-53589) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2024-53589](https://nvd.nist.gov/vuln/detail/CVE-2024-53589) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2024-53589](https://nvd.nist.gov/vuln/detail/CVE-2024-53589) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2024-53589](https://nvd.nist.gov/vuln/detail/CVE-2024-53589) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2024-53589](https://nvd.nist.gov/vuln/detail/CVE-2024-53589) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2024-53589](https://nvd.nist.gov/vuln/detail/CVE-2024-53589) | Negligible | `binutils` | `2.40-2` |
| [CVE-2024-57360](https://nvd.nist.gov/vuln/detail/CVE-2024-57360) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2024-57360](https://nvd.nist.gov/vuln/detail/CVE-2024-57360) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2024-57360](https://nvd.nist.gov/vuln/detail/CVE-2024-57360) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2024-57360](https://nvd.nist.gov/vuln/detail/CVE-2024-57360) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2024-57360](https://nvd.nist.gov/vuln/detail/CVE-2024-57360) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2024-57360](https://nvd.nist.gov/vuln/detail/CVE-2024-57360) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2024-57360](https://nvd.nist.gov/vuln/detail/CVE-2024-57360) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-0725](https://nvd.nist.gov/vuln/detail/CVE-2025-0725) | Negligible | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-0725](https://nvd.nist.gov/vuln/detail/CVE-2025-0725) | Negligible | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-0725](https://nvd.nist.gov/vuln/detail/CVE-2025-0725) | Negligible | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-0725](https://nvd.nist.gov/vuln/detail/CVE-2025-0725) | Negligible | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-0840](https://nvd.nist.gov/vuln/detail/CVE-2025-0840) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-0840](https://nvd.nist.gov/vuln/detail/CVE-2025-0840) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-0840](https://nvd.nist.gov/vuln/detail/CVE-2025-0840) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-0840](https://nvd.nist.gov/vuln/detail/CVE-2025-0840) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-0840](https://nvd.nist.gov/vuln/detail/CVE-2025-0840) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-0840](https://nvd.nist.gov/vuln/detail/CVE-2025-0840) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-0840](https://nvd.nist.gov/vuln/detail/CVE-2025-0840) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-10966](https://nvd.nist.gov/vuln/detail/CVE-2025-10966) | Negligible | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-10966](https://nvd.nist.gov/vuln/detail/CVE-2025-10966) | Negligible | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-10966](https://nvd.nist.gov/vuln/detail/CVE-2025-10966) | Negligible | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-10966](https://nvd.nist.gov/vuln/detail/CVE-2025-10966) | Negligible | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-11081](https://nvd.nist.gov/vuln/detail/CVE-2025-11081) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11081](https://nvd.nist.gov/vuln/detail/CVE-2025-11081) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11081](https://nvd.nist.gov/vuln/detail/CVE-2025-11081) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11081](https://nvd.nist.gov/vuln/detail/CVE-2025-11081) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11081](https://nvd.nist.gov/vuln/detail/CVE-2025-11081) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11081](https://nvd.nist.gov/vuln/detail/CVE-2025-11081) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11081](https://nvd.nist.gov/vuln/detail/CVE-2025-11081) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11082](https://nvd.nist.gov/vuln/detail/CVE-2025-11082) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11082](https://nvd.nist.gov/vuln/detail/CVE-2025-11082) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11082](https://nvd.nist.gov/vuln/detail/CVE-2025-11082) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11082](https://nvd.nist.gov/vuln/detail/CVE-2025-11082) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11082](https://nvd.nist.gov/vuln/detail/CVE-2025-11082) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11082](https://nvd.nist.gov/vuln/detail/CVE-2025-11082) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11082](https://nvd.nist.gov/vuln/detail/CVE-2025-11082) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11083](https://nvd.nist.gov/vuln/detail/CVE-2025-11083) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11083](https://nvd.nist.gov/vuln/detail/CVE-2025-11083) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11083](https://nvd.nist.gov/vuln/detail/CVE-2025-11083) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11083](https://nvd.nist.gov/vuln/detail/CVE-2025-11083) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11083](https://nvd.nist.gov/vuln/detail/CVE-2025-11083) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11083](https://nvd.nist.gov/vuln/detail/CVE-2025-11083) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11083](https://nvd.nist.gov/vuln/detail/CVE-2025-11083) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11412](https://nvd.nist.gov/vuln/detail/CVE-2025-11412) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11412](https://nvd.nist.gov/vuln/detail/CVE-2025-11412) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11412](https://nvd.nist.gov/vuln/detail/CVE-2025-11412) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11412](https://nvd.nist.gov/vuln/detail/CVE-2025-11412) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11412](https://nvd.nist.gov/vuln/detail/CVE-2025-11412) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11412](https://nvd.nist.gov/vuln/detail/CVE-2025-11412) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11412](https://nvd.nist.gov/vuln/detail/CVE-2025-11412) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11413](https://nvd.nist.gov/vuln/detail/CVE-2025-11413) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11413](https://nvd.nist.gov/vuln/detail/CVE-2025-11413) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11413](https://nvd.nist.gov/vuln/detail/CVE-2025-11413) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11413](https://nvd.nist.gov/vuln/detail/CVE-2025-11413) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11413](https://nvd.nist.gov/vuln/detail/CVE-2025-11413) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11413](https://nvd.nist.gov/vuln/detail/CVE-2025-11413) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11413](https://nvd.nist.gov/vuln/detail/CVE-2025-11413) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11414](https://nvd.nist.gov/vuln/detail/CVE-2025-11414) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11414](https://nvd.nist.gov/vuln/detail/CVE-2025-11414) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11414](https://nvd.nist.gov/vuln/detail/CVE-2025-11414) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11414](https://nvd.nist.gov/vuln/detail/CVE-2025-11414) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11414](https://nvd.nist.gov/vuln/detail/CVE-2025-11414) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11414](https://nvd.nist.gov/vuln/detail/CVE-2025-11414) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11414](https://nvd.nist.gov/vuln/detail/CVE-2025-11414) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11468](https://nvd.nist.gov/vuln/detail/CVE-2025-11468) | Negligible | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-11468](https://nvd.nist.gov/vuln/detail/CVE-2025-11468) | Negligible | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-11468](https://nvd.nist.gov/vuln/detail/CVE-2025-11468) | Negligible | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-11468](https://nvd.nist.gov/vuln/detail/CVE-2025-11468) | Negligible | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-1147](https://nvd.nist.gov/vuln/detail/CVE-2025-1147) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1147](https://nvd.nist.gov/vuln/detail/CVE-2025-1147) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1147](https://nvd.nist.gov/vuln/detail/CVE-2025-1147) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1147](https://nvd.nist.gov/vuln/detail/CVE-2025-1147) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1147](https://nvd.nist.gov/vuln/detail/CVE-2025-1147) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1147](https://nvd.nist.gov/vuln/detail/CVE-2025-1147) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1147](https://nvd.nist.gov/vuln/detail/CVE-2025-1147) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1148](https://nvd.nist.gov/vuln/detail/CVE-2025-1148) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1148](https://nvd.nist.gov/vuln/detail/CVE-2025-1148) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1148](https://nvd.nist.gov/vuln/detail/CVE-2025-1148) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1148](https://nvd.nist.gov/vuln/detail/CVE-2025-1148) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1148](https://nvd.nist.gov/vuln/detail/CVE-2025-1148) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1148](https://nvd.nist.gov/vuln/detail/CVE-2025-1148) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1148](https://nvd.nist.gov/vuln/detail/CVE-2025-1148) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1149](https://nvd.nist.gov/vuln/detail/CVE-2025-1149) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1149](https://nvd.nist.gov/vuln/detail/CVE-2025-1149) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1149](https://nvd.nist.gov/vuln/detail/CVE-2025-1149) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1149](https://nvd.nist.gov/vuln/detail/CVE-2025-1149) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1149](https://nvd.nist.gov/vuln/detail/CVE-2025-1149) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1149](https://nvd.nist.gov/vuln/detail/CVE-2025-1149) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1149](https://nvd.nist.gov/vuln/detail/CVE-2025-1149) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11494](https://nvd.nist.gov/vuln/detail/CVE-2025-11494) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11494](https://nvd.nist.gov/vuln/detail/CVE-2025-11494) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11494](https://nvd.nist.gov/vuln/detail/CVE-2025-11494) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11494](https://nvd.nist.gov/vuln/detail/CVE-2025-11494) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11494](https://nvd.nist.gov/vuln/detail/CVE-2025-11494) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11494](https://nvd.nist.gov/vuln/detail/CVE-2025-11494) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11494](https://nvd.nist.gov/vuln/detail/CVE-2025-11494) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11495](https://nvd.nist.gov/vuln/detail/CVE-2025-11495) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11495](https://nvd.nist.gov/vuln/detail/CVE-2025-11495) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11495](https://nvd.nist.gov/vuln/detail/CVE-2025-11495) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11495](https://nvd.nist.gov/vuln/detail/CVE-2025-11495) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11495](https://nvd.nist.gov/vuln/detail/CVE-2025-11495) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11495](https://nvd.nist.gov/vuln/detail/CVE-2025-11495) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11495](https://nvd.nist.gov/vuln/detail/CVE-2025-11495) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1150](https://nvd.nist.gov/vuln/detail/CVE-2025-1150) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1150](https://nvd.nist.gov/vuln/detail/CVE-2025-1150) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1150](https://nvd.nist.gov/vuln/detail/CVE-2025-1150) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1150](https://nvd.nist.gov/vuln/detail/CVE-2025-1150) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1150](https://nvd.nist.gov/vuln/detail/CVE-2025-1150) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1150](https://nvd.nist.gov/vuln/detail/CVE-2025-1150) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1150](https://nvd.nist.gov/vuln/detail/CVE-2025-1150) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1151](https://nvd.nist.gov/vuln/detail/CVE-2025-1151) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1151](https://nvd.nist.gov/vuln/detail/CVE-2025-1151) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1151](https://nvd.nist.gov/vuln/detail/CVE-2025-1151) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1151](https://nvd.nist.gov/vuln/detail/CVE-2025-1151) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1151](https://nvd.nist.gov/vuln/detail/CVE-2025-1151) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1151](https://nvd.nist.gov/vuln/detail/CVE-2025-1151) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1151](https://nvd.nist.gov/vuln/detail/CVE-2025-1151) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1152](https://nvd.nist.gov/vuln/detail/CVE-2025-1152) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1152](https://nvd.nist.gov/vuln/detail/CVE-2025-1152) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1152](https://nvd.nist.gov/vuln/detail/CVE-2025-1152) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1152](https://nvd.nist.gov/vuln/detail/CVE-2025-1152) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1152](https://nvd.nist.gov/vuln/detail/CVE-2025-1152) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1152](https://nvd.nist.gov/vuln/detail/CVE-2025-1152) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1152](https://nvd.nist.gov/vuln/detail/CVE-2025-1152) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1153](https://nvd.nist.gov/vuln/detail/CVE-2025-1153) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1153](https://nvd.nist.gov/vuln/detail/CVE-2025-1153) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1153](https://nvd.nist.gov/vuln/detail/CVE-2025-1153) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1153](https://nvd.nist.gov/vuln/detail/CVE-2025-1153) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1153](https://nvd.nist.gov/vuln/detail/CVE-2025-1153) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1153](https://nvd.nist.gov/vuln/detail/CVE-2025-1153) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1153](https://nvd.nist.gov/vuln/detail/CVE-2025-1153) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1176](https://nvd.nist.gov/vuln/detail/CVE-2025-1176) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1176](https://nvd.nist.gov/vuln/detail/CVE-2025-1176) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1176](https://nvd.nist.gov/vuln/detail/CVE-2025-1176) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1176](https://nvd.nist.gov/vuln/detail/CVE-2025-1176) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1176](https://nvd.nist.gov/vuln/detail/CVE-2025-1176) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1176](https://nvd.nist.gov/vuln/detail/CVE-2025-1176) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1176](https://nvd.nist.gov/vuln/detail/CVE-2025-1176) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1178](https://nvd.nist.gov/vuln/detail/CVE-2025-1178) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1178](https://nvd.nist.gov/vuln/detail/CVE-2025-1178) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1178](https://nvd.nist.gov/vuln/detail/CVE-2025-1178) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1178](https://nvd.nist.gov/vuln/detail/CVE-2025-1178) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1178](https://nvd.nist.gov/vuln/detail/CVE-2025-1178) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1178](https://nvd.nist.gov/vuln/detail/CVE-2025-1178) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1178](https://nvd.nist.gov/vuln/detail/CVE-2025-1178) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1179](https://nvd.nist.gov/vuln/detail/CVE-2025-1179) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1179](https://nvd.nist.gov/vuln/detail/CVE-2025-1179) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1179](https://nvd.nist.gov/vuln/detail/CVE-2025-1179) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1179](https://nvd.nist.gov/vuln/detail/CVE-2025-1179) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1179](https://nvd.nist.gov/vuln/detail/CVE-2025-1179) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1179](https://nvd.nist.gov/vuln/detail/CVE-2025-1179) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1179](https://nvd.nist.gov/vuln/detail/CVE-2025-1179) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1180](https://nvd.nist.gov/vuln/detail/CVE-2025-1180) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1180](https://nvd.nist.gov/vuln/detail/CVE-2025-1180) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1180](https://nvd.nist.gov/vuln/detail/CVE-2025-1180) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1180](https://nvd.nist.gov/vuln/detail/CVE-2025-1180) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1180](https://nvd.nist.gov/vuln/detail/CVE-2025-1180) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1180](https://nvd.nist.gov/vuln/detail/CVE-2025-1180) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1180](https://nvd.nist.gov/vuln/detail/CVE-2025-1180) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1181](https://nvd.nist.gov/vuln/detail/CVE-2025-1181) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1181](https://nvd.nist.gov/vuln/detail/CVE-2025-1181) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1181](https://nvd.nist.gov/vuln/detail/CVE-2025-1181) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1181](https://nvd.nist.gov/vuln/detail/CVE-2025-1181) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1181](https://nvd.nist.gov/vuln/detail/CVE-2025-1181) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1181](https://nvd.nist.gov/vuln/detail/CVE-2025-1181) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1181](https://nvd.nist.gov/vuln/detail/CVE-2025-1181) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1182](https://nvd.nist.gov/vuln/detail/CVE-2025-1182) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1182](https://nvd.nist.gov/vuln/detail/CVE-2025-1182) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1182](https://nvd.nist.gov/vuln/detail/CVE-2025-1182) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1182](https://nvd.nist.gov/vuln/detail/CVE-2025-1182) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1182](https://nvd.nist.gov/vuln/detail/CVE-2025-1182) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1182](https://nvd.nist.gov/vuln/detail/CVE-2025-1182) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1182](https://nvd.nist.gov/vuln/detail/CVE-2025-1182) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11839](https://nvd.nist.gov/vuln/detail/CVE-2025-11839) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11839](https://nvd.nist.gov/vuln/detail/CVE-2025-11839) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11839](https://nvd.nist.gov/vuln/detail/CVE-2025-11839) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11839](https://nvd.nist.gov/vuln/detail/CVE-2025-11839) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11839](https://nvd.nist.gov/vuln/detail/CVE-2025-11839) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11839](https://nvd.nist.gov/vuln/detail/CVE-2025-11839) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11839](https://nvd.nist.gov/vuln/detail/CVE-2025-11839) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11840](https://nvd.nist.gov/vuln/detail/CVE-2025-11840) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11840](https://nvd.nist.gov/vuln/detail/CVE-2025-11840) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11840](https://nvd.nist.gov/vuln/detail/CVE-2025-11840) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11840](https://nvd.nist.gov/vuln/detail/CVE-2025-11840) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11840](https://nvd.nist.gov/vuln/detail/CVE-2025-11840) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11840](https://nvd.nist.gov/vuln/detail/CVE-2025-11840) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11840](https://nvd.nist.gov/vuln/detail/CVE-2025-11840) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1352](https://nvd.nist.gov/vuln/detail/CVE-2025-1352) | Negligible | `libelf1` | `0.188-2.1` |
| [CVE-2025-1365](https://nvd.nist.gov/vuln/detail/CVE-2025-1365) | Negligible | `libelf1` | `0.188-2.1` |
| [CVE-2025-1371](https://nvd.nist.gov/vuln/detail/CVE-2025-1371) | Negligible | `libelf1` | `0.188-2.1` |
| [CVE-2025-1372](https://nvd.nist.gov/vuln/detail/CVE-2025-1372) | Negligible | `libelf1` | `0.188-2.1` |
| [CVE-2025-1376](https://nvd.nist.gov/vuln/detail/CVE-2025-1376) | Negligible | `libelf1` | `0.188-2.1` |
| [CVE-2025-1377](https://nvd.nist.gov/vuln/detail/CVE-2025-1377) | Negligible | `libelf1` | `0.188-2.1` |
| [CVE-2025-14017](https://nvd.nist.gov/vuln/detail/CVE-2025-14017) | Negligible | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-14017](https://nvd.nist.gov/vuln/detail/CVE-2025-14017) | Negligible | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-14017](https://nvd.nist.gov/vuln/detail/CVE-2025-14017) | Negligible | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-14017](https://nvd.nist.gov/vuln/detail/CVE-2025-14017) | Negligible | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-15079](https://nvd.nist.gov/vuln/detail/CVE-2025-15079) | Negligible | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-15079](https://nvd.nist.gov/vuln/detail/CVE-2025-15079) | Negligible | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-15079](https://nvd.nist.gov/vuln/detail/CVE-2025-15079) | Negligible | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-15079](https://nvd.nist.gov/vuln/detail/CVE-2025-15079) | Negligible | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-15224](https://nvd.nist.gov/vuln/detail/CVE-2025-15224) | Negligible | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-15224](https://nvd.nist.gov/vuln/detail/CVE-2025-15224) | Negligible | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-15224](https://nvd.nist.gov/vuln/detail/CVE-2025-15224) | Negligible | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-15224](https://nvd.nist.gov/vuln/detail/CVE-2025-15224) | Negligible | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-27587](https://nvd.nist.gov/vuln/detail/CVE-2025-27587) | Negligible | `libssl-dev` | `3.0.18-1~deb12u2` |
| [CVE-2025-27587](https://nvd.nist.gov/vuln/detail/CVE-2025-27587) | Negligible | `libssl3` | `3.0.18-1~deb12u2` |
| [CVE-2025-27587](https://nvd.nist.gov/vuln/detail/CVE-2025-27587) | Negligible | `openssl` | `3.0.18-1~deb12u2` |
| [CVE-2025-28162](https://nvd.nist.gov/vuln/detail/CVE-2025-28162) | Negligible | `libpng16-16` | `1.6.39-2+deb12u1` |
| [CVE-2025-28162](https://nvd.nist.gov/vuln/detail/CVE-2025-28162) | Negligible | `libpng-dev` | `1.6.39-2+deb12u1` |
| [CVE-2025-28164](https://nvd.nist.gov/vuln/detail/CVE-2025-28164) | Negligible | `libpng-dev` | `1.6.39-2+deb12u1` |
| [CVE-2025-28164](https://nvd.nist.gov/vuln/detail/CVE-2025-28164) | Negligible | `libpng16-16` | `1.6.39-2+deb12u1` |
| [CVE-2025-29070](https://nvd.nist.gov/vuln/detail/CVE-2025-29070) | Negligible | `liblcms2-2` | `2.14-2` |
| [CVE-2025-29070](https://nvd.nist.gov/vuln/detail/CVE-2025-29070) | Negligible | `liblcms2-dev` | `2.14-2` |
| [CVE-2025-29088](https://nvd.nist.gov/vuln/detail/CVE-2025-29088) | Negligible | `libsqlite3-dev` | `3.40.1-2+deb12u2` |
| [CVE-2025-29088](https://nvd.nist.gov/vuln/detail/CVE-2025-29088) | Negligible | `libsqlite3-0` | `3.40.1-2+deb12u2` |
| [CVE-2025-3198](https://nvd.nist.gov/vuln/detail/CVE-2025-3198) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-3198](https://nvd.nist.gov/vuln/detail/CVE-2025-3198) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-3198](https://nvd.nist.gov/vuln/detail/CVE-2025-3198) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-3198](https://nvd.nist.gov/vuln/detail/CVE-2025-3198) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-3198](https://nvd.nist.gov/vuln/detail/CVE-2025-3198) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-3198](https://nvd.nist.gov/vuln/detail/CVE-2025-3198) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-3198](https://nvd.nist.gov/vuln/detail/CVE-2025-3198) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-4516](https://nvd.nist.gov/vuln/detail/CVE-2025-4516) | Negligible | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-4516](https://nvd.nist.gov/vuln/detail/CVE-2025-4516) | Negligible | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-4516](https://nvd.nist.gov/vuln/detail/CVE-2025-4516) | Negligible | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-4516](https://nvd.nist.gov/vuln/detail/CVE-2025-4516) | Negligible | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-5244](https://nvd.nist.gov/vuln/detail/CVE-2025-5244) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-5244](https://nvd.nist.gov/vuln/detail/CVE-2025-5244) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-5244](https://nvd.nist.gov/vuln/detail/CVE-2025-5244) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-5244](https://nvd.nist.gov/vuln/detail/CVE-2025-5244) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-5244](https://nvd.nist.gov/vuln/detail/CVE-2025-5244) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-5244](https://nvd.nist.gov/vuln/detail/CVE-2025-5244) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-5244](https://nvd.nist.gov/vuln/detail/CVE-2025-5244) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-5245](https://nvd.nist.gov/vuln/detail/CVE-2025-5245) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-5245](https://nvd.nist.gov/vuln/detail/CVE-2025-5245) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-5245](https://nvd.nist.gov/vuln/detail/CVE-2025-5245) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-5245](https://nvd.nist.gov/vuln/detail/CVE-2025-5245) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-5245](https://nvd.nist.gov/vuln/detail/CVE-2025-5245) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-5245](https://nvd.nist.gov/vuln/detail/CVE-2025-5245) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-5245](https://nvd.nist.gov/vuln/detail/CVE-2025-5245) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-5278](https://nvd.nist.gov/vuln/detail/CVE-2025-5278) | Negligible | `coreutils` | `9.1-1` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-66861](https://nvd.nist.gov/vuln/detail/CVE-2025-66861) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-66861](https://nvd.nist.gov/vuln/detail/CVE-2025-66861) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-66861](https://nvd.nist.gov/vuln/detail/CVE-2025-66861) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-66861](https://nvd.nist.gov/vuln/detail/CVE-2025-66861) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-66861](https://nvd.nist.gov/vuln/detail/CVE-2025-66861) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-66861](https://nvd.nist.gov/vuln/detail/CVE-2025-66861) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-66861](https://nvd.nist.gov/vuln/detail/CVE-2025-66861) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-66862](https://nvd.nist.gov/vuln/detail/CVE-2025-66862) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-66862](https://nvd.nist.gov/vuln/detail/CVE-2025-66862) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-66862](https://nvd.nist.gov/vuln/detail/CVE-2025-66862) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-66862](https://nvd.nist.gov/vuln/detail/CVE-2025-66862) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-66862](https://nvd.nist.gov/vuln/detail/CVE-2025-66862) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-66862](https://nvd.nist.gov/vuln/detail/CVE-2025-66862) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-66862](https://nvd.nist.gov/vuln/detail/CVE-2025-66862) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-66863](https://nvd.nist.gov/vuln/detail/CVE-2025-66863) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-66863](https://nvd.nist.gov/vuln/detail/CVE-2025-66863) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-66863](https://nvd.nist.gov/vuln/detail/CVE-2025-66863) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-66863](https://nvd.nist.gov/vuln/detail/CVE-2025-66863) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-66863](https://nvd.nist.gov/vuln/detail/CVE-2025-66863) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-66863](https://nvd.nist.gov/vuln/detail/CVE-2025-66863) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-66863](https://nvd.nist.gov/vuln/detail/CVE-2025-66863) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-66864](https://nvd.nist.gov/vuln/detail/CVE-2025-66864) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-66864](https://nvd.nist.gov/vuln/detail/CVE-2025-66864) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-66864](https://nvd.nist.gov/vuln/detail/CVE-2025-66864) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-66864](https://nvd.nist.gov/vuln/detail/CVE-2025-66864) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-66864](https://nvd.nist.gov/vuln/detail/CVE-2025-66864) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-66864](https://nvd.nist.gov/vuln/detail/CVE-2025-66864) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-66864](https://nvd.nist.gov/vuln/detail/CVE-2025-66864) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-66865](https://nvd.nist.gov/vuln/detail/CVE-2025-66865) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-66865](https://nvd.nist.gov/vuln/detail/CVE-2025-66865) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-66865](https://nvd.nist.gov/vuln/detail/CVE-2025-66865) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-66865](https://nvd.nist.gov/vuln/detail/CVE-2025-66865) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-66865](https://nvd.nist.gov/vuln/detail/CVE-2025-66865) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-66865](https://nvd.nist.gov/vuln/detail/CVE-2025-66865) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-66865](https://nvd.nist.gov/vuln/detail/CVE-2025-66865) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-66866](https://nvd.nist.gov/vuln/detail/CVE-2025-66866) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-66866](https://nvd.nist.gov/vuln/detail/CVE-2025-66866) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-66866](https://nvd.nist.gov/vuln/detail/CVE-2025-66866) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-66866](https://nvd.nist.gov/vuln/detail/CVE-2025-66866) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-66866](https://nvd.nist.gov/vuln/detail/CVE-2025-66866) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-66866](https://nvd.nist.gov/vuln/detail/CVE-2025-66866) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-66866](https://nvd.nist.gov/vuln/detail/CVE-2025-66866) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-7545](https://nvd.nist.gov/vuln/detail/CVE-2025-7545) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-7545](https://nvd.nist.gov/vuln/detail/CVE-2025-7545) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-7545](https://nvd.nist.gov/vuln/detail/CVE-2025-7545) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-7545](https://nvd.nist.gov/vuln/detail/CVE-2025-7545) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-7545](https://nvd.nist.gov/vuln/detail/CVE-2025-7545) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-7545](https://nvd.nist.gov/vuln/detail/CVE-2025-7545) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-7545](https://nvd.nist.gov/vuln/detail/CVE-2025-7545) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-7546](https://nvd.nist.gov/vuln/detail/CVE-2025-7546) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-7546](https://nvd.nist.gov/vuln/detail/CVE-2025-7546) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-7546](https://nvd.nist.gov/vuln/detail/CVE-2025-7546) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-7546](https://nvd.nist.gov/vuln/detail/CVE-2025-7546) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-7546](https://nvd.nist.gov/vuln/detail/CVE-2025-7546) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-7546](https://nvd.nist.gov/vuln/detail/CVE-2025-7546) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-7546](https://nvd.nist.gov/vuln/detail/CVE-2025-7546) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-8176](https://nvd.nist.gov/vuln/detail/CVE-2025-8176) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8176](https://nvd.nist.gov/vuln/detail/CVE-2025-8176) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2025-8176](https://nvd.nist.gov/vuln/detail/CVE-2025-8176) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8177](https://nvd.nist.gov/vuln/detail/CVE-2025-8177) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8177](https://nvd.nist.gov/vuln/detail/CVE-2025-8177) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8177](https://nvd.nist.gov/vuln/detail/CVE-2025-8177) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2025-8224](https://nvd.nist.gov/vuln/detail/CVE-2025-8224) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-8224](https://nvd.nist.gov/vuln/detail/CVE-2025-8224) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-8224](https://nvd.nist.gov/vuln/detail/CVE-2025-8224) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-8224](https://nvd.nist.gov/vuln/detail/CVE-2025-8224) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-8224](https://nvd.nist.gov/vuln/detail/CVE-2025-8224) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-8224](https://nvd.nist.gov/vuln/detail/CVE-2025-8224) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-8224](https://nvd.nist.gov/vuln/detail/CVE-2025-8224) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-8225](https://nvd.nist.gov/vuln/detail/CVE-2025-8225) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-8225](https://nvd.nist.gov/vuln/detail/CVE-2025-8225) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-8225](https://nvd.nist.gov/vuln/detail/CVE-2025-8225) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-8225](https://nvd.nist.gov/vuln/detail/CVE-2025-8225) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-8225](https://nvd.nist.gov/vuln/detail/CVE-2025-8225) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-8225](https://nvd.nist.gov/vuln/detail/CVE-2025-8225) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-8225](https://nvd.nist.gov/vuln/detail/CVE-2025-8225) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-8534](https://nvd.nist.gov/vuln/detail/CVE-2025-8534) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8534](https://nvd.nist.gov/vuln/detail/CVE-2025-8534) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2025-8534](https://nvd.nist.gov/vuln/detail/CVE-2025-8534) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8732](https://nvd.nist.gov/vuln/detail/CVE-2025-8732) | Negligible | `libxml2-dev` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2025-8732](https://nvd.nist.gov/vuln/detail/CVE-2025-8732) | Negligible | `libxml2` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2025-8851](https://nvd.nist.gov/vuln/detail/CVE-2025-8851) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2025-8851](https://nvd.nist.gov/vuln/detail/CVE-2025-8851) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8851](https://nvd.nist.gov/vuln/detail/CVE-2025-8851) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8961](https://nvd.nist.gov/vuln/detail/CVE-2025-8961) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8961](https://nvd.nist.gov/vuln/detail/CVE-2025-8961) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8961](https://nvd.nist.gov/vuln/detail/CVE-2025-8961) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2025-9165](https://nvd.nist.gov/vuln/detail/CVE-2025-9165) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2025-9165](https://nvd.nist.gov/vuln/detail/CVE-2025-9165) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2025-9165](https://nvd.nist.gov/vuln/detail/CVE-2025-9165) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2026-1757](https://nvd.nist.gov/vuln/detail/CVE-2026-1757) | Negligible | `libxml2-dev` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-1757](https://nvd.nist.gov/vuln/detail/CVE-2026-1757) | Negligible | `libxml2` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-22185](https://nvd.nist.gov/vuln/detail/CVE-2026-22185) | Negligible | `libldap-2.5-0` | `2.5.13+dfsg-5` |
</details>

<details>
<summary><strong>📦 Installed Packages (611 total)</strong></summary>

| Package | Version |
|---------|---------|
| @isaacs/cliui | 8.0.2 |
| @isaacs/fs-minipass | 4.0.1 |
| @isaacs/string-locale-compare | 1.1.0 |
| @npmcli/agent | 3.0.0 |
| @npmcli/arborist | 8.0.1 |
| @npmcli/config | 9.0.0 |
| @npmcli/fs | 4.0.0 |
| @npmcli/git | 6.0.3 |
| @npmcli/installed-package-contents | 3.0.0 |
| @npmcli/map-workspaces | 4.0.2 |
| @npmcli/metavuln-calculator | 8.0.1 |
| @npmcli/name-from-folder | 3.0.0 |
| @npmcli/node-gyp | 4.0.0 |
| @npmcli/package-json | 6.2.0 |
| @npmcli/promise-spawn | 8.0.2 |
| @npmcli/query | 4.0.1 |
| @npmcli/redact | 3.2.2 |
| @npmcli/run-script | 9.1.0 |
| @pkgjs/parseargs | 0.11.0 |
| @sigstore/bundle | 3.1.0 |
| @sigstore/core | 2.0.0 |
| @sigstore/protobuf-specs | 0.4.3 |
| @sigstore/sign | 3.1.0 |
| @sigstore/tuf | 3.1.1 |
| @sigstore/verify | 2.1.1 |
| @tufjs/canonical-json | 2.0.0 |
| @tufjs/models | 3.0.1 |
| abbrev | 3.0.1 |
| adduser | 3.134 |
| agent-base | 7.1.3 |
| ansi-regex | 5.0.1 |
| ansi-regex | 6.1.0 |
| ansi-styles | 4.3.0 |
| ansi-styles | 6.2.1 |
| aproba | 2.0.0 |
| apt | 2.6.1 |
| archy | 1.0.0 |
| autoconf | 2.71-3 |
| automake | 1:1.16.5-1.3 |
| autotools-dev | 20220109.1 |
| balanced-match | 1.0.2 |
| base-files | 12.4+deb12u13 |
| base-passwd | 3.6.1 |
| bash | 5.2.15-2+b10 |
| bin-links | 5.0.0 |
| binary-extensions | 2.3.0 |
| binutils | 2.40-2 |
| binutils-aarch64-linux-gnu | 2.40-2 |
| binutils-common | 2.40-2 |
| brace-expansion | 2.0.2 |
| bsdutils | 1:2.38.1-5+deb12u3 |
| bzip2 | 1.0.8-5+b1 |
| ca-certificates | 20230311+deb12u1 |
| cacache | 19.0.1 |
| chalk | 5.4.1 |
| chownr | 2.0.0 |
| chownr | 3.0.0 |
| ci-info | 4.2.0 |
| cidr-regex | 4.1.3 |
| cli-columns | 4.0.0 |
| cmd-shim | 7.0.0 |
| color-convert | 2.0.1 |
| color-name | 1.1.4 |
| comerr-dev | 2.1-1.47.0-2+b2 |
| common-ancestor-path | 1.0.1 |
| corepack | 0.34.0 |
| coreutils | 9.1-1 |
| cpp | 4:12.2.0-3 |
| cpp-12 | 12.2.0-14+deb12u1 |
| cross-spawn | 7.0.6 |
| cssesc | 3.0.0 |
| curl | 7.88.1-10+deb12u14 |
| dash | 0.5.12-2 |
| debconf | 1.5.82 |
| debian-archive-keyring | 2023.3+deb12u2 |
| debianutils | 5.7-0.5~deb12u1 |
| debug | 4.4.1 |
| default-libmysqlclient-dev | 1.1.0 |
| diff | 5.2.0 |
| diffutils | 1:3.8-4 |
| dirmngr | 2.2.40-1.1+deb12u2 |
| dpkg | 1.21.22 |
| dpkg-dev | 1.21.22 |
| e2fsprogs | 1.47.0-2+b2 |
| eastasianwidth | 0.2.0 |
| emoji-regex | 8.0.0 |
| emoji-regex | 9.2.2 |
| encoding | 0.1.13 |
| env-paths | 2.2.1 |
| err-code | 2.0.3 |
| exponential-backoff | 3.1.2 |
| fastest-levenshtein | 1.0.16 |
| fdir | 6.4.6 |
| file | 1:5.44-3 |
| findutils | 4.9.0-4 |
| fontconfig | 2.14.1-4 |
| fontconfig-config | 2.14.1-4 |
| fonts-dejavu-core | 2.37-6 |
| foreground-child | 3.3.1 |
| fs-minipass | 2.1.0 |
| fs-minipass | 3.0.3 |
| g++ | 4:12.2.0-3 |
| g++-12 | 12.2.0-14+deb12u1 |
| gcc | 4:12.2.0-3 |
| gcc-12 | 12.2.0-14+deb12u1 |
| gcc-12-base | 12.2.0-14+deb12u1 |
| gir1.2-freedesktop | 1.74.0-3 |
| gir1.2-gdkpixbuf-2.0 | 2.42.10+dfsg-1+deb12u3 |
| gir1.2-glib-2.0 | 1.74.0-3 |
| gir1.2-rsvg-2.0 | 2.54.7+dfsg-1~deb12u1 |
| git | 1:2.39.5-0+deb12u3 |
| git-man | 1:2.39.5-0+deb12u3 |
| glob | 10.4.5 |
| gnupg | 2.2.40-1.1+deb12u2 |
| gnupg-l10n | 2.2.40-1.1+deb12u2 |
| gnupg-utils | 2.2.40-1.1+deb12u2 |
| gpg | 2.2.40-1.1+deb12u2 |
| gpg-agent | 2.2.40-1.1+deb12u2 |
| gpg-wks-client | 2.2.40-1.1+deb12u2 |
| gpg-wks-server | 2.2.40-1.1+deb12u2 |
| gpgconf | 2.2.40-1.1+deb12u2 |
| gpgsm | 2.2.40-1.1+deb12u2 |
| gpgv | 2.2.40-1.1+deb12u2 |
| graceful-fs | 4.2.11 |
| grep | 3.8-5 |
| gzip | 1.12-1 |
| hicolor-icon-theme | 0.17-2 |
| hosted-git-info | 8.1.0 |
| hostname | 3.23+nmu1 |
| http-cache-semantics | 4.2.0 |
| http-proxy-agent | 7.0.2 |
| https-proxy-agent | 7.0.6 |
| iconv-lite | 0.6.3 |
| icu-devtools | 72.1-3+deb12u1 |
| ignore-walk | 7.0.0 |
| imagemagick | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| imagemagick-6-common | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| imagemagick-6.q16 | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| imurmurhash | 0.1.4 |
| ini | 5.0.0 |
| init-package-json | 7.0.2 |
| init-system-helpers | 1.65.2+deb12u1 |
| ip-address | 9.0.5 |
| ip-regex | 5.0.0 |
| is-cidr | 5.1.1 |
| is-fullwidth-code-point | 3.0.0 |
| isexe | 2.0.0 |
| isexe | 3.1.1 |
| jackspeak | 3.4.3 |
| jsbn | 1.1.0 |
| json-parse-even-better-errors | 4.0.0 |
| json-stringify-nice | 1.1.4 |
| jsonparse | 1.3.1 |
| just-diff | 6.0.2 |
| just-diff-apply | 5.5.0 |
| krb5-multidev | 1.20.1-2+deb12u4 |
| libacl1 | 2.3.1-3 |
| libaom3 | 3.6.0-1+deb12u2 |
| libapr1 | 1.7.2-3+deb12u1 |
| libaprutil1 | 1.6.3-1 |
| libapt-pkg6.0 | 2.6.1 |
| libasan8 | 12.2.0-14+deb12u1 |
| libassuan0 | 2.5.5-5 |
| libatomic1 | 12.2.0-14+deb12u1 |
| libattr1 | 1:2.5.1-4 |
| libaudit-common | 1:3.0.9-1 |
| libaudit1 | 1:3.0.9-1 |
| libbinutils | 2.40-2 |
| libblkid-dev | 2.38.1-5+deb12u3 |
| libblkid1 | 2.38.1-5+deb12u3 |
| libbrotli-dev | 1.0.9-2+b6 |
| libbrotli1 | 1.0.9-2+b6 |
| libbsd0 | 0.11.7-2 |
| libbz2-1.0 | 1.0.8-5+b1 |
| libbz2-dev | 1.0.8-5+b1 |
| libc-bin | 2.36-9+deb12u13 |
| libc-dev-bin | 2.36-9+deb12u13 |
| libc6 | 2.36-9+deb12u13 |
| libc6-dev | 2.36-9+deb12u13 |
| libcairo-gobject2 | 1.16.0-7 |
| libcairo-script-interpreter2 | 1.16.0-7 |
| libcairo2 | 1.16.0-7 |
| libcairo2-dev | 1.16.0-7 |
| libcap-ng0 | 0.8.3-1+b3 |
| libcap2 | 1:2.66-4+deb12u2+b2 |
| libcbor0.8 | 0.8.0-2+b1 |
| libcc1-0 | 12.2.0-14+deb12u1 |
| libcom-err2 | 1.47.0-2+b2 |
| libcrypt-dev | 1:4.4.33-2 |
| libcrypt1 | 1:4.4.33-2 |
| libctf-nobfd0 | 2.40-2 |
| libctf0 | 2.40-2 |
| libcurl3-gnutls | 7.88.1-10+deb12u14 |
| libcurl4 | 7.88.1-10+deb12u14 |
| libcurl4-openssl-dev | 7.88.1-10+deb12u14 |
| libdatrie1 | 0.2.13-2+b1 |
| libdav1d6 | 1.0.0-2+deb12u1 |
| libdb-dev | 5.3.2 |
| libdb5.3 | 5.3.28+dfsg2-1 |
| libdb5.3-dev | 5.3.28+dfsg2-1 |
| libde265-0 | 1.0.11-1+deb12u2 |
| libdebconfclient0 | 0.270 |
| libdeflate-dev | 1.14-1 |
| libdeflate0 | 1.14-1 |
| libdjvulibre-dev | 3.5.28-2.2~deb12u1 |
| libdjvulibre-text | 3.5.28-2.2~deb12u1 |
| libdjvulibre21 | 3.5.28-2.2~deb12u1 |
| libdpkg-perl | 1.21.22 |
| libedit2 | 3.1-20221030-2 |
| libelf1 | 0.188-2.1 |
| liberror-perl | 0.17029-2 |
| libevent-2.1-7 | 2.1.12-stable-8 |
| libevent-core-2.1-7 | 2.1.12-stable-8 |
| libevent-dev | 2.1.12-stable-8 |
| libevent-extra-2.1-7 | 2.1.12-stable-8 |
| libevent-openssl-2.1-7 | 2.1.12-stable-8 |
| libevent-pthreads-2.1-7 | 2.1.12-stable-8 |
| libexif-dev | 0.6.24-1+b1 |
| libexif12 | 0.6.24-1+b1 |
| libexpat1 | 2.5.0-1+deb12u2 |
| libexpat1-dev | 2.5.0-1+deb12u2 |
| libext2fs2 | 1.47.0-2+b2 |
| libffi-dev | 3.4.4-1 |
| libffi8 | 3.4.4-1 |
| libfftw3-double3 | 3.3.10-1 |
| libfido2-1 | 1.12.0-2+b1 |
| libfontconfig-dev | 2.14.1-4 |
| libfontconfig1 | 2.14.1-4 |
| libfreetype-dev | 2.12.1+dfsg-5+deb12u4 |
| libfreetype6 | 2.12.1+dfsg-5+deb12u4 |
| libfreetype6-dev | 2.12.1+dfsg-5+deb12u4 |
| libfribidi0 | 1.0.8-2.1 |
| libgcc-12-dev | 12.2.0-14+deb12u1 |
| libgcc-s1 | 12.2.0-14+deb12u1 |
| libgcrypt20 | 1.10.1-3 |
| libgdbm-compat4 | 1.23-3 |
| libgdbm-dev | 1.23-3 |
| libgdbm6 | 1.23-3 |
| libgdk-pixbuf-2.0-0 | 2.42.10+dfsg-1+deb12u3 |
| libgdk-pixbuf-2.0-dev | 2.42.10+dfsg-1+deb12u3 |
| libgdk-pixbuf2.0-bin | 2.42.10+dfsg-1+deb12u3 |
| libgdk-pixbuf2.0-common | 2.42.10+dfsg-1+deb12u3 |
| libgirepository-1.0-1 | 1.74.0-3 |
| libglib2.0-0 | 2.74.6-2+deb12u8 |
| libglib2.0-bin | 2.74.6-2+deb12u8 |
| libglib2.0-data | 2.74.6-2+deb12u8 |
| libglib2.0-dev | 2.74.6-2+deb12u8 |
| libglib2.0-dev-bin | 2.74.6-2+deb12u8 |
| libgmp-dev | 2:6.2.1+dfsg1-1.1 |
| libgmp10 | 2:6.2.1+dfsg1-1.1 |
| libgmpxx4ldbl | 2:6.2.1+dfsg1-1.1 |
| libgnutls30 | 3.7.9-2+deb12u5 |
| libgomp1 | 12.2.0-14+deb12u1 |
| libgpg-error0 | 1.46-1 |
| libgprofng0 | 2.40-2 |
| libgraphite2-3 | 1.3.14-1 |
| libgssapi-krb5-2 | 1.20.1-2+deb12u4 |
| libgssrpc4 | 1.20.1-2+deb12u4 |
| libharfbuzz0b | 6.0.0+dfsg-3 |
| libheif1 | 1.15.1-1+deb12u1 |
| libhogweed6 | 3.8.1-2 |
| libhwasan0 | 12.2.0-14+deb12u1 |
| libice-dev | 2:1.0.10-1 |
| libice6 | 2:1.0.10-1 |
| libicu-dev | 72.1-3+deb12u1 |
| libicu72 | 72.1-3+deb12u1 |
| libidn2-0 | 2.3.3-1+b1 |
| libimath-3-1-29 | 3.1.6-1 |
| libimath-dev | 3.1.6-1 |
| libisl23 | 0.25-1.1 |
| libitm1 | 12.2.0-14+deb12u1 |
| libjansson4 | 2.14-2 |
| libjbig-dev | 2.1-6.1 |
| libjbig0 | 2.1-6.1 |
| libjpeg-dev | 1:2.1.5-2 |
| libjpeg62-turbo | 1:2.1.5-2 |
| libjpeg62-turbo-dev | 1:2.1.5-2 |
| libk5crypto3 | 1.20.1-2+deb12u4 |
| libkadm5clnt-mit12 | 1.20.1-2+deb12u4 |
| libkadm5srv-mit12 | 1.20.1-2+deb12u4 |
| libkdb5-10 | 1.20.1-2+deb12u4 |
| libkeyutils1 | 1.6.3-2 |
| libkrb5-3 | 1.20.1-2+deb12u4 |
| libkrb5-dev | 1.20.1-2+deb12u4 |
| libkrb5support0 | 1.20.1-2+deb12u4 |
| libksba8 | 1.6.3-2 |
| liblcms2-2 | 2.14-2 |
| liblcms2-dev | 2.14-2 |
| libldap-2.5-0 | 2.5.13+dfsg-5 |
| liblerc-dev | 4.0.0+ds-2 |
| liblerc4 | 4.0.0+ds-2 |
| liblqr-1-0 | 0.4.2-2.1 |
| liblqr-1-0-dev | 0.4.2-2.1 |
| liblsan0 | 12.2.0-14+deb12u1 |
| libltdl-dev | 2.4.7-7~deb12u1 |
| libltdl7 | 2.4.7-7~deb12u1 |
| liblz4-1 | 1.9.4-1 |
| liblzma-dev | 5.4.1-1 |
| liblzma5 | 5.4.1-1 |
| liblzo2-2 | 2.10-2 |
| libmagic-mgc | 1:5.44-3 |
| libmagic1 | 1:5.44-3 |
| libmagickcore-6-arch-config | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickcore-6-headers | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickcore-6.q16-6 | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickcore-6.q16-6-extra | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickcore-6.q16-dev | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickcore-dev | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickwand-6-headers | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickwand-6.q16-6 | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickwand-6.q16-dev | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickwand-dev | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmariadb-dev | 1:10.11.14-0+deb12u2 |
| libmariadb-dev-compat | 1:10.11.14-0+deb12u2 |
| libmariadb3 | 1:10.11.14-0+deb12u2 |
| libmaxminddb-dev | 1.7.1-1 |
| libmaxminddb0 | 1.7.1-1 |
| libmd0 | 1.0.4-2 |
| libmount-dev | 2.38.1-5+deb12u3 |
| libmount1 | 2.38.1-5+deb12u3 |
| libmpc3 | 1.3.1-1 |
| libmpfr6 | 4.2.0-1 |
| libncurses-dev | 6.4-4 |
| libncurses5-dev | 6.4-4 |
| libncurses6 | 6.4-4 |
| libncursesw5-dev | 6.4-4 |
| libncursesw6 | 6.4-4 |
| libnettle8 | 3.8.1-2 |
| libnghttp2-14 | 1.52.0-1+deb12u2 |
| libnpmaccess | 9.0.0 |
| libnpmdiff | 7.0.1 |
| libnpmexec | 9.0.1 |
| libnpmfund | 6.0.1 |
| libnpmhook | 11.0.0 |
| libnpmorg | 7.0.0 |
| libnpmpack | 8.0.1 |
| libnpmpublish | 10.0.1 |
| libnpmsearch | 8.0.0 |
| libnpmteam | 7.0.0 |
| libnpmversion | 7.0.0 |
| libnpth0 | 1.6-3 |
| libnsl-dev | 1.3.0-2 |
| libnsl2 | 1.3.0-2 |
| libnuma1 | 2.0.16-1 |
| libopenexr-3-1-30 | 3.1.5-5 |
| libopenexr-dev | 3.1.5-5 |
| libopenjp2-7 | 2.5.0-2+deb12u2 |
| libopenjp2-7-dev | 2.5.0-2+deb12u2 |
| libp11-kit0 | 0.24.1-2 |
| libpam-modules | 1.5.2-6+deb12u2 |
| libpam-modules-bin | 1.5.2-6+deb12u2 |
| libpam-runtime | 1.5.2-6+deb12u2 |
| libpam0g | 1.5.2-6+deb12u2 |
| libpango-1.0-0 | 1.50.12+ds-1 |
| libpangocairo-1.0-0 | 1.50.12+ds-1 |
| libpangoft2-1.0-0 | 1.50.12+ds-1 |
| libpcre2-16-0 | 10.42-1 |
| libpcre2-32-0 | 10.42-1 |
| libpcre2-8-0 | 10.42-1 |
| libpcre2-dev | 10.42-1 |
| libpcre2-posix3 | 10.42-1 |
| libperl5.36 | 5.36.0-7+deb12u3 |
| libpixman-1-0 | 0.42.2-1 |
| libpixman-1-dev | 0.42.2-1 |
| libpkgconf3 | 1.8.1-1 |
| libpng-dev | 1.6.39-2+deb12u1 |
| libpng16-16 | 1.6.39-2+deb12u1 |
| libpq-dev | 15.15-0+deb12u1 |
| libpq5 | 15.15-0+deb12u1 |
| libproc2-0 | 2:4.0.2-3 |
| libpsl5 | 0.21.2-1 |
| libpthread-stubs0-dev | 0.4-1 |
| libpython3-stdlib | 3.11.2-1+b1 |
| libpython3.11-minimal | 3.11.2-6+deb12u6 |
| libpython3.11-stdlib | 3.11.2-6+deb12u6 |
| libreadline-dev | 8.2-1.3 |
| libreadline8 | 8.2-1.3 |
| librsvg2-2 | 2.54.7+dfsg-1~deb12u1 |
| librsvg2-common | 2.54.7+dfsg-1~deb12u1 |
| librsvg2-dev | 2.54.7+dfsg-1~deb12u1 |
| librtmp1 | 2.4+20151223.gitfa8646d.1-2+b2 |
| libsasl2-2 | 2.1.28+dfsg-10 |
| libsasl2-modules-db | 2.1.28+dfsg-10 |
| libseccomp2 | 2.5.4-1+deb12u1 |
| libselinux1 | 3.4-1+b6 |
| libselinux1-dev | 3.4-1+b6 |
| libsemanage-common | 3.4-1 |
| libsemanage2 | 3.4-1+b5 |
| libsepol-dev | 3.4-2.1 |
| libsepol2 | 3.4-2.1 |
| libserf-1-1 | 1.3.9-11 |
| libsm-dev | 2:1.2.3-1 |
| libsm6 | 2:1.2.3-1 |
| libsmartcols1 | 2.38.1-5+deb12u3 |
| libsqlite3-0 | 3.40.1-2+deb12u2 |
| libsqlite3-dev | 3.40.1-2+deb12u2 |
| libss2 | 1.47.0-2+b2 |
| libssh2-1 | 1.10.0-3+b1 |
| libssl-dev | 3.0.18-1~deb12u2 |
| libssl3 | 3.0.18-1~deb12u2 |
| libstdc++-12-dev | 12.2.0-14+deb12u1 |
| libstdc++6 | 12.2.0-14+deb12u1 |
| libsvn1 | 1.14.2-4+deb12u1 |
| libsystemd0 | 252.39-1~deb12u1 |
| libtasn1-6 | 4.19.0-2+deb12u1 |
| libthai-data | 0.1.29-1 |
| libthai0 | 0.1.29-1 |
| libtiff-dev | 4.5.0-6+deb12u3 |
| libtiff6 | 4.5.0-6+deb12u3 |
| libtiffxx6 | 4.5.0-6+deb12u3 |
| libtinfo6 | 6.4-4 |
| libtirpc-common | 1.3.3+ds-1 |
| libtirpc-dev | 1.3.3+ds-1 |
| libtirpc3 | 1.3.3+ds-1 |
| libtool | 2.4.7-7~deb12u1 |
| libtsan2 | 12.2.0-14+deb12u1 |
| libubsan1 | 12.2.0-14+deb12u1 |
| libudev1 | 252.39-1~deb12u1 |
| libunistring2 | 1.0-2 |
| libutf8proc2 | 2.8.0-1 |
| libuuid1 | 2.38.1-5+deb12u3 |
| libwebp-dev | 1.2.4-0.2+deb12u1 |
| libwebp7 | 1.2.4-0.2+deb12u1 |
| libwebpdemux2 | 1.2.4-0.2+deb12u1 |
| libwebpmux3 | 1.2.4-0.2+deb12u1 |
| libwmf-0.2-7 | 0.2.12-5.1 |
| libwmf-dev | 0.2.12-5.1 |
| libwmflite-0.2-7 | 0.2.12-5.1 |
| libx11-6 | 2:1.8.4-2+deb12u2 |
| libx11-data | 2:1.8.4-2+deb12u2 |
| libx11-dev | 2:1.8.4-2+deb12u2 |
| libx265-199 | 3.5-2+b1 |
| libxau-dev | 1:1.0.9-1 |
| libxau6 | 1:1.0.9-1 |
| libxcb-render0 | 1.15-1 |
| libxcb-render0-dev | 1.15-1 |
| libxcb-shm0 | 1.15-1 |
| libxcb-shm0-dev | 1.15-1 |
| libxcb1 | 1.15-1 |
| libxcb1-dev | 1.15-1 |
| libxdmcp-dev | 1:1.1.2-3 |
| libxdmcp6 | 1:1.1.2-3 |
| libxext-dev | 2:1.3.4-1+b1 |
| libxext6 | 2:1.3.4-1+b1 |
| libxml2 | 2.9.14+dfsg-1.3~deb12u5 |
| libxml2-dev | 2.9.14+dfsg-1.3~deb12u5 |
| libxrender-dev | 1:0.9.10-1.1 |
| libxrender1 | 1:0.9.10-1.1 |
| libxslt1-dev | 1.1.35-1+deb12u3 |
| libxslt1.1 | 1.1.35-1+deb12u3 |
| libxt-dev | 1:1.2.1-1.1 |
| libxt6 | 1:1.2.1-1.1 |
| libxxhash0 | 0.8.1-1 |
| libyaml-0-2 | 0.2.5-1 |
| libyaml-dev | 0.2.5-1 |
| libzstd-dev | 1.5.4+dfsg2-5 |
| libzstd1 | 1.5.4+dfsg2-5 |
| linux-libc-dev | 6.1.159-1 |
| login | 1:4.13+dfsg1-1+deb12u2 |
| logsave | 1.47.0-2+b2 |
| lru-cache | 10.4.3 |
| m4 | 1.4.19-3 |
| make | 4.3-4.1 |
| make-fetch-happen | 14.0.3 |
| mariadb-common | 1:10.11.14-0+deb12u2 |
| mawk | 1.3.4.20200120-3.1 |
| media-types | 10.0.0 |
| mercurial | 6.3.2 |
| mercurial | 6.3.2-1+deb12u1 |
| mercurial-common | 6.3.2-1+deb12u1 |
| minimatch | 9.0.5 |
| minipass | 3.3.6 |
| minipass | 5.0.0 |
| minipass | 7.1.2 |
| minipass-collect | 2.0.1 |
| minipass-fetch | 4.0.1 |
| minipass-flush | 1.0.5 |
| minipass-pipeline | 1.2.4 |
| minipass-sized | 1.0.3 |
| minizlib | 2.1.2 |
| minizlib | 3.0.2 |
| mkdirp | 1.0.4 |
| mkdirp | 3.0.1 |
| mount | 2.38.1-5+deb12u3 |
| ms | 2.1.3 |
| mute-stream | 2.0.0 |
| mysql-common | 5.8+1.1.0 |
| ncurses-base | 6.4-4 |
| ncurses-bin | 6.4-4 |
| negotiator | 1.0.0 |
| netbase | 6.4 |
| node | 22.22.0 |
| node-gyp | 11.2.0 |
| nopt | 8.1.0 |
| normalize-package-data | 7.0.0 |
| npm | 10.9.4 |
| npm-audit-report | 6.0.0 |
| npm-bundled | 4.0.0 |
| npm-install-checks | 7.1.1 |
| npm-normalize-package-bin | 4.0.0 |
| npm-package-arg | 12.0.2 |
| npm-packlist | 9.0.0 |
| npm-pick-manifest | 10.0.0 |
| npm-profile | 11.0.1 |
| npm-registry-fetch | 18.0.2 |
| npm-user-validate | 3.0.0 |
| openssh-client | 1:9.2p1-2+deb12u7 |
| openssl | 3.0.18-1~deb12u2 |
| p-map | 7.0.3 |
| package-json-from-dist | 1.0.1 |
| pacote | 19.0.1 |
| pacote | 20.0.0 |
| parse-conflict-json | 4.0.0 |
| passwd | 1:4.13+dfsg1-1+deb12u2 |
| patch | 2.7.6-7 |
| path-key | 3.1.1 |
| path-scurry | 1.11.1 |
| perl | 5.36.0-7+deb12u3 |
| perl-base | 5.36.0-7+deb12u3 |
| perl-modules-5.36 | 5.36.0-7+deb12u3 |
| picomatch | 4.0.2 |
| pinentry-curses | 1.2.1-1 |
| pkg-config | 1.8.1-1 |
| pkgconf | 1.8.1-1 |
| pkgconf-bin | 1.8.1-1 |
| postcss-selector-parser | 7.1.0 |
| proc-log | 5.0.0 |
| procps | 2:4.0.2-3 |
| proggy | 3.0.0 |
| promise-all-reject-late | 1.0.1 |
| promise-call-limit | 3.0.2 |
| promise-retry | 2.0.1 |
| promzard | 2.0.0 |
| python3 | 3.11.2-1+b1 |
| python3-distutils | 3.11.2-3 |
| python3-lib2to3 | 3.11.2-3 |
| python3-minimal | 3.11.2-1+b1 |
| python3.11 | 3.11.2-6+deb12u6 |
| python3.11-minimal | 3.11.2-6+deb12u6 |
| qrcode-terminal | 0.12.0 |
| read | 4.1.0 |
| read-cmd-shim | 5.0.0 |
| read-package-json-fast | 4.0.0 |
| readline-common | 8.2-1.3 |
| retry | 0.12.0 |
| rpcsvc-proto | 1.4.3-1 |
| safer-buffer | 2.1.2 |
| sed | 4.9-1 |
| semver | 7.7.2 |
| sensible-utils | 0.0.17+nmu1 |
| shared-mime-info | 2.2-1 |
| shebang-command | 2.0.0 |
| shebang-regex | 3.0.0 |
| signal-exit | 4.1.0 |
| sigstore | 3.1.0 |
| smart-buffer | 4.2.0 |
| socks | 2.8.5 |
| socks-proxy-agent | 8.0.5 |
| spdx-correct | 3.2.0 |
| spdx-exceptions | 2.5.0 |
| spdx-expression-parse | 3.0.1 |
| spdx-expression-parse | 4.0.0 |
| spdx-license-ids | 3.0.21 |
| sprintf-js | 1.1.3 |
| sq | 0.27.0-2+b1 |
| ssri | 12.0.0 |
| string-width | 4.2.3 |
| string-width | 5.1.2 |
| strip-ansi | 6.0.1 |
| strip-ansi | 7.1.0 |
| subversion | 1.14.2-4+deb12u1 |
| supports-color | 9.4.0 |
| sysvinit-utils | 3.06-4 |
| tar | 1.34+dfsg-1.2+deb12u1 |
| tar | 6.2.1 |
| tar | 7.4.3 |
| text-table | 0.2.0 |
| tiny-relative-date | 1.3.0 |
| tinyglobby | 0.2.14 |
| treeverse | 3.0.0 |
| tuf-js | 3.0.1 |
| tzdata | 2025b-0+deb12u2 |
| ucf | 3.0043+nmu1+deb12u1 |
| unique-filename | 4.0.0 |
| unique-slug | 5.0.0 |
| unzip | 6.0-28 |
| usr-is-merged | 37~deb12u1 |
| util-deprecate | 1.0.2 |
| util-linux | 2.38.1-5+deb12u3 |
| util-linux-extra | 2.38.1-5+deb12u3 |
| uuid-dev | 2.38.1-5+deb12u3 |
| validate-npm-package-license | 3.0.4 |
| validate-npm-package-name | 6.0.1 |
| walk-up-path | 3.0.1 |
| wget | 1.21.3-1+deb12u1 |
| which | 2.0.2 |
| which | 5.0.0 |
| wrap-ansi | 7.0.0 |
| wrap-ansi | 8.1.0 |
| write-file-atomic | 6.0.0 |
| x11-common | 1:7.7+23 |
| x11proto-core-dev | 2022.1-1 |
| x11proto-dev | 2022.1-1 |
| xorg-sgml-doctools | 1:1.11-1.1 |
| xtrans-dev | 1.4.0-1 |
| xz-utils | 5.4.1-1 |
| yallist | 4.0.0 |
| yallist | 5.0.0 |
| yarn | 1.22.22 |
| zlib1g | 1:1.2.13.dfsg-1 |
| zlib1g-dev | 1:1.2.13.dfsg-1 |
</details>

</details>

<details>
<summary><strong>🔍 Full Report: node:24</strong></summary>

## 🛡️ Security & Efficiency

**Base Image:** `Debian GNU/Linux 12 (bookworm) (linux/arm64)`
**Supported Architectures:** `linux/amd64, linux/arm64, linux/ppc64le, linux/s390x, unknown/unknown`
**Efficiency Score:** 99.4%

### Vulnerabilities
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| 🛑 6 | 🟠 66 | 🟡 161 | 🔵 57 |

<details>
<summary><strong>👇 Expand Vulnerability Details (1020 found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [CVE-2023-45853](https://nvd.nist.gov/vuln/detail/CVE-2023-45853) | Critical | `zlib1g-dev` | `1:1.2.13.dfsg-1` |
| [CVE-2023-5841](https://nvd.nist.gov/vuln/detail/CVE-2023-5841) | Critical | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2023-5841](https://nvd.nist.gov/vuln/detail/CVE-2023-5841) | Critical | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2023-6879](https://nvd.nist.gov/vuln/detail/CVE-2023-6879) | Critical | `libaom3` | `3.6.0-1+deb12u2` |
| [CVE-2025-7458](https://nvd.nist.gov/vuln/detail/CVE-2025-7458) | Critical | `libsqlite3-dev` | `3.40.1-2+deb12u2` |
| [CVE-2025-7458](https://nvd.nist.gov/vuln/detail/CVE-2025-7458) | Critical | `libsqlite3-0` | `3.40.1-2+deb12u2` |
| [CVE-2023-25193](https://nvd.nist.gov/vuln/detail/CVE-2023-25193) | High | `libharfbuzz0b` | `6.0.0+dfsg-3` |
| [CVE-2023-2953](https://nvd.nist.gov/vuln/detail/CVE-2023-2953) | High | `libldap-2.5-0` | `2.5.13+dfsg-5` |
| [CVE-2023-39616](https://nvd.nist.gov/vuln/detail/CVE-2023-39616) | High | `libaom3` | `3.6.0-1+deb12u2` |
| [CVE-2023-52355](https://nvd.nist.gov/vuln/detail/CVE-2023-52355) | High | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2023-52355](https://nvd.nist.gov/vuln/detail/CVE-2023-52355) | High | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2023-52355](https://nvd.nist.gov/vuln/detail/CVE-2023-52355) | High | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2025-12495](https://nvd.nist.gov/vuln/detail/CVE-2025-12495) | High | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2025-12495](https://nvd.nist.gov/vuln/detail/CVE-2025-12495) | High | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2025-12839](https://nvd.nist.gov/vuln/detail/CVE-2025-12839) | High | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2025-12839](https://nvd.nist.gov/vuln/detail/CVE-2025-12839) | High | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2025-12840](https://nvd.nist.gov/vuln/detail/CVE-2025-12840) | High | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2025-12840](https://nvd.nist.gov/vuln/detail/CVE-2025-12840) | High | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2025-13151](https://nvd.nist.gov/vuln/detail/CVE-2025-13151) | High | `libtasn1-6` | `4.19.0-2+deb12u1` |
| [CVE-2025-13699](https://nvd.nist.gov/vuln/detail/CVE-2025-13699) | High | `libmariadb-dev-compat` | `1:10.11.14-0+deb12u2` |
| [CVE-2025-13699](https://nvd.nist.gov/vuln/detail/CVE-2025-13699) | High | `libmariadb-dev` | `1:10.11.14-0+deb12u2` |
| [CVE-2025-13699](https://nvd.nist.gov/vuln/detail/CVE-2025-13699) | High | `libmariadb3` | `1:10.11.14-0+deb12u2` |
| [CVE-2025-13699](https://nvd.nist.gov/vuln/detail/CVE-2025-13699) | High | `mariadb-common` | `1:10.11.14-0+deb12u2` |
| [CVE-2025-13836](https://nvd.nist.gov/vuln/detail/CVE-2025-13836) | High | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-13836](https://nvd.nist.gov/vuln/detail/CVE-2025-13836) | High | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-13836](https://nvd.nist.gov/vuln/detail/CVE-2025-13836) | High | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-13836](https://nvd.nist.gov/vuln/detail/CVE-2025-13836) | High | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-15281](https://nvd.nist.gov/vuln/detail/CVE-2025-15281) | High | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2025-15281](https://nvd.nist.gov/vuln/detail/CVE-2025-15281) | High | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2025-15281](https://nvd.nist.gov/vuln/detail/CVE-2025-15281) | High | `libc6` | `2.36-9+deb12u13` |
| [CVE-2025-15281](https://nvd.nist.gov/vuln/detail/CVE-2025-15281) | High | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2025-46835](https://nvd.nist.gov/vuln/detail/CVE-2025-46835) | High | `git-man` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-46835](https://nvd.nist.gov/vuln/detail/CVE-2025-46835) | High | `git` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-48384](https://nvd.nist.gov/vuln/detail/CVE-2025-48384) | High | `git-man` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-48384](https://nvd.nist.gov/vuln/detail/CVE-2025-48384) | High | `git` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-48385](https://nvd.nist.gov/vuln/detail/CVE-2025-48385) | High | `git` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-48385](https://nvd.nist.gov/vuln/detail/CVE-2025-48385) | High | `git-man` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-59375](https://nvd.nist.gov/vuln/detail/CVE-2025-59375) | High | `libexpat1` | `2.5.0-1+deb12u2` |
| [CVE-2025-59375](https://nvd.nist.gov/vuln/detail/CVE-2025-59375) | High | `libexpat1-dev` | `2.5.0-1+deb12u2` |
| [CVE-2025-6297](https://nvd.nist.gov/vuln/detail/CVE-2025-6297) | High | `libdpkg-perl` | `1.21.22` |
| [CVE-2025-6297](https://nvd.nist.gov/vuln/detail/CVE-2025-6297) | High | `dpkg` | `1.21.22` |
| [CVE-2025-6297](https://nvd.nist.gov/vuln/detail/CVE-2025-6297) | High | `dpkg-dev` | `1.21.22` |
| [CVE-2025-64181](https://nvd.nist.gov/vuln/detail/CVE-2025-64181) | High | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2025-64181](https://nvd.nist.gov/vuln/detail/CVE-2025-64181) | High | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2025-7425](https://nvd.nist.gov/vuln/detail/CVE-2025-7425) | High | `libxslt1-dev` | `1.1.35-1+deb12u3` |
| [CVE-2025-7425](https://nvd.nist.gov/vuln/detail/CVE-2025-7425) | High | `libxslt1.1` | `1.1.35-1+deb12u3` |
| [CVE-2025-8194](https://nvd.nist.gov/vuln/detail/CVE-2025-8194) | High | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-8194](https://nvd.nist.gov/vuln/detail/CVE-2025-8194) | High | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-8194](https://nvd.nist.gov/vuln/detail/CVE-2025-8194) | High | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-8194](https://nvd.nist.gov/vuln/detail/CVE-2025-8194) | High | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2026-0861](https://nvd.nist.gov/vuln/detail/CVE-2026-0861) | High | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2026-0861](https://nvd.nist.gov/vuln/detail/CVE-2026-0861) | High | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2026-0861](https://nvd.nist.gov/vuln/detail/CVE-2026-0861) | High | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2026-0861](https://nvd.nist.gov/vuln/detail/CVE-2026-0861) | High | `libc6` | `2.36-9+deb12u13` |
| [CVE-2026-0915](https://nvd.nist.gov/vuln/detail/CVE-2026-0915) | High | `libc6` | `2.36-9+deb12u13` |
| [CVE-2026-0915](https://nvd.nist.gov/vuln/detail/CVE-2026-0915) | High | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2026-0915](https://nvd.nist.gov/vuln/detail/CVE-2026-0915) | High | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2026-0915](https://nvd.nist.gov/vuln/detail/CVE-2026-0915) | High | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2026-2004](https://nvd.nist.gov/vuln/detail/CVE-2026-2004) | High | `libpq5` | `15.15-0+deb12u1` |
| [CVE-2026-2004](https://nvd.nist.gov/vuln/detail/CVE-2026-2004) | High | `libpq-dev` | `15.15-0+deb12u1` |
| [CVE-2026-2005](https://nvd.nist.gov/vuln/detail/CVE-2026-2005) | High | `libpq-dev` | `15.15-0+deb12u1` |
| [CVE-2026-2005](https://nvd.nist.gov/vuln/detail/CVE-2026-2005) | High | `libpq5` | `15.15-0+deb12u1` |
| [CVE-2026-2006](https://nvd.nist.gov/vuln/detail/CVE-2026-2006) | High | `libpq5` | `15.15-0+deb12u1` |
| [CVE-2026-2006](https://nvd.nist.gov/vuln/detail/CVE-2026-2006) | High | `libpq-dev` | `15.15-0+deb12u1` |
| [CVE-2026-22695](https://nvd.nist.gov/vuln/detail/CVE-2026-22695) | High | `libpng16-16` | `1.6.39-2+deb12u1` |
| [CVE-2026-22695](https://nvd.nist.gov/vuln/detail/CVE-2026-22695) | High | `libpng-dev` | `1.6.39-2+deb12u1` |
| [CVE-2026-22801](https://nvd.nist.gov/vuln/detail/CVE-2026-22801) | High | `libpng16-16` | `1.6.39-2+deb12u1` |
| [CVE-2026-22801](https://nvd.nist.gov/vuln/detail/CVE-2026-22801) | High | `libpng-dev` | `1.6.39-2+deb12u1` |
| [CVE-2026-25646](https://nvd.nist.gov/vuln/detail/CVE-2026-25646) | High | `libpng16-16` | `1.6.39-2+deb12u1` |
| [CVE-2026-25646](https://nvd.nist.gov/vuln/detail/CVE-2026-25646) | High | `libpng-dev` | `1.6.39-2+deb12u1` |
| [GHSA-34x7-hfp2-rc4v](https://nvd.nist.gov/vuln/detail/GHSA-34x7-hfp2-rc4v) | High | `tar` | `7.5.4` |
| [GHSA-7h2j-956f-4vf2](https://nvd.nist.gov/vuln/detail/GHSA-7h2j-956f-4vf2) | High | `@isaacs/brace-expansion` | `5.0.0` |
| [CVE-2007-3996](https://nvd.nist.gov/vuln/detail/CVE-2007-3996) | Medium | `libwmf-dev` | `0.2.12-5.1` |
| [CVE-2007-3996](https://nvd.nist.gov/vuln/detail/CVE-2007-3996) | Medium | `libwmf-0.2-7` | `0.2.12-5.1` |
| [CVE-2007-3996](https://nvd.nist.gov/vuln/detail/CVE-2007-3996) | Medium | `libwmflite-0.2-7` | `0.2.12-5.1` |
| [CVE-2009-3546](https://nvd.nist.gov/vuln/detail/CVE-2009-3546) | Medium | `libwmf-0.2-7` | `0.2.12-5.1` |
| [CVE-2009-3546](https://nvd.nist.gov/vuln/detail/CVE-2009-3546) | Medium | `libwmf-dev` | `0.2.12-5.1` |
| [CVE-2009-3546](https://nvd.nist.gov/vuln/detail/CVE-2009-3546) | Medium | `libwmflite-0.2-7` | `0.2.12-5.1` |
| [CVE-2021-31879](https://nvd.nist.gov/vuln/detail/CVE-2021-31879) | Medium | `wget` | `1.21.3-1+deb12u1` |
| [CVE-2023-32570](https://nvd.nist.gov/vuln/detail/CVE-2023-32570) | Medium | `libdav1d6` | `1.0.0-2+deb12u1` |
| [CVE-2023-39327](https://nvd.nist.gov/vuln/detail/CVE-2023-39327) | Medium | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2023-39327](https://nvd.nist.gov/vuln/detail/CVE-2023-39327) | Medium | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2023-39328](https://nvd.nist.gov/vuln/detail/CVE-2023-39328) | Medium | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2023-39328](https://nvd.nist.gov/vuln/detail/CVE-2023-39328) | Medium | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2023-39329](https://nvd.nist.gov/vuln/detail/CVE-2023-39329) | Medium | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2023-39329](https://nvd.nist.gov/vuln/detail/CVE-2023-39329) | Medium | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `libncursesw6` | `6.4-4` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `libncurses6` | `6.4-4` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `ncurses-base` | `6.4-4` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `libtinfo6` | `6.4-4` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `libncurses5-dev` | `6.4-4` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `libncurses-dev` | `6.4-4` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `libncursesw5-dev` | `6.4-4` |
| [CVE-2023-50495](https://nvd.nist.gov/vuln/detail/CVE-2023-50495) | Medium | `ncurses-bin` | `6.4-4` |
| [CVE-2023-6277](https://nvd.nist.gov/vuln/detail/CVE-2023-6277) | Medium | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2023-6277](https://nvd.nist.gov/vuln/detail/CVE-2023-6277) | Medium | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2023-6277](https://nvd.nist.gov/vuln/detail/CVE-2023-6277) | Medium | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2024-10041](https://nvd.nist.gov/vuln/detail/CVE-2024-10041) | Medium | `libpam-modules-bin` | `1.5.2-6+deb12u2` |
| [CVE-2024-10041](https://nvd.nist.gov/vuln/detail/CVE-2024-10041) | Medium | `libpam-runtime` | `1.5.2-6+deb12u2` |
| [CVE-2024-10041](https://nvd.nist.gov/vuln/detail/CVE-2024-10041) | Medium | `libpam-modules` | `1.5.2-6+deb12u2` |
| [CVE-2024-10041](https://nvd.nist.gov/vuln/detail/CVE-2024-10041) | Medium | `libpam0g` | `1.5.2-6+deb12u2` |
| [CVE-2024-10524](https://nvd.nist.gov/vuln/detail/CVE-2024-10524) | Medium | `wget` | `1.21.3-1+deb12u1` |
| [CVE-2024-38949](https://nvd.nist.gov/vuln/detail/CVE-2024-38949) | Medium | `libde265-0` | `1.0.11-1+deb12u2` |
| [CVE-2024-38950](https://nvd.nist.gov/vuln/detail/CVE-2024-38950) | Medium | `libde265-0` | `1.0.11-1+deb12u2` |
| [CVE-2025-10148](https://nvd.nist.gov/vuln/detail/CVE-2025-10148) | Medium | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-10148](https://nvd.nist.gov/vuln/detail/CVE-2025-10148) | Medium | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-10148](https://nvd.nist.gov/vuln/detail/CVE-2025-10148) | Medium | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-10148](https://nvd.nist.gov/vuln/detail/CVE-2025-10148) | Medium | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-10911](https://nvd.nist.gov/vuln/detail/CVE-2025-10911) | Medium | `libxslt1.1` | `1.1.35-1+deb12u3` |
| [CVE-2025-10911](https://nvd.nist.gov/vuln/detail/CVE-2025-10911) | Medium | `libxslt1-dev` | `1.1.35-1+deb12u3` |
| [CVE-2025-12084](https://nvd.nist.gov/vuln/detail/CVE-2025-12084) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-12084](https://nvd.nist.gov/vuln/detail/CVE-2025-12084) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-12084](https://nvd.nist.gov/vuln/detail/CVE-2025-12084) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-12084](https://nvd.nist.gov/vuln/detail/CVE-2025-12084) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-12781](https://nvd.nist.gov/vuln/detail/CVE-2025-12781) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-12781](https://nvd.nist.gov/vuln/detail/CVE-2025-12781) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-12781](https://nvd.nist.gov/vuln/detail/CVE-2025-12781) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-12781](https://nvd.nist.gov/vuln/detail/CVE-2025-12781) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-13837](https://nvd.nist.gov/vuln/detail/CVE-2025-13837) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-13837](https://nvd.nist.gov/vuln/detail/CVE-2025-13837) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-13837](https://nvd.nist.gov/vuln/detail/CVE-2025-13837) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-13837](https://nvd.nist.gov/vuln/detail/CVE-2025-13837) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `bsdutils` | `1:2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libuuid1` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libmount1` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `util-linux` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `util-linux-extra` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `uuid-dev` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libmount-dev` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libblkid1` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libblkid-dev` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `mount` | `2.38.1-5+deb12u3` |
| [CVE-2025-14104](https://nvd.nist.gov/vuln/detail/CVE-2025-14104) | Medium | `libsmartcols1` | `2.38.1-5+deb12u3` |
| [CVE-2025-14524](https://nvd.nist.gov/vuln/detail/CVE-2025-14524) | Medium | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-14524](https://nvd.nist.gov/vuln/detail/CVE-2025-14524) | Medium | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-14524](https://nvd.nist.gov/vuln/detail/CVE-2025-14524) | Medium | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-14524](https://nvd.nist.gov/vuln/detail/CVE-2025-14524) | Medium | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-14819](https://nvd.nist.gov/vuln/detail/CVE-2025-14819) | Medium | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-14819](https://nvd.nist.gov/vuln/detail/CVE-2025-14819) | Medium | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-14819](https://nvd.nist.gov/vuln/detail/CVE-2025-14819) | Medium | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-14819](https://nvd.nist.gov/vuln/detail/CVE-2025-14819) | Medium | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-14831](https://nvd.nist.gov/vuln/detail/CVE-2025-14831) | Medium | `libgnutls30` | `3.7.9-2+deb12u5` |
| [CVE-2025-15282](https://nvd.nist.gov/vuln/detail/CVE-2025-15282) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-15282](https://nvd.nist.gov/vuln/detail/CVE-2025-15282) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-15282](https://nvd.nist.gov/vuln/detail/CVE-2025-15282) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-15282](https://nvd.nist.gov/vuln/detail/CVE-2025-15282) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-15366](https://nvd.nist.gov/vuln/detail/CVE-2025-15366) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-15366](https://nvd.nist.gov/vuln/detail/CVE-2025-15366) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-15366](https://nvd.nist.gov/vuln/detail/CVE-2025-15366) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-15366](https://nvd.nist.gov/vuln/detail/CVE-2025-15366) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-15367](https://nvd.nist.gov/vuln/detail/CVE-2025-15367) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-15367](https://nvd.nist.gov/vuln/detail/CVE-2025-15367) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-15367](https://nvd.nist.gov/vuln/detail/CVE-2025-15367) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-15367](https://nvd.nist.gov/vuln/detail/CVE-2025-15367) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gpg-wks-client` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gnupg-utils` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gpgconf` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gpgsm` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `dirmngr` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gnupg` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gnupg-l10n` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gpgv` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gpg-wks-server` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gpg-agent` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-30258](https://nvd.nist.gov/vuln/detail/CVE-2025-30258) | Medium | `gpg` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-48074](https://nvd.nist.gov/vuln/detail/CVE-2025-48074) | Medium | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2025-48074](https://nvd.nist.gov/vuln/detail/CVE-2025-48074) | Medium | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2025-6069](https://nvd.nist.gov/vuln/detail/CVE-2025-6069) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-6069](https://nvd.nist.gov/vuln/detail/CVE-2025-6069) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-6069](https://nvd.nist.gov/vuln/detail/CVE-2025-6069) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-6069](https://nvd.nist.gov/vuln/detail/CVE-2025-6069) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-6075](https://nvd.nist.gov/vuln/detail/CVE-2025-6075) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-6075](https://nvd.nist.gov/vuln/detail/CVE-2025-6075) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-6075](https://nvd.nist.gov/vuln/detail/CVE-2025-6075) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-6075](https://nvd.nist.gov/vuln/detail/CVE-2025-6075) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libncurses5-dev` | `6.4-4` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libncurses-dev` | `6.4-4` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libncursesw5-dev` | `6.4-4` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `ncurses-bin` | `6.4-4` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libtinfo6` | `6.4-4` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libncursesw6` | `6.4-4` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `libncurses6` | `6.4-4` |
| [CVE-2025-6141](https://nvd.nist.gov/vuln/detail/CVE-2025-6141) | Medium | `ncurses-base` | `6.4-4` |
| [CVE-2025-66382](https://nvd.nist.gov/vuln/detail/CVE-2025-66382) | Medium | `libexpat1-dev` | `2.5.0-1+deb12u2` |
| [CVE-2025-66382](https://nvd.nist.gov/vuln/detail/CVE-2025-66382) | Medium | `libexpat1` | `2.5.0-1+deb12u2` |
| [CVE-2025-68431](https://nvd.nist.gov/vuln/detail/CVE-2025-68431) | Medium | `libheif1` | `1.15.1-1+deb12u1` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gpgconf` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gnupg-l10n` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `dirmngr` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gnupg-utils` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gnupg` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gpgv` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gpgsm` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gpg` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gpg-wks-server` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gpg-agent` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-68972](https://nvd.nist.gov/vuln/detail/CVE-2025-68972) | Medium | `gpg-wks-client` | `2.2.40-1.1+deb12u2` |
| [CVE-2025-7709](https://nvd.nist.gov/vuln/detail/CVE-2025-7709) | Medium | `libsqlite3-dev` | `3.40.1-2+deb12u2` |
| [CVE-2025-7709](https://nvd.nist.gov/vuln/detail/CVE-2025-7709) | Medium | `libsqlite3-0` | `3.40.1-2+deb12u2` |
| [CVE-2025-8291](https://nvd.nist.gov/vuln/detail/CVE-2025-8291) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-8291](https://nvd.nist.gov/vuln/detail/CVE-2025-8291) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-8291](https://nvd.nist.gov/vuln/detail/CVE-2025-8291) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-8291](https://nvd.nist.gov/vuln/detail/CVE-2025-8291) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-9820](https://nvd.nist.gov/vuln/detail/CVE-2025-9820) | Medium | `libgnutls30` | `3.7.9-2+deb12u5` |
| [CVE-2026-0672](https://nvd.nist.gov/vuln/detail/CVE-2026-0672) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2026-0672](https://nvd.nist.gov/vuln/detail/CVE-2026-0672) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2026-0672](https://nvd.nist.gov/vuln/detail/CVE-2026-0672) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2026-0672](https://nvd.nist.gov/vuln/detail/CVE-2026-0672) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2026-0865](https://nvd.nist.gov/vuln/detail/CVE-2026-0865) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2026-0865](https://nvd.nist.gov/vuln/detail/CVE-2026-0865) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2026-0865](https://nvd.nist.gov/vuln/detail/CVE-2026-0865) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2026-0865](https://nvd.nist.gov/vuln/detail/CVE-2026-0865) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2026-0990](https://nvd.nist.gov/vuln/detail/CVE-2026-0990) | Medium | `libxml2` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-0990](https://nvd.nist.gov/vuln/detail/CVE-2026-0990) | Medium | `libxml2-dev` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-1299](https://nvd.nist.gov/vuln/detail/CVE-2026-1299) | Medium | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2026-1299](https://nvd.nist.gov/vuln/detail/CVE-2026-1299) | Medium | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2026-1299](https://nvd.nist.gov/vuln/detail/CVE-2026-1299) | Medium | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2026-1299](https://nvd.nist.gov/vuln/detail/CVE-2026-1299) | Medium | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2026-1484](https://nvd.nist.gov/vuln/detail/CVE-2026-1484) | Medium | `libglib2.0-dev` | `2.74.6-2+deb12u8` |
| [CVE-2026-1484](https://nvd.nist.gov/vuln/detail/CVE-2026-1484) | Medium | `libglib2.0-0` | `2.74.6-2+deb12u8` |
| [CVE-2026-1484](https://nvd.nist.gov/vuln/detail/CVE-2026-1484) | Medium | `libglib2.0-dev-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-1484](https://nvd.nist.gov/vuln/detail/CVE-2026-1484) | Medium | `libglib2.0-data` | `2.74.6-2+deb12u8` |
| [CVE-2026-1484](https://nvd.nist.gov/vuln/detail/CVE-2026-1484) | Medium | `libglib2.0-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-1489](https://nvd.nist.gov/vuln/detail/CVE-2026-1489) | Medium | `libglib2.0-dev` | `2.74.6-2+deb12u8` |
| [CVE-2026-1489](https://nvd.nist.gov/vuln/detail/CVE-2026-1489) | Medium | `libglib2.0-dev-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-1489](https://nvd.nist.gov/vuln/detail/CVE-2026-1489) | Medium | `libglib2.0-data` | `2.74.6-2+deb12u8` |
| [CVE-2026-1489](https://nvd.nist.gov/vuln/detail/CVE-2026-1489) | Medium | `libglib2.0-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-1489](https://nvd.nist.gov/vuln/detail/CVE-2026-1489) | Medium | `libglib2.0-0` | `2.74.6-2+deb12u8` |
| [CVE-2026-2003](https://nvd.nist.gov/vuln/detail/CVE-2026-2003) | Medium | `libpq5` | `15.15-0+deb12u1` |
| [CVE-2026-2003](https://nvd.nist.gov/vuln/detail/CVE-2026-2003) | Medium | `libpq-dev` | `15.15-0+deb12u1` |
| [CVE-2026-22693](https://nvd.nist.gov/vuln/detail/CVE-2026-22693) | Medium | `libharfbuzz0b` | `6.0.0+dfsg-3` |
| [CVE-2026-25210](https://nvd.nist.gov/vuln/detail/CVE-2026-25210) | Medium | `libexpat1-dev` | `2.5.0-1+deb12u2` |
| [CVE-2026-25210](https://nvd.nist.gov/vuln/detail/CVE-2026-25210) | Medium | `libexpat1` | `2.5.0-1+deb12u2` |
| [CVE-2007-3476](https://nvd.nist.gov/vuln/detail/CVE-2007-3476) | Low | `libwmf-dev` | `0.2.12-5.1` |
| [CVE-2007-3476](https://nvd.nist.gov/vuln/detail/CVE-2007-3476) | Low | `libwmflite-0.2-7` | `0.2.12-5.1` |
| [CVE-2007-3476](https://nvd.nist.gov/vuln/detail/CVE-2007-3476) | Low | `libwmf-0.2-7` | `0.2.12-5.1` |
| [CVE-2007-3477](https://nvd.nist.gov/vuln/detail/CVE-2007-3477) | Low | `libwmflite-0.2-7` | `0.2.12-5.1` |
| [CVE-2007-3477](https://nvd.nist.gov/vuln/detail/CVE-2007-3477) | Low | `libwmf-dev` | `0.2.12-5.1` |
| [CVE-2007-3477](https://nvd.nist.gov/vuln/detail/CVE-2007-3477) | Low | `libwmf-0.2-7` | `0.2.12-5.1` |
| [CVE-2016-2781](https://nvd.nist.gov/vuln/detail/CVE-2016-2781) | Low | `coreutils` | `9.1-1` |
| [CVE-2017-7475](https://nvd.nist.gov/vuln/detail/CVE-2017-7475) | Low | `libcairo2` | `1.16.0-7` |
| [CVE-2017-7475](https://nvd.nist.gov/vuln/detail/CVE-2017-7475) | Low | `libcairo2-dev` | `1.16.0-7` |
| [CVE-2017-7475](https://nvd.nist.gov/vuln/detail/CVE-2017-7475) | Low | `libcairo-gobject2` | `1.16.0-7` |
| [CVE-2017-7475](https://nvd.nist.gov/vuln/detail/CVE-2017-7475) | Low | `libcairo-script-interpreter2` | `1.16.0-7` |
| [CVE-2019-6461](https://nvd.nist.gov/vuln/detail/CVE-2019-6461) | Low | `libcairo2` | `1.16.0-7` |
| [CVE-2019-6461](https://nvd.nist.gov/vuln/detail/CVE-2019-6461) | Low | `libcairo-script-interpreter2` | `1.16.0-7` |
| [CVE-2019-6461](https://nvd.nist.gov/vuln/detail/CVE-2019-6461) | Low | `libcairo-gobject2` | `1.16.0-7` |
| [CVE-2019-6461](https://nvd.nist.gov/vuln/detail/CVE-2019-6461) | Low | `libcairo2-dev` | `1.16.0-7` |
| [CVE-2019-6462](https://nvd.nist.gov/vuln/detail/CVE-2019-6462) | Low | `libcairo2` | `1.16.0-7` |
| [CVE-2019-6462](https://nvd.nist.gov/vuln/detail/CVE-2019-6462) | Low | `libcairo-script-interpreter2` | `1.16.0-7` |
| [CVE-2019-6462](https://nvd.nist.gov/vuln/detail/CVE-2019-6462) | Low | `libcairo-gobject2` | `1.16.0-7` |
| [CVE-2019-6462](https://nvd.nist.gov/vuln/detail/CVE-2019-6462) | Low | `libcairo2-dev` | `1.16.0-7` |
| [CVE-2019-6988](https://nvd.nist.gov/vuln/detail/CVE-2019-6988) | Low | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2019-6988](https://nvd.nist.gov/vuln/detail/CVE-2019-6988) | Low | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2023-4016](https://nvd.nist.gov/vuln/detail/CVE-2023-4016) | Low | `procps` | `2:4.0.2-3` |
| [CVE-2023-4016](https://nvd.nist.gov/vuln/detail/CVE-2023-4016) | Low | `libproc2-0` | `2:4.0.2-3` |
| [CVE-2023-51792](https://nvd.nist.gov/vuln/detail/CVE-2023-51792) | Low | `libde265-0` | `1.0.11-1+deb12u2` |
| [CVE-2024-13978](https://nvd.nist.gov/vuln/detail/CVE-2024-13978) | Low | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2024-13978](https://nvd.nist.gov/vuln/detail/CVE-2024-13978) | Low | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2024-13978](https://nvd.nist.gov/vuln/detail/CVE-2024-13978) | Low | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2024-31047](https://nvd.nist.gov/vuln/detail/CVE-2024-31047) | Low | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2024-31047](https://nvd.nist.gov/vuln/detail/CVE-2024-31047) | Low | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2024-56433](https://nvd.nist.gov/vuln/detail/CVE-2024-56433) | Low | `passwd` | `1:4.13+dfsg1-1+deb12u2` |
| [CVE-2024-56433](https://nvd.nist.gov/vuln/detail/CVE-2024-56433) | Low | `login` | `1:4.13+dfsg1-1+deb12u2` |
| [CVE-2025-11731](https://nvd.nist.gov/vuln/detail/CVE-2025-11731) | Low | `libxslt1-dev` | `1.1.35-1+deb12u3` |
| [CVE-2025-11731](https://nvd.nist.gov/vuln/detail/CVE-2025-11731) | Low | `libxslt1.1` | `1.1.35-1+deb12u3` |
| [CVE-2025-27613](https://nvd.nist.gov/vuln/detail/CVE-2025-27613) | Low | `git` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-27613](https://nvd.nist.gov/vuln/detail/CVE-2025-27613) | Low | `git-man` | `1:2.39.5-0+deb12u3` |
| [CVE-2025-50422](https://nvd.nist.gov/vuln/detail/CVE-2025-50422) | Low | `libcairo-gobject2` | `1.16.0-7` |
| [CVE-2025-50422](https://nvd.nist.gov/vuln/detail/CVE-2025-50422) | Low | `libcairo-script-interpreter2` | `1.16.0-7` |
| [CVE-2025-50422](https://nvd.nist.gov/vuln/detail/CVE-2025-50422) | Low | `libcairo2` | `1.16.0-7` |
| [CVE-2025-50422](https://nvd.nist.gov/vuln/detail/CVE-2025-50422) | Low | `libcairo2-dev` | `1.16.0-7` |
| [CVE-2025-61984](https://nvd.nist.gov/vuln/detail/CVE-2025-61984) | Low | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2025-61985](https://nvd.nist.gov/vuln/detail/CVE-2025-61985) | Low | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2026-0988](https://nvd.nist.gov/vuln/detail/CVE-2026-0988) | Low | `libglib2.0-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-0988](https://nvd.nist.gov/vuln/detail/CVE-2026-0988) | Low | `libglib2.0-dev-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-0988](https://nvd.nist.gov/vuln/detail/CVE-2026-0988) | Low | `libglib2.0-0` | `2.74.6-2+deb12u8` |
| [CVE-2026-0988](https://nvd.nist.gov/vuln/detail/CVE-2026-0988) | Low | `libglib2.0-dev` | `2.74.6-2+deb12u8` |
| [CVE-2026-0988](https://nvd.nist.gov/vuln/detail/CVE-2026-0988) | Low | `libglib2.0-data` | `2.74.6-2+deb12u8` |
| [CVE-2026-0989](https://nvd.nist.gov/vuln/detail/CVE-2026-0989) | Low | `libxml2-dev` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-0989](https://nvd.nist.gov/vuln/detail/CVE-2026-0989) | Low | `libxml2` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-0992](https://nvd.nist.gov/vuln/detail/CVE-2026-0992) | Low | `libxml2` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-0992](https://nvd.nist.gov/vuln/detail/CVE-2026-0992) | Low | `libxml2-dev` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-1485](https://nvd.nist.gov/vuln/detail/CVE-2026-1485) | Low | `libglib2.0-0` | `2.74.6-2+deb12u8` |
| [CVE-2026-1485](https://nvd.nist.gov/vuln/detail/CVE-2026-1485) | Low | `libglib2.0-dev-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-1485](https://nvd.nist.gov/vuln/detail/CVE-2026-1485) | Low | `libglib2.0-dev` | `2.74.6-2+deb12u8` |
| [CVE-2026-1485](https://nvd.nist.gov/vuln/detail/CVE-2026-1485) | Low | `libglib2.0-data` | `2.74.6-2+deb12u8` |
| [CVE-2026-1485](https://nvd.nist.gov/vuln/detail/CVE-2026-1485) | Low | `libglib2.0-bin` | `2.74.6-2+deb12u8` |
| [CVE-2026-24515](https://nvd.nist.gov/vuln/detail/CVE-2026-24515) | Low | `libexpat1-dev` | `2.5.0-1+deb12u2` |
| [CVE-2026-24515](https://nvd.nist.gov/vuln/detail/CVE-2026-24515) | Low | `libexpat1` | `2.5.0-1+deb12u2` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-0406](https://nvd.nist.gov/vuln/detail/CVE-2005-0406) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2005-2541](https://nvd.nist.gov/vuln/detail/CVE-2005-2541) | Negligible | `tar` | `1.34+dfsg-1.2+deb12u1` |
| [CVE-2007-2243](https://nvd.nist.gov/vuln/detail/CVE-2007-2243) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2007-2768](https://nvd.nist.gov/vuln/detail/CVE-2007-2768) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2007-5686](https://nvd.nist.gov/vuln/detail/CVE-2007-5686) | Negligible | `passwd` | `1:4.13+dfsg1-1+deb12u2` |
| [CVE-2007-5686](https://nvd.nist.gov/vuln/detail/CVE-2007-5686) | Negligible | `login` | `1:4.13+dfsg1-1+deb12u2` |
| [CVE-2008-1687](https://nvd.nist.gov/vuln/detail/CVE-2008-1687) | Negligible | `m4` | `1.4.19-3` |
| [CVE-2008-1688](https://nvd.nist.gov/vuln/detail/CVE-2008-1688) | Negligible | `m4` | `1.4.19-3` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3134](https://nvd.nist.gov/vuln/detail/CVE-2008-3134) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2008-3234](https://nvd.nist.gov/vuln/detail/CVE-2008-3234) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2010-4651](https://nvd.nist.gov/vuln/detail/CVE-2010-4651) | Negligible | `patch` | `2.7.6-7` |
| [CVE-2010-4756](https://nvd.nist.gov/vuln/detail/CVE-2010-4756) | Negligible | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2010-4756](https://nvd.nist.gov/vuln/detail/CVE-2010-4756) | Negligible | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2010-4756](https://nvd.nist.gov/vuln/detail/CVE-2010-4756) | Negligible | `libc6` | `2.36-9+deb12u13` |
| [CVE-2010-4756](https://nvd.nist.gov/vuln/detail/CVE-2010-4756) | Negligible | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2011-3374](https://nvd.nist.gov/vuln/detail/CVE-2011-3374) | Negligible | `libapt-pkg6.0` | `2.6.1` |
| [CVE-2011-3374](https://nvd.nist.gov/vuln/detail/CVE-2011-3374) | Negligible | `apt` | `2.6.1` |
| [CVE-2011-3389](https://nvd.nist.gov/vuln/detail/CVE-2011-3389) | Negligible | `libgnutls30` | `3.7.9-2+deb12u5` |
| [CVE-2011-4116](https://nvd.nist.gov/vuln/detail/CVE-2011-4116) | Negligible | `libperl5.36` | `5.36.0-7+deb12u3` |
| [CVE-2011-4116](https://nvd.nist.gov/vuln/detail/CVE-2011-4116) | Negligible | `perl-modules-5.36` | `5.36.0-7+deb12u3` |
| [CVE-2011-4116](https://nvd.nist.gov/vuln/detail/CVE-2011-4116) | Negligible | `perl-base` | `5.36.0-7+deb12u3` |
| [CVE-2011-4116](https://nvd.nist.gov/vuln/detail/CVE-2011-4116) | Negligible | `perl` | `5.36.0-7+deb12u3` |
| [CVE-2012-0039](https://nvd.nist.gov/vuln/detail/CVE-2012-0039) | Negligible | `libglib2.0-data` | `2.74.6-2+deb12u8` |
| [CVE-2012-0039](https://nvd.nist.gov/vuln/detail/CVE-2012-0039) | Negligible | `libglib2.0-dev-bin` | `2.74.6-2+deb12u8` |
| [CVE-2012-0039](https://nvd.nist.gov/vuln/detail/CVE-2012-0039) | Negligible | `libglib2.0-dev` | `2.74.6-2+deb12u8` |
| [CVE-2012-0039](https://nvd.nist.gov/vuln/detail/CVE-2012-0039) | Negligible | `libglib2.0-bin` | `2.74.6-2+deb12u8` |
| [CVE-2012-0039](https://nvd.nist.gov/vuln/detail/CVE-2012-0039) | Negligible | `libglib2.0-0` | `2.74.6-2+deb12u8` |
| [CVE-2013-4392](https://nvd.nist.gov/vuln/detail/CVE-2013-4392) | Negligible | `libudev1` | `252.39-1~deb12u1` |
| [CVE-2013-4392](https://nvd.nist.gov/vuln/detail/CVE-2013-4392) | Negligible | `libsystemd0` | `252.39-1~deb12u1` |
| [CVE-2015-3276](https://nvd.nist.gov/vuln/detail/CVE-2015-3276) | Negligible | `libldap-2.5-0` | `2.5.13+dfsg-5` |
| [CVE-2015-9019](https://nvd.nist.gov/vuln/detail/CVE-2015-9019) | Negligible | `libxslt1.1` | `1.1.35-1+deb12u3` |
| [CVE-2015-9019](https://nvd.nist.gov/vuln/detail/CVE-2015-9019) | Negligible | `libxslt1-dev` | `1.1.35-1+deb12u3` |
| [CVE-2016-10505](https://nvd.nist.gov/vuln/detail/CVE-2016-10505) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2016-10505](https://nvd.nist.gov/vuln/detail/CVE-2016-10505) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-20012](https://nvd.nist.gov/vuln/detail/CVE-2016-20012) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-8678](https://nvd.nist.gov/vuln/detail/CVE-2016-8678) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2016-9113](https://nvd.nist.gov/vuln/detail/CVE-2016-9113) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-9113](https://nvd.nist.gov/vuln/detail/CVE-2016-9113) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2016-9114](https://nvd.nist.gov/vuln/detail/CVE-2016-9114) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2016-9114](https://nvd.nist.gov/vuln/detail/CVE-2016-9114) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-9115](https://nvd.nist.gov/vuln/detail/CVE-2016-9115) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2016-9115](https://nvd.nist.gov/vuln/detail/CVE-2016-9115) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-9116](https://nvd.nist.gov/vuln/detail/CVE-2016-9116) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2016-9116](https://nvd.nist.gov/vuln/detail/CVE-2016-9116) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-9117](https://nvd.nist.gov/vuln/detail/CVE-2016-9117) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2016-9117](https://nvd.nist.gov/vuln/detail/CVE-2016-9117) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-9580](https://nvd.nist.gov/vuln/detail/CVE-2016-9580) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-9580](https://nvd.nist.gov/vuln/detail/CVE-2016-9580) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2016-9581](https://nvd.nist.gov/vuln/detail/CVE-2016-9581) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2016-9581](https://nvd.nist.gov/vuln/detail/CVE-2016-9581) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11754](https://nvd.nist.gov/vuln/detail/CVE-2017-11754) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-11755](https://nvd.nist.gov/vuln/detail/CVE-2017-11755) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-13716](https://nvd.nist.gov/vuln/detail/CVE-2017-13716) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2017-13716](https://nvd.nist.gov/vuln/detail/CVE-2017-13716) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2017-13716](https://nvd.nist.gov/vuln/detail/CVE-2017-13716) | Negligible | `binutils` | `2.40-2` |
| [CVE-2017-13716](https://nvd.nist.gov/vuln/detail/CVE-2017-13716) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2017-13716](https://nvd.nist.gov/vuln/detail/CVE-2017-13716) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2017-13716](https://nvd.nist.gov/vuln/detail/CVE-2017-13716) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2017-13716](https://nvd.nist.gov/vuln/detail/CVE-2017-13716) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2017-14159](https://nvd.nist.gov/vuln/detail/CVE-2017-14159) | Negligible | `libldap-2.5-0` | `2.5.13+dfsg-5` |
| [CVE-2017-14988](https://nvd.nist.gov/vuln/detail/CVE-2017-14988) | Negligible | `libopenexr-dev` | `3.1.5-5` |
| [CVE-2017-14988](https://nvd.nist.gov/vuln/detail/CVE-2017-14988) | Negligible | `libopenexr-3-1-30` | `3.1.5-5` |
| [CVE-2017-16232](https://nvd.nist.gov/vuln/detail/CVE-2017-16232) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2017-16232](https://nvd.nist.gov/vuln/detail/CVE-2017-16232) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2017-16232](https://nvd.nist.gov/vuln/detail/CVE-2017-16232) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2017-17740](https://nvd.nist.gov/vuln/detail/CVE-2017-17740) | Negligible | `libldap-2.5-0` | `2.5.13+dfsg-5` |
| [CVE-2017-18018](https://nvd.nist.gov/vuln/detail/CVE-2017-18018) | Negligible | `coreutils` | `9.1-1` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-7275](https://nvd.nist.gov/vuln/detail/CVE-2017-7275) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2017-9937](https://nvd.nist.gov/vuln/detail/CVE-2017-9937) | Negligible | `libjbig-dev` | `2.1-6.1` |
| [CVE-2017-9937](https://nvd.nist.gov/vuln/detail/CVE-2017-9937) | Negligible | `libjbig0` | `2.1-6.1` |
| [CVE-2018-1000021](https://nvd.nist.gov/vuln/detail/CVE-2018-1000021) | Negligible | `git-man` | `1:2.39.5-0+deb12u3` |
| [CVE-2018-1000021](https://nvd.nist.gov/vuln/detail/CVE-2018-1000021) | Negligible | `git` | `1:2.39.5-0+deb12u3` |
| [CVE-2018-10126](https://nvd.nist.gov/vuln/detail/CVE-2018-10126) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2018-10126](https://nvd.nist.gov/vuln/detail/CVE-2018-10126) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2018-10126](https://nvd.nist.gov/vuln/detail/CVE-2018-10126) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15607](https://nvd.nist.gov/vuln/detail/CVE-2018-15607) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2018-15919](https://nvd.nist.gov/vuln/detail/CVE-2018-15919) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2018-16376](https://nvd.nist.gov/vuln/detail/CVE-2018-16376) | Negligible | `libopenjp2-7` | `2.5.0-2+deb12u2` |
| [CVE-2018-16376](https://nvd.nist.gov/vuln/detail/CVE-2018-16376) | Negligible | `libopenjp2-7-dev` | `2.5.0-2+deb12u2` |
| [CVE-2018-18064](https://nvd.nist.gov/vuln/detail/CVE-2018-18064) | Negligible | `libcairo-script-interpreter2` | `1.16.0-7` |
| [CVE-2018-18064](https://nvd.nist.gov/vuln/detail/CVE-2018-18064) | Negligible | `libcairo2-dev` | `1.16.0-7` |
| [CVE-2018-18064](https://nvd.nist.gov/vuln/detail/CVE-2018-18064) | Negligible | `libcairo2` | `1.16.0-7` |
| [CVE-2018-18064](https://nvd.nist.gov/vuln/detail/CVE-2018-18064) | Negligible | `libcairo-gobject2` | `1.16.0-7` |
| [CVE-2018-20673](https://nvd.nist.gov/vuln/detail/CVE-2018-20673) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2018-20673](https://nvd.nist.gov/vuln/detail/CVE-2018-20673) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2018-20673](https://nvd.nist.gov/vuln/detail/CVE-2018-20673) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2018-20673](https://nvd.nist.gov/vuln/detail/CVE-2018-20673) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2018-20673](https://nvd.nist.gov/vuln/detail/CVE-2018-20673) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2018-20673](https://nvd.nist.gov/vuln/detail/CVE-2018-20673) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2018-20673](https://nvd.nist.gov/vuln/detail/CVE-2018-20673) | Negligible | `binutils` | `2.40-2` |
| [CVE-2018-20712](https://nvd.nist.gov/vuln/detail/CVE-2018-20712) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2018-20712](https://nvd.nist.gov/vuln/detail/CVE-2018-20712) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2018-20712](https://nvd.nist.gov/vuln/detail/CVE-2018-20712) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2018-20712](https://nvd.nist.gov/vuln/detail/CVE-2018-20712) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2018-20712](https://nvd.nist.gov/vuln/detail/CVE-2018-20712) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2018-20712](https://nvd.nist.gov/vuln/detail/CVE-2018-20712) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2018-20712](https://nvd.nist.gov/vuln/detail/CVE-2018-20712) | Negligible | `binutils` | `2.40-2` |
| [CVE-2018-20796](https://nvd.nist.gov/vuln/detail/CVE-2018-20796) | Negligible | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2018-20796](https://nvd.nist.gov/vuln/detail/CVE-2018-20796) | Negligible | `libc6` | `2.36-9+deb12u13` |
| [CVE-2018-20796](https://nvd.nist.gov/vuln/detail/CVE-2018-20796) | Negligible | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2018-20796](https://nvd.nist.gov/vuln/detail/CVE-2018-20796) | Negligible | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libkdb5-10` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libkrb5-dev` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libkrb5support0` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libkrb5-3` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libkadm5srv-mit12` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libkadm5clnt-mit12` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libk5crypto3` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libgssrpc4` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `libgssapi-krb5-2` | `1.20.1-2+deb12u4` |
| [CVE-2018-5709](https://nvd.nist.gov/vuln/detail/CVE-2018-5709) | Negligible | `krb5-multidev` | `1.20.1-2+deb12u4` |
| [CVE-2018-6829](https://nvd.nist.gov/vuln/detail/CVE-2018-6829) | Negligible | `libgcrypt20` | `1.10.1-3` |
| [CVE-2018-6951](https://nvd.nist.gov/vuln/detail/CVE-2018-6951) | Negligible | `patch` | `2.7.6-7` |
| [CVE-2018-6952](https://nvd.nist.gov/vuln/detail/CVE-2018-6952) | Negligible | `patch` | `2.7.6-7` |
| [CVE-2018-9996](https://nvd.nist.gov/vuln/detail/CVE-2018-9996) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2018-9996](https://nvd.nist.gov/vuln/detail/CVE-2018-9996) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2018-9996](https://nvd.nist.gov/vuln/detail/CVE-2018-9996) | Negligible | `binutils` | `2.40-2` |
| [CVE-2018-9996](https://nvd.nist.gov/vuln/detail/CVE-2018-9996) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2018-9996](https://nvd.nist.gov/vuln/detail/CVE-2018-9996) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2018-9996](https://nvd.nist.gov/vuln/detail/CVE-2018-9996) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2018-9996](https://nvd.nist.gov/vuln/detail/CVE-2018-9996) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2019-1010022](https://nvd.nist.gov/vuln/detail/CVE-2019-1010022) | Negligible | `libc6` | `2.36-9+deb12u13` |
| [CVE-2019-1010022](https://nvd.nist.gov/vuln/detail/CVE-2019-1010022) | Negligible | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2019-1010022](https://nvd.nist.gov/vuln/detail/CVE-2019-1010022) | Negligible | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2019-1010022](https://nvd.nist.gov/vuln/detail/CVE-2019-1010022) | Negligible | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2019-1010023](https://nvd.nist.gov/vuln/detail/CVE-2019-1010023) | Negligible | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2019-1010023](https://nvd.nist.gov/vuln/detail/CVE-2019-1010023) | Negligible | `libc6` | `2.36-9+deb12u13` |
| [CVE-2019-1010023](https://nvd.nist.gov/vuln/detail/CVE-2019-1010023) | Negligible | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2019-1010023](https://nvd.nist.gov/vuln/detail/CVE-2019-1010023) | Negligible | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2019-1010024](https://nvd.nist.gov/vuln/detail/CVE-2019-1010024) | Negligible | `libc6` | `2.36-9+deb12u13` |
| [CVE-2019-1010024](https://nvd.nist.gov/vuln/detail/CVE-2019-1010024) | Negligible | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2019-1010024](https://nvd.nist.gov/vuln/detail/CVE-2019-1010024) | Negligible | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2019-1010024](https://nvd.nist.gov/vuln/detail/CVE-2019-1010024) | Negligible | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2019-1010025](https://nvd.nist.gov/vuln/detail/CVE-2019-1010025) | Negligible | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2019-1010025](https://nvd.nist.gov/vuln/detail/CVE-2019-1010025) | Negligible | `libc6` | `2.36-9+deb12u13` |
| [CVE-2019-1010025](https://nvd.nist.gov/vuln/detail/CVE-2019-1010025) | Negligible | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2019-1010025](https://nvd.nist.gov/vuln/detail/CVE-2019-1010025) | Negligible | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2019-6110](https://nvd.nist.gov/vuln/detail/CVE-2019-6110) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2019-9192](https://nvd.nist.gov/vuln/detail/CVE-2019-9192) | Negligible | `libc6-dev` | `2.36-9+deb12u13` |
| [CVE-2019-9192](https://nvd.nist.gov/vuln/detail/CVE-2019-9192) | Negligible | `libc-bin` | `2.36-9+deb12u13` |
| [CVE-2019-9192](https://nvd.nist.gov/vuln/detail/CVE-2019-9192) | Negligible | `libc-dev-bin` | `2.36-9+deb12u13` |
| [CVE-2019-9192](https://nvd.nist.gov/vuln/detail/CVE-2019-9192) | Negligible | `libc6` | `2.36-9+deb12u13` |
| [CVE-2020-14145](https://nvd.nist.gov/vuln/detail/CVE-2020-14145) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2020-15719](https://nvd.nist.gov/vuln/detail/CVE-2020-15719) | Negligible | `libldap-2.5-0` | `2.5.13+dfsg-5` |
| [CVE-2020-15778](https://nvd.nist.gov/vuln/detail/CVE-2020-15778) | Negligible | `openssh-client` | `1:9.2p1-2+deb12u7` |
| [CVE-2020-36325](https://nvd.nist.gov/vuln/detail/CVE-2020-36325) | Negligible | `libjansson4` | `2.14-2` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-20311](https://nvd.nist.gov/vuln/detail/CVE-2021-20311) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2021-32256](https://nvd.nist.gov/vuln/detail/CVE-2021-32256) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2021-32256](https://nvd.nist.gov/vuln/detail/CVE-2021-32256) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2021-32256](https://nvd.nist.gov/vuln/detail/CVE-2021-32256) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2021-32256](https://nvd.nist.gov/vuln/detail/CVE-2021-32256) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2021-32256](https://nvd.nist.gov/vuln/detail/CVE-2021-32256) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2021-32256](https://nvd.nist.gov/vuln/detail/CVE-2021-32256) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2021-32256](https://nvd.nist.gov/vuln/detail/CVE-2021-32256) | Negligible | `binutils` | `2.40-2` |
| [CVE-2021-4214](https://nvd.nist.gov/vuln/detail/CVE-2021-4214) | Negligible | `libpng16-16` | `1.6.39-2+deb12u1` |
| [CVE-2021-4214](https://nvd.nist.gov/vuln/detail/CVE-2021-4214) | Negligible | `libpng-dev` | `1.6.39-2+deb12u1` |
| [CVE-2021-4217](https://nvd.nist.gov/vuln/detail/CVE-2021-4217) | Negligible | `unzip` | `6.0-28` |
| [CVE-2021-45261](https://nvd.nist.gov/vuln/detail/CVE-2021-45261) | Negligible | `patch` | `2.7.6-7` |
| [CVE-2021-45346](https://nvd.nist.gov/vuln/detail/CVE-2021-45346) | Negligible | `libsqlite3-0` | `3.40.1-2+deb12u2` |
| [CVE-2021-45346](https://nvd.nist.gov/vuln/detail/CVE-2021-45346) | Negligible | `libsqlite3-dev` | `3.40.1-2+deb12u2` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libmount1` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `util-linux-extra` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `bsdutils` | `1:2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libsmartcols1` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libuuid1` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `mount` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `util-linux` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libmount-dev` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libblkid-dev` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `uuid-dev` | `2.38.1-5+deb12u3` |
| [CVE-2022-0563](https://nvd.nist.gov/vuln/detail/CVE-2022-0563) | Negligible | `libblkid1` | `2.38.1-5+deb12u3` |
| [CVE-2022-1210](https://nvd.nist.gov/vuln/detail/CVE-2022-1210) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2022-1210](https://nvd.nist.gov/vuln/detail/CVE-2022-1210) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2022-1210](https://nvd.nist.gov/vuln/detail/CVE-2022-1210) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2022-24975](https://nvd.nist.gov/vuln/detail/CVE-2022-24975) | Negligible | `git-man` | `1:2.39.5-0+deb12u3` |
| [CVE-2022-24975](https://nvd.nist.gov/vuln/detail/CVE-2022-24975) | Negligible | `git` | `1:2.39.5-0+deb12u3` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libhwasan0` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libgcc-s1` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libtsan2` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libstdc++6` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libstdc++-12-dev` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `liblsan0` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libitm1` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libgomp1` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libubsan1` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `cpp-12` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `g++-12` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `gcc-12` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `gcc-12-base` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libgcc-12-dev` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libcc1-0` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libatomic1` | `12.2.0-14+deb12u1` |
| [CVE-2022-27943](https://nvd.nist.gov/vuln/detail/CVE-2022-27943) | Negligible | `libasan8` | `12.2.0-14+deb12u1` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `dirmngr` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gpg-agent` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gpgv` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gpg-wks-client` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gnupg-utils` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gnupg-l10n` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gnupg` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gpgsm` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gpg` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gpg-wks-server` | `2.2.40-1.1+deb12u2` |
| [CVE-2022-3219](https://nvd.nist.gov/vuln/detail/CVE-2022-3219) | Negligible | `gpgconf` | `2.2.40-1.1+deb12u2` |
| [CVE-2023-1916](https://nvd.nist.gov/vuln/detail/CVE-2023-1916) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2023-1916](https://nvd.nist.gov/vuln/detail/CVE-2023-1916) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2023-1916](https://nvd.nist.gov/vuln/detail/CVE-2023-1916) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2023-1972](https://nvd.nist.gov/vuln/detail/CVE-2023-1972) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2023-1972](https://nvd.nist.gov/vuln/detail/CVE-2023-1972) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2023-1972](https://nvd.nist.gov/vuln/detail/CVE-2023-1972) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2023-1972](https://nvd.nist.gov/vuln/detail/CVE-2023-1972) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2023-1972](https://nvd.nist.gov/vuln/detail/CVE-2023-1972) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2023-1972](https://nvd.nist.gov/vuln/detail/CVE-2023-1972) | Negligible | `binutils` | `2.40-2` |
| [CVE-2023-1972](https://nvd.nist.gov/vuln/detail/CVE-2023-1972) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2023-31437](https://nvd.nist.gov/vuln/detail/CVE-2023-31437) | Negligible | `libsystemd0` | `252.39-1~deb12u1` |
| [CVE-2023-31437](https://nvd.nist.gov/vuln/detail/CVE-2023-31437) | Negligible | `libudev1` | `252.39-1~deb12u1` |
| [CVE-2023-31438](https://nvd.nist.gov/vuln/detail/CVE-2023-31438) | Negligible | `libudev1` | `252.39-1~deb12u1` |
| [CVE-2023-31438](https://nvd.nist.gov/vuln/detail/CVE-2023-31438) | Negligible | `libsystemd0` | `252.39-1~deb12u1` |
| [CVE-2023-31439](https://nvd.nist.gov/vuln/detail/CVE-2023-31439) | Negligible | `libudev1` | `252.39-1~deb12u1` |
| [CVE-2023-31439](https://nvd.nist.gov/vuln/detail/CVE-2023-31439) | Negligible | `libsystemd0` | `252.39-1~deb12u1` |
| [CVE-2023-31486](https://nvd.nist.gov/vuln/detail/CVE-2023-31486) | Negligible | `perl-modules-5.36` | `5.36.0-7+deb12u3` |
| [CVE-2023-31486](https://nvd.nist.gov/vuln/detail/CVE-2023-31486) | Negligible | `perl-base` | `5.36.0-7+deb12u3` |
| [CVE-2023-31486](https://nvd.nist.gov/vuln/detail/CVE-2023-31486) | Negligible | `perl` | `5.36.0-7+deb12u3` |
| [CVE-2023-31486](https://nvd.nist.gov/vuln/detail/CVE-2023-31486) | Negligible | `libperl5.36` | `5.36.0-7+deb12u3` |
| [CVE-2023-3164](https://nvd.nist.gov/vuln/detail/CVE-2023-3164) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2023-3164](https://nvd.nist.gov/vuln/detail/CVE-2023-3164) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2023-3164](https://nvd.nist.gov/vuln/detail/CVE-2023-3164) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-34152](https://nvd.nist.gov/vuln/detail/CVE-2023-34152) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2023-37769](https://nvd.nist.gov/vuln/detail/CVE-2023-37769) | Negligible | `libpixman-1-0` | `0.42.2-1` |
| [CVE-2023-37769](https://nvd.nist.gov/vuln/detail/CVE-2023-37769) | Negligible | `libpixman-1-dev` | `0.42.2-1` |
| [CVE-2023-49463](https://nvd.nist.gov/vuln/detail/CVE-2023-49463) | Negligible | `libheif1` | `1.15.1-1+deb12u1` |
| [CVE-2023-52426](https://nvd.nist.gov/vuln/detail/CVE-2023-52426) | Negligible | `libexpat1-dev` | `2.5.0-1+deb12u2` |
| [CVE-2023-52426](https://nvd.nist.gov/vuln/detail/CVE-2023-52426) | Negligible | `libexpat1` | `2.5.0-1+deb12u2` |
| [CVE-2023-6228](https://nvd.nist.gov/vuln/detail/CVE-2023-6228) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2023-6228](https://nvd.nist.gov/vuln/detail/CVE-2023-6228) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2023-6228](https://nvd.nist.gov/vuln/detail/CVE-2023-6228) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2024-2236](https://nvd.nist.gov/vuln/detail/CVE-2024-2236) | Negligible | `libgcrypt20` | `1.10.1-3` |
| [CVE-2024-2379](https://nvd.nist.gov/vuln/detail/CVE-2024-2379) | Negligible | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2024-2379](https://nvd.nist.gov/vuln/detail/CVE-2024-2379) | Negligible | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2024-2379](https://nvd.nist.gov/vuln/detail/CVE-2024-2379) | Negligible | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2024-2379](https://nvd.nist.gov/vuln/detail/CVE-2024-2379) | Negligible | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2024-25260](https://nvd.nist.gov/vuln/detail/CVE-2024-25260) | Negligible | `libelf1` | `0.188-2.1` |
| [CVE-2024-25269](https://nvd.nist.gov/vuln/detail/CVE-2024-25269) | Negligible | `libheif1` | `1.15.1-1+deb12u1` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libgssrpc4` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libkadm5clnt-mit12` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libk5crypto3` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libgssapi-krb5-2` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libkrb5support0` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libkrb5-dev` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libkrb5-3` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libkdb5-10` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `krb5-multidev` | `1.20.1-2+deb12u4` |
| [CVE-2024-26458](https://nvd.nist.gov/vuln/detail/CVE-2024-26458) | Negligible | `libkadm5srv-mit12` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libkdb5-10` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libkrb5support0` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `krb5-multidev` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libgssrpc4` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libk5crypto3` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libkadm5clnt-mit12` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libkadm5srv-mit12` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libgssapi-krb5-2` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libkrb5-3` | `1.20.1-2+deb12u4` |
| [CVE-2024-26461](https://nvd.nist.gov/vuln/detail/CVE-2024-26461) | Negligible | `libkrb5-dev` | `1.20.1-2+deb12u4` |
| [CVE-2024-28757](https://nvd.nist.gov/vuln/detail/CVE-2024-28757) | Negligible | `libexpat1` | `2.5.0-1+deb12u2` |
| [CVE-2024-28757](https://nvd.nist.gov/vuln/detail/CVE-2024-28757) | Negligible | `libexpat1-dev` | `2.5.0-1+deb12u2` |
| [CVE-2024-52005](https://nvd.nist.gov/vuln/detail/CVE-2024-52005) | Negligible | `git-man` | `1:2.39.5-0+deb12u3` |
| [CVE-2024-52005](https://nvd.nist.gov/vuln/detail/CVE-2024-52005) | Negligible | `git` | `1:2.39.5-0+deb12u3` |
| [CVE-2024-53589](https://nvd.nist.gov/vuln/detail/CVE-2024-53589) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2024-53589](https://nvd.nist.gov/vuln/detail/CVE-2024-53589) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2024-53589](https://nvd.nist.gov/vuln/detail/CVE-2024-53589) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2024-53589](https://nvd.nist.gov/vuln/detail/CVE-2024-53589) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2024-53589](https://nvd.nist.gov/vuln/detail/CVE-2024-53589) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2024-53589](https://nvd.nist.gov/vuln/detail/CVE-2024-53589) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2024-53589](https://nvd.nist.gov/vuln/detail/CVE-2024-53589) | Negligible | `binutils` | `2.40-2` |
| [CVE-2024-57360](https://nvd.nist.gov/vuln/detail/CVE-2024-57360) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2024-57360](https://nvd.nist.gov/vuln/detail/CVE-2024-57360) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2024-57360](https://nvd.nist.gov/vuln/detail/CVE-2024-57360) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2024-57360](https://nvd.nist.gov/vuln/detail/CVE-2024-57360) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2024-57360](https://nvd.nist.gov/vuln/detail/CVE-2024-57360) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2024-57360](https://nvd.nist.gov/vuln/detail/CVE-2024-57360) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2024-57360](https://nvd.nist.gov/vuln/detail/CVE-2024-57360) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-0725](https://nvd.nist.gov/vuln/detail/CVE-2025-0725) | Negligible | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-0725](https://nvd.nist.gov/vuln/detail/CVE-2025-0725) | Negligible | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-0725](https://nvd.nist.gov/vuln/detail/CVE-2025-0725) | Negligible | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-0725](https://nvd.nist.gov/vuln/detail/CVE-2025-0725) | Negligible | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-0840](https://nvd.nist.gov/vuln/detail/CVE-2025-0840) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-0840](https://nvd.nist.gov/vuln/detail/CVE-2025-0840) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-0840](https://nvd.nist.gov/vuln/detail/CVE-2025-0840) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-0840](https://nvd.nist.gov/vuln/detail/CVE-2025-0840) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-0840](https://nvd.nist.gov/vuln/detail/CVE-2025-0840) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-0840](https://nvd.nist.gov/vuln/detail/CVE-2025-0840) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-0840](https://nvd.nist.gov/vuln/detail/CVE-2025-0840) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-10966](https://nvd.nist.gov/vuln/detail/CVE-2025-10966) | Negligible | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-10966](https://nvd.nist.gov/vuln/detail/CVE-2025-10966) | Negligible | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-10966](https://nvd.nist.gov/vuln/detail/CVE-2025-10966) | Negligible | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-10966](https://nvd.nist.gov/vuln/detail/CVE-2025-10966) | Negligible | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-11081](https://nvd.nist.gov/vuln/detail/CVE-2025-11081) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11081](https://nvd.nist.gov/vuln/detail/CVE-2025-11081) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11081](https://nvd.nist.gov/vuln/detail/CVE-2025-11081) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11081](https://nvd.nist.gov/vuln/detail/CVE-2025-11081) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11081](https://nvd.nist.gov/vuln/detail/CVE-2025-11081) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11081](https://nvd.nist.gov/vuln/detail/CVE-2025-11081) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11081](https://nvd.nist.gov/vuln/detail/CVE-2025-11081) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11082](https://nvd.nist.gov/vuln/detail/CVE-2025-11082) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11082](https://nvd.nist.gov/vuln/detail/CVE-2025-11082) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11082](https://nvd.nist.gov/vuln/detail/CVE-2025-11082) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11082](https://nvd.nist.gov/vuln/detail/CVE-2025-11082) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11082](https://nvd.nist.gov/vuln/detail/CVE-2025-11082) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11082](https://nvd.nist.gov/vuln/detail/CVE-2025-11082) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11082](https://nvd.nist.gov/vuln/detail/CVE-2025-11082) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11083](https://nvd.nist.gov/vuln/detail/CVE-2025-11083) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11083](https://nvd.nist.gov/vuln/detail/CVE-2025-11083) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11083](https://nvd.nist.gov/vuln/detail/CVE-2025-11083) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11083](https://nvd.nist.gov/vuln/detail/CVE-2025-11083) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11083](https://nvd.nist.gov/vuln/detail/CVE-2025-11083) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11083](https://nvd.nist.gov/vuln/detail/CVE-2025-11083) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11083](https://nvd.nist.gov/vuln/detail/CVE-2025-11083) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11412](https://nvd.nist.gov/vuln/detail/CVE-2025-11412) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11412](https://nvd.nist.gov/vuln/detail/CVE-2025-11412) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11412](https://nvd.nist.gov/vuln/detail/CVE-2025-11412) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11412](https://nvd.nist.gov/vuln/detail/CVE-2025-11412) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11412](https://nvd.nist.gov/vuln/detail/CVE-2025-11412) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11412](https://nvd.nist.gov/vuln/detail/CVE-2025-11412) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11412](https://nvd.nist.gov/vuln/detail/CVE-2025-11412) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11413](https://nvd.nist.gov/vuln/detail/CVE-2025-11413) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11413](https://nvd.nist.gov/vuln/detail/CVE-2025-11413) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11413](https://nvd.nist.gov/vuln/detail/CVE-2025-11413) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11413](https://nvd.nist.gov/vuln/detail/CVE-2025-11413) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11413](https://nvd.nist.gov/vuln/detail/CVE-2025-11413) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11413](https://nvd.nist.gov/vuln/detail/CVE-2025-11413) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11413](https://nvd.nist.gov/vuln/detail/CVE-2025-11413) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11414](https://nvd.nist.gov/vuln/detail/CVE-2025-11414) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11414](https://nvd.nist.gov/vuln/detail/CVE-2025-11414) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11414](https://nvd.nist.gov/vuln/detail/CVE-2025-11414) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11414](https://nvd.nist.gov/vuln/detail/CVE-2025-11414) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11414](https://nvd.nist.gov/vuln/detail/CVE-2025-11414) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11414](https://nvd.nist.gov/vuln/detail/CVE-2025-11414) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11414](https://nvd.nist.gov/vuln/detail/CVE-2025-11414) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11468](https://nvd.nist.gov/vuln/detail/CVE-2025-11468) | Negligible | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-11468](https://nvd.nist.gov/vuln/detail/CVE-2025-11468) | Negligible | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-11468](https://nvd.nist.gov/vuln/detail/CVE-2025-11468) | Negligible | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-11468](https://nvd.nist.gov/vuln/detail/CVE-2025-11468) | Negligible | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-1147](https://nvd.nist.gov/vuln/detail/CVE-2025-1147) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1147](https://nvd.nist.gov/vuln/detail/CVE-2025-1147) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1147](https://nvd.nist.gov/vuln/detail/CVE-2025-1147) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1147](https://nvd.nist.gov/vuln/detail/CVE-2025-1147) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1147](https://nvd.nist.gov/vuln/detail/CVE-2025-1147) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1147](https://nvd.nist.gov/vuln/detail/CVE-2025-1147) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1147](https://nvd.nist.gov/vuln/detail/CVE-2025-1147) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1148](https://nvd.nist.gov/vuln/detail/CVE-2025-1148) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1148](https://nvd.nist.gov/vuln/detail/CVE-2025-1148) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1148](https://nvd.nist.gov/vuln/detail/CVE-2025-1148) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1148](https://nvd.nist.gov/vuln/detail/CVE-2025-1148) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1148](https://nvd.nist.gov/vuln/detail/CVE-2025-1148) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1148](https://nvd.nist.gov/vuln/detail/CVE-2025-1148) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1148](https://nvd.nist.gov/vuln/detail/CVE-2025-1148) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1149](https://nvd.nist.gov/vuln/detail/CVE-2025-1149) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1149](https://nvd.nist.gov/vuln/detail/CVE-2025-1149) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1149](https://nvd.nist.gov/vuln/detail/CVE-2025-1149) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1149](https://nvd.nist.gov/vuln/detail/CVE-2025-1149) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1149](https://nvd.nist.gov/vuln/detail/CVE-2025-1149) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1149](https://nvd.nist.gov/vuln/detail/CVE-2025-1149) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1149](https://nvd.nist.gov/vuln/detail/CVE-2025-1149) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11494](https://nvd.nist.gov/vuln/detail/CVE-2025-11494) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11494](https://nvd.nist.gov/vuln/detail/CVE-2025-11494) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11494](https://nvd.nist.gov/vuln/detail/CVE-2025-11494) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11494](https://nvd.nist.gov/vuln/detail/CVE-2025-11494) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11494](https://nvd.nist.gov/vuln/detail/CVE-2025-11494) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11494](https://nvd.nist.gov/vuln/detail/CVE-2025-11494) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11494](https://nvd.nist.gov/vuln/detail/CVE-2025-11494) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11495](https://nvd.nist.gov/vuln/detail/CVE-2025-11495) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11495](https://nvd.nist.gov/vuln/detail/CVE-2025-11495) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11495](https://nvd.nist.gov/vuln/detail/CVE-2025-11495) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11495](https://nvd.nist.gov/vuln/detail/CVE-2025-11495) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11495](https://nvd.nist.gov/vuln/detail/CVE-2025-11495) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11495](https://nvd.nist.gov/vuln/detail/CVE-2025-11495) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11495](https://nvd.nist.gov/vuln/detail/CVE-2025-11495) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1150](https://nvd.nist.gov/vuln/detail/CVE-2025-1150) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1150](https://nvd.nist.gov/vuln/detail/CVE-2025-1150) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1150](https://nvd.nist.gov/vuln/detail/CVE-2025-1150) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1150](https://nvd.nist.gov/vuln/detail/CVE-2025-1150) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1150](https://nvd.nist.gov/vuln/detail/CVE-2025-1150) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1150](https://nvd.nist.gov/vuln/detail/CVE-2025-1150) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1150](https://nvd.nist.gov/vuln/detail/CVE-2025-1150) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1151](https://nvd.nist.gov/vuln/detail/CVE-2025-1151) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1151](https://nvd.nist.gov/vuln/detail/CVE-2025-1151) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1151](https://nvd.nist.gov/vuln/detail/CVE-2025-1151) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1151](https://nvd.nist.gov/vuln/detail/CVE-2025-1151) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1151](https://nvd.nist.gov/vuln/detail/CVE-2025-1151) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1151](https://nvd.nist.gov/vuln/detail/CVE-2025-1151) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1151](https://nvd.nist.gov/vuln/detail/CVE-2025-1151) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1152](https://nvd.nist.gov/vuln/detail/CVE-2025-1152) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1152](https://nvd.nist.gov/vuln/detail/CVE-2025-1152) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1152](https://nvd.nist.gov/vuln/detail/CVE-2025-1152) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1152](https://nvd.nist.gov/vuln/detail/CVE-2025-1152) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1152](https://nvd.nist.gov/vuln/detail/CVE-2025-1152) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1152](https://nvd.nist.gov/vuln/detail/CVE-2025-1152) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1152](https://nvd.nist.gov/vuln/detail/CVE-2025-1152) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1153](https://nvd.nist.gov/vuln/detail/CVE-2025-1153) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1153](https://nvd.nist.gov/vuln/detail/CVE-2025-1153) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1153](https://nvd.nist.gov/vuln/detail/CVE-2025-1153) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1153](https://nvd.nist.gov/vuln/detail/CVE-2025-1153) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1153](https://nvd.nist.gov/vuln/detail/CVE-2025-1153) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1153](https://nvd.nist.gov/vuln/detail/CVE-2025-1153) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1153](https://nvd.nist.gov/vuln/detail/CVE-2025-1153) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1176](https://nvd.nist.gov/vuln/detail/CVE-2025-1176) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1176](https://nvd.nist.gov/vuln/detail/CVE-2025-1176) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1176](https://nvd.nist.gov/vuln/detail/CVE-2025-1176) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1176](https://nvd.nist.gov/vuln/detail/CVE-2025-1176) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1176](https://nvd.nist.gov/vuln/detail/CVE-2025-1176) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1176](https://nvd.nist.gov/vuln/detail/CVE-2025-1176) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1176](https://nvd.nist.gov/vuln/detail/CVE-2025-1176) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1178](https://nvd.nist.gov/vuln/detail/CVE-2025-1178) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1178](https://nvd.nist.gov/vuln/detail/CVE-2025-1178) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1178](https://nvd.nist.gov/vuln/detail/CVE-2025-1178) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1178](https://nvd.nist.gov/vuln/detail/CVE-2025-1178) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1178](https://nvd.nist.gov/vuln/detail/CVE-2025-1178) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1178](https://nvd.nist.gov/vuln/detail/CVE-2025-1178) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1178](https://nvd.nist.gov/vuln/detail/CVE-2025-1178) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1179](https://nvd.nist.gov/vuln/detail/CVE-2025-1179) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1179](https://nvd.nist.gov/vuln/detail/CVE-2025-1179) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1179](https://nvd.nist.gov/vuln/detail/CVE-2025-1179) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1179](https://nvd.nist.gov/vuln/detail/CVE-2025-1179) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1179](https://nvd.nist.gov/vuln/detail/CVE-2025-1179) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1179](https://nvd.nist.gov/vuln/detail/CVE-2025-1179) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1179](https://nvd.nist.gov/vuln/detail/CVE-2025-1179) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1180](https://nvd.nist.gov/vuln/detail/CVE-2025-1180) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1180](https://nvd.nist.gov/vuln/detail/CVE-2025-1180) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1180](https://nvd.nist.gov/vuln/detail/CVE-2025-1180) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1180](https://nvd.nist.gov/vuln/detail/CVE-2025-1180) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1180](https://nvd.nist.gov/vuln/detail/CVE-2025-1180) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1180](https://nvd.nist.gov/vuln/detail/CVE-2025-1180) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1180](https://nvd.nist.gov/vuln/detail/CVE-2025-1180) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1181](https://nvd.nist.gov/vuln/detail/CVE-2025-1181) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1181](https://nvd.nist.gov/vuln/detail/CVE-2025-1181) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1181](https://nvd.nist.gov/vuln/detail/CVE-2025-1181) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-1181](https://nvd.nist.gov/vuln/detail/CVE-2025-1181) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1181](https://nvd.nist.gov/vuln/detail/CVE-2025-1181) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1181](https://nvd.nist.gov/vuln/detail/CVE-2025-1181) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1181](https://nvd.nist.gov/vuln/detail/CVE-2025-1181) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1182](https://nvd.nist.gov/vuln/detail/CVE-2025-1182) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-1182](https://nvd.nist.gov/vuln/detail/CVE-2025-1182) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-1182](https://nvd.nist.gov/vuln/detail/CVE-2025-1182) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-1182](https://nvd.nist.gov/vuln/detail/CVE-2025-1182) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-1182](https://nvd.nist.gov/vuln/detail/CVE-2025-1182) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1182](https://nvd.nist.gov/vuln/detail/CVE-2025-1182) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-1182](https://nvd.nist.gov/vuln/detail/CVE-2025-1182) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11839](https://nvd.nist.gov/vuln/detail/CVE-2025-11839) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-11839](https://nvd.nist.gov/vuln/detail/CVE-2025-11839) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11839](https://nvd.nist.gov/vuln/detail/CVE-2025-11839) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11839](https://nvd.nist.gov/vuln/detail/CVE-2025-11839) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11839](https://nvd.nist.gov/vuln/detail/CVE-2025-11839) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11839](https://nvd.nist.gov/vuln/detail/CVE-2025-11839) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11839](https://nvd.nist.gov/vuln/detail/CVE-2025-11839) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11840](https://nvd.nist.gov/vuln/detail/CVE-2025-11840) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-11840](https://nvd.nist.gov/vuln/detail/CVE-2025-11840) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-11840](https://nvd.nist.gov/vuln/detail/CVE-2025-11840) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-11840](https://nvd.nist.gov/vuln/detail/CVE-2025-11840) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-11840](https://nvd.nist.gov/vuln/detail/CVE-2025-11840) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-11840](https://nvd.nist.gov/vuln/detail/CVE-2025-11840) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-11840](https://nvd.nist.gov/vuln/detail/CVE-2025-11840) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-1352](https://nvd.nist.gov/vuln/detail/CVE-2025-1352) | Negligible | `libelf1` | `0.188-2.1` |
| [CVE-2025-1365](https://nvd.nist.gov/vuln/detail/CVE-2025-1365) | Negligible | `libelf1` | `0.188-2.1` |
| [CVE-2025-1371](https://nvd.nist.gov/vuln/detail/CVE-2025-1371) | Negligible | `libelf1` | `0.188-2.1` |
| [CVE-2025-1372](https://nvd.nist.gov/vuln/detail/CVE-2025-1372) | Negligible | `libelf1` | `0.188-2.1` |
| [CVE-2025-1376](https://nvd.nist.gov/vuln/detail/CVE-2025-1376) | Negligible | `libelf1` | `0.188-2.1` |
| [CVE-2025-1377](https://nvd.nist.gov/vuln/detail/CVE-2025-1377) | Negligible | `libelf1` | `0.188-2.1` |
| [CVE-2025-14017](https://nvd.nist.gov/vuln/detail/CVE-2025-14017) | Negligible | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-14017](https://nvd.nist.gov/vuln/detail/CVE-2025-14017) | Negligible | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-14017](https://nvd.nist.gov/vuln/detail/CVE-2025-14017) | Negligible | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-14017](https://nvd.nist.gov/vuln/detail/CVE-2025-14017) | Negligible | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-15079](https://nvd.nist.gov/vuln/detail/CVE-2025-15079) | Negligible | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-15079](https://nvd.nist.gov/vuln/detail/CVE-2025-15079) | Negligible | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-15079](https://nvd.nist.gov/vuln/detail/CVE-2025-15079) | Negligible | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-15079](https://nvd.nist.gov/vuln/detail/CVE-2025-15079) | Negligible | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-15224](https://nvd.nist.gov/vuln/detail/CVE-2025-15224) | Negligible | `curl` | `7.88.1-10+deb12u14` |
| [CVE-2025-15224](https://nvd.nist.gov/vuln/detail/CVE-2025-15224) | Negligible | `libcurl4-openssl-dev` | `7.88.1-10+deb12u14` |
| [CVE-2025-15224](https://nvd.nist.gov/vuln/detail/CVE-2025-15224) | Negligible | `libcurl4` | `7.88.1-10+deb12u14` |
| [CVE-2025-15224](https://nvd.nist.gov/vuln/detail/CVE-2025-15224) | Negligible | `libcurl3-gnutls` | `7.88.1-10+deb12u14` |
| [CVE-2025-27587](https://nvd.nist.gov/vuln/detail/CVE-2025-27587) | Negligible | `libssl-dev` | `3.0.18-1~deb12u2` |
| [CVE-2025-27587](https://nvd.nist.gov/vuln/detail/CVE-2025-27587) | Negligible | `libssl3` | `3.0.18-1~deb12u2` |
| [CVE-2025-27587](https://nvd.nist.gov/vuln/detail/CVE-2025-27587) | Negligible | `openssl` | `3.0.18-1~deb12u2` |
| [CVE-2025-28162](https://nvd.nist.gov/vuln/detail/CVE-2025-28162) | Negligible | `libpng16-16` | `1.6.39-2+deb12u1` |
| [CVE-2025-28162](https://nvd.nist.gov/vuln/detail/CVE-2025-28162) | Negligible | `libpng-dev` | `1.6.39-2+deb12u1` |
| [CVE-2025-28164](https://nvd.nist.gov/vuln/detail/CVE-2025-28164) | Negligible | `libpng-dev` | `1.6.39-2+deb12u1` |
| [CVE-2025-28164](https://nvd.nist.gov/vuln/detail/CVE-2025-28164) | Negligible | `libpng16-16` | `1.6.39-2+deb12u1` |
| [CVE-2025-29070](https://nvd.nist.gov/vuln/detail/CVE-2025-29070) | Negligible | `liblcms2-2` | `2.14-2` |
| [CVE-2025-29070](https://nvd.nist.gov/vuln/detail/CVE-2025-29070) | Negligible | `liblcms2-dev` | `2.14-2` |
| [CVE-2025-29088](https://nvd.nist.gov/vuln/detail/CVE-2025-29088) | Negligible | `libsqlite3-dev` | `3.40.1-2+deb12u2` |
| [CVE-2025-29088](https://nvd.nist.gov/vuln/detail/CVE-2025-29088) | Negligible | `libsqlite3-0` | `3.40.1-2+deb12u2` |
| [CVE-2025-3198](https://nvd.nist.gov/vuln/detail/CVE-2025-3198) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-3198](https://nvd.nist.gov/vuln/detail/CVE-2025-3198) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-3198](https://nvd.nist.gov/vuln/detail/CVE-2025-3198) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-3198](https://nvd.nist.gov/vuln/detail/CVE-2025-3198) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-3198](https://nvd.nist.gov/vuln/detail/CVE-2025-3198) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-3198](https://nvd.nist.gov/vuln/detail/CVE-2025-3198) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-3198](https://nvd.nist.gov/vuln/detail/CVE-2025-3198) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-4516](https://nvd.nist.gov/vuln/detail/CVE-2025-4516) | Negligible | `python3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-4516](https://nvd.nist.gov/vuln/detail/CVE-2025-4516) | Negligible | `libpython3.11-minimal` | `3.11.2-6+deb12u6` |
| [CVE-2025-4516](https://nvd.nist.gov/vuln/detail/CVE-2025-4516) | Negligible | `libpython3.11-stdlib` | `3.11.2-6+deb12u6` |
| [CVE-2025-4516](https://nvd.nist.gov/vuln/detail/CVE-2025-4516) | Negligible | `python3.11` | `3.11.2-6+deb12u6` |
| [CVE-2025-5244](https://nvd.nist.gov/vuln/detail/CVE-2025-5244) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-5244](https://nvd.nist.gov/vuln/detail/CVE-2025-5244) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-5244](https://nvd.nist.gov/vuln/detail/CVE-2025-5244) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-5244](https://nvd.nist.gov/vuln/detail/CVE-2025-5244) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-5244](https://nvd.nist.gov/vuln/detail/CVE-2025-5244) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-5244](https://nvd.nist.gov/vuln/detail/CVE-2025-5244) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-5244](https://nvd.nist.gov/vuln/detail/CVE-2025-5244) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-5245](https://nvd.nist.gov/vuln/detail/CVE-2025-5245) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-5245](https://nvd.nist.gov/vuln/detail/CVE-2025-5245) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-5245](https://nvd.nist.gov/vuln/detail/CVE-2025-5245) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-5245](https://nvd.nist.gov/vuln/detail/CVE-2025-5245) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-5245](https://nvd.nist.gov/vuln/detail/CVE-2025-5245) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-5245](https://nvd.nist.gov/vuln/detail/CVE-2025-5245) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-5245](https://nvd.nist.gov/vuln/detail/CVE-2025-5245) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-5278](https://nvd.nist.gov/vuln/detail/CVE-2025-5278) | Negligible | `coreutils` | `9.1-1` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `imagemagick-6-common` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `imagemagick-6.q16` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickcore-6-arch-config` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickwand-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickwand-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickwand-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickcore-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickcore-6.q16-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickcore-6.q16-6-extra` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickcore-6-headers` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickcore-6.q16-6` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `libmagickwand-dev` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-55160](https://nvd.nist.gov/vuln/detail/CVE-2025-55160) | Negligible | `imagemagick` | `8:6.9.11.60+dfsg-1.6+deb12u6` |
| [CVE-2025-66861](https://nvd.nist.gov/vuln/detail/CVE-2025-66861) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-66861](https://nvd.nist.gov/vuln/detail/CVE-2025-66861) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-66861](https://nvd.nist.gov/vuln/detail/CVE-2025-66861) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-66861](https://nvd.nist.gov/vuln/detail/CVE-2025-66861) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-66861](https://nvd.nist.gov/vuln/detail/CVE-2025-66861) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-66861](https://nvd.nist.gov/vuln/detail/CVE-2025-66861) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-66861](https://nvd.nist.gov/vuln/detail/CVE-2025-66861) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-66862](https://nvd.nist.gov/vuln/detail/CVE-2025-66862) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-66862](https://nvd.nist.gov/vuln/detail/CVE-2025-66862) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-66862](https://nvd.nist.gov/vuln/detail/CVE-2025-66862) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-66862](https://nvd.nist.gov/vuln/detail/CVE-2025-66862) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-66862](https://nvd.nist.gov/vuln/detail/CVE-2025-66862) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-66862](https://nvd.nist.gov/vuln/detail/CVE-2025-66862) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-66862](https://nvd.nist.gov/vuln/detail/CVE-2025-66862) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-66863](https://nvd.nist.gov/vuln/detail/CVE-2025-66863) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-66863](https://nvd.nist.gov/vuln/detail/CVE-2025-66863) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-66863](https://nvd.nist.gov/vuln/detail/CVE-2025-66863) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-66863](https://nvd.nist.gov/vuln/detail/CVE-2025-66863) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-66863](https://nvd.nist.gov/vuln/detail/CVE-2025-66863) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-66863](https://nvd.nist.gov/vuln/detail/CVE-2025-66863) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-66863](https://nvd.nist.gov/vuln/detail/CVE-2025-66863) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-66864](https://nvd.nist.gov/vuln/detail/CVE-2025-66864) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-66864](https://nvd.nist.gov/vuln/detail/CVE-2025-66864) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-66864](https://nvd.nist.gov/vuln/detail/CVE-2025-66864) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-66864](https://nvd.nist.gov/vuln/detail/CVE-2025-66864) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-66864](https://nvd.nist.gov/vuln/detail/CVE-2025-66864) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-66864](https://nvd.nist.gov/vuln/detail/CVE-2025-66864) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-66864](https://nvd.nist.gov/vuln/detail/CVE-2025-66864) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-66865](https://nvd.nist.gov/vuln/detail/CVE-2025-66865) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-66865](https://nvd.nist.gov/vuln/detail/CVE-2025-66865) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-66865](https://nvd.nist.gov/vuln/detail/CVE-2025-66865) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-66865](https://nvd.nist.gov/vuln/detail/CVE-2025-66865) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-66865](https://nvd.nist.gov/vuln/detail/CVE-2025-66865) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-66865](https://nvd.nist.gov/vuln/detail/CVE-2025-66865) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-66865](https://nvd.nist.gov/vuln/detail/CVE-2025-66865) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-66866](https://nvd.nist.gov/vuln/detail/CVE-2025-66866) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-66866](https://nvd.nist.gov/vuln/detail/CVE-2025-66866) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-66866](https://nvd.nist.gov/vuln/detail/CVE-2025-66866) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-66866](https://nvd.nist.gov/vuln/detail/CVE-2025-66866) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-66866](https://nvd.nist.gov/vuln/detail/CVE-2025-66866) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-66866](https://nvd.nist.gov/vuln/detail/CVE-2025-66866) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-66866](https://nvd.nist.gov/vuln/detail/CVE-2025-66866) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-7545](https://nvd.nist.gov/vuln/detail/CVE-2025-7545) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-7545](https://nvd.nist.gov/vuln/detail/CVE-2025-7545) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-7545](https://nvd.nist.gov/vuln/detail/CVE-2025-7545) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-7545](https://nvd.nist.gov/vuln/detail/CVE-2025-7545) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-7545](https://nvd.nist.gov/vuln/detail/CVE-2025-7545) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-7545](https://nvd.nist.gov/vuln/detail/CVE-2025-7545) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-7545](https://nvd.nist.gov/vuln/detail/CVE-2025-7545) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-7546](https://nvd.nist.gov/vuln/detail/CVE-2025-7546) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-7546](https://nvd.nist.gov/vuln/detail/CVE-2025-7546) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-7546](https://nvd.nist.gov/vuln/detail/CVE-2025-7546) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-7546](https://nvd.nist.gov/vuln/detail/CVE-2025-7546) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-7546](https://nvd.nist.gov/vuln/detail/CVE-2025-7546) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-7546](https://nvd.nist.gov/vuln/detail/CVE-2025-7546) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-7546](https://nvd.nist.gov/vuln/detail/CVE-2025-7546) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-8176](https://nvd.nist.gov/vuln/detail/CVE-2025-8176) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8176](https://nvd.nist.gov/vuln/detail/CVE-2025-8176) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2025-8176](https://nvd.nist.gov/vuln/detail/CVE-2025-8176) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8177](https://nvd.nist.gov/vuln/detail/CVE-2025-8177) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8177](https://nvd.nist.gov/vuln/detail/CVE-2025-8177) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8177](https://nvd.nist.gov/vuln/detail/CVE-2025-8177) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2025-8224](https://nvd.nist.gov/vuln/detail/CVE-2025-8224) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-8224](https://nvd.nist.gov/vuln/detail/CVE-2025-8224) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-8224](https://nvd.nist.gov/vuln/detail/CVE-2025-8224) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-8224](https://nvd.nist.gov/vuln/detail/CVE-2025-8224) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-8224](https://nvd.nist.gov/vuln/detail/CVE-2025-8224) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-8224](https://nvd.nist.gov/vuln/detail/CVE-2025-8224) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-8224](https://nvd.nist.gov/vuln/detail/CVE-2025-8224) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-8225](https://nvd.nist.gov/vuln/detail/CVE-2025-8225) | Negligible | `binutils-aarch64-linux-gnu` | `2.40-2` |
| [CVE-2025-8225](https://nvd.nist.gov/vuln/detail/CVE-2025-8225) | Negligible | `libctf0` | `2.40-2` |
| [CVE-2025-8225](https://nvd.nist.gov/vuln/detail/CVE-2025-8225) | Negligible | `libctf-nobfd0` | `2.40-2` |
| [CVE-2025-8225](https://nvd.nist.gov/vuln/detail/CVE-2025-8225) | Negligible | `libbinutils` | `2.40-2` |
| [CVE-2025-8225](https://nvd.nist.gov/vuln/detail/CVE-2025-8225) | Negligible | `binutils-common` | `2.40-2` |
| [CVE-2025-8225](https://nvd.nist.gov/vuln/detail/CVE-2025-8225) | Negligible | `binutils` | `2.40-2` |
| [CVE-2025-8225](https://nvd.nist.gov/vuln/detail/CVE-2025-8225) | Negligible | `libgprofng0` | `2.40-2` |
| [CVE-2025-8534](https://nvd.nist.gov/vuln/detail/CVE-2025-8534) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8534](https://nvd.nist.gov/vuln/detail/CVE-2025-8534) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2025-8534](https://nvd.nist.gov/vuln/detail/CVE-2025-8534) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8732](https://nvd.nist.gov/vuln/detail/CVE-2025-8732) | Negligible | `libxml2-dev` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2025-8732](https://nvd.nist.gov/vuln/detail/CVE-2025-8732) | Negligible | `libxml2` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2025-8851](https://nvd.nist.gov/vuln/detail/CVE-2025-8851) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2025-8851](https://nvd.nist.gov/vuln/detail/CVE-2025-8851) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8851](https://nvd.nist.gov/vuln/detail/CVE-2025-8851) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8961](https://nvd.nist.gov/vuln/detail/CVE-2025-8961) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8961](https://nvd.nist.gov/vuln/detail/CVE-2025-8961) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2025-8961](https://nvd.nist.gov/vuln/detail/CVE-2025-8961) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2025-9165](https://nvd.nist.gov/vuln/detail/CVE-2025-9165) | Negligible | `libtiffxx6` | `4.5.0-6+deb12u3` |
| [CVE-2025-9165](https://nvd.nist.gov/vuln/detail/CVE-2025-9165) | Negligible | `libtiff6` | `4.5.0-6+deb12u3` |
| [CVE-2025-9165](https://nvd.nist.gov/vuln/detail/CVE-2025-9165) | Negligible | `libtiff-dev` | `4.5.0-6+deb12u3` |
| [CVE-2026-1757](https://nvd.nist.gov/vuln/detail/CVE-2026-1757) | Negligible | `libxml2-dev` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-1757](https://nvd.nist.gov/vuln/detail/CVE-2026-1757) | Negligible | `libxml2` | `2.9.14+dfsg-1.3~deb12u5` |
| [CVE-2026-22185](https://nvd.nist.gov/vuln/detail/CVE-2026-22185) | Negligible | `libldap-2.5-0` | `2.5.13+dfsg-5` |
</details>

<details>
<summary><strong>📦 Installed Packages (576 total)</strong></summary>

| Package | Version |
|---------|---------|
| @isaacs/balanced-match | 4.0.1 |
| @isaacs/brace-expansion | 5.0.0 |
| @isaacs/fs-minipass | 4.0.1 |
| @isaacs/string-locale-compare | 1.1.0 |
| @npmcli/agent | 4.0.0 |
| @npmcli/arborist | 9.1.10 |
| @npmcli/config | 10.5.0 |
| @npmcli/fs | 5.0.0 |
| @npmcli/git | 7.0.1 |
| @npmcli/installed-package-contents | 4.0.0 |
| @npmcli/map-workspaces | 5.0.3 |
| @npmcli/metavuln-calculator | 9.0.3 |
| @npmcli/name-from-folder | 4.0.0 |
| @npmcli/node-gyp | 5.0.0 |
| @npmcli/package-json | 7.0.4 |
| @npmcli/promise-spawn | 9.0.1 |
| @npmcli/query | 5.0.0 |
| @npmcli/redact | 4.0.0 |
| @npmcli/run-script | 10.0.3 |
| @sigstore/bundle | 4.0.0 |
| @sigstore/core | 3.1.0 |
| @sigstore/protobuf-specs | 0.5.0 |
| @sigstore/sign | 4.1.0 |
| @sigstore/tuf | 4.0.1 |
| @sigstore/verify | 3.1.0 |
| @tufjs/canonical-json | 2.0.0 |
| @tufjs/models | 4.1.0 |
| abbrev | 4.0.0 |
| adduser | 3.134 |
| agent-base | 7.1.4 |
| ansi-regex | 5.0.1 |
| aproba | 2.1.0 |
| apt | 2.6.1 |
| archy | 1.0.0 |
| autoconf | 2.71-3 |
| automake | 1:1.16.5-1.3 |
| autotools-dev | 20220109.1 |
| base-files | 12.4+deb12u13 |
| base-passwd | 3.6.1 |
| bash | 5.2.15-2+b10 |
| bin-links | 6.0.0 |
| binary-extensions | 3.1.0 |
| binutils | 2.40-2 |
| binutils-aarch64-linux-gnu | 2.40-2 |
| binutils-common | 2.40-2 |
| bsdutils | 1:2.38.1-5+deb12u3 |
| bzip2 | 1.0.8-5+b1 |
| ca-certificates | 20230311+deb12u1 |
| cacache | 20.0.3 |
| chalk | 5.6.2 |
| chownr | 3.0.0 |
| ci-info | 4.3.1 |
| cidr-regex | 5.0.1 |
| cli-columns | 4.0.0 |
| cmd-shim | 8.0.0 |
| comerr-dev | 2.1-1.47.0-2+b2 |
| common-ancestor-path | 2.0.0 |
| corepack | 0.34.6 |
| coreutils | 9.1-1 |
| cpp | 4:12.2.0-3 |
| cpp-12 | 12.2.0-14+deb12u1 |
| cssesc | 3.0.0 |
| curl | 7.88.1-10+deb12u14 |
| dash | 0.5.12-2 |
| debconf | 1.5.82 |
| debian-archive-keyring | 2023.3+deb12u2 |
| debianutils | 5.7-0.5~deb12u1 |
| debug | 4.4.3 |
| default-libmysqlclient-dev | 1.1.0 |
| diff | 8.0.3 |
| diffutils | 1:3.8-4 |
| dirmngr | 2.2.40-1.1+deb12u2 |
| dpkg | 1.21.22 |
| dpkg-dev | 1.21.22 |
| e2fsprogs | 1.47.0-2+b2 |
| emoji-regex | 8.0.0 |
| encoding | 0.1.13 |
| env-paths | 2.2.1 |
| err-code | 2.0.3 |
| exponential-backoff | 3.1.3 |
| fastest-levenshtein | 1.0.16 |
| fdir | 6.5.0 |
| file | 1:5.44-3 |
| findutils | 4.9.0-4 |
| fontconfig | 2.14.1-4 |
| fontconfig-config | 2.14.1-4 |
| fonts-dejavu-core | 2.37-6 |
| fs-minipass | 3.0.3 |
| g++ | 4:12.2.0-3 |
| g++-12 | 12.2.0-14+deb12u1 |
| gcc | 4:12.2.0-3 |
| gcc-12 | 12.2.0-14+deb12u1 |
| gcc-12-base | 12.2.0-14+deb12u1 |
| gir1.2-freedesktop | 1.74.0-3 |
| gir1.2-gdkpixbuf-2.0 | 2.42.10+dfsg-1+deb12u3 |
| gir1.2-glib-2.0 | 1.74.0-3 |
| gir1.2-rsvg-2.0 | 2.54.7+dfsg-1~deb12u1 |
| git | 1:2.39.5-0+deb12u3 |
| git-man | 1:2.39.5-0+deb12u3 |
| glob | 13.0.0 |
| gnupg | 2.2.40-1.1+deb12u2 |
| gnupg-l10n | 2.2.40-1.1+deb12u2 |
| gnupg-utils | 2.2.40-1.1+deb12u2 |
| gpg | 2.2.40-1.1+deb12u2 |
| gpg-agent | 2.2.40-1.1+deb12u2 |
| gpg-wks-client | 2.2.40-1.1+deb12u2 |
| gpg-wks-server | 2.2.40-1.1+deb12u2 |
| gpgconf | 2.2.40-1.1+deb12u2 |
| gpgsm | 2.2.40-1.1+deb12u2 |
| gpgv | 2.2.40-1.1+deb12u2 |
| graceful-fs | 4.2.11 |
| grep | 3.8-5 |
| gzip | 1.12-1 |
| hicolor-icon-theme | 0.17-2 |
| hosted-git-info | 9.0.2 |
| hostname | 3.23+nmu1 |
| http-cache-semantics | 4.2.0 |
| http-proxy-agent | 7.0.2 |
| https-proxy-agent | 7.0.6 |
| iconv-lite | 0.6.3 |
| icu-devtools | 72.1-3+deb12u1 |
| ignore-walk | 8.0.0 |
| imagemagick | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| imagemagick-6-common | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| imagemagick-6.q16 | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| imurmurhash | 0.1.4 |
| ini | 6.0.0 |
| init-package-json | 8.2.4 |
| init-system-helpers | 1.65.2+deb12u1 |
| ip-address | 10.1.0 |
| ip-regex | 5.0.0 |
| is-cidr | 6.0.1 |
| is-fullwidth-code-point | 3.0.0 |
| isexe | 3.1.1 |
| json-parse-even-better-errors | 5.0.0 |
| json-stringify-nice | 1.1.4 |
| jsonparse | 1.3.1 |
| just-diff | 6.0.2 |
| just-diff-apply | 5.5.0 |
| krb5-multidev | 1.20.1-2+deb12u4 |
| libacl1 | 2.3.1-3 |
| libaom3 | 3.6.0-1+deb12u2 |
| libapr1 | 1.7.2-3+deb12u1 |
| libaprutil1 | 1.6.3-1 |
| libapt-pkg6.0 | 2.6.1 |
| libasan8 | 12.2.0-14+deb12u1 |
| libassuan0 | 2.5.5-5 |
| libatomic1 | 12.2.0-14+deb12u1 |
| libattr1 | 1:2.5.1-4 |
| libaudit-common | 1:3.0.9-1 |
| libaudit1 | 1:3.0.9-1 |
| libbinutils | 2.40-2 |
| libblkid-dev | 2.38.1-5+deb12u3 |
| libblkid1 | 2.38.1-5+deb12u3 |
| libbrotli-dev | 1.0.9-2+b6 |
| libbrotli1 | 1.0.9-2+b6 |
| libbsd0 | 0.11.7-2 |
| libbz2-1.0 | 1.0.8-5+b1 |
| libbz2-dev | 1.0.8-5+b1 |
| libc-bin | 2.36-9+deb12u13 |
| libc-dev-bin | 2.36-9+deb12u13 |
| libc6 | 2.36-9+deb12u13 |
| libc6-dev | 2.36-9+deb12u13 |
| libcairo-gobject2 | 1.16.0-7 |
| libcairo-script-interpreter2 | 1.16.0-7 |
| libcairo2 | 1.16.0-7 |
| libcairo2-dev | 1.16.0-7 |
| libcap-ng0 | 0.8.3-1+b3 |
| libcap2 | 1:2.66-4+deb12u2+b2 |
| libcbor0.8 | 0.8.0-2+b1 |
| libcc1-0 | 12.2.0-14+deb12u1 |
| libcom-err2 | 1.47.0-2+b2 |
| libcrypt-dev | 1:4.4.33-2 |
| libcrypt1 | 1:4.4.33-2 |
| libctf-nobfd0 | 2.40-2 |
| libctf0 | 2.40-2 |
| libcurl3-gnutls | 7.88.1-10+deb12u14 |
| libcurl4 | 7.88.1-10+deb12u14 |
| libcurl4-openssl-dev | 7.88.1-10+deb12u14 |
| libdatrie1 | 0.2.13-2+b1 |
| libdav1d6 | 1.0.0-2+deb12u1 |
| libdb-dev | 5.3.2 |
| libdb5.3 | 5.3.28+dfsg2-1 |
| libdb5.3-dev | 5.3.28+dfsg2-1 |
| libde265-0 | 1.0.11-1+deb12u2 |
| libdebconfclient0 | 0.270 |
| libdeflate-dev | 1.14-1 |
| libdeflate0 | 1.14-1 |
| libdjvulibre-dev | 3.5.28-2.2~deb12u1 |
| libdjvulibre-text | 3.5.28-2.2~deb12u1 |
| libdjvulibre21 | 3.5.28-2.2~deb12u1 |
| libdpkg-perl | 1.21.22 |
| libedit2 | 3.1-20221030-2 |
| libelf1 | 0.188-2.1 |
| liberror-perl | 0.17029-2 |
| libevent-2.1-7 | 2.1.12-stable-8 |
| libevent-core-2.1-7 | 2.1.12-stable-8 |
| libevent-dev | 2.1.12-stable-8 |
| libevent-extra-2.1-7 | 2.1.12-stable-8 |
| libevent-openssl-2.1-7 | 2.1.12-stable-8 |
| libevent-pthreads-2.1-7 | 2.1.12-stable-8 |
| libexif-dev | 0.6.24-1+b1 |
| libexif12 | 0.6.24-1+b1 |
| libexpat1 | 2.5.0-1+deb12u2 |
| libexpat1-dev | 2.5.0-1+deb12u2 |
| libext2fs2 | 1.47.0-2+b2 |
| libffi-dev | 3.4.4-1 |
| libffi8 | 3.4.4-1 |
| libfftw3-double3 | 3.3.10-1 |
| libfido2-1 | 1.12.0-2+b1 |
| libfontconfig-dev | 2.14.1-4 |
| libfontconfig1 | 2.14.1-4 |
| libfreetype-dev | 2.12.1+dfsg-5+deb12u4 |
| libfreetype6 | 2.12.1+dfsg-5+deb12u4 |
| libfreetype6-dev | 2.12.1+dfsg-5+deb12u4 |
| libfribidi0 | 1.0.8-2.1 |
| libgcc-12-dev | 12.2.0-14+deb12u1 |
| libgcc-s1 | 12.2.0-14+deb12u1 |
| libgcrypt20 | 1.10.1-3 |
| libgdbm-compat4 | 1.23-3 |
| libgdbm-dev | 1.23-3 |
| libgdbm6 | 1.23-3 |
| libgdk-pixbuf-2.0-0 | 2.42.10+dfsg-1+deb12u3 |
| libgdk-pixbuf-2.0-dev | 2.42.10+dfsg-1+deb12u3 |
| libgdk-pixbuf2.0-bin | 2.42.10+dfsg-1+deb12u3 |
| libgdk-pixbuf2.0-common | 2.42.10+dfsg-1+deb12u3 |
| libgirepository-1.0-1 | 1.74.0-3 |
| libglib2.0-0 | 2.74.6-2+deb12u8 |
| libglib2.0-bin | 2.74.6-2+deb12u8 |
| libglib2.0-data | 2.74.6-2+deb12u8 |
| libglib2.0-dev | 2.74.6-2+deb12u8 |
| libglib2.0-dev-bin | 2.74.6-2+deb12u8 |
| libgmp-dev | 2:6.2.1+dfsg1-1.1 |
| libgmp10 | 2:6.2.1+dfsg1-1.1 |
| libgmpxx4ldbl | 2:6.2.1+dfsg1-1.1 |
| libgnutls30 | 3.7.9-2+deb12u5 |
| libgomp1 | 12.2.0-14+deb12u1 |
| libgpg-error0 | 1.46-1 |
| libgprofng0 | 2.40-2 |
| libgraphite2-3 | 1.3.14-1 |
| libgssapi-krb5-2 | 1.20.1-2+deb12u4 |
| libgssrpc4 | 1.20.1-2+deb12u4 |
| libharfbuzz0b | 6.0.0+dfsg-3 |
| libheif1 | 1.15.1-1+deb12u1 |
| libhogweed6 | 3.8.1-2 |
| libhwasan0 | 12.2.0-14+deb12u1 |
| libice-dev | 2:1.0.10-1 |
| libice6 | 2:1.0.10-1 |
| libicu-dev | 72.1-3+deb12u1 |
| libicu72 | 72.1-3+deb12u1 |
| libidn2-0 | 2.3.3-1+b1 |
| libimath-3-1-29 | 3.1.6-1 |
| libimath-dev | 3.1.6-1 |
| libisl23 | 0.25-1.1 |
| libitm1 | 12.2.0-14+deb12u1 |
| libjansson4 | 2.14-2 |
| libjbig-dev | 2.1-6.1 |
| libjbig0 | 2.1-6.1 |
| libjpeg-dev | 1:2.1.5-2 |
| libjpeg62-turbo | 1:2.1.5-2 |
| libjpeg62-turbo-dev | 1:2.1.5-2 |
| libk5crypto3 | 1.20.1-2+deb12u4 |
| libkadm5clnt-mit12 | 1.20.1-2+deb12u4 |
| libkadm5srv-mit12 | 1.20.1-2+deb12u4 |
| libkdb5-10 | 1.20.1-2+deb12u4 |
| libkeyutils1 | 1.6.3-2 |
| libkrb5-3 | 1.20.1-2+deb12u4 |
| libkrb5-dev | 1.20.1-2+deb12u4 |
| libkrb5support0 | 1.20.1-2+deb12u4 |
| libksba8 | 1.6.3-2 |
| liblcms2-2 | 2.14-2 |
| liblcms2-dev | 2.14-2 |
| libldap-2.5-0 | 2.5.13+dfsg-5 |
| liblerc-dev | 4.0.0+ds-2 |
| liblerc4 | 4.0.0+ds-2 |
| liblqr-1-0 | 0.4.2-2.1 |
| liblqr-1-0-dev | 0.4.2-2.1 |
| liblsan0 | 12.2.0-14+deb12u1 |
| libltdl-dev | 2.4.7-7~deb12u1 |
| libltdl7 | 2.4.7-7~deb12u1 |
| liblz4-1 | 1.9.4-1 |
| liblzma-dev | 5.4.1-1 |
| liblzma5 | 5.4.1-1 |
| liblzo2-2 | 2.10-2 |
| libmagic-mgc | 1:5.44-3 |
| libmagic1 | 1:5.44-3 |
| libmagickcore-6-arch-config | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickcore-6-headers | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickcore-6.q16-6 | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickcore-6.q16-6-extra | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickcore-6.q16-dev | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickcore-dev | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickwand-6-headers | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickwand-6.q16-6 | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickwand-6.q16-dev | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmagickwand-dev | 8:6.9.11.60+dfsg-1.6+deb12u6 |
| libmariadb-dev | 1:10.11.14-0+deb12u2 |
| libmariadb-dev-compat | 1:10.11.14-0+deb12u2 |
| libmariadb3 | 1:10.11.14-0+deb12u2 |
| libmaxminddb-dev | 1.7.1-1 |
| libmaxminddb0 | 1.7.1-1 |
| libmd0 | 1.0.4-2 |
| libmount-dev | 2.38.1-5+deb12u3 |
| libmount1 | 2.38.1-5+deb12u3 |
| libmpc3 | 1.3.1-1 |
| libmpfr6 | 4.2.0-1 |
| libncurses-dev | 6.4-4 |
| libncurses5-dev | 6.4-4 |
| libncurses6 | 6.4-4 |
| libncursesw5-dev | 6.4-4 |
| libncursesw6 | 6.4-4 |
| libnettle8 | 3.8.1-2 |
| libnghttp2-14 | 1.52.0-1+deb12u2 |
| libnpmaccess | 10.0.3 |
| libnpmdiff | 8.0.13 |
| libnpmexec | 10.1.12 |
| libnpmfund | 7.0.13 |
| libnpmorg | 8.0.1 |
| libnpmpack | 9.0.13 |
| libnpmpublish | 11.1.3 |
| libnpmsearch | 9.0.1 |
| libnpmteam | 8.0.2 |
| libnpmversion | 8.0.3 |
| libnpth0 | 1.6-3 |
| libnsl-dev | 1.3.0-2 |
| libnsl2 | 1.3.0-2 |
| libnuma1 | 2.0.16-1 |
| libopenexr-3-1-30 | 3.1.5-5 |
| libopenexr-dev | 3.1.5-5 |
| libopenjp2-7 | 2.5.0-2+deb12u2 |
| libopenjp2-7-dev | 2.5.0-2+deb12u2 |
| libp11-kit0 | 0.24.1-2 |
| libpam-modules | 1.5.2-6+deb12u2 |
| libpam-modules-bin | 1.5.2-6+deb12u2 |
| libpam-runtime | 1.5.2-6+deb12u2 |
| libpam0g | 1.5.2-6+deb12u2 |
| libpango-1.0-0 | 1.50.12+ds-1 |
| libpangocairo-1.0-0 | 1.50.12+ds-1 |
| libpangoft2-1.0-0 | 1.50.12+ds-1 |
| libpcre2-16-0 | 10.42-1 |
| libpcre2-32-0 | 10.42-1 |
| libpcre2-8-0 | 10.42-1 |
| libpcre2-dev | 10.42-1 |
| libpcre2-posix3 | 10.42-1 |
| libperl5.36 | 5.36.0-7+deb12u3 |
| libpixman-1-0 | 0.42.2-1 |
| libpixman-1-dev | 0.42.2-1 |
| libpkgconf3 | 1.8.1-1 |
| libpng-dev | 1.6.39-2+deb12u1 |
| libpng16-16 | 1.6.39-2+deb12u1 |
| libpq-dev | 15.15-0+deb12u1 |
| libpq5 | 15.15-0+deb12u1 |
| libproc2-0 | 2:4.0.2-3 |
| libpsl5 | 0.21.2-1 |
| libpthread-stubs0-dev | 0.4-1 |
| libpython3-stdlib | 3.11.2-1+b1 |
| libpython3.11-minimal | 3.11.2-6+deb12u6 |
| libpython3.11-stdlib | 3.11.2-6+deb12u6 |
| libreadline-dev | 8.2-1.3 |
| libreadline8 | 8.2-1.3 |
| librsvg2-2 | 2.54.7+dfsg-1~deb12u1 |
| librsvg2-common | 2.54.7+dfsg-1~deb12u1 |
| librsvg2-dev | 2.54.7+dfsg-1~deb12u1 |
| librtmp1 | 2.4+20151223.gitfa8646d.1-2+b2 |
| libsasl2-2 | 2.1.28+dfsg-10 |
| libsasl2-modules-db | 2.1.28+dfsg-10 |
| libseccomp2 | 2.5.4-1+deb12u1 |
| libselinux1 | 3.4-1+b6 |
| libselinux1-dev | 3.4-1+b6 |
| libsemanage-common | 3.4-1 |
| libsemanage2 | 3.4-1+b5 |
| libsepol-dev | 3.4-2.1 |
| libsepol2 | 3.4-2.1 |
| libserf-1-1 | 1.3.9-11 |
| libsm-dev | 2:1.2.3-1 |
| libsm6 | 2:1.2.3-1 |
| libsmartcols1 | 2.38.1-5+deb12u3 |
| libsqlite3-0 | 3.40.1-2+deb12u2 |
| libsqlite3-dev | 3.40.1-2+deb12u2 |
| libss2 | 1.47.0-2+b2 |
| libssh2-1 | 1.10.0-3+b1 |
| libssl-dev | 3.0.18-1~deb12u2 |
| libssl3 | 3.0.18-1~deb12u2 |
| libstdc++-12-dev | 12.2.0-14+deb12u1 |
| libstdc++6 | 12.2.0-14+deb12u1 |
| libsvn1 | 1.14.2-4+deb12u1 |
| libsystemd0 | 252.39-1~deb12u1 |
| libtasn1-6 | 4.19.0-2+deb12u1 |
| libthai-data | 0.1.29-1 |
| libthai0 | 0.1.29-1 |
| libtiff-dev | 4.5.0-6+deb12u3 |
| libtiff6 | 4.5.0-6+deb12u3 |
| libtiffxx6 | 4.5.0-6+deb12u3 |
| libtinfo6 | 6.4-4 |
| libtirpc-common | 1.3.3+ds-1 |
| libtirpc-dev | 1.3.3+ds-1 |
| libtirpc3 | 1.3.3+ds-1 |
| libtool | 2.4.7-7~deb12u1 |
| libtsan2 | 12.2.0-14+deb12u1 |
| libubsan1 | 12.2.0-14+deb12u1 |
| libudev1 | 252.39-1~deb12u1 |
| libunistring2 | 1.0-2 |
| libutf8proc2 | 2.8.0-1 |
| libuuid1 | 2.38.1-5+deb12u3 |
| libwebp-dev | 1.2.4-0.2+deb12u1 |
| libwebp7 | 1.2.4-0.2+deb12u1 |
| libwebpdemux2 | 1.2.4-0.2+deb12u1 |
| libwebpmux3 | 1.2.4-0.2+deb12u1 |
| libwmf-0.2-7 | 0.2.12-5.1 |
| libwmf-dev | 0.2.12-5.1 |
| libwmflite-0.2-7 | 0.2.12-5.1 |
| libx11-6 | 2:1.8.4-2+deb12u2 |
| libx11-data | 2:1.8.4-2+deb12u2 |
| libx11-dev | 2:1.8.4-2+deb12u2 |
| libx265-199 | 3.5-2+b1 |
| libxau-dev | 1:1.0.9-1 |
| libxau6 | 1:1.0.9-1 |
| libxcb-render0 | 1.15-1 |
| libxcb-render0-dev | 1.15-1 |
| libxcb-shm0 | 1.15-1 |
| libxcb-shm0-dev | 1.15-1 |
| libxcb1 | 1.15-1 |
| libxcb1-dev | 1.15-1 |
| libxdmcp-dev | 1:1.1.2-3 |
| libxdmcp6 | 1:1.1.2-3 |
| libxext-dev | 2:1.3.4-1+b1 |
| libxext6 | 2:1.3.4-1+b1 |
| libxml2 | 2.9.14+dfsg-1.3~deb12u5 |
| libxml2-dev | 2.9.14+dfsg-1.3~deb12u5 |
| libxrender-dev | 1:0.9.10-1.1 |
| libxrender1 | 1:0.9.10-1.1 |
| libxslt1-dev | 1.1.35-1+deb12u3 |
| libxslt1.1 | 1.1.35-1+deb12u3 |
| libxt-dev | 1:1.2.1-1.1 |
| libxt6 | 1:1.2.1-1.1 |
| libxxhash0 | 0.8.1-1 |
| libyaml-0-2 | 0.2.5-1 |
| libyaml-dev | 0.2.5-1 |
| libzstd-dev | 1.5.4+dfsg2-5 |
| libzstd1 | 1.5.4+dfsg2-5 |
| linux-libc-dev | 6.1.159-1 |
| login | 1:4.13+dfsg1-1+deb12u2 |
| logsave | 1.47.0-2+b2 |
| lru-cache | 11.2.4 |
| m4 | 1.4.19-3 |
| make | 4.3-4.1 |
| make-fetch-happen | 15.0.3 |
| mariadb-common | 1:10.11.14-0+deb12u2 |
| mawk | 1.3.4.20200120-3.1 |
| media-types | 10.0.0 |
| mercurial | 6.3.2 |
| mercurial | 6.3.2-1+deb12u1 |
| mercurial-common | 6.3.2-1+deb12u1 |
| minimatch | 10.1.1 |
| minipass | 3.3.6 |
| minipass | 7.1.2 |
| minipass-collect | 2.0.1 |
| minipass-fetch | 5.0.0 |
| minipass-flush | 1.0.5 |
| minipass-pipeline | 1.2.4 |
| minipass-sized | 1.0.3 |
| minizlib | 3.1.0 |
| mount | 2.38.1-5+deb12u3 |
| ms | 2.1.3 |
| mute-stream | 3.0.0 |
| mysql-common | 5.8+1.1.0 |
| ncurses-base | 6.4-4 |
| ncurses-bin | 6.4-4 |
| negotiator | 1.0.0 |
| netbase | 6.4 |
| node | 24.13.1 |
| node-gyp | 12.1.0 |
| nopt | 9.0.0 |
| npm | 11.8.0 |
| npm-audit-report | 7.0.0 |
| npm-bundled | 5.0.0 |
| npm-install-checks | 8.0.0 |
| npm-normalize-package-bin | 5.0.0 |
| npm-package-arg | 13.0.2 |
| npm-packlist | 10.0.3 |
| npm-pick-manifest | 11.0.3 |
| npm-profile | 12.0.1 |
| npm-registry-fetch | 19.1.1 |
| npm-user-validate | 4.0.0 |
| openssh-client | 1:9.2p1-2+deb12u7 |
| openssl | 3.0.18-1~deb12u2 |
| p-map | 7.0.4 |
| pacote | 21.0.4 |
| parse-conflict-json | 5.0.1 |
| passwd | 1:4.13+dfsg1-1+deb12u2 |
| patch | 2.7.6-7 |
| path-scurry | 2.0.1 |
| perl | 5.36.0-7+deb12u3 |
| perl-base | 5.36.0-7+deb12u3 |
| perl-modules-5.36 | 5.36.0-7+deb12u3 |
| picomatch | 4.0.3 |
| pinentry-curses | 1.2.1-1 |
| pkg-config | 1.8.1-1 |
| pkgconf | 1.8.1-1 |
| pkgconf-bin | 1.8.1-1 |
| postcss-selector-parser | 7.1.1 |
| proc-log | 6.1.0 |
| procps | 2:4.0.2-3 |
| proggy | 4.0.0 |
| promise-all-reject-late | 1.0.1 |
| promise-call-limit | 3.0.2 |
| promise-retry | 2.0.1 |
| promzard | 3.0.1 |
| python3 | 3.11.2-1+b1 |
| python3-distutils | 3.11.2-3 |
| python3-lib2to3 | 3.11.2-3 |
| python3-minimal | 3.11.2-1+b1 |
| python3.11 | 3.11.2-6+deb12u6 |
| python3.11-minimal | 3.11.2-6+deb12u6 |
| qrcode-terminal | 0.12.0 |
| read | 5.0.1 |
| read-cmd-shim | 6.0.0 |
| readline-common | 8.2-1.3 |
| retry | 0.12.0 |
| rpcsvc-proto | 1.4.3-1 |
| safer-buffer | 2.1.2 |
| sed | 4.9-1 |
| semver | 7.7.3 |
| sensible-utils | 0.0.17+nmu1 |
| shared-mime-info | 2.2-1 |
| signal-exit | 4.1.0 |
| sigstore | 4.1.0 |
| smart-buffer | 4.2.0 |
| socks | 2.8.7 |
| socks-proxy-agent | 8.0.5 |
| spdx-correct | 3.2.0 |
| spdx-exceptions | 2.5.0 |
| spdx-expression-parse | 3.0.1 |
| spdx-expression-parse | 4.0.0 |
| spdx-license-ids | 3.0.22 |
| sq | 0.27.0-2+b1 |
| ssri | 13.0.0 |
| string-width | 4.2.3 |
| strip-ansi | 6.0.1 |
| subversion | 1.14.2-4+deb12u1 |
| supports-color | 10.2.2 |
| sysvinit-utils | 3.06-4 |
| tar | 1.34+dfsg-1.2+deb12u1 |
| tar | 7.5.4 |
| text-table | 0.2.0 |
| tiny-relative-date | 2.0.2 |
| tinyglobby | 0.2.15 |
| treeverse | 3.0.0 |
| tuf-js | 4.1.0 |
| tzdata | 2025b-0+deb12u2 |
| ucf | 3.0043+nmu1+deb12u1 |
| unique-filename | 5.0.0 |
| unique-slug | 6.0.0 |
| unzip | 6.0-28 |
| usr-is-merged | 37~deb12u1 |
| util-deprecate | 1.0.2 |
| util-linux | 2.38.1-5+deb12u3 |
| util-linux-extra | 2.38.1-5+deb12u3 |
| uuid-dev | 2.38.1-5+deb12u3 |
| validate-npm-package-license | 3.0.4 |
| validate-npm-package-name | 7.0.2 |
| walk-up-path | 4.0.0 |
| wget | 1.21.3-1+deb12u1 |
| which | 6.0.0 |
| write-file-atomic | 7.0.0 |
| x11-common | 1:7.7+23 |
| x11proto-core-dev | 2022.1-1 |
| x11proto-dev | 2022.1-1 |
| xorg-sgml-doctools | 1:1.11-1.1 |
| xtrans-dev | 1.4.0-1 |
| xz-utils | 5.4.1-1 |
| yallist | 4.0.0 |
| yallist | 5.0.0 |
| yarn | 1.22.22 |
| zlib1g | 1:1.2.13.dfsg-1 |
| zlib1g-dev | 1:1.2.13.dfsg-1 |
</details>

</details>

<!-- END: dock-docs:comparison -->
