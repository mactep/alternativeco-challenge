package transportoutbound

import (
	"github.com/mactep/alternativeco-challenge/email/internal/transport-outbound/publisher"
	"go.uber.org/fx"
)

var Module = fx.Options(
	publisher.Module,
)
