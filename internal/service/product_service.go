package service

import (
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/model"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProducts(page int, limit int) ([]model.ProductModel, error) {
	return s.repo.GetProducts(page, limit)
}
