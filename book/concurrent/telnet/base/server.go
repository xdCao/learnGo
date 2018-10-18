package base

import (
	"fmt"
	"net"
)

func Server(address string, exitChan chan int) {
	listener, e := net.Listen("tcp", address)
	if e != nil {
		fmt.Println(e.Error())
		exitChan <- 1
	}
	fmt.Println("listen , address : ", address)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		go HandleSession(conn, exitChan)
	}
}
