package handler

import (
	"api/genproto/order_service"
	"api/pkg/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router /order-status [POST]
// @Summary		Update order status by ID
// @Description	This API create the status of an order
// @Tags		order_status
// @Produce		json
// @Param        order body order_service.CreateStatusRequest true "Order Status"
// @Success		200 {object} models.Response
// @Failure		400 {object} models.Response
// @Failure		404 {object} models.Response
// @Failure		500 {object} models.Response
func (h *Handler) CreateOrderstatus(c *gin.Context) {

	order := &order_service.CreateStatusRequest{}

	if err := c.ShouldBindJSON(order); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while reading request body")
		return
	}

	if err := helper.Validstatus(order.Status); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "some error input status err ")
		return
	}

	fmt.Println(order)
	resp, err := h.grpcClient.OrderStatus().Create(c.Request.Context(), order)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while creating order")
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Router		/order-status/{id} [GET]
// @Summary		Get by id a orderstatus
// @Description	This api get a order by id
// @Tags		order_status
// @Produce		json
// @Param 		id path order_service.OrderPrimaryKeyRequest true "Order Status"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) OrderStatusById(c *gin.Context) {

	id := &order_service.OrderPrimaryKeyRequest{}
	id.Id = c.Param("id")
	fmt.Println(id, "id______________________")
	data, err := h.grpcClient.OrderStatus().GetByID(c.Request.Context(), id)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while reading request id")
		return
	}

	c.JSON(http.StatusCreated, data)
}

// @Router     /order-status [PATCH]
// @Summary    Update order status by ID
// @Description This API updates the status of an order by ID
// @Tags       order_status
// @Produce    json
// @Param      orders body order_service.OrderPrimaryStatusKeyRequest true "Order Status"
// @Success    200 {object} models.Response
// @Failure    400 {object} models.Response
// @Failure    404 {object} models.Response
// @Failure    500 {object} models.Response
func (h *Handler) OrderStatusPatch(c *gin.Context) {
	// Create an instance of OrderPrimaryStatusKeyRequest
	order := &order_service.OrderPrimaryStatusKeyRequest{}

	// Bind JSON input to the order object
	if err := c.ShouldBindJSON(order); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while reading request body")
		return
	}
	//// Get the current status of the order
	data, err := h.grpcClient.OrderStatus().GetStatusByID(c.Request.Context(), order)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while fetching current order status")
		return
	}

	fmt.Println("Current status:", data.Status)
	fmt.Println("New status:", order.Status)

	// Validate the new status against the current status
	if err := helper.Validstatusorder(data.Status, order.Status); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "invalid status transition")
		return
	}
	// Update the order status
	resp, err := h.grpcClient.OrderStatus().PUTCH(c.Request.Context(), order)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while updating order status")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/order-statusp/{id} [GET]
// @Summary		Get by id a orderstatusputch
// @Description	This API gets an order by ID
// @Tags		order_statusp
// @Produce		json
// @Param 		id path string true "Order Status ID"
// @Success		200 {object} models.Response
// @Failure		400 {object} models.Response
// @Failure		404 {object} models.Response
// @Failure		500 {object} models.Response
func (h *Handler) OrderStatusputch(c *gin.Context) {
	// Create an instance of OrderPrimaryStatusKeyRequest
	order := &order_service.OrderPrimaryStatusKeyRequest{}

	// Capture the 'id' parameter from the URL path
	order.OrderId = c.Param("id")

	// Log the captured ID for debugging
	fmt.Println(order.OrderId, "id______________________")

	// Fetch the status by ID using gRPC client
	data, err := h.grpcClient.OrderStatus().GetStatusByID(c.Request.Context(), order)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while reading request id")
		return
	}

	// Return the response
	c.JSON(http.StatusOK, data)
}
