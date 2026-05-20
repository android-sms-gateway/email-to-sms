package smtp

import "errors"

var (
	// ErrInvalidEmailFormat is returned when the recipient email address cannot be parsed.
	ErrInvalidEmailFormat = errors.New("invalid email format")

	// ErrDomainMismatch is returned when the email domain does not match the configured domain.
	ErrDomainMismatch = errors.New("email domain does not match configured domain")

	// ErrReadFailed is returned when reading the email data fails.
	ErrReadFailed = errors.New("failed to read email data")
)
