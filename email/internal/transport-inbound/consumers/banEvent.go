package consumers

import (
	"context"

	thunderEvents "github.com/gothunder/thunder/pkg/events"
	"github.com/mactep/alternativeco-challenge/ban/pkg/events"
	"github.com/rs/zerolog/log"
)

func (c *ConsumerGroup) banEvent(ctx context.Context, payload events.BanPayload) thunderEvents.HandlerResponse {
	log.Ctx(ctx).Info().Msgf("banEvent: %v", payload)
	id := payload.ID
	err := c.emailService.Delete(ctx, id)
	if err != nil {
		return thunderEvents.DeadLetter
	}

	return thunderEvents.Success
}
