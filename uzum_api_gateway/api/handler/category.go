package handler

import (
	"api/genproto/category_service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /getlistCategory [POST]
// @Summary Get List Categories
// @Description API for getting list categories
// @Tags category
// @Accept json
// @Produce json
// @Param category body category_service.GetListCategoryRequest true "category"
// @Success 200 {object} category_service.GetListCategoryResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) GetListCategory(c *gin.Context) {
	var (
		req  category_service.GetListCategoryRequest
		resp *category_service.GetListCategoryResponse
		err  error
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	resp, err = h.grpcClient.CategoryService().GetAll(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /category [POST]
// @Summary Create a category
// @Description API for creating a category
// @Tags category
// @Accept json
// @Produce json
// @Param category body category_service.CreateCategory true "category"
// @Success 200 {object} category_service.Category
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) CreateCategory(c *gin.Context) {
	var (
		req  category_service.CreateCategory
		resp *category_service.Category
		err  error
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}
	fmt.Println(req, "1")
	resp, err = h.grpcClient.CategoryService().Create(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
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
// @Success 200 {object} category_service.Category
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) GetCategoryByID(c *gin.Context) {
	var (
		id   = c.Param("id")
		resp *category_service.Category
		err  error
	)

	req := &category_service.CategoryPrimaryKey{
		Id: id,
	}

	resp, err = h.grpcClient.CategoryService().GetByID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
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
// @Param category body category_service.UpdateCategory true "category"
// @Success 200 {object} category_service.Category
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) UpdateCategory(c *gin.Context) {
	var (
		id   = c.Param("id")
		req  category_service.UpdateCategory
		resp *category_service.Category
		err  error
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	req.Id = id
	resp, err = h.grpcClient.CategoryService().Update(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
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
// @Success 200 {object} category_service.Empty
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) DeleteCategory(c *gin.Context) {
	var (
		id   = c.Param("id")
		err  error
		resp *category_service.Empty
	)

	req := &category_service.CategoryPrimaryKey{
		Id: id,
	}

	resp, err = h.grpcClient.CategoryService().Delete(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
