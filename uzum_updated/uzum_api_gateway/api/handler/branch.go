package handler

import (
	"api/pkg/helper"
	"api/genproto/user_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router        /createBranch [post]
// @Summary       Create branch
// @Description   API for creating branch
// @Tags          branch
// @Accept        json
// @Produce       json
// @Param branch  body     user_service.CreateBranch true "branch"
// @Success 200   {object} user_service.CreateBranch
// @Failure 404   {object} models.ResponseError
// @Failure 500   {object} models.ResponseError
func (h *Handler) CreateBranch(c *gin.Context) {

	var (
		req  user_service.CreateBranch
		resp *user_service.Branch
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

	resp, err = h.grpcClient.BranchService().Create(c.Request.Context(), &req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "failed to create branch")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router         /GetListBranch [GET]
// @Summary        Get List Branch
// @Description    API for getting list branch
// @Tags           branch
// @Accept         json
// @Produce        json
// @Param		   branch query string false "branches"
// @Param		   page query int false "page"
// @Param		   limit query int false "limit"
// @Success 200    {object} user_service.GetListBranchResponse
// @Failure 404    {object} models.ResponseError
// @Failure 500    {object} models.ResponseError
func (h *Handler) GetListBranch(c *gin.Context) {
	var (
		req  user_service.GetListBranchRequest
		resp *user_service.GetListBranchResponse
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

	resp, err = h.grpcClient.BranchService().GetList(c, &req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router         /getbyidbranch/{id} [GET]
// @Summary        Get a single branch by ID
// @Description    API for getting a single branch by ID
// @Tags           branch
// @Accept         json
// @Produce        json
// @Param          id path string true "branch ID"
// @Success        200 {object} user_service.Branch
// @Failure        404 {object} models.ResponseError
// @Failure        500 {object} models.ResponseError
func (h *Handler) GetBranchByID(c *gin.Context) {
	var (
		id   = c.Param("id")
		resp *user_service.Branch
		err  error
	)

	req := &user_service.BranchPrimaryKey{
		Id: id,
	}

	resp, err = h.grpcClient.BranchService().GetByID(c, req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router          /updateBranch/{id} [PUT]
// @Summary         Update a branch by ID
// @Description     API for updating a branch by ID
// @Tags            branch
// @Accept          json
// @Produce         json
// @Param           id path string true "branch ID"
// @Param           branch body user_service.UpdateBranch true "branch"
// @Success         200 {object} user_service.Branch
// @Failure         404 {object} models.ResponseError
// @Failure         500 {object} models.ResponseError
func (h *Handler) UpdateBranch(c *gin.Context) {
	var (
		id   = c.Param("id")
		req  user_service.UpdateBranch
		resp *user_service.Branch
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

	req.Id = id
	resp, err = h.grpcClient.BranchService().Update(c, &req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router        /deleteBranch/{id} [DELETE]
// @Summary       Delete a branch by ID
// @Description   API for deleting a branch by ID
// @Tags          branch
// @Accept        json
// @Produce       json
// @Param         id path string true "branch ID"
// @Success       200 {object} user_service.Empty3
// @Failure       404 {object} models.ResponseError
// @Failure       500 {object} models.ResponseError
func (h *Handler) DeleteBranch(c *gin.Context) {
	var (
		id   = c.Param("id")
		err  error
		resp *user_service.Empty3
	)

	req := &user_service.BranchPrimaryKey{
		Id: id,
	}

	resp, err = h.grpcClient.BranchService().Delete(c, req)
	if err != nil {
		HandleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
