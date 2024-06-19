package handler

import (
	"api/genproto/catalog_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /getlistProductCategory [GET]
// @Summary Get List Product Categorys
// @Description API for getting list category categorys
// @Tags category
// @Accept json
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Param search query string false "Search term"
// @Success 200 {object} catalog_service.GetAllProductCategoryResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *Handler) GetListProductCategory(c *gin.Context) {
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

	in := &catalog_service.GetAllProductCategoryRequest{
		Limit:  int64(limit),
		Offset: int64(offset),
		Search: search,
	}

	resp, err := h.grpcClient.ProductCategoryService().GetAll(c, in)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while getting all category categorys")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /ProductCategory [POST]
// @Summary Create a category category
// @Description API for creating a category category
// @Tags category
// @Accept json
// @Produce json
// @Param ProductCategory body catalog_service.CreateProductCategory true "category category"
// @Success 200 {object} catalog_service.ProductCategory
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *Handler) CreateProductCategory(c *gin.Context) {
	var (
		req  catalog_service.CreateProductCategory
		resp *catalog_service.ProductCategory
		err  error
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	resp, err = h.grpcClient.ProductCategoryService().Create(c, &req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while creating category category")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /ProductCategory/{id} [GET]
// @Summary Get a single category category by ID
// @Description API for getting a single category category by ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "category category ID"
// @Success 200 {object} catalog_service.ProductCategory
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *Handler) GetProductCategoryByID(c *gin.Context) {
	var (
		id   = c.Param("id")
		resp *catalog_service.ProductCategory
		err  error
	)

	req := &catalog_service.ProductCategoryPrimaryKey{
		Id: id,
	}

	resp, err = h.grpcClient.ProductCategoryService().GetByID(c, req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while getting category category by id")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /ProductCategory/{id} [PUT]
// @Summary Update a category category by ID
// @Description API for updating a category category by ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "category category ID"
// @Param ProductCategory body catalog_service.UpdateProductCategory true "category category"
// @Success 200 {object} catalog_service.ProductCategory
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *Handler) UpdateProductCategory(c *gin.Context) {
	var (
		id   = c.Param("id")
		req  catalog_service.UpdateProductCategory
		resp *catalog_service.ProductCategory
		err  error
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	req.Id = id
	resp, err = h.grpcClient.ProductCategoryService().Update(c, &req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while updating category category")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /ProductCategory/{id} [DELETE]
// @Summary Delete a category category by ID
// @Description API for deleting a category category by ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "category category ID"
// @Success 200 {object} catalog_service.Empty
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *Handler) DeleteProductCategory(c *gin.Context) {
	var (
		id   = c.Param("id")
		err  error
		resp *catalog_service.Empty
	)

	req := &catalog_service.ProductCategoryPrimaryKey{
		Id: id,
	}

	resp, err = h.grpcClient.ProductCategoryService().Delete(c, req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "error while deleting category category")
		return
	}

	c.JSON(http.StatusOK, resp)
}


