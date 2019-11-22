package log

import "strings"

type Filter struct {
	IsBlackList bool
	Topics      map[string]bool
}

func NewFilter(topics []string) (f *Filter) {
	f = &Filter{
		Topics: map[string]bool{},
	}
	if len(topics) == 0 {
		return
	}

	if strings.HasPrefix(topics[0], "-") {
		f.IsBlackList = true
	}

	for _, t := range topics {
		if strings.HasPrefix(t, "-") {
			f.Topics[t[1:]] = true
		} else {
			f.Topics[t] = true
		}
	}
	return
}

func (f *Filter) IsTopicEnabled(topic string) bool {
	if f == nil {
		return true
	}
	if f.IsBlackList {
		return !f.Topics[topic]
	}
	return f.Topics[topic]
}
