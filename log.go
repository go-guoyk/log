package log

import (
	"context"
	"fmt"
	"os"
	"time"
)

// Event log event
type Event struct {
	Timestamp time.Time `json:"timestamp"`
	Project   string    `json:"project"`
	Env       string    `json:"env"`
	Hostname  string    `json:"hostname"`
	Topic     string    `json:"topic"`
	Labels    Labels    `json:"labels"`
	Message   string    `json:"message"`
}

// Adapter adapter to the actual log facility
type Adapter interface {
	// Log log a event, returns error
	Log(e Event) error
}

var (
	activeProject     = "noname"
	activeEnv         = "noname"
	activeHostname, _ = os.Hostname()
	activeAdapter     = ConsoleAdapter()
)

// SetProject set project
func SetProject(project string) {
	activeProject = project
}

// SetEnv set environment
func SetEnv(env string) {
	activeEnv = env
}

// SetHostname set hostname
func SetHostname(hostname string) {
	activeHostname = hostname
}

// SetAdapter set adapter
func SetAdapter(adapter Adapter) {
	activeAdapter = adapter
}

// Loglf log a message with additional labels and format
func Loglf(ctx context.Context, topic string, addLabels Labels, format string, items ...interface{}) {
	e := Event{
		Timestamp: time.Now(),
		Hostname:  activeHostname,
		Project:   activeProject,
		Env:       activeEnv,
		Topic:     topic,
	}

	ctxLabels := GetAllLabels(ctx)
	if len(ctxLabels)+len(addLabels) > 0 {
		e.Labels = make(Labels)
		for k, v := range ctxLabels {
			e.Labels[k] = v
		}
		for k, v := range addLabels {
			e.Labels[k] = v
		}
	}

	if len(items) == 0 {
		e.Message = format
	} else {
		e.Message = fmt.Sprintf(format, items...)
	}
	_ = activeAdapter.Log(e)
}

// Logl log a message with additional labels
func Logl(ctx context.Context, topic string, addLabels Labels) {
	Loglf(ctx, topic, addLabels, "")
}

// Logf log a message with format
func Logf(ctx context.Context, topic string, format string, items ...interface{}) {
	Loglf(ctx, topic, nil, format, items...)
}

// Log log a message
func Log(ctx context.Context, topic string, message string) {
	Loglf(ctx, topic, nil, message)
}
