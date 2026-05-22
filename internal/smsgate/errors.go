package smsgate

import "errors"

var (
	// ErrAuthenticationFailed is returned when SMSGate returns 401/403.
	ErrAuthenticationFailed = errors.New("authentication failed")

	// ErrSendFailed is returned when SMSGate returns a 5xx error.
	ErrSendFailed = errors.New("failed to send SMS")

	// ErrTemporaryFailure is returned when SMSGate returns a 4xx error.
	ErrTemporaryFailure = errors.New("temporary failure")

	// ErrTimeout is returned when the SMSGate request times out.
	ErrTimeout = errors.New("request timeout")
)
