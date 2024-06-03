package grpc

import (
	"uzum_clone/config"
	"uzum_clone/genproto/category_service"
	"uzum_clone/grpc/client"
	"uzum_clone/grpc/service"
	"uzum_clone/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()

	categoryClient := srvc.Category()

	category_service.RegisterCategoryServiceServer(grpcServer, service.NewCategoryService(cfg, log, strg, categoryClient))
	reflection.Register(grpcServer)
	return
}
