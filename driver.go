package log

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Driver interface {
	Log(t time.Time, project, topic string, meta map[string]interface{}, message string) error
}

func SimpleDriver() Driver {
	return &simpleDriver{}
}

type simpleDriver struct{}

func (s *simpleDriver) Log(t time.Time, project, topic string, labels map[string]interface{}, message string) error {
	var mLabels []byte
	if len(labels) > 0 {
		mLabels, _ = json.Marshal(labels)
		mLabels = append(mLabels, ' ')
	}
	_, err := fmt.Printf("%s [%s.%s] %s%s\n", t.Format(time.RFC3339), project, topic, mLabels, strings.TrimSpace(message))
	return err
}
