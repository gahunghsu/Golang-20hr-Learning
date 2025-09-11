# Day 15：HTTP (Server & Client) 🌐

本日目標：學會使用 Go 的 `net/http` 建立 **Web 伺服器** 與 **簡單的 HTTP Client**。

---

## 🧩 程式概念

1. **HTTP Server**
   - 使用 `http.HandleFunc` 註冊路由
   - 使用 `http.ListenAndServe(":8080", nil)` 啟動伺服器
   - 回傳固定字串 `"Hello from Go!"`

2. **HTTP Client**
   - 使用 `http.Get` 發送 GET 請求
   - 使用 `io.ReadAll` 讀取回應 Body（取代舊的 `ioutil.ReadAll`）

---

## 🛠 程式碼

```go
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from Go!")
}

func main() {
    http.HandleFunc("/", helloHandler)
    go http.ListenAndServe(":8080", nil)

    resp, _ := http.Get("http://localhost:8080")
    defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)
    fmt.Println("📡 Client received:", string(body))
}
````

---

## ▶️ 執行結果

```bash
🚀 Server running at http://localhost:8080
📡 Client received: Hello from Go!
```

---

## 🤔 學習心得

今天第一次用 Go 寫 Web Server，發現它的 **標準庫就能啟動一個伺服器**，完全不用額外框架！
比起我熟悉的 Node.js，Go 的程式碼更精簡，而且 `net/http` 已經足夠處理很多基礎需求。

👉 一句話總結：
**Go 開 Web Server 超快上手，內建 Client/Server 一把抓！**

```