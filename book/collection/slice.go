package main

import "fmt"

func main() {

	var highRiseBuilding [30]int

	for i := 0; i < 30; i++ {
		highRiseBuilding[i] = i + 1
	}

	fmt.Println(highRiseBuilding[10:15])

	fmt.Println(highRiseBuilding[20:])

	fmt.Println(highRiseBuilding[:2])

	//a:= []int{1,2,3}

	//fmt.Println(a[:])
	//
	//fmt.Println(a[0:0])

	var strList []string

	var numList []int

	var numListEmpty = []int{}

	fmt.Println(strList, numList, numListEmpty)

	fmt.Println(len(strList), len(numList), len(numListEmpty))

	fmt.Println(strList == nil)

	fmt.Println(numList == nil)

	fmt.Println(numListEmpty == nil)

	a := make([]int, 2)
	b := make([]int, 2, 10)

	fmt.Println(a, b)
	fmt.Println(len(a), len(b))

	var numbers []int

	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d  cap:  %d  pointer:  %p\n", len(numbers), cap(numbers), numbers)
	}

	var car []string

	car = append(car, "old driver")

	car = append(car, "ice", "Sniper", "Monk")

	team := []string{"pig", "cake"}

	car = append(car, team...)

	fmt.Println(car)

	const elementCount = 1000

	srcData := make([]int, elementCount)

	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}

	refData := srcData

	copyData := make([]int, elementCount)

	copy(copyData, srcData)

	srcData[0] = 999

	fmt.Println(refData[0])

	fmt.Println(copyData[0], copyData[elementCount-1])

	copy(copyData, srcData[4:6])

	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}

	seq := []int{0, 1, 2, 3, 4, 5, 6, 7}

	seq = append(seq[:3], seq[4:]...)

	fmt.Println(seq)

}
