package utils

import (
	"database/sql"
	"net/http"

	"github.com/sunnygosdk/go-chi-fullcycle-api/configs"
)

type Controller struct {
	db *sql.DB
}

func NewController(db *sql.DB) *Controller {
	return &Controller{db: db}
}

func (c *Controller) HealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	configs.Logger(ctx, configs.LogInfo, "Starting HealthCheck")
	err := c.db.Ping()
	if err != nil {
		WriteResponse(ctx, w, http.StatusInternalServerError, databaseErrorMessage)
		configs.Logger(ctx, configs.LogError, "HealthCheck failed: %v", err)
		return
	}

	configs.Logger(ctx, configs.LogInfo, "HealthCheck completed successfully")
	WriteResponse(ctx, w, http.StatusOK, healthCheckMessage)
}
