// Package example provides a demonstration of a well-structured Go module using
// the Uber FX framework for dependency injection. This module showcases common
// patterns and best practices for organizing Go code in a modular, maintainable way.
//
// The example module is organized into several files, each with a specific responsibility:
//
//   - config.go: Contains configuration structures for the module
//   - domain.go: Defines domain entities and core business logic types
//   - errors.go: Custom error definitions for the module
//   - metrics.go: Prometheus metrics collection and reporting
//   - models.go: Data models used by the module
//   - module.go: FX module definition for dependency injection
//   - repository.go: Data access layer implementation
//   - service.go: Business logic and service layer implementation
//
// This structure follows a clean architecture approach with clear separation of concerns.
package example

// Overview
//
// The example module demonstrates a typical structure for a Go service module:
//
// 1. Configuration (config.go)
//    - Defines the Config struct that holds module-specific configuration
//    - Typically loaded from environment variables or configuration files
//
// 2. Domain Layer (domain.go)
//    - Contains core business entities and types
//    - Represents the business domain concepts
//
// 3. Error Handling (errors.go)
//    - Defines module-specific error types
//    - Provides consistent error handling throughout the module
//
// 4. Metrics (metrics.go)
//    - Implements Prometheus metrics for monitoring
//    - Provides observability into module operations
//
// 5. Data Models (models.go)
//    - Defines data structures used internally
//    - May represent database entities or DTOs
//
// 6. Module Definition (module.go)
//    - Uses Uber FX for dependency injection
//    - Wires up all components and their dependencies
//
// 7. Repository Layer (repository.go)
//    - Handles data access and persistence
//    - Abstracts the data source from the service layer
//
// 8. Service Layer (service.go)
//    - Implements business logic
//    - Orchestrates operations between domain, repository, and other components
//
// Usage
//
// To use this module in your application, import it and include the FX module:
//
//   app := fx.New(
//       example.Module(),
//       // other modules...
//   )
//
// The module will automatically wire up all dependencies and make the Service available
// for use by other modules.
//
// Dependencies
//
// This module depends on:
//   - go.uber.org/fx: For dependency injection
//   - go.uber.org/zap: For structured logging
//   - github.com/prometheus/client_golang: For metrics collection
//   - github.com/go-core-fx/logger: For enhanced logging capabilities
//
// Example
//
// Here's a basic example of how to use the example module:
//
//   package main
//
//   import (
//       "context"
//       "go.uber.org/fx"
//
//       "yourproject/internal/example"
//   )
//
//   func main() {
//       app := fx.New(
//           example.Module(),
//           fx.Invoke(run),
//       )
//
//       app.Run()
//   }
//
//   func run(lc fx.Lifecycle, service *example.Service) {
//       lc.Append(fx.Hook{
//           OnStart: func(ctx context.Context) error {
//               // Use the service here
//               return nil
//           },
//       })
//   }
