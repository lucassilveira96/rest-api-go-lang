package mapper

import (
	"rest-api-go-lang/internal/app/domain/product"
	"rest-api-go-lang/internal/app/web/data/request"
)

func CreateProductRequestToProduct(req *request.CreateProduct) *product.Product {
	return &product.Product{
		Description:   req.Description,
		Value_Product: req.Value_Product,
	}
}
