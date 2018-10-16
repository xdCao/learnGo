package main

import "fmt"

func main() {
	p := new(Property)
	p.SetValue(100)
	fmt.Println(p.Value())

	p1 := Point{1, 1}
	p2 := Point{2, 2}

	result := p1.Add(p2)

	fmt.Println(result)
}

type Bag struct {
	items []int
}

func (b *Bag) Insert(itemId int) {
	b.items = append(b.items, itemId)
}

type Property struct {
	value int
}

func (p *Property) SetValue(v int) {
	p.value = v
}

func (p *Property) Value() int {
	return p.value
}

type Point struct {
	X int
	Y int
}

func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}
