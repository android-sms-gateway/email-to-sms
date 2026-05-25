package bridge

import (
	"github.com/android-sms-gateway/email-to-sms/internal/smtp"
	"github.com/go-core-fx/logger"
	"go.uber.org/fx"
)

// Module creates and returns an FX module for the bridge package.
func Module() fx.Option {
	return fx.Module(
		"bridge",
		logger.WithNamedLogger("bridge"),

		fx.Provide(NewMetrics, fx.Private),
		fx.Provide(NewService),
		fx.Provide(func(s *Service) smtp.Handler { return s }),
	)
}
