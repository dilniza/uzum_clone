package order

import (
	"context"
	ct "service/genproto/order_service"
	"service/config"
	"service/grpc/client"
	"service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type OrderService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewOrderService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *OrderService {
	return &OrderService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *OrderService) Create(ctx context.Context, req *ct.CreateOrder) (resp *ct.GetOrder, err error) {

	f.log.Info("---CreateCategory--->>>", logger.Any("req", req))

	resp, err = f.strg.Order().Create(ctx, req)
	if err != nil {
		f.log.Error("---CreateCategory--->>>", logger.Error(err))
		return &ct.GetOrder{}, err
	}

	return resp, nil
}

func (f *OrderService) GetByID(ctx context.Context, req *ct.OrderPrimaryKey) (resp *ct.GetOrder, err error) {
	f.log.Info("---GetSingleCategory--->>>", logger.Any("req", req))

	resp, err = f.strg.Order().GetByID(ctx, req)
	if err != nil {
		f.log.Error("---GetSingleCategory--->>>", logger.Error(err))
		return &ct.GetOrder{}, err
	}

	return resp, nil
}
