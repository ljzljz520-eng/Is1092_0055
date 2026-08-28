package editorial

import (
	"sort"
	"time"
)

type Task struct {
	ID, Kind string
	Priority int
	Created  time.Time
}
type Tasks []Task

func (t Tasks) Len() int { return len(t) }
func (t Tasks) Less(i, j int) bool {
	if t[i].Priority == t[j].Priority {
		return t[i].Created.Before(t[j].Created)
	}
	return t[i].Priority > t[j].Priority
}
func (t Tasks) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func Order(t Tasks) Tasks     { o := append(Tasks{}, t...); sort.Sort(o); return o }
func NextTask(t Tasks) (Task, bool) {
	o := Order(t)
	if len(o) == 0 {
		return Task{}, false
	}
	return o[0], true
}
