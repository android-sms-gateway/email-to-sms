package smsgate

import (
	"github.com/go-core-fx/logger"
	"go.uber.org/fx"
)

// Module creates and returns an FX module for the SMSGate package.
func Module() fx.Option {
	return fx.Module(
		"smsgate",
		logger.WithNamedLogger("smsgate"),

		fx.Provide(NewMetrics, fx.Private),
		fx.Provide(NewService),
	)
}
