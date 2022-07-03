package helper

import (
	"context"

	"github.com/MovieStoreGuy/homify/data"
	"github.com/MovieStoreGuy/homify/device"
	"github.com/MovieStoreGuy/homify/internal/sort"
	"github.com/MovieStoreGuy/homify/internal/worker"
)

// CaptureStatistics filteres the list of devices and then concurrently polls
// one of the devices for information then reporting back each sorted by timestamp.
func CaptureStatistics(ctx context.Context, list *device.List) ([]data.Statistic, error) {
	var monitors []device.MonitoredDevice
	list.Range(func(d device.Device) bool {
		if monitor, ok := d.(device.MonitoredDevice); ok {
			monitors = append(monitors, monitor)
		}
		return true
	})

	w := worker.NewPool[[]data.Statistic]()
	for _, m := range monitors {
		w.Add(worker.JobFunc[[]data.Statistic](m.GetStatistics))
	}

	stats, err := w.Collect(ctx)
	if err != nil {
		return nil, err

	}
	results := []data.Statistic{}
	for _, s := range stats {
		results = append(results, s...)
	}
	return sort.SortBy(results, func(a, b data.Statistic) bool {
		return a.GetTimestamp().Before(b.GetTimestamp())
	}), nil
}
