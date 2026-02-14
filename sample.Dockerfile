FROM golang:1.21

# @name: HTTP_PORT
# @description: The port the server listens on
# @default: 8080
# @required: true
ENV HTTP_PORT=8080

# @description: Database connection string
ARG DB_URL

LABEL org.opencontainers.image.authors="Eddie"

EXPOSE 8080
