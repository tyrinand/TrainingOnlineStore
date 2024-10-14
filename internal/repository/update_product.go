package repository

import (
	"context"
	"trainingOnlineStore/internal/domain"
	"trainingOnlineStore/internal/repository/models"
)

func (r *repository) UpdateProduct(ctx context.Context, product domain.Product) error {

	productDto := models.ProductDto{Id: product.Id, Name: product.Name, Price: product.Price}
	_, err := r.db.NamedExec("UPDATE product SET name = :name, price = :price WHERE id = :id", &productDto)
	return err
}
