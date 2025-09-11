# Day 10 — 測試（Testing）

## 學習重點
- 使用 `go test` 撰寫與執行單元測試
- **Table-driven tests**：以表格列出多組輸入/輸出案例
- **子測試** `t.Run(name, func)`：讓每組案例有清楚名稱
- **錯誤比對**：`errors.Is/As` 準確檢查預期錯誤
- **Example 測試**：文件即測試，檢查輸出是否一致
- **Benchmark 測試**：`go test -bench=.` 量測效能
- **Fuzz 測試**（Go 1.18+）：隨機輸入找邊角錯


## 指令速查
```bash
# 執行所有測試
go test ./10-testing -v

# 顯示覆蓋率
go test ./10-testing -cover

# 產生覆蓋率報告（HTML）
go test ./10-testing -coverprofile=cover.out
go tool cover -html=cover.out

# 基準測試（benchmark）
go test ./10-testing -bench=. -benchmem

# Fuzz 測試（Go 1.18+）
go test ./10-testing -fuzz=Fuzz -fuzztime=5s

