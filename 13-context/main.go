// 13-context/timeout.go
package main

import (
	"context"
	"flag"
	"fmt"
	"sync"
	"time"
)

/*
觀察目標：
1) 了解 context.WithTimeout 的取消傳遞：一但超時/取消，所有持有 ctx 的 goroutine 都應盡快退出。
2) 正確撰寫「可中斷的工作」：用 select 同步監聽 <-ctx.Done() 與 time.After(workDur)。
3) Producer / Worker / Consumer 全流程都要「傳遞 ctx」並尊重取消，避免 goroutine 洩漏。
*/

func main() {
	var (
		numJobs   = flag.Int("jobs", 8, "任務數量")
		workers   = flag.Int("workers", 3, "worker 數量")
		workMS    = flag.Int("workms", 400, "每個任務模擬耗時(ms)")
		timeoutMS = flag.Int("timeout", 1500, "整體流程逾時(ms)")
		bufsize   = flag.Int("bufsize", 0, "jobs/results channel 緩衝大小（0=無緩衝）")
	)
	flag.Parse()

	start := time.Now()
	ts := func() string { return time.Since(start).Truncate(time.Millisecond).String() }

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeoutMS)*time.Millisecond)
	defer cancel()

	jobs := make(chan int, *bufsize)
	results := make(chan string, *bufsize)

	fmt.Printf("[SETUP %s] jobs=%d workers=%d work=%dms timeout=%dms buf=%d\n",
		ts(), *numJobs, *workers, *workMS, *timeoutMS, *bufsize)

	// 啟動 workers
	var wg sync.WaitGroup
	for w := 1; w <= *workers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("[W%d %s] ctx done: %v，worker exit\n", id, ts(), ctx.Err())
					return
				case j, ok := <-jobs:
					if !ok {
						fmt.Printf("[W%d %s] jobs closed，worker exit\n", id, ts())
						return
					}
					fmt.Printf("[W%d %s] <- job %d\n", id, ts(), j)

					// 可中斷的工作：同時等待工作耗時與取消訊號
					select {
					case <-time.After(time.Duration(*workMS) * time.Millisecond):
						out := fmt.Sprintf("[W%d %s] done job %d -> %d", id, ts(), j, j*j)
						// 寫出結果也要可中斷
						select {
						case <-ctx.Done():
							fmt.Printf("[W%d %s] ctx done before results send (job %d)\n", id, ts(), j)
							return
						case results <- out:
							fmt.Printf("[W%d %s] -> results (job %d)\n", id, ts(), j)
						}
					case <-ctx.Done():
						fmt.Printf("[W%d %s] ctx done while working (job %d)\n", id, ts(), j)
						return
					}
				}
			}
		}(w)
	}

	// Producer：送工作（尊重 ctx）
	go func() {
		defer close(jobs)
		for j := 1; j <= *numJobs; j++ {
			select {
			case <-ctx.Done():
				fmt.Printf("[P  %s] ctx done: 停止 enqueue，剩餘任務未送出\n", ts())
				return
			case jobs <- j:
				fmt.Printf("[P  %s] -> jobs (enqueue %d)\n", ts(), j)
			}
		}
		fmt.Printf("[P  %s] close(jobs)\n", ts())
	}()

	// 收集結果關閉協調
	go func() {
		wg.Wait()
		close(results)
		fmt.Printf("[C  %s] close(results)\n", ts())
	}()

	// Consumer：讀取結果（尊重 ctx）
	total := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[C  %s] ctx done: %v，停止消費\n", ts(), ctx.Err())
			// 等 results 消耗完再離開（避免遺失已產生的結果）
			for line := range results {
				fmt.Printf("[C  %s] %s\n", ts(), line)
				total++
			}
			fmt.Printf("[DONE %s] collected=%d (timeout)\n", ts(), total)
			return
		case line, ok := <-results:
			if !ok {
				fmt.Printf("[DONE %s] collected=%d\n", ts(), total)
				return
			}
			fmt.Printf("[C  %s] %s\n", ts(), line)
			total++
		}
	}
}
