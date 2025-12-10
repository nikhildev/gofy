# Multi-stage build for minimal image size
# Stage 1: Build the Go application
FROM golang:alpine AS builder

# Set working directory
WORKDIR /build

# Copy go mod files (build context is parent directory)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the web application source
COPY web/ ./web/

# Build the application
WORKDIR /build/web
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o weather-app .

# Stage 2: Create the runtime image
FROM alpine:latest

# Install ca-certificates for HTTPS requests to external APIs
RUN apk update && \
    apk --no-cache add ca-certificates tzdata && \
    rm -rf /var/cache/apk/*

# Create a non-root user
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

# Set working directory
WORKDIR /app

# Copy the binary from builder
COPY --from=builder /build/web/weather-app .

# Copy templates
COPY --from=builder /build/web/templates ./templates/

# Change ownership to non-root user
RUN chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

# Expose port 8080
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/?location=Espoo || exit 1

# Run the application
CMD ["./weather-app"]

