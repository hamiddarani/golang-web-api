package handlers

import (
	"golang-web-api/api/dtos"
	"golang-web-api/api/helper"
	"golang-web-api/config"
	"golang-web-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CountryHandler struct {
	service *services.CountryService
}

func NewCountryHandler(cfg *config.Config) *CountryHandler {
	return &CountryHandler{
		service: services.NewCountryService(cfg),
	}
}

// CreateCountry godoc
// @Summary Create a country
// @Description Create a country
// @Tags Countries
// @Accept json
// @produces json
// @Param Request body dtos.CreateUpdateCountryRequest true "Create a country"
// @Success 201 {object} helper.BaseHttpResponse{result=dtos.CountryResponse} "Country response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/countries/ [post]
// @Security AuthBearer
func (h *CountryHandler) Create(c *gin.Context) {
	req := dtos.CreateUpdateCountryRequest{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	res, err := h.service.Create(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))
}

// UpdateCountry godoc
// @Summary Update a country
// @Description Update a country
// @Tags Countries
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dtos.CreateUpdateCountryRequest true "Update a country"
// @Success 200 {object} helper.BaseHttpResponse{result=dtos.CountryResponse} "Country response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/countries/{id} [put]
// @Security AuthBearer
func (h *CountryHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	req := dtos.CreateUpdateCountryRequest{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	res, err := h.service.Update(c, id, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))
}

// DeleteCountry godoc
// @Summary Delete a country
// @Description Delete a country
// @Tags Countries
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/countries/{id} [delete]
// @Security AuthBearer
func (h *CountryHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	c.JSON(http.StatusNoContent, helper.GenerateBaseResponse(nil, true, 0))
}

// GetCountry godoc
// @Summary Get a country
// @Description Get a country
// @Tags Countries
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dtos.CountryResponse} "Country response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/countries/{id} [get]
// @Security AuthBearer
func (h *CountryHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	res, err := h.service.GetById(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))
}

// GetCountries godoc
// @Summary Get Countries
// @Description Get Countries
// @Tags Countries
// @Accept json
// @produces json
// @Param Request body dtos.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dtos.PagedList[dtos.CountryResponse]} "Country response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/countries/get-by-filter [post]
// @Security AuthBearer
func (h *CountryHandler) GetByFilter(c *gin.Context) {
	req := dtos.PaginationInputWithFilter{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	res, err := h.service.GetByFilter(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))
}

