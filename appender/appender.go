package appender

import (
	"github.com/novakit/log/event"
)

// Appender appender to the actual log facility
type Appender interface {
	// Log log a event, returns error
	Log(e event.Event) error
	// Close close the appender
	Close() error
}
