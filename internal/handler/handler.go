package handler

import (
	"context"
	"fmt"
	"trainingOnlineStore/internal/domain"
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

func (h *Handler) CreateProduct(ctx context.Context, in *pb.ProductMessage) (*pb.OperationStatusResponse, error) {
	product := domain.Product{
		Id:    in.GetId(),
		Name:  in.GetName(),
		Price: in.GetPrice(),
	}
	err := h.createUc.CreateProduct(ctx, product)

	if err != nil {
		var errorMessage error = fmt.Errorf("useCase CreateProduct : %v", err)
		var errorMessageStr string = errorMessage.Error()

		Status := pb.OperationStatusResponse{Status: false, Error: &errorMessageStr}
		return &Status, errorMessage
	}

	return &pb.OperationStatusResponse{Status: true}, nil
}

func (h *Handler) GetProductById(ctx context.Context, in *pb.SearchRequestById) (*pb.ProductItemResponse, error) {
	var id uint32 = in.GetId()

	var product, err = h.getByIdUc.GetProductById(ctx, id)

	if err != nil {
		var errorMessage error = fmt.Errorf("useCase GetById : %v", err)
		return nil, errorMessage
	}

	var responseProduct pb.ProductMessage = pb.ProductMessage{Id: product.Id, Name: product.Name, Price: product.Price}
	var response pb.ProductItemResponse = pb.ProductItemResponse{Item: &responseProduct}

	return &response, nil
}

func (h *Handler) GetProductList(ctx context.Context, in *pb.SearchRequestList) (*pb.ProductListResponse, error) {
	page := domain.PageInfo{
		Page:     in.GetPage(),
		PageSize: in.GetPageSize(),
	}

	list, err := h.getListUc.GetProductList(ctx, page)

	if err != nil {
		return nil, err
	}

	resultList := make([]*pb.ProductMessage, len(list))

	for index, item := range list {
		resultList[index] = &pb.ProductMessage{Id: item.Id, Name: item.Name, Price: item.Price}
	}

	return &pb.ProductListResponse{List: resultList}, nil

}

func (h *Handler) UpdateProduct(ctx context.Context, in *pb.ProductMessage) (*pb.OperationStatusResponse, error) {
	product := domain.Product{Id: in.GetId(), Name: in.GetName(), Price: in.GetPrice()}

	err := h.updateUc.UpdateProduct(ctx, product)

	if err != nil {
		var errorMessage error = fmt.Errorf("useCase UpdateProduct : %v", err)
		var errorMessageStr string = errorMessage.Error()
		return &pb.OperationStatusResponse{Status: false, Error: &errorMessageStr}, errorMessage
	}

	return &pb.OperationStatusResponse{Status: true}, nil
}

func (h *Handler) RemoveProductById(ctx context.Context, in *pb.SearchRequestById) (*pb.OperationStatusResponse, error) {
	var id uint32 = in.GetId()
	err := h.removeUc.RemoveProductById(ctx, id)

	if err != nil {
		var errorMessage error = fmt.Errorf("useCase RemoveProductById : %v", err)
		var errorMessageStr string = errorMessage.Error()
		return &pb.OperationStatusResponse{Status: false, Error: &errorMessageStr}, errorMessage
	}

	return &pb.OperationStatusResponse{Status: true}, nil
}
