package logger

import (
	"context"
	"testing"
)

func TestLogger(t *testing.T) {
	ctx := context.Background()
	for i := 1; i <= 1; i++ {
		Infof(ctx, "log by default logger %d", i)
		Info(ctx, "log by default logger ", i)
		Debugf(ctx, "log by default logger %d", i)
		Debug(ctx, "log by default logger ", i)
		Warnf(ctx, "log by default logger %d", i)
		Warn(ctx, "log by default logger ", i)
		Errorf(ctx, "log by default logger %d", i)
		Error(ctx, "log by default logger ", i)
	}

	SetLogger(NewLogrusLogger())
	for i := 1; i <= 1; i++ {
		Infof(ctx, "log by logrus logger %d", i)
		Info(ctx, "log by logrus logger ", i)
		Debugf(ctx, "log by logrus logger %d", i)
		Debug(ctx, "log by logrus logger ", i)
		Warnf(ctx, "log by logrus logger %d", i)
		Warn(ctx, "log by logrus logger ", i)
		Errorf(ctx, "log by logrus logger %d", i)
		Error(ctx, "log by logrus logger ", i)
	}

	SetLogger(NewZapLogger())
	for i := 1; i <= 1; i++ {
		Infof(ctx, "log by zap logger %d", i)
		Info(ctx, "log by zap logger ", i)
		Debugf(ctx, "log by zap logger %d", i)
		Debug(ctx, "log by zap logger ", i)
		Warnf(ctx, "log by zap logger %d", i)
		Warn(ctx, "log by zap logger ", i)
		Errorf(ctx, "log by zap logger %d", i)
		Error(ctx, "log by zap logger ", i)
	}
}
