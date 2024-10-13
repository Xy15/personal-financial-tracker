# Use the official Golang image as the base image
FROM golang:1.23-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Install swag for generating Swagger documentation
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.12

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Generate Swagger documentation
RUN swag init

# Build the Go app
RUN go build -o main .

# Final Stage: Run the Go application
FROM alpine:latest

# Copy the Go application from the builder stage
COPY --from=builder /app/main /app/main

# Expose the port
EXPOSE 8080

# Command to run the executable
CMD ["/app/main"]