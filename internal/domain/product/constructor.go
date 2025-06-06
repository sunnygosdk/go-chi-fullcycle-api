package product

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// NewProduct creates a new product instance with the provided ID, name, price, and department ID.
// It initializes the product with the given ID, name, price, and department ID,
// and sets the creation and update timestamps to the current time.
//
// The function also validates the product before returning it. If validation fails,
// it returns an error.
//
// Parameters:
//   - Name: Name of the product.
//   - Price: Price of the product.
//   - DepartmentID: Department ID of the product.
//
// Returns:
//   - *product: A pointer to the newly created and validated product.
//   - error: An error if the product validation fails.
func NewProduct(Name string, Price float64, DepartmentID entity.ID) (*product, error) {
	product := &product{
		ID:           entity.NewID(),
		Name:         Name,
		Price:        Price,
		DepartmentID: DepartmentID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	product, err := product.validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

// validate validates the product instance.
// It checks if the product ID is required and valid,
// and if the product name, price, and department ID are required.
//
// Parameters:
//   - p: The product instance to validate.
//
// Returns:
//   - *product: A pointer to the validated product instance.
//   - error: An error if the product validation fails.
func (p *product) validate() (*product, error) {
	if p.ID.String() == "" {
		return nil, ErrProductIDisRequired
	}

	if p.DepartmentID.String() == "" {
		return nil, ErrProductDepartmentIDisRequired
	}

	_, err := entity.ParseID(p.ID.String())
	if err != nil {
		return nil, ErrProductInvalidID
	}

	_, err = entity.ParseID(p.DepartmentID.String())
	if err != nil {
		return nil, ErrProductInvalidDepartmentID
	}

	if p.Name == "" {
		return nil, ErrProductNameRequired
	}

	if p.Price <= 0 {
		return nil, ErrProductPriceLessOrZero
	}

	return p, nil
}
