package server

import "time"

// TimeToStr takes a time and converts it to a date string in the format
// Jan 02, 2006
func TimeToStr(t time.Time) string {
	return t.Format("Jan, 02 2006")
}

// StrToTime takes a date string in the following format and converts it to Time
// Jan 02, 2006
func StrToTime(s string) (time.Time, error) {
	return time.Parse("Jan, 02 2006", s)
}
