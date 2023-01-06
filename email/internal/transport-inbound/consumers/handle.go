package consumers

import (
	"context"

	thunderEvents "github.com/gothunder/thunder/pkg/events"
	"github.com/mactep/alternativeco-challenge/ban/pkg/events"
	emailEvents "github.com/mactep/alternativeco-challenge/email/pkg/events"
	"github.com/rs/zerolog/log"
)

func (c *ConsumerGroup) Handle(ctx context.Context, topic string, decoder thunderEvents.EventDecoder) thunderEvents.HandlerResponse {
	switch {
	case topic == events.BanTopic:
		log.Ctx(ctx).Info().Msg("got ban event")
		var formattedPayload events.BanPayload
		err := decoder.Decode(&formattedPayload)
		if err != nil {
			log.Ctx(ctx).Error().Err(err).Msg("failed to decode payload")
			return thunderEvents.DeadLetter
		}

		return c.banEvent(ctx, formattedPayload)
	case topic == emailEvents.EmailTopic:
		log.Ctx(ctx).Info().Msg("got email event")
		return thunderEvents.Retry
	default:
		return thunderEvents.DeadLetter
	}
}
