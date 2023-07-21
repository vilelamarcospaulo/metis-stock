# Use a multi-stage build to keep the final image small
# Stage 1: Build the Go binary
FROM golang:1.17-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy the entire project into the container
COPY ./src .

RUN ls

# Build the Go binary
RUN go build -o server main/main.go


# Stage 2: Create the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go binary from the build stage
COPY --from=build /app/server .

# Expose the port if your Go application listens on a specific port
EXPOSE 8080

# Start the server
CMD ["./server"]
