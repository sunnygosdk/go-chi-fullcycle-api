package entity

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// Stock represents a stock within an store.
type Stock struct {
	ID        entity.ID
	Quantity  int
	ProductID entity.ID
	StoreID   entity.ID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// NewStock creates a new stock.
// It validates the stock before creating it. If validation fails,
// it returns an error.
//
// Parameters:
//   - quantity: Quantity of the stock.
//   - productID: ID of the product the stock belongs to.
//   - storeID: ID of the store the stock belongs to.
//
// Returns:
//   - *Stock: A pointer to the newly created and validated stock.
//   - error: An error if the stock validation fails.
func NewStock(quantity int, productID string, storeID string) (*Stock, error) {
	err := validateQuantity(quantity)
	if err != nil {
		return nil, err
	}

	prodID, err := entity.ParseID(productID)
	if err != nil {
		return nil, ErrorStockInvalidProductID
	}

	stoID, err := entity.ParseID(storeID)
	if err != nil {
		return nil, ErrorStockInvalidStoreID
	}

	stock := &Stock{
		ID:        entity.NewID(),
		Quantity:  quantity,
		ProductID: prodID,
		StoreID:   stoID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}

	return stock, nil
}

// Update updates the stock with the provided values.
// It validates the stock before updating it. If validation fails,
// it returns an error.
//
// Parameters:
//   - quantity: Quantity of the stock.
//   - productID: ID of the product the stock belongs to.
//   - storeID: ID of the store the stock belongs to.
//
// Returns:
//   - error: An error if the stock validation fails.
func (s *Stock) Update(quantity *int, productID *string, storeID *string) error {
	if quantity == nil && productID == nil && storeID == nil {
		return ErrorStockAtLeastOneField
	}

	if quantity != nil {
		err := validateQuantity(*quantity)
		if err != nil {
			return err
		}
		s.Quantity = *quantity
	}

	if productID != nil {
		prodID, err := entity.ParseID(*productID)
		if err != nil {
			return ErrorStockInvalidProductID
		}
		s.ProductID = prodID
	}

	if storeID != nil {
		stoID, err := entity.ParseID(*storeID)
		if err != nil {
			return ErrorStockInvalidStoreID
		}
		s.StoreID = stoID
	}

	s.UpdatedAt = time.Now()
	return nil
}

// Delete marks the stock as deleted by setting the deletedAt timestamp to the current time.
// It also validates the stock before deleting it. If validation fails,
// it returns an error.
func (s *Stock) Delete() error {
	if s.DeletedAt != nil {
		return ErrorStockIsDeleted
	}

	deletedAt := time.Now()
	s.DeletedAt = &deletedAt
	return nil
}

// validateQuantity validates the quantity of the stock.
// It returns an error if the quantity is less than zero.
//
// Parameters:
//   - quantity: Quantity of the stock.
//
// Returns:
//   - error: An error if the stock quantity validation fails.
func validateQuantity(quantity int) error {
	if quantity < 0 {
		return ErrorStockQuantityLessOfZero
	}
	return nil
}
