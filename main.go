package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	originServer := "127.0.0.1:8080"
	fmt.Println("Server started")

	listener, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println("cannot open the port at 3000")
	}
	CLientconn, err := listener.Accept()

	reader := bufio.NewReader(CLientconn)

	buffer := make([]byte, 4000)

	reader.Read(buffer)
	// connecting with the origin server

	MainConn, err := net.Dial("tcp", originServer)
	MainConn.Write(buffer)
	backendToProxy := make([]byte, 4000)
	MainConn.Read(backendToProxy)
	CLientconn.Write(backendToProxy)
}
