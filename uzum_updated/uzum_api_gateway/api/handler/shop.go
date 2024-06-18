package handler

import (
	"api/genproto/user_service"
	"api/pkg/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router        /CreateShop [post]
// @Summary       Create shop
// @Description   API for creating shop
// @Tags          shop
// @Accept        json
// @Produce       json
// @Param         shop body user_service.CreateShop true "shop"
// @Success 200   {object} user_service.Shop
// @Failure 404   {object} models.ResponseError
// @Failure 500   {object} models.ResponseError
func (h *Handler) CreateShop(c *gin.Context) {
	var (
		req  user_service.CreateShop
		resp *user_service.Shop
		err  error
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	if err := helper.ValidatePhone(req.Phone); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while validating phone number "+req.Phone)
		return
	}

	resp, err = h.grpcClient.ShopService().Create(c.Request.Context(), &req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "failed to create shop")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router         /GetListShop [GET]
// @Summary        Get list of shops
// @Description    API for getting list of shops
// @Tags           shop
// @Accept         json
// @Produce        json
// @Param		   shop query string false "shops"
// @Param		   page query int false "page"
// @Param		   limit query int false "limit"
// @Success 200    {object} user_service.GetListShopResponse
// @Failure 404    {object} models.ResponseError
// @Failure 500    {object} models.ResponseError
func (h *Handler) GetListShop(c *gin.Context) {

	var (
		req  user_service.GetListShopRequest
		resp *user_service.GetListShopResponse
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

	resp, err = h.grpcClient.ShopService().GetList(c.Request.Context(), &req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router         /GetByIdShop/{id} [GET]
// @Summary        Get a single shop by ID
// @Description    API for getting a single shop by ID
// @Tags           shop
// @Accept         json
// @Produce        json
// @Param          id path string true "shop ID"
// @Success        200 {object} user_service.Shop
// @Failure        404 {object} models.ResponseError
// @Failure        500 {object} models.ResponseError
func (h *Handler) GetShopByID(c *gin.Context) {
	var (
		id   = c.Param("id")
		resp *user_service.Shop
		err  error
	)

	req := &user_service.ShopPrimaryKey{
		Id: id,
	}

	resp, err = h.grpcClient.ShopService().GetByID(c.Request.Context(), req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router          /UpdateShop/{id} [PUT]
// @Summary         Update a shop by ID
// @Description     API for updating a shop by ID
// @Tags            shop
// @Accept          json
// @Produce         json
// @Param           id path string true "shop ID"
// @Param           shop body user_service.UpdateShop true "shop"
// @Success         200 {object} user_service.Shop
// @Failure         404 {object} models.ResponseError
// @Failure         500 {object} models.ResponseError
func (h *Handler) UpdateShop(c *gin.Context) {
	var (
		id   = c.Param("id")
		req  user_service.UpdateShop
		resp *user_service.Shop
		err  error
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	if err := helper.ValidatePhone(req.Phone); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while validating phone number "+req.Phone)
		return
	}

	req.Id = id
	resp, err = h.grpcClient.ShopService().Update(c.Request.Context(), &req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router        /DeleteShop/{id} [DELETE]
// @Summary       Delete a shop by ID
// @Description   API for deleting a shop by ID
// @Tags          shop
// @Accept        json
// @Produce       json
// @Param         id path string true "shop ID"
// @Success       200 {object} user_service.Empty4
// @Failure       404 {object} models.ResponseError
// @Failure       500 {object} models.ResponseError
func (h *Handler) DeleteShop(c *gin.Context) {
	var (
		id   = c.Param("id")
		err  error
		resp *user_service.Empty4
	)

	req := &user_service.ShopPrimaryKey{
		Id: id,
	}

	resp, err = h.grpcClient.ShopService().Delete(c.Request.Context(), req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
