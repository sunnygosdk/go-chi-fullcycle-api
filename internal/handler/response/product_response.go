package response

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/sunnygosdk/go-chi-fullcycle-api/config"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/model"
)

type GetProductsResponse struct {
	RequestID string               `json:"request_id"`
	Message   string               `json:"message"`
	Page      int                  `json:"page"`
	Limit     int                  `json:"limit"`
	Total     int                  `json:"total"`
	Products  []model.ProductModel `json:"products"`
}

type CreateProductResponse struct {
	RequestID string             `json:"request_id"`
	Message   string             `json:"message"`
	Product   model.ProductModel `json:"product"`
}

func ErrorGetProductsResponse(
	ctx context.Context,
	w http.ResponseWriter,
	statusCode int,
	message string,
) {
	config.Logger(ctx, config.LogError, "GetProducts failed: %v", message)
	w.WriteHeader(statusCode)
	res := &GetProductsResponse{
		RequestID: middleware.GetReqID(ctx),
		Message:   message,
		Page:      0,
		Limit:     0,
		Total:     0,
		Products:  []model.ProductModel{},
	}
	json.NewEncoder(w).Encode(res)
}

func SuccessGetProductsResponse(
	ctx context.Context,
	w http.ResponseWriter,
	statusCode int,
	message string,
	products []model.ProductModel,
) {
	config.Logger(ctx, config.LogInfo, "GetProducts completed successfully: %v", message)
	w.WriteHeader(statusCode)
	res := &GetProductsResponse{
		RequestID: middleware.GetReqID(ctx),
		Message:   message,
		Page:      0,
		Limit:     0,
		Total:     0,
		Products:  products,
	}
	json.NewEncoder(w).Encode(res)
}

func ErrorCreateProductResponse(
	ctx context.Context,
	w http.ResponseWriter,
	statusCode int,
	message string,
) {
	config.Logger(ctx, config.LogError, "CreateProduct failed: %v", message)
	w.WriteHeader(statusCode)
	res := &CreateProductResponse{
		RequestID: middleware.GetReqID(ctx),
		Message:   message,
		Product:   model.ProductModel{},
	}
	json.NewEncoder(w).Encode(res)
}

func SuccessCreateProductResponse(
	ctx context.Context,
	w http.ResponseWriter,
	statusCode int,
	message string,
	product *model.ProductModel,
) {
	config.Logger(ctx, config.LogInfo, "CreateProduct completed successfully: %v", message)
	w.WriteHeader(statusCode)
	res := &CreateProductResponse{
		RequestID: middleware.GetReqID(ctx),
		Message:   message,
		Product:   *product,
	}
	json.NewEncoder(w).Encode(res)
}
