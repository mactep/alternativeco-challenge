package consumers

import (
	thunderEvents "github.com/gothunder/thunder/pkg/events"
	"github.com/mactep/alternativeco-challenge/email/internal/features"
)

type ConsumerGroup struct {
	emailService features.EmailService
}

func newConsumerGroup(emailService features.EmailService) thunderEvents.Handler {
	return &ConsumerGroup{
		emailService: emailService,
	}
}
