package service

import (
	"context"
	"uzum_clone/config"
	"uzum_clone/genproto/category_service"
	"uzum_clone/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type CategoryService struct {
	cfg            config.Config
	log            logger.LoggerI
	strg           storage.StorageI
	categoryClient category_service.CategoryServiceClient
	*category_service.UnimplementedCategoryServiceServer
}

func NewCategoryService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, categoryClient category_service.CategoryServiceClient) *CategoryService {
	return &CategoryService{
		cfg:            cfg,
		log:            log,
		strg:           strg,
		categoryClient: categoryClient,
	}
}

func (f *CategoryService) Create(ctx context.Context, req *category_service.CreateCategory) (resp *category_service.Category, err error) {

	f.log.Info("---CreateCategory--->>>", logger.Any("req", req))

	resp, err = f.strg.Category().Create(ctx, req)
	if err != nil {
		f.log.Error("---CreateCategory--->>>", logger.Error(err))
		return &category_service.Category{}, err
	}

	return resp, nil
}

func (f *CategoryService) GetByID(ctx context.Context, req *category_service.CategoryPrimaryKey) (resp *category_service.Category, err error) {
	f.log.Info("---GetSingleCategory--->>>", logger.Any("req", req))

	resp, err = f.strg.Category().GetByID(ctx, req)
	if err != nil {
		f.log.Error("---GetSingleCategory--->>>", logger.Error(err))
		return &category_service.Category{}, err
	}

	return resp, nil
}

func (f *CategoryService) GetAll(ctx context.Context, req *category_service.GetListCategoryRequest) (resp *category_service.GetListCategoryResponse, err error) {
	f.log.Info("---GetAllCategories--->>>", logger.Any("req", req))

	resp, err = f.strg.Category().GetAll(ctx, req)
	if err != nil {
		f.log.Error("---GetAllCategories--->>>", logger.Error(err))
		return &category_service.GetListCategoryResponse{}, err
	}

	return resp, nil
}

func (f *CategoryService) Update(ctx context.Context, req *category_service.UpdateCategory) (resp *category_service.Category, err error) {
	f.log.Info("---UpdateCategory--->>>", logger.Any("req", req))

	resp, err = f.strg.Category().Update(ctx, req)
	if err != nil {
		f.log.Error("---UpdateCategory--->>>", logger.Error(err))
		return &category_service.Category{}, err
	}

	return resp, nil
}

func (f *CategoryService) Delete(ctx context.Context, req *category_service.CategoryPrimaryKey) (resp *category_service.Empty, err error) {
	f.log.Info("---DeleteCategory--->>>", logger.Any("req", req))

	err = f.strg.Category().Delete(ctx, req)
	if err != nil {
		f.log.Error("---DeleteCategory--->>>", logger.Error(err))
		return &category_service.Empty{}, err
	}

	return &category_service.Empty{}, nil
}
