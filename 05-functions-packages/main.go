package main

import (
	"fmt"
	"golang-20hr-learning/05-functions-packages/mathutil"
)

func main() {
	a, b := 10, 5

	fmt.Printf("%d + %d = %d\n", a, b, mathutil.Add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, mathutil.Subtract(a, b))

	// 測試多回傳值
	if result, ok := mathutil.Divide(a, b); ok {
		fmt.Printf("%d ÷ %d = %d\n", a, b, result)
	} else {
		fmt.Println("除數不可為 0")
	}
}
