package main

import (
	"io"
	"os"

	thunderEventRabbitmq "github.com/gothunder/thunder/pkg/events/rabbitmq"
	thunderLogs "github.com/gothunder/thunder/pkg/log"
	"github.com/mactep/alternativeco-challenge/ban/internal/features"
	transportinbound "github.com/mactep/alternativeco-challenge/ban/internal/transport-inbound"
	transportoutbound "github.com/mactep/alternativeco-challenge/ban/internal/transport-outbound"

	"github.com/rs/zerolog/diode"
	"go.uber.org/fx"
)

func main() {
	var w diode.Writer

	// read from file and return the io.Reader
	bannedDomainsSource, err := os.Open("banned_domains.txt")
	if err != nil {
		panic(err)
	}

	app := fx.New(
		// The order of these options isn't important.
		thunderLogs.Module,
		fx.Populate(&w),

		fx.Provide(
			func() io.Reader {
				return bannedDomainsSource
			},
		),

		transportinbound.Module,
		transportoutbound.Module,
		features.Module,

		thunderEventRabbitmq.PublisherModule,
		thunderEventRabbitmq.InvokeConsumer,
	)
	app.Run()

	// This is required to flush the logs to stdout.
	// We only want to do this after the app has exited.
	thunderLogs.DiodeShutdown(w)
}
