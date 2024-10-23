package log

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Logger wraps logrus.Logger and adds custom logging methods
type Logger struct {
	logrus *logrus.Logger
}

// NewLogger creates a new instance of Logger
func NewLogger() *Logger {
	logger := logrus.New()

	// Set log format if needed, for example as TextFormatter or JSONFormatter
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	return &Logger{logrus: logger}
}

// Error logs an error message with multiple variables
func (l *Logger) Error(message string, variables ...interface{}) {
	l.logrus.Error(formatMessage(message, variables...))
}

// Warn logs a warning message with multiple variables
func (l *Logger) Warn(message string, variables ...interface{}) {
	l.logrus.Warn(formatMessage(message, variables...))
}

// Info logs an info message with multiple variables
func (l *Logger) Info(message string, variables ...interface{}) {
	l.logrus.Info(formatMessage(message, variables...))
}

func (l *Logger) Debug(message string, variables ...interface{}) {
	l.logrus.Debug(formatMessage(message, variables...))
}

// formatMessage formats the message and appends variables
func formatMessage(message string, variables ...interface{}) string {
	if len(variables) > 0 {
		return fmt.Sprintf(message, variables...)
	}
	return message
}
