package controller

import (
	"database/sql"
	"net/http"

	"github.com/sunnygosdk/go-chi-fullcycle-api/config"
)

type UtilsController struct {
	db *sql.DB
}

func NewUtilsController(db *sql.DB) *UtilsController {
	return &UtilsController{db: db}
}

func (c *UtilsController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	config.Logger(ctx, config.LogInfo, "Starting HealthCheck")
	err := c.db.Ping()
	if err != nil {
		WriteUtilsResponse(ctx, w, http.StatusInternalServerError, databaseErrorMessage)
		config.Logger(ctx, config.LogError, "HealthCheck failed: %v", err)
		return
	}

	config.Logger(ctx, config.LogInfo, "HealthCheck completed successfully")
	WriteUtilsResponse(ctx, w, http.StatusOK, healthCheckMessage)
}
