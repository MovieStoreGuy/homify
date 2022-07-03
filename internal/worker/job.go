package worker

import "context"

type (
	// Job is a single run task
	// the returns either an an error or the result
	Job[T any] interface {
		Do(ctx context.Context) (T, error)
	}

	// JobFunc allows you to define a function as a job
	JobFunc[T any] func(ctx context.Context) (T, error)
)

func (jb JobFunc[T]) Do(ctx context.Context) (T, error) {
	return jb(ctx)
}
