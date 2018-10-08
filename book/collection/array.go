package main

import (
	"fmt"
)

func main() {

	//	声明数组
	var team [3]string
	team[0] = "hammer"
	team[1] = "soldier"
	team[2] = "mum"
	fmt.Println(team)

	//	初始化数组
	var team1 = [...]string{"hammer", "soldier", "mum"}

	//	遍历数组
	for k, v := range team1 {
		fmt.Println(k, v)
	}

}
