package main

import "os"
import "awesomeProject/book/concurrent/telnet/base"

func main() {
	exitChan := make(chan int)
	go base.Server("127.0.0.1:7001", exitChan)
	code := <-exitChan
	os.Exit(code)
}
