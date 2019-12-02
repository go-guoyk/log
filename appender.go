package log

// Appender appender to the actual log facility
type Appender interface {
	// Log log a event, returns error
	Log(e Event) error
	// Close close the appender
	Close() error
}
