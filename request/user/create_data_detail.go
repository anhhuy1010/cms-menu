package user

type (
	GetCreateRequest struct {
		Uuid        string   `json:"uuid" `
		Name        string   `json:"name"`
		Image       string   `json:"image"`
		IsActive    int      `json:"is_active"`
		Price       int      `json:"price"`
		Sequence    int      `json:"sequence"`
		Quantity    int      `json:"quantity"`
		Description string   `json:"description"`
		Gallery     []string `json:"gallery"`
	}
	CreateResponse struct {
		Uuid        string   `json:"uuid" `
		Name        string   `json:"name"`
		Image       string   `json:"image"`
		IsActive    int      `json:"is_active"`
		IsDelete    int      `json:"is_delete"`
		Price       int      `json:"price"`
		Sequence    int      `json:"sequence"`
		Quantity    int      `json:"quantity"`
		Description string   `json:"description"`
		Gallery     []string `json:"gallery"`
	}
)
