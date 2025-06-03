package product

import (
	productModel "github.com/sunnygosdk/go-chi-fullcycle-api/internal/model/product"
	productRepository "github.com/sunnygosdk/go-chi-fullcycle-api/internal/repository/product"
)

type Service struct {
	repo *productRepository.Repository
}

func NewService(repo *productRepository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetProducts(page int, limit int) ([]productModel.Model, error) {
	return s.repo.GetProducts(page, limit)
}
