# Multi-stage docker build
# https://docs.docker.com/develop/develop-images/multistage-build/

# Use the official Golang image from the Docker Hub
# alpine for small image size
# build stage
FROM golang:1.21.5-alpine3.19 as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the entire project from the current directory to the working dir in the container
COPY . .

# Build the Go app. This will create an executable named 'main'
RUN go build -o main main.go

# run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable when the container starts
CMD ["./app/main"]