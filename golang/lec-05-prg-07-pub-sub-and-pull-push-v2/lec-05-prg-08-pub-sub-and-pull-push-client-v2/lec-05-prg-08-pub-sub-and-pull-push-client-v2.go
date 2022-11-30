package main

import (
	"fmt"
	"math/rand"
	"os"
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

	clientID := os.Args[1]

	poller := zmq.NewPoller()
	poller.Add(subscriber, zmq.POLLIN)
	for {
		sockets, _ := poller.Poll(100)
		if len(sockets) != 0 {
			for _, socket := range sockets {
				sock := socket.Socket
				msg, _ := sock.Recv(0)
				fmt.Printf("%s: received status => %s\n", clientID, msg)
			}
		} else {
			rand_int := rand.Intn(100)
			if rand_int < 10 {
				time.Sleep(time.Second)
				msg := fmt.Sprintf("(%s: ON)", clientID)
				publisher.Send(msg, 0)
				fmt.Printf("%s: send status - activated\n ", clientID)
			} else if rand_int > 90 {
				time.Sleep(time.Second)
				msg := fmt.Sprintf("(%s: OFF)", clientID)
				publisher.Send(msg, 0)
				fmt.Printf("%s: send status - deactivated\n ", clientID)
			}
		}
	}
}
