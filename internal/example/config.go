package example

// Config holds the configuration for the example module.
//
// This struct contains all the configuration parameters needed to initialize
// and configure the example module. Typically, these values would be loaded
// from environment variables, configuration files, or other configuration sources.
//
// Example:
//
//	cfg := example.Config{
//	    Example: "demo-value",
//	}
//
//	service := example.New(cfg, repo, metrics, logger)
type Config struct {
	// Example is a configuration parameter that demonstrates how to include
	// custom configuration values in the module. This could be used for
	// feature flags, API endpoints, timeouts, or any other configurable
	// aspect of the module.
	Example string
}
