package main

import (
	"fmt"
	"os"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func set_id(soc *zmq.Socket, client_name string) {
	identity := client_name
	soc.SetIdentity(identity)
}

func client_task(client_name string) {

	client, _ := zmq.NewSocket(zmq.DEALER)
	defer client.Close()

	set_id(client, client_name)
	id, _ := client.GetIdentity()
	client.Connect("tcp://localhost:5570")
	fmt.Printf("Client %s started\n", id)

	poller := zmq.NewPoller()
	poller.Add(client, zmq.POLLIN)

	go func() {
		for {
			sockets, _ := poller.Poll(1000)
			if len(sockets) != 0 {
				for _, socket := range sockets {
					sock := socket.Socket
					msg, _ := sock.Recv(0)
					fmt.Printf("%s received : %s\n", id, msg)
				}
			}
		}

	}()

	for reqs := 1; true; reqs++ {
		fmt.Printf("Req #%d sent..\n ", reqs)
		msg := fmt.Sprintf("request #%d", reqs)
		client.Send(msg, 0)
		time.Sleep(1000 * time.Millisecond)
	}

}

func main() {
	client_name := os.Args[1]
	client_task(client_name)
	time.Sleep(20 * time.Second)
}
