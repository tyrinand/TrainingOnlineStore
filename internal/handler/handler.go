package handler

import (
	pb "trainingOnlineStore/internal/generated/grpc/product"
)

type Handler struct {
	pb.UnimplementedProductServiceServer
	createUc  createProductUseCase
	getByIdUc getProductByIdUseCase
	getListUc getProductListUseCase
	updateUc  updateProductUseCase
	removeUc  removeProductUseCase
}

func New(
	createUseCase createProductUseCase,
	getByIdUseCase getProductByIdUseCase,
	getListUseCase getProductListUseCase,
	updateUseCase updateProductUseCase,
	removeUseCase removeProductUseCase,
) *Handler {
	return &Handler{
		createUc:  createUseCase,
		getByIdUc: getByIdUseCase,
		getListUc: getListUseCase,
		updateUc:  updateUseCase,
		removeUc:  removeUseCase,
	}
}
