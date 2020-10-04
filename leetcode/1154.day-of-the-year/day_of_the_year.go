package leetcode

import "time"

func dayOfYear(date string) int {
	todayZero, _ := time.ParseInLocation("2006-01-02", date, time.Local)
	return todayZero.YearDay()
}
