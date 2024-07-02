package user

type (
	UpdateUri struct {
		Uuid string `uri:"uuid"`
	}
	UpdateRequest struct {
		IsDelete *int `json:"is_delete"`
		IsActive *int `json:"is_active"`
	}
)
