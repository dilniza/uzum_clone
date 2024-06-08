package grpc

import (
	"service/config"
	"service/genproto/catalog_service"
	"service/grpc/client"
	"service/grpc/service/catalog_service"
	"service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()

	catalog_service.RegisterCategoryServiceServer(grpcServer, catalog.NewCategoryService(cfg, log, strg, srvc))
	catalog_service.RegisterProductServiceServer(grpcServer, catalog.NewProductService(cfg, log, strg, srvc))
	reflection.Register(grpcServer)
	return
}
