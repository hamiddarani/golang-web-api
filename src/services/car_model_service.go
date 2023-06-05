package services

import (
	"golang-web-api/api/dtos"
	"golang-web-api/config"
	"golang-web-api/data/db"
	"golang-web-api/data/models"
	"golang-web-api/pkg/logging"

	"github.com/gin-gonic/gin"
)

type CarModelService struct {
	base *BaseService[models.CarModel, dtos.CreateCarModelRequest, dtos.UpdateCarModelRequest, dtos.CarModelResponse]
}

func NewCarModelService(cfg *config.Config) *CarModelService {
	return &CarModelService{
		base: &BaseService[models.CarModel, dtos.CreateCarModelRequest, dtos.UpdateCarModelRequest, dtos.CarModelResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Company.Country"},
				{string: "CarType"},
				{string: "Gearbox"},
				{string: "CarModelColors.Color"},
				{string: "CarModelYears.PersianYear"},
				{string: "CarModelYears.CarModelPriceHistories"},
				{string: "CarModelProperties.Property.Category"},
				{string: "CarModelImages.Image"},
				{string: "CarModelComments.User"},
			},
		},
	}
}

// Create
func (s *CarModelService) Create(ctx *gin.Context, req *dtos.CreateCarModelRequest) (*dtos.CarModelResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelService) Update(ctx *gin.Context, id int, req *dtos.UpdateCarModelRequest) (*dtos.CarModelResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelService) Delete(ctx *gin.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelService) GetById(ctx *gin.Context, id int) (*dtos.CarModelResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarModelService) GetByFilter(ctx *gin.Context, req *dtos.PaginationInputWithFilter) (*dtos.PagedList[dtos.CarModelResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
