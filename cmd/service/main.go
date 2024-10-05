package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	pb "trainingOnlineStore/internal/generated/grpc/product"
	"trainingOnlineStore/internal/handler"
	"trainingOnlineStore/internal/repository"
	create "trainingOnlineStore/internal/usecase/create_product"
	getbyid "trainingOnlineStore/internal/usecase/get_product_by_id"
	getlist "trainingOnlineStore/internal/usecase/get_product_list"
	remove "trainingOnlineStore/internal/usecase/remove_product_by_id"
	update "trainingOnlineStore/internal/usecase/update_product"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	repo := repository.New()

	productCreateUc := create.New(repo)
	productGetByIdUc := getbyid.New(repo)
	productGetListUc := getlist.New(repo)
	productUpdateUc := update.New(repo)
	productRemoveUc := remove.New(repo)

	productHandler := handler.New(productCreateUc, productGetByIdUc, productGetListUc, productUpdateUc, productRemoveUc)

	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, productHandler)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
