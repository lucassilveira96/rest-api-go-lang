package product

import "context"

type Repository interface {
	ListProducts(ctx context.Context) ([]*Product, error)
	CreateProduct(ctx context.Context, product *Product) error
}
