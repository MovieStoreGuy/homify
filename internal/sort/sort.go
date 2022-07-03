package sort

import "sort"

type sorter[T any] struct {
	slice []T
	less  func(a, b T) bool
}

var (
	_ sort.Interface = (*sorter[any])(nil)
)

func (s *sorter[T]) Len() int           { return len(s.slice) }
func (s *sorter[T]) Swap(i, j int)      { s.slice[i], s.slice[j] = s.slice[j], s.slice[i] }
func (s *sorter[T]) Less(i, j int) bool { return s.less(s.slice[i], s.slice[j]) }

func SortBy[T any](slice []T, by func(a, b T) bool) []T {
	s := &sorter[T]{
		slice: slice,
		less:  by,
	}
	sort.Sort(s)
	return s.slice
}
