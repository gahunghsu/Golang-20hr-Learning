# Day4 – 陣列、Slice、Map

## 學習重點
- Array：固定長度
- Slice：動態長度，可 append，支援 len / cap
- Map：類似 dictionary/hash table，使用 make 初始化

## 學習心得
學到 rune 是用來處理 Unicode 字元的，這樣才不會中文或 emoji 出錯。
map 必須用 make 初始化，不然會 panic。
除了map, make還可以初始化slice和channel; panic是 runtime error的意思