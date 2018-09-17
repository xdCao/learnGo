package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	fmt.Println(eval(16, 4, "/"))
	fmt.Println(div(13, 3))
	fmt.Println(apply(func(a int, b int) int {
		return a - b
	}, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5, 6))

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)

}

func eval(a, b int, op string) (int, error) {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		q, _ := div(a, b)
		result = q
	default:
		return 0, fmt.Errorf("unsupported operatopr: %s", op)
	}
	return result, nil
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	pointer := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(pointer).Name()
	fmt.Printf("Calling function %s with args "+"(%d , %d)\n", opName, a, b)
	return op(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b *int) {
	*b, *a = *a, *b
}
