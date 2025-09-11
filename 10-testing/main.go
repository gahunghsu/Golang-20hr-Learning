package main

import (
	"errors"
	"fmt"
)

// 簡單的數學工具 - 為了示範單元測試
func Add(a, b int) int      { return a + b }
func Subtract(a, b int) int { return a - b }
func Multiply(a, b int) int { return a * b }

var ErrDivideByZero = errors.New("divide by zero")

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func main() {
	fmt.Println("Demo:", Add(2, 3), Subtract(5, 1), Multiply(2, 4))
	if q, err := Divide(10, 2); err == nil {
		fmt.Println("Divide(10,2) =", q)
	}
}
