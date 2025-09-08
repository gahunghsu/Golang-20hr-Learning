package main

import "fmt"

// User 定義一個自訂型別，包含 Name 與 Age
type User struct {
	Name string
	Age  int
}

// 值接收者 (Value Receiver)
// 這裡的 u 是一份複製品，修改不會影響原本的 User
func (u User) SayHello() {
	fmt.Printf("哈囉，我是 %s，今年 %d 歲！\n", u.Name, u.Age)
}

// 指標接收者 (Pointer Receiver)
// 這裡的 u 是指標，可以修改原本的 User 資料
func (u *User) SetAge(newAge int) {
	u.Age = newAge
}

func main() {
	// 建立一個 User
	user := User{Name: "Alice", Age: 20}

	// 使用值接收者的方法
	user.SayHello()

	// 嘗試用值接收者修改 (失敗示範)
	copyUser := user
	copyUser.Age = 30
	fmt.Printf("copyUser.Age = %d，但原本 user.Age = %d\n", copyUser.Age, user.Age)

	// 使用指標接收者修改 (成功)
	user.SetAge(25)
	fmt.Printf("使用 SetAge 修改後，user.Age = %d ✅\n", user.Age)

	// 再次呼叫方法
	user.SayHello()
}
