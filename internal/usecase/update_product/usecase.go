package updateproduct

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

func (u *usecase) UpdateProduct(ctx context.Context, product domain.Product) error {
	return u.repo.UpdateProduct(ctx, product)
}
