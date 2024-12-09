package logger

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

// CustomFormatter struct to format logs with color
type CustomFormatter struct {
	logrus.TextFormatter
}

// Format formats the log entry with colors for different log levels
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Define color codes
	var color string
	levelName := map[logrus.Level]string{
		logrus.DebugLevel: "DEBUG",
		logrus.InfoLevel:  "INFO",
		logrus.WarnLevel:  "WARN",
		logrus.ErrorLevel: "ERROR",
		logrus.FatalLevel: "FATAL",
		logrus.PanicLevel: "PANIC",
	}[entry.Level]

	if levelName == "" {
		levelName = entry.Level.String() // Fallback to default
	}

	switch entry.Level {
	case logrus.DebugLevel:
		color = getColor("DEBUG") // Gray
	case logrus.InfoLevel:
		color = getColor("INFO") // Green
	case logrus.WarnLevel:
		color = getColor("WARN") // Yellow
	case logrus.ErrorLevel:
		color = getColor("ERROR") // Red
	case logrus.FatalLevel, logrus.PanicLevel:
		color = getColor("FATAL") // Magenta
	default:
		color = getColor("DEFAULT") // Default color (reset)
	}

	// Get the timestamp
	timestamp := entry.Time.Format("2006-01-02 15:04:05.000")

	// Format the log level
	level := fmt.Sprintf("%s[%s]\033[0m", color, strings.ToUpper(levelName))

	// Format the message
	message := entry.Message

	// Construct the final log line
	line := fmt.Sprintf("%s %s %s\n", timestamp, level, message)

	return []byte(line), nil
}

// logrusLogger struct
type logrusLogger struct {
	log *logrus.Logger
}

// NewLogrusLogger initializes a logrusLogger with a custom formatter
func NewLogrusLogger() Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&CustomFormatter{})
	return &logrusLogger{
		log: logger,
	}
}

// logrusLogger methods implementation
func (d *logrusLogger) Debug(ctx context.Context, args ...interface{}) {
	d.log.Debug(args...)
}

func (d *logrusLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	d.log.Debugf(format, args...)
}

func (d *logrusLogger) Info(ctx context.Context, args ...interface{}) {
	d.log.Info(args...)
}

func (d *logrusLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	d.log.Infof(format, args...)
}

func (d *logrusLogger) Warn(ctx context.Context, args ...interface{}) {
	d.log.Warn(args...)
}

func (d *logrusLogger) Warnf(ctx context.Context, format string, args ...interface{}) {
	d.log.Warnf(format, args...)
}

func (d *logrusLogger) Error(ctx context.Context, args ...interface{}) {
	d.log.Error(args...)
}

func (d *logrusLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	d.log.Errorf(format, args...)
}

func (d *logrusLogger) Fatal(ctx context.Context, args ...interface{}) {
	d.log.Fatal(args...)
}

func (d *logrusLogger) Fatalf(ctx context.Context, format string, args ...interface{}) {
	d.log.Fatalf(format, args...)
}
