package consumers

import (
	"context"

	thunderEvents "github.com/gothunder/thunder/pkg/events"
	"github.com/mactep/alternativeco-challenge/email/pkg/events"
	"github.com/rs/zerolog/log"
)

func (c *ConsumerGroup) emailEvent(ctx context.Context, payload events.EmailPayload) thunderEvents.HandlerResponse {
	log.Ctx(ctx).Info().Msgf("emailEvent: %v", payload)
	c.banService.CheckBan(ctx, payload.Email, payload.ID)

	return thunderEvents.Success
}
