package controller

import (
	"database/sql"

	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/repository"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/service"
)

func InjectProductController(db *sql.DB) *ProductController {
	repo := repository.NewProductRepository(db)
	service := service.NewProductService(repo)
	controller := NewProductController(service)
	return controller
}
