package product

import "context"

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) ListProducts(ctx context.Context) ([]*Product, error) {
	return s.repository.ListProducts(ctx)
}

func (s *Service) CreateProduct(ctx context.Context, product *Product) error {
	return s.repository.CreateProduct(ctx, product)
}
