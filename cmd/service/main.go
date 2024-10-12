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

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {

	db, err := sqlx.Connect("postgres", "user=tyrinand password=12faga_LA123 dbname=mydatabase sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	if !tableExists(db, "product") {
		createStartTable(db)
	}

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	repo := repository.New(db)

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

	defer db.Close()
}

func createStartTable(db *sqlx.DB) {
	var schema = `
	CREATE TABLE product (
			Id SERIAL PRIMARY KEY,
			Name text,
			Price INTEGER
	);
	`

	db.MustExec(schema)

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO product (Name, Price) VALUES ($1, $2)", "Хлеб", "46")
	tx.MustExec("INSERT INTO product (Name, Price) VALUES ($1, $2)", "Сыр", "270")
	tx.MustExec("INSERT INTO product (Name, Price) VALUES ($1, $2)", "Масло", "200")
	tx.Commit()
}

func tableExists(db *sqlx.DB, tableName string) bool {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = '%s')", tableName)
	err := db.Get(&exists, query)
	if err != nil {
		log.Fatalf("Ошибка при проверке существования таблицы: %v", err)
	}
	return exists
}
