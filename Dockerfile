# Start from the official Golang image (change to whichever version is required, if needed)
FROM golang:latest AS builder

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o listingd ./cmd/server

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/listingd .

# Expose the API port
EXPOSE 8080

# Run the binary
CMD ["./listingd"]
