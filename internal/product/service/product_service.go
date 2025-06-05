package service

import (
	"context"
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/config"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/handler/request"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/product/model"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/product/repository"
	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProducts(ctx context.Context, request *request.GetProductsRequest) ([]model.ProductModel, error) {
	config.Logger(ctx, config.LogInfo, "Starting GetProducts")
	products, err := s.repo.GetProducts(request.Page, request.Limit)
	if err != nil {
		config.Logger(ctx, config.LogError, "GetProducts failed: %v", err)
		return nil, err
	}
	return products, nil
}

func (s *ProductService) GetProductByID(id entity.ID) (*model.ProductModel, error) {
	return s.repo.GetProductByID(id)
}

func (s *ProductService) Create(ctx context.Context, product *request.CreateProductRequest) (*model.ProductModel, error) {
	productModel := &model.ProductModel{
		ID:        entity.NewID(),
		Name:      product.Name,
		Price:     product.Price,
		CreatedAt: time.Now(),
	}
	return s.repo.Create(productModel)
}

func (s *ProductService) Update(id entity.ID, product *model.ProductModel) error {
	return s.repo.Update(id, product)
}

func (s *ProductService) Delete(id entity.ID) error {
	return s.repo.Delete(id)
}
