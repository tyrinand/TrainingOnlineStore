package getproductlist

import (
	"context"
	"trainingOnlineStore/internal/domain"
)

type repository interface {
	GetProductList(ctx context.Context, page domain.PageInfo) ([]domain.Product, error)
}
