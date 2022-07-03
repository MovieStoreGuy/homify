package device

import (
	"github.com/MovieStoreGuy/homify/data"
	"github.com/MovieStoreGuy/homify/internal/context"
)

type (
	State int

	// Device represents an IoT device information
	Device interface {
		GetInformation(ctx context.Context) (*Information, error)
		// Heartbeat checks the connectivity
		// and returns what state the device is in.
		Heartbeat(ctx context.Context) (State, error)
		// Power will turn on or off the device
		Power(ctx context.Context, on bool) error
	}

	// MonitoredDevice allows for the device to report
	// statitics that would be collected by it
	MonitoredDevice interface {
		Device
		// GetStatstics returns all avaliable statistics from the device
		// to be reported back to a telemetry provider
		GetStatistics(ctx context.Context) ([]data.Statistic, error)
	}
)

const (
	Unconfigured State = iota
	PoweredOff
	PoweredOn
)
