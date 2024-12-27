# Use the official Golang image as the base image
FROM golang:1.20-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files (go.mod and go.sum)
COPY go.mod go.sum ./

# Download and cache Go dependencies
RUN go mod tidy

# Copy the entire project into the container
COPY . .

# Build the Go application (replace "main.go" with your entry point)
RUN go build -o main .

# Start a new stage to build the final image
FROM alpine:latest

# Install the necessary dependencies (e.g., ca-certificates)
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your application will listen on
EXPOSE 8080

# Command to run the Go application
CMD ["./main"]