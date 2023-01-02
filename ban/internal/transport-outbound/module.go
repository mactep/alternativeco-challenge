package transportoutbound

import (
	"github.com/mactep/alternativeco-challenge/ban/internal/transport-outbound/publisher"
	"go.uber.org/fx"
)

var Module = fx.Options(
	publisher.Module,
)
