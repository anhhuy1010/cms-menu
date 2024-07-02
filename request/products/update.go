package products

type (
	UpdateUri struct {
		Uuid string `uri:"uuid"`
	}
	UpdateRequest struct {
		IsActive *int `json:"is_active"`
	}
)
