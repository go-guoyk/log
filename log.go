package log

import (
	"context"
	"fmt"
	"time"
)

const (
	TopicDebug = "debug"
	TopicInfo  = "info"
	TopicError = "error"
)

var (
	activeProject = "noname"
	activeAdapter = SimpleAdapter()
)

func SetProject(project string) {
	activeProject = project
}

func SetAdapter(adapter Adapter) {
	activeAdapter = adapter
}

func Log(ctx context.Context, topic string, format string, items ...interface{}) {
	if len(items) == 0 {
		_ = activeAdapter.Log(time.Now(), activeProject, topic, GetAllLabels(ctx), format)
	} else {
		_ = activeAdapter.Log(time.Now(), activeProject, topic, GetAllLabels(ctx), fmt.Sprintf(format, items...))
	}
}

func Debug(ctx context.Context, format string, items ...interface{}) {
	Log(ctx, TopicDebug, format, items...)
}

func Info(ctx context.Context, format string, items ...interface{}) {
	Log(ctx, TopicInfo, format, items...)
}

func Error(ctx context.Context, format string, items ...interface{}) {
	Log(ctx, TopicError, format, items...)
}
