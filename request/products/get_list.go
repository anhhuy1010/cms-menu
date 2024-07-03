package products

import "time"

type (
	GetListRequest struct {
		Keyword     string     `form:"keyword"`
		Name        *string    `form:"name"`
		Page        int        `form:"page"`
		Limit       int        `form:"limit"`
		Sort        string     `form:"sort"`
		IsActive    *int       `form:"is_active" `
		Sequence    *int       `form:"sequence"`
		MinQuantity *int       `form:"min_quantity"`
		MaxQuantity *int       `form:"max_quantity"`
		Date        *time.Time `form:"date"`
		Price       int        `form:"price"`
	}
	ListResponse struct {
		Uuid      string    `json:"uuid" `
		Name      string    `json:"name"`
		Image     string    `json:"image"`
		Price     int       `json:"price"`
		IsActive  int       `json:"is_active"`
		Sequence  int       `json:"sequence"`
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Quantity  int       `json:"quantity"`
	}
)
