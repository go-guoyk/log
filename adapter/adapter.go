package adapter

import (
	"time"
)

type Adapter interface {
	Log(t time.Time, project, topic string, labels map[string]interface{}, message string) error
}
