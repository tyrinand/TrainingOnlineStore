package repository

import (
	"context"
	"trainingOnlineStore/internal/domain"
	"trainingOnlineStore/internal/repository/models"
)

func (r *repository) GetProductList(ctx context.Context, page uint32, pageSize uint32) ([]domain.Product, error) {

	offset := (page - 1) * pageSize
	query := "SELECT P.id, P.name, P.price FROM product P ORDER BY P.id LIMIT $1 OFFSET $2"
	products := []models.ProductDto{}

	err := r.db.Select(&products, query, pageSize, offset)

	if err != nil {
		return []domain.Product{}, err
	}

	result := make([]domain.Product, len(products))

	for index, item := range products {
		result[index] = domain.Product{Id: item.Id, Name: item.Name, Price: item.Price}
	}

	return result, nil
}
