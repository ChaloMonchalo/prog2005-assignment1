package handler

import "time"

// startTime is set once from main.go.
var startTime time.Time

func SetStartTime(t time.Time) {
	startTime = t
}

func UptimeSeconds() int64 {
	if startTime.IsZero() {
		return 0
	}
	return int64(time.Since(startTime).Seconds())
}
