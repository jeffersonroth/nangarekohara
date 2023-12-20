FROM golang:1.21-bookworm

LABEL maintainer="Jefferson Johannes Roth Filho"

# Install git
RUN apt-get update && apt-get install -y git

# Install Air
RUN go install github.com/cosmtrek/air@latest

# Set the Current Working Directory inside the container
RUN mkdir /app
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY ../. .
COPY ../.env .

# Build the Go app
RUN go build -o main ./cmd/nangarekohara

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["air"]
