# Day 12：Channel 練習 🎯

今天的重點是 **Channel**，特別觀察 **無緩衝** (`-bufsize=0`) 與 **有緩衝** (`-bufsize=N`) 的差異。

---

## 📌 程式說明

這個範例模擬一個 **Producer-Worker-Consumer** 的流程：

1. **Producer (P)**  
   將任務送進 `jobs` channel。
2. **Workers (W1, W2, W3 …)**  
   從 `jobs` 拿到任務後，執行計算（模擬耗時 `workms` 毫秒）。
3. **Consumer (C)**  
   收集所有結果並輸出。

程式可透過參數調整：

- `-jobs`：總任務數量  
- `-workers`：Worker 數量  
- `-bufsize`：`jobs` channel 緩衝大小（0=無緩衝）  
- `-workms`：每個任務模擬耗時 (ms)  
- `-samebuf`：是否讓 `results` 的緩衝與 `jobs` 相同  

---

## 🚀 執行方式

```bash
# 無緩衝 channel
go run worker.go -jobs=5 -workers=3 -bufsize=0 -workms=200

# 有緩衝 channel（例如 bufsize=3）
go run worker.go -jobs=5 -workers=3 -bufsize=3 -workms=200
````

---

## 🔍 觀察重點

### 1. 無緩衝 (`-bufsize=0`)

* Producer `jobs <- j` **會阻塞**，直到有 Worker 正在接收。
* Log 會呈現 **「一收一送」交錯**：

```text
[P  0s] -> jobs (enqueue 1)
[W1 0s] <- job 1
[P  0s] enqueued 1
```

### 2. 有緩衝 (`-bufsize=N`)

* Producer 可以先連續丟 N 個任務進 `jobs`，不會馬上被阻塞。
* Log 會呈現 **「Producer 先連發，Worker 慢慢消化」**：

```text
[P  0s] -> jobs (enqueue 1)
[P  0s] enqueued 1
[P  0s] -> jobs (enqueue 2)
[P  0s] enqueued 2
[P  0s] -> jobs (enqueue 3)
[P  0s] enqueued 3
[W1 0s] <- job 1
[W2 0s] <- job 2
[W3 0s] <- job 3
```

---

## 🤔 學習心得

* **無緩衝**：強制同步，Producer 與 Worker 一一配對，常用於「確認對方已收到」的情境。
* **有緩衝**：允許短暫的「工作堆積」，Producer 可以先跑一段，Worker 再慢慢處理。
* 對應到現實世界：

  * 無緩衝 → 工人必須在場才能接單。
  * 有緩衝 → 可以先把貨放到倉庫，工人之後再來處理。

```
