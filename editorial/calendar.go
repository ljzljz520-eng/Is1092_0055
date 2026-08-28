package editorial

import (
	"sort"
	"time"
)

type Slot struct {
	Day      time.Time
	Capacity int
	Used     int
}

func PlanSlots(start time.Time, days, capacity int) []Slot {
	o := make([]Slot, 0, days)
	for i := 0; i < days; i++ {
		o = append(o, Slot{start.AddDate(0, 0, i), capacity, 0})
	}
	return o
}
func Reserve(slots []Slot, n int) (int, bool) {
	for i := range slots {
		if slots[i].Used+n <= slots[i].Capacity {
			slots[i].Used += n
			return i, true
		}
	}
	return -1, false
}
func Available(slots []Slot) []Slot {
	o := []Slot{}
	for _, s := range slots {
		if s.Used < s.Capacity {
			o = append(o, s)
		}
	}
	sort.Slice(o, func(i, j int) bool { return o[i].Day.Before(o[j].Day) })
	return o
}
func Shift(s Slot, delta int) Slot { s.Day = s.Day.AddDate(0, 0, delta); return s }
