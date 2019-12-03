package filter

import (
	"github.com/novakit/log/event"
)

type Filter interface {
	Check(e event.Event) bool
}
