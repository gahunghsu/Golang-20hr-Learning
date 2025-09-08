package main

import "fmt"

// swap 透過指標交換兩個整數的值
// a、b 是「位址」，在函式內用 * 取值後再互換
func swap(a *int, b *int) {
	*a, *b = *b, *a
}

// swapWrong 示範「值傳遞」：a、b 是複製品，交換不會影響外部變數
func swapWrong(a, b int) {
	a, b = b, a
}

// updateMap：map 是參考型別，直接修改內容會反映到外部
func updateMap(m map[string]int) {
	m["count"] += 1
}

func main() {
	// ------- 指標交換示範 -------
	x, y := 10, 20
	fmt.Printf("交換前  → x=%d, y=%d\n", x, y)

	swap(&x, &y) // 傳入位址
	fmt.Printf("交換後  → x=%d, y=%d (使用指標 ✅)\n", x, y)

	swapWrong(x, y) // 傳入值（複製品）
	fmt.Printf("交換錯誤→ x=%d, y=%d (值傳遞，不會變 ❌)\n\n", x, y)

	// ------- map 參考行為示範 -------
	data := map[string]int{"count": 0}
	fmt.Printf("map 初始：%v\n", data)
	updateMap(data)
	fmt.Printf("map 更新：%v (函式內修改直接生效 ✅)\n", data)
}
