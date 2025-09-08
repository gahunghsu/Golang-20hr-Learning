package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "hello world"

	// 使用 map 統計字母出現次數
	counts := make(map[rune]int)

	for _, ch := range text {
		if ch != ' ' { // 忽略空白
			counts[ch]++
		}
	}

	// 輸出結果
	fmt.Println("字母出現次數：")
	for k, v := range counts {
		fmt.Printf("%c : %d\n", k, v)
	}

	// 額外範例: Slice 操作
	words := strings.Split(text, " ")
	fmt.Println("\n使用 Slice 拆分字串:")
	fmt.Println(words)
	fmt.Printf("Split: %#v\n", words)

	// 陣列範例
	arr := [3]int{1, 2, 3}
	fmt.Println("\n陣列內容:", arr)

	// Slice 範例
	slice := []int{1, 2}
	slice = append(slice, 3, 4)
	fmt.Println("Slice 內容:", slice, "長度:", len(slice), "容量:", cap(slice))
}
