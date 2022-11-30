package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	context, _ := zmq.NewContext()
	fmt.Printf("Connecting to hello world server…\n")
	socket, _ := context.NewSocket(zmq.REQ)
	socket.Connect("tcp://localhost:5555")

	for i := 0; i < 10; i++ {
		fmt.Printf("Sending request %d...\n", i)
		socket.Send("Hello", 0)

		msg, _ := socket.Recv(0)
		fmt.Printf("Received reply %d [ %s ]\n", i, msg)
	}
}
