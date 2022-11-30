package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"log"
	"os"
	"strconv"
)

func pop(msg []string) (head, tail []string) {
	if msg[1] == "" {
		head = msg[:2]
		tail = msg[2:]
	} else {
		head = msg[:1]
		tail = msg[1:]
	}
	return
}

func server_task(num_worker int) {
	frontend, _ := zmq.NewSocket(zmq.ROUTER)
	defer frontend.Close()
	frontend.Bind("tcp://*:5570")

	backend, _ := zmq.NewSocket(zmq.DEALER)
	defer backend.Close()
	backend.Bind("inproc://backend")

	for i := 0; i < num_worker; i++ {
		go server_worker(i)
	}

	err := zmq.Proxy(frontend, backend, nil)
	log.Fatalln("Proxy interrupted:", err)
}

func server_worker(num int) {
	worker, _ := zmq.NewSocket(zmq.DEALER)
	defer worker.Close()
	worker.Connect("inproc://backend")
	fmt.Printf("Worker#%d started\n", num)
	for {
		msg, _ := worker.RecvMessage(0)
		identity, content := pop(msg)
		fmt.Printf("Worker#%d received %s from %s\n", num, content, identity)
		worker.SendMessage(identity, content)
	}
}

func main() {
	input_msg := os.Args[1]
	num_server, _ := strconv.Atoi(input_msg)
	server_task(num_server)
}
