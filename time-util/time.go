package time_util

import (
	"time"
)
// time.Time -> 毫秒时间戳
func timeToMillisecond(t time.Time) int64 {
	if s := t.UnixNano()/1e6; s > 0 {
		return s
	}
	return 0
}
// 毫秒时间戳 -> time.Time
func millisecondTotime(ms int64) time.Time {
	if ms > 0 {
		return time.Unix(0, ms*int64(time.Millisecond))
	}
	return time.Time{}
}
// 字符串转时间戳
func strToTime(dateTime string) (int64, error) {
	layout := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(layout, dateTime, time.Local)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}
