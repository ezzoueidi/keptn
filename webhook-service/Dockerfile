# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.19.5-alpine3.16 as builder

WORKDIR /go/src/github.com/keptn/keptn/webhook-service

# Force the go compiler to use modules
ENV GO111MODULE=on
ENV BUILDFLAGS=""
ENV GOPROXY=https://proxy.golang.org

RUN apk add --no-cache gcc libc-dev git

# Copy `go.mod` for definitions and `go.sum` to invalidate the next layer
# in case of a change in the dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy local code to the container image.
COPY . .

# `skaffold debug` sets SKAFFOLD_GO_GCFLAGS to disable compiler optimizations
ARG SKAFFOLD_GO_GCFLAGS

# Build the command inside the container.
# (You may fetch or manage dependencies here, either manually or with a tool like "godep".)
RUN GOOS=linux go build -ldflags '-linkmode=external' -gcflags="${SKAFFOLD_GO_GCFLAGS}" -v -o webhook-service

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3.17 as production
LABEL org.opencontainers.image.source="https://github.com/keptn/keptn" \
    org.opencontainers.image.url="https://keptn.sh" \
    org.opencontainers.image.title="Keptn Webhook Service" \
    org.opencontainers.image.vendor="Keptn" \
    org.opencontainers.image.documentation="https://keptn.sh/docs/" \
    org.opencontainers.image.licenses="Apache-2.0"

# we need to install ca-certificates and libc6-compat for go programs to work properly
RUN apk add --no-cache ca-certificates libc6-compat curl

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/src/github.com/keptn/keptn/webhook-service/webhook-service /webhook-service

EXPOSE 8080

# required for external tools to detect this as a go binary
ENV GOTRACEBACK=all

RUN adduser -D nonroot -u 65532
USER nonroot

ARG version=develop
LABEL org.opencontainers.image.version="${version}"

# Run the web service on container startup.
CMD ["/webhook-service"]
