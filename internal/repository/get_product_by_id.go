package repository

import (
	"context"
	"trainingOnlineStore/internal/domain"
	"trainingOnlineStore/internal/repository/models"
)

func (r *repository) GetProductById(ctx context.Context, id uint32) (domain.Product, error) {

	productDto := models.ProductDto{}

	err := r.db.Get(&productDto, "SELECT P.id, P.name, P.price FROM product P WHERE P.Id = $1", id)

	if err != nil {
		return domain.Product{}, err
	}

	return domain.Product{Id: productDto.Id, Name: productDto.Name, Price: productDto.Price}, nil
}
