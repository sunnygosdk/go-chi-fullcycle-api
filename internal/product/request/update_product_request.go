package request

type UpdateProductRequest struct {
	Name  *string  `json:"name"`
	Price *float64 `json:"price"`
}
