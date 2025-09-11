package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

/*
觀察目標：
1) -bufsize=0（無緩衝）：Producer 的 send 會等到 Worker 正在接收，log 會「一收一送」交錯。
2) -bufsize=N（有緩衝）：Producer 可以先連續塞滿 N 個任務，log 會出現 Producer 先連發、Worker 之後慢慢消化。
*/

func main() {
	var (
		numJobs  = flag.Int("jobs", 5, "任務數量")
		workers  = flag.Int("workers", 3, "worker 數量")
		bufsize  = flag.Int("bufsize", 0, "channel 緩衝大小（0=無緩衝）")
		workMS   = flag.Int("workms", 200, "每個任務模擬耗時(ms)")
		sameBuff = flag.Bool("samebuf", true, "results 是否與 jobs 使用相同緩衝大小")
	)
	flag.Parse()

	start := time.Now()
	ts := func() string { return time.Since(start).Truncate(time.Millisecond).String() }

	jobs := make(chan int, *bufsize)
	resSize := *bufsize
	if !*sameBuff {
		resSize = 0 // 想觀察 results 無緩衝時，設成 0
	}
	results := make(chan string, resSize)

	fmt.Printf("[SETUP %s] jobs=%d workers=%d bufsize=%d work=%dms (results buf=%d)\n",
		ts(), *numJobs, *workers, *bufsize, *workMS, resSize)

	var wg sync.WaitGroup
	// 啟動 workers
	for w := 1; w <= *workers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := range jobs {
				fmt.Printf("[W%d %s] <- job %d\n", id, ts(), j)
				time.Sleep(time.Duration(*workMS) * time.Millisecond)
				out := fmt.Sprintf("[W%d %s] done job %d -> %d", id, ts(), j, j*j)
				// 觀察 results 的阻塞行為
				fmt.Printf("[W%d %s] -> results (job %d)\n", id, ts(), j)
				results <- out
			}
			fmt.Printf("[W%d %s] jobs closed, worker exit\n", id, ts())
		}(w)
	}

	// Producer：送工作
	go func() {
		for j := 1; j <= *numJobs; j++ {
			fmt.Printf("[P  %s] -> jobs (enqueue %d)\n", ts(), j)
			jobs <- j // 無緩衝時，這裡會等到有 worker 收到
			fmt.Printf("[P  %s] enqueued %d\n", ts(), j)
		}
		close(jobs)
		fmt.Printf("[P  %s] close(jobs)\n", ts())
	}()

	// 收集結果
	go func() {
		wg.Wait()
		close(results)
		fmt.Printf("[C  %s] close(results)\n", ts())
	}()

	// 消費 results
	for line := range results {
		fmt.Printf("[C  %s] %s\n", ts(), line)
	}

	fmt.Printf("[DONE %s]\n", ts())
}
