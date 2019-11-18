package log

import "strings"

type Filter struct {
	Blacklist bool
	Topics    map[string]bool
}

func (f *Filter) IsTopicEnabled(topic string) bool {
	if f.Blacklist {
		return !f.Topics[topic]
	}
	return f.Topics[topic]
}

type Filters map[string]*Filter

func (fs Filters) IsTopicEnabled(scope string, topic string) bool {
	if fs == nil {
		return false
	}
	f := fs[scope]
	if f == nil {
		f = fs[DefaultScopeName]
	}
	if f == nil {
		return false
	}
	return f.IsTopicEnabled(topic)
}

func NewFilters(def map[string][]string) Filters {
	fs := Filters{}
	for k, v := range def {
		f := &Filter{
			Topics: map[string]bool{},
		}
		fs[k] = f
		if len(v) == 0 {
			continue
		}

		if strings.HasPrefix(v[0], "-") {
			f.Blacklist = true
		}

		for _, t := range v {
			if strings.HasPrefix(t, "-") {
				f.Topics[t[1:]] = true
			} else {
				f.Topics[t] = true
			}
		}
	}
	return fs
}
