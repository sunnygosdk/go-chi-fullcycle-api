package product

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sunnygosdk/go-chi-fullcycle-api/configs"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/service/product"
)

type Controller struct {
	service *product.Service
}

func NewController(service *product.Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	configs.Logger(ctx, configs.LogInfo, "Starting GetProducts")
	page := chi.URLParam(r, "page")
	limit := chi.URLParam(r, "limit")
	request := ValidateGetProductsRequest(page, limit)

	products, err := c.service.GetProducts(request.Page, request.Limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		configs.Logger(ctx, configs.LogError, "GetProducts failed: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
