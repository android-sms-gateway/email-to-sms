package smsgate

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	metricsNamespace = "email2sms"
	metricsSubsystem = "smsgate"
)

// Metrics handles Prometheus metrics for SMSGate operations.
type Metrics struct {
	sentTotal    prometheus.Counter
	failedTotal  prometheus.Counter
	authFailures prometheus.Counter
}

// NewMetrics creates and initializes metrics collectors.
func NewMetrics() *Metrics {
	return &Metrics{
		sentTotal: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "sms_sent_total",
			Help:      "Total number of SMS sent successfully",
		}),
		failedTotal: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "sms_failed_total",
			Help:      "Total number of SMS send failures",
		}),
		authFailures: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "auth_failed_total",
			Help:      "Total number of SMSGate authentication failures",
		}),
	}
}

// IncSent increments the successful sends counter.
func (m *Metrics) IncSent() {
	m.sentTotal.Inc()
}

// IncFailed increments the failures counter.
func (m *Metrics) IncFailed() {
	m.failedTotal.Inc()
}

// IncAuthFailures increments authentication failures counter.
func (m *Metrics) IncAuthFailures() {
	m.authFailures.Inc()
}
