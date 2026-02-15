# Docker Docs Samples

Here are examples of documentation generated for common Dockerfiles.

## Python

<!-- BEGIN: docker-docs:python -->

# 🐳 Docker Image Analysis: Dockerfile

## ⚙️ Configuration
### Environment Variables
| Name | Description | Default | Required |
|------|-------------|---------|:--------:|
| `NAME` | The name to greet | `World` | ❌ |
### Exposed Ports
| Port | Description |
|------|-------------|
| `80` | The port where the FastAPI application listens |

<!-- END: docker-docs:python -->

## Node.js

<!-- BEGIN: docker-docs:node -->

# 🐳 Docker Image Analysis: Dockerfile

## ⚙️ Configuration
### Build Arguments
| Name | Description | Default | Required |
|------|-------------|---------|:--------:|
| `NODE_ENV=development` | Environment (development/production) | `development` | ❌ |
### Exposed Ports
| Port | Description |
|------|-------------|
| `3000` | The default port for the Node.js application |

<!-- END: docker-docs:node -->

## Go (Multi-stage)

<!-- BEGIN: docker-docs:go -->

# 🐳 Docker Image Analysis: Dockerfile

## ⚙️ Configuration
### Build Arguments
| Name | Description | Default | Required |
|------|-------------|---------|:--------:|
| `CGO_ENABLED=0` | Enable CGO for building (default false) | `` | ❌ |
### Exposed Ports
| Port | Description |
|------|-------------|
| `8080` | The default port for the Go application |

<!-- END: docker-docs:go -->
