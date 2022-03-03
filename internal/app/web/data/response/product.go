package response

type (
	ListProducts struct {
		ID            uint64  `json:"id"`
		Description   string  `json:"description"`
		Value_Product float64 `json:"value"`
	}

	CreateProduct struct {
		ID uint64 `json:"id"`
	}
)
