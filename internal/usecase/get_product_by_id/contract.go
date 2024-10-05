package getproductbyid

import (
	"context"
	"trainingOnlineStore/internal/domain"
)

type repository interface {
	GetProductById(ctx context.Context, id uint32) (*domain.Product, error)
}
