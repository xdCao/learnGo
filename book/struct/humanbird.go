package main

import "fmt"

func main() {
	b := new(Bird)
	fmt.Println("bird: ")
	b.Fly()
	b.Walk()

	h := new(Human)
	fmt.Println("human: ")
	h.Walk()
}

type Flying struct {
}

func (f *Flying) Fly() {
	fmt.Println("can fly")
}

type Walkable struct {
}

func (w *Walkable) Walk() {
	fmt.Println("can walk")
}

type Human struct {
	Walkable
}

type Bird struct {
	Flying
	Walkable
}
