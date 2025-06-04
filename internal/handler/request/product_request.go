package request

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type GetProductsRequest struct {
	Page  int
	Limit int
}

type CreateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func ParseGetProductsRequest(r *http.Request) *GetProductsRequest {
	page := chi.URLParam(r, "page")
	limit := chi.URLParam(r, "limit")
	return ValidateGetProductsRequest(page, limit)
}

func ValidateGetProductsRequest(page string, limit string) *GetProductsRequest {
	const (
		defaultPage  = 1
		defaultLimit = 10
		minPage      = 1
		minLimit     = 1
		maxLimit     = 100
	)

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < minPage {
		pageInt = defaultPage
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < minLimit || limitInt > maxLimit {
		limitInt = defaultLimit
	}

	return &GetProductsRequest{
		Page:  pageInt,
		Limit: limitInt,
	}
}

func ParseCreateProductRequest(r *http.Request) (*CreateProductRequest, error) {
	var req CreateProductRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	validator, err := ValidateCreateProductRequest(req.Name, req.Price)
	if err != nil {
		return nil, err
	}
	return validator, nil
}

func ValidateCreateProductRequest(name string, price float64) (*CreateProductRequest, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	if price <= 0 {
		return nil, errors.New("price must be greater than 0")
	}
	return &CreateProductRequest{
		Name:  name,
		Price: price,
	}, nil
}
