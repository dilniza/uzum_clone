package catalog

import (
	"context"
	"service/config"
	"service/genproto/catalog_service"
	"service/grpc/client"
	"service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type ProductService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ProductService {
	return &ProductService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *ProductService) Create(ctx context.Context, req *catalog_service.CreateProduct) (resp *catalog_service.Product, err error) {
	f.log.Info("---CreateProduct--->>>", logger.Any("req", req))

	resp, err = f.strg.Product().Create(ctx, req)
	if err != nil {
		f.log.Error("---CreateProduct--->>>", logger.Error(err))
		return &catalog_service.Product{}, err
	}
	return resp, nil
}

func (f *ProductService) GetByID(ctx context.Context, req *catalog_service.ProductPrimaryKey) (resp *catalog_service.Product, err error) {
	f.log.Info("---GetSingleProduct--->>>", logger.Any("req", req))

	resp, err = f.strg.Product().GetByID(ctx, req)
	if err != nil {
		f.log.Error("---GetSingleProduct--->>>", logger.Error(err))
		return &catalog_service.Product{}, err
	}

	return resp, nil
}

func (f *ProductService) GetAll(ctx context.Context, req *catalog_service.GetAllProductRequest) (resp *catalog_service.GetAllProductResponse, err error) {
	f.log.Info("---GetAllCategories--->>>", logger.Any("req", req))

	resp, err = f.strg.Product().GetAll(ctx, req)
	if err != nil {
		f.log.Error("---GetAllCategories--->>>", logger.Error(err))
		return &catalog_service.GetAllProductResponse{}, err
	}

	return resp, nil
}

func (f *ProductService) Update(ctx context.Context, req *catalog_service.UpdateProduct) (resp *catalog_service.Product, err error) {
	f.log.Info("---UpdateProduct--->>>", logger.Any("req", req))

	resp, err = f.strg.Product().Update(ctx, req)
	if err != nil {
		f.log.Error("---UpdateProduct--->>>", logger.Error(err))
		return &catalog_service.Product{}, err
	}

	return resp, nil
}

func (f *ProductService) Delete(ctx context.Context, req *catalog_service.ProductPrimaryKey) (resp *catalog_service.Empty, err error) {
	f.log.Info("---DeleteProduct--->>>", logger.Any("req", req))

	err = f.strg.Product().Delete(ctx, req)
	if err != nil {
		f.log.Error("---DeleteProduct--->>>", logger.Error(err))
		return &catalog_service.Empty{}, err
	}

	return &catalog_service.Empty{}, nil
}
