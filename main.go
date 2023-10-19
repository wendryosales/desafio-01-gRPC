package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/wendryosales/desafio-01-gRPC/application/grpc"
	"github.com/wendryosales/desafio-01-gRPC/infrastructure/db"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))

	grpc.StartGrpcServer(database, 50051)

}
