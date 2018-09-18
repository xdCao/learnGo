package main

import "fmt"

func main() {

	var arr1 [5]int
	arr2 := [3]int{1, 3, 4}
	arr3 := [...]int{2, 3, 4, 5, 6}

	fmt.Println(arr1, arr2, arr3)

	var grid [4][5]bool

	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	for _, v := range arr3 {
		fmt.Println(v)
	}

	/*数组是一个值类型，调用方法传参是会拷贝数组的*/
	fmt.Println("print arr1:")
	printArray(&arr1)
	//printArray(arr2)
	fmt.Println("print arr3:")
	printArray(&arr3)

	fmt.Println(arr1, arr3)

	/*注意，go语言一般不使用数组*/

}

func printArray(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
