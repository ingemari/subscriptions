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
