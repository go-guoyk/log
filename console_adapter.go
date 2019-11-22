package log

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type consoleAdapter struct {
	w      io.Writer
	filter *Filter
}

func NewConsoleAdapter(w io.Writer, filter *Filter) Adapter {
	return &consoleAdapter{
		w:      w,
		filter: filter,
	}
}

func (a *consoleAdapter) Log(e Event) error {
	if !a.filter.IsTopicEnabled(e.Topic) {
		return nil
	}
	var labels []byte
	if len(e.Labels) > 0 {
		labels, _ = json.Marshal(e.Labels)
		labels = append(labels, ' ')
	}
	_, err := fmt.Fprintf(
		a.w,
		"%s [%s] %s%s\n",
		e.Timestamp.Format("2006-01-02T15:04:05.000-0700"),
		e.Topic,
		labels,
		strings.TrimSpace(e.Message),
	)
	return err
}

func (a *consoleAdapter) Close() error {
	return nil
}
