# Multi-stage build for minimal image size
# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod files
COPY ../go.mod ../go.sum ./

# Download dependencies
RUN go mod download

# Copy the web application source
COPY . ./web/

# Build the application
WORKDIR /app/web
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o weather-app .

# Stage 2: Create the runtime image
FROM alpine:latest

# Install ca-certificates for HTTPS requests to external APIs
RUN apk --no-cache add ca-certificates tzdata

# Create a non-root user
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

# Set working directory
WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/web/weather-app .

# Copy templates
COPY --from=builder /app/web/templates ./templates/

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

