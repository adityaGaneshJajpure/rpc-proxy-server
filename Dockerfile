# Start from golang base image
FROM golang:1.19.6-alpine as builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
RUN apk add --update --no-cache openssh

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates dumb-init

WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8000 to the outside world
EXPOSE 3000

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["./main"]
