package app

import (
	"database/sql"

	"github.com/go-chi/chi/v5"

	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/handler/controller/product"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/handler/controller/utils"
)

func SetupRoutes(r chi.Router, db *sql.DB) {
	SetUtilsRoutes(r, db)
	SetProductRoutes(r, db)
}

func SetUtilsRoutes(r chi.Router, db *sql.DB) {
	utilsController := utils.InjectController(db)
	r.Get("/health", utilsController.HealthCheck)
}

func SetProductRoutes(r chi.Router, db *sql.DB) {
	productController := product.InjectController(db)
	r.Get("/products", productController.GetProducts)
}
