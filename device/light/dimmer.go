package light

import (
	"github.com/MovieStoreGuy/homify/internal/context"
)

type (
	// Dimmer represents a light that is dimmable
	Dimmer interface {
		Light
		// SetBrightness allows you to set the brightness
		// from a percentage of 0.0 to 100.0
		SetBrightness(ctx context.Context, percentage float64) error
	}
)
