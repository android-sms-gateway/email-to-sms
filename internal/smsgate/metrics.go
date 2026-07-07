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
	apiDuration  *prometheus.HistogramVec
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
		apiDuration: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "api_duration_seconds",
			Help:      "Duration of SMSGate API calls",
			Buckets:   []float64{0.05, 0.1, 0.25, 0.5, 1, 2, 5, 10},
		}, []string{"error_type"}),
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

// ObserveAPIDuration records the SMSGate API call latency with the error type label.
func (m *Metrics) ObserveAPIDuration(seconds float64, errorType string) {
	m.apiDuration.WithLabelValues(errorType).Observe(seconds)
}
