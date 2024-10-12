package handler

import (
	"context"
	"fmt"
	"trainingOnlineStore/internal/domain"
	pb "trainingOnlineStore/internal/generated/grpc/product"
)

func (h *Handler) UpdateProduct(ctx context.Context, in *pb.ProductMessage) (*pb.OperationStatusResponse, error) {
	product := domain.Product{Id: in.GetId(), Name: in.GetName(), Price: in.GetPrice()}

	err := h.updateUc.UpdateProduct(ctx, product)

	if err != nil {
		errorMessage := fmt.Errorf("useCase UpdateProduct : %v", err)
		errorMessageStr := errorMessage.Error()
		return &pb.OperationStatusResponse{Status: false, Error: &errorMessageStr}, errorMessage
	}

	return &pb.OperationStatusResponse{Status: true}, nil
}
