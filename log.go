package log

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	activeProject     = "noname"
	activeEnv         = "noname"
	activeHostname, _ = os.Hostname()
	activeFilter      = &Filter{IsBlackList: true}
	activeAppenders   []Appender
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
	}
}

func setActiveFilter(filter []string) {
	activeFilter = NewFilter(filter)
}

func setActiveAppenders(appenders []Appender) {
	existedAppenders := activeAppenders
	activeAppenders = appenders
	for _, a := range existedAppenders {
		if err := a.Close(); err != nil {
			stderr("failed to close appender: %s", err.Error())
		}
	}
}

func Setup(opts Options) {
	setActiveProject(opts.Project)
	setActiveEnv(opts.Env)
	setActiveHostname(opts.Hostname)
	setActiveFilter(opts.Topics)

	var appenders []Appender
	if opts.Console != nil {
		if opts.Console.Enabled {
			appenders = append(appenders, NewConsoleAppender(os.Stdout, NewFilter(opts.Console.Topics)))
		}
	}
	if opts.File != nil {
		if opts.File.Enabled {
			if err := os.MkdirAll(opts.File.Dir, 0755); err != nil {
			}
			appenders = append(appenders, NewFileAppender(opts.File.Dir, NewFilter(opts.File.Topics)))
		}
	}
	setActiveAppenders(appenders)
}

// Loglf log a message with additional labels and format
func Loglf(ctx context.Context, topic string, addLabels Labels, format string, items ...interface{}) {
	if !activeFilter.IsTopicEnabled(topic) {
		return
	}
	e := Event{
		Timestamp: time.Now(),
		Hostname:  activeHostname,
		Project:   activeProject,
		Env:       activeEnv,
		Topic:     topic,
	}

	ctxLabels := GetAllLabels(ctx)
	if len(ctxLabels)+len(addLabels) > 0 {
		e.Labels = make(Labels)
		for k, v := range ctxLabels {
			e.Labels[k] = v
		}
		for k, v := range addLabels {
			e.Labels[k] = v
		}
	}

	if len(items) == 0 {
		e.Message = format
	} else {
		e.Message = fmt.Sprintf(format, items...)
	}
	for _, a := range activeAppenders {
		if err := a.Log(e); err != nil {
			stderr("failed to append appender: %s", err.Error())
		}
	}
}

// Logl log a message with additional labels
func Logl(ctx context.Context, topic string, addLabels Labels) {
	Loglf(ctx, topic, addLabels, "")
}

// Logf log a message with format
func Logf(ctx context.Context, topic string, format string, items ...interface{}) {
	Loglf(ctx, topic, nil, format, items...)
}

// Log log a message
func Log(ctx context.Context, topic string, message string) {
	Loglf(ctx, topic, nil, message)
}
