# Start from a base Go image
FROM golang:1.19-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download and cache Go dependencies
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN go mod tidy
WORKDIR /app/cmd
RUN go build -o main

# Create a new stage for the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the previous stage
COPY --from=build /app/cmd/main .

# Expose the port that the application listens on
EXPOSE 3000

# Run the Go application
CMD ["./main"]