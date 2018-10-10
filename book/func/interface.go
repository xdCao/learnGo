package main

import "fmt"

type Invoker interface {
	Call(interface{})
}

type funcCaller func(interface{})

type Struct struct {
}

func main() {

	//s := new(Struct)
	//invoker := s
	//invoker.Call("hello")
	//s.Call("123")

	var j int = 5
	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i,j:%d,%d\n", i, j)
		}
	}() //将一个无需参数返回值为匿名函数的函数赋值给a()

	a()
	j *= 2
	// i*=2这样是错的
	a()

}

func (s *Struct) Call(p interface{}) {

	fmt.Println("from struct", p)

}

func (f funcCaller) Call(p interface{}) {
	f(p)
}
