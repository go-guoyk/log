package log

import (
	"context"
	"fmt"
	"time"
)

type Scope struct {
	Name string
}

func NewScope(name string) *Scope {
	return &Scope{Name: name}
}

// Loglf log a message with additional labels and format
func (s *Scope) Loglf(ctx context.Context, topic string, addLabels Labels, format string, items ...interface{}) {
	if !activeFilters.IsTopicEnabled(s.Name, topic) {
		return
	}
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
	for _, a := range activeAdapters {
		_ = a.Log(e)
	}
}

// Logl log a message with additional labels
func (s *Scope) Logl(ctx context.Context, topic string, addLabels Labels) {
	s.Loglf(ctx, topic, addLabels, "")
}

// Logf log a message with format
func (s *Scope) Logf(ctx context.Context, topic string, format string, items ...interface{}) {
	s.Loglf(ctx, topic, nil, format, items...)
}

// Log log a message
func (s *Scope) Log(ctx context.Context, topic string, message string) {
	s.Loglf(ctx, topic, nil, message)
}
