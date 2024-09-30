package create_product

import (
	"context"
	"trainingOnlineStore/internal/domain"
)

type repository interface {
	CreateProduct(ctx context.Context, product domain.Product) error
}
