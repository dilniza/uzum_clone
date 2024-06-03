package grpc_client

import (
	"fmt"
	"uzum_clone/config"
	"uzum_clone/genproto/category_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GrpcClientI ...
type GrpcClientI interface {
	CategoryService() category_service.CategoryServiceClient
}

// GrpcClient ...
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {

	connCatalog, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.CatalogServiceHost, cfg.CatalogServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("catalog service dial host: %s port:%s err: %s",
			cfg.CatalogServiceHost, cfg.CatalogServicePort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"category_service": category_service.NewCategoryServiceClient(connCatalog),
		},
	}, nil
}

func (g *GrpcClient) CategoryService() category_service.CategoryServiceClient {
	return g.connections["category_service"].(category_service.CategoryServiceClient)
}

// Implement the Category() method
func (g *GrpcClient) Category() category_service.CategoryServiceClient {
	return g.CategoryService() // Return the category client
}
