package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

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
	euler()
	triangle()
	consts()
	enums()
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

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	pow := cmplx.Exp(1i*math.Pi) + 1
	fmt.Println(pow)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const filenname = "abc.txt"
	const a, b = 3, 4
	var c int = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

func enums() {
	const (
		cpp = iota
		java
		python
		golang
	)
	fmt.Println(cpp, java, python, golang)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)

}
