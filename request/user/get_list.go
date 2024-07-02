package user

import "time"

type (
	GetListRequest struct {
		Keyword   string     `form:"keyword"`
		Name      *string    `form:"name"`
		Page      int        `form:"page"`
		Limit     int        `form:"limit"`
		Sort      string     `form:"sort"`
		IsActive  *int       `form:"is_active" `
		Sequence  int        `form:"sequence"`
		Quantity  *int       `form:"quantity"`
		StartDate *time.Time `form:"start_date"`
		EndDate   *time.Time `form:"end_date"`
	}
	ListResponse struct {
		Uuid       string `json:"uuid" `
		ClientUuid string `json:"client_uuid"`
		Name       string `json:"name"`
		Image      string `json:"image"`
		Price      int    `json:"price"`
		IsActive   int    `json:"is_active"`
		Sequence   int    `json:"sequence"`
	}
)
