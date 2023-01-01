package transportinbound

import (
	"github.com/mactep/alternativeco-challenge/email/internal/transport-inbound/consumers"
	"github.com/mactep/alternativeco-challenge/email/internal/transport-inbound/router"
	"go.uber.org/fx"
)

var Module = fx.Options(
	consumers.Module,
	router.Module,
)
