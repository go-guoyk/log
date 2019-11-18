package log

import "time"

// Event log event
type Event struct {
	Timestamp time.Time `json:"timestamp"`
	Project   string    `json:"project"`
	Env       string    `json:"env"`
	Hostname  string    `json:"hostname"`
	Scope     string    `json:"scope"`
	Topic     string    `json:"topic"`
	Labels    Labels    `json:"labels"`
	Message   string    `json:"message"`
}
