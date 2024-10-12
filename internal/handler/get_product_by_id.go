package handler

import (
	"context"
	"fmt"
	pb "trainingOnlineStore/internal/generated/grpc/product"
)

func (h *Handler) GetProductById(ctx context.Context, in *pb.SearchRequestById) (*pb.ProductItemResponse, error) {
	id := in.GetId()

	product, err := h.getByIdUc.GetProductById(ctx, id)

	if err != nil {
		errorMessage := fmt.Errorf("useCase GetById : %v", err)
		return nil, errorMessage
	}

	responseProduct := pb.ProductMessage{Id: product.Id, Name: product.Name, Price: product.Price}
	response := pb.ProductItemResponse{Item: &responseProduct}

	return &response, nil
}
