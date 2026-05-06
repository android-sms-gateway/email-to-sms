package example

import "go.uber.org/zap"

// Service implements the business logic for the example module.
//
// This struct represents the service layer in a clean architecture pattern.
// The service is responsible for implementing business rules, orchestrating
// operations between different components, and exposing functionality to
// the rest of the application.
//
// The Service depends on:
//   - Config: For configuration parameters
//   - Repository: For data access operations
//   - Metrics: For collecting and reporting metrics
//   - Logger: For structured logging
//
// In a real application, this struct would contain methods for:
//   - Business operations and workflows
//   - Coordinating between repositories and other services
//   - Enforcing business rules and validation
//   - Handling errors and logging
type Service struct {
	// config holds the configuration for the service
	config Config

	// examples provides access to data operations
	examples *Repository

	// metrics is used for collecting and reporting metrics
	metrics *Metrics

	// logger is used for structured logging
	logger *zap.Logger
}

// New creates and initializes a new Service instance.
//
// This function serves as a constructor for the Service struct, accepting
// all its dependencies as parameters. This approach, known as dependency
// injection, makes the code more testable and maintainable.
//
// Parameters:
//   - config: Configuration for the service
//   - examples: Repository for data access
//   - metrics: Metrics collector for monitoring
//   - logger: Logger for structured logging
//
// Returns:
//   - *Service: A pointer to the newly created Service instance
//
// Example:
//
//	config := example.Config{Example: "demo"}
//	repo := example.NewRepository()
//	metrics := example.NewMetrics()
//	logger, _ := zap.NewProduction()
//
//	service := example.New(config, repo, metrics, logger)
func New(config Config, examples *Repository, metrics *Metrics, logger *zap.Logger) *Service {
	return &Service{
		config:   config,
		examples: examples,
		metrics:  metrics,
		logger:   logger,
	}
}

// Example returns the example value from the configuration.
func (s *Service) Example() string {
	return s.config.Example
}
