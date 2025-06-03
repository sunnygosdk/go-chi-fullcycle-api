package app

import (
	"database/sql"

	"github.com/go-chi/chi/v5"

	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/handler/controller"
)

func SetupRoutes(r chi.Router, db *sql.DB) {
	SetUtilsRoutes(r, db)
	SetProductRoutes(r, db)
}

func SetUtilsRoutes(r chi.Router, db *sql.DB) {
	utilsController := controller.InjectUtilsController(db)
	r.Get("/health", utilsController.HealthCheck)
}

func SetProductRoutes(r chi.Router, db *sql.DB) {
	productController := controller.InjectProductController(db)
	r.Get("/products", productController.GetProducts)
}
