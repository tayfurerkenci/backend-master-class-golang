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

# install curl
RUN apk add curl
# install migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz

# run stage
FROM alpine:3.19
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

COPY --from=builder /app/migrate ./migrate
COPY ./app.env .
COPY ./start.sh .
RUN chmod +x /app/start.sh
COPY ./wait-for.sh .
RUN chmod +x /app/wait-for.sh
COPY ./db/migration db/migration

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable when the container starts
CMD ["./main"]

ENTRYPOINT [ "/app/start.sh" ]