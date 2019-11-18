package log

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// SimpleAdapter returns the default console adapter
func SimpleAdapter() Adapter {
	return defaultSimpleAdapter
}

var defaultSimpleAdapter = simpleAdapter{
	w: os.Stdout,
}

type simpleAdapter struct {
	w io.Writer
}

func (a simpleAdapter) Log(e Event) error {
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
