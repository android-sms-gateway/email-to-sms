package example

import "sync"

// Repository handles data access operations for the example module.
//
// This struct represents the repository layer in a clean architecture pattern.
// The repository is responsible for all data access operations, abstracting
// the details of data storage and retrieval from the rest of the application.
//
// Benefits of using a repository pattern:
//   - Separates data access logic from business logic
//   - Makes it easier to switch data sources (e.g., from SQL to NoSQL)
//   - Centralizes data access operations
//   - Improves testability by allowing mock repositories
//
// In a real application, this struct would contain methods for:
//   - Creating, reading, updating, and deleting (CRUD) data
//   - Querying data with various filters
//   - Handling transactions
//   - Mapping between domain entities and data models
type Repository struct {
	// In a real application, this would contain fields like:
	// - db *sql.DB (for SQL databases)
	// - client *mongo.Client (for MongoDB)
	// - cache *redis.Client (for caching)
	// - logger *zap.Logger (for logging)
	items []exampleModel
	mu    sync.Mutex
}

// NewRepository creates and initializes a new Repository instance.
//
// This function serves as a constructor for the Repository struct. In a real
// application, it would typically accept dependencies like database connections,
// clients, or configuration needed to initialize the repository.
//
// Returns:
//   - *Repository: A pointer to the newly created Repository instance
//
// Example:
//
//	repo := example.NewRepository()
//	// Use the repository in your service
//	service := example.New(config, repo, metrics, logger)
func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Add(item Example) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items = append(r.items, exampleModel{ID: len(r.items) + 1, Value: item.Value})
}
