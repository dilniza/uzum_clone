package handler

import (
	"api/pkg/helper"
	"api/genproto/user_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router        /createCustomer [post]
// @Summary       Create customer
// @Description   API for creating customer
// @Tags          customer
// @Accept        json
// @Produce       json
// @Param order   body     user_service.CreateCustomer true "customer"
// @Success 200   {object} user_service.CreateCustomer
// @Failure 404   {object} models.ResponseError
// @Failure 500   {object} models.ResponseError
func (h *Handler) CreateCustomer(c *gin.Context) {

	var (
		req  user_service.CreateCustomer
		resp *user_service.Customer
		err  error
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	if err := helper.ValidatePhone(req.Phone); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while validating phone number"+req.Phone)
		return
	}

	if err := helper.ValidateEmailAddress(req.Email); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "erroe while validating email"+req.Email)
	}

	resp, err = h.grpcClient.UserService().Create(c.Request.Context(), &req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "failed to create customer")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router         /getlistcustomer [GET]
// @Summary        Get List Customer
// @Description    API for getting list customer
// @Tags           customer
// @Accept         json
// @Produce        json
// @Param		   customer query string false "customers"
// @Param		   page query int false "page"
// @Param		   limit query int false "limit"
// @Success 200    {object} user_service.GetListCustomerResponse
// @Failure 404    {object} models.ResponseError
// @Failure 500    {object} models.ResponseError
func (h *Handler) GetListCustomer(c *gin.Context) {
	var (
		req  user_service.GetListCustomerRequest
		resp *user_service.GetListCustomerResponse
		err  error
	)

	req.Search = c.Query("search")

	page, err := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while parsing page")
		return
	}

	limit, err := strconv.ParseUint(c.DefaultQuery("limit", "10"), 10, 64)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while parsing limit")
		return
	}

	req.Page = page
	req.Limit = limit

	resp, err = h.grpcClient.UserService().GetList(c, &req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router         /getbyidcustomer/{id} [GET]
// @Summary        Get a single customer by ID
// @Description    API for getting a single category by ID
// @Tags           customer
// @Accept         json
// @Produce        json
// @Param          id path string true "customer ID"
// @Success        200 {object} user_service.Customer
// @Failure        404 {object} models.ResponseError
// @Failure        500 {object} models.ResponseError
func (h *Handler) GetCustomerByID(c *gin.Context) {
	var (
		id   = c.Param("id")
		resp *user_service.Customer
		err  error
	)

	req := &user_service.CustomerPrimaryKey{
		Id: id,
	}

	resp, err = h.grpcClient.UserService().GetByID(c, req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router          /updateCustomer/{id} [PUT]
// @Summary         Update a customer by ID
// @Description     API for updating a customer by ID
// @Tags            customer
// @Accept          json
// @Produce         json
// @Param           id path string true "customer ID"
// @Param           category body user_service.UpdateCustomer true "customer"
// @Success         200 {object} user_service.Customer
// @Failure         404 {object} models.ResponseError
// @Failure         500 {object} models.ResponseError
func (h *Handler) UpdateCustomer(c *gin.Context) {
	var (
		id   = c.Param("id")
		req  user_service.UpdateCustomer
		resp *user_service.Customer
		err  error
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	if err := helper.ValidatePhone(req.Phone); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while validating phone number"+req.Phone)
		return
	}

	if err := helper.ValidateEmailAddress(req.Email); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "erroe while validating email"+req.Email)
	}

	req.Id = id
	resp, err = h.grpcClient.UserService().Update(c, &req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router        /deleteCustomer/{id} [DELETE]
// @Summary       Delete a customer by ID
// @Description   API for deleting a customer by ID
// @Tags          customer
// @Accept        json
// @Produce       json
// @Param         id path string true "customer ID"
// @Success       200 {object} user_service.Empty
// @Failure       404 {object} models.ResponseError
// @Failure       500 {object} models.ResponseError
func (h *Handler) DeleteCustomer(c *gin.Context) {
	var (
		id   = c.Param("id")
		err  error
		resp *user_service.Empty
	)

	req := &user_service.CustomerPrimaryKey{
		Id: id,
	}

	resp, err = h.grpcClient.UserService().Delete(c, req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
