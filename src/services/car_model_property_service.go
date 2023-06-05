package services

import (
	"golang-web-api/api/dtos"
	"golang-web-api/config"
	"golang-web-api/data/db"
	"golang-web-api/data/models"
	"golang-web-api/pkg/logging"

	"github.com/gin-gonic/gin"
)

type CarModelPropertyService struct {
	base *BaseService[models.CarModelProperty, dtos.CreateCarModelPropertyRequest, dtos.UpdateCarModelPropertyRequest, dtos.CarModelPropertyResponse]
}

func NewCarModelPropertyService(cfg *config.Config) *CarModelPropertyService {
	return &CarModelPropertyService{
		base: &BaseService[models.CarModelProperty, dtos.CreateCarModelPropertyRequest, dtos.UpdateCarModelPropertyRequest, dtos.CarModelPropertyResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Property.Category"},
			},
		},
	}
}

// Create
func (s *CarModelPropertyService) Create(ctx *gin.Context, req *dtos.CreateCarModelPropertyRequest) (*dtos.CarModelPropertyResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelPropertyService) Update(ctx *gin.Context, id int, req *dtos.UpdateCarModelPropertyRequest) (*dtos.CarModelPropertyResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelPropertyService) Delete(ctx *gin.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelPropertyService) GetById(ctx *gin.Context, id int) (*dtos.CarModelPropertyResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarModelPropertyService) GetByFilter(ctx *gin.Context, req *dtos.PaginationInputWithFilter) (*dtos.PagedList[dtos.CarModelPropertyResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
