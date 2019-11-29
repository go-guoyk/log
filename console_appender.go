package log

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type consoleAppender struct {
	w      io.Writer
	filter *Filter
}

// NewConsoleAppender create a console appender
func NewConsoleAppender(w io.Writer, filter *Filter) Appender {
	return &consoleAppender{
		w:      w,
		filter: filter,
	}
}

func (a *consoleAppender) Log(e Event) error {
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

func (a *consoleAppender) Close() error {
	return nil
}
