package storage

import (
	"context"
	ct "uzum_clone/genproto/category_service"
)

type StorageI interface {
	CloseDB()
	Category() CategoryRepoI
}

type CategoryRepoI interface {
	Create(ctx context.Context, req *ct.CreateCategory) (resp *ct.Category, err error)
	GetByID(ctx context.Context, req *ct.CategoryPrimaryKey) (resp *ct.Category, err error)
	Update(ctx context.Context, req *ct.UpdateCategory) (resp *ct.Category, err error)
	GetAll(ctx context.Context, req *ct.GetListCategoryRequest) (resp *ct.GetListCategoryResponse, err error)
	Delete(ctx context.Context, req *ct.CategoryPrimaryKey) error
}
