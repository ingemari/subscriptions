package entities

import (
	"subscriptions/internal/model"
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID          int       `db:"id"`
	ServiceName string    `db:"service_name"`
	Price       int       `db:"price"`
	UserID      uuid.UUID `db:"user_id"`
	StartDate   time.Time `db:"start_date"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func ModelToEntity(m model.Subscription) Subscription {
	return Subscription{
		ID:          m.ID,
		ServiceName: m.ServiceName,
		Price:       m.Price,
		UserID:      m.UserID,
		StartDate:   m.StartDate,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func EntityToModel(e Subscription) model.Subscription {
	return model.Subscription{
		ID:          e.ID,
		ServiceName: e.ServiceName,
		Price:       e.Price,
		UserID:      e.UserID,
		StartDate:   e.StartDate,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}
