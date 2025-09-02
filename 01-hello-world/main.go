// main.go
// 說明：這是一個最小可執行的 Go 程式。編譯後會產生一個可執行檔，執行時在終端機印出一行文字。

// package 宣告：
// 1) 在 Go 中，所有檔案都必須屬於某個 package。
// 2) "main" 是特別的套件名稱，代表此套件可以被編譯成「可執行檔」(executable)。
// 3) 只有 package main 且包含 func main() 的程式，才有「程式進入點」。
package main

// import 匯入套件：
// - 這裡匯入的是標準函式庫的 "fmt"（format 的縮寫）。
// - fmt 提供常用的輸出入與格式化函式，例如 Print、Println、Printf。
// - 若有多個套件，可用括號分行：
//     import (
//         "fmt"
//         "os"
//     )
import "fmt"

// main 函式：
// - Go 程式的「進入點」(entry point)，執行檔從這裡開始跑。
// - 函式簽章固定為：func main()，不接受參數、不回傳值。
// - 同一個 package 只能有一個 main()，否則編譯錯誤。
func main() {
    // 呼叫 fmt.Println：將內容輸出到「標準輸出」（stdout），並在結尾自動加上換行符號 '\n'。
    // 字串常值使用雙引號包住，Go 預設採用 UTF-8，因此可直接包含中文與 Emoji。
    // Println 的回傳值為 (n int, err error)：
    //   - n 代表實際寫入的位元組數（包含換行）。
    //   - err 代表是否發生錯誤（多半為 I/O 錯誤）。這裡我們不需要，故忽略回傳值。
    fmt.Println("Hello, Golang! 🚀")
}
