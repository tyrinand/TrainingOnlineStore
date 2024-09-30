package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	pb "trainingOnlineStore/internal/generated/grpc/product"
	"trainingOnlineStore/internal/handler"
	"trainingOnlineStore/internal/usecase/create_product"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	// flag.Parse()
	// lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// s := grpc.NewServer()
	// pb.RegisterGreeterServer(s, &server{})
	// log.Printf("server listening at %v", lis.Addr())
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }

	// под handler

	//TODO repository

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	productCreateUc := create_product.New(repository)
	productHandler := handler.New(productCreateUc)

	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, productHandler)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
