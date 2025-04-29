###############################################################################
# Stage 1 – build the Go binary
###############################################################################
FROM golang:1.22-alpine AS build

# 1. Enable simple, reproducible builds
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOFLAGS=-buildvcs=false

WORKDIR /src

# 2. Download deps first (for better layer-level caching)
COPY go.mod go.sum ./
RUN go mod download

# 3. Copy the remaining source and compile
COPY . .
RUN go build -o mqtt-router ./cmd/mqtt-router


###############################################################################
# Stage 2 – scratch (distroless) runtime
###############################################################################
FROM gcr.io/distroless/base

# 4. Copy only the static binary (no shell, no libc, ~8 MB total image)
COPY --from=build /src/mqtt-router /mqtt-router

# 5. Non-root for safety; UID 65532 = “nonroot” in distroless
USER 65532:65532

# 6. The binary becomes the container entry-point
ENTRYPOINT ["/mqtt-router"]
