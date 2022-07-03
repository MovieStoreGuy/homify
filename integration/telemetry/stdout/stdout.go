package stdout

import (
	"context"

	"go.uber.org/zap"

	"github.com/MovieStoreGuy/homify/device"
	"github.com/MovieStoreGuy/homify/integration/telemetry"
	"github.com/MovieStoreGuy/homify/integration/telemetry/helper"
)

type Logger struct {
	log *zap.Logger
}

var (
	_ telemetry.Provider = (*Logger)(nil)
)

func NewLogger(log *zap.Logger) *Logger {
	return &Logger{log: log}
}

func (l *Logger) CaptureStatistics(ctx context.Context, devices device.List) error {
	stats, err := helper.CaptureStatistics(ctx, devices)
	if err != nil {
		return err
	}
	for _, s := range stats {
		l.log.Info("Statistic", zap.Any("value", s))
	}
	return nil
}

func (l *Logger) Flush(_ context.Context) error {
	return l.log.Sync()
}
