package client

import (
	"fmt"
	"service/config"
	"service/genproto/catalog_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	Category() catalog_service.CategoryServiceClient
	Product() catalog_service.ProductServiceClient
}

type grpcClients struct {
	categoryClient catalog_service.CategoryServiceClient
	productClient  catalog_service.ProductServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connCatalog, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.CategoryServiceHost, cfg.CategoryGRPCPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("catalog service dial host: %s port:%s err: %s",
			cfg.CategoryServiceHost, cfg.CategoryGRPCPort, err)
	}

	categoryClient := catalog_service.NewCategoryServiceClient(connCatalog)
	productClient := catalog_service.NewProductServiceClient(connCatalog)

	return &grpcClients{
		categoryClient: categoryClient,
		productClient:  productClient,
	}, nil
}

func (g *grpcClients) Category() catalog_service.CategoryServiceClient {
	return g.categoryClient
}

func (g *grpcClients) Product() catalog_service.ProductServiceClient {
	return g.productClient
}
