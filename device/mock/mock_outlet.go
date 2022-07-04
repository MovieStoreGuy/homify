package mock

import (
	"testing"

	"github.com/MovieStoreGuy/homify/device/outlet"
)

type (
	MockOutlet struct {
		MockDevice
		outlet.EmbedableOutlet
	}

	MockOutletOption func(mo *MockOutlet)

	MockMonitor struct {
		MockMonitoredDevice
		outlet.EmbedableMonitor
	}

	MockMonitorOption func(mo *MockMonitor)
)

func WithOutletDeviceOptions(opts ...MockDeviceOption) MockOutletOption {
	return func(mo *MockOutlet) {
		for _, opt := range opts {
			opt(&mo.MockDevice)
		}
	}
}

func WithMonitorDeviceOptions(opts ...MockMonitoredDeviceOption) MockMonitorOption {
	return func(mo *MockMonitor) {
		for _, opt := range opts {
			opt(&mo.MockMonitoredDevice)
		}
	}
}

func NewOutlet(tb testing.TB, opts ...MockOutletOption) outlet.Outlet {
	mo := &MockOutlet{}
	for _, opt := range opts {
		opt(mo)
	}
	tb.Cleanup(func() {
		mo.AssertExpectations(tb)
	})
	return mo
}

func NewMonitor(tb testing.TB, opts ...MockMonitorOption) outlet.Monitor {
	mm := &MockMonitor{}
	for _, opt := range opts {
		opt(mm)
	}
	tb.Cleanup(func() {
		mm.AssertExpectations(tb)
	})
	return mm
}
