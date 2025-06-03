package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var (
	databaseErrorMessage = Response{
		Status:  "error",
		Message: "Database connection failed",
	}

	healthCheckMessage = Response{
		Status:  "success",
		Message: "Server Healthly",
	}
)

func WriteUtilsResponse(ctx context.Context, w http.ResponseWriter, status int, response Response) {
	requestID := middleware.GetReqID(ctx)
	if requestID != "" {
		w.Header().Set("X-Request-ID", requestID)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
