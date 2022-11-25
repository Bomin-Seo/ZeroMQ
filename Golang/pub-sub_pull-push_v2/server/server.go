package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {

	publisher, _ := zmq.NewSocket(zmq.PUB)
	defer publisher.Close()
	publisher.Bind("tcp://*:5557")

	collector, _ := zmq.NewSocket(zmq.PULL)
	defer collector.Close()
	collector.Bind("tcp://*:5558")

	for {
		message, _ := collector.Recv(0)
		fmt.Println("Server: publishing update => ", message)
		publisher.Send(message, 0)
	}
}
