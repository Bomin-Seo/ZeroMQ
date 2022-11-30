package main

import (
	"fmt"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.REP)
	socket.Bind("tcp://*:5555")

	for {
		message, _ := socket.Recv(0)
		fmt.Printf("Received request: %s\n", message)

		time.Sleep(time.Second * 1)

		socket.Send("World", 0)
	}
}
