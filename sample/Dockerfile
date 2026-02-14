# syntax=docker/dockerfile:1
# -----------------------------------------------------------------------------
# BASE IMAGE
# We use an old Alpine version (3.14) to GUARANTEE strictly "High/Critical"
# vulnerabilities for Grype to find, testing your badge logic.
# -----------------------------------------------------------------------------
FROM alpine:3.14

# -----------------------------------------------------------------------------
# DYNAMIC ANALYSIS GENERATORS
# -----------------------------------------------------------------------------

# 1. GENERATE PACKAGES (For Syft)
# Installing common tools to populate the "Installed Packages" table.
RUN apk add --no-cache \
    curl \
    git \
    python3 \
    bash

# 2. GENERATE INEFFICIENCY (For Dive)
# We create a 10MB file in one layer and delete it in the next.
# This creates 10MB of "Wasted Bytes" and lowers the "Efficiency Score".
RUN dd if=/dev/zero of=/tmp/large-waste-file bs=1M count=10
RUN rm /tmp/large-waste-file

# -----------------------------------------------------------------------------
# BUILD ARGUMENTS
# -----------------------------------------------------------------------------

# @name: VERSION
# @description: The application version to build.
# @default: 1.0.0
ARG VERSION=1.0.0

# @name: BUILD_DATE
# @description: The date the image was built (RFC3339).
# @required: true
ARG BUILD_DATE

# -----------------------------------------------------------------------------
# ENVIRONMENT VARIABLES
# -----------------------------------------------------------------------------

# @name: APP_ENV
# @description: The environment the app is running in (dev, staging, prod).
# @default: production
ENV APP_ENV="production"

# @name: DATABASE_URL
# @description: Connection string for the primary database.
# @required: true
ENV DATABASE_URL=""

# @name: LOG_LEVEL
# @description: Global logging level. Options: DEBUG, INFO, WARN, ERROR.
# @default: INFO
ENV LOG_LEVEL INFO

# Testing equality vs space separation
# @name: API_TIMEOUT
# @description: Timeout in seconds for external API calls.
ENV API_TIMEOUT=30

# -----------------------------------------------------------------------------
# COMPLEX CASES (The "Torture" Section)
# -----------------------------------------------------------------------------

# @name: FEATURE_FLAGS
# @description: A JSON string of enabled feature flags.
# @default: {"new_ui": false}
ENV FEATURE_FLAGS='{"new_ui": false, "beta_access": true}'

# @name: PATH_additions
# @description: Extensions to the system path.
# @description: Other Stuff
# This is a multi-line instruction test.
ENV PATH="/opt/myapp/bin:$PATH" \
    LD_LIBRARY_PATH="/opt/myapp/lib"

# @name: MARKDOWN_TEST
# @description: This description contains **bold text**, `code blocks`, and [links](https://example.com) to test the renderer.
ENV MARKDOWN_TEST="true"

# @name: MISMATCH_TEST
# @description: Tests if the tool prefers the @default tag over the actual value.
# @default: "documented_value"
ENV MISMATCH_TEST="actual_value"

# @name: UNDOCUMENTED_VAR
# This variable has a comment but no magic @description.
ENV SECRET_SAUCE="szechuan"

# This variable has NO comments at all.
ENV RAW_VAR="raw"

# -----------------------------------------------------------------------------
# METADATA & PORTS
# -----------------------------------------------------------------------------

# @name: org.opencontainers.image.authors
LABEL org.opencontainers.image.authors="platform-team@example.com"

# @name: Multi-Line Label
LABEL org.opencontainers.image.description="This is a really long description \
that spans multiple lines in the Dockerfile \
to test parsing."

# @name: HTTP Service
# @description: The main web server port.
EXPOSE 8080

# @name: Metrics
# @description: Prometheus metrics endpoint.
EXPOSE 9090/tcp

CMD ["/bin/bash"]