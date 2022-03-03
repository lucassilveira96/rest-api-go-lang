package presenter

import (
	"rest-api-go-lang/internal/app/domain/product"
	"rest-api-go-lang/internal/app/web/data/response"
)

func ListProducts(products []*product.Product) []*response.ListProducts {
	resp := make([]*response.ListProducts, len(products), len(products))
	for i, p := range products {
		resp[i] = &response.ListProducts{
			ID:            p.ID,
			Description:   p.Description,
			Value_Product: p.Value_Product,
		}
	}

	return resp
}

func CreateProduct(p *product.Product) *response.CreateProduct {
	return &response.CreateProduct{
		ID: p.ID,
	}
}
