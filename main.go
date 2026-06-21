package main

import (
	"cachify/handler"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Server started")

	listener, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println("cannot open the port at 3000")
		listener.Close()
	}
	for {
		Clientconn, err := listener.Accept()
		if err != nil {
			fmt.Println("could not accept the request")
			Clientconn.Close()

		}
		go handler.HandleReq(Clientconn)

	}
}
