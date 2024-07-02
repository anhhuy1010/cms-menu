package user

type (
	GetDetailUri struct {
		Uuid string `uri:"uuid"`
	}
	GetDetailResponse struct {
		Name  string `json:"name"`
		Image string `json:"image"`
		Price int    `json:"price"`
	}
)
