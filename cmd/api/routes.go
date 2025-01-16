package main

import (
	"github.com/flink/flink-backend-assingment/internal/data"
	"github.com/flink/flink-backend-assingment/internal/health"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"net/http"
)

func setupRoutes(router *mux.Router, conn *pgxpool.Pool, logger *slog.Logger) {
	healthHandler := health.NewHandler()
	var healthRepository data.HealthRepository
	healthRepository = data.HealthModel{DB: conn}

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		healthHandler.CheckHealth(healthRepository, logger, w, r)
	}).Methods(http.MethodGet)
}
