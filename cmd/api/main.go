package main

import (
	"context"
	"fmt"
	"github.com/flink/flink-backend-assingment/internal/config"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// environment variables
	cfg, err := config.Load()
	if err != nil {
		log.Panicf("failed to load config: %v", err)
	}

	// HTTP router
	router := mux.NewRouter()

	// database connection
	conn, err := connectToDatabase(&cfg)
	if err != nil {
		logger.Error("unable to connect to database: %v", err)
	} else {
		logger.Info("connected to database successfully.")
	}
	defer func() {
		fmt.Println("closing database connection")
		conn.Close()
	}()

	setupRoutes(router, conn, logger)

	// HTTP server
	server := getHttpServer(cfg, router)
	logger.Info("listening", slog.String("port", cfg.HTTPPort))

	log.Fatal(server.ListenAndServe())
}

func getHttpServer(cfg config.Config, router *mux.Router) *http.Server {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.HTTPPort),
		Handler: router,
	}

	return server
}

func connectToDatabase(cfg *config.Config) (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(context.Background(), cfg.Database.ConnectionString())
	if err != nil {
		log.Panicf("unable to connect to database: %v", err)
	}

	return conn, err
}
