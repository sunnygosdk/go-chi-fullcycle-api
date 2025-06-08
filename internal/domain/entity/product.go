package entity

import (
	"errors"
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// Error messages for product validation
var (
	ErrorProductNameRequired         = errors.New("product: name is required")
	ErrorProductDescriptionRequired  = errors.New("product: description is required")
	ErrorProductPriceLessOrZero      = errors.New("product: price must be greater than zero")
	ErrorProductInvalidDepartmentID  = errors.New("product: invalid department ID")
	ErrorProductIsDeleted            = errors.New("product: product is already deleted")
	ErrorProductAtLeastOneField      = errors.New("product: at least one field must be provided")
	ErrorProductDescriptionMinLength = errors.New("product: description must be at least 2 characters long")
)

// Product represents a product within an department.
type Product struct {
	ID           entity.ID
	Name         string
	Description  string
	Price        float64
	DepartmentID entity.ID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

// NewProduct creates a new product.
// It validates the product before creating it. If validation fails,
// it returns an error.
//
// Parameters:
//   - name: Name of the product.
//   - description: Description of the product.
//   - price: Price of the product.
//   - departmentID: ID of the department the product belongs to.
//
// Returns:
//   - *Product: A pointer to the newly created and validated product.
//   - error: An error if the product validation fails.
func NewProduct(name string, description string, price float64, departmentID string) (*Product, error) {
	err := validateProductName(name)
	if err != nil {
		return nil, err
	}

	err = validateProductDescription(description)
	if err != nil {
		return nil, err
	}

	err = validateProductPrice(price)
	if err != nil {
		return nil, err
	}

	depID, err := entity.ParseID(departmentID)
	if err != nil {
		return nil, ErrorProductInvalidDepartmentID
	}

	product := &Product{
		ID:           entity.NewID(),
		Name:         name,
		Description:  description,
		Price:        price,
		DepartmentID: depID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    nil,
	}

	return product, nil
}

// Update updates the product with the provided values.
// It validates the product before updating it. If validation fails,
// it returns an error.
//
// Parameters:
//   - name: Name of the product.
//   - description: Description of the product.
//   - price: Price of the product.
//   - departmentID: ID of the department the product belongs to.
//
// Returns:
//   - error: An error if the product validation fails.
func (p *Product) Update(name *string, description *string, price *float64, departmentID *string) error {
	if name == nil &&
		description == nil &&
		price == nil &&
		departmentID == nil {
		return ErrorProductAtLeastOneField
	}

	if name != nil {
		err := validateProductName(*name)
		if err != nil {
			return err
		}
		p.Name = *name
	}

	if description != nil {
		err := validateProductDescription(*description)
		if err != nil {
			return err
		}
		p.Description = *description
	}

	if price != nil {
		err := validateProductPrice(*price)
		if err != nil {
			return err
		}
		p.Price = *price
	}

	if departmentID != nil {
		depID, err := entity.ParseID(*departmentID)
		if err != nil {
			return ErrorProductInvalidDepartmentID
		}
		p.DepartmentID = depID
	}

	p.UpdatedAt = time.Now()
	return nil
}

// Delete marks the product as deleted by setting the deletedAt timestamp to the current time.
// It also validates the product before deleting it. If validation fails,
// it returns an error.
func (p *Product) Delete() error {
	if p.DeletedAt != nil {
		return ErrorProductIsDeleted
	}

	deletedAt := time.Now()
	p.DeletedAt = &deletedAt
	return nil
}

// validateProductName validates the name of the product.
// It returns an error if the name is empty.
//
// Parameters:
//   - name: Name of the product.
//
// Returns:
//   - error: An error if the product name validation fails.
func validateProductName(name string) error {
	if name == "" {
		return ErrorProductNameRequired
	}
	return nil
}

// validateProductDescription validates the description of the product.
// It returns an error if the description is empty.
//
// Parameters:
//   - description: Description of the product.
//
// Returns:
//   - error: An error if the product description validation fails.
func validateProductDescription(description string) error {
	if description == "" {
		return ErrorProductDescriptionRequired
	}

	if len(description) < 2 {
		return ErrorProductDescriptionMinLength
	}

	return nil
}

// validateProductPrice validates the price of the product.
// It returns an error if the price is less than or equal to zero.
//
// Parameters:
//   - price: Price of the product.
//
// Returns:
//   - error: An error if the product price validation fails.
func validateProductPrice(price float64) error {
	if price <= 0 {
		return ErrorProductPriceLessOrZero
	}
	return nil
}
