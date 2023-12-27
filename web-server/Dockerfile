# Use the official Golang image to create a build artifact.
# This is stage is named as builder.
FROM golang:1.21 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# Start a new stage from scratch
FROM alpine:latest  

# Install ca-certificates in case you need to call HTTPS endpoints.
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/server .

COPY .env /root/

# Expose port 8080 to the outside world
EXPOSE 4000

# Command to run the executable
CMD ["./server"]
