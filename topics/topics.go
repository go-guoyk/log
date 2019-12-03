package topics

import (
	"strings"
)

type Topics interface {
	Contains(topic string) bool
}

type topics struct {
	black  bool
	topics map[string]bool
}

// New create a set of topics
func New(ts []string) Topics {
	f := &topics{
		topics: map[string]bool{},
	}
	if len(ts) == 0 {
		return f
	}

	if strings.HasPrefix(ts[0], "-") {
		f.black = true
	}

	for _, t := range ts {
		if strings.HasPrefix(t, "-") {
			f.topics[t[1:]] = true
		} else {
			f.topics[t] = true
		}
	}
	return f
}

func (f *topics) Contains(topic string) bool {
	if f == nil {
		return true
	}
	if f.black {
		return !f.topics[topic]
	}
	return f.topics[topic]
}
