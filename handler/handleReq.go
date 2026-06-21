package handler

import (
	"fmt"
	"io"
	"net"
)

func HandleReq(ClientConn net.Conn) {
	MainConn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("could not connect to main server")
		MainConn.Close()
		return
	}
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	go func() {
		io.Copy(MainConn, ClientConn)
		ch1 <- true

	}()
	//but the browser expects  a response till it closes the connection
	go func() {
		go io.Copy(ClientConn, MainConn)
		ch2 <- true
	}()
	<-ch1
	<-ch2
	ClientConn.Close()
	MainConn.Close()
}
