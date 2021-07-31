package logging

// Logger exposes a generic interface for logging
type Logger interface {
	Log(log string)
	Logf(formatString string, params ...interface{})
}
