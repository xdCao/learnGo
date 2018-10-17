package main

import "awesomeProject/book/pkg/base"
import _ "awesomeProject/book/pkg/cls1"
import _ "awesomeProject/book/pkg/cls2"

func main() {
	c1 := base.Create("Class1")
	c1.Do()
	c2 := base.Create("Class2")
	c2.Do()
}
