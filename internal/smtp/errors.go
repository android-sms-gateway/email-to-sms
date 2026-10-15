package smtp

import "errors"

var (
	// ErrIncompleteTLSConfig is returned when only one of the TLS certificate or key is configured.
	ErrIncompleteTLSConfig = errors.New("TLS certificate and key must be configured together")

	// ErrInvalidEmailFormat is returned when the recipient email address cannot be parsed.
	ErrInvalidEmailFormat = errors.New("invalid email format")

	// ErrDomainMismatch is returned when the email domain does not match the configured domain.
	ErrDomainMismatch = errors.New("email domain does not match configured domain")

	// ErrReadFailed is returned when reading the email data fails.
	ErrReadFailed = errors.New("failed to read email data")
)
