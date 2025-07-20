package service

import (
	"context"
	"log/slog"
	"subscriptions/internal/model"
	"time"
)

type SubRepository interface {
	CreateSub(ctx context.Context, sub model.Subscription) (model.Subscription, error)
	GetByID(ctx context.Context, id int) (model.Subscription, error)
	UpdateSubPrice(ctx context.Context, sub model.Subscription) (model.Subscription, error)
	DeleteByID(ctx context.Context, id int) error
	ListSubs(ctx context.Context, sum model.Subscription) ([]model.Subscription, error)
	SumSubs(ctx context.Context, from, to time.Time) (int, error)
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

func (s *SubService) GetByID(ctx context.Context, id int) (model.Subscription, error) {
	sub, err := s.subRepo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("Failed to find subscription in service layer", slog.Any("err", err))
		return model.Subscription{}, err
	}

	return sub, nil
}

func (s *SubService) UpdateSubPrice(ctx context.Context, sub model.Subscription) (model.Subscription, error) {
	sub, err := s.subRepo.UpdateSubPrice(ctx, sub)
	if err != nil {
		s.logger.Error("Failed to change subscription in service layer", slog.Any("err", err))
		return model.Subscription{}, err
	}

	return sub, nil
}

func (s *SubService) DeleteSub(ctx context.Context, id int) error {
	err := s.subRepo.DeleteByID(ctx, id)
	if err != nil {
		s.logger.Error("Service failed to delete subscription", slog.Int("id", id), slog.Any("err", err))
		return err
	}

	return nil
}

func (s *SubService) ListSubs(ctx context.Context, sub model.Subscription) ([]model.Subscription, error) {
	subs, err := s.subRepo.ListSubs(ctx, sub)
	if err != nil {
		s.logger.Error("Failed to fetch subscriptions in service layer", slog.Any("err", err))
		return nil, err
	}
	return subs, nil
}

func (s *SubService) SumSubs(ctx context.Context, from, to time.Time) (int, error) {
	return s.subRepo.SumSubs(ctx, from, to)
}
