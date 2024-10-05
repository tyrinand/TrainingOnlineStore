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

func (u *usecase) GetProductList(ctx context.Context, pageInfo domain.PageInfo) ([]domain.Product, error) {
	return u.repo.GetProductList(ctx, pageInfo)
}
