package devices

import "time"

func fixTimezone(parsedTime time.Time) time.Time {
	return time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day(),
		parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second(),
		parsedTime.Nanosecond(), time.Local)
}

func boolToFloat64(b bool) float64 {
	if b {
		return 1
	}
	return 0
}
