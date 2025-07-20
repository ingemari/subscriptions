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

	//price, err := strconv.Atoi(req.Price)
	//if err != nil {
	//	return model.Subscription{}, fmt.Errorf("invalid price: %w", err)
	//}

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
