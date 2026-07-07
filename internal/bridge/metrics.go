package bridge

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	metricsNamespace = "email2sms"
	metricsSubsystem = "bridge"
)

// Metrics handles Prometheus metrics for bridge operations.
type Metrics struct {
	emailsReceived          prometheus.Counter
	smsSent                 prometheus.Counter
	smsFailed               prometheus.Counter
	emailProcessingDuration prometheus.Histogram
	inflightRequests        prometheus.Gauge
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
		emailProcessingDuration: promauto.NewHistogram(prometheus.HistogramOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "email_processing_duration_seconds",
			Help:      "Duration of email processing from receipt to SMS send result",
			Buckets:   []float64{0.1, 0.5, 1, 2, 5, 10, 30},
		}),
		inflightRequests: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "inflight_requests",
			Help:      "Current number of emails being processed",
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

func (m *Metrics) ObserveEmailProcessingDuration() func() {
	start := time.Now()

	return func() {
		m.emailProcessingDuration.Observe(time.Since(start).Seconds())
	}
}

func (m *Metrics) InflightRequest() func() {
	m.inflightRequests.Inc()

	return func() {
		m.inflightRequests.Dec()
	}
}
