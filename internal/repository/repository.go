package repository

import (
	"context"
	"trainingOnlineStore/internal/domain"
)

type repository struct {
}

func New() *repository {
	return &repository{}
}

func (r *repository) CreateProduct(ctx context.Context, product domain.Product) error {
	return nil
}

func (r *repository) GetProductById(ctx context.Context, id uint32) (*domain.Product, error) {
	return &domain.Product{Id: 0, Name: "Хлеб", Price: 46}, nil
}

func (r *repository) GetProductList(ctx context.Context, page domain.PageInfo) ([]domain.Product, error) {
	var list []domain.Product = make([]domain.Product, 3)

	list[0] = domain.Product{Id: 0, Name: "Хлеб", Price: 46}
	list[1] = domain.Product{Id: 1, Name: "Сыр", Price: 270}
	list[2] = domain.Product{Id: 2, Name: "Масло", Price: 200}

	return list, nil
}

func (r *repository) UpdateProduct(ctx context.Context, product domain.Product) error {
	return nil
}

func (r *repository) RemoveProductById(ctx context.Context, id uint32) error {
	return nil
}
