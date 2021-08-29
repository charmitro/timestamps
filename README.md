# Matching Timestamps of a Periodic Task

# Build

- local: `go build`
- docker: `docker build -t timestamps .`

# Run

- local: `go run main.go -port=<port number>` or `./timestamps.{out|exe} -port=<port number>"`
- docker: `docker run -it timestamps`

# Tests

`go test ./...`
