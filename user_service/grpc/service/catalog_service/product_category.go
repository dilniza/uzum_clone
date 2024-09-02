package catalog

import (
	"context"
	"service/config"
	"service/genproto/catalog_service"
	"service/grpc/client"
	"service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type ProductCategoryService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewProductCategoryService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ProductCategoryService {
	return &ProductCategoryService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *ProductCategoryService) Create(ctx context.Context, req *catalog_service.CreateProductCategory) (resp *catalog_service.ProductCategory, err error) {
	f.log.Info("---CreateProductCategory--->>>", logger.Any("req", req))

	resp, err = f.strg.ProductCategory().Create(ctx, req)
	if err != nil {
		f.log.Error("---CreateProductCategory--->>>", logger.Error(err))
		return &catalog_service.ProductCategory{}, err
	}
	return resp, nil
}

func (f *ProductCategoryService) GetByID(ctx context.Context, req *catalog_service.ProductCategoryPrimaryKey) (resp *catalog_service.ProductCategory, err error) {
	f.log.Info("---GetSingleProductCategory--->>>", logger.Any("req", req))

	resp, err = f.strg.ProductCategory().GetByID(ctx, req)
	if err != nil {
		f.log.Error("---GetSingleProductCategory--->>>", logger.Error(err))
		return &catalog_service.ProductCategory{}, err
	}

	return resp, nil
}

func (f *ProductCategoryService) GetAll(ctx context.Context, req *catalog_service.GetAllProductCategoryRequest) (resp *catalog_service.GetAllProductCategoryResponse, err error) {
	f.log.Info("---GetAllCategories--->>>", logger.Any("req", req))

	resp, err = f.strg.ProductCategory().GetAll(ctx, req)
	if err != nil {
		f.log.Error("---GetAllCategories--->>>", logger.Error(err))
		return &catalog_service.GetAllProductCategoryResponse{}, err
	}

	return resp, nil
}

func (f *ProductCategoryService) Update(ctx context.Context, req *catalog_service.UpdateProductCategory) (resp *catalog_service.ProductCategory, err error) {
	f.log.Info("---UpdateProductCategory--->>>", logger.Any("req", req))

	resp, err = f.strg.ProductCategory().Update(ctx, req)
	if err != nil {
		f.log.Error("---UpdateProductCategory--->>>", logger.Error(err))
		return &catalog_service.ProductCategory{}, err
	}

	return resp, nil
}

func (f *ProductCategoryService) Delete(ctx context.Context, req *catalog_service.ProductCategoryPrimaryKey) (resp *catalog_service.Empty, err error) {
	f.log.Info("---DeleteProductCategory--->>>", logger.Any("req", req))

	err = f.strg.ProductCategory().Delete(ctx, req)
	if err != nil {
		f.log.Error("---DeleteProductCategory--->>>", logger.Error(err))
		return &catalog_service.Empty{}, err
	}

	return &catalog_service.Empty{}, nil
}
