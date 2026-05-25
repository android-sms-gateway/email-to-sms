package bridge

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	metricsNamespace = "email2sms"
	metricsSubsystem = "bridge"
)

// Metrics handles Prometheus metrics for bridge operations.
type Metrics struct {
	emailsReceived prometheus.Counter
	smsSent        prometheus.Counter
	smsFailed      prometheus.Counter
}

// NewMetrics creates and initializes metrics collectors.
func NewMetrics() *Metrics {
	return &Metrics{
		emailsReceived: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "emails_received_total",
			Help:      "Total number of emails received by the bridge",
		}),
		smsSent: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "sms_sent_total",
			Help:      "Total number of SMS sent successfully",
		}),
		smsFailed: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "sms_failed_total",
			Help:      "Total number of SMS send failures",
		}),
	}
}

// IncEmailsReceived increments the emails received counter.
func (m *Metrics) IncEmailsReceived() {
	m.emailsReceived.Inc()
}

// IncSMSSent increments the successful sends counter.
func (m *Metrics) IncSMSSent() {
	m.smsSent.Inc()
}

// IncSMSFailed increments the failures counter.
func (m *Metrics) IncSMSFailed() {
	m.smsFailed.Inc()
}
