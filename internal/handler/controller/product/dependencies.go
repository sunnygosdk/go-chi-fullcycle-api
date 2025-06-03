package product

import (
	"database/sql"

	productRepository "github.com/sunnygosdk/go-chi-fullcycle-api/internal/repository/product"
	productService "github.com/sunnygosdk/go-chi-fullcycle-api/internal/service/product"
)

func InjectController(db *sql.DB) *Controller {
	repo := productRepository.NewRepository(db)
	service := productService.NewService(repo)
	controller := NewController(service)
	return controller
}
