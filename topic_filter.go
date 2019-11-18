package log

type TopicFilter map[string][]string

func (tf TopicFilter) IsTopicEnabled(scope string, topic string) bool {
	return false
}
