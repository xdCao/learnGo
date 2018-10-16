package main

//func main() {
//
//}

type Cat struct {
	Color string
	Name  string
}

func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

func NewCatByColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}

type BlackCat struct {
	Cat
}

func NewBlackCatByColor(color string) *BlackCat {
	blackCat := &BlackCat{}
	blackCat.Color = color
	return blackCat
}
