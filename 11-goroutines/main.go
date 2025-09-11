package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Day 11 目標：
1. 體驗 goroutine 的並行特性
2. 學會 WaitGroup + channel 來收集結果
3. 示範 worker pool（多工人處理多工作）
*/

func main() {
	// 使用 rand.New 建立本地亂數產生器（Go 1.20+ 推薦）
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	nums := []int{1, 2, 3, 4, 5}
	results := make(chan string, len(nums))
	var wg sync.WaitGroup

	// 啟動 goroutines 計算平方
	for _, n := range nums {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(r.Intn(250)) * time.Millisecond)
			results <- fmt.Sprintf("%d^2 = %d", n, n*n)
		}()
	}

	// 等待完成後關閉 channel
	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("== 並行計算平方（輸出順序不保證）==")
	for line := range results {
		fmt.Println(line)
	}

	// --- Worker Pool 範例 ---
	fmt.Println("\n== Worker Pool 範例（3 workers）==")
	jobs := make(chan int, len(nums))
	out := make(chan string, len(nums))

	workers := 3
	var ww sync.WaitGroup
	for i := 1; i <= workers; i++ {
		ww.Add(1)
		go func(id int) {
			defer ww.Done()
			for j := range jobs {
				time.Sleep(50 * time.Millisecond)
				out <- fmt.Sprintf("[W%d] %d^3 = %d", id, j, j*j*j)
			}
		}(i)
	}

	for _, n := range nums {
		jobs <- n
	}
	close(jobs)

	go func() {
		ww.Wait()
		close(out)
	}()

	for line := range out {
		fmt.Println(line)
	}
}
