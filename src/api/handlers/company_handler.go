package handlers

import (
	"golang-web-api/config"
	"golang-web-api/services"

	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	service *services.CompanyService
}

func NewCompanyHandler(cfg *config.Config) *CompanyHandler {
	return &CompanyHandler{
		service: services.NewCompanyService(cfg),
	}
}

// CreateCompany godoc
// @Summary Create a Company
// @Description Create a Company
// @Tags Companies
// @Accept json
// @produces json
// @Param Request body dtos.CreateCompanyRequest true "Create a Company"
// @Success 201 {object} helper.BaseHttpResponse{result=dtos.CompanyResponse} "Company response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/companies/ [post]
// @Security AuthBearer
func (h *CompanyHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// UpdateCompany godoc
// @Summary Update a Company
// @Description Update a Company
// @Tags Companies
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dtos.UpdateCompanyRequest true "Update a Company"
// @Success 200 {object} helper.BaseHttpResponse{result=dtos.CompanyResponse} "Company response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/companies/{id} [put]
// @Security AuthBearer
func (h *CompanyHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// DeleteCompany godoc
// @Summary Delete a Company
// @Description Delete a Company
// @Tags Companies
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/companies/{id} [delete]
// @Security AuthBearer
func (h *CompanyHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetCompany godoc
// @Summary Get a Company
// @Description Get a Company
// @Tags Companies
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dtos.CompanyResponse} "Company response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/companies/{id} [get]
// @Security AuthBearer
func (h *CompanyHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetCompanies godoc
// @Summary Get Companies
// @Description Get Companies
// @Tags Companies
// @Accept json
// @produces json
// @Param Request body dtos.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dtos.PagedList[dtos.CompanyResponse]} "Company response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/companies/get-by-filter [post]
// @Security AuthBearer
func (h *CompanyHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
