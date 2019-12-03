package log

import (
	"context"
	"fmt"
	"github.com/novakit/log/appender"
	"github.com/novakit/log/event"
	"github.com/novakit/log/filter"
	"github.com/novakit/log/labels"
	"os"
	"strings"
	"time"
)

var (
	activeProject     = "unknown"
	activeEnv         = "unknown"
	activeHostname, _ = os.Hostname()
	activeAppender    = appender.NOP()
)

func stderr(format string, items ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, "[novakit/log] "+format+"\n", items...)
}

func setActiveProject(project string) {
	if project = strings.TrimSpace(project); len(project) > 0 {
		activeProject = project
	}
}

func setActiveEnv(env string) {
	if env = strings.TrimSpace(env); len(env) > 0 {
		activeEnv = env
	}
}

func setActiveHostname(hostname string) {
	if hostname = strings.TrimSpace(hostname); len(hostname) > 0 {
		activeHostname = hostname
	} else if activeHostname, _ = os.Hostname(); len(hostname) > 0 {
		activeHostname = "unknown"
	}
}

func setActiveAppender(a appender.Appender) {
	if a == nil {
		a = appender.NOP()
	}
	oldAppender := activeAppender
	activeAppender = a
	_ = oldAppender.Close()
}

func Setup(opts Options) {
	setActiveProject(opts.Project)
	setActiveEnv(opts.Env)
	setActiveHostname(opts.Hostname)

	var appenders []appender.Appender
	if opts.Console != nil {
		if opts.Console.Enabled {
			appenders = append(appenders, appender.Filter(filter.Topic(opts.Console.Topics), appender.Console(os.Stdout)))
		}
	}
	if opts.File != nil {
		if opts.File.Enabled {
			if err := os.MkdirAll(opts.File.Dir, 0755); err != nil {
			}
			appenders = append(appenders, appender.Filter(filter.Topic(opts.File.Topics), appender.File(opts.File.Dir)))
		}
	}
	setActiveAppender(appender.Filter(filter.Topic(opts.Topics), appender.Multi(appenders...)))
}

func AutoSetup(load func(name string, out interface{}) error) error {
	var opts Options
	if err := load("log", &opts); err != nil {
		return nil
	}
	Setup(opts)
	return nil
}

// log log a message with additional labels and format
func Loglf(topic string, l labels.Labels, format string, items ...interface{}) {
	// build event
	e := event.Event{
		Timestamp: time.Now(),
		Hostname:  activeHostname,
		Project:   activeProject,
		Env:       activeEnv,
		Topic:     topic,
		Labels:    l,
	}

	// build message
	if len(items) == 0 {
		e.Message = format
	} else {
		e.Message = fmt.Sprintf(format, items...)
	}

	if err := activeAppender.Log(e); err != nil {
		stderr("failed to append appender: %s", err.Error())
	}
}

// Logl log a message with label
func Logl(ctx context.Context, topic string, l labels.Labels, merge bool) {
	if merge {
		l = l.Merge(labels.GetAll(ctx))
	}
	Loglf(topic, l, "")
}

// Logf log a message with format
func Logf(ctx context.Context, topic string, format string, items ...interface{}) {
	Loglf(topic, labels.GetAll(ctx), format, items...)
}

// Log log a message
func Log(ctx context.Context, topic string, message string) {
	Loglf(topic, labels.GetAll(ctx), message)
}
