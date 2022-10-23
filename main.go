package main

import (
	"os"
	"strconv"

	"github.com/iverly/go-mcping/mcping"
)

func main() {
	address := "localhost"
	if len(os.Args) >= 2 {
		address = os.Args[1]
	}

	port := 25565
	if len(os.Args) >= 3 {
		rawPort := os.Args[2]
		extractedPort, err := strconv.Atoi(rawPort)
		if err == nil && extractedPort <= 64738 {
			port = extractedPort
		}
	}

	pinger := mcping.NewPinger()
	_, err := pinger.Ping(address, uint16(port))
	if err != nil {
		println("Server is not running")
		os.Exit(1)
	} else {
		println("Server is running")
		os.Exit(0)
	}
}
