package order

import (
	"context"
	op "service/genproto/order_service"
	"service/config"
	"service/grpc/client"
	"service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type ProducOrderService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewProductOrderService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvices client.ServiceManagerI) *ProducOrderService {

	return &ProducOrderService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvices,
	}
}

func (s *ProducOrderService) Create(ctx context.Context, req *op.CreateOrderProducts) (resp *op.GetOrderProducts, err error) {

	s.log.Info("---CreateProductOrder--->>>", logger.Any("req", req))
	resp, err = s.strg.OrderProduct().Create(ctx, req)
	if err != nil {
		s.log.Error("---CreateProductOrder--->>>", logger.Error(err))
		return &op.GetOrderProducts{}, err
	}

	return resp, nil
}

func (s *ProducOrderService) GetByID(ctx context.Context, req *op.OrderProductsPrimaryKey) (resp *op.GetOrderProducts, err error) {
	s.log.Info("---GetbyidProductOrder--->>>", logger.Any("req", req))

	resp, err = s.strg.OrderProduct().GetByID(ctx, req)

	if err != nil {
		s.log.Error("---GetbyidProductOrder--->>>", logger.Error(err))
		return &op.GetOrderProducts{}, err

	}

	return resp, nil

}
