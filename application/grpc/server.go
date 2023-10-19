package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"github.com/wendryosales/desafio-01-gRPC/application/grpc/pb"
	"github.com/wendryosales/desafio-01-gRPC/application/usecase"
	"github.com/wendryosales/desafio-01-gRPC/infrastructure/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	productRepository := repository.ProductRepositoryDb{Db: database}
	productUseCase := usecase.ProductUseCase{ProductRepository: productRepository}
	productGrpcService := NewProductGrpcService(productUseCase)

	pb.RegisterProductServiceServer(grpcServer, productGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, error := net.Listen("tcp", address)
	if error != nil {
		log.Fatalf("Could not listen on port %d: %v", port, error)
	}

	log.Printf("gRPC server running on port %d", port)

	error = grpcServer.Serve(listener)
	if error != nil {
		log.Fatalf("Could not start gRPC server: %v", error)
	}
}
