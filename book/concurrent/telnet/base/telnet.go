package base

import (
	"fmt"
	"strings"
)

func ProcessTelnetCommand(str string, exitChan chan int) bool {
	if strings.HasPrefix(str, "@close") {
		fmt.Println("Session closed")
		return false
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("Server shutdown")
		exitChan <- 0
		return false
	}
	fmt.Println(str)
	return true
}
