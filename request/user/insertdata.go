package user

import "time"

type (
	GetInsertRequest struct {
		Uuid      string    `json:"uuid" `
		Name      string    `json:"name"`
		Image     string    `json:"image"`
		IsActive  int       `json:"is_active"`
		Price     int       `json:"price"`
		Sequence  int       `json:"sequence"`
		Quantity  int       `json:"quantity"`
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
	}
	InsertResponse struct {
		Uuid     string `json:"uuid" `
		Name     string `json:"name"`
		Image    string `json:"image"`
		Price    int    `json:"price"`
		Sequence int    `json:"sequence"`
		Quantity int    `json:"quantity"`
	}
)
