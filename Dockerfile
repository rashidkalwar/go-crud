# Offical Golang image as a build stage
FROM golang:1.23-alpine AS builder

# Set the current working directory inside the conatiner
WORKDIR /app

# Copy the GO modules files first to cache the dependencies
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the rest of the source code into container
COPY . .

# Build the Go app
RUN go build -o go-crud ./cmd

# Use a minimal image for running the application
FROM alpine:3.18

# Set the Current Working Directory inside the contianer
WORKDIR /app

# Copy the pre-built binary file from builder stage
COPY --from=builder /app/go-crud .

# Expose the port the app runs on
EXPOSE ${PORT}

# Command to run the executable
CMD ["./go-crud"]
