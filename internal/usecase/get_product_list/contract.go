package getproductlist

import (
	"context"
	"trainingOnlineStore/internal/domain"
)

type repository interface {
	GetProductList(ctx context.Context, page uint32, pageSize uint32) ([]domain.Product, error)
}
