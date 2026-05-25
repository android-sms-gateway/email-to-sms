package bridge

import "errors"

var (
	// ErrInvalidPhoneNumber is returned when the phone number contains non-digit characters.
	ErrInvalidPhoneNumber = errors.New("invalid phone number")

	// ErrEmptyBody is returned when the email body is empty.
	ErrEmptyBody = errors.New("email body is empty")
)
