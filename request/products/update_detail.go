package products

import "time"

type (
	UpdateUriDetail struct {
		Uuid string `uri:"uuid"`
	}
	UpdateDetailRequest struct {
		Name        string     `json:"name"`
		Image       string     `json:"image"`
		IsActive    *int       `json:"is_active"`
		IsDelete    *int       `json:"is_delete"`
		Price       int        `json:"price"`
		Sequence    int        `json:"sequence"`
		Quantity    int        `json:"quantity"`
		Description string     `json:"description"`
		Gallery     []string   `json:"gallery"`
		StartDate   *time.Time `json:"start_date"`
		EndDate     *time.Time `json:"end_date"`
	}
)
