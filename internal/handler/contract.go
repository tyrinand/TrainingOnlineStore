package handler

import "trainingOnlineStore/internal/domain"

type usecase interface {
	GetProductById(id uint32) (product domain.Product) //что возвращать если не найден ???
	GetProductList(pageInfo domain.PageInfo) (list []domain.Product)
	CreateProduct(product domain.Product) error
	UpdateProduct(product domain.Product) error
	RemoveProduct(id uint32) error
}
