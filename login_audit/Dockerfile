FROM golang:1.24.6 AS builder

WORKDIR /app

# Copy go.mod and go.sum first (for caching)
COPY go.mod /

# Copy all source files
COPY . .

# Now run go mod tidy, with all code present
RUN go mod tidy

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app main.go

FROM alpine:3.20

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 3000

CMD ["./app"]
