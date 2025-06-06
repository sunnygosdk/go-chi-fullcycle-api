package product

// ProductRepository is the interface for product repository.
type ProductRepository interface {
	// Create creates a new product.
	Create(product *product) error
}
