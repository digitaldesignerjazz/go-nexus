# Multi-stage Dockerfile for go-nexus orchestrator
# Produces a small, secure runtime image

FROM golang:1.21-alpine AS builder

WORKDIR /src

# Install build dependencies if needed (git, etc.)
RUN apk add --no-cache git ca-certificates tzdata

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /go-nexus .

# --- Runtime stage ---
FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata docker-cli docker-compose

WORKDIR /app

# Copy binary
COPY --from=builder /go-nexus /usr/local/bin/go-nexus

# Create non-root user (recommended)
RUN addgroup -S nexus && adduser -S nexus -G nexus
USER nexus

ENTRYPOINT ["go-nexus"]
CMD ["help"]

# Healthcheck example (uncomment when doctor supports exit codes properly)
# HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
#   CMD go-nexus doctor || exit 1

# Build:
#   docker build -t go-nexus:latest .
#
# Run:
#   docker run --rm go-nexus:latest doctor
#   docker run -it --network host go-nexus:latest start --component=mesh