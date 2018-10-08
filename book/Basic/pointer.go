package Basic

import "fmt"

func main() {

	//var cat int = 1
	//
	//var str string = "banana"
	//
	//fmt.Printf("%p %p ", &cat, &str)
	//
	//
	//house := "Malibu Point 10880, 90265"
	//ptr := &house
	//fmt.Printf("ptr type: %T\n",ptr)
	//value := *ptr
	//fmt.Printf("value type: %T\n",value)
	//fmt.Printf("value : %s\n",value)
	//
	//
	//a,b:=1,2
	//swap(&a,&b)
	//fmt.Println(a,b)

	//var mode = flag.String("mode","","process mode")
	//
	//
	//flag.Parse()
	//fmt.Println(*mode)

	//var a int
	//void()
	//fmt.Println(a,dummy(0))

	pi := 3.1415926
	variant := fmt.Sprintf("%v %v %v ", "月球基地", pi, true)
	fmt.Println(variant)

	profile := &struct {
		Name string
		HP   int
	}{
		Name: "cat",
		HP:   150,
	}

	fmt.Printf("使用%%v: %v\n", profile)
	fmt.Printf("使用%%+v: %+v", profile)
	fmt.Printf("使用%%#v: %#v", profile)
	fmt.Printf("使用%%： %T", profile)

}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func dummy(b int) int {
	var c int
	c = b
	return c
}

func void() {

}
