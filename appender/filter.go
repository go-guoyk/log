package appender

import (
	"github.com/novakit/log/event"
	"github.com/novakit/log/filter"
)

type filterAppender struct {
	filter   filter.Filter
	appender Appender
}

func (f *filterAppender) Log(e event.Event) error {
	if f.filter.Check(e) {
		return f.appender.Log(e)
	}
	return nil
}

func (f *filterAppender) Close() error {
	return f.appender.Close()
}

func Filter(f filter.Filter, a Appender) Appender {
	return &filterAppender{
		filter:   f,
		appender: a,
	}
}
