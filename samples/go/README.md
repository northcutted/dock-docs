# Go Docker Analysis — Template Showcase

This document demonstrates all 6 built-in dock-docs templates applied to the Go sample project.
Each template is shown for both **image** (single image analysis) and **comparison** (multi-image comparison) section types.

---

## Default Template

### Image Analysis
<!-- BEGIN: dock-docs:default-image -->

# 🐳 Docker Image Analysis: golang:1.26-alpine
![Size](https://img.shields.io/static/v1?label=Size&message=68.04+MB&color=blue) ![Layers](https://img.shields.io/static/v1?label=Layers&message=5&color=blue) ![Vulns](https://img.shields.io/static/v1?label=Security&message=3+Vulns+%280+Crit%29&color=green) ![Efficiency](https://img.shields.io/static/v1?label=Efficiency&message=99.9%&color=green)

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

**Base Image:** `Alpine Linux (linux/amd64)`
**Supported Architectures:** `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown`
**Efficiency Score:** 99.9%

### Vulnerabilities
**Last scanned:** 2026-02-18T04:36:13Z

| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| 🟢 0 | 🟢 0 | 🟡 3 | 🟢 0 |

<details>
<summary><strong>👇 Expand Vulnerability Details (3 found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox` | `1.37.0-r30` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox-binsh` | `1.37.0-r30` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `ssl_client` | `1.37.0-r30` |
</details>

<details>
<summary><strong>📦 Installed Packages (29 total)</strong></summary>

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
| cmd/asm | UNKNOWN |
| cmd/cgo | UNKNOWN |
| cmd/compile | UNKNOWN |
| cmd/cover | UNKNOWN |
| cmd/fix | UNKNOWN |
| cmd/go | UNKNOWN |
| cmd/gofmt | UNKNOWN |
| cmd/link | UNKNOWN |
| cmd/preprofile | UNKNOWN |
| cmd/vet | UNKNOWN |
| go | 1.26.0 |
| libapk | 3.0.3-r1 |
| libcrypto3 | 3.5.5-r0 |
| libssl3 | 3.5.5-r0 |
| musl | 1.2.5-r21 |
| musl-utils | 1.2.5-r21 |
| scanelf | 1.3.8-r2 |
| ssl_client | 1.37.0-r30 |
| stdlib | go1.26.0 |
| zlib | 1.3.1-r2 |
</details>

<!-- END: dock-docs:default-image -->

### Comparison
<!-- BEGIN: dock-docs:default-comparison -->

### 🏷️ Supported Tags

| Tag | Size | Vulns | Efficiency | Architectures |
|-----|------|-------|------------|---------------|
| `golang:1.24-alpine` | ![Size](https://img.shields.io/static/v1?label=Size&message=79.44+MB&color=blue) | ![Vulns](https://img.shields.io/static/v1?label=Security&message=3+Vulns+%280+Crit%29&color=green) | 99.9% | `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown` |
| `golang:1.25-alpine` | ![Size](https://img.shields.io/static/v1?label=Size&message=61.35+MB&color=blue) | ![Vulns](https://img.shields.io/static/v1?label=Security&message=3+Vulns+%280+Crit%29&color=green) | 99.9% | `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown` |
| `golang:1.26-alpine` | ![Size](https://img.shields.io/static/v1?label=Size&message=68.04+MB&color=blue) | ![Vulns](https://img.shields.io/static/v1?label=Security&message=3+Vulns+%280+Crit%29&color=green) | 99.9% | `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown` |

<details>
<summary><strong>🔍 Full Report: golang:1.24-alpine</strong></summary>

## 🛡️ Security & Efficiency

**Base Image:** `Alpine Linux (linux/amd64)`
**Supported Architectures:** `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown`
**Efficiency Score:** 99.9%

### Vulnerabilities
**Last scanned:** 2026-02-18T04:37:03Z

| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| 🟢 0 | 🟢 0 | 🟡 3 | 🟢 0 |

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

<details>
<summary><strong>🔍 Full Report: golang:1.25-alpine</strong></summary>

## 🛡️ Security & Efficiency

**Base Image:** `Alpine Linux (linux/amd64)`
**Supported Architectures:** `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown`
**Efficiency Score:** 99.9%

### Vulnerabilities
**Last scanned:** 2026-02-18T04:37:01Z

| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| 🟢 0 | 🟢 0 | 🟡 3 | 🟢 0 |

<details>
<summary><strong>👇 Expand Vulnerability Details (3 found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox` | `1.37.0-r30` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox-binsh` | `1.37.0-r30` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `ssl_client` | `1.37.0-r30` |
</details>

<details>
<summary><strong>📦 Installed Packages (28 total)</strong></summary>

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
| cmd/asm | UNKNOWN |
| cmd/cgo | UNKNOWN |
| cmd/compile | UNKNOWN |
| cmd/cover | UNKNOWN |
| cmd/go | UNKNOWN |
| cmd/gofmt | UNKNOWN |
| cmd/link | UNKNOWN |
| cmd/preprofile | UNKNOWN |
| cmd/vet | UNKNOWN |
| go | 1.25.7 |
| libapk | 3.0.3-r1 |
| libcrypto3 | 3.5.5-r0 |
| libssl3 | 3.5.5-r0 |
| musl | 1.2.5-r21 |
| musl-utils | 1.2.5-r21 |
| scanelf | 1.3.8-r2 |
| ssl_client | 1.37.0-r30 |
| stdlib | go1.25.7 |
| zlib | 1.3.1-r2 |
</details>

</details>

<details>
<summary><strong>🔍 Full Report: golang:1.26-alpine</strong></summary>

## 🛡️ Security & Efficiency

**Base Image:** `Alpine Linux (linux/amd64)`
**Supported Architectures:** `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown`
**Efficiency Score:** 99.9%

### Vulnerabilities
**Last scanned:** 2026-02-18T04:36:56Z

| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| 🟢 0 | 🟢 0 | 🟡 3 | 🟢 0 |

<details>
<summary><strong>👇 Expand Vulnerability Details (3 found)</strong></summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox` | `1.37.0-r30` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox-binsh` | `1.37.0-r30` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `ssl_client` | `1.37.0-r30` |
</details>

<details>
<summary><strong>📦 Installed Packages (29 total)</strong></summary>

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
| cmd/asm | UNKNOWN |
| cmd/cgo | UNKNOWN |
| cmd/compile | UNKNOWN |
| cmd/cover | UNKNOWN |
| cmd/fix | UNKNOWN |
| cmd/go | UNKNOWN |
| cmd/gofmt | UNKNOWN |
| cmd/link | UNKNOWN |
| cmd/preprofile | UNKNOWN |
| cmd/vet | UNKNOWN |
| go | 1.26.0 |
| libapk | 3.0.3-r1 |
| libcrypto3 | 3.5.5-r0 |
| libssl3 | 3.5.5-r0 |
| musl | 1.2.5-r21 |
| musl-utils | 1.2.5-r21 |
| scanelf | 1.3.8-r2 |
| ssl_client | 1.37.0-r30 |
| stdlib | go1.26.0 |
| zlib | 1.3.1-r2 |
</details>

</details>

<!-- END: dock-docs:default-comparison -->

---

## Minimal Template

### Image Analysis
<!-- BEGIN: dock-docs:minimal-image -->
## Configuration: golang:1.26-alpine

### Build Arguments

| Name | Default |
|------|---------|
| `CGO_ENABLED=0` | `` |

### Exposed Ports

| Port | Description |
|------|-------------|
| `8080` | The default port for the Go application |

<!-- END: dock-docs:minimal-image -->

### Comparison
<!-- BEGIN: dock-docs:minimal-comparison -->
### Supported Tags

| Tag | Size | Vulns | Efficiency |
|-----|------|-------|------------|
| `golang:1.24-alpine` | 79.44 MB | 0C / 0H | 99.9% |
| `golang:1.25-alpine` | 61.35 MB | 0C / 0H | 99.9% |
| `golang:1.26-alpine` | 68.04 MB | 0C / 0H | 99.9% |

<!-- END: dock-docs:minimal-comparison -->

---

## Detailed Template

### Image Analysis
<!-- BEGIN: dock-docs:detailed-image -->

# 🐳 Docker Image Analysis: golang:1.26-alpine
![Size](https://img.shields.io/static/v1?label=Size&message=68.04+MB&color=blue) ![Layers](https://img.shields.io/static/v1?label=Layers&message=5&color=blue) ![Vulns](https://img.shields.io/static/v1?label=Security&message=3+Vulns+%280+Crit%29&color=green) ![Efficiency](https://img.shields.io/static/v1?label=Efficiency&message=99.9%&color=green)

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

### Image Metadata

| Property | Value |
|----------|-------|
| **Tag** | `golang:1.26-alpine` |
| **Base Image OS** | `Alpine Linux` |
| **Architecture** | `amd64` |
| **OS** | `linux` |
| **Supported Architectures** | `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown` |
| **Image Size** | 68.04 MB |
| **Total Layers** | 5 |
| **Efficiency Score** | 99.9% |
| **Wasted Space** | 0.47 MB |

### Vulnerability Summary

**Last scanned:** 2026-02-18T04:36:29Z

| Critical | High | Medium | Low | Total |
|:---:|:---:|:---:|:---:|:---:|
| 🟢 0 | 🟢 0 | 🟡 3 | 🟢 0 | 3 |

### Vulnerability Details

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox` | `1.37.0-r30` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox-binsh` | `1.37.0-r30` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `ssl_client` | `1.37.0-r30` |

### Installed Packages (29 total)

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
| cmd/asm | UNKNOWN |
| cmd/cgo | UNKNOWN |
| cmd/compile | UNKNOWN |
| cmd/cover | UNKNOWN |
| cmd/fix | UNKNOWN |
| cmd/go | UNKNOWN |
| cmd/gofmt | UNKNOWN |
| cmd/link | UNKNOWN |
| cmd/preprofile | UNKNOWN |
| cmd/vet | UNKNOWN |
| go | 1.26.0 |
| libapk | 3.0.3-r1 |
| libcrypto3 | 3.5.5-r0 |
| libssl3 | 3.5.5-r0 |
| musl | 1.2.5-r21 |
| musl-utils | 1.2.5-r21 |
| scanelf | 1.3.8-r2 |
| ssl_client | 1.37.0-r30 |
| stdlib | go1.26.0 |
| zlib | 1.3.1-r2 |

<!-- END: dock-docs:detailed-image -->

### Comparison
<!-- BEGIN: dock-docs:detailed-comparison -->
### 🏷️ Supported Tags

| Tag | Size | Vulns | Efficiency | Architectures |
|-----|------|-------|------------|---------------|
| `golang:1.24-alpine` | ![Size](https://img.shields.io/static/v1?label=Size&message=79.44+MB&color=blue) | ![Vulns](https://img.shields.io/static/v1?label=Security&message=3+Vulns+%280+Crit%29&color=green) | 99.9% | `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown` |
| `golang:1.25-alpine` | ![Size](https://img.shields.io/static/v1?label=Size&message=61.35+MB&color=blue) | ![Vulns](https://img.shields.io/static/v1?label=Security&message=3+Vulns+%280+Crit%29&color=green) | 99.9% | `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown` |
| `golang:1.26-alpine` | ![Size](https://img.shields.io/static/v1?label=Size&message=68.04+MB&color=blue) | ![Vulns](https://img.shields.io/static/v1?label=Security&message=3+Vulns+%280+Crit%29&color=green) | 99.9% | `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown` |

---

## 🔍 Full Report: golang:1.24-alpine

### Image Metadata

| Property | Value |
|----------|-------|
| **Tag** | `golang:1.24-alpine` |
| **Base Image OS** | `Alpine Linux` |
| **Architecture** | `amd64` |
| **OS** | `linux` |
| **Supported Architectures** | `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown` |
| **Image Size** | 79.44 MB |
| **Total Layers** | 5 |
| **Efficiency Score** | 99.9% |
| **Wasted Space** | 0.47 MB |

### Vulnerability Summary

**Last scanned:** 2026-02-18T04:37:43Z

| Critical | High | Medium | Low | Total |
|:---:|:---:|:---:|:---:|:---:|
| 🟢 0 | 🟢 0 | 🟡 3 | 🟢 0 | 3 |

### Vulnerability Details

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox` | `1.37.0-r30` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox-binsh` | `1.37.0-r30` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `ssl_client` | `1.37.0-r30` |

### Installed Packages (39 total)

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

---

## 🔍 Full Report: golang:1.25-alpine

### Image Metadata

| Property | Value |
|----------|-------|
| **Tag** | `golang:1.25-alpine` |
| **Base Image OS** | `Alpine Linux` |
| **Architecture** | `amd64` |
| **OS** | `linux` |
| **Supported Architectures** | `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown` |
| **Image Size** | 61.35 MB |
| **Total Layers** | 5 |
| **Efficiency Score** | 99.9% |
| **Wasted Space** | 0.47 MB |

### Vulnerability Summary

**Last scanned:** 2026-02-18T04:37:42Z

| Critical | High | Medium | Low | Total |
|:---:|:---:|:---:|:---:|:---:|
| 🟢 0 | 🟢 0 | 🟡 3 | 🟢 0 | 3 |

### Vulnerability Details

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox` | `1.37.0-r30` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox-binsh` | `1.37.0-r30` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `ssl_client` | `1.37.0-r30` |

### Installed Packages (28 total)

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
| cmd/asm | UNKNOWN |
| cmd/cgo | UNKNOWN |
| cmd/compile | UNKNOWN |
| cmd/cover | UNKNOWN |
| cmd/go | UNKNOWN |
| cmd/gofmt | UNKNOWN |
| cmd/link | UNKNOWN |
| cmd/preprofile | UNKNOWN |
| cmd/vet | UNKNOWN |
| go | 1.25.7 |
| libapk | 3.0.3-r1 |
| libcrypto3 | 3.5.5-r0 |
| libssl3 | 3.5.5-r0 |
| musl | 1.2.5-r21 |
| musl-utils | 1.2.5-r21 |
| scanelf | 1.3.8-r2 |
| ssl_client | 1.37.0-r30 |
| stdlib | go1.25.7 |
| zlib | 1.3.1-r2 |

---

## 🔍 Full Report: golang:1.26-alpine

### Image Metadata

| Property | Value |
|----------|-------|
| **Tag** | `golang:1.26-alpine` |
| **Base Image OS** | `Alpine Linux` |
| **Architecture** | `amd64` |
| **OS** | `linux` |
| **Supported Architectures** | `linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64le, linux/riscv64, linux/s390x, unknown/unknown` |
| **Image Size** | 68.04 MB |
| **Total Layers** | 5 |
| **Efficiency Score** | 99.9% |
| **Wasted Space** | 0.47 MB |

### Vulnerability Summary

**Last scanned:** 2026-02-18T04:37:43Z

| Critical | High | Medium | Low | Total |
|:---:|:---:|:---:|:---:|:---:|
| 🟢 0 | 🟢 0 | 🟡 3 | 🟢 0 | 3 |

### Vulnerability Details

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox` | `1.37.0-r30` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `busybox-binsh` | `1.37.0-r30` |
| [CVE-2025-60876](https://nvd.nist.gov/vuln/detail/CVE-2025-60876) | Medium | `ssl_client` | `1.37.0-r30` |

### Installed Packages (29 total)

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
| cmd/asm | UNKNOWN |
| cmd/cgo | UNKNOWN |
| cmd/compile | UNKNOWN |
| cmd/cover | UNKNOWN |
| cmd/fix | UNKNOWN |
| cmd/go | UNKNOWN |
| cmd/gofmt | UNKNOWN |
| cmd/link | UNKNOWN |
| cmd/preprofile | UNKNOWN |
| cmd/vet | UNKNOWN |
| go | 1.26.0 |
| libapk | 3.0.3-r1 |
| libcrypto3 | 3.5.5-r0 |
| libssl3 | 3.5.5-r0 |
| musl | 1.2.5-r21 |
| musl-utils | 1.2.5-r21 |
| scanelf | 1.3.8-r2 |
| ssl_client | 1.37.0-r30 |
| stdlib | go1.26.0 |
| zlib | 1.3.1-r2 |

---

<!-- END: dock-docs:detailed-comparison -->

---

## Compact Template

### Image Analysis
<!-- BEGIN: dock-docs:compact-image -->
**golang:1.26-alpine** | Size: 68.04 MB | Layers: 5 | Efficiency: 99.9% | Vulns: 0C/0H/3M/0L
| ARG | Default | Req |
|-----|---------|:---:|
| `CGO_ENABLED=0` | `` | ❌ |
| Port |
|------|
| `8080` |

<!-- END: dock-docs:compact-image -->

### Comparison
<!-- BEGIN: dock-docs:compact-comparison -->
| Tag | Size | Vulns | Efficiency |
|-----|------|-------|------------|
| `golang:1.24-alpine` | 79.44 MB | 0C/0H/3M | 99.9% |
| `golang:1.25-alpine` | 61.35 MB | 0C/0H/3M | 99.9% |
| `golang:1.26-alpine` | 68.04 MB | 0C/0H/3M | 99.9% |

<!-- END: dock-docs:compact-comparison -->
