package devicetest

import (
	"testing"

	"github.com/MovieStoreGuy/homify/device"
	"github.com/MovieStoreGuy/homify/device/outlet"
)

type (
	MockOutlet = MockDevice

	MockOutletOption = MockDeviceOption

	MockOutletMonitor = MockMonitoredDevice

	MockOutletMonitorOption func(mo *MockOutletMonitor)
)

var (
	_ device.Device = (*MockOutlet)(nil)
	_ outlet.Outlet = (*MockOutlet)(nil)

	_ device.Device          = (*MockOutletMonitor)(nil)
	_ device.MonitoredDevice = (*MockOutletMonitor)(nil)
	_ outlet.Outlet          = (*MockOutletMonitor)(nil)
	_ outlet.OutletMonitor   = (*MockOutletMonitor)(nil)
)

func NewOutlet(tb testing.TB, opts ...MockOutletOption) *MockOutlet {
	return NewDevice(tb, opts...)
}

func NewOutletMonitor(tb testing.TB, opts ...MockOutletMonitorOption) *MockOutletMonitor {
	mo := &MockOutletMonitor{}
	for _, opt := range opts {
		opt(mo)
	}
	tb.Cleanup(func() {
		mo.AssertExpectations(tb)
	})
	return mo
}
