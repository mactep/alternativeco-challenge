package features

import (
	"context"

	"github.com/mactep/alternativeco-challenge/ban/internal/transport-outbound/publisher"
	"github.com/mactep/alternativeco-challenge/ban/pkg/events"
	"github.com/rs/zerolog/log"
)

type BanService struct {
	publisherGroup *publisher.PublisherGroup
}

func NewBanService(publisherGroup *publisher.PublisherGroup) BanService {
	return BanService{
		publisherGroup: publisherGroup,
	}
}

func (b BanService) CheckBan(ctx context.Context, email string, id int) {
	log.Ctx(ctx).Info().Msgf("Checking ban for %s", email)
	isBanned := true

	if isBanned {
		log.Ctx(ctx).Info().Msg("Email is banned!")
		b.publisherGroup.SendBanEvent(
			ctx,
			events.BanPayload{
				ID: id,
			},
		)
	}
}
