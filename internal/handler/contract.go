package handler

import (
	"context"
	"trainingOnlineStore/internal/domain"
)

//TO DO
// add ctx - на всех уровнях

//ctx.candcel - отвалился при timeOut
// снижение нагрузки

// type usecase interface {
// 	GetProductById(id uint32) (product domain.Product) //что возвращать если не найден ???
// 	GetProductList(pageInfo domain.PageInfo) (list []domain.Product)
// 	CreateProduct(ctx context.Context, product domain.Product) error
// 	UpdateProduct(product domain.Product) error
// 	RemoveProduct(id uint32) error
// }

type createProductUseCase interface {
	CreateProduct(ctx context.Context, product domain.Product) error
}

// handler - 1
// repository
