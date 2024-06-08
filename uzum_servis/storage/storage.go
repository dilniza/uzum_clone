package storage

import (
	"context"
	ct "service/genproto/catalog_service"
)

type StorageI interface {
	CloseDB()
	Category() CategoryRepoI
	Product() ProductRepoI
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
