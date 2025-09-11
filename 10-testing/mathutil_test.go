package main

import (
	"errors"
	"fmt"
	"testing"
)

// Table-driven test for Add
func TestAdd(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{"small", 1, 2, 3},
		{"zero", 0, 0, 0},
		{"negative", -3, 5, 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := Add(c.a, c.b); got != c.want {
				t.Fatalf("Add(%d,%d) = %d; want %d", c.a, c.b, got, c.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	if got := Subtract(10, 4); got != 6 {
		t.Fatalf("Subtract(10,4) = %d; want 6", got)
	}
}

func TestMultiply(t *testing.T) {
	if got := Multiply(3, 4); got != 12 {
		t.Fatalf("Multiply(3,4) = %d; want 12", got)
	}
}

func TestDivide(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		got, err := Divide(10, 2)
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		if got != 5 {
			t.Fatalf("Divide(10,2) = %d; want 5", got)
		}
	})
	t.Run("divide_by_zero", func(t *testing.T) {
		_, err := Divide(10, 0)
		if !errors.Is(err, ErrDivideByZero) {
			t.Fatalf("want ErrDivideByZero, got %v", err)
		}
	})
}

// 📚 範例測試：`go test` 會檢查輸出是否與註解相同
func ExampleAdd() {
	fmt.Println(Add(2, 3))
	// Output: 5
}

// 🚀 基準測試：`go test -bench=.` 執行
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Add(123, 456)
	}
}

/*
🧪 Fuzz 測試（Go 1.18+）
啟用方式：go test -fuzz=Fuzz -fuzztime=5s
*/
func FuzzDivide(f *testing.F) {
	f.Add(10, 2)
	f.Add(0, 1)
	f.Fuzz(func(t *testing.T, a, b int) {
		if b == 0 {
			t.Skip() // 我們定義 b=0 屬於預期錯誤，不 fuzz 這個情況
		}
		_, err := Divide(a, b)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
