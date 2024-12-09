package logger

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

// CustomLogger struct to format logs with color
type CustomLogger struct {
	log *log.Logger
}

// NewDefaultLogger initializes a CustomLogger
func NewDefaultLogger() Logger {
	return &CustomLogger{
		log: log.New(os.Stdout, "", 0),
	}
}

// formatLog formats the log message with color and timestamp
func formatLog(level string, message string) string {

	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	return fmt.Sprintf("%s %s[%s]\033[0m %s\n", timestamp, getColor(level), level, message)
}

// logWithLevel logs a message with a specific log level
func (d *CustomLogger) logWithLevel(level string, args ...interface{}) {
	message := fmt.Sprint(args...)
	d.log.Print(formatLog(level, message))
}

// logWithLevelf logs a formatted message with a specific log level
func (d *CustomLogger) logWithLevelf(level string, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	d.log.Print(formatLog(level, message))
}

// Debug logs a message at DEBUG level
func (d *CustomLogger) Debug(ctx context.Context, args ...interface{}) {
	d.logWithLevel("DEBUG", args...)
}

// Debugf logs a formatted message at DEBUG level
func (d *CustomLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	d.logWithLevelf("DEBUG", format, args...)
}

// Info logs a message at INFO level
func (d *CustomLogger) Info(ctx context.Context, args ...interface{}) {
	d.logWithLevel("INFO", args...)
}

// Infof logs a formatted message at INFO level
func (d *CustomLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	d.logWithLevelf("INFO", format, args...)
}

// Warn logs a message at WARN level
func (d *CustomLogger) Warn(ctx context.Context, args ...interface{}) {
	d.logWithLevel("WARN", args...)
}

// Warnf logs a formatted message at WARN level
func (d *CustomLogger) Warnf(ctx context.Context, format string, args ...interface{}) {
	d.logWithLevelf("WARN", format, args...)
}

// Error logs a message at ERROR level
func (d *CustomLogger) Error(ctx context.Context, args ...interface{}) {
	d.logWithLevel("ERROR", args...)
}

// Errorf logs a formatted message at ERROR level
func (d *CustomLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	d.logWithLevelf("ERROR", format, args...)
}

// Fatal logs a message at FATAL level and exits the program
func (d *CustomLogger) Fatal(ctx context.Context, args ...interface{}) {
	d.logWithLevel("FATAL", args...)
	os.Exit(1) // Ensure the program exits after a fatal error
}

// Fatalf logs a formatted message at FATAL level and exits the program
func (d *CustomLogger) Fatalf(ctx context.Context, format string, args ...interface{}) {
	d.logWithLevelf("FATAL", format, args...)
	os.Exit(1) // Ensure the program exits after a fatal error
}
