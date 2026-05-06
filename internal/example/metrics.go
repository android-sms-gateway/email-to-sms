package example

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics handles Prometheus metrics collection for the example module.
//
// This struct encapsulates all Prometheus metrics related to the example module.
// Having a dedicated metrics struct provides a clean way to organize and manage
// metrics, making it easier to add new metrics and maintain existing ones.
//
// Metrics are important for:
//   - Monitoring application health and performance
//   - Tracking business metrics and KPIs
//   - Setting up alerts and dashboards
//   - Debugging and troubleshooting issues
//
// Example:
//
//	metrics := example.NewMetrics()
//	metrics.IncTotal() // Increment the counter
type Metrics struct {
	// totalCounter is a Prometheus counter that tracks the total number
	// of examples processed or created by this module.
	//
	// Counters are used for values that only increase, such as request counts,
	// error counts, or completed operations.
	totalCounter prometheus.Counter
}

// NewMetrics creates and initializes a new Metrics instance.
//
// This function serves as a constructor for the Metrics struct, initializing
// all the Prometheus metrics with their appropriate configuration.
//
// The metrics defined here are:
//   - example_total: A counter that tracks the total number of examples
//
// Returns:
//   - *Metrics: A pointer to the newly created Metrics instance
//
// Example:
//
//	metrics := example.NewMetrics()
//	// Use metrics in your service
//	service := example.New(config, repo, metrics, logger)
func NewMetrics() *Metrics {
	return &Metrics{
		totalCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "example_total",
			Help: "Total number of examples",
		}),
	}
}

// IncTotal increments the total example counter.
//
// This method should be called whenever an example is processed, created,
// or any other event occurs that should be tracked by the total counter.
//
// Example:
//
//	func (s *Service) ProcessExample() error {
//	    // Process the example
//	    s.metrics.IncTotal() // Record the metric
//	    return nil
//	}
func (m *Metrics) IncTotal() {
	m.totalCounter.Inc()
}
