package controller

import (
	"net/http"

	"github.com/sunnygosdk/go-chi-fullcycle-api/config"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/handler/request"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/handler/response"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/model"
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
	request := request.ParseGetProductsRequest(r)

	products, err := c.service.GetProducts(ctx, request)
	if err != nil {
		response.ErrorGetProductsResponse(ctx, w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(products) == 0 {
		response.SuccessGetProductsResponse(ctx, w, http.StatusOK, "No products found", []model.ProductModel{})
		return
	}

	response.SuccessGetProductsResponse(ctx, w, http.StatusOK, "Products found", products)
}

func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	config.Logger(ctx, config.LogInfo, "Starting CreateProduct")
	request, err := request.ParseCreateProductRequest(r)
	if err != nil {
		response.ErrorCreateProductResponse(ctx, w, http.StatusBadRequest, err.Error())
		return
	}

	product, err := c.service.Create(ctx, request)
	if err != nil {
		response.ErrorCreateProductResponse(ctx, w, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessCreateProductResponse(ctx, w, http.StatusOK, "Product created", product)
}
