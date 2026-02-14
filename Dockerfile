# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache curl tar

# 1. Build the Go application
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o docker-docs .

# 2. Download and prepare external tools
WORKDIR /tmp/tools

# Install syft
RUN curl -sSfL https://raw.githubusercontent.com/anchore/syft/main/install.sh | sh -s -- -b /tmp/tools

# Install grype
RUN curl -sSfL https://raw.githubusercontent.com/anchore/grype/main/install.sh | sh -s -- -b /tmp/tools

# Install dive
RUN curl -L https://github.com/wagoodman/dive/releases/download/v0.12.0/dive_0.12.0_linux_amd64.tar.gz -o dive.tar.gz && \
    tar -xzf dive.tar.gz dive && \
    rm dive.tar.gz

# Get Docker CLI (static binary)
# We need docker to run 'docker inspect'.
ENV DOCKER_VERSION=24.0.5
RUN curl -fsSLO https://download.docker.com/linux/static/stable/x86_64/docker-${DOCKER_VERSION}.tgz && \
    tar xzvf docker-${DOCKER_VERSION}.tgz --strip 1 -C /tmp/tools docker/docker && \
    rm docker-${DOCKER_VERSION}.tgz

# Final Stage: Distroless
# We use 'static:nonroot' for security, but we need to ensure our binary and tools work without a shell.
# Note: 'static' image does NOT contain a shell or libc.
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /

# Copy our Go binary
COPY --from=builder /app/docker-docs /usr/local/bin/docker-docs

# Copy external tools
COPY --from=builder /tmp/tools/syft /usr/local/bin/syft
COPY --from=builder /tmp/tools/grype /usr/local/bin/grype
COPY --from=builder /tmp/tools/dive /usr/local/bin/dive
COPY --from=builder /tmp/tools/docker /usr/local/bin/docker

# Add /usr/local/bin to PATH (Distroless might not have this set by default for our usage)
ENV PATH="/usr/local/bin:${PATH}"

ENTRYPOINT ["docker-docs"]
