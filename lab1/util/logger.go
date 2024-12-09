package logger

import (
	"context"
	"fmt"
	"runtime"
	"strings"
)

// Logger interface definition
type Logger interface {
	Debug(ctx context.Context, args ...interface{})
	Debugf(ctx context.Context, format string, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Warn(ctx context.Context, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	Fatal(ctx context.Context, args ...interface{})
	Fatalf(ctx context.Context, format string, args ...interface{})
}

// by default i use zap logger
var defaultLogger = NewDefaultLogger()

func SetLogger(logger Logger) {
	defaultLogger = logger
}

// Color constants for log levels
const (
	green  = "\033[32m"
	yellow = "\033[33m"
	red    = "\033[31m"
	black  = "\033[30m"
	reset  = "\033[0m"
)

// formatMessage formats the log message with timestamp, log level, goroutine count, and file info
func formatMessage() string {
	_, file, line, _ := runtime.Caller(2) // Caller depth adjusted for correct file and line number
	numGoroutines := runtime.NumGoroutine()
	fileInfo := fmt.Sprintf("%s:%d", file, line)
	return fmt.Sprintf("[%d] [%s] - ", numGoroutines, fileInfo)
}

// Info logs an informational message
func Info(ctx context.Context, args ...interface{}) {
	combinedArgs := append([]interface{}{formatMessage()}, args...)
	defaultLogger.Info(ctx, combinedArgs...)
}

// Infof logs a formatted informational message
func Infof(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.Infof(ctx, formatMessage()+format, args...)
}

// Debug logs a debug message
func Debug(ctx context.Context, args ...interface{}) {
	combinedArgs := append([]interface{}{formatMessage()}, args...)
	defaultLogger.Debug(ctx, combinedArgs...)
}

// Debugf logs a formatted debug message
func Debugf(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.Debugf(ctx, formatMessage()+format, args...)
}

// Warn logs a warning message
func Warn(ctx context.Context, args ...interface{}) {
	combinedArgs := append([]interface{}{formatMessage()}, args...)
	defaultLogger.Warn(ctx, combinedArgs...)
}

// Warnf logs a formatted warning message
func Warnf(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.Warnf(ctx, formatMessage()+format, args...)
}

// Error logs an error message
func Error(ctx context.Context, args ...interface{}) {
	combinedArgs := append([]interface{}{formatMessage()}, args...)
	defaultLogger.Error(ctx, combinedArgs...)
}

// Errorf logs a formatted error message
func Errorf(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.Errorf(ctx, formatMessage()+format, args...)
}

// Fatal logs a fatal message and exits the program
func Fatal(ctx context.Context, args ...interface{}) {
	combinedArgs := append([]interface{}{formatMessage()}, args...)
	defaultLogger.Fatal(ctx, combinedArgs...)
}

// Fatalf logs a formatted fatal message and exits the program
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.Fatalf(ctx, formatMessage()+format, args...)
}

func getColor(level string) string {
	var color string
	switch strings.ToUpper(level) {
	case "DEBUG":
		color = "\033[37m" // Gray
	case "INFO":
		color = "\033[32m" // Green
	case "WARN":
		color = "\033[33m" // Yellow
	case "ERROR":
		color = "\033[31m" // Red
	case "FATAL":
		color = "\033[35m" // Magenta
	default:
		color = "\033[0m" // Default color (reset)
	}
	return color
}
