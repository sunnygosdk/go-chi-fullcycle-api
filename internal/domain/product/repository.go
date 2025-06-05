package domain

type ProductRepository interface {
	Create(product *product) error
}
