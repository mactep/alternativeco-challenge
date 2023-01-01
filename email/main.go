package main

import (
	thunderEventRabbitmq "github.com/gothunder/thunder/pkg/events/rabbitmq"
	thunderLogs "github.com/gothunder/thunder/pkg/log"
	thunderChi "github.com/gothunder/thunder/pkg/router/chi"
	transportinbound "github.com/mactep/alternativeco-challenge/email/internal/transport-inbound"
	transportoutbound "github.com/mactep/alternativeco-challenge/email/internal/transport-outbound"

	"github.com/rs/zerolog/diode"
	"go.uber.org/fx"
)

func main() {
	var w diode.Writer

	app := fx.New(
		fx.Populate(&w),
		thunderLogs.Module,
		thunderChi.Module,
		fx.Invoke(thunderChi.StartListener),

		transportinbound.Module,
		transportoutbound.Module,

		thunderEventRabbitmq.PublisherModule,
		thunderEventRabbitmq.InvokeConsumer,
	)
	app.Run()

	// This is required to flush the logs to stdout.
	// We only want to do this after the app has exited.
	thunderLogs.DiodeShutdown(w)
}
