package appender

import (
	"errors"
	"github.com/novakit/log/event"
	"strings"
)

type multiAppender struct {
	appenders []Appender
}

func (m *multiAppender) Log(e event.Event) error {
	errs := make([]string, 0, len(m.appenders))
	for _, a := range m.appenders {
		if err := a.Log(e); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) == 0 {
		return nil
	} else {
		return errors.New(strings.Join(errs, "; "))
	}
}

func (m *multiAppender) Close() error {
	errs := make([]string, 0, len(m.appenders))
	for _, a := range m.appenders {
		if err := a.Close(); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) == 0 {
		return nil
	} else {
		return errors.New(strings.Join(errs, "; "))
	}
}

func Multi(a ...Appender) Appender {
	return &multiAppender{appenders: a}
}
