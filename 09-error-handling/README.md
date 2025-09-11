# Day 9 — 錯誤處理（Error Handling）

## 學習重點
- `error` 是 Go 內建介面，慣用 **回傳值 + error** 型式處理錯誤  
- 產生錯誤的方式：
  - `errors.New("...")` 建立固定錯誤（sentinel）
  - `fmt.Errorf("...: %w", err)` 包裝錯誤（wrapping）
- 錯誤判斷：
  - `errors.Is(err, target)`：判斷是否為某個已知錯誤（支援被包裝）
  - `errors.As(err, &targetType)`：解構出特定錯誤型別（常用於自訂錯誤）
- 自訂錯誤型別：實作 `Error() string`，可攜帶更多語意與欄位（如資源與 ID）
- `panic / recover`：
  - 僅用於**不可回復**、程式無法繼續的狀況
  - 一般邏輯錯誤請回傳 `error`，不要濫用 `panic`

## 學習心得
以前常把例外當流程控制，今天體會到 Go 的哲學：**錯誤就是資料**。  
以回傳 `error` 的方式讓呼叫端自行決策（重試、回退、忽略或中止），比起隱性拋例外更可預測。  

實作上也更有結構：  
- 用 **sentinel**（如 `ErrDivideByZero`）做穩定比對  
- 透過 **`%w` 包裝** 保留上下文，搭配 `errors.Is/As` 拆解  
- `panic/recover` 只在「真的走不下去」時作為保護網，避免濫用  

總之，Go 的錯誤處理把「控制權」交還給呼叫端，讓模組之間的契約更清楚、行為更可預期，維護起來也更放心 👍
