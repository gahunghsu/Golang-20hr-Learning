package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

/*
Day14 重點：
1) strings 常用操作：Split、TrimSpace、ToLower、Contains、ReplaceAll
2) time 常用操作：Now、Parse/Format、Duration、Ticker、sleep、每天 HH:MM 固定提醒

兩種模式：
- ticker 模式（預設）：每隔 -every 提醒一次，共 -count 次
- daily 模式：指定 -times="08:30,14:00" 這些每日時間點提醒（會依當前時間自動選下一個）

範例：
go run reminder.go -tasks=" Buy milk,Send Email , code review " -filter=mail
go run reminder.go -every=3s -count=5 -tasks="喝水, 伸展, 休息"
go run reminder.go -mode=daily -times="09:00,13:30,21:10" -tasks="review PR,standup"
*/

func main() {
	var (
		mode   = flag.String("mode", "ticker", "提醒模式：ticker | daily")
		every  = flag.String("every", "3s", "ticker 模式的間隔，例如 3s、1m、2h")
		count  = flag.Int("count", 5, "ticker 模式要提醒的次數")
		times  = flag.String("times", "", "daily 模式的時間列表，格式 HH:MM，多個以逗號分隔，例如 08:30,14:00")
		tasks  = flag.String("tasks", "喝水,伸展,休息", "要提醒的任務清單（逗號分隔）")
		filter = flag.String("filter", "", "任務關鍵字過濾（不分大小寫，包含即算符合）")
		zone   = flag.String("zone", "Local", "時區（例如 Asia/Taipei；預設 Local）")
	)
	flag.Parse()

	loc := mustLoadLocation(*zone)
	now := time.Now().In(loc)
	fmt.Printf("[SETUP %s] mode=%s zone=%s\n", ts(now), *mode, loc.String())

	// strings：清理與過濾任務
	allTasks := parseTasks(*tasks)
	if *filter != "" {
		allTasks = filterTasks(allTasks, *filter)
	}
	if len(allTasks) == 0 {
		fmt.Println("[INFO] 沒有任務可提醒（可能被 filter 全部過濾掉了）")
		return
	}
	fmt.Printf("[TASKS] %v\n", allTasks)

	switch strings.ToLower(*mode) {
	case "ticker":
		runTickerMode(loc, *every, *count, allTasks)
	case "daily":
		if strings.TrimSpace(*times) == "" {
			fmt.Println("[ERROR] daily 模式請提供 -times，例如 -times=\"09:00,13:30\"")
			return
		}
		runDailyMode(loc, *times, allTasks)
	default:
		fmt.Println("[ERROR] mode 僅支援 ticker | daily")
	}
}

// ---------- strings 區 ----------

func parseTasks(csv string) []string {
	parts := strings.Split(csv, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		p = strings.ReplaceAll(p, "  ", " ") // 將連續空白稍微清一下（簡單示範）
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func filterTasks(tasks []string, keyword string) []string {
	kw := strings.ToLower(strings.TrimSpace(keyword))
	if kw == "" {
		return tasks
	}
	out := make([]string, 0, len(tasks))
	for _, t := range tasks {
		if strings.Contains(strings.ToLower(t), kw) {
			out = append(out, t)
		}
	}
	return out
}

// ---------- ticker 模式 ----------

func runTickerMode(loc *time.Location, every string, count int, tasks []string) {
	d, err := time.ParseDuration(every)
	if err != nil || d <= 0 {
		fmt.Printf("[ERROR] -every 解析失敗或非正數：%q\n", every)
		return
	}

	ticker := time.NewTicker(d)
	defer ticker.Stop()

	fmt.Printf("[MODE] ticker every=%s count=%d\n", d, count)

	i := 0
	for range ticker.C {
		i++
		now := time.Now().In(loc)
		task := tasks[(i-1)%len(tasks)]
		fmt.Printf("[REMIND %s] (%d/%d) 該：%s\n", ts(now), i, count, task)
		if i >= count {
			fmt.Println("[DONE] ticker 模式完成")
			return
		}
	}
}

// ---------- daily 模式 ----------

func runDailyMode(loc *time.Location, timesCSV string, tasks []string) {
	targets, err := parseDailyTimes(loc, timesCSV)
	if err != nil || len(targets) == 0 {
		fmt.Printf("[ERROR] -times 解析失敗：%v\n", err)
		return
	}
	fmt.Printf("[MODE] daily times=%v (%s)\n", prettyTimes(targets), loc)

	i := 0
	for {
		// 找下一個即將到來的時間點
		now := time.Now().In(loc)
		next := nextOccurrence(now, targets)
		wait := next.Sub(now)
		fmt.Printf("[WAIT] 現在 %s，下一次在 %s（還有 %s）\n", ts(now), ts(next), wait.Truncate(time.Second))

		time.Sleep(wait) // 簡化示範；需要可中斷可搭配 context

		i++
		now = time.Now().In(loc)
		task := tasks[(i-1)%len(tasks)]
		fmt.Printf("[REMIND %s] 每日提醒：%s\n", ts(now), task)
		// 下一圈會重新計算 nextOccurrence
	}
}

func parseDailyTimes(loc *time.Location, csv string) ([]time.Time, error) {
	parts := strings.Split(csv, ",")
	out := make([]time.Time, 0, len(parts))
	now := time.Now().In(loc)
	y, m, d := now.Date()

	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		// 解析 HH:MM
		hhmm := strings.Split(p, ":")
		if len(hhmm) != 2 {
			return nil, fmt.Errorf("時間格式錯誤：%q（需 HH:MM）", p)
		}
		hh := strings.TrimSpace(hhmm[0])
		mm := strings.TrimSpace(hhmm[1])

		hour, err1 := parseInt(hh)
		min, err2 := parseInt(mm)
		if err1 != nil || err2 != nil || hour < 0 || hour > 23 || min < 0 || min > 59 {
			return nil, fmt.Errorf("時間格式錯誤：%q（需 HH:MM 且範圍正確）", p)
		}
		out = append(out, time.Date(y, m, d, hour, min, 0, 0, loc))
	}
	return out, nil
}

func parseInt(s string) (int, error) {
	var n int
	_, err := fmt.Sscanf(s, "%d", &n)
	return n, err
}

func prettyTimes(ts []time.Time) []string {
	out := make([]string, 0, len(ts))
	for _, t := range ts {
		out = append(out, t.Format("15:04"))
	}
	return out
}

// 回傳「接下來」最近的一個時間點（如已過，往明天同一時間）
func nextOccurrence(now time.Time, candidates []time.Time) time.Time {
	next := candidates[0]
	soonest := time.Duration(1<<63 - 1) // max duration
	for _, c := range candidates {
		// 將 candidate 的日期設為今天（parseDailyTimes 已是今天）
		cToday := time.Date(now.Year(), now.Month(), now.Day(), c.Hour(), c.Minute(), 0, 0, now.Location())
		if !cToday.After(now) {
			cToday = cToday.Add(24 * time.Hour)
		}
		delta := cToday.Sub(now)
		if delta < soonest {
			soonest = delta
			next = cToday
		}
	}
	return next
}

// ---------- utils ----------

func mustLoadLocation(name string) *time.Location {
	if name == "Local" || strings.TrimSpace(name) == "" {
		return time.Local
	}
	loc, err := time.LoadLocation(name)
	if err != nil {
		fmt.Printf("[WARN] 無法載入時區 %q，改用 Local：%v\n", name, err)
		return time.Local
	}
	return loc
}

func ts(t time.Time) string {
	// 範例：2025-09-12 00:15:04.123
	return t.Format("2006-01-02 15:04:05.000")
}
