package grpc_client

import (
	"fmt"
	pc "service/genproto/catalog_service"

	"service/config"

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

	connCategory, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.CategoryServiceHost, cfg.CategoryGRPCPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("category service dial host: %s port:%s err: %s",
			cfg.CategoryServiceHost, cfg.CategoryGRPCPort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"category_service": pc.NewCategoryServiceClient(connCategory),
		},
	}, nil
}

func (g *GrpcClient) CategoryService() pc.CategoryServiceClient {
	return g.connections["category_service"].(pc.CategoryServiceClient)
}
