package outlet

import (
	"github.com/MovieStoreGuy/homify/device"
)

type (
	// Monitor represents an outlet that can track its energy usage
	Monitor interface {
		device.MonitoredDevice

		privateMonitor()
	}

	// EmbedableMonitor must be embeded into the concrete struct type
	// in order for the type to adhere to the Monitor interface
	EmbedableMonitor struct{}
)

func (EmbedableMonitor) privateMonitor() {}
