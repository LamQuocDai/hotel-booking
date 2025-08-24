package utils

import (
	"fmt"
	"time"
)

func DateToDayInt(t time.Time) int32 {
	t = t.UTC()
	return int32(t.Year()*10000 + int(t.Month())*100 + t.Day())
}

func DayIntToTime(day int32) (time.Time, error) {
	if day <= 0 {
		return time.Time{}, fmt.Errorf("invalid day %d", day)
	}
	y := int(day / 10000)
	m := int((day / 100) % 100)
	d := int(day % 100)
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC), nil
}

func ParseFlexibleDate(s string) (time.Time, error) {
	layouts := []string{"2006-01-02", "2/1/2006", time.RFC3339}
	var lastErr error
	for _, l := range layouts {
		if t, err := time.Parse(l, s); err == nil {
			return t, nil
		} else {
			lastErr = err
		}
	}
	return time.Time{}, lastErr
}
