package scheduler

import (
	"context"
	"time"
)

type Policy struct {
	MaxAttempts int
	Backoff     time.Duration
}

func DefaultPolicy() Policy { return Policy{MaxAttempts: 3, Backoff: time.Minute} }
func (p Policy) Delay(attempt int) time.Duration {
	if attempt < 1 {
		return 0
	}
	return time.Duration(attempt) * p.Backoff
}
func (p Policy) Execute(ctx context.Context, fn func(context.Context) error) error {
	var err error
	for n := 0; n < p.MaxAttempts; n++ {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		err = fn(ctx)
		if err == nil {
			return nil
		}
		if p.Backoff > 0 {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(p.Backoff):
			}
		}
	}
	return err
}
