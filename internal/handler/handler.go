package handler

import (
	"context"
	"fmt"
	"trainingOnlineStore/internal/domain"
	pb "trainingOnlineStore/internal/generated/grpc/product"
)

type Handler struct {
	pb.UnimplementedProductServiceServer
	uc createProductUseCase
}

func New(uc createProductUseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) CreateProduct(ctx context.Context, in *pb.ProductMessage) (*pb.OperationStatusResponse, error) {
	product := domain.Product{
		Id:    in.GetId(),
		Name:  in.GetName(),
		Price: in.GetPrice(),
	}
	err := h.uc.CreateProduct(ctx, product)

	if err != nil {
		Status := pb.OperationStatusResponse{Status: false, Error: fmt.Sprintf("create product faild : %v", err)}
		return &Status, fmt.Errorf("useCase CreateProduct : %v", err)
	}

	return &pb.OperationStatusResponse{Status: true, Error: ""}, nil
}
