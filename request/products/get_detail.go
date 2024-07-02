package products

import "time"

type (
	GetDetailUri struct {
		Uuid string `uri:"uuid"`
	}
	GetDetailResponse struct {
		Uuid        string    `json:"uuid"`
		Price       int       `json:"price"`
		Image       string    `json:"image"`
		Name        string    `json:"name"`
		Sequence    int       `json:"sequence"`
		Quantity    int       `json:"quantity"`
		Description string    `json:"description"`
		Gallery     []string  `json:"gallery"`
		IsActive    int       `json:"is_active"`
		StartDate   time.Time `json:"start_date"`
		EndDate     time.Time `json:"end_date"`
	}
)
