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

func (l *StandardLogger) GalError(message string, error error) {
	l.Errorf(galErrorMessage.message, message, error)
}

func (l *StandardLogger) DalError(message string, error error) {
	l.Errorf(dalErrorMessage.message, message, error)
}

func (l *StandardLogger) UtilsError(message string, error error) {
	l.Errorf(utilsErrorMessage.message, message, error)
}

func (l *StandardLogger) MarshError(message string, error error) {
	l.Errorf(marshErrorMessage.message, message, error)
}
