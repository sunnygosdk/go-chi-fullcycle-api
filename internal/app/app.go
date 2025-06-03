package app

import (
	"fmt"
	"log"
	"net/http"

	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func StartServer(port string, db *sql.DB) {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	SetupRoutes(r, db)

	log.Println("Server Running on Port", port)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), r)

}
