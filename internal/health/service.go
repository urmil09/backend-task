package health

import (
	"github.com/flink/flink-backend-assingment/internal/data"
	"log/slog"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Check(healthRepository data.HealthRepository, logger *slog.Logger) error {
	err := healthRepository.Select()
	if err != nil {
		logger.Error("Error in fetching data: ", err)
	}
	return err
}
