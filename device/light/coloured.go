package light

import (
	"image/color"

	"github.com/MovieStoreGuy/homify/internal/context"
)

type (
	// Colour represents a light device that
	// can configure the light's colour.
	Colour interface {
		Light
		// Set Colour will update the device's colour
		SetColour(ctx context.Context, c color.Color) error
	}
)
