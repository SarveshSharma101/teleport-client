# Build the manager binary
FROM golang:1.19 as builder
ARG TARGETOS
ARG TARGETARCH

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY . .

# Build
# the GOARCH has not a default value to allow the binary be built according to the host where the command
# was called. For example, if we call make docker-build in a local env which has the Apple Silicon M1 SO
# the docker BUILDPLATFORM arg will be linux/arm64 when for Apple x86 it will be linux/amd64. Therefore,
# by leaving it empty we can ensure that the container and binary shipped on it will have the same platform.
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -a -o teleportClient main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM ubuntu:20.04
WORKDIR /
COPY --from=builder /workspace/teleportClient .

RUN mkdir status

RUN apt-get update && apt-get install software-properties-common -y && apt-get install -y gnupg2 && apt-get install -y wget

RUN wget -qO- https://deb.releases.teleport.dev/teleport-pubkey.asc \
| gpg --dearmor > /etc/apt/trusted.gpg.d/teleport.gpg

RUN echo "deb https://deb.releases.teleport.dev/ stable main" > /etc/apt/sources.list.d/teleport.list
RUN apt-get update 
RUN apt-get install teleport 
# USER 65532:65532
ENTRYPOINT ["/teleportClient"]