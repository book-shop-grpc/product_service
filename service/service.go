package service

import (
	"log"
	"net"

	"github.com/book-shop-grpc/product_service/config"
	"github.com/book-shop-grpc/product_service/genproto/product_service"
	db "github.com/book-shop-grpc/product_service/pkg/postgres"
	"github.com/book-shop-grpc/product_service/storage"
	"google.golang.org/grpc"
)

func Service() {
	log.Println("service is running:")

	lis, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		log.Println("error on listen service", err)
		return
	}

	cfg := config.Load()

	db, err := db.ConnectToDb(cfg.PgConfig)
	if err != nil {
		log.Println("error on connect to db:",err)
		return
	}

	storage := storage.NewStorage(db)

	ProductService := NewProductService(storage)
	server := grpc.NewServer()

	product_service.RegisterProductServiceServer(server, ProductService)

	log.Println("service is running:",lis.Addr())
	server.Serve(lis)
}
