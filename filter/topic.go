package filter

import (
	"github.com/novakit/log/event"
	"strings"
)

type topicFilter struct {
	isBlackList bool
	topics      map[string]bool
}

// Topic create a topic filter
func Topic(topics []string) Filter {
	f := &topicFilter{
		topics: map[string]bool{},
	}
	if len(topics) == 0 {
		return f
	}

	if strings.HasPrefix(topics[0], "-") {
		f.isBlackList = true
	}

	for _, t := range topics {
		if strings.HasPrefix(t, "-") {
			f.topics[t[1:]] = true
		} else {
			f.topics[t] = true
		}
	}
	return f
}

func (f *topicFilter) Check(e event.Event) bool {
	if f == nil {
		return true
	}
	if f.isBlackList {
		return !f.topics[e.Topic]
	}
	return f.topics[e.Topic]
}
