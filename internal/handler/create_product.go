package handler

import (
	"context"
	"fmt"
	"trainingOnlineStore/internal/domain"
	pb "trainingOnlineStore/internal/generated/grpc/product"
)

func (h *Handler) CreateProduct(ctx context.Context, in *pb.ProductMessage) (*pb.OperationStatusResponse, error) {
	product := domain.Product{
		Id:    in.GetId(),
		Name:  in.GetName(),
		Price: in.GetPrice(),
	}
	err := h.createUc.CreateProduct(ctx, product)

	if err != nil {
		errorMessage := fmt.Errorf("useCase CreateProduct : %v", err)
		errorMessageStr := errorMessage.Error()

		Status := pb.OperationStatusResponse{Status: false, Error: &errorMessageStr}
		return &Status, errorMessage
	}

	return &pb.OperationStatusResponse{Status: true}, nil
}
