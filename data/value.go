package data

import "fmt"

type (
	// Valuetype defines what statistical records type
	ValueType interface {
		~int64 | ~float64
	}

	// Value defines a statistic value with a statistical unit
	Value[vt ValueType] struct {
		value vt
		unit  StatisticUnit
	}
)

func NewValue[VT ValueType](value VT, unit StatisticUnit) *Value[VT] {
	return &Value[VT]{value: value, unit: unit}
}

func (v *Value[VT]) Value() VT {
	return v.value
}

func (v *Value[VT]) Unit() StatisticUnit {
	return v.unit
}

func (v *Value[VT]) String() string {
	return fmt.Sprint("value:", v.value, v.unit.String())
}
