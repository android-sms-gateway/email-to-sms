package smtp

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	metricsNamespace = "email2sms"
	metricsSubsystem = "smtp"
)

type Metrics struct {
	activeConnections prometheus.Gauge
	connectionsTotal  prometheus.Counter
	sessionDuration   prometheus.Histogram
}

func NewMetrics() *Metrics {
	return &Metrics{
		activeConnections: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "active_connections",
			Help:      "Current number of active SMTP connections",
		}),
		connectionsTotal: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "connections_total",
			Help:      "Total number of SMTP connections accepted",
		}),
		sessionDuration: promauto.NewHistogram(prometheus.HistogramOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "session_duration_seconds",
			Help:      "Duration of SMTP sessions",
			Buckets:   []float64{0.1, 0.5, 1, 2, 5, 10, 30},
		}),
	}
}

func (m *Metrics) IncActiveConnections() {
	m.activeConnections.Inc()
}

func (m *Metrics) DecActiveConnections() {
	m.activeConnections.Dec()
}

func (m *Metrics) IncConnectionsTotal() {
	m.connectionsTotal.Inc()
}

func (m *Metrics) ObserveSessionDuration(seconds float64) {
	m.sessionDuration.Observe(seconds)
}
