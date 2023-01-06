package consumers

import "github.com/mactep/alternativeco-challenge/email/pkg/events"

func (c *ConsumerGroup) Topics() []string {
	return []string{
		events.EmailTopic,
	}
}
