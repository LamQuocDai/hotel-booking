package util

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type CustomTime time.Time

const customTimeLayout = "1/2/2006"

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil || s == "" {
		return err
	}
	t, err := time.Parse(customTimeLayout, s)
	if err != nil {
		return err
	}
	*ct = CustomTime(t)
	return nil
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(ct).Format(customTimeLayout) + `"`), nil
}

func (ct CustomTime) GormDataType() string {
	return "date"
}

func (ct CustomTime) Value() (driver.Value, error) {
	return time.Time(ct), nil
}
