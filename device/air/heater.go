package air

import (
	"github.com/MovieStoreGuy/homify/data"
	"github.com/MovieStoreGuy/homify/internal/context"
)

type (
	Heater[V data.ValueType] interface {
		Fan
		// SetTemperature accepts a temperature tht the heater will then
		// update to closest setting for the device.
		SetTemperature(ctx context.Context, temp data.Value[V]) error
	}
)
