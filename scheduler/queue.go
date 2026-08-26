package scheduler

import (
	"context"
	"errors"
	"sync"
	"time"
)

type Job struct {
	ID       string
	Run      func(context.Context) error
	Attempts int
	Due      time.Time
}
type Queue struct {
	mu   sync.Mutex
	jobs []Job
}

func New() *Queue { return &Queue{jobs: []Job{}} }
func (q *Queue) Enqueue(j Job) error {
	if j.ID == "" || j.Run == nil {
		return errors.New("invalid job")
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	q.jobs = append(q.jobs, j)
	return nil
}
func (q *Queue) Len() int { q.mu.Lock(); defer q.mu.Unlock(); return len(q.jobs) }
func (q *Queue) Next(now time.Time) (Job, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	for n, j := range q.jobs {
		if !j.Due.After(now) {
			q.jobs = append(q.jobs[:n], q.jobs[n+1:]...)
			return j, true
		}
	}
	return Job{}, false
}
func (q *Queue) RunOne(ctx context.Context, now time.Time) error {
	j, ok := q.Next(now)
	if !ok {
		return nil
	}
	err := j.Run(ctx)
	if err != nil && j.Attempts < 3 {
		j.Attempts++
		j.Due = now.Add(time.Minute)
		_ = q.Enqueue(j)
	}
	return err
}
