package model

import (
	"encoding/json"
	"time"
)

func (r Record) MarshalJSON() ([]byte, error) { type alias Record; return json.Marshal(alias(r)) }
func (r *Record) UnmarshalJSON(b []byte) error {
	type alias Record
	var a alias
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	*r = Record(a)
	return nil
}
func (r Record) Age(now time.Time) time.Duration {
	if now.Before(r.CreatedAt) {
		return 0
	}
	return now.Sub(r.CreatedAt)
}
