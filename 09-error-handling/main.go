package main

import (
	"errors"
	"fmt"
)

// ---- 1) Sentinel error（可比對的固定錯誤）----
var ErrDivideByZero = errors.New("divide by zero")

// SafeDiv：安全除法，b 為 0 時回傳錯誤（並用 %w 包裝 sentinel）
func SafeDiv(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("SafeDiv: %w (a=%.2f, b=%.2f)", ErrDivideByZero, a, b)
	}
	return a / b, nil
}

// ---- 2) 自訂錯誤型別 + errors.As 範例 ----
type NotFoundError struct {
	Resource string
	ID       int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s with id %d not found", e.Resource, e.ID)
}

func GetUserName(id int) (string, error) {
	// demo：只有 id=42 有資料，其餘視為找不到
	if id == 42 {
		return "Alice", nil
	}
	return "", &NotFoundError{Resource: "User", ID: id}
}

// ---- 3) panic / recover 範例（僅示範，實務少用）----
func mustPositive(n int) {
	if n < 0 {
		panic(fmt.Sprintf("invalid n: %d (must be >= 0)", n))
	}
	fmt.Println("mustPositive OK, n =", n)
}

func runWithRecover(f func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()
	f()
}

func main() {
	// --- SafeDiv: ok case ---
	fmt.Println("== SafeDiv：正常案例 ==")
	if v, err := SafeDiv(10, 2); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n\n", v)
	}

	// --- SafeDiv: divide by zero ---
	fmt.Println("== SafeDiv：錯誤案例（除以 0） ==")
	if _, err := SafeDiv(10, 0); err != nil {
		// 用 errors.Is 比對包裝後的錯誤
		if errors.Is(err, ErrDivideByZero) {
			fmt.Println("caught divide-by-zero (via errors.Is):", err)
		} else {
			fmt.Println("other error:", err)
		}
	}
	fmt.Println()

	// --- 自訂錯誤型別 + errors.As ---
	fmt.Println("== GetUserName：自訂錯誤型別 + errors.As ==")
	if name, err := GetUserName(7); err != nil {
		var nf *NotFoundError
		if errors.As(err, &nf) {
			fmt.Printf("not found: resource=%s id=%d\n", nf.Resource, nf.ID)
		} else {
			fmt.Println("unexpected error:", err)
		}
	} else {
		fmt.Println("username =", name)
	}
	if name, err := GetUserName(42); err == nil {
		fmt.Println("username =", name)
	}
	fmt.Println()

	// --- panic / recover 示範 ---
	fmt.Println("== panic / recover 示範（僅示範用） ==")
	runWithRecover(func() { mustPositive(5) })
	runWithRecover(func() { mustPositive(-1) }) // 將被 recover
}
