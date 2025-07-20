package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"subscriptions/internal/model"
	"subscriptions/internal/repository/entities"
	"time"

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

func (r *SubRepository) GetByID(ctx context.Context, id int) (model.Subscription, error) {
	var ent entities.Subscription
	ent.ID = id

	query := `
		SELECT user_id, price, service_name, start_date
		FROM subscriptions
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, ent.ID).Scan(&ent.UserID, &ent.Price, &ent.ServiceName, &ent.StartDate)
	if err != nil {
		r.logger.Error("Failed to find subscription in repo layer", slog.Any("err", err))
		return model.Subscription{}, err
	}

	sub := entities.EntityToModel(ent)

	return sub, nil
}

func (r *SubRepository) UpdateSubPrice(ctx context.Context, sub model.Subscription) (model.Subscription, error) {
	query := `
		UPDATE subscriptions
		SET price = $1, updated_at = CURRENT_TIMESTAMP
		WHERE id = $2
		RETURNING user_id, service_name, start_date, updated_at
	`

	var ent entities.Subscription
	ent.ID = sub.ID
	ent.Price = sub.Price

	err := r.db.QueryRow(ctx, query, ent.Price, ent.ID).Scan(
		&ent.UserID,
		&ent.ServiceName,
		&ent.StartDate,
		&ent.UpdatedAt,
	)
	if err != nil {
		r.logger.Error("Failed to change subscription in repo layer", slog.Any("err", err))
		return model.Subscription{}, err
	}

	sub = entities.EntityToModel(ent)
	return sub, nil
}

func (r *SubRepository) DeleteByID(ctx context.Context, id int) error {
	query := `DELETE FROM subscriptions WHERE id = $1`

	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		r.logger.Error("Repo failed to delete subscription", slog.Int("id", id), slog.Any("err", err))
		return err
	}
	return nil
}

func (r *SubRepository) ListSubs(ctx context.Context, sub model.Subscription) ([]model.Subscription, error) {
	var rows pgx.Rows
	var err error

	if sub.UserID != uuid.Nil {
		query := `
			SELECT id, user_id, price, service_name, start_date, created_at, updated_at
			FROM subscriptions
			WHERE user_id = $1
		`
		rows, err = r.db.Query(ctx, query, sub.UserID)
	} else {
		query := `
			SELECT id, user_id, price, service_name, start_date, created_at, updated_at
			FROM subscriptions
		`
		rows, err = r.db.Query(ctx, query)
	}

	if err != nil {
		r.logger.Error("Failed to fetch subscriptions from DB", slog.Any("err", err))
		return nil, err
	}
	defer rows.Close()

	var results []model.Subscription
	for rows.Next() {
		var ent entities.Subscription
		if err := rows.Scan(&ent.ID, &ent.UserID, &ent.Price, &ent.ServiceName, &ent.StartDate, &ent.CreatedAt, &ent.UpdatedAt); err != nil {
			r.logger.Error("Failed to scan subscription row", slog.Any("err", err))
			continue
		}
		results = append(results, entities.EntityToModel(ent))
	}

	return results, nil
}

func (r *SubRepository) SumSubs(ctx context.Context, from, to time.Time) (int, error) {
	query := `
        SELECT COALESCE(SUM(price), 0)
        FROM subscriptions
        WHERE start_date >= $1 AND start_date <= $2
    `
	var sum int
	err := r.db.QueryRow(ctx, query, from, to).Scan(&sum)
	if err != nil {
		r.logger.Error("Failed to sum subscriptions", slog.Any("err", err))
		return 0, err
	}
	return sum, nil
}
