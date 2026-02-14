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

<!-- END: docker-docs -->

Footer info.
