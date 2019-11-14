package adapter

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func SimpleAdapter() Adapter {
	return &simpleAdapter{}
}

type simpleAdapter struct{}

func (s *simpleAdapter) Log(t time.Time, project, topic string, labels map[string]interface{}, message string) error {
	var labelsJ []byte
	if len(labels) > 0 {
		labelsJ, _ = json.Marshal(labels)
		labelsJ = append(labelsJ, ' ')
	}
	_, err := fmt.Printf("%s [%s.%s] %s%s\n", t.Format(time.RFC3339), project, topic, labelsJ, strings.TrimSpace(message))
	return err
}
