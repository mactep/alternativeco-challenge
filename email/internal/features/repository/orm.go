package repository

import (
	"context"

	"github.com/mactep/alternativeco-challenge/email/internal/features/repository/ent"
	"github.com/rs/zerolog"

	_ "github.com/lib/pq"
)

func NewClient(logger *zerolog.Logger) *ent.Client {
	client := Connect(logger)
	defer client.Close()
	return client
}

func Connect(logger *zerolog.Logger) *ent.Client {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=email sslmode=disable password=password")
	if err != nil {
		logger.Error().Err(err).Msg("failed opening connection to postgres")
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		logger.Error().Err(err).Msg("failed creating schema resources")
	}

	logger.Info().Msg("connected to postgres")
	return client
}
