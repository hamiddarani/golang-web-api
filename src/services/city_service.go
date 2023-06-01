package services

import (
	"golang-web-api/api/dtos"
	"golang-web-api/config"
	"golang-web-api/data/db"
	"golang-web-api/data/models"
	"golang-web-api/pkg/logging"

	"github.com/gin-gonic/gin"
)

type CityService struct {
	base *BaseService[models.City, dtos.CreateCityRequest, dtos.UpdateCityRequest, dtos.CityResponse]
}

func NewCityService(cfg *config.Config) *CityService {
	return &CityService{
		base: &BaseService[models.City, dtos.CreateCityRequest, dtos.UpdateCityRequest, dtos.CityResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Country"},
			},
		},
	}
}

func (s *CityService) Create(ctx *gin.Context, req *dtos.CreateCityRequest) (*dtos.CityResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *CityService) Update(ctx *gin.Context, id int, req *dtos.UpdateCityRequest) (*dtos.CityResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CityService) Delete(ctx *gin.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CityService) GetById(ctx *gin.Context, id int) (*dtos.CityResponse, error) {
	return s.base.GetById(ctx, id)
}

func (s *CityService) GetByFilter(ctx *gin.Context, req *dtos.PaginationInputWithFilter) (*dtos.PagedList[dtos.CityResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
