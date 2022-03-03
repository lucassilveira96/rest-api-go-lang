package repository

import (
	"context"
	"database/sql"
	"fmt"
	"rest-api-go-lang/internal/app/domain/product"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) ListProducts(ctx context.Context) ([]*product.Product, error) {
	rows, err := r.db.QueryContext(ctx, listProductsQuery)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()
	products := make([]*product.Product, 0)

	for rows.Next() {
		var entity product.Product
		if err := rows.Scan(&entity.ID, &entity.Description, &entity.Value_Product); err != nil {
			return nil, err
		}

		products = append(products, &entity)
	}

	return products, nil
}

func (r *ProductRepository) CreateProduct(ctx context.Context, product *product.Product) error {
	return r.db.QueryRowContext(ctx, createProductQuery, product.Description, product.Value_Product).Scan(&product.ID)
}

const (
	listProductsQuery  = "select id,description,value_product from products"
	createProductQuery = "insert into products (description, value_product) values ($1, $2) RETURNING id"
)
