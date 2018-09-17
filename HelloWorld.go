package main

import "fmt"

var (
	aa = 1
	bb = 2
	cc = "sss"
)

func main() {
	fmt.Println("Hello world!")
	varZeroValue()
	varInit()
	typeDeduction()
	shorter()
	println(aa, bb, cc)
}

func varZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func varInit() {

	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)

}

func typeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	println(a, b, c, s)
}

func shorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}
