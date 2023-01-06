package features

import (
	"bufio"
	"context"
	"io"

	"github.com/mactep/alternativeco-challenge/ban/internal/transport-outbound/publisher"
	"github.com/mactep/alternativeco-challenge/ban/pkg/events"
	"github.com/rs/zerolog/log"
)

type BanService struct {
	publisherGroup *publisher.PublisherGroup
	bannedDomains  map[string]interface{}
}

func NewBanService(publisherGroup *publisher.PublisherGroup, bannedDomainsSource io.Reader) BanService {
	bannedDomains := make(map[string]interface{})
	bannedDomainsIterator := bufio.NewScanner(bannedDomainsSource)
	for bannedDomainsIterator.Scan() {
		bannedDomains[bannedDomainsIterator.Text()] = nil
	}

	return BanService{
		publisherGroup: publisherGroup,
		bannedDomains:  bannedDomains,
	}
}

func (b BanService) CheckBan(ctx context.Context, domain string, id int) {
	log.Ctx(ctx).Debug().Msgf("Checking ban for '%s'", domain)
	_, isBanned := b.bannedDomains[domain]

	if isBanned {
		log.Ctx(ctx).Debug().Msgf("Domain '%s' is banned!", domain)
		b.publisherGroup.SendBanEvent(
			ctx,
			events.BanPayload{
				ID: id,
			},
		)
	} else {
		log.Ctx(ctx).Debug().Msgf("Domain '%s' is not banned!", domain)
	}
}
