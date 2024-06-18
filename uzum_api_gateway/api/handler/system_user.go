package handler

import (
	"api/genproto/user_service"
	"api/pkg/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router        /CreateUser [post]
// @Summary       Create user
// @Description   API for creating user
// @Tags          user
// @Accept        json
// @Produce       json
// @Param order   body     user_service.CreateUs true "user"
// @Success 200   {object} user_service.CreateUs
// @Failure 404   {object} models.ResponseError
// @Failure 500   {object} models.ResponseError
func (h *Handler) CreateUser(c *gin.Context) {
	var (
		req  user_service.CreateUs
		resp *user_service.Us
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

	if err := helper.ValidateEmailAddress(req.Gmail); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "erroe while validating email"+req.Gmail)
	}

	System_user := h.grpcClient.SystemUserService()
	if System_user == nil {
		HandleGrpcErrWithDescription(c, h.log, err, "system service client not initialized")
		return
	}

	resp, err = System_user.Create(c.Request.Context(), &req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "failed to create customer")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router         /GetListUser [GET]
// @Summary        Get List user
// @Description    API for getting list user
// @Tags           user
// @Accept         json
// @Produce        json
// @Param		   system_user query string false "system_users"
// @Param		   page query int false "page"
// @Param		   limit query int false "limit"
// @Success 200    {object} user_service.GetListUsResponse
// @Failure 404    {object} models.ResponseError
// @Failure 500    {object} models.ResponseError
func (h *Handler) GetListUser(c *gin.Context) {
	var (
		req  user_service.GetListUsRequest
		resp *user_service.GetListUsResponse
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

	System_user := h.grpcClient.SystemUserService()
	if System_user == nil {
		HandleGrpcErrWithDescription(c, h.log, err, "system service client not initialized")
		return
	}

	resp, err = System_user.GetList(c.Request.Context(), &req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router         /GetByIdUser/{id} [GET]
// @Summary        Get a single user by ID
// @Description    API for getting a single user by ID
// @Tags           user
// @Accept         json
// @Produce        json
// @Param          id path string true "user ID"
// @Success        200 {object} user_service.Us
// @Failure        404 {object} models.ResponseError
// @Failure        500 {object} models.ResponseError
func (h *Handler) GetUserByID(c *gin.Context) {
	var (
		id   = c.Param("id")
		resp *user_service.Us
		err  error
	)

	req := &user_service.UsPrimaryKey{
		Id: id,
	}

	System_user := h.grpcClient.SystemUserService()
	if System_user == nil {
		HandleGrpcErrWithDescription(c, h.log, err, "system service client not initialized")
		return
	}

	resp, err = System_user.GetByID(c.Request.Context(), req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router          /UpdateUser/{id} [PUT]
// @Summary         Update a user by ID
// @Description     API for updating a user by ID
// @Tags            user
// @Accept          json
// @Produce         json
// @Param           id path string true "user ID"
// @Param           category body user_service.UpdateUs true "user"
// @Success         200 {object} user_service.Us
// @Failure         404 {object} models.ResponseError
// @Failure         500 {object} models.ResponseError
func (h *Handler) UpdateUser(c *gin.Context) {
	var (
		id   = c.Param("id")
		req  user_service.UpdateUs
		resp *user_service.Us
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

	if err := helper.ValidateEmailAddress(req.Gmail); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "erroe while validating email"+req.Gmail)
	}

	System_user := h.grpcClient.SystemUserService()
	if System_user == nil {
		HandleGrpcErrWithDescription(c, h.log, err, "system service client not initialized")
		return
	}

	req.Id = id
	resp, err = System_user.Update(c.Request.Context(), &req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router        /DeleteUser/{id} [DELETE]
// @Summary       Delete a user by ID
// @Description   API for deleting a user by ID
// @Tags          user
// @Accept        json
// @Produce       json
// @Param         id path string true "user ID"
// @Success       200 {object} user_service.Empty
// @Failure       404 {object} models.ResponseError
// @Failure       500 {object} models.ResponseError
func (h *Handler) DeleteUser(c *gin.Context) {
	var (
		id   = c.Param("id")
		err  error
		resp *user_service.Empty1
	)

	req := &user_service.UsPrimaryKey{
		Id: id,
	}

	System_user := h.grpcClient.SystemUserService()
	if System_user == nil {
		HandleGrpcErrWithDescription(c, h.log, err, "system service client not initialized")
		return
	}

	resp, err = System_user.Delete(c.Request.Context(), req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
