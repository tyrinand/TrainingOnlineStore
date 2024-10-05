package updateproduct

import (
	"context"
	"trainingOnlineStore/internal/domain"
)

type repository interface {
	UpdateProduct(ctx context.Context, product domain.Product) error
}
