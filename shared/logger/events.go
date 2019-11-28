package logger

// Declare variables to store log messages as new Events
var (
	configErrorMessage = Event{1, "Config setup error"}
	galErrorMessage    = Event{2, "GAL error: %s -- [%v]"}
	dalErrorMessage    = Event{3, "DAL error: %s -- [%v]"}
	utilsErrorMessage  = Event{4, "Utils error: %s -- [%v]"}
	marshErrorMessage  = Event{5, "Marsh error: %s -- [%v]"}
)
