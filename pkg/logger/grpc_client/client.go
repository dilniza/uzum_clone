package grpc_client

import (
	"fmt"
	pc "uzum_clone/genproto/category_service"

	"uzum_clone/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GrpcClientI ...
type GrpcClientI interface {
	CategoryService() pc.CategoryServiceClient
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
			"category_service": pc.NewCategoryServiceClient(connCatalog),
		},
	}, nil
}

func (g *GrpcClient) CategoryService() pc.CategoryServiceClient {
	return g.connections["category_service"].(pc.CategoryServiceClient)
}
