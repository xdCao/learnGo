package main

import (
	"flag"
	"fmt"
)

//func main() {
//
//	visit([]int{1,2,3,4,5}, func(ele int) {
//		fmt.Println(ele)
//	})
//
//}
//
//func visit(list []int, f func(ele int))  {
//	for _,v := range list{
//		f(v)
//	}
//}

var skillParam = flag.String("skill", "", "skill to perform")

func main() {

	flag.Parse()

	skill := map[string]func(){
		"fire": func() {
			fmt.Println("fire")
		},
		"run": func() {
			fmt.Println("run")
		},
		"fly": func() {
			fmt.Println("fly")
		},
	}

	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}

}
