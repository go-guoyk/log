package topics

var all allTopics

type allTopics struct {
}

func (allTopics) Contains(topic string) bool {
	return true
}

func All() Topics {
	return all
}
