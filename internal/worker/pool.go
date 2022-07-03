package worker

import (
	"context"
	"sync"

	"go.uber.org/multierr"
)

type (
	// Pool defines a worker
	// that is able to compute
	// a list of jobss concurrently
	// and capture those results
	Pool[Result any] struct {
		wg   sync.WaitGroup
		jobs []Job[Result]
	}
)

// NewPool returns a worker pool
// that collections results of type T
func NewPool[T any]() *Pool[T] {
	return &Pool[T]{}
}

func (p *Pool[T]) Add(jobs ...Job[T]) {
	p.wg.Add(len(jobs))
	p.jobs = append(p.jobs, jobs...)
}

func (p *Pool[T]) Collect(ctx context.Context) ([]T, error) {
	var (
		results = make(chan T, len(p.jobs))
		issues  = make(chan error, len(p.jobs))
	)
	for _, j := range p.jobs {
		go func(ctx context.Context, j Job[T]) {
			defer p.wg.Done()

			t, err := j.Do(ctx)
			if err != nil {
				issues <- err
				return
			}
			results <- t
		}(ctx, j)
	}
	p.wg.Wait()
	close(results)
	close(issues)

	var (
		data []T
		errs error
	)
	for d := range results {
		data = append(data, d)
	}
	for err := range issues {
		errs = multierr.Append(errs, err)
	}
	p.jobs = p.jobs[:0]
	return data, errs
}

func (p *Pool[T]) Do(ctx context.Context) error {
	_, err := p.Collect(ctx)
	return err
}
