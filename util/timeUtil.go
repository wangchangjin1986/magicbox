package util

import (
	"time"
)

// timeToMs returns an integer number, which represents t in milliseconds.
func TimeToMs(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// msToTime returns the Location time corresponding to the given Unix time,
func MsToTime(t int64, loc *time.Location) time.Time {
	return time.Unix(0, t*int64(time.Millisecond)).In(loc)
}

// msToTime returns the UTC time corresponding to the given Unix time,
// t milliseconds since January 1, 1970 UTC.
func MsToUTCTime(t int64) time.Time {
	return time.Unix(0, t*int64(time.Millisecond)).In(time.UTC)
}
