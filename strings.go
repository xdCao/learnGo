package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "hahhahhaha我是谁"
	fmt.Println(len(s))

	fmt.Printf("%x", []byte(s))
	fmt.Println()

	for i, ch := range s {
		fmt.Printf("( %d %x)", i, ch)
	}
	fmt.Println()

	n := utf8.RuneCountInString(s)
	fmt.Println(n)

	bytes := []byte(s)

	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d,%c)", i, ch)
	}
	fmt.Println()

}
