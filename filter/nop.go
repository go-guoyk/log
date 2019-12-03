package filter

import (
	"github.com/novakit/log/event"
)

var nop nopFilter

type nopFilter struct {
}

func (nopFilter) Check(e event.Event) bool {
	return true
}

func NOP() Filter {
	return nop
}
