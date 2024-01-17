# Use the official Golang image for the build stage
FROM golang:1.21 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working directory inside the container
COPY . .

# Build the Go app
ARG APP_ENV
ARG PORT
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server -ldflags "-X main.Environment=$APP_ENV" .

# Start a new stage from Alpine Linux
FROM alpine:latest  

# Install ca-certificates in case you need to call HTTPS endpoints
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the pre-built binary file and necessary environment files from the builder stage
COPY --from=builder /app/server .
COPY .env.$APP_ENV /root/
COPY mysql.env /root/

# Expose the port the app runs on
EXPOSE $PORT

# Command to run the executable
CMD ["./server"]
