package order

import (
	"context"
	st "service/genproto/order_service"
	"service/config"
	"service/grpc/client"
	"service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type Order_statusService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewOrderstatus(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *Order_statusService {
	return &Order_statusService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *Order_statusService) Create(ctx context.Context, req *st.CreateStatusRequest) (resp *st.GetOrderStatusResponse, err error) {

	f.log.Info("---CreateStatus--->>>", logger.Any("req", req))
	resp, err = f.strg.OrderStatus().Create(ctx, req)

	if err != nil {
		f.log.Error("---CreateStatus--->>>", logger.Error(err))
		return &st.GetOrderStatusResponse{}, err
	}

	return resp, nil
}

func (f *Order_statusService) GetByID(ctx context.Context, req *st.OrderPrimaryKeyRequest) (resp *st.GetOrderStatusResponse, err error) {

	f.log.Info("---GetsinglStatus--->>>", logger.Any("req", req))

	resp, err = f.strg.OrderStatus().GetByID(ctx, req)

	if err != nil {
		f.log.Error("---GetSingleStatus--->>>", logger.Error(err))

		return &st.GetOrderStatusResponse{}, err
	}
	return resp, nil
}

func (f *Order_statusService) PUTCH(ctx context.Context, req *st.OrderPrimaryStatusKeyRequest) (resp *st.GetOrderStatusResponse, err error) {

	f.log.Info("---PutchStatus--->>>", logger.Any("req", req))

	resp, err = f.strg.OrderStatus().PUTCH(ctx, req)

	if err != nil {
		f.log.Error("---PutchStatus--->>>", logger.Error(err))

		return &st.GetOrderStatusResponse{}, err
	}
	return resp, nil
}

func (f *Order_statusService) GetStatusByID(ctx context.Context, req *st.OrderPrimaryStatusKeyRequest) (resp *st.GetOrderStatusResponse, err error) {

	f.log.Info("---PutchStatus--->>>", logger.Any("req", req))

	resp, err = f.strg.OrderStatus().GetStatusByID(ctx, req)

	if err != nil {
		f.log.Error("---PutchStatus--->>>", logger.Error(err))

		return &st.GetOrderStatusResponse{}, err
	}
	return resp, nil
}
