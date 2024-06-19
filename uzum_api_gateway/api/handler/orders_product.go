package handler

import (
	"api/genproto/order_service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router /order [post]
// @Summary Create Order_product
// @Description API for creating a order_service
// @Tags  order_service
// @Accept       json
// @Produce      json
// @Param        order body order_service.CreateOrderProducts true "Order_Product"
// @Success      201 {object} models.ResponseSuccess
// @Failure      404 {object} models.ResponseError
// @Failure      500 {object} models.ResponseError
func (h *Handler) CreateOrderProducts(c *gin.Context) {

	order := &order_service.CreateOrderProducts{}

	if err := c.ShouldBindJSON(order); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while reading request body")
		return
	}
	fmt.Println(order)
	resp, err := h.grpcClient.ProducOrderService().Create(c.Request.Context(), order)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while creating order")
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Router		/order/{id} [get]
// @Summary		Get by id a order
// @Description	This api get a order by id
// @Tags		order_service
// @Produce		json
// @Param 		id path order_service.OrderProductsPrimaryKey true "order_service.OrderProductsPrimaryKey"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) OrderProductetById(c *gin.Context) {

	id := &order_service.OrderProductsPrimaryKey{}
	id.Id = c.Param("id")
	fmt.Println(id, "id______________________")
	data, err := h.grpcClient.ProducOrderService().GetByID(c.Request.Context(), id)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while reading request id")
		return
	}

	c.JSON(http.StatusCreated, data)
}
