package main

import (
	"errors"
	"flag"
	"log"
	"strconv"

	"github.com/charmitro/timestamps/server"
)

func main() {
	port := flag.String("port", "8080", "port to listen to")
	flag.Parse()

	if _, err := strconv.Atoi(*port); err != nil {
		log.Fatal(errors.New("-port flag value must contain only numbers. e.g. 8080"))
	}

	server.Server(*port)
}
