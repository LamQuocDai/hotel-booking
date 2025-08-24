package utils

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	layoutDMYSlash = "2/1/2006"   // e.g. 7/7/2025
	layoutISODate  = "2006-01-02" // e.g. 2025-07-07
)

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" || s == "null" {
		ct.Time = time.Time{}
		return nil
	}
	layouts := []string{
		layoutDMYSlash,
		layoutISODate,
		time.RFC3339,
	}
	var err error
	for _, l := range layouts {
		var t time.Time
		t, err = time.Parse(l, s)
		if err == nil {
			ct.Time = t
			return nil
		}
	}
	return fmt.Errorf("invalid date format: %s", s)
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return json.Marshal("")
	}
	// choose one format to return
	return json.Marshal(ct.Time.Format(layoutDMYSlash))
}
