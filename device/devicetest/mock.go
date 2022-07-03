package devicetest

import (
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/tilinna/clock"
)

// MethodCallOption allows to extend the
// assertion of the method
type MethodCallOption func(call *mock.Call)

func WithMethodTimes(count int) MethodCallOption {
	return func(call *mock.Call) {
		if count < 1 {
			call.Maybe()
			count = 0
		}
		call.Times(count)
	}
}

func WithMethodVirtualWait(clk clock.Clock, d time.Duration) MethodCallOption {
	return func(call *mock.Call) {
		call.WaitUntil(clk.After(d))
	}
}

func WithMethodPanic(message string) MethodCallOption {
	return func(call *mock.Call) {
		call.Panic(message)
	}
}
