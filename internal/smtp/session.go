package smtp

import (
	"context"
	"fmt"
	"io"
	"net/mail"
	"strings"

	"github.com/emersion/go-smtp"
	"go.uber.org/zap"
)

// session handles a single SMTP connection.
type session struct {
	backend *backend
	logger  *zap.Logger

	username string
	password string
	from     string
	to       string
}

// AuthPlain handles SMTP AUTH PLAIN authentication.
func (s *session) AuthPlain(username, password string) error {
	s.username = username
	s.password = password
	return nil
}

// Mail handles the MAIL FROM command.
func (s *session) Mail(from string, _ *smtp.MailOptions) error {
	s.from = from
	return nil
}

// Rcpt handles the RCPT TO command.
func (s *session) Rcpt(addr string, _ *smtp.RcptOptions) error {
	s.to = addr
	return nil
}

// Data handles the DATA command, receiving the full email and processing it.
func (s *session) Data(r io.Reader) error {
	const maxMessageBytes = 1 << 20 // 1 MiB
	msg, err := io.ReadAll(io.LimitReader(r, maxMessageBytes+1))
	if err != nil {
		return ErrReadFailed
	}
	if len(msg) > maxMessageBytes {
		return ErrReadFailed
	}

	addr, err := mail.ParseAddress(s.to)
	if err != nil {
		return ErrInvalidEmailFormat
	}

	localPart, domain, found := strings.Cut(addr.Address, "@")
	if !found {
		return ErrInvalidEmailFormat
	}

	if !strings.EqualFold(domain, s.backend.config.Domain) {
		return ErrDomainMismatch
	}

	phone := strings.TrimSpace(localPart)

	rawEmail, err := mail.ReadMessage(strings.NewReader(string(msg)))
	if err != nil {
		return ErrInvalidEmailFormat
	}

	from := rawEmail.Header.Get("From")
	body, err := io.ReadAll(rawEmail.Body)
	if err != nil {
		return ErrReadFailed
	}

	ctx := context.Background()
	emailData := EmailData{
		Phone:    phone,
		From:     from,
		Body:     string(body),
		Username: s.username,
		Password: s.password,
	}

	if handleErr := s.backend.handler.Handle(ctx, emailData); handleErr != nil {
		return fmt.Errorf("handler failed: %w", handleErr)
	}
	return nil
}

// Reset handles the RSET command.
func (s *session) Reset() {
	s.username = ""
	s.password = ""
	s.from = ""
	s.to = ""
}

// Logout handles session cleanup.
func (s *session) Logout() error {
	s.Reset()
	return nil
}
