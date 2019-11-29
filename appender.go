package log

// Appender appender to the actual log facility
type Appender interface {
	// Log log a event, returns error
	Log(e Event) error
	// Close close the appender
	Close() error
}

// DiscardAppender a appender does nothing
var DiscardAppender Appender = discardAppender{}

type discardAppender struct {
}

func (discardAppender) Log(e Event) error {
	return nil
}

func (discardAppender) Close() error {
	return nil
}
