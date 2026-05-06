package example

import "errors"

// This file defines module-specific error types and values.
//
// Having dedicated error types for a module provides several benefits:
//   - Enables error handling specific to this module's domain
//   - Allows for programmatic error type checking
//   - Improves error message consistency
//   - Makes debugging and troubleshooting easier
//
// Example:
//
//	if err := processExample(); err != nil {
//	    if errors.Is(err, example.ErrExample) {
//	        // Handle specific example error
//	    }
//	    // Handle other errors
//	}
var (
	// ErrExample is a predefined error for the example module.
	//
	// This demonstrates how to create module-specific errors that can be
	// used throughout the codebase. In a real application, you might have
	// multiple error types for different error conditions.
	//
	// Usage:
	//   return fmt.Errorf("failed to process: %w", example.ErrExample)
	//
	// Checking:
	//   if errors.Is(err, example.ErrExample) {
	//       // Handle example error
	//   }
	ErrExample = errors.New("example error")
)
