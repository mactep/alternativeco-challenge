package consumers

import (
	"context"

	thunderEvents "github.com/gothunder/thunder/pkg/events"
	"github.com/mactep/alternativeco-challenge/email/pkg/events"
	"github.com/rs/zerolog/log"
)

func (c *ConsumerGroup) Handle(ctx context.Context, topic string, decoder thunderEvents.EventDecoder) thunderEvents.HandlerResponse {
	switch {
	case topic == events.EmailTopic:
		log.Ctx(ctx).Info().Msg("got email event")
		var formattedPayload events.EmailPayload
		err := decoder.Decode(&formattedPayload)
		if err != nil {
			log.Ctx(ctx).Error().Err(err).Msg("failed to decode payload")
			return thunderEvents.DeadLetter
		}

		return c.emailEvent(ctx, formattedPayload)
	default:
		return thunderEvents.DeadLetter
	}
}
