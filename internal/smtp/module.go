package smtp

import (
	"context"

	"github.com/go-core-fx/logger"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module creates and returns an FX module for the SMTP package.
func Module() fx.Option {
	return fx.Module(
		"smtp",
		logger.WithNamedLogger("smtp"),

		fx.Provide(NewServer),

		fx.Invoke(func(lc fx.Lifecycle, server *Server, logger *zap.Logger) {
			lc.Append(fx.Hook{
				OnStart: func(_ context.Context) error {
					logger.Info("starting SMTP server", zap.String("addr", server.server.Addr))
					go func() {
						if err := server.server.ListenAndServe(); err != nil {
							logger.Error("SMTP server error", zap.Error(err))
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					logger.Info("stopping SMTP server")
					return server.server.Shutdown(ctx)
				},
			})
		}),
	)
}
