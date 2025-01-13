package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

func setupRoutes(router *mux.Router, conn *pgxpool.Pool) {
	// health endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := conn.Exec(context.Background(), "select 1")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprintf("ok reponse!")))
	}).Methods(http.MethodGet)
}
