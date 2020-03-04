# Build the manager binary
FROM golang:1.13 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY api/ api/
COPY cmd/ cmd/
COPY pkg/ pkg/

# Test
RUN CGO_ENABLED=0 GO111MODULE=on go test ./cmd/... ./pkg/...
# Build
RUN CGO_ENABLED=0 GO111MODULE=on go build -a -o octopus ./octopus/main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/robot .
USER nonroot:nonroot

ENTRYPOINT ["/octopus"]