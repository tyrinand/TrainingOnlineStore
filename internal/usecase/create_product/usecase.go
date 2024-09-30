package create_product

import (
	"context"
	"trainingOnlineStore/internal/domain"
)

type usecase struct {
	repo repository
}

func New(r repository) *usecase {
	return &usecase{repo: r}
}

func (u *usecase) CreateProduct(ctx context.Context, product domain.Product) error {
	return nil // реализовать метод
}
