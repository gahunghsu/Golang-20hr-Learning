# Day 13：Context（逾時與取消）⏱️

本日目標：學會使用 `context` 在多個 goroutine 間**傳遞逾時/取消**，並撰寫**可中斷的工作**，避免 goroutine 洩漏。

---

## 🧩 程式概念

這個範例延續 Day12 的 Producer–Worker–Consumer 流程，加入：

* `context.WithTimeout`：整體流程的最長時間
* **Producer / Worker / Consumer 全程傳遞 `ctx` 並尊重取消**
* **工作可中斷**：以 `select` 同步等待 `time.After` 與 `<-ctx.Done()`

> **關鍵心法**：只要 `ctx` 逾時或取消，所有持有 `ctx` 的 goroutine 都應盡快返回，避免卡住或資源洩漏。

---

## 🚀 執行方式

```bash
# 逾時較短，容易看到中途被取消（timeout=1500ms, 每項工作 400ms）
go run timeout.go -jobs=8 -workers=3 -workms=400 -timeout=1500 -bufsize=0

# 拉長逾時，觀察完整跑完（timeout=6000ms）
go run timeout.go -jobs=8 -workers=3 -workms=400 -timeout=6000 -bufsize=0

# 加上緩衝，看結果堆積行為（bufsize=3）
go run timeout.go -jobs=8 -workers=3 -workms=400 -timeout=1500 -bufsize=3
```

---

## 🔍 觀察重點

### 1. 逾時觸發後的傳遞效果

* 你會看到 `ctx done` 的輸出先在 **Worker** 出現，之後 **Producer** 停止 enqueue、**Consumer** 停止消費。
* 每個角色都要**主動**監聽 `select { case <-ctx.Done(): ... }`，才會「聽得見」取消訊號。

### 2. 可中斷的 Sleep / I/O

用以下模式避免被動等待無法被中斷：

```go
select {
case <-time.After(workDur):
  // 正常完成
case <-ctx.Done():
  // 被取消
}
```

⚠️ 千萬別直接呼叫 `time.Sleep(workDur)`，那樣中途不會理會取消。

### 3. 寫入/讀取 channel 也要可中斷

* 寫入 `results` 時一樣用 `select` 同步監聽 `<-ctx.Done()>`，避免在滿佇列時**永遠卡住**。

---

## 📖 Context 基本概念

`context` 是 Go 標準庫提供的機制，用來在 goroutine 之間傳遞**取消**與**截止時間**訊號，並附帶少量**值（metadata）**。

三個核心概念：

1. **取消（Cancellation）**
   一旦某個 goroutine 透過 `cancel()` 取消，所有繼承該 `ctx` 的子 goroutine 都會收到通知並應盡快退出。

2. **逾時（Timeout / Deadline）**
   以 `WithTimeout` 或 `WithDeadline` 設定最長時間，時間到自動取消，統一整條流程的生命週期。

3. **值傳遞（Values）**
   以 `context.WithValue` 傳遞少量 metadata（如 trace ID、user ID）。**不建議**放大型物件或業務資料。

---

## 🌍 常見用途

* **API 請求逾時控制**
  避免 HTTP 呼叫卡死太久，設定 2 秒超時，自動取消。

* **資料庫操作**
  查詢/交易時間過長時中斷，避免拖垮系統資源。

* **微服務之間的呼叫**
  在請求鏈路傳遞 trace ID 與逾時設定，讓整條鏈路有一致的退出行為。

* **長時間任務管理**
  排程/批次任務可由外部發出取消，所有 goroutine 接到 `ctx` 訊號後有序結束。

---

## ✅ 最佳實務（快速清單）

* **一律傳遞 ctx**：`func work(ctx context.Context, ...)`
* **最外層使用 `WithTimeout` / `WithDeadline`**，並 `defer cancel()`
* **所有等待都用 `select` 同步監聽**：`time.After` / I/O / channel 與 `<-ctx.Done()`
* **Producer / Worker / Consumer 全部尊重取消**
* **關閉順序**：`close(jobs)` → `wg.Wait()` → `close(results)`
* **避免把大量資料塞進 `context.Value`**（只放 metadata，不放大物件）

---

## 🧪 建議實驗

* 縮短 `-timeout`，觀察誰先印出 `ctx done`。
* 增大 `-bufsize`，觀察結果先堆在 `results`，`Consumer` 被逾時中斷後如何收尾。
* 加大 `-workms`，感受「長工時」中途被取消的退出路徑。

---

## 🧾 一句話總結

👉 把 `ctx` 當成「流程生命線」。只要可能需要提前終止，就把 `ctx` 傳下去並尊重它。
