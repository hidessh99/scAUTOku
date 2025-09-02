# Use the official Golang image as the base image
FROM golang:1.21-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files (if they exist)
COPY go.mod go.sum ./ 2>/dev/null || echo "No go.mod file found"

# Download all dependencies
RUN go mod download 2>/dev/null || echo "No go.mod file to download"

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vpn-api .

# Use a minimal base image for the final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create a non-root user
RUN adduser -D -s /bin/sh vpnuser

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/vpn-api .

# Copy the rest-go.sh script if needed
COPY --from=builder /app/rest-go.sh .

# Make the binary executable
RUN chmod +x ./vpn-api
RUN chmod +x ./rest-go.sh

# Expose the port the app runs on
EXPOSE 3000

# Switch to the non-root user
USER vpnuser

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:3000/health || exit 1

# Run the binary
CMD ["./vpn-api"]