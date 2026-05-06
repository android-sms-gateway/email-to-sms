package example

// exampleModel represents a data model used by the example module.
//
// This struct demonstrates how to define data models that might be used for
// database entities, DTOs (Data Transfer Objects), or other data structures
// that need to be serialized/deserialized or mapped to external systems.
//
// In a typical application:
//   - Models represent the shape of your data
//   - They often include tags for serialization (JSON, database, etc.)
//   - They might include validation rules
//   - They separate the internal representation from external APIs
//
// Example with tags:
//
//	type exampleModel struct {
//	    ID        int    `json:"id" db:"id"`
//	    Name      string `json:"name" db:"name"`
//	    CreatedAt time.Time `json:"created_at" db:"created_at"`
//	}
//
// Note: This model is unexported (lowercase 'e') to indicate it's for
// internal use within the module. If it needed to be accessed from outside
// the module, it would be exported (uppercase 'E').
type exampleModel struct {
	// In a real application, this would contain fields that represent
	// the data structure, possibly with tags for serialization,
	// database mapping, validation, etc.
	ID    int
	Value string
}
