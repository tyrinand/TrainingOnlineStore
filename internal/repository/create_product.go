package repository

import (
	"context"
	"trainingOnlineStore/internal/domain"
	"trainingOnlineStore/internal/repository/models"
)

func (r *repository) CreateProduct(ctx context.Context, product domain.Product) error {
	productDto := models.ProductDto{Name: product.Name, Price: product.Price}
	_, err := r.db.NamedExec("INSERT INTO product (name, price) VALUES (:name, :price)", &productDto)
	return err
}
