package time_util

import (
	"time"
)

func timeToMillisecond(t time.Time) int64 {
	if s := t.UnixNano()/1e6; s > 0 {
		return s
	}
	return 0
}

func millisecondTotime(ms int64) time.Time {
	if ms > 0 {
		return time.Unix(0, ms*int64(time.Millisecond))
	}
	return time.Time{}
}

