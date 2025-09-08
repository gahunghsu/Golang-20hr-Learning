package main

import "fmt"

func main() {
	// FizzBuzz: 印出 1~100
	for i := 1; i <= 100; i++ {
		// 判斷是否為 3 和 5 的倍數
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// 額外範例: switch 寫法
	fmt.Println("\n--- 使用 switch 範例 ---")
	for i := 1; i <= 15; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
