package logger

import (
	"context"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// zapLogger struct using zap without worker pool
type zapLogger struct {
	log *zap.SugaredLogger
}

// NewZapLogger initializes a zapLogger
func NewZapLogger() Logger {
	cfg := zap.Config{
		Encoding:    "console",
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		OutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:          "time",
			EncodeTime:       SyslogTimeEncoder,
			LevelKey:         "level",
			EncodeLevel:      CustomLevelEncoder,
			MessageKey:       "message",
			ConsoleSeparator: " ",
		},
	}

	logger, _ := cfg.Build()
	sugar := logger.Sugar()

	defer sugar.Sync()

	return &zapLogger{
		log: sugar,
	}
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// CustomLevelEncoder with colors for different levels
func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var color string
	switch level {
	case zapcore.DebugLevel:
		color = getColor("DEBUG") // Gray
	case zapcore.InfoLevel:
		color = getColor("INFO") // Green
	case zapcore.WarnLevel:
		color = getColor("WARN") // Yellow
	case zapcore.ErrorLevel:
		color = getColor("ERROR") // Red
	case zapcore.FatalLevel, zapcore.PanicLevel:
		color = getColor("FATAL") // Magenta
	default:
		color = getColor("DEFAULT") // Default color (reset)
	}

	// Append color, level, and reset color code
	enc.AppendString(color + "[" + level.CapitalString() + "]" + "\033[0m")
}

// zapLogger methods implementation
func (d *zapLogger) Debug(ctx context.Context, args ...interface{}) {
	d.log.Debug(args...)
}

func (d *zapLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	d.log.Debugf(format, args...)
}

func (d *zapLogger) Info(ctx context.Context, args ...interface{}) {
	d.log.Info(args...)
}

func (d *zapLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	d.log.Infof(format, args...)
}

func (d *zapLogger) Warn(ctx context.Context, args ...interface{}) {
	d.log.Warn(args...)
}

func (d *zapLogger) Warnf(ctx context.Context, format string, args ...interface{}) {
	d.log.Warnf(format, args...)
}

func (d *zapLogger) Error(ctx context.Context, args ...interface{}) {
	d.log.Error(args...)
}

func (d *zapLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	d.log.Errorf(format, args...)
}

func (d *zapLogger) Fatal(ctx context.Context, args ...interface{}) {
	d.log.Fatal(args...)
}

func (d *zapLogger) Fatalf(ctx context.Context, format string, args ...interface{}) {
	d.log.Fatalf(format, args...)
}
