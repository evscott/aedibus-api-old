package logger

// Declare variables to store log messages as new Events
var (
	configErrorMessage = Event{1, "Config setup error"}
	galErrorMessage    = Event{2, "Github access layer error: %s"}
)
