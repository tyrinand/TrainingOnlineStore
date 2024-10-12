package handler

import (
	"context"
	"fmt"
	pb "trainingOnlineStore/internal/generated/grpc/product"
)

func (h *Handler) RemoveProductById(ctx context.Context, in *pb.SearchRequestById) (*pb.OperationStatusResponse, error) {
	id := in.GetId()
	err := h.removeUc.RemoveProductById(ctx, id)

	if err != nil {
		errorMessage := fmt.Errorf("useCase RemoveProductById : %v", err)
		errorMessageStr := errorMessage.Error()
		return &pb.OperationStatusResponse{Status: false, Error: &errorMessageStr}, errorMessage
	}

	return &pb.OperationStatusResponse{Status: true}, nil
}
