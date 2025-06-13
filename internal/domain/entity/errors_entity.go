package entity

import "errors"

// Generic error messages
var (
	ErrorEntityInvalidID = errors.New("invalid ID")
)

// Error messages for department validation
var (
	ErrorDepartmentNameRequired         = errors.New("department: department name is required")
	ErrorDepartmentDescriptionRequired  = errors.New("department: department description is required")
	ErrorDepartmentNameMinLength        = errors.New("department: department name must be at least 2 characters long")
	ErrorDepartmentDescriptionMinLength = errors.New("department: department description must be at least 2 characters long")
	ErrorDepartmentIsDeleted            = errors.New("department: department is already deleted")
	ErrorDepartmentAtLeastOneField      = errors.New("department: at least one field must be provided")
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

// Error messages for stock validation
var (
	ErrorStockInvalidID          = errors.New("stock: invalid ID")
	ErrorStockQuantityLessOfZero = errors.New("stock: quantity must be greater than zero")
	ErrorStockInvalidProductID   = errors.New("stock: invalid product ID")
	ErrorStockInvalidStoreID     = errors.New("stock: invalid store ID")
	ErrorStockAtLeastOneField    = errors.New("stock: at least one field must be provided")
	ErrorStockIsDeleted          = errors.New("stock: stock is already deleted")
)

// Error messages for store_department_map validation
var (
	ErrorSDMInvalidStoreID      = errors.New("store_department_map: invalid store ID")
	ErrorSDMInvalidDepartmentID = errors.New("store_department_map: invalid department ID")
	ErrorSDMIsDeleted           = errors.New("store_department_map: store is already deleted")
	ErrorSDMAtLeastOneField     = errors.New("store_department_map: at least one field must be provided")
)

// Error messages for store validation
var (
	ErrorStoreInvalidID        = errors.New("store: invalid ID")
	ErrorStoreIsDeleted        = errors.New("store: store is already deleted")
	ErrorStoreInvalidName      = errors.New("store: invalid name")
	ErrorStoreInvalidAddress   = errors.New("store: invalid address")
	ErrorStoreMinLengthName    = errors.New("store: name must be at least 3 characters long")
	ErrorStoreMinLengthAddress = errors.New("store: address must be at least 3 characters long")
	ErrorStoreAtLeastOneField  = errors.New("store: at least one field must be provided")
)

// Error messages for transaction validation
var (
	ErrorTransactionQuantityIsZero         = errors.New("transaction: quantity must not be zero")
	ErrorTransactionInvalidStockID         = errors.New("transaction: invalid stock ID")
	ErrorTransactionInvalidTransactionType = errors.New("transaction: invalid transaction type")
	ErrorTransactionAtLeastOneField        = errors.New("transaction: at least one field must be provided")
	ErrorTransactionIsDeleted              = errors.New("transaction: transaction is already deleted")
)
