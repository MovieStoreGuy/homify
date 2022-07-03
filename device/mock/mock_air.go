package mock

import (
	"context"
	"testing"

	"github.com/MovieStoreGuy/homify/data"
	"github.com/MovieStoreGuy/homify/device"
	"github.com/MovieStoreGuy/homify/device/air"
)

type (
	MockFan struct {
		MockDevice
	}

	MockFanOption func(mf *MockFan)

	MockAirMonitor = MockMonitoredDevice

	MockAirMonitorOption func(mam *MockAirMonitor)

	MockHeater[V data.ValueType] struct {
		MockFan
	}

	MockHeaterOption[V data.ValueType] func(mh *MockHeater[V])
)

var (
	_ device.Device = (*MockFan)(nil)
	_ air.Fan       = (*MockFan)(nil)

	_ device.Device          = (*MockAirMonitor)(nil)
	_ device.MonitoredDevice = (*MockAirMonitor)(nil)

	_ device.Device       = (*MockHeater[float64])(nil)
	_ air.Fan             = (*MockHeater[float64])(nil)
	_ air.Heater[float64] = (*MockHeater[float64])(nil)
)

func WithAssertDeviceOptions(opts ...MockDeviceOption) MockFanOption {
	return func(mf *MockFan) {
		for _, opt := range opts {
			opt(&mf.MockDevice)
		}
	}
}

func WithAssertFanSetSpeed(ctx context.Context, speed float64, err error, opts ...MethodCallOption) MockFanOption {
	return func(mf *MockFan) {
		call := mf.On("SetSpped", ctx, speed).Return(err)
		for _, opt := range opts {
			opt(call)
		}
	}
}

func WithHeaterFanOptions[V data.ValueType](opts ...MockFanOption) MockHeaterOption[V] {
	return func(mh *MockHeater[V]) {
		for _, opt := range opts {
			opt(&mh.MockFan)
		}
	}
}

func WithAssertHeaterSetTemperature[V data.ValueType](ctx context.Context, temp data.Value[V], err error, opts ...MethodCallOption) MockHeaterOption[V] {
	return func(mh *MockHeater[V]) {
		call := mh.On("SetTemperature", ctx, temp).Return(err)
		for _, opt := range opts {
			opt(call)
		}
	}
}

func NewFan(tb testing.TB, opts ...MockFanOption) *MockFan {
	mf := &MockFan{}
	for _, opt := range opts {
		opt(mf)
	}
	tb.Cleanup(func() {
		mf.AssertExpectations(tb)
	})
	return mf
}

func NewAirMonitor[V data.ValueType](tb testing.TB, opts ...MockAirMonitorOption) *MockAirMonitor {
	ma := &MockAirMonitor{}
	for _, opt := range opts {
		opt(ma)
	}
	tb.Cleanup(func() {
		ma.AssertExpectations(tb)
	})
	return ma
}

func NewHeater[V data.ValueType](tb testing.TB, opts ...MockHeaterOption[V]) *MockHeater[V] {
	mh := &MockHeater[V]{}
	for _, opt := range opts {
		opt(mh)
	}
	tb.Cleanup(func() {
		mh.AssertExpectations(tb)
	})
	return mh
}

func (mf *MockFan) SetSpeed(ctx context.Context, speed float64) error {
	return mf.Called(ctx, speed).Error(0)
}

func (mh *MockHeater[V]) SetTemperature(ctx context.Context, temp data.Value[V]) error {
	return mh.Called(ctx, temp).Error(0)
}
