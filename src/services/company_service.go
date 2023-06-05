package services

import (
	"golang-web-api/api/dtos"
	"golang-web-api/config"
	"golang-web-api/data/db"
	"golang-web-api/data/models"
	"golang-web-api/pkg/logging"

	"github.com/gin-gonic/gin"
)

type CompanyService struct {
	base *BaseService[models.Company, dtos.CreateCompanyRequest, dtos.UpdateCompanyRequest, dtos.CompanyResponse]
}

func NewCompanyService(cfg *config.Config) *CompanyService {
	return &CompanyService{
		base: &BaseService[models.Company, dtos.CreateCompanyRequest, dtos.UpdateCompanyRequest, dtos.CompanyResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Country"},
			},
		},
	}
}

// Create
func (s *CompanyService) Create(ctx *gin.Context, req *dtos.CreateCompanyRequest) (*dtos.CompanyResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CompanyService) Update(ctx *gin.Context, id int, req *dtos.UpdateCompanyRequest) (*dtos.CompanyResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CompanyService) Delete(ctx *gin.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CompanyService) GetById(ctx *gin.Context, id int) (*dtos.CompanyResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CompanyService) GetByFilter(ctx *gin.Context, req *dtos.PaginationInputWithFilter) (*dtos.PagedList[dtos.CompanyResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
