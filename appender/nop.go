package appender

import (
	"github.com/novakit/log/event"
)

var nop nopAppender

type nopAppender struct {
}

func (n nopAppender) Log(e event.Event) error {
	return nil
}

func (n nopAppender) Close() error {
	return nil
}

func NOP() Appender {
	return nop
}
