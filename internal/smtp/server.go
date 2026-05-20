package smtp

import (
	"crypto/tls"
	"fmt"

	gosmtp "github.com/emersion/go-smtp"
	"go.uber.org/zap"
)

// Server wraps the go-smtp server with handler integration.
type Server struct {
	server  *gosmtp.Server
	backend *backend
}

// NewServer creates and configures the SMTP server.
func NewServer(cfg Config, h Handler, logger *zap.Logger) (*Server, error) {
	b := &backend{
		config:  cfg,
		handler: h,
		logger:  logger,
	}

	srv := gosmtp.NewServer(b)
	srv.Addr = fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	if cfg.TLSCert != "" && cfg.TLSKey != "" {
		cert, err := tls.LoadX509KeyPair(cfg.TLSCert, cfg.TLSKey)
		if err != nil {
			return nil, fmt.Errorf("failed to load TLS certificate: %w", err)
		}

		srv.TLSConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
			MinVersion:   tls.VersionTLS12,
		}
	}

	return &Server{
		server:  srv,
		backend: b,
	}, nil
}
