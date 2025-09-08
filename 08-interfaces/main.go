package main

import (
	"fmt"
	"math"
)

// Shape 介面：定義所有形狀都要有 Area 方法
type Shape interface {
	Area() float64
}

// Circle 結構
type Circle struct {
	Radius float64
}

// Rectangle 結構
type Rectangle struct {
	Width, Height float64
}

// Circle 實作 Shape 介面
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Rectangle 實作 Shape 介面
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 傳入任意 Shape，呼叫 Area()
func printArea(s Shape) {
	fmt.Printf("面積 = %.2f\n", s.Area())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}

	fmt.Println("計算不同形狀的面積：")
	printArea(c) // Circle
	printArea(r) // Rectangle
}
