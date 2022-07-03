package data

import (
	"fmt"
	"strings"
	"time"

	"go.opentelemetry.io/otel/attribute"
)

type (
	// Statistic is the base interface type
	// that defines the shared methods.
	Statistic interface {
		GetName() string
		GetAttributes() []attribute.KeyValue
		GetTimestamp() time.Time
		GetMonotonic() bool
	}

	// StatisticDouble
	StatisticDouble interface {
		Statistic

		GetValue() *Value[float64]
	}

	StatisticInt interface {
		Statistic

		GetValue() *Value[int64]
	}

	stat struct {
		Name       string
		Attributes []attribute.KeyValue
		Timestamp  time.Time
		Monotonic  bool
	}

	statValue[T ValueType] struct {
		stat
		Value *Value[T]
	}

	// StatisticBuilder allows for additional properities
	// to be applied to the statistic.
	StatisticBuilder func(s *stat)
)

func WithTimestamp(ts time.Time) StatisticBuilder {
	return func(s *stat) {
		s.Timestamp = ts
	}
}

func WithMonotonic() StatisticBuilder {
	return func(s *stat) {
		s.Monotonic = true
	}
}

func WithAttributes(attrs ...attribute.KeyValue) StatisticBuilder {
	return func(s *stat) {
		s.Attributes = append(s.Attributes[:0], attrs...)
	}
}

// NewStatistic creates a new statistic value with a fixed name
// and allows to set a custom properities for it.
func NewStatisticDouble(name string, value float64, unit StatisticUnit, opts ...StatisticBuilder) StatisticDouble {
	return newStatistic(name, NewValue(value, unit), opts...)
}

func NewStatisticInt(name string, value int64, unit StatisticUnit, opts ...StatisticBuilder) StatisticInt {
	return newStatistic(name, NewValue(value, unit), opts...)
}

func newStatistic[VT ValueType](name string, value *Value[VT], opts ...StatisticBuilder) *statValue[VT] {
	stat := &statValue[VT]{
		stat: stat{
			Name:      name,
			Monotonic: false,
			Timestamp: time.Now(),
		},
		Value: value,
	}
	for _, opt := range opts {
		opt(&stat.stat)
	}
	return stat
}

func (s *stat) GetName() string                     { return s.Name }
func (s *stat) GetAttributes() []attribute.KeyValue { return s.Attributes }
func (s *stat) GetTimestamp() time.Time             { return s.Timestamp }
func (s *stat) GetMonotonic() bool                  { return s.Monotonic }

func (s *statValue[T]) GetValue() *Value[T] {
	return s.Value
}

func (s *statValue[T]) String() string {
	return strings.Join([]string{
		fmt.Sprint("name:", s.Name),
		fmt.Sprint("attributes:", s.Attributes),
		fmt.Sprint("timestamp:", s.Timestamp.Format(time.RFC3339Nano)),
		fmt.Sprint("monotonic:", s.Monotonic),
		fmt.Sprint("value:", s.Value.String()),
	}, " ")
}
