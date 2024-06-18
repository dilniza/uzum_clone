package handler

import (
	"api/genproto/catalog_service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /category [POST]
// @Summary Create a category
// @Description API for creating a category
// @Tags category
// @Accept json
// @Produce json
// @Param category body catalog_service.CreateCategory true "category"
// @Success 200 {object} catalog_service.Category
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *Handler) CreateCategory(c *gin.Context) {
	var (
		req  catalog_service.CreateCategory
		resp *catalog_service.Category
		err  error
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	resp, err = h.grpcClient.CategoryService().Create(c, &req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}
	fmt.Println(resp, "1")
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /category/{id} [GET]
// @Summary Get a single category by ID
// @Description API for getting a single category by ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "category ID"
// @Success 200 {object} catalog_service.Category
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *Handler) GetCategoryByID(c *gin.Context) {
	var (
		id   = c.Param("id")
		resp *catalog_service.Category
		err  error
	)

	req := &catalog_service.CategoryPrimaryKey{
		Id: id,
	}

	resp, err = h.grpcClient.CategoryService().GetByID(c, req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /getallcategory [GET]
// @Summary Get All Categories
// @Description API for getting all categories
// @Tags category
// @Accept json
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Param search query string false "Search"
// @Success 200 {object} catalog_service.GetAllCategoryResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *Handler) GetAllCategory(c *gin.Context) {
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

	req := &catalog_service.GetAllCategoryRequest{
		Limit:  int64(limit),
		Offset: int64(offset),
		Search: search,
	}

	resp, err := h.grpcClient.CategoryService().GetAll(c, req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /category/{id} [PUT]
// @Summary Update a category by ID
// @Description API for updating a category by ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "category ID"
// @Param category body catalog_service.UpdateCategory true "category"
// @Success 200 {object} catalog_service.Category
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *Handler) UpdateCategory(c *gin.Context) {
	var (
		id   = c.Param("id")
		req  catalog_service.UpdateCategory
		resp *catalog_service.Category
		err  error
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	req.Id = id
	resp, err = h.grpcClient.CategoryService().Update(c, &req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /category/{id} [DELETE]
// @Summary Delete a category by ID
// @Description API for deleting a category by ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "category ID"
// @Success 200 {object} catalog_service.Empty
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *Handler) DeleteCategory(c *gin.Context) {
	var (
		id   = c.Param("id")
		err  error
		resp *catalog_service.Empty
	)

	req := &catalog_service.CategoryPrimaryKey{
		Id: id,
	}

	resp, err = h.grpcClient.CategoryService().Delete(c, req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
