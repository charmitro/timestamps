# Matching Timestamps of a Periodic Task

# Build

- local: `go build`
- docker: `docker build -t timestamps:latest .`

# Run

- local: `go run main.go -port=<port number, default is 8080>` or `./timestamps.{out|exe} -port=<port number, default is 8080>"`, e.g. `go run main.go -port=8000`
- docker: `docker run -it -p 8080:8080 timestamps:latest`

# Tests

`go test ./...`
