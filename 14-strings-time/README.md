# Day 14：strings + time 🧵⏰

本日目標：  
1) 熟悉 `strings` 常用 API（切割、清理、搜尋）。  
2) 熟悉 `time` 常用 API（時間戳、Duration、Ticker、每天固定時間提醒）。

---

## 🧩 程式概念

這個範例提供 **兩種提醒模式**：

- **ticker 模式（預設）**：每隔固定時間提醒一次、提醒固定次數。  
- **daily 模式**：每天在指定的 `HH:MM` 時間點提醒（多個時間以逗號分隔）。

同時用 `strings` 來把 `-tasks`（逗號分隔）整理成任務清單，並可用 `-filter` 做不分大小寫的關鍵字過濾。

---

## 🚀 執行方式

```bash
# 1) 預設 ticker 模式：每 3 秒提醒一次，共 5 次
go run reminder.go -every=3s -count=5 -tasks="喝水, 伸展, 休息"

# 2) ticker + 關鍵字過濾（不分大小寫，包含即算符合）
go run reminder.go -every=5s -count=3 -tasks="Buy milk,Send Email,Code Review" -filter=mail

# 3) daily 模式：每天 09:00、13:30、21:10 提醒
go run reminder.go -mode=daily -times="09:00,13:30,21:10" -tasks="review PR, standup"

# 4) 指定時區（例如 Asia/Taipei）
go run reminder.go -mode=daily -times="08:30" -zone=Asia/Taipei -tasks="晨會"
````

> Windows / WSL / macOS 都可執行；時區未指定時使用系統 Local。

---

## 🔍 strings 常用招式

* `strings.Split(csv, ",")`：把逗號分隔的任務字串拆成 slice。
* `strings.TrimSpace(s)`：移除首尾空白。
* `strings.ReplaceAll(s, "  ", " ")`：簡單清理連續空白。
* `strings.ToLower(s)` + `strings.Contains(s, kw)`：不分大小寫的關鍵字搜尋。

> 小提醒：`strings.Fields` 也很實用，可把多空白自動切字。

---

## ⏱️ time 常用招式

* `time.Now()` / `t.Format("2006-01-02 15:04:05.000")`：取得/格式化時間。
* `time.ParseDuration("3s")`：把字串轉成 `time.Duration`。
* `time.NewTicker(d)`：每隔 `d` 觸發一次。
* **每天固定時間**的技巧：

  1. 解析 `HH:MM` → 轉為「今天」的 `time.Time`；
  2. 若已過該時間，往後加 24 小時；
  3. 從多個時間中挑「下一個最近的」，`Sleep` 到點後再重新計算下一次。

---

## 🧪 觀察重點

1. **ticker vs daily 行為差異**

   * ticker：固定間隔、固定次數。
   * daily：固定時刻、無限循環（ Ctrl+C 中斷 ）。

2. **字串清理的重要性**

   * 來源常帶空白或大小寫不一致，先 `TrimSpace` + `ToLower` 可避免邊界 bug。

3. **時區影響**

   * 若部署在伺服器，請明確指定 `-zone`（例如 `Asia/Taipei`），避免因時區不同導致提醒時間偏移。

---

## 🧱 邊界情況（可自行嘗試）

* `-times` 格式錯誤（非 `HH:MM`） → 會提示錯誤並結束。
* `-filter` 把所有任務都排除了 → 會印出「沒有任務可提醒」。
* `-every` 不是合法的 `Duration`（例如 `3秒`） → 會提示錯誤並結束。

---

## 📂 專案結構（建議）

```
14-strings-time/
├── reminder.go
└── README.md
```

---

## ✅ 總結

* `strings`：處理輸入與清理字串是日常必備。
* `time`：理解 `Duration`、格式化、固定時刻排程，是寫排程/提醒器的基礎。

> 下一步可以把 Day13 的 `context` 加進來，讓 `daily` 模式可被安全中斷（例如收到信號就取消），並避免長時間 `Sleep` 無法提早退出。

```
