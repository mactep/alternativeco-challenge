package publisher

import (
	"context"

	"github.com/mactep/alternativeco-challenge/ban/pkg/events"
)

func (pg *PublisherGroup) SendTestEvent(ctx context.Context, event events.TestPayload) error {
	return pg.publisher.Publish(ctx, events.TestTopic, event)
}
