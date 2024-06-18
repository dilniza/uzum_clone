package handler

import (
	"api/genproto/catalog_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /getlistProductReview [GET]
// @Summary Get List Category Reviews
// @Description API for getting list category reviews
// @Tags category
// @Accept json
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Param search query string false "Search term"
// @Success 200 {object} catalog_service.GetAllProductReviewResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *Handler) GetListProductReview(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "Invalid limit")
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "Invalid offset")
		return
	}

	search := c.DefaultQuery("search", "")

	in := &catalog_service.GetAllProductReviewRequest{
		Limit:  int64(limit),
		Offset: int64(offset),
		Search: search,
	}

	resp, err := h.grpcClient.ProductReviewService().GetAll(c, in)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while getting all category reviews")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /ProductReview [POST]
// @Summary Create a category review
// @Description API for creating a category review
// @Tags category
// @Accept json
// @Produce json
// @Param ProductReview body catalog_service.CreateProductReview true "category review"
// @Success 200 {object} catalog_service.ProductReview
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *Handler) CreateProductReview(c *gin.Context) {
	var (
		req  catalog_service.CreateProductReview
		resp *catalog_service.ProductReview
		err  error
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	resp, err = h.grpcClient.ProductReviewService().Create(c, &req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while creating category review")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /ProductReview/{id} [GET]
// @Summary Get a single category review by ID
// @Description API for getting a single category review by ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "category review ID"
// @Success 200 {object} catalog_service.ProductReview
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *Handler) GetProductReviewByID(c *gin.Context) {
	var (
		id   = c.Param("id")
		resp *catalog_service.ProductReview
		err  error
	)

	req := &catalog_service.ProductReviewPrimaryKey{
		Id: id,
	}

	resp, err = h.grpcClient.ProductReviewService().GetByID(c, req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while getting category review by id")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /ProductReview/{id} [PUT]
// @Summary Update a category review by ID
// @Description API for updating a category review by ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "category review ID"
// @Param ProductReview body catalog_service.UpdateProductReview true "category review"
// @Success 200 {object} catalog_service.ProductReview
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *Handler) UpdateProductReview(c *gin.Context) {
	var (
		id   = c.Param("id")
		req  catalog_service.UpdateProductReview
		resp *catalog_service.ProductReview
		err  error
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	req.Id = id
	resp, err = h.grpcClient.ProductReviewService().Update(c, &req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while updating category review")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /ProductReview/{id} [DELETE]
// @Summary Delete a category review by ID
// @Description API for deleting a category review by ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "category review ID"
// @Success 200 {object} catalog_service.Empty
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *Handler) DeleteProductReview(c *gin.Context) {
	var (
		id   = c.Param("id")
		err  error
		resp *catalog_service.Empty
	)

	req := &catalog_service.ProductReviewPrimaryKey{
		Id: id,
	}

	resp, err = h.grpcClient.ProductReviewService().Delete(c, req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while deleting category review")
		return
	}

	c.JSON(http.StatusOK, resp)
}


