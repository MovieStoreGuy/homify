package influx

import (
	"context"

	influxdb "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"

	"github.com/MovieStoreGuy/homify/data"
	"github.com/MovieStoreGuy/homify/device"
	"github.com/MovieStoreGuy/homify/integration/telemetry"
	"github.com/MovieStoreGuy/homify/integration/telemetry/helper"
)

type (
	ClientV2 struct {
		client      influxdb.Client
		org, bucket string
	}
)

func NewInfluxDBv2(host, token, org, bucket string) telemetry.Provider {
	return &ClientV2{
		client: influxdb.NewClient(host, token),
		org:    org,
		bucket: bucket,
	}
}

func (c2 *ClientV2) CaptureStatistics(ctx context.Context, list *device.List) error {
	stats, err := helper.CaptureStatistics(ctx, list)
	if err != nil {
		return err
	}
	metrics := c2.client.WriteAPI(c2.org, c2.bucket)
	for i := 0; i < len(stats); i++ {
		p := write.NewPointWithMeasurement(stats[i].GetName())
		for a := 0; a < len(stats[i].GetAttributes()); a++ {
			p.AddTag(string(stats[i].GetAttributes()[a].Key), stats[i].GetAttributes()[a].Value.AsString())
		}
		switch s := stats[i].(type) {
		case data.StatisticDouble:
			p.AddField("value", s.GetValue().Value())
			p.AddTag("unit", s.GetValue().Unit().String())
		case data.StatisticInt:
			p.AddField("value", s.GetValue().Value())
			p.AddTag("unit", s.GetValue().Unit().String())
		}
		p.SetTime(stats[i].GetTimestamp())
		metrics.WritePoint(p)
	}
	metrics.Flush()
	return nil
}

func (c2 *ClientV2) Flush(_ context.Context) error {
	return nil
}
