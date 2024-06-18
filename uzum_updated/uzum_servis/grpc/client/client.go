package client

import (
	"fmt"
	"service/config"
	"service/genproto/user_service"
	"service/genproto/catalog_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	Category() catalog_service.CategoryServiceClient
	Product() catalog_service.ProductServiceClient
	ProductCategory() catalog_service.ProductCategoryServiceClient
	ProductReview() catalog_service.ProductReviewServiceClient
	
	User() user_service.UsServiceClient
	Customer() user_service.CustomerServiceClient
	Seller() user_service.SellerServiceClient
	Branch() user_service.BranchServiceClient
	Shop() user_service.ShopServiceClient
}

type grpcClients struct {
	categoryClient catalog_service.CategoryServiceClient
	productClient  catalog_service.ProductServiceClient
	productCategoryClient catalog_service.ProductCategoryServiceClient
	productReviewClient catalog_service.ProductReviewServiceClient
	
	userClient user_service.UsServiceClient
	customerClient user_service.CustomerServiceClient
	sellerClient user_service.SellerServiceClient
	branchClient user_service.BranchServiceClient
	shopClient user_service.ShopServiceClient	
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connCatalog, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.CategoryServiceHost, cfg.CategoryGRPCPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("catalog service dial host: %s port:%s err: %s",
			cfg.CategoryServiceHost, cfg.CategoryGRPCPort, err)
	}
	
	connUser, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("user service dial host: %s port:%s err: %s",
			cfg.UserServiceHost, cfg.UserServicePort, err)
	}

	categoryClient := catalog_service.NewCategoryServiceClient(connCatalog)
	productClient := catalog_service.NewProductServiceClient(connCatalog)
	productCategoryClient := catalog_service.NewProductCategoryServiceClient(connCatalog)
	productReviewClient := catalog_service.NewProductReviewServiceClient(connCatalog)
	
	userClient := user_service.NewUsServiceClient(connUser)
	
	

	return &grpcClients{
		categoryClient: categoryClient,
		productClient:  productClient,
		productCategoryClient: productCategoryClient,
		productReviewClient: productReviewClient,
		userClient: userClient,
	}, nil
}

func (g *grpcClients) Category() catalog_service.CategoryServiceClient {
	return g.categoryClient
}

func (g *grpcClients) Product() catalog_service.ProductServiceClient {
	return g.productClient
}

func (g *grpcClients) ProductCategory() catalog_service.ProductCategoryServiceClient {
	return g.productCategoryClient
}

func (g *grpcClients) ProductReview() catalog_service.ProductReviewServiceClient {
	return g.productReviewClient
}

func (g *grpcClients) User() user_service.UsServiceClient {
	return g.userClient
}


func (g *grpcClients) Customer() user_service.CustomerServiceClient {
	return g.customerClient
}

func (g *grpcClients) Seller() user_service.SellerServiceClient {
	return g.sellerClient
}

func (g *grpcClients) Branch() user_service.BranchServiceClient {
	return g.branchClient
}


func (g *grpcClients) Shop() user_service.ShopServiceClient {
	return g.shopClient
}