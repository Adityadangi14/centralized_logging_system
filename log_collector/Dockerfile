# ---------- Stage 1: Build ----------
FROM golang:1.24.6 AS builder

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the entire project
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app main.go

# ---------- Stage 2: Run ----------
FROM alpine:3.20

# Install any certificates (needed for DB, HTTPS, etc.)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY .env .env

# Copy binary from builder stage
COPY --from=builder /app/app .

# Copy .env if you use one (optional)



# Expose the port (change to your app's port if needed)

EXPOSE 3000



# Run the binary
CMD ["./app"]
