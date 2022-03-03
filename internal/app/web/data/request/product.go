package request

type CreateProduct struct {
	Description   string  `json:"description" validate:"required"`
	Value_Product float64 `json:"value" validate:"required"`
}
