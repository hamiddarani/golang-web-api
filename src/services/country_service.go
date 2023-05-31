package services

import (
	"golang-web-api/api/dtos"
	"golang-web-api/config"
	"golang-web-api/data/db"
	"golang-web-api/data/models"
	"golang-web-api/pkg/logging"

	"github.com/gin-gonic/gin"
)

type CountryService struct {
	base *BaseService[models.Country, dtos.CreateUpdateCountryRequest, dtos.CreateUpdateCountryRequest, dtos.CountryResponse]
}

func NewCountryService(cfg *config.Config) *CountryService {
	return &CountryService{
		base: &BaseService[models.Country, dtos.CreateUpdateCountryRequest, dtos.CreateUpdateCountryRequest, dtos.CountryResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

func (s *CountryService) Create(ctx *gin.Context, req *dtos.CreateUpdateCountryRequest) (*dtos.CountryResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *CountryService) Update(ctx *gin.Context, id int, req *dtos.CreateUpdateCountryRequest) (*dtos.CountryResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CountryService) Delete(ctx *gin.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CountryService) GetById(ctx *gin.Context, id int) (*dtos.CountryResponse, error) {
	return s.base.GetById(ctx, id)
}
