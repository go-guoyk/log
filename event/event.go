package event

import (
	"time"
)

// Event log event
type Event struct {
	Timestamp time.Time              `json:"timestamp"`
	Project   string                 `json:"project"`
	Env       string                 `json:"env"`
	Hostname  string                 `json:"hostname"`
	Topic     string                 `json:"topic"`
	Labels    map[string]interface{} `json:"labels"`
	Message   string                 `json:"message"`
}
