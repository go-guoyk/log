package appender

import (
	"encoding/json"
	"fmt"
	"github.com/novakit/log/event"
	"io"
	"strings"
)

type consoleAppender struct {
	w io.Writer
}

func (a *consoleAppender) Log(e event.Event) error {
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

// Console create a console appender
func Console(w io.Writer) Appender {
	return &consoleAppender{
		w: w,
	}
}
