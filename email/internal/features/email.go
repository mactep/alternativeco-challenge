package features

import (
	"context"
	"regexp"

	"github.com/mactep/alternativeco-challenge/email/internal/features/repository/ent"
	"github.com/rs/zerolog"
)

type EmailService struct {
	client *ent.Client
	logger *zerolog.Logger
}

func NewEmailService(client *ent.Client, logger *zerolog.Logger) EmailService {
	return EmailService{
		client: client,
		logger: logger,
	}
}

func (e EmailService) Create(ctx context.Context, email string) (*ent.Email, error) {
	emailRegistry, err := e.client.Email.Create().SetEmail(email).Save(ctx)
	if err != nil {
		e.logger.Error().Err(err).Msg("failed creating email")
		return nil, err
	}
	e.logger.Info().Msg("email created")
	return emailRegistry, nil
}

func (e EmailService) Delete(ctx context.Context, id int) error {
	err := e.client.Email.DeleteOneID(id).Exec(ctx)
	if err != nil {
		e.logger.Error().Err(err).Msg("failed deleting email")
		return err
	}
	e.logger.Info().Msg("email deleted")
	return nil
}

// IsValidEmail checks if the email is valid using a regex pattern
func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)
	return emailRegex.MatchString(email)
}
