package main

import (
	"fmt"
)

func main() {
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}

	for name, obj := range animals {
		flyer, isFlyer := obj.(Flyer)
		walker, isWalker := obj.(Walker)

		fmt.Println(name, " isFlyer: ", isFlyer)
		fmt.Println(name, " isWalker: ", isWalker)

		if isWalker {
			walker.Walk()
		}

		if isFlyer {
			flyer.Fly()
		}

	}

}

type Flyer interface {
	Fly()
}

type Walker interface {
	Walk()
}

type bird struct {
}

func (b *bird) Fly() {
	fmt.Println("bird fly")
}

func (b *bird) Walk() {
	fmt.Println("bird walk")
}

type pig struct {
}

func (p *pig) Walk() {
	fmt.Println("pig walk")
}
