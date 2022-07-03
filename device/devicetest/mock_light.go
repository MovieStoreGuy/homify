package devicetest

import (
	"context"
	"image/color"
	"testing"

	"github.com/MovieStoreGuy/homify/device"
	"github.com/MovieStoreGuy/homify/device/light"
)

type (
	MockLight = MockDevice

	MockLightOption = MockDeviceOption

	MockDimmer struct {
		MockLight
	}

	MockDimmerOption func(md *MockDimmer)

	MockColour struct {
		MockLight
	}

	MockColouredOption func(mc *MockColour)
)

var (
	_ device.Device = (*MockLight)(nil)
	_ light.Light   = (*MockLight)(nil)

	_ device.Device = (*MockDimmer)(nil)
	_ light.Light   = (*MockDimmer)(nil)
	_ light.Dimmer  = (*MockDimmer)(nil)

	_ device.Device = (*MockColour)(nil)
	_ light.Light   = (*MockColour)(nil)
	_ light.Colour  = (*MockColour)(nil)
)

func WithLightDeviceOptions(opts ...MockDeviceOption) MockLightOption {
	return func(md *MockDevice) {
		for _, opt := range opts {
			opt(md)
		}
	}
}

func WithDimmerDeviceOptions(opts ...MockDeviceOption) MockDimmerOption {
	return func(md *MockDimmer) {
		for _, opt := range opts {
			opt(&md.MockLight)
		}
	}
}

func WithLightAssertSetBrightness(ctx context.Context, percentage float64, err error, opts ...MethodCallOption) MockDimmerOption {
	return func(md *MockDimmer) {
		call := md.On("SetBrightness", ctx, percentage).Return(err)
		for _, opt := range opts {
			opt(call)
		}
	}
}

func WithColourDeviceOptions(opts ...MockDeviceOption) MockColouredOption {
	return func(mc *MockColour) {
		for _, opt := range opts {
			opt(&mc.MockLight)
		}
	}
}

func WithColourAssertSetColour(ctx context.Context, c color.Color, err error, opts ...MethodCallOption) MockColouredOption {
	return func(mc *MockColour) {
		call := mc.On("SetColour", ctx, c).Return(err)
		for _, opt := range opts {
			opt(call)
		}
	}
}

func NewLight(tb testing.TB, opts ...MockLightOption) *MockLight {
	return NewDevice(tb, opts...)
}

func NewDimmer(tb testing.TB, opts ...MockDimmerOption) *MockDimmer {
	md := &MockDimmer{}
	for _, opt := range opts {
		opt(md)
	}
	tb.Cleanup(func() {
		md.AssertExpectations(tb)
	})
	return md
}

func NewColour(tb testing.TB, opts ...MockColouredOption) *MockColour {
	mc := &MockColour{}
	for _, opt := range opts {
		opt(mc)
	}
	tb.Cleanup(func() {
		mc.AssertExpectations(tb)
	})
	return mc
}

func (md *MockDimmer) SetBrightness(ctx context.Context, percentage float64) error {
	return md.Called(ctx, percentage).Error(0)
}

func (mc *MockColour) SetColour(ctx context.Context, c color.Color) error {
	return mc.Called(ctx, c).Error(0)
}
