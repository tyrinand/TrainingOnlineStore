package removeproductbyid

import "context"

type repository interface {
	RemoveProductById(ctx context.Context, id uint32) error
}
