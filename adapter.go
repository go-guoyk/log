package log

// Adapter adapter to the actual log facility
type Adapter interface {
	// Log log a event, returns error
	Log(e Event) error
	// Close close the adapter
	Close() error
}
