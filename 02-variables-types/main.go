package main

import (
    "fmt"
    "math"
)

func main() {
    // 標準宣告方式
	var name string
    var age int
    var radius float64 = 5.5

	// 宣告變數 (方式二：使用 := 自動推斷型別)
    greeting := "歡迎來學 Go！"

	// 輸出提示訊息，並讀取使用者輸入
    fmt.Print("請輸入名字: ")
    fmt.Scanln(&name) // 注意：要用 & 取變數的記憶體位址

    fmt.Print("請輸入年齡: ")
    fmt.Scanln(&age)

	// 使用 Printf 格式化輸出
    fmt.Printf("哈囉 %s，你今年 %d 歲，%s\n", name, age, greeting)

	fmt.Print("你知道圓面積是pi * 半徑平方嗎?\n")
	fmt.Printf("請輸入圓的半徑 (例如 %.2f): ", radius)
	fmt.Scanln(&radius)

    // 簡短宣告（Go 會自動推斷型別）
    pi := 3.14159

    // 計算圓面積
    area := pi * math.Pow(radius, 2)

    // 使用 Printf 格式化輸出
    fmt.Printf("半徑: %.2f\n", radius)
    fmt.Printf("圓面積: %.2f\n", area)

    // 其他型別示範
    var isBigCircle bool = area > 50

    fmt.Printf("面積是否大於 50? %t\n", isBigCircle)
}
