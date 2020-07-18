package shortener

// Repository is an interface for a repository for the url package.
// It's made to have methods that expose Creating, Reading, Updating and deleting a url shortener.
type Repository interface {
	// Create inserts a model in the repository.
	Create(urlsh Model) error
	// Remove removes a model from the repository
	Remove(id string) error
	// Get fetches a model from the repository
	Get(id string) (Model, error)
	// GetAll fetches all models from the repository
	GetAll() ([]Model, error)
	// Update updates a model in the repository
	Update(id string, m Model) error
}
