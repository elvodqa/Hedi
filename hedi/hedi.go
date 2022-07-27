package hedi

import (
	"fmt"
	"net"
)

type Hedi struct {
	Server net.Listener
}

func CreateHedi() *Hedi {
	hedi := new(Hedi)

	// initialize services

	listener, err := net.Listen("tcp", "127.0.0.1:9467")

	if err != nil {
		fmt.Printf("Failed to create TCP server on 127.0.0.1:9467")
	}

	hedi.Server = listener

	return hedi
}

func (hedi *Hedi) RunHedi() {
	fmt.Printf("Running Hedi on 127.0.0.1:9467\n")

	for {
		conn, err := hedi.Server.Accept()
		fmt.Printf("Connection Accepted!\n")

		if err != nil {
			continue
		}

		// handle new client
	}
}
