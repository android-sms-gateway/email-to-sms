package config

import (
	"github.com/android-sms-gateway/email-to-sms/internal/smsgate"
	"github.com/android-sms-gateway/email-to-sms/internal/smtp"
	"github.com/go-core-fx/fiberfx"
	"github.com/go-core-fx/fiberfx/openapi"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"config",
		fx.Provide(New, fx.Private),
		fx.Provide(
			func(cfg Config) fiberfx.Config {
				return fiberfx.Config{
					Address:     cfg.HTTP.Address,
					ProxyHeader: cfg.HTTP.ProxyHeader,
					Proxies:     cfg.HTTP.Proxies,
				}
			},
			func(cfg Config) openapi.Config {
				return openapi.Config{
					Enabled:    cfg.HTTP.OpenAPI.Enabled,
					PublicHost: cfg.HTTP.OpenAPI.PublicHost,
					PublicPath: cfg.HTTP.OpenAPI.PublicPath,
				}
			},
		),
		fx.Provide(
			func(cfg Config) smtp.Config {
				return smtp.Config{
					Host:    cfg.SMTP.Host,
					Port:    cfg.SMTP.Port,
					Domain:  cfg.SMTP.Domain,
					TLSCert: cfg.SMTP.TLSCert,
					TLSKey:  cfg.SMTP.TLSKey,
				}
			},
			func(cfg Config) smsgate.Config {
				return smsgate.Config{
					URL:                 cfg.SMSGate.URL,
					SkipPhoneValidation: cfg.SMSGate.SkipPhoneValidation,
				}
			},
		),
	)
}
