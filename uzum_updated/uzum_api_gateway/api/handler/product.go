package handler

import (
	"api/genproto/catalog_service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /product [POST]
// @Summary Create a product
// @Description API for creating a product
// @Tags product
// @Accept json
// @Produce json
// @Param product body catalog_service.CreateProduct true "product"
// @Success 200 {object} catalog_service.Product
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *Handler) CreateProduct(c *gin.Context) {
	var (
		req  catalog_service.CreateProduct
		resp *catalog_service.Product
		err  error
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	resp, err = h.grpcClient.ProductService().Create(c, &req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}
	fmt.Println(resp, "1")
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /product/{id} [GET]
// @Summary Get a single product by ID
// @Description API for getting a single product by ID
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "product ID"
// @Success 200 {object} catalog_service.Product
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *Handler) GetProductByID(c *gin.Context) {
	var (
		id   = c.Param("id")
		resp *catalog_service.Product
		err  error
	)

	req := &catalog_service.ProductPrimaryKey{
		Id: id,
	}

	resp, err = h.grpcClient.ProductService().GetByID(c, req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /getallproduct [GET]
// @Summary Get All Products
// @Description API for getting all products
// @Tags product
// @Accept json
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Param search query string false "Search"
// @Success 200 {object} catalog_service.GetAllProductResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *Handler) GetAllProduct(c *gin.Context) {
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

	req := &catalog_service.GetAllProductRequest{
		Limit:  int64(limit),
		Offset: int64(offset),
		Search: search,
	}

	resp, err := h.grpcClient.ProductService().GetAll(c, req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /product/{id} [PUT]
// @Summary Update a product by ID
// @Description API for updating a product by ID
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "product ID"
// @Param product body catalog_service.UpdateProduct true "product"
// @Success 200 {object} catalog_service.Product
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *Handler) UpdateProduct(c *gin.Context) {
	var (
		id   = c.Param("id")
		req  catalog_service.UpdateProduct
		resp *catalog_service.Product
		err  error
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	req.Id = id
	resp, err = h.grpcClient.ProductService().Update(c, &req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /product/{id} [DELETE]
// @Summary Delete a product by ID
// @Description API for deleting a product by ID
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "product ID"
// @Success 200 {object} catalog_service.Empty
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *Handler) DeleteProduct(c *gin.Context) {
	var (
		id   = c.Param("id")
		err  error
		resp *catalog_service.Empty
	)

	req := &catalog_service.ProductPrimaryKey{
		Id: id,
	}

	resp, err = h.grpcClient.ProductService().Delete(c, req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
