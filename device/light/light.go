package light

import (
	"github.com/MovieStoreGuy/homify/device"
)

type (
	// Light represents a IoT device that is lightbulb.
	Light interface {
		device.Device

		privateLight()
	}

	// EmbedableLight must be used to be embeded into
	// the concrete type so the Light interface is satisified
	EmbedableLight struct{}
)

func (EmbedableLight) privateLight() {}
