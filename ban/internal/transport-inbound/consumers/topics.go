package consumers

import "github.com/mactep/alternativeco-challenge/ban/pkg/events"

func (c *ConsumerGroup) Topics() []string {
	return []string{
		events.TestTopic,
	}
}
