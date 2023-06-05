package handlers

import (
	"golang-web-api/config"
	"golang-web-api/services"

	"github.com/gin-gonic/gin"
)

type CarModelCommentHandler struct {
	service *services.CarModelCommentService
}

func NewCarModelCommentHandler(cfg *config.Config) *CarModelCommentHandler {
	return &CarModelCommentHandler{
		service: services.NewCarModelCommentService(cfg),
	}
}

// CreateCarModelComment godoc
// @Summary Create a CarModelComment
// @Description Create a CarModelComment
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param Request body dtos.CreateCarModelCommentRequest true "Create a CarModelComment"
// @Success 201 {object} helper.BaseHttpResponse{result=dtos.CarModelCommentResponse} "CarModelComment response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-comments/ [post]
// @Security AuthBearer
func (h *CarModelCommentHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// UpdateCarModelComment godoc
// @Summary Update a CarModelComment
// @Description Update a CarModelComment
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dtos.UpdateCarModelCommentRequest true "Update a CarModelComment"
// @Success 200 {object} helper.BaseHttpResponse{result=dtos.CarModelCommentResponse} "CarModelComment response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-comments/{id} [put]
// @Security AuthBearer
func (h *CarModelCommentHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// DeleteCarModelComment godoc
// @Summary Delete a CarModelComment
// @Description Delete a CarModelComment
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-comments/{id} [delete]
// @Security AuthBearer
func (h *CarModelCommentHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetCarModelComment godoc
// @Summary Get a CarModelComment
// @Description Get a CarModelComment
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dtos.CarModelCommentResponse} "CarModelComment response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-comments/{id} [get]
// @Security AuthBearer
func (h *CarModelCommentHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetCarModelComments godoc
// @Summary Get CarModelComments
// @Description Get CarModelComments
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param Request body dtos.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dtos.PagedList[dtos.CarModelCommentResponse]} "CarModelComment response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-comments/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelCommentHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
