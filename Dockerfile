# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Brandon Wofford <gwoff11@gmail.com>"

# Arguments for database config
ARG DATABASE_HOST
ARG DATABASE_PORT
ARG DATABASE_USER
ARG DATABASE_NAME
ARG DATABASE_PASSWORD

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Set the environment variables for the database from the arguments
ENV DATABASE_HOST=$DATABASE_HOST
ENV DATABASE_PORT=$DATABASE_PORT
ENV DATABASE_USER=$DATABASE_USER
ENV DATABASE_NAME=$DATABASE_NAME
ENV DATABASE_PASSWORD=$DATABASE_PASSWORD

# Expose port 8080 to the outside world (or whichever port your app runs on)
EXPOSE 32500

# Command to run the executable
CMD ["./main"]
