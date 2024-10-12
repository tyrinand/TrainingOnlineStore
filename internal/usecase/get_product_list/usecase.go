package getproductlist

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

func (u *usecase) GetProductList(ctx context.Context, page uint32, pageSize uint32) ([]domain.Product, error) {
	return u.repo.GetProductList(ctx, page, pageSize)
}
