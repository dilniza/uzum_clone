package handler

import (
	"api/genproto/order_service"
	"api/pkg/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router /orders [post]
// @Summary Create Order
// @Description API for creating a order
// @Tags   orders
// @Accept       json
// @Produce      json
// @Param        order body order_service.CreateOrder true "Order"
// @Success      201 {object} models.ResponseSuccess
// @Failure      404 {object} models.ResponseError
// @Failure      500 {object} models.ResponseError
func (h *Handler) CreateOrders(c *gin.Context) {

	order := &order_service.CreateOrder{}

	if err := c.ShouldBindJSON(order); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while reading request body")
		return
	}

	if err := helper.ValidatePhone(order.CourierPhone); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "some err  CourierPhone ")

		return
	}

	if err := helper.ValidatePhone(order.CustomerPhone); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "some err  CustomerPhone ")

		return
	}

	if err := helper.Validatetype(order.Type); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "some error input type ")
		return
	}

	if err := helper.Validatepayment_type(order.PaymentType); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "some error input payment_type ")
		return
	}
	if err := helper.Validstatus(order.Status); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "some error input status err ")
		return
	}

	fmt.Println(order)
	resp, err := h.grpcClient.OrderService().Create(c.Request.Context(), order)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while creating order")
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Router		/orders/{id} [get]
// @Summary		Get by id a order
// @Description	This api get a order by id
// @Tags		orders
// @Produce		json
// @Param 		id path order_service.OrderPrimaryKey true "order_product.OrderProductsPrimaryKey"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) OrderstById(c *gin.Context) {

	id := &order_service.OrderPrimaryKey{}
	id.Id = c.Param("id")
	fmt.Println(id, "id______________________")
	data, err := h.grpcClient.OrderService().GetByID(c.Request.Context(), id)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while reading request id")
		return
	}

	c.JSON(http.StatusCreated, data)
}
