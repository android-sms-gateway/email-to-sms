package smtp

import (
	"github.com/emersion/go-smtp"
	"go.uber.org/zap"
)

// backend implements the go-smtp Backend interface.
type backend struct {
	config  Config
	handler Handler
	logger  *zap.Logger
}

// NewSession creates a new SMTP session.
func (b *backend) NewSession(_ *smtp.Conn) (smtp.Session, error) {
	return &session{
		backend:  b,
		logger:   b.logger,
		username: "",
		password: "",
		from:     "",
		to:       "",
	}, nil
}
