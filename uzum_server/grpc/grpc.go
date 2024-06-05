package grpc

import (
	"server/config"
	"server/genproto/category_service"
	"server/grpc/client"
	"server/grpc/service"
	"server/storage"

	"github.com/saidamir98/udevs_pkg/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()

	category_service.RegisterCategoryServiceServer(grpcServer, service.NewCategoryService(cfg, log, strg, srvc))
	reflection.Register(grpcServer)
	return
}
