package repository

import (
	"context"
	"trainingOnlineStore/internal/domain"
	"trainingOnlineStore/internal/repository/models"
)

func (r *repository) CreateProduct(ctx context.Context, product domain.Product) error {
	// как реализовать подобие try{} catch{} ???
	// try-catch - нужен ли

	productDto := models.ProductDto{Name: product.Name, Price: product.Price}

	tx := r.db.MustBegin()
	tx.NamedExec("INSERT INTO product (name, price) VALUES (:name, :price)", &productDto)

	return tx.Commit()
}
