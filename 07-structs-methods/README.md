# Day 7 — Struct 與方法

## 學習重點
- 學會使用 `struct` 定義自訂型別  
- 瞭解 **方法 (method)** 與 **函式 (function)** 的差別  
  - 函式：`func doSomething(x int) int`  
  - 方法：`func (u User) SayHello() string`  
- 方法的 **receiver** 有兩種：
  - 值接收者 (value receiver)：操作的是複製品
  - 指標接收者 (pointer receiver)：可以修改原本的資料
- 練習範例：
  - 定義 `User` struct，欄位包含 `Name`、`Age`
  - 為 `User` 寫一個 `SayHello()` 方法
  - 嘗試用值接收者 vs 指標接收者的差別

## 學習心得
今天終於進入到物件導向的味道 🎉  
Go 雖然沒有 class，但透過 **struct + method** 就能達到類似效果。  

有趣的是：  
- 如果用 **值接收者**，在方法裡修改 struct 不會影響外部 → 適合只讀操作  
- 如果用 **指標接收者**，在方法裡修改 struct 就會反映到外部 → 適合需要更新資料  

這讓我覺得 Go 的設計又回到「嚴謹 vs 鬆散」的差異：  
不像 JS 的物件隨手就能改，Go 把「要不要能改」的選擇權交給開發者，用 receiver 來控制，非常清楚 👍  

整體感覺：  
- **Go 沒有 class，但 struct + method 足夠靈活**  
- 更偏向資料結構的組合 (composition) 而不是繼承 (inheritance)  
- 這讓我開始體會 Go 的哲學：簡單、直接，但也很強大。
