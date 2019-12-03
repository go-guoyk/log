package log

import (
	"context"
	"github.com/novakit/log/labels"
)

// Error shorthand for Log with topic error
func Error(ctx context.Context, message string) {
	Log(ctx, "error", message)
}

// Errorf shorthand for Logf with topic error
func Errorf(ctx context.Context, format string, items ...interface{}) {
	Logf(ctx, "error", format, items...)
}

// Errorl shorthand for Logl with topic error
func Errorl(ctx context.Context, l labels.Labels, merge bool) {
	Logl(ctx, "error", l, merge)
}

// Errorl shorthand for Loglf with topic error
func Errorlf(topic string, l labels.Labels, format string, items ...interface{}) {
	Loglf("error", l, format, items...)
}

// Info shorthand for Log with topic info
func Info(ctx context.Context, message string) {
	Log(ctx, "info", message)
}

// Infof shorthand for Logf with topic info
func Infof(ctx context.Context, format string, items ...interface{}) {
	Logf(ctx, "info", format, items...)
}

// Infol shorthand for Logl with topic info
func Infol(ctx context.Context, l labels.Labels, merge bool) {
	Logl(ctx, "info", l, merge)
}

// Infol shorthand for Loglf with topic info
func Infolf(topic string, l labels.Labels, format string, items ...interface{}) {
	Loglf("info", l, format, items...)
}

// Debug shorthand for Log with topic debug
func Debug(ctx context.Context, message string) {
	Log(ctx, "debug", message)
}

// Debugf shorthand for Logf with topic debug
func Debugf(ctx context.Context, format string, items ...interface{}) {
	Logf(ctx, "debug", format, items...)
}

// Debugl shorthand for Logl with topic debug
func Debugl(ctx context.Context, l labels.Labels, merge bool) {
	Logl(ctx, "debug", l, merge)
}

// Debugl shorthand for Loglf with topic debug
func Debuglf(topic string, l labels.Labels, format string, items ...interface{}) {
	Loglf("debug", l, format, items...)
}
