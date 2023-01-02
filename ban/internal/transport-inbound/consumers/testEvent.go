package consumers

import (
	"context"

	thunderEvents "github.com/gothunder/thunder/pkg/events"
	"github.com/mactep/alternativeco-challenge/ban/pkg/events"
)

func (c *ConsumerGroup) testEvent(ctx context.Context, payload events.TestPayload) thunderEvents.HandlerResponse {
	// implement your logic here

	return thunderEvents.Success
}
