package dto

type SubReq struct {
	ServiceName string `json:"service_name" db:"service_name"`
	Price       string `json:"price" db:"price"`
	UserID      string `json:"user_id" db:"user_id"`
	StartDate   string `json:"start_date" db:"start_date"`
}

type SubResp struct {
	ServiceName string `json:"service_name" db:"service_name"`
	Price       string `json:"price" db:"price"`
	UserID      string `json:"user_id" db:"user_id"`
	StartDate   string `json:"start_date" db:"start_date"`
}
