package main

import "fmt"

func main() {

	var delegate func(int)

	c := new(Class)

	delegate = c.Do

	delegate(100)

	delegate = funcDo

	delegate(100)

}

type Class struct {
}

func (c *Class) Do(v int) {
	fmt.Println("Call method Do: ", v)
}

func funcDo(v int) {
	fmt.Println("Call funtion Do: ", v)
}
