package base

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func HandleSession(conn net.Conn, exitChan chan int) {
	fmt.Println("Session started")
	reader := bufio.NewReader(conn)
	for {
		readString, e := reader.ReadString('\n')
		if e == nil {
			readString = strings.TrimSpace(readString)
			if !ProcessTelnetCommand(readString, exitChan) {
				conn.Close()
				break
			}
			conn.Write([]byte(readString + "\r\n"))
		} else {
			fmt.Println("Session closed")
			conn.Close()
			break
		}
	}
}
