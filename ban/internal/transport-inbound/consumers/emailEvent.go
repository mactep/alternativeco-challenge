package consumers

import (
	"context"
	"strings"

	thunderEvents "github.com/gothunder/thunder/pkg/events"
	"github.com/mactep/alternativeco-challenge/email/pkg/events"
	"github.com/rs/zerolog/log"
)

func (c *ConsumerGroup) emailEvent(ctx context.Context, payload events.EmailPayload) thunderEvents.HandlerResponse {
	log.Ctx(ctx).Info().Msgf("emailEvent: %v", payload)
	c.banService.CheckBan(ctx, payload.Email, payload.ID)
	domain := strings.Split(payload.Email, "@")[1]
	c.banService.CheckBan(ctx, domain, payload.ID)

	return thunderEvents.Success
}
