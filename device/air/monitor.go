package air

import (
	"github.com/MovieStoreGuy/homify/device"
)

type (
	// Monitor represents a device that captures
	// information about the environment that it is in.
	AirMonitor interface {
		device.MonitoredDevice

		privateAitMonitor()
	}

	// EmbedableAirMonitor is used to embed inside the AirMonitor concrete
	// type so that it can be cast to an AirMonitor value
	EmbedableAirMonitor struct{}
)

func (EmbedableAirMonitor) privateAitMonitor() {}
