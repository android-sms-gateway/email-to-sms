package bridge

import (
	"context"
	"fmt"
	"strings"

	"github.com/android-sms-gateway/email-to-sms/internal/smsgate"
	"github.com/android-sms-gateway/email-to-sms/internal/smtp"
	"go.uber.org/zap"
)

// Service implements smtp.Handler and processes emails for SMS delivery.
type Service struct {
	sender  *smsgate.Service
	metrics *Metrics
	logger  *zap.Logger
}

// NewService creates and initializes a new Service instance.
func NewService(sender *smsgate.Service, metrics *Metrics, logger *zap.Logger) *Service {
	return &Service{
		sender:  sender,
		metrics: metrics,
		logger:  logger,
	}
}

// Handle processes a validated email and sends the SMS via SMSGate.
// The SMTP server has already validated the domain match.
func (s *Service) Handle(ctx context.Context, email smtp.EmailData) error {
	s.metrics.IncEmailsReceived()

	if err := validatePhone(email.Phone); err != nil {
		return err
	}

	body := strings.TrimSpace(email.Body)
	if body == "" {
		return ErrEmptyBody
	}

	if err := s.sender.Send(ctx, email.Phone, body, email.Username, email.Password); err != nil {
		s.metrics.IncSMSFailed()
		s.logger.Error("SMS send failed",
			zap.String("phone", maskPhone(email.Phone)),
			zap.Error(err),
		)

		return fmt.Errorf("failed to send SMS: %w", err)
	}

	s.metrics.IncSMSSent()
	s.logger.Info("SMS sent successfully",
		zap.String("phone", maskPhone(email.Phone)),
	)

	return nil
}

// validatePhone validates that the phone number contains only digits.
func validatePhone(phone string) error {
	phone = strings.TrimSpace(phone)
	if len(phone) == 0 || len(phone) > 128 {
		return ErrInvalidPhoneNumber
	}

	for _, c := range phone {
		if c < '0' || c > '9' {
			return ErrInvalidPhoneNumber
		}
	}

	return nil
}

func maskPhone(phone string) string {
	const visibleDigits = 4

	if len(phone) <= visibleDigits {
		return "****"
	}
	return strings.Repeat("*", len(phone)-visibleDigits) + phone[len(phone)-visibleDigits:]
}
