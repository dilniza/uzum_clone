package grpc_client

import (
	pc "api/genproto/catalog_service"
	os "api/genproto/order_service"
	us "api/genproto/user_service"
	"log"

	"fmt"

	"api/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GrpcClientI ...
type GrpcClientI interface {
	CategoryService() pc.CategoryServiceClient
	ProductService() pc.ProductServiceClient
	ProductReviewService() pc.ProductReviewServiceClient
	ProductCategoryService() pc.ProductCategoryServiceClient

	UserService() us.CustomerServiceClient
	SystemUserService() us.UsServiceClient
	SellerServive() us.SellerServiceClient
	BranchService() us.BranchServiceClient
	ShopService() us.ShopServiceClient

	ProducOrderService() os.OrderProductsServiceClient
	OrderService() os.OrderServiceServer
	OrderStatus() os.OrderStatusServiceClient
}

// GrpcClient ...
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {

	connCategory, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.CatalogServiceHost, cfg.CatalogServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("category service dial host: %s port:%s err: %s",
			cfg.CatalogServiceHost, cfg.CatalogServicePort, err)
	}

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("user service dial host: %s port:%s err: %s",
			cfg.UserServiceHost, cfg.UserServicePort, err)
	}

	connOrder, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.OrderServiceHost, cfg.OrderServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("order service dial hsot:%s port :%s err:%s",
			cfg.OrderServiceHost, cfg.OrderServicePort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"catalog_service":          pc.NewCategoryServiceClient(connCategory),
			"product_service":          pc.NewProductServiceClient(connCategory),
			"product_review_service":   pc.NewProductReviewServiceClient(connCategory),
			"product_category_service": pc.NewProductCategoryServiceClient(connCategory),

			"user_service": us.NewCustomerServiceClient(connUser),
			"system_user":  us.NewUsServiceClient(connUser),
			"seller":       us.NewSellerServiceClient(connUser),
			"branch":       us.NewBranchServiceClient(connUser),
			"shop":         us.NewShopServiceClient(connUser),

			"order_product_service": os.NewOrderProductsServiceClient(connOrder),
			"order_service":         os.NewOrderServiceClient(connOrder),
			"order_status_service":  os.NewOrderStatusServiceClient(connOrder),
		},
	}, nil
}

func (g *GrpcClient) CategoryService() pc.CategoryServiceClient {
	return g.connections["catalog_service"].(pc.CategoryServiceClient)
}

func (g *GrpcClient) ProductService() pc.ProductServiceClient {
	return g.connections["product_service"].(pc.ProductServiceClient)
}

func (g *GrpcClient) ProductReviewService() pc.ProductReviewServiceClient {
	return g.connections["product_review_service"].(pc.ProductReviewServiceClient)
}

func (g *GrpcClient) ProductCategoryService() pc.ProductCategoryServiceClient {
	return g.connections["product_category_service"].(pc.ProductCategoryServiceClient)
}

func (g *GrpcClient) UserService() us.CustomerServiceClient {
	client, ok := g.connections["user_service"].(us.CustomerServiceClient)
	if !ok {
		log.Println("failed to assert type for user_service")
		return nil
	}
	return client
}

func (g *GrpcClient) SystemUserService() us.UsServiceClient {
	client, ok := g.connections["system_user"].(us.UsServiceClient)
	if !ok {
		log.Println("failed to assert type for system_user")
		return nil
	}
	return client
}

func (g *GrpcClient) SellerService() us.SellerServiceClient {
	client, ok := g.connections["seller"].(us.SellerServiceClient)
	if !ok {
		log.Println("failed to assert type for seller")
		return nil
	}
	return client
}

func (g *GrpcClient) BranchService() us.BranchServiceClient {
	client, ok := g.connections["branch"].(us.BranchServiceClient)
	if !ok {
		log.Println("failed to assert type for branch")
		return nil
	}
	return client
}

func (g *GrpcClient) ShopService() us.ShopServiceClient {
	client, ok := g.connections["shop"].(us.ShopServiceClient)
	if !ok {
		log.Println("failed to assert type for shop")
		return nil
	}
	return client
}

func (o *GrpcClient) ProducOrderService() os.OrderProductsServiceClient {
	return o.connections["order_product_service"].(os.OrderProductsServiceClient)
}

func (o *GrpcClient) OrderService() os.OrderServiceClient {
	return o.connections["order_service"].(os.OrderServiceClient)
}

func (o *GrpcClient) OrderStatus() os.OrderStatusServiceClient {
	return o.connections["order_status_service"].(os.OrderStatusServiceClient)
}
