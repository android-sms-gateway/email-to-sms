package smsgate

// Config holds configuration for the SMSGate module.
type Config struct {
	URL                 string
	SkipPhoneValidation bool
}
