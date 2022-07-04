package mock

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/MovieStoreGuy/homify/data"
	"github.com/MovieStoreGuy/homify/device"
)

type (
	// MockDevice implements all the methods that
	// a potential device can be.
	// Using this mock, you must use the highest order
	// value otherwise, you could potentially cast to the wrong type
	MockDevice struct {
		mock.Mock
	}

	MockDeviceOption func(md *MockDevice)

	// MockMonitoredDevice embeds the MockDevice
	// an implements the method GetStatistics
	// Please refere to MockDevice's warning
	MockMonitoredDevice struct {
		MockDevice
	}

	MockMonitoredDeviceOption func(md *MockMonitoredDevice)
)

var (
	_ device.Device = (*MockDevice)(nil)

	_ device.Device          = (*MockMonitoredDevice)(nil)
	_ device.MonitoredDevice = (*MockMonitoredDevice)(nil)
)

func WithAsserDeviceGetInformation(ctx context.Context, info *device.Information, err error, opts ...MethodCallOption) MockDeviceOption {
	return func(md *MockDevice) {
		call := md.On("GetInformation", ctx).Return(info, err)
		for _, opt := range opts {
			opt(call)
		}
	}
}

func WithAssertDevicePower(ctx context.Context, on bool, err error, opts ...MethodCallOption) MockDeviceOption {
	return func(md *MockDevice) {
		call := md.On("Power", ctx, on).Return(err)
		for _, opt := range opts {
			opt(call)
		}
	}
}

func WithAssertDeviceHeartbeat(ctx context.Context, state device.State, err error, opts ...MethodCallOption) MockDeviceOption {
	return func(md *MockDevice) {
		call := md.On("Heatbeat", ctx).Return(state, err)
		for _, opt := range opts {
			opt(call)
		}
	}
}

func WithMonitoredDeviceOptions(opts ...MockDeviceOption) MockMonitoredDeviceOption {
	return func(md *MockMonitoredDevice) {
		for _, opt := range opts {
			opt(&md.MockDevice)
		}
	}
}

func WithAssertMonitoredDeviceGetStatstics(ctx context.Context, stats []data.Statistic, err error, opts ...MethodCallOption) MockMonitoredDeviceOption {
	return func(md *MockMonitoredDevice) {
		call := md.On("GetStatistics", ctx).Return(stats, err)
		for _, opt := range opts {
			opt(call)
		}
	}
}

func NewDevice(tb testing.TB, opts ...MockDeviceOption) *MockDevice {
	md := &MockDevice{}
	for _, opt := range opts {
		opt(md)
	}

	tb.Cleanup(func() {
		md.AssertExpectations(tb)
	})
	return md
}

func NewMonitoredDevice(tb testing.TB, opts ...MockMonitoredDeviceOption) *MockMonitoredDevice {
	md := &MockMonitoredDevice{}
	for _, opt := range opts {
		opt(md)
	}
	tb.Cleanup(func() {
		md.AssertExpectations(tb)
	})
	return md
}

func (md *MockDevice) GetInformation(ctx context.Context) (*device.Information, error) {
	args := md.Called(ctx)
	return args.Get(0).(*device.Information), args.Error(1)
}
func (md *MockDevice) Power(ctx context.Context, on bool) error {
	return md.Called(ctx, on).Error(0)
}

func (md *MockDevice) Heartbeat(context.Context) (device.State, error) {
	args := md.Called()
	return args.Get(0).(device.State), args.Error(1)
}

func (md *MockMonitoredDevice) GetStatistics(ctx context.Context) ([]data.Statistic, error) {
	args := md.Mock.Called(ctx)
	return args.Get(0).([]data.Statistic), args.Error(1)
}
