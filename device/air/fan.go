package air

import (
	"github.com/MovieStoreGuy/homify/device"
	"github.com/MovieStoreGuy/homify/internal/context"
)

type (
	// Fan represents an IoT device that works as a fan
	Fan interface {
		device.Device
		// SetSpeed will set the fan speed based on the percentage
		// being passed in.
		SetSpeed(ctx context.Context, speed float64) error
	}

	OscillatingFan interface {
		Fan
		// SetOscillation will start the fans oscillation
		SetOscillation(ctx context.Context, on bool) error
	}
)
