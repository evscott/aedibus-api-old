package logger

import (
	"github.com/sirupsen/logrus"
)

// Event stores messages to log later, from our standard interface
type Event struct {
	id      int
	message string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {
	var baseLogger = logrus.New()

	var standardLogger = &StandardLogger{baseLogger}

	standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}

func (l *StandardLogger) ConfigError(error error) {
	l.Fatalf(configErrorMessage.message, error)
}

func (l *StandardLogger) GalError(error error) {
	l.Errorf(galErrorMessage.message, error)
}
