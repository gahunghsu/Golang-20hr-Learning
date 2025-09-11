package main

import (
	"fmt"
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go!")
}

func main() {
	// 啟動 HTTP Server
	http.HandleFunc("/", helloHandler)
	go func() {
		fmt.Println("🚀 Server running at http://localhost:8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
	}()

	// 簡單 HTTP Client 請求自己
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("📡 Client received:", string(body))
}
