package main

import (
	"fmt"
)

func main() {

	for y := 1; y <= 9; y++ {

		for x := 1; x <= y; x++ {
			fmt.Printf("%d * %d = %d ", x, y, x*y)
		}

		fmt.Println()

	}

	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d, value:%d\n", key, value)
	}

	var str = "你好"

	for k, v := range str {
		fmt.Printf("key: %d  value: 0x%x", k, v)
	}

	m := map[string]int{
		"hello": 100,
		"world": 200,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	c := make(chan int)

	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				goto breakhere
			}
		}
	}

	return

breakhere:
	fmt.Println("done")

	err := firstCheckError()

	if err != nil {
		goto onExit
	}

	err = secondCheckError()

	if err != nil {
		goto onExit
	}

onExit:
	fmt.Println(err)
	exitProcess()

}
func exitProcess() {

}
func secondCheckError() interface{} {

}
func firstCheckError() interface{} {

}
