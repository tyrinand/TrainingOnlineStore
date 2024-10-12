package getproductbyid

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

func (u *usecase) GetProductById(ctx context.Context, id uint32) (domain.Product, error) {
	return u.repo.GetProductById(ctx, id)
}
