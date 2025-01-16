package data

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HealthRepository interface {
	Select() error
}

type HealthModel struct {
	DB *pgxpool.Pool
}

func (healthModel HealthModel) Select() error {
	query := "select 1"

	_, err := healthModel.DB.Exec(context.Background(), query)
	return err
}
