package repository

import (
	"context"
	"log/slog"
	"subscriptions/internal/model"
	"subscriptions/internal/repository/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SubRepository struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func NewSubRepository(db *pgxpool.Pool, logger *slog.Logger) *SubRepository {
	return &SubRepository{db: db, logger: logger}
}

func (r *SubRepository) CreateSub(ctx context.Context, sub model.Subscription) (model.Subscription, error) {
	ent := entities.ModelToEntity(sub)

	query := `
		INSERT INTO subscriptions (user_id, price, service_name, start_date)
		VALUES ($1, $2, $3,$4)
		RETURNING start_date
	`

	err := r.db.QueryRow(ctx, query, ent.UserID, ent.Price, ent.ServiceName, ent.StartDate).Scan(&ent.StartDate)
	if err != nil {
		r.logger.Error("Failed to create sub in repo layer", slog.Any("err", err))
		return model.Subscription{}, err
	}

	sub = entities.EntityToModel(ent)

	return sub, nil
}
