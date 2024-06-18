package storage

import (
	"context"
	ct "service/genproto/catalog_service"
	us "service/genproto/user_service"
)

type StorageI interface {
	CloseDB()

	Category() CategoryRepoI
	Product() ProductRepoI
	ProductCategory() ProductCategoryRepoI
	ProductReview() ProductReviewRepoI

	Customer() CustomerRepoI
	User() UserRepoI
	Seller() SellerRepoI
	Branch() BranchRepoI
	Shop() ShopRepoI
}

type CategoryRepoI interface {
	Create(ctx context.Context, req *ct.CreateCategory) (resp *ct.Category, err error)
	GetByID(ctx context.Context, req *ct.CategoryPrimaryKey) (resp *ct.Category, err error)
	Update(ctx context.Context, req *ct.UpdateCategory) (resp *ct.Category, err error)
	GetAll(ctx context.Context, req *ct.GetAllCategoryRequest) (resp *ct.GetAllCategoryResponse, err error)
	Delete(ctx context.Context, req *ct.CategoryPrimaryKey) error
}

type ProductRepoI interface {
	Create(ctx context.Context, req *ct.CreateProduct) (resp *ct.Product, err error)
	GetByID(ctx context.Context, req *ct.ProductPrimaryKey) (resp *ct.Product, err error)
	Update(ctx context.Context, req *ct.UpdateProduct) (resp *ct.Product, err error)
	GetAll(ctx context.Context, req *ct.GetAllProductRequest) (resp *ct.GetAllProductResponse, err error)
	Delete(ctx context.Context, req *ct.ProductPrimaryKey) error
}

type ProductCategoryRepoI interface {
	Create(ctx context.Context, req *ct.CreateProductCategory) (resp *ct.ProductCategory, err error)
	GetByID(ctx context.Context, req *ct.ProductCategoryPrimaryKey) (resp *ct.ProductCategory, err error)
	Update(ctx context.Context, req *ct.UpdateProductCategory) (resp *ct.ProductCategory, err error)
	GetAll(ctx context.Context, req *ct.GetAllProductCategoryRequest) (resp *ct.GetAllProductCategoryResponse, err error)
	Delete(ctx context.Context, req *ct.ProductCategoryPrimaryKey) error
}

type ProductReviewRepoI interface {
	Create(ctx context.Context, req *ct.CreateProductReview) (resp *ct.ProductReview, err error)
	GetByID(ctx context.Context, req *ct.ProductReviewPrimaryKey) (resp *ct.ProductReview, err error)
	Update(ctx context.Context, req *ct.UpdateProductReview) (resp *ct.ProductReview, err error)
	GetAll(ctx context.Context, req *ct.GetAllProductReviewRequest) (resp *ct.GetAllProductReviewResponse, err error)
	Delete(ctx context.Context, req *ct.ProductReviewPrimaryKey) error
}

type CustomerRepoI interface {
	Create(ctx context.Context, req *us.CreateCustomer) (*us.Customer, error)
	GetByID(ctx context.Context, req *us.CustomerPrimaryKey) (*us.Customer, error)
	GetList(ctx context.Context, req *us.GetListCustomerRequest) (*us.GetListCustomerResponse, error)
	Update(ctx context.Context, req *us.UpdateCustomer) (*us.Customer, error)
	Delete(ctx context.Context, req *us.CustomerPrimaryKey) error
}

type UserRepoI interface {
	Create(ctx context.Context, req *us.CreateUs) (*us.Us, error)
	GetByID(ctx context.Context, req *us.UsPrimaryKey) (*us.Us, error)
	GetList(ctx context.Context, req *us.GetListUsRequest) (*us.GetListUsResponse, error)
	Update(ctx context.Context, req *us.UpdateUs) (*us.Us, error)
	Delete(ctx context.Context, req *us.UsPrimaryKey) error
}

type SellerRepoI interface {
	Create(ctx context.Context, req *us.CreateSeller) (*us.Seller, error)
	GetByID(ctx context.Context, req *us.SellerPrimaryKey) (*us.Seller, error)
	GetList(ctx context.Context, req *us.GetListSellerRequest) (*us.GetListSellerResponse, error)
	Update(ctx context.Context, req *us.UpdateSeller) (*us.Seller, error)
	Delete(ctx context.Context, req *us.SellerPrimaryKey) error
}

type BranchRepoI interface {
	Create(ctx context.Context, req *us.CreateBranch) (*us.Branch, error)
	GetByID(ctx context.Context, req *us.BranchPrimaryKey) (*us.Branch, error)
	GetList(ctx context.Context, req *us.GetListBranchRequest) (*us.GetListBranchResponse, error)
	Update(ctx context.Context, req *us.UpdateBranch) (*us.Branch, error)
	Delete(ctx context.Context, req *us.BranchPrimaryKey) error
}

type ShopRepoI interface {
	Create(ctx context.Context, req *us.CreateShop) (*us.Shop, error)
	GetById(ctx context.Context, req *us.ShopPrimaryKey) (*us.Shop, error)
	GetList(ctx context.Context, req *us.GetListShopRequest) (*us.GetListShopResponse, error)
	Update(ctx context.Context, req *us.UpdateShop) (*us.Shop, error)
	Delete(ctx context.Context, req *us.ShopPrimaryKey) error
}
