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

	// database connection
	conn, err := pgxpool.New(context.Background(), cfg.Database.ConnectionString())
	if err != nil {
		log.Panicf("unable to connect to database: %v", err)
	}
	defer conn.Close()

	// HTTP router
	router := mux.NewRouter()

	// HTTP server
	server, deferFn := getHttpServer(cfg, router, logger)
	logger.Info("listening", slog.String("port", cfg.HTTPPort))

	defer deferFn()
	log.Fatal(server.ListenAndServe())
}

func getHttpServer(cfg config.Config, router *mux.Router, logger *slog.Logger) (*http.Server, func()) {
	conn := connectToDatabase(&cfg)
	logger.Info("connected to database successfully")

	setupRoutes(router, conn)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.HTTPPort),
		Handler: router,
	}

	deferFn := func() {
		fmt.Println("closing database connection")
		conn.Close()
	}
	return server, deferFn
}

func connectToDatabase(cfg *config.Config) *pgxpool.Pool {
	conn, err := pgxpool.New(context.Background(), cfg.Database.ConnectionString())
	if err != nil {
		log.Panicf("unable to connect to database: %v", err)
	}

	return conn
}
