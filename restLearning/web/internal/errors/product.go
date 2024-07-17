package errors

import (
	"fmt"
)

type ProductError struct {
	Status  string
	Message string
}

func NewProductError(status, message string) error {
	return &ProductError{
		Status:  status,
		Message: message,
	}
}

func (pe *ProductError) Error() string {
	return fmt.Sprintf("status: %s, message: %s", pe.Status, pe.Message)
}

// Common errors
var (
	ErrProductNotFound   = NewProductError("404", "Product not found")
	ErrProductExists     = NewProductError("409", "Product already exists")
	ErrProductNotAdded   = NewProductError("500", "Failed to add product")
	ErrProductNotUpdated = NewProductError("500", "Failed to update product")
	ErrProductNotDeleted = NewProductError("500", "Failed to delete product")
	ErrProductNotPatched = NewProductError("500", "Failed to patch product")
	ErrBadRequest        = NewProductError("400", "Bad request")
	ErrorInternal        = NewProductError("500", "Internal server error")
)
