package service

import (
	"context"
	"log/slog"
	"subscriptions/internal/model"
)

type SubRepository interface {
	CreateSub(ctx context.Context, sub model.Subscription) (model.Subscription, error)
}

type SubService struct {
	subRepo SubRepository
	logger  *slog.Logger
}

func NewSubService(pr SubRepository, logger *slog.Logger) *SubService {
	return &SubService{subRepo: pr, logger: logger}
}

func (s *SubService) CreateSub(ctx context.Context, sub model.Subscription) (model.Subscription, error) {
	sub, err := s.subRepo.CreateSub(ctx, sub)
	if err != nil {
		s.logger.Error("Failed to create subscription in service layer", slog.Any("err", err))
		return model.Subscription{}, err
	}

	return sub, nil
}
