package catalog

import (
	"context"
	"service/config"
	"service/genproto/catalog_service"
	"service/grpc/client"
	"service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type ProductReviewService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewProductReviewService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ProductReviewService {
	return &ProductReviewService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *ProductReviewService) Create(ctx context.Context, req *catalog_service.CreateProductReview) (resp *catalog_service.ProductReview, err error) {
	f.log.Info("---CreateProductReview--->>>", logger.Any("req", req))

	resp, err = f.strg.ProductReview().Create(ctx, req)
	if err != nil {
		f.log.Error("---CreateProductReview--->>>", logger.Error(err))
		return &catalog_service.ProductReview{}, err
	}
	return resp, nil
}

func (f *ProductReviewService) GetByID(ctx context.Context, req *catalog_service.ProductReviewPrimaryKey) (resp *catalog_service.ProductReview, err error) {
	f.log.Info("---GetSingleProductReview--->>>", logger.Any("req", req))

	resp, err = f.strg.ProductReview().GetByID(ctx, req)
	if err != nil {
		f.log.Error("---GetSingleProductReview--->>>", logger.Error(err))
		return &catalog_service.ProductReview{}, err
	}

	return resp, nil
}

func (f *ProductReviewService) GetAll(ctx context.Context, req *catalog_service.GetAllProductReviewRequest) (resp *catalog_service.GetAllProductReviewResponse, err error) {
	f.log.Info("---GetAllCategories--->>>", logger.Any("req", req))

	resp, err = f.strg.ProductReview().GetAll(ctx, req)
	if err != nil {
		f.log.Error("---GetAllCategories--->>>", logger.Error(err))
		return &catalog_service.GetAllProductReviewResponse{}, err
	}

	return resp, nil
}

func (f *ProductReviewService) Update(ctx context.Context, req *catalog_service.UpdateProductReview) (resp *catalog_service.ProductReview, err error) {
	f.log.Info("---UpdateProductReview--->>>", logger.Any("req", req))

	resp, err = f.strg.ProductReview().Update(ctx, req)
	if err != nil {
		f.log.Error("---UpdateProductReview--->>>", logger.Error(err))
		return &catalog_service.ProductReview{}, err
	}

	return resp, nil
}

func (f *ProductReviewService) Delete(ctx context.Context, req *catalog_service.ProductReviewPrimaryKey) (resp *catalog_service.Empty, err error) {
	f.log.Info("---DeleteProductReview--->>>", logger.Any("req", req))

	err = f.strg.ProductReview().Delete(ctx, req)
	if err != nil {
		f.log.Error("---DeleteProductReview--->>>", logger.Error(err))
		return &catalog_service.Empty{}, err
	}

	return &catalog_service.Empty{}, nil
}
