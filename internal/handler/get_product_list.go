package handler

import (
	"context"
	pb "trainingOnlineStore/internal/generated/grpc/product"
)

func (h *Handler) GetProductList(ctx context.Context, in *pb.SearchRequestList) (*pb.ProductListResponse, error) {

	page := in.GetPage()
	pageSize := in.GetPageSize()

	list, err := h.getListUc.GetProductList(ctx, page, pageSize)

	if err != nil {
		return nil, err
	}

	resultList := make([]*pb.ProductMessage, len(list))

	for index, item := range list {
		resultList[index] = &pb.ProductMessage{Id: item.Id, Name: item.Name, Price: item.Price}
	}

	return &pb.ProductListResponse{List: resultList}, nil

}
