package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sunnygosdk/go-chi-fullcycle-api/config"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/service"
)

type ProductController struct {
	service *service.ProductService
}

func NewProductController(service *service.ProductService) *ProductController {
	return &ProductController{service: service}
}

func (c *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	config.Logger(ctx, config.LogInfo, "Starting GetProducts")
	page := chi.URLParam(r, "page")
	limit := chi.URLParam(r, "limit")
	request := ValidateGetProductsRequest(page, limit)

	products, err := c.service.GetProducts(request.Page, request.Limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		config.Logger(ctx, config.LogError, "GetProducts failed: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
