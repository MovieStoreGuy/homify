package outlet

import "github.com/MovieStoreGuy/homify/device"

type (
	// Outlet represents a device that be controlled
	Outlet interface {
		device.Device

		privateOutlet()
	}

	// EmbedableOutlet must be used in concrete types for
	// for the concrete type to implement the Outlet interface
	EmbedableOutlet struct{}
)

func (EmbedableOutlet) privateOutlet() {}
