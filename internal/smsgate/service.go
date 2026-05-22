package smsgate

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/android-sms-gateway/client-go/rest"
	"github.com/android-sms-gateway/client-go/smsgateway"
	"go.uber.org/zap"
)

// Service wraps the SMSGate API client for sending SMS messages.
type Service struct {
	baseURL             string
	skipPhoneValidation bool
	metrics             *Metrics
	logger              *zap.Logger
}

// NewService creates and initializes a new Sender instance.
func NewService(cfg Config, metrics *Metrics, logger *zap.Logger) *Service {
	return &Service{
		baseURL:             cfg.URL,
		skipPhoneValidation: cfg.SkipPhoneValidation,
		metrics:             metrics,
		logger:              logger,
	}
}

// Send sends an SMS via SMSGate using the provided credentials.
func (s *Service) Send(ctx context.Context, phone, message, username, password string) error {
	s.logger.Info("sending SMS",
		zap.String("username", username),
	)

	client := smsgateway.NewClient(smsgateway.Config{
		BaseURL:  s.baseURL,
		User:     username,
		Password: password,
		Client:   nil,
		Token:    "",
	})

	msg := smsgateway.Message{
		PhoneNumbers: []string{phone},
		TextMessage: &smsgateway.TextMessage{
			Text: message,
		},
		ID:                 "",
		DeviceID:           "",
		Message:            "",
		DataMessage:        nil,
		IsEncrypted:        false,
		SimNumber:          nil,
		WithDeliveryReport: nil,
		Priority:           smsgateway.PriorityDefault,
		TTL:                nil,
		ValidUntil:         nil,
		ScheduleAt:         nil,
	}

	opts := []smsgateway.SendOption{}
	if s.skipPhoneValidation {
		opts = append(opts, smsgateway.WithSkipPhoneValidation(true))
	}

	if _, err := client.Send(ctx, msg, opts...); err != nil {
		s.logger.Error("SMSGate send failed", zap.Error(err))
		return s.mapError(err)
	}

	s.metrics.IncSent()
	return nil
}

// SendWithTimeout sends an SMS with a configurable timeout.
func (s *Service) SendWithTimeout(
	ctx context.Context, phone, message, username, password string, timeout time.Duration,
) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return s.Send(ctx, phone, message, username, password)
}

// mapError maps SMSGate client errors to module-specific sentinel errors.
func (s *Service) mapError(err error) error {
	if isAuthFailure(err) {
		s.metrics.IncAuthFailures()
		return ErrAuthenticationFailed
	}

	s.metrics.IncFailed()

	switch {
	case isTimeout(err):
		return ErrTimeout
	case rest.IsServerError(err):
		return ErrSendFailed
	case rest.IsClientError(err):
		return ErrTemporaryFailure
	default:
		return ErrSendFailed
	}
}

func isTimeout(err error) bool {
	if err == nil {
		return false
	}

	var netErr interface{ Timeout() bool }
	if errors.As(err, &netErr) && netErr.Timeout() {
		return true
	}

	return errors.Is(err, context.DeadlineExceeded)
}

func isAuthFailure(err error) bool {
	if err == nil {
		return false
	}

	if !rest.IsClientError(err) {
		return false
	}

	errStr := err.Error()

	return strings.Contains(errStr, "401") || strings.Contains(errStr, "403")
}
