FROM golang:1.8.1-alpine

# Add all source code
COPY . ./src/udp/

# Indicate the binary as our entrypoint
ENTRYPOINT go run ./src/udp/tcp_server.go

# Expose your port
EXPOSE 10080
