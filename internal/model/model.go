package model

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID          int
	ServiceName string
	Price       int
	UserID      uuid.UUID
	StartDate   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
