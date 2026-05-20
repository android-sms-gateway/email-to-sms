package smtp

import "context"

// Handler processes validated emails from the SMTP server.
type Handler interface {
	Handle(ctx context.Context, email EmailData) error
}

// EmailData contains email fields extracted by the SMTP server.
type EmailData struct {
	Phone    string
	From     string
	Body     string
	Username string
	Password string
}
