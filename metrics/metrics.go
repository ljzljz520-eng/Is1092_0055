package metrics

import (
	"magazine-editor/model"
	"sync"
	"time"
)

type Counter struct {
	mu     sync.Mutex
	values map[string]int64
}

func New() *Counter                         { return &Counter{values: map[string]int64{}} }
func (c *Counter) Inc(name string)          { c.mu.Lock(); c.values[name]++; c.mu.Unlock() }
func (c *Counter) Add(name string, n int64) { c.mu.Lock(); c.values[name] += n; c.mu.Unlock() }
func (c *Counter) Get(name string) int64    { c.mu.Lock(); defer c.mu.Unlock(); return c.values[name] }
func (c *Counter) Snapshot() map[string]int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	o := map[string]int64{}
	for k, v := range c.values {
		o[k] = v
	}
	return o
}
func Age(r model.Record, now time.Time) time.Duration {
	if now.Before(r.CreatedAt) {
		return 0
	}
	return now.Sub(r.CreatedAt)
}
