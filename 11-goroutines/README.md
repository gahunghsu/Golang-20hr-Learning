# Day 11 — Goroutine（並行）

## 學習重點
- 使用 `go func()` 啟動 goroutine，讓工作並行執行
- 用 `sync.WaitGroup` 等待所有 goroutine 完成
- 用 **channel** 收集結果（fan-in）
- 了解 Go 1.22 之後 for 迴圈變數捕捉已經修正，不再需要 `n := n`
- 認識 **Worker Pool**：多個 worker 從同一個 channel 處理工作

## 注意：for 迴圈變數捕捉
- **Go 1.21 以前**：迴圈變數是同一個位址，goroutine 可能讀到錯的值 → 解法是 `n := n`  
- **Go 1.22 以後**：語言本身修正了，每次迭代都會生成新變數 → 不需要再寫 `n := n`  

## 學習心得
- goroutine 很輕量，開多個也不會心虛  
- `WaitGroup + channel` 幾乎是最常用的同步組合  
- for 迴圈變數捕捉在 **Go 1.22 以前**要小心，但新版已修正  
- Worker Pool 讓我看到 Go 並行程式設計的彈性
