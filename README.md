# My Project

Here is the configuration:

<!-- BEGIN: docker-docs -->

| Name | Type | Description | Default | Required |
|------|------|-------------|---------|----------|
| VERSION | ARG | The application version to build. | 1.0.0 | false |
| BUILD_DATE | ARG | The date the image was built (RFC3339). |  | true |
| APP_ENV | ENV | The environment the app is running in (dev, staging, prod). | production | false |
| DATABASE_URL | ENV | Connection string for the primary database. |  | true |
| LOG_LEVEL | ENV | Global logging level. Options: DEBUG, INFO, WARN, ERROR. | INFO | false |
| API_TIMEOUT | ENV | Timeout in seconds for external API calls. | 30 | false |
| FEATURE_FLAGS | ENV | A JSON string of enabled feature flags. | {"new_ui": false} | false |
| PATH_additions | ENV | Extensions to the system path. | /opt/myapp/bin:$PATH | false |
| UNDOCUMENTED_VAR | ENV |  | szechuan | false |
| RAW_VAR | ENV |  | raw | false |
| org.opencontainers.image.authors | LABEL |  | platform-team@example.com | false |
| 8080 | EXPOSE | The main web server port. | 8080 | false |
| 9090/tcp | EXPOSE | Prometheus metrics endpoint. | 9090/tcp | false |

## Image Analysis (my-app:latest)

| Metric | Value |
|--------|-------|
| Size | 7.66 MB |
| Architecture | arm64/linux |
| Efficiency | 100.0% (0.00 MB wasted) |
| Total Layers | 1 |

### Security Summary
Critical: 0 | High: 0 | Medium: 3
<details>
<summary>Vulnerabilities Details (9 found)</summary>

| ID | Severity | Package | Version |
|----|----------|---------|---------|
| CVE-2025-60876 | Medium | busybox | 1.36.1-r20 |
| CVE-2025-60876 | Medium | busybox-binsh | 1.36.1-r20 |
| CVE-2025-60876 | Medium | ssl_client | 1.36.1-r20 |
| CVE-2024-58251 | Low | busybox | 1.36.1-r20 |
| CVE-2024-58251 | Low | busybox-binsh | 1.36.1-r20 |
| CVE-2024-58251 | Low | ssl_client | 1.36.1-r20 |
| CVE-2025-46394 | Low | busybox | 1.36.1-r20 |
| CVE-2025-46394 | Low | busybox-binsh | 1.36.1-r20 |
| CVE-2025-46394 | Low | ssl_client | 1.36.1-r20 |
</details>
<details>
<summary>Packages (15 total)</summary>

| Package | Version |
|---------|---------|
| alpine-baselayout | 3.4.3-r2 |
| alpine-baselayout-data | 3.4.3-r2 |
| alpine-keys | 2.4-r1 |
| apk-tools | 2.14.4-r0 |
| busybox | 1.36.1-r20 |
| busybox-binsh | 1.36.1-r20 |
| ca-certificates-bundle | 20250911-r0 |
| libc-utils | 0.7.2-r5 |
| libcrypto3 | 3.1.8-r1 |
| libssl3 | 3.1.8-r1 |
| musl | 1.2.4_git20230717-r5 |
| musl-utils | 1.2.4_git20230717-r5 |
| scanelf | 1.3.7-r2 |
| ssl_client | 1.36.1-r20 |
| zlib | 1.3.1-r0 |
</details>

<!-- END: docker-docs -->

Footer info.
