package main

import (
	"fmt"
)

type Data struct {
	complax  []int
	instance InnerData
	ptr      *InnerData
}

type InnerData struct {
	a int
}

func main() {

	in := Data{
		complax: []int{1, 2, 3},
		instance: InnerData{
			5,
		},
		ptr: &InnerData{1},
	}

	fmt.Printf("in value : %+v\n", in)
	fmt.Printf("in ptr: %p\n", &in)

	out := passByValue(in)

	fmt.Printf("out value : %+v\n", out)
	fmt.Printf("out ptr: %p\n", &out)

	var f func() = fire

	f()

}

func fire() {
	fmt.Println("fire")
}

func passByValue(inFunc Data) Data {
	fmt.Printf("inFUnc value : %+v\n", inFunc)
	fmt.Printf("inFunc ptr: %p", &inFunc)
	return inFunc
}
