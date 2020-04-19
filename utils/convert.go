package utils

import "time"

// TimeToMs ...
func TimeToMs(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// MsToTime ...
func MsToTime(ms int64) time.Time {
	return time.Unix(0, ms*int64(time.Millisecond))
}
