package health

import (
	"fmt"
	"github.com/flink/flink-backend-assingment/internal/data"
	"log/slog"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CheckHealth(healthRepository data.HealthRepository, logger *slog.Logger, w http.ResponseWriter, r *http.Request) {
	healthService := NewService()
	err := healthService.Check(healthRepository, logger)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(fmt.Sprintf("ok response!")))
}
