package main

import (
	"fmt"
)

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:6]
	fmt.Println("arr[2:6] = ", s)
	updateSlice(s)
	fmt.Println("arr[2:6] = ", s)
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:]", arr[2:])
	fmt.Println("arr[:]", arr[:])

	/*reslice*/
	s2 := arr[:]
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	/*slice 是可以扩展的*/
	s1 := arr[2:6]
	fmt.Println(s1)
	s2 = s1[3:5]
	fmt.Println(s2)
	s3 := append(s2, 10)
	/*s4,s5指向的不再是arr了，而是一个新的arr*/
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println(s3, s4, s5)
	fmt.Println("arr = ", arr)

	/*创建slice*/
	var slice []int
	for i := 0; i < 100; i++ {
		slice = append(slice, 2*i+1)
	}
	fmt.Println(slice)
	printSlice(slice)

	slice1 := []int{2, 4, 6, 8}
	printSlice(slice1)

	slice2 := make([]int, 16)
	slice3 := make([]int, 10, 32)
	printSlice(slice2)
	printSlice(slice3)

	/*slice拷贝*/
	copy(slice2, slice1)
	printSlice(slice2)

	/*slice中删除元素*/
	slice2 = append(slice2[:3], slice2[4:]...)
	printSlice(slice2)

}

func printSlice(s []int) {
	fmt.Printf("%v", s)
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))
}

func updateSlice(s []int) {
	s[0] = 100
}
