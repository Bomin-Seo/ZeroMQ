package main

import (
	"fmt"
	"math/rand"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()
	subscriber.SetSubscribe("")
	subscriber.Connect("tcp://localhost:5557")

	publisher, _ := zmq.NewSocket(zmq.PUSH)
	defer publisher.Close()
	publisher.Connect("tcp://localhost:5558")

	rand.Seed(time.Now().UnixNano())

	poller := zmq.NewPoller()
	poller.Add(subscriber, zmq.POLLIN)
	for {
		sockets, _ := poller.Poll(100)
		if len(sockets) != 0 {
			for _, socket := range sockets {
				sock := socket.Socket
				msg, _ := sock.Recv(0)
				fmt.Println("I: received message ", msg)
			}
		} else {
			rand_int := rand.Intn(100)
			if rand_int < 10 {
				msg := fmt.Sprintf("%d", rand_int)
				publisher.Send(msg, 0)
				fmt.Println("I: sending message ", rand_int)
			}
		}

	}
}
