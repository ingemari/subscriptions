package dto

type SubReq struct {
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	UserID      string `json:"user_id"`
	StartDate   string `json:"start_date"`
}

type SubResp struct {
	ID          string `json:"id"`
	ServiceName string `json:"service_name"`
	Price       string `json:"price"`
	UserID      string `json:"user_id"`
	StartDate   string `json:"start_date"`
}

type UpdatePriceRequest struct {
	Price int `json:"price"`
}

type UpdateResp struct {
	ID          string `json:"id"`
	ServiceName string `json:"service_name"`
	Price       string `json:"price"`
	UserID      string `json:"user_id"`
	StartDate   string `json:"start_date"`
	UpdatedAt   string `json:"updated_at"`
}

type ListReq struct {
	UserID string `json:"user_id"`
}

type SumReq struct {
	StartDateFrom string `form:"start_date_from"`
	StartDateTo   string `form:"start_date_to"`
}
