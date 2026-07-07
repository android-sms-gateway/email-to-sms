package smtp

import (
	"time"

	"github.com/emersion/go-smtp"
	"go.uber.org/zap"
)

// backend implements the go-smtp Backend interface.
type backend struct {
	config  Config
	handler Handler
	metrics *Metrics
	logger  *zap.Logger
}

// NewSession creates a new SMTP session.
func (b *backend) NewSession(_ *smtp.Conn) (smtp.Session, error) {
	b.metrics.IncActiveConnections()
	b.metrics.IncConnectionsTotal()
	return &session{
		backend:   b,
		metrics:   b.metrics,
		logger:    b.logger,
		username:  "",
		password:  "",
		from:      "",
		to:        "",
		startTime: time.Now(),
	}, nil
}
