package removeproductbyid

import (
	"context"
)

type usecase struct {
	repo repository
}

func New(r repository) *usecase {
	return &usecase{repo: r}
}

func (u *usecase) RemoveProductById(ctx context.Context, id uint32) error {
	return u.repo.RemoveProductById(ctx, id)
}
