package mapper

import (
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"subscriptions/internal/handler/dto"
	"subscriptions/internal/model"
)

func CreateReqToModel(req dto.SubReq) (model.Subscription, error) {
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return model.Subscription{}, fmt.Errorf("invalid user_id: %w", err)
	}

	startDate, err := parseMonthYear(req.StartDate)
	if err != nil {
		return model.Subscription{}, fmt.Errorf("invalid start_date: %w", err)
	}

	return model.Subscription{
		ServiceName: req.ServiceName,
		Price:       req.Price,
		UserID:      userID,
		StartDate:   startDate,
	}, nil
}

func ModelToResp(sub model.Subscription) dto.SubResp {
	return dto.SubResp{
		ID:          sub.UserID.String(),
		ServiceName: sub.ServiceName,
		Price:       strconv.Itoa(sub.Price),
		UserID:      sub.UserID.String(),
		StartDate:   sub.StartDate.Format("01-2006"), // "07-2025"
	}
}

func parseMonthYear(input string) (time.Time, error) {
	full := "01-" + input
	return time.Parse("02-01-2006", full)
}

func UpdatePriceReqToModel(req dto.UpdatePriceRequest, id int) model.Subscription {
	return model.Subscription{
		ID:    id,
		Price: req.Price,
	}
}

func ModelToUpdatePriceResp(sub model.Subscription) dto.UpdateResp {
	return dto.UpdateResp{
		ID:          sub.UserID.String(),
		ServiceName: sub.ServiceName,
		Price:       strconv.Itoa(sub.Price),
		UserID:      sub.UserID.String(),
		StartDate:   sub.StartDate.Format("01-2006"),
		UpdatedAt:   sub.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
