package context

import (
	"context"
	"time"
)

type (
	Context    = context.Context
	CancelFunc = context.CancelFunc

	optimised struct {
		Context

		done <-chan struct{}
	}
)

func (o *optimised) Done() <-chan struct{} {
	return o.done
}

var (
	Background       = context.Background
	TODO             = context.TODO
	Canceled         = context.Canceled
	DeadlineExceeded = context.DeadlineExceeded

	WithValue = context.WithValue
)

func WithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	ctx, done := context.WithCancel(ctx)
	return &optimised{ctx, ctx.Done()}, done
}

func WithDeadline(ctx context.Context, t time.Time) (context.Context, context.CancelFunc) {
	ctx, done := context.WithDeadline(ctx, t)
	return &optimised{ctx, ctx.Done()}, done
}

func WithTimeout(ctx context.Context, t time.Duration) (context.Context, context.CancelFunc) {
	ctx, done := context.WithTimeout(ctx, t)
	return &optimised{ctx, ctx.Done()}, done
}
