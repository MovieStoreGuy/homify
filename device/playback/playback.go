package playback

import (
	"github.com/MovieStoreGuy/homify/device"
	"github.com/MovieStoreGuy/homify/internal/context"
)

type Device interface {
	device.Device

	Playback(ctx context.Context) error
}
