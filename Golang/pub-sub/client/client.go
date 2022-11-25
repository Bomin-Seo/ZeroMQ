package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	zmq "github.com/pebbe/zmq4"
)

func main() {

	total_temp := 0
	update_nbr := 0

	socket, _ := zmq.NewSocket(zmq.SUB)
	defer socket.Close()
	socket.Connect("tcp://localhost:5556")

	zip_filter := "10001"
	if len(os.Args) > 1 {
		zip_filter = os.Args[1] + " "
	}

	fmt.Println("Collecting updates from weather server...")
	socket.SetSubscribe(zip_filter)

	for update_nbr < 20 {
		msg, _ := socket.Recv(0)
		if msgs := strings.Fields(msg); len(msgs) > 1 {
			if temp, err := strconv.Atoi(msgs[1]); err == nil {
				total_temp += temp
				update_nbr++
				fmt.Printf("Receive temperature for zipcode %s was %dF \n\n", zip_filter, temp)
			}
		}
	}
	fmt.Printf("Average temperature for zipcode %s was %dF \n\n", zip_filter, total_temp/update_nbr)
}
