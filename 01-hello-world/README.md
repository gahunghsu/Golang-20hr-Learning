# Hour 01 - 環境設置與 Hello World

## 學習重點
- 安裝並驗證 Go (`go version`)
- 建立專案模組 (`go mod init`)
- 撰寫第一個 Go 程式 (Hello World)
- 使用 `go run` 執行程式

## 指令回顧
```bash
# 建立模組
go mod init golang-20hr-learning

# 到指定目錄
cd .\01-hello-world\

# 執行程式
go run main.go

# 編譯產生執行檔
go build

---

## 延伸閱讀：為什麼需要 `go build`？

在 Go 開發流程中，常見兩種方式：

### 🔹 `go run`
- 用途：快速測試程式  
- 行為：自動編譯到暫存 → 立即執行 → 結束後刪除  
- 特點：
  - 不會留下執行檔
  - 適合開發測試
  - 每次執行都需要重新編譯

### 🔹 `go build`
- 用途：編譯並產生獨立的執行檔  
- 行為：把程式編譯成機器碼，輸出單一二進位檔（Windows 會有 `.exe`）  
- 特點：
  - 不需安裝 Go 環境即可執行
  - 跨平台（Windows 開發，可編譯 Linux/Mac 執行檔）
  - 效能高、啟動快，適合部署到生產環境

### 📌 什麼時候該用哪個？
- **開發測試 →** `go run`  
- **部署上線 →** `go build`

### 🌍 為什麼 Go 特別強調 build？
- **跨平台部署**：一行指令就能輸出不同 OS/CPU 架構的程式  
- **不依賴外部環境**：不像 Python/Node.js 需要 runtime，Go build 出來的檔案可以直接執行  
- **高效能**：接近 C 語言效能，適合長時間運行的後端服務  
- **容器化便利**：只需一個小型二進位檔，Docker 映像檔能非常精簡  

👉 簡單來說：  
- `go run` 幫你「跑起來」方便開發  
- `go build` 幫你「交付出去」方便部署
