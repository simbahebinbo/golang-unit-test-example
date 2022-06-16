package testable

import "time"

//  judgeRate 报警速率决策函数
func judgeRate() int {
	now := time.Now()
	return judgeRateByTime(now)
}

//  judgeRateByTime 报警速率决策函数,时间作为参数
func judgeRateByTime(now time.Time) int {
	switch hour := now.Hour(); {
	case hour >= 8 && hour < 20:
		return 10
	case hour >= 20 && hour <= 23:
		return 1
	}
	return -1
}
