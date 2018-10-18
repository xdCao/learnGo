package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go RPCServer(ch)

	recv, err := RPCClient(ch, "hi")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("client received: ", recv)
	}
}

func RPCClient(ch chan string, req string) (string, error) {
	ch <- req
	select {
	case ack := <-ch:
		return ack, nil
	case <-time.After(time.Second):
		return "", errors.New("timeout")
	}
}

func RPCServer(ch chan string) {
	for {
		data := <-ch
		fmt.Println("server received: ", data)
		ch <- "roger"
	}
}
