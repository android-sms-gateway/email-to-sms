package example

import (
	"github.com/go-core-fx/logger"
	"go.uber.org/fx"
)

// Module creates and returns an FX module for the example package.
//
// This function defines how the example module should be wired into an application
// using the Uber FX dependency injection framework. It specifies all the components
// that make up the module and how they depend on each other.
//
// The module includes:
//   - A named logger for structured logging
//   - A repository for data access (provided privately)
//   - A service for business logic (provided publicly)
//
// FX will automatically resolve dependencies and inject them where needed.
// For example, the Service depends on Config, Repository, Metrics, and Logger,
// so FX will ensure these are available when creating the Service.
//
// Usage:
//
//	app := fx.New(
//	    example.Module(),
//	    // other modules...
//	)
//
//	// The Service can then be injected into other components:
//	fx.Invoke(func(service *example.Service) {
//	    // Use the service
//	})
func Module() fx.Option {
	return fx.Module(
		"example",
		// Add a named logger for this module
		logger.WithNamedLogger("example"),

		// Provide the metrics collector as a private dependency
		// This means it can only be used within this module
		fx.Provide(NewMetrics, fx.Private),

		// Provide the repository as a private dependency
		// This means it can only be used within this module
		fx.Provide(NewRepository, fx.Private),

		// Provide the service as a public dependency
		// This means it can be injected into other modules
		fx.Provide(New),
	)
}
