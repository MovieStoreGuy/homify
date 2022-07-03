package helper

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/attribute"

	"github.com/MovieStoreGuy/homify/data"
	"github.com/MovieStoreGuy/homify/device"
	"github.com/MovieStoreGuy/homify/device/mock"
)

func TestCaptureStatistics(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tests := []struct {
		scenario string
		devices  []device.Device
		expect   []data.Statistic
		err      error
	}{
		{
			scenario: "No Device Reported Statistics",
			devices: []device.Device{
				mock.NewDevice(t),
				mock.NewDevice(t),
				mock.NewDevice(t),
			},
			expect: []data.Statistic{},
			err:    nil,
		},
		{
			scenario: "Only one device reporting data",
			devices: []device.Device{
				mock.NewDevice(t),
				mock.NewDevice(t),
				mock.NewDevice(t),
				mock.NewMonitoredDevice(t,
					mock.WithAssertMonitoredDeviceGetStatstics(
						ctx,
						[]data.Statistic{
							data.NewStatisticInt(
								"outlet.power.usage",
								303,
								data.StatisticWatts,
								data.WithAttributes(
									attribute.String("model", "p110"),
									attribute.String("name", "office-outlet"),
									attribute.String("vendor", "tp-link"),
								),
								data.WithTimestamp(time.Unix(100, 0)),
							),
						},
						nil,
						mock.WithMethodTimes(1),
					),
				),
			},
			expect: []data.Statistic{
				data.NewStatisticInt(
					"outlet.power.usage",
					303,
					data.StatisticWatts,
					data.WithAttributes(
						attribute.String("model", "p110"),
						attribute.String("name", "office-outlet"),
						attribute.String("vendor", "tp-link"),
					),
					data.WithTimestamp(time.Unix(100, 0)),
				),
			},
			err: nil,
		},
		{
			scenario: "All devices reporting statistics",
			devices: []device.Device{
				mock.NewMonitoredDevice(t,
					mock.WithAssertMonitoredDeviceGetStatstics(
						ctx,
						[]data.Statistic{
							data.NewStatisticInt(
								"outlet.power.usage",
								1600,
								data.StatisticWatts,
								data.WithAttributes(
									attribute.String("model", "p110"),
									attribute.String("name", "coffee-machine"),
									attribute.String("vendor", "tp-link"),
								),
								data.WithTimestamp(time.Unix(100, 0)),
							),
							data.NewStatisticInt(
								"outlet.power.usage",
								200,
								data.StatisticWatts,
								data.WithAttributes(
									attribute.String("model", "p110"),
									attribute.String("name", "coffee-machine"),
									attribute.String("vendor", "tp-link"),
								),
								data.WithTimestamp(time.Unix(120, 0)),
							),
							data.NewStatisticInt(
								"outlet.power.usage",
								10,
								data.StatisticWatts,
								data.WithAttributes(
									attribute.String("model", "p110"),
									attribute.String("name", "coffee-machine"),
									attribute.String("vendor", "tp-link"),
								),
								data.WithTimestamp(time.Unix(180, 0)),
							),
						},
						nil,
						mock.WithMethodTimes(1),
					),
				),
				mock.NewMonitoredDevice(t,
					mock.WithAssertMonitoredDeviceGetStatstics(
						ctx,
						[]data.Statistic{
							data.NewStatisticInt(
								"outlet.power.usage",
								303,
								data.StatisticWatts,
								data.WithAttributes(
									attribute.String("model", "p110"),
									attribute.String("name", "office-outlet"),
									attribute.String("vendor", "tp-link"),
								),
								data.WithTimestamp(time.Unix(100, 10)),
							),
							data.NewStatisticInt(
								"outlet.power.usage",
								100,
								data.StatisticWatts,
								data.WithAttributes(
									attribute.String("model", "p110"),
									attribute.String("name", "office-outlet"),
									attribute.String("vendor", "tp-link"),
								),
								data.WithTimestamp(time.Unix(200, 10)),
							),
						},
						nil,
						mock.WithMethodTimes(1),
					),
				),
			},
			expect: []data.Statistic{
				data.NewStatisticInt(
					"outlet.power.usage",
					1600,
					data.StatisticWatts,
					data.WithAttributes(
						attribute.String("model", "p110"),
						attribute.String("name", "coffee-machine"),
						attribute.String("vendor", "tp-link"),
					),
					data.WithTimestamp(time.Unix(100, 0)),
				),
				data.NewStatisticInt(
					"outlet.power.usage",
					303,
					data.StatisticWatts,
					data.WithAttributes(
						attribute.String("model", "p110"),
						attribute.String("name", "office-outlet"),
						attribute.String("vendor", "tp-link"),
					),
					data.WithTimestamp(time.Unix(100, 10)),
				),
				data.NewStatisticInt(
					"outlet.power.usage",
					200,
					data.StatisticWatts,
					data.WithAttributes(
						attribute.String("model", "p110"),
						attribute.String("name", "coffee-machine"),
						attribute.String("vendor", "tp-link"),
					),
					data.WithTimestamp(time.Unix(120, 0)),
				),
				data.NewStatisticInt(
					"outlet.power.usage",
					10,
					data.StatisticWatts,
					data.WithAttributes(
						attribute.String("model", "p110"),
						attribute.String("name", "coffee-machine"),
						attribute.String("vendor", "tp-link"),
					),
					data.WithTimestamp(time.Unix(180, 0)),
				),
				data.NewStatisticInt(
					"outlet.power.usage",
					100,
					data.StatisticWatts,
					data.WithAttributes(
						attribute.String("model", "p110"),
						attribute.String("name", "office-outlet"),
						attribute.String("vendor", "tp-link"),
					),
					data.WithTimestamp(time.Unix(200, 10)),
				),
			},
			err: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			stats, err := CaptureStatistics(ctx, device.NewList(tc.devices...))
			assert.ErrorIs(t, err, tc.err)
			assert.EqualValues(t, tc.expect, stats)
		})
	}
}
