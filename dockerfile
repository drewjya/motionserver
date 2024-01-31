# Use the official Go image as the base image
FROM golang:1.19-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download and install Go dependencies
RUN go mod download

# Copy the entire application directory into the container
COPY . .

# Build the Go Fiber application
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go

# Use a smaller base image for the final image
FROM alpine:3.14

# Copy the built application from the builder stage
COPY --from=builder /app /app

# Set the working directory for the application
WORKDIR /app

# Set the entry point for the container
CMD ["./app"]
