package log

import (
	"context"
	"fmt"
	"github.com/novakit/log/appender"
	"github.com/novakit/log/event"
	"github.com/novakit/log/labels"
	"github.com/novakit/log/topics"
	"os"
	"strings"
	"time"
)

var (
	activeProject     = "unknown"
	activeEnv         = "unknown"
	activeHostname, _ = os.Hostname()
	activeTopics      = topics.All()
	activeAppenders   []AppenderRegistration
)

type AppenderRegistration struct {
	Topics   topics.Topics
	Appender appender.Appender
}

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

func setActiveTopics(ts topics.Topics) {
	if ts == nil {
		ts = topics.All()
	}
	activeTopics = ts
}

func setActiveAppender(as []AppenderRegistration) {
	oldAppenders := activeAppenders
	activeAppenders = as
	for _, r := range oldAppenders {
		if err := r.Appender.Close(); err != nil {
			stderr("failed to close appender: %s", err.Error())
		}
	}
}

func Setup(opts Options) {
	setActiveProject(opts.Project)
	setActiveEnv(opts.Env)
	setActiveHostname(opts.Hostname)
	setActiveTopics(topics.New(opts.Topics))

	var appenders []AppenderRegistration
	if opts.Console != nil {
		if opts.Console.Enabled {
			appenders = append(appenders, AppenderRegistration{
				Topics:   topics.New(opts.Console.Topics),
				Appender: appender.Console(os.Stdout),
			})
		}
	}
	if opts.File != nil {
		if opts.File.Enabled {
			if err := os.MkdirAll(opts.File.Dir, 0755); err != nil {
				stderr("failed to create log directory '%s': %s", opts.File.Dir, err.Error())
			} else {
				appenders = append(appenders, AppenderRegistration{
					Topics:   topics.New(opts.File.Topics),
					Appender: appender.File(opts.File.Dir),
				})
			}
		}
	}
	setActiveAppender(appenders)
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
	if !activeTopics.Contains(topic) {
		return
	}

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

	for _, r := range activeAppenders {
		if !r.Topics.Contains(topic) {
			continue
		}

		if err := r.Appender.Log(e); err != nil {
			stderr("failed to append appender: %s", err.Error())
		}
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
