package main

import (
	"github.com/fullcycle-dev/codepix/application/grpc"
	"github.com/fullcycle-dev/codepix/infrastructure/db"
	"github.com/jinzhu/gorm"
	"os"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
