package mock

import (
	"context"
	"image/color"
	"testing"

	"github.com/MovieStoreGuy/homify/device/light"
)

type (
	MockLight struct {
		MockDevice
		light.EmbedableLight
	}

	MockLightOption func(ml *MockLight)

	MockDimmer struct {
		MockLight
	}

	MockDimmerOption func(md *MockDimmer)

	MockColour struct {
		MockLight
	}

	MockColouredOption func(mc *MockColour)
)

func WithLightDeviceOptions(opts ...MockLightOption) MockLightOption {
	return func(md *MockLight) {
		for _, opt := range opts {
			opt(md)
		}
	}
}

func WithDimmerDeviceOptions(opts ...MockDeviceOption) MockDimmerOption {
	return func(md *MockDimmer) {
		for _, opt := range opts {
			opt(&md.MockDevice)
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
			opt(&mc.MockDevice)
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

func NewLight(tb testing.TB, opts ...MockLightOption) light.Light {
	ml := &MockLight{}
	for _, opt := range opts {
		opt(ml)
	}
	tb.Cleanup(func() {
		ml.AssertExpectations(tb)
	})
	return ml
}

func NewDimmer(tb testing.TB, opts ...MockDimmerOption) light.Dimmer {
	md := &MockDimmer{}
	for _, opt := range opts {
		opt(md)
	}
	tb.Cleanup(func() {
		md.AssertExpectations(tb)
	})
	return md
}

func NewColour(tb testing.TB, opts ...MockColouredOption) light.Colour {
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
