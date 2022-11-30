package main

import (
	"fmt"
	"math/rand"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	fmt.Println("Publishing updates at weather server...")

	socket, _ := zmq.NewSocket(zmq.PUB)
	defer socket.Close()

	socket.Bind("tcp://*:5556")
	socket.Bind("ipc://weather.ipc")

	rand.Seed(time.Now().UnixNano())

	for {
		zipcode := rand.Intn(100000)
		temperature := -80 + rand.Intn(216)
		relhumidity := 10 + rand.Intn(51)

		msg := fmt.Sprintf("%05d %d %d", zipcode, temperature, relhumidity)

		socket.Send(msg, 0)
	}
}
