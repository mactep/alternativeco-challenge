package transportinbound

import (
	"github.com/mactep/alternativeco-challenge/ban/internal/transport-inbound/consumers"
	"go.uber.org/fx"
)

var Module = fx.Options(
	consumers.Module,
)
