package log

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type consoleAdapter struct {
	w       io.Writer
	filters Filters
}

func NewConsoleAdapter(w io.Writer, filters Filters) Adapter {
	return &consoleAdapter{
		w:       w,
		filters: filters,
	}
}

func (a consoleAdapter) Log(e Event) error {
	if !a.filters.IsTopicEnabled(e.Scope, e.Topic) {
		return nil
	}
	var labels []byte
	if len(e.Labels) > 0 {
		labels, _ = json.Marshal(e.Labels)
		labels = append(labels, ' ')
	}
	_, err := fmt.Fprintf(
		a.w,
		"%s [%s:%s:%s:%s] %s%s\n",
		e.Timestamp.Format("2006-01-02T15:04:05.999-0700"),
		e.Project,
		e.Env,
		e.Hostname,
		e.Topic,
		labels,
		strings.TrimSpace(e.Message),
	)
	return err
}
