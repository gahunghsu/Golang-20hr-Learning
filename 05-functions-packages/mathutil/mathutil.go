package mathutil

// Add 回傳兩數相加結果
func Add(a, b int) int {
	return a + b
}

// Subtract 回傳兩數相減結果
func Subtract(a, b int) int {
	return a - b
}

// Divide 提供多回傳值 (商數, 是否成功)
func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}
