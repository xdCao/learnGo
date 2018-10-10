package main

import "fmt"

func main() {
	print(1, 2, 3)
}

func rawPrint(rawList ...interface{}) {
	for _, s := range rawList {
		fmt.Println(s)
	}
}

func print(slist ...interface{}) {
	rawPrint(slist...)
}
