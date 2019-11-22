package log

import "context"

// Error shorthand for Log with topic error
func Error(ctx context.Context, message string) {
	Log(ctx, "error", message)
}

// Errorf shorthand for Logf with topic error
func Errorf(ctx context.Context, format string, items ...interface{}) {
	Logf(ctx, "error", format, items...)
}

// Errorl shorthand for Logl with topic error
func Errorl(ctx context.Context, addLabels Labels) {
	Logl(ctx, "error", addLabels)
}

// Errorlf shorthand for Loglf with topic error
func Errorlf(ctx context.Context, addLabels Labels, format string, items ...interface{}) {
	Loglf(ctx, "error", addLabels, format, items...)
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
func Infol(ctx context.Context, addLabels Labels) {
	Logl(ctx, "info", addLabels)
}

// Infolf shorthand for Loglf with topic info
func Infolf(ctx context.Context, addLabels Labels, format string, items ...interface{}) {
	Loglf(ctx, "info", addLabels, format, items...)
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
func Debugl(ctx context.Context, addLabels Labels) {
	Logl(ctx, "debug", addLabels)
}

// Debuglf shorthand for Loglf with topic debug
func Debuglf(ctx context.Context, addLabels Labels, format string, items ...interface{}) {
	Loglf(ctx, "debug", addLabels, format, items...)
}
