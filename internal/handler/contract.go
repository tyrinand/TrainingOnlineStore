package handler

import (
	"context"
	"trainingOnlineStore/internal/domain"
)

type createProductUseCase interface {
	CreateProduct(ctx context.Context, product domain.Product) error
}

type getProductByIdUseCase interface {
	GetProductById(ctx context.Context, id uint32) (domain.Product, error)
}

type getProductListUseCase interface {
	GetProductList(ctx context.Context, page uint32, pageSize uint32) ([]domain.Product, error)
}

type updateProductUseCase interface {
	UpdateProduct(ctx context.Context, product domain.Product) error
}

type removeProductUseCase interface {
	RemoveProductById(ctx context.Context, id uint32) error
}
